// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type ConfirmHoldRequest struct {
	ConfirmHoldRequest *shared.ConfirmHoldRequest `request:"mediaType=application/json"`
	HoldID             string                     `pathParam:"style=simple,explode=false,name=hold_id"`
}

func (o *ConfirmHoldRequest) GetConfirmHoldRequest() *shared.ConfirmHoldRequest {
	if o == nil {
		return nil
	}
	return o.ConfirmHoldRequest
}

func (o *ConfirmHoldRequest) GetHoldID() string {
	if o == nil {
		return ""
	}
	return o.HoldID
}

type ConfirmHoldResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Error
	WalletsErrorResponse *shared.WalletsErrorResponse
}

func (o *ConfirmHoldResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ConfirmHoldResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ConfirmHoldResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ConfirmHoldResponse) GetWalletsErrorResponse() *shared.WalletsErrorResponse {
	if o == nil {
		return nil
	}
	return o.WalletsErrorResponse
}
