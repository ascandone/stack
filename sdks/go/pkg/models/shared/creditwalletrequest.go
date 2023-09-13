// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreditWalletRequest struct {
	Amount Monetary `json:"amount"`
	// The balance to credit
	Balance *string `json:"balance,omitempty"`
	// Metadata associated with the wallet.
	Metadata  map[string]string `json:"metadata"`
	Reference *string           `json:"reference,omitempty"`
	Sources   []Subject         `json:"sources"`
}
