package monitor

import (
	"context"

	"github.com/formancehq/ledger/pkg/core"
)

type Monitor interface {
	CommittedTransactions(ctx context.Context, ledger string, res ...core.ExpandedTransaction)
	SavedMetadata(ctx context.Context, ledger, targetType, id string, metadata core.Metadata)
	RevertedTransaction(ctx context.Context, ledger string, reverted, revert *core.ExpandedTransaction)
}

type noOpMonitor struct{}

func (n noOpMonitor) CommittedTransactions(ctx context.Context, s string, res ...core.ExpandedTransaction) {
}
func (n noOpMonitor) SavedMetadata(ctx context.Context, ledger string, targetType string, id string, metadata core.Metadata) {
}
func (n noOpMonitor) RevertedTransaction(ctx context.Context, ledger string, reverted, revert *core.ExpandedTransaction) {
}

var _ Monitor = &noOpMonitor{}

func NewNoOpMonitor() *noOpMonitor {
	return &noOpMonitor{}
}