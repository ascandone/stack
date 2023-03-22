/*
Provides `Machine`, which executes programs and outputs postings.
1: Create New Machine
2: Set Variables (with `core.Value`s or JSON)
3: Resolve Resources (answer requests on channel)
4: Resolve Balances (answer requests on channel)
6: Execute
*/
package vm

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/formancehq/ledger/pkg/core"
	"github.com/formancehq/ledger/pkg/machine/vm/program"
	"github.com/logrusorgru/aurora"
	"github.com/pkg/errors"
)

const (
	EXIT_OK = byte(iota + 1)
	EXIT_FAIL
	EXIT_FAIL_INVALID
	EXIT_FAIL_INSUFFICIENT_FUNDS
)

type Machine struct {
	P                   uint
	Program             program.Program
	Vars                map[string]core.Value
	UnresolvedResources []program.Resource
	Resources           []core.Value // Constants and Variables
	resolveCalled       bool
	Balances            map[core.AccountAddress]map[core.Asset]*core.MonetaryInt // keeps track of balances throughout execution
	setBalanceCalled    bool
	Stack               []core.Value
	Postings            []Posting                                     // accumulates postings throughout execution
	TxMeta              map[string]core.Value                         // accumulates transaction meta throughout execution
	AccountsMeta        map[core.AccountAddress]map[string]core.Value // accumulates accounts meta throughout execution
	Printer             func(chan core.Value)
	printChan           chan core.Value
	Debug               bool
}

type Posting struct {
	Source      string            `json:"source"`
	Destination string            `json:"destination"`
	Amount      *core.MonetaryInt `json:"amount"`
	Asset       string            `json:"asset"`
}

type Metadata map[string]any

func NewMachine(p program.Program) *Machine {
	printChan := make(chan core.Value)

	m := Machine{
		Program:             p,
		UnresolvedResources: p.Resources,
		Resources:           make([]core.Value, 0),
		printChan:           printChan,
		Printer:             StdOutPrinter,
		Postings:            make([]Posting, 0),
		TxMeta:              map[string]core.Value{},
		AccountsMeta:        map[core.AccountAddress]map[string]core.Value{},
	}

	return &m
}

func StdOutPrinter(c chan core.Value) {
	for v := range c {
		fmt.Println("OUT:", v)
	}
}

func (m *Machine) GetTxMetaJSON() Metadata {
	meta := make(Metadata)
	for k, v := range m.TxMeta {
		valJSON, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		v, err := json.Marshal(core.ValueJSON{
			Type:  v.GetType().String(),
			Value: valJSON,
		})
		if err != nil {
			panic(err)
		}
		meta[k] = v
	}
	return meta
}

func (m *Machine) GetAccountsMetaJSON() Metadata {
	res := Metadata{}
	for account, meta := range m.AccountsMeta {
		for k, v := range meta {
			if _, ok := res[account.String()]; !ok {
				res[account.String()] = map[string][]byte{}
			}
			valJSON, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}
			v, err := json.Marshal(core.ValueJSON{
				Type:  v.GetType().String(),
				Value: valJSON,
			})
			if err != nil {
				panic(err)
			}
			res[account.String()].(map[string][]byte)[k] = v
		}
	}

	return res
}

func (m *Machine) getResource(addr core.Address) (*core.Value, bool) {
	a := int(addr)
	if a >= len(m.Resources) {
		return nil, false
	}
	return &m.Resources[a], true
}

func (m *Machine) withdrawAll(account core.AccountAddress, asset core.Asset, overdraft *core.MonetaryInt) (*core.Funding, error) {
	if accBalances, ok := m.Balances[account]; ok {
		if balance, ok := accBalances[asset]; ok {
			amountTaken := core.NewMonetaryInt(0)
			if balance.Add(overdraft).Gt(core.NewMonetaryInt(0)) {
				amountTaken = balance.Add(overdraft)
				accBalances[asset] = overdraft.Neg()
			}

			return &core.Funding{
				Asset: asset,
				Parts: []core.FundingPart{{
					Account: account,
					Amount:  amountTaken,
				}},
			}, nil
		}
	}
	return nil, fmt.Errorf("missing %v balance from %v", asset, account)
}

