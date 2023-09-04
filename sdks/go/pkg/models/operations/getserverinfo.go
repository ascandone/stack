// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type GetServerInfoResponse struct {
	ContentType string
	// Server information
	ServerInfo  *shared.ServerInfo
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetServerInfoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetServerInfoResponse) GetServerInfo() *shared.ServerInfo {
	if o == nil {
		return nil
	}
	return o.ServerInfo
}

func (o *GetServerInfoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetServerInfoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
