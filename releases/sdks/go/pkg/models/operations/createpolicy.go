// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type CreatePolicyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	PolicyResponse *shared.PolicyResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreatePolicyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePolicyResponse) GetPolicyResponse() *shared.PolicyResponse {
	if o == nil {
		return nil
	}
	return o.PolicyResponse
}

func (o *CreatePolicyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePolicyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