func (m *Machine) withdrawAlways(account core.AccountAddress, mon core.Monetary) (*core.Funding, error) {
	if accBalance, ok := m.Balances[account]; ok {
		if balance, ok := accBalance[mon.Asset]; ok {
			accBalance[mon.Asset] = balance.Sub(mon.Amount)
			return &core.Funding{
				Asset: mon.Asset,
				Parts: []core.FundingPart{{
					Account: account,
					Amount:  mon.Amount,
				}},
			}, nil
		}
	}
	return nil, fmt.Errorf("missing %v balance from %v", mon.Asset, account)
}

func (m *Machine) credit(account core.AccountAddress, funding core.Funding) {
	if account == "world" {
		return
	}
	if accBalance, ok := m.Balances[account]; ok {
		if _, ok := accBalance[funding.Asset]; ok {
			for _, part := range funding.Parts {
				balance := accBalance[funding.Asset]
				accBalance[funding.Asset] = balance.Add(part.Amount)
			}
		}
	}
}

func (m *Machine) repay(funding core.Funding) {
	for _, part := range funding.Parts {
		if part.Account == "world" {
			continue
		}
		balance := m.Balances[part.Account][funding.Asset]
		m.Balances[part.Account][funding.Asset] = balance.Add(part.Amount)
	}
}

