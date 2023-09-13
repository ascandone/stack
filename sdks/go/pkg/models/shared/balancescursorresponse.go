// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"math/big"
)

type BalancesCursorResponseCursor struct {
	Data     []map[string]map[string]*big.Int `json:"data"`
	HasMore  bool                             `json:"hasMore"`
	Next     *string                          `json:"next,omitempty"`
	PageSize int64                            `json:"pageSize"`
	Previous *string                          `json:"previous,omitempty"`
}

// BalancesCursorResponse - OK
type BalancesCursorResponse struct {
	Cursor BalancesCursorResponseCursor `json:"cursor"`
}
