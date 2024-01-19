// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type GetConnectorTaskV1Request struct {
	// The name of the connector.
	Connector shared.Connector `pathParam:"style=simple,explode=false,name=connector"`
	// The connector ID.
	ConnectorID string `pathParam:"style=simple,explode=false,name=connectorId"`
	// The task ID.
	TaskID string `pathParam:"style=simple,explode=false,name=taskId"`
}

func (o *GetConnectorTaskV1Request) GetConnector() shared.Connector {
	if o == nil {
		return shared.Connector("")
	}
	return o.Connector
}

func (o *GetConnectorTaskV1Request) GetConnectorID() string {
	if o == nil {
		return ""
	}
	return o.ConnectorID
}

func (o *GetConnectorTaskV1Request) GetTaskID() string {
	if o == nil {
		return ""
	}
	return o.TaskID
}

type GetConnectorTaskV1Response struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TaskResponse *shared.TaskResponse
}

func (o *GetConnectorTaskV1Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetConnectorTaskV1Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConnectorTaskV1Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetConnectorTaskV1Response) GetTaskResponse() *shared.TaskResponse {
	if o == nil {
		return nil
	}
	return o.TaskResponse
}