func (m *Machine) tick() (bool, byte) {
	op := m.Program.Instructions[m.P]

	if m.Debug {
		fmt.Println("STATE ---------------------------------------------------------------------")
		fmt.Printf("    %v\n", aurora.Blue(m.Stack))
		fmt.Printf("    %v\n", aurora.Cyan(m.Balances))
		fmt.Printf("    %v\n", program.OpcodeName(op))
	}

	switch op {
	case program.OP_APUSH:
		bytes := m.Program.Instructions[m.P+1 : m.P+3]
		v, ok := m.getResource(core.Address(binary.LittleEndian.Uint16(bytes)))
		if !ok {
			return true, EXIT_FAIL
		}
		m.Stack = append(m.Stack, *v)
		m.P += 2

	case program.OP_BUMP:
		n := big.Int(*pop[core.Number](m))
		idx := len(m.Stack) - int(n.Uint64()) - 1
		v := m.Stack[idx]
		m.Stack = append(m.Stack[:idx], m.Stack[idx+1:]...)
		m.Stack = append(m.Stack, v)

	case program.OP_DELETE:
		n := m.popValue()
		if n.GetType() == core.TypeFunding {
			return true, EXIT_FAIL_INVALID
		}

	case program.OP_IADD:
		b := pop[core.Number](m)
		a := pop[core.Number](m)
		m.pushValue(a.Add(b))

	case program.OP_ISUB:
		b := pop[core.Number](m)
		a := pop[core.Number](m)
		m.pushValue(a.Sub(b))

	case program.OP_PRINT:
		a := m.popValue()
		m.printChan <- a

	case program.OP_FAIL:
		return true, EXIT_FAIL

	case program.OP_ASSET:
		v := m.popValue()
		switch v := v.(type) {
		case core.Asset:
			m.pushValue(v)
		case core.Monetary:
			m.pushValue(v.Asset)
		case core.Funding:
			m.pushValue(v.Asset)
		default:
			return true, EXIT_FAIL_INVALID
		}

	case program.OP_MONETARY_NEW:
		amount := pop[core.Number](m)
		asset := pop[core.Asset](m)
		m.pushValue(core.Monetary{
			Asset:  asset,
			Amount: amount,
		})

	case program.OP_MONETARY_ADD:
		b := pop[core.Monetary](m)
		a := pop[core.Monetary](m)
		if a.Asset != b.Asset {
			return true, EXIT_FAIL_INVALID
		}
		m.pushValue(core.Monetary{
			Asset:  a.Asset,
			Amount: a.Amount.Add(b.Amount),
		})

	case program.OP_MAKE_ALLOTMENT:
		n := pop[core.Number](m)
		portions := make([]core.Portion, n.Uint64())
		for i := uint64(0); i < n.Uint64(); i++ {
			p := pop[core.Portion](m)
			portions[i] = p
		}
		allotment, err := core.NewAllotment(portions)
		if err != nil {
			return true, EXIT_FAIL_INVALID
		}
		m.pushValue(*allotment)

	case program.OP_TAKE_ALL:
		overdraft := pop[core.Monetary](m)
		account := pop[core.AccountAddress](m)
		funding, err := m.withdrawAll(account, overdraft.Asset, overdraft.Amount)
		if err != nil {
			return true, EXIT_FAIL_INVALID
		}
		m.pushValue(*funding)

	case program.OP_TAKE_ALWAYS:
		mon := pop[core.Monetary](m)
		account := pop[core.AccountAddress](m)
		funding, err := m.withdrawAlways(account, mon)
		if err != nil {
			return true, EXIT_FAIL_INVALID
		}
		m.pushValue(*funding)

	case program.OP_TAKE:
		mon := pop[core.Monetary](m)
		funding := pop[core.Funding](m)
		if funding.Asset != mon.Asset {
			return true, EXIT_FAIL_INVALID
		}
		result, remainder, err := funding.Take(mon.Amount)
		if err != nil {
			return true, EXIT_FAIL_INSUFFICIENT_FUNDS
		}
		m.pushValue(remainder)
		m.pushValue(result)

	case program.OP_TAKE_MAX:
		mon := pop[core.Monetary](m)
		funding := pop[core.Funding](m)
		if funding.Asset != mon.Asset {
			return true, EXIT_FAIL_INVALID
		}
		missing := core.NewMonetaryInt(0)
		total := funding.Total()
		if mon.Amount.Gt(total) {
			missing = mon.Amount.Sub(total)
		}
		result, remainder := funding.TakeMax(mon.Amount)
		m.pushValue(core.Monetary{
			Asset:  mon.Asset,
			Amount: missing,
		})
		m.pushValue(remainder)
		m.pushValue(result)

	case program.OP_FUNDING_ASSEMBLE:
		num := pop[core.Number](m)
		n := int(num.Uint64())
		if n == 0 {
			return true, EXIT_FAIL_INVALID
		}
		first := pop[core.Funding](m)
		result := core.Funding{
			Asset: first.Asset,
		}
		fundings_rev := make([]core.Funding, n)
		fundings_rev[0] = first
		for i := 1; i < n; i++ {
			f := pop[core.Funding](m)
			if f.Asset != result.Asset {
				return true, EXIT_FAIL_INVALID
			}
			fundings_rev[i] = f
		}
		for i := 0; i < n; i++ {
			res, err := result.Concat(fundings_rev[n-1-i])
			if err != nil {
				return true, EXIT_FAIL_INVALID
			}
			result = res
		}
		m.pushValue(result)

	case program.OP_FUNDING_SUM:
		funding := pop[core.Funding](m)
		sum := funding.Total()
		m.pushValue(funding)
		m.pushValue(core.Monetary{
			Asset:  funding.Asset,
			Amount: sum,
		})

	case program.OP_FUNDING_REVERSE:
		funding := pop[core.Funding](m)
		result := funding.Reverse()
		m.pushValue(result)

	case program.OP_ALLOC:
		allotment := pop[core.Allotment](m)
		monetary := pop[core.Monetary](m)
		total := monetary.Amount
		parts := allotment.Allocate(total)
		for i := len(parts) - 1; i >= 0; i-- {
			m.pushValue(core.Monetary{
				Asset:  monetary.Asset,
				Amount: parts[i],
			})
		}

	case program.OP_REPAY:
		m.repay(pop[core.Funding](m))

	case program.OP_SEND:
		dest := pop[core.AccountAddress](m)
		funding := pop[core.Funding](m)
		m.credit(dest, funding)
		for _, part := range funding.Parts {
			src := part.Account
			amt := part.Amount
			if amt.Eq(core.NewMonetaryInt(0)) {
				continue
			}
			m.Postings = append(m.Postings, Posting{
				Source:      string(src),
				Destination: string(dest),
				Asset:       string(funding.Asset),
				Amount:      amt,
			})
		}

	case program.OP_TX_META:
		k := pop[core.String](m)
		v := m.popValue()
		m.TxMeta[string(k)] = v

	case program.OP_ACCOUNT_META:
		a := pop[core.AccountAddress](m)
		k := pop[core.String](m)
		v := m.popValue()
		if m.AccountsMeta[a] == nil {
			m.AccountsMeta[a] = map[string]core.Value{}
		}
		m.AccountsMeta[a][string(k)] = v

	default:
		return true, EXIT_FAIL_INVALID
	}

	m.P += 1

	if int(m.P) >= len(m.Program.Instructions) {
		return true, EXIT_OK
	}

	return false, 0
}

