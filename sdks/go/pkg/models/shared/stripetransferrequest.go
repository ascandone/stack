// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"math/big"
)

type StripeTransferRequest struct {
	Amount      *big.Int `json:"amount,omitempty"`
	Asset       *string  `json:"asset,omitempty"`
	Destination *string  `json:"destination,omitempty"`
	// A set of key/value pairs that you can attach to a transfer object.
	// It can be useful for storing additional information about the transfer in a structured format.
	//
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}
