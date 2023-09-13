// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type PaymentslistAccountsRequest struct {
	// Parameter used in pagination requests. Maximum page size is set to 15.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when this parameter is set.
	//
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// The maximum number of results to return per page.
	//
	PageSize *int64 `queryParam:"style=form,explode=true,name=pageSize"`
	// Fields used to sort payments (default is date:desc).
	Sort []string `queryParam:"style=form,explode=true,name=sort"`
}

type PaymentslistAccountsResponse struct {
	// OK
	AccountsCursor *shared.AccountsCursor
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
}
