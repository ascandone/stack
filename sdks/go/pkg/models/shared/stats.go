// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Stats struct {
	Accounts     int64 `json:"accounts"`
	Transactions int64 `json:"transactions"`
}

func (o *Stats) GetAccounts() int64 {
	if o == nil {
		return 0
	}
	return o.Accounts
}

func (o *Stats) GetTransactions() int64 {
	if o == nil {
		return 0
	}
	return o.Transactions
}