func (m *Machine) Execute() (byte, error) {
	go m.Printer(m.printChan)
	defer close(m.printChan)

	if len(m.Resources) != len(m.UnresolvedResources) {
		return 0, errors.New("resources haven't been initialized")
	} else if m.Balances == nil {
		return 0, errors.New("balances haven't been initialized")
	}

	for {
		finished, exitCode := m.tick()
		if finished {
			if exitCode == EXIT_OK && len(m.Stack) != 0 {
				return EXIT_FAIL_INVALID, nil
			} else {
				return exitCode, nil
			}
		}
	}
}

type BalanceRequest struct {
	Account  string
	Asset    string
	Response chan *core.MonetaryInt
	Error    error
}

func (m *Machine) ResolveBalances(ctx context.Context, store Store) error {
	if len(m.Resources) != len(m.UnresolvedResources) {
		return errors.New("tried to resolve balances before resources")
	}
	if m.setBalanceCalled {
		return errors.New("tried to call ResolveBalances twice")
	}
	m.setBalanceCalled = true
	m.Balances = make(map[core.AccountAddress]map[core.Asset]*core.MonetaryInt)
	// for every account that we need balances of, check if it's there
	for addr, neededAssets := range m.Program.NeededBalances {
		account, ok := m.getResource(addr)
		if !ok {
			return errors.New("invalid program (resolve balances: invalid address of account)")
		}
		accountAddress := (*account).(core.AccountAddress)
		m.Balances[accountAddress] = make(map[core.Asset]*core.MonetaryInt)
		// for every asset, send request
		for addr := range neededAssets {
			mon, ok := m.getResource(addr)
			if !ok {
				return errors.New("invalid program (resolve balances: invalid address of monetary)")
			}

			asset := (*mon).(core.HasAsset).GetAsset()
			if string(accountAddress) == "world" {
				m.Balances[accountAddress][asset] = core.NewMonetaryInt(0)
				continue
			}

			account, err := store.GetAccountWithVolumes(ctx, string(accountAddress))
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("could not get account %q", account))
			}
			amt := account.Volumes[string(asset)].Balance().OrZero()

			m.Balances[accountAddress][asset] = amt
		}
	}
	return nil
}

type ResourceRequest struct {
	Account  string
	Key      string
	Asset    string
	Response chan core.Value
	Error    error
}

