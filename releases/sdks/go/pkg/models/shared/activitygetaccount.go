// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ActivityGetAccount struct {
	ID     string `json:"id"`
	Ledger string `json:"ledger"`
}

func (o *ActivityGetAccount) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ActivityGetAccount) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}