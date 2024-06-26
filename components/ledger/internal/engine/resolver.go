package engine

import (
	"context"
	"sync"

	"github.com/pkg/errors"

	"github.com/formancehq/ledger/internal/storage/ledgerstore"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/formancehq/ledger/internal/engine/command"
	"github.com/formancehq/ledger/internal/opentelemetry/metrics"
	"github.com/formancehq/ledger/internal/storage/driver"
	"github.com/formancehq/stack/libs/go-libs/logging"
	"github.com/sirupsen/logrus"
)

type option func(r *Resolver)

func WithMessagePublisher(publisher message.Publisher) option {
	return func(r *Resolver) {
		r.publisher = publisher
	}
}

func WithMetricsRegistry(registry metrics.GlobalRegistry) option {
	return func(r *Resolver) {
		r.metricsRegistry = registry
	}
}

func WithCompiler(compiler *command.Compiler) option {
	return func(r *Resolver) {
		r.compiler = compiler
	}
}

func WithLogger(logger logging.Logger) option {
	return func(r *Resolver) {
		r.logger = logger
	}
}

func WithLedgerConfig(config LedgerConfig) option {
	return func(r *Resolver) {
		r.ledgerConfig = config
	}
}

var defaultOptions = []option{
	WithMetricsRegistry(metrics.NewNoOpRegistry()),
	WithCompiler(command.NewCompiler(1024)),
	WithLogger(logging.NewLogrus(logrus.New())),
}

type Resolver struct {
	storageDriver   *driver.Driver
	lock            sync.RWMutex
	metricsRegistry metrics.GlobalRegistry
	//TODO(gfyrag): add a routine to clean old ledger
	ledgers      map[string]*Ledger
	ledgerConfig LedgerConfig
	compiler     *command.Compiler
	logger       logging.Logger
	publisher    message.Publisher
}

func NewResolver(storageDriver *driver.Driver, options ...option) *Resolver {
	r := &Resolver{
		storageDriver: storageDriver,
		ledgers:       map[string]*Ledger{},
		ledgerConfig:  defaultLedgerConfig,
	}
	for _, opt := range append(defaultOptions, options...) {
		opt(r)
	}

	return r
}

func (r *Resolver) startLedgerUsingStore(ctx context.Context, name string, store *ledgerstore.Store) *Ledger {
	ledger := New(store, r.publisher, r.compiler, r.ledgerConfig)
	ledger.Start(logging.ContextWithLogger(context.Background(), r.logger))
	r.ledgers[name] = ledger
	r.metricsRegistry.ActiveLedgers().Add(ctx, +1)

	return ledger
}

func (r *Resolver) GetLedger(ctx context.Context, name string) (*Ledger, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}
	r.lock.RLock()
	ledger, ok := r.ledgers[name]
	r.lock.RUnlock()

	if !ok {
		r.lock.Lock()
		defer r.lock.Unlock()

		ledger, ok = r.ledgers[name]
		if ok {
			return ledger, nil
		}

		store, err := r.storageDriver.GetLedgerStore(ctx, name)
		if err != nil {
			return nil, err
		}

		ledger = r.startLedgerUsingStore(ctx, name, store)
	}

	return ledger, nil
}

func (r *Resolver) CreateLedger(ctx context.Context, name string, configuration driver.LedgerConfiguration) (*Ledger, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	store, err := r.storageDriver.CreateLedgerStore(ctx, name, configuration)
	if err != nil {
		return nil, err
	}

	return r.startLedgerUsingStore(ctx, name, store), nil
}

func (r *Resolver) CloseLedgers(ctx context.Context) error {
	r.logger.Info("Close all ledgers")
	defer func() {
		r.logger.Info("All ledgers closed")
	}()
	for name, ledger := range r.ledgers {
		r.logger.Infof("Close ledger %s", name)
		ledger.Close(logging.ContextWithLogger(ctx, r.logger.WithField("ledger", name)))
		delete(r.ledgers, name)
	}

	return nil
}
