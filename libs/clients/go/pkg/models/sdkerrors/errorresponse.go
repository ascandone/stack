// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

type ErrorResponse struct {
	Details      *string    `json:"details,omitempty"`
	ErrorCode    ErrorsEnum `json:"errorCode"`
	ErrorMessage string     `json:"errorMessage"`
}

var _ error = &ErrorResponse{}

func (e *ErrorResponse) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}