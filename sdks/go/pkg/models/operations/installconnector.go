// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type InstallConnectorRequest struct {
	RequestBody interface{} `request:"mediaType=application/json"`
	// The name of the connector.
	Connector shared.Connector `pathParam:"style=simple,explode=false,name=connector"`
}

func (o *InstallConnectorRequest) GetRequestBody() interface{} {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *InstallConnectorRequest) GetConnector() shared.Connector {
	if o == nil {
		return shared.Connector("")
	}
	return o.Connector
}

type InstallConnectorResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *InstallConnectorResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *InstallConnectorResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *InstallConnectorResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}