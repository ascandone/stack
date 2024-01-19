// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
)

// WebhooksErrorResponse - Error
type WebhooksErrorResponse struct {
	Details      *string                   `json:"details,omitempty"`
	ErrorCode    shared.WebhooksErrorsEnum `json:"errorCode"`
	ErrorMessage string                    `json:"errorMessage"`
}

var _ error = &WebhooksErrorResponse{}

func (e *WebhooksErrorResponse) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
