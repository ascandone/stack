/*
Search API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: latest
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ResponseCursor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCursor{}

// ResponseCursor struct for ResponseCursor
type ResponseCursor struct {
	PageSize interface{} `json:"pageSize,omitempty"`
	HasMore interface{} `json:"hasMore,omitempty"`
	Total *ResponseCursorTotal `json:"total,omitempty"`
	Next interface{} `json:"next,omitempty"`
	Previous interface{} `json:"previous,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// NewResponseCursor instantiates a new ResponseCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCursor() *ResponseCursor {
	this := ResponseCursor{}
	return &this
}

// NewResponseCursorWithDefaults instantiates a new ResponseCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCursorWithDefaults() *ResponseCursor {
	this := ResponseCursor{}
	return &this
}

// GetPageSize returns the PageSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseCursor) GetPageSize() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseCursor) GetPageSizeOk() (*interface{}, bool) {
	if o == nil || isNil(o.PageSize) {
		return nil, false
	}
	return &o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ResponseCursor) HasPageSize() bool {
	if o != nil && isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given interface{} and assigns it to the PageSize field.
func (o *ResponseCursor) SetPageSize(v interface{}) {
	o.PageSize = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseCursor) GetHasMore() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseCursor) GetHasMoreOk() (*interface{}, bool) {
	if o == nil || isNil(o.HasMore) {
		return nil, false
	}
	return &o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *ResponseCursor) HasHasMore() bool {
	if o != nil && isNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given interface{} and assigns it to the HasMore field.
func (o *ResponseCursor) SetHasMore(v interface{}) {
	o.HasMore = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ResponseCursor) GetTotal() ResponseCursorTotal {
	if o == nil || isNil(o.Total) {
		var ret ResponseCursorTotal
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCursor) GetTotalOk() (*ResponseCursorTotal, bool) {
	if o == nil || isNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ResponseCursor) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given ResponseCursorTotal and assigns it to the Total field.
func (o *ResponseCursor) SetTotal(v ResponseCursorTotal) {
	o.Total = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseCursor) GetNext() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseCursor) GetNextOk() (*interface{}, bool) {
	if o == nil || isNil(o.Next) {
		return nil, false
	}
	return &o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ResponseCursor) HasNext() bool {
	if o != nil && isNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given interface{} and assigns it to the Next field.
func (o *ResponseCursor) SetNext(v interface{}) {
	o.Next = v
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseCursor) GetPrevious() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseCursor) GetPreviousOk() (*interface{}, bool) {
	if o == nil || isNil(o.Previous) {
		return nil, false
	}
	return &o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *ResponseCursor) HasPrevious() bool {
	if o != nil && isNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given interface{} and assigns it to the Previous field.
func (o *ResponseCursor) SetPrevious(v interface{}) {
	o.Previous = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseCursor) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseCursor) GetDataOk() (*interface{}, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResponseCursor) HasData() bool {
	if o != nil && isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *ResponseCursor) SetData(v interface{}) {
	o.Data = v
}

func (o ResponseCursor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCursor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}
	if o.HasMore != nil {
		toSerialize["hasMore"] = o.HasMore
	}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableResponseCursor struct {
	value *ResponseCursor
	isSet bool
}

func (v NullableResponseCursor) Get() *ResponseCursor {
	return v.value
}

func (v *NullableResponseCursor) Set(val *ResponseCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCursor(val *ResponseCursor) *NullableResponseCursor {
	return &NullableResponseCursor{value: val, isSet: true}
}

func (v NullableResponseCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