func (m *Machine) ResolveResources(ctx context.Context, store Store) error {
	if m.resolveCalled {
		return errors.New("tried to call ResolveResources twice")
	}

	m.resolveCalled = true
	for len(m.Resources) != len(m.UnresolvedResources) {
		idx := len(m.Resources)
		res := m.UnresolvedResources[idx]
		var val core.Value
		switch res := res.(type) {
		case program.Constant:
			val = res.Inner
		case program.Variable:
			var ok bool
			val, ok = m.Vars[res.Name]
			if !ok {
				return fmt.Errorf("missing variable '%s'", res.Name)
			}
		case program.VariableAccountMetadata:
			sourceAccount, ok := m.getResource(res.Account)
			if !ok {
				return fmt.Errorf(
					"variable '%s': tried to request metadata of an account which has not yet been solved",
					res.Name)
			}
			if (*sourceAccount).GetType() != core.TypeAccount {
				return fmt.Errorf(
					"variable '%s': tried to request metadata on wrong entity: %v instead of account",
					res.Name, (*sourceAccount).GetType())
			}

			address := string((*sourceAccount).(core.AccountAddress))
			account, err := store.GetAccountWithVolumes(ctx, address)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("could not get account %s", address))
			}

			entry, ok := account.Metadata[res.Key]
			if !ok {
				return NewScriptError(ScriptErrorCompilationFailed,
					fmt.Sprintf("missing key %v in metadata for account %s", res.Key, address))
			}

			data, err := json.Marshal(entry)
			if err != nil {
				return errors.Wrap(err, "marshaling metadata")
			}
			val, err = core.NewValueFromTypedJSON(data)
			if err != nil {
				return NewScriptError(ScriptErrorCompilationFailed,
					errors.Wrap(err, fmt.Sprintf("invalid format for metadata at key %v for account %s",
						res.Key, address)).Error())
			}
		case program.VariableAccountBalance:
			sourceAccount, ok := m.getResource(res.Account)
			if !ok {
				return fmt.Errorf(
					"variable '%s': tried to request metadata of an account which has not yet been solved",
					res.Name)
			}
			if (*sourceAccount).GetType() != core.TypeAccount {
				return fmt.Errorf(
					"variable '%s': tried to request balance of an account which has not yet been solved",
					res.Name)
			}

			address := string((*sourceAccount).(core.AccountAddress))
			account, err := store.GetAccountWithVolumes(ctx, address)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("could not get account %s", address))
			}

			ass, ok := m.getResource(res.Asset)
			if !ok {
				return fmt.Errorf(
					"variable '%s': tried to request balance of an account for an asset which has not yet been solved",
					res.Name)
			}
			if (*ass).GetType() != core.TypeAsset {
				return fmt.Errorf(
					"variable '%s': tried to request account balance on wrong entity: %v instead of asset",
					res.Name, (*ass).GetType())
			}

			amt := account.Volumes[string((*ass).(core.Asset))].Balance().OrZero()
			if amt.Ltz() {
				return fmt.Errorf(
					"variable '%s': tried to request the balance of account %s for asset %s: received %s: monetary amounts must be non-negative",
					res.Name, address, string((*ass).(core.Asset)), amt)
			}

			val = core.Monetary{
				Asset:  (*ass).(core.Asset),
				Amount: amt,
			}
		case program.Monetary:
			ass, ok := m.getResource(res.Asset)
			if !ok {
				return fmt.Errorf("tried to resolve an asset which has not yet been solved")
			}
			if (*ass).GetType() != core.TypeAsset {
				return fmt.Errorf(
					"tried to resolve an asset on wrong type '%v'",
					(*ass).GetType())
			}
			asset := (*ass).(core.Asset)
			val = core.Monetary{
				Asset:  asset,
				Amount: res.Amount,
			}
		default:
			panic(fmt.Errorf("type %T not implemented", res))
		}
		m.Resources = append(m.Resources, val)
	}
	return nil
}

func (m *Machine) SetVars(vars map[string]core.Value) error {
	v, err := m.Program.ParseVariables(vars)
	if err != nil {
		return err
	}
	m.Vars = v
	return nil
}

func (m *Machine) SetVarsFromJSON(vars map[string]json.RawMessage) error {
	v, err := m.Program.ParseVariablesJSON(vars)
	if err != nil {
		return err
	}
	m.Vars = v
	return nil
}
