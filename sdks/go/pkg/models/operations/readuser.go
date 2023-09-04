// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type ReadUserRequest struct {
	// User ID
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

func (o *ReadUserRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type ReadUserResponse struct {
	ContentType string
	// Retrieved user
	ReadUserResponse *shared.ReadUserResponse
	StatusCode       int
	RawResponse      *http.Response
}

func (o *ReadUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReadUserResponse) GetReadUserResponse() *shared.ReadUserResponse {
	if o == nil {
		return nil
	}
	return o.ReadUserResponse
}

func (o *ReadUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReadUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
