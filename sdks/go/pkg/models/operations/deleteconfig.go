// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type DeleteConfigRequest struct {
	// Config ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteConfigRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteConfigResponse struct {
	ContentType string
	// Error
	ErrorResponse *shared.ErrorResponse
	StatusCode    int
	RawResponse   *http.Response
}

func (o *DeleteConfigResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteConfigResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *DeleteConfigResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteConfigResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
