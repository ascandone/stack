// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type GetTransferInitiationRequest struct {
	// The transfer ID.
	TransferID string `pathParam:"style=simple,explode=false,name=transferId"`
}

func (o *GetTransferInitiationRequest) GetTransferID() string {
	if o == nil {
		return ""
	}
	return o.TransferID
}

type GetTransferInitiationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TransferInitiationResponse *shared.TransferInitiationResponse
}

func (o *GetTransferInitiationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTransferInitiationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTransferInitiationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTransferInitiationResponse) GetTransferInitiationResponse() *shared.TransferInitiationResponse {
	if o == nil {
		return nil
	}
	return o.TransferInitiationResponse
}
