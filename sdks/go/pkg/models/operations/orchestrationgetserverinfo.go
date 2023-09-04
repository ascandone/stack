// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type OrchestrationgetServerInfoResponse struct {
	ContentType string
	// General error
	Error *shared.Error
	// Server information
	ServerInfo  *shared.ServerInfo
	StatusCode  int
	RawResponse *http.Response
}

func (o *OrchestrationgetServerInfoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *OrchestrationgetServerInfoResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *OrchestrationgetServerInfoResponse) GetServerInfo() *shared.ServerInfo {
	if o == nil {
		return nil
	}
	return o.ServerInfo
}

func (o *OrchestrationgetServerInfoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *OrchestrationgetServerInfoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}