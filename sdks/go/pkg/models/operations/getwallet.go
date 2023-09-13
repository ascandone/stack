// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type GetWalletRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetWalletResponse struct {
	ContentType string
	// Wallet
	GetWalletResponse *shared.GetWalletResponse
	StatusCode        int
	RawResponse       *http.Response
	// Error
	WalletsErrorResponse *shared.WalletsErrorResponse
}
