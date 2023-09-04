// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type UpdateScopeRequest struct {
	UpdateScopeRequest *shared.UpdateScopeRequest `request:"mediaType=application/json"`
	// Scope ID
	ScopeID string `pathParam:"style=simple,explode=false,name=scopeId"`
}

func (o *UpdateScopeRequest) GetUpdateScopeRequest() *shared.UpdateScopeRequest {
	if o == nil {
		return nil
	}
	return o.UpdateScopeRequest
}

func (o *UpdateScopeRequest) GetScopeID() string {
	if o == nil {
		return ""
	}
	return o.ScopeID
}

type UpdateScopeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Updated scope
	UpdateScopeResponse *shared.UpdateScopeResponse
}

func (o *UpdateScopeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateScopeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateScopeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateScopeResponse) GetUpdateScopeResponse() *shared.UpdateScopeResponse {
	if o == nil {
		return nil
	}
	return o.UpdateScopeResponse
}