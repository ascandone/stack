// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type PaymentsgetServerInfoResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Server information
	ServerInfo *shared.ServerInfo
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PaymentsgetServerInfoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PaymentsgetServerInfoResponse) GetServerInfo() *shared.ServerInfo {
	if o == nil {
		return nil
	}
	return o.ServerInfo
}

func (o *PaymentsgetServerInfoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PaymentsgetServerInfoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
