/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: develop
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the ScopeAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScopeAllOf{}

// ScopeAllOf struct for ScopeAllOf
type ScopeAllOf struct {
	Id string `json:"id"`
	Transient []string `json:"transient,omitempty"`
}

// NewScopeAllOf instantiates a new ScopeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopeAllOf(id string) *ScopeAllOf {
	this := ScopeAllOf{}
	this.Id = id
	return &this
}

// NewScopeAllOfWithDefaults instantiates a new ScopeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopeAllOfWithDefaults() *ScopeAllOf {
	this := ScopeAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *ScopeAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScopeAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScopeAllOf) SetId(v string) {
	o.Id = v
}

// GetTransient returns the Transient field value if set, zero value otherwise.
func (o *ScopeAllOf) GetTransient() []string {
	if o == nil || IsNil(o.Transient) {
		var ret []string
		return ret
	}
	return o.Transient
}

// GetTransientOk returns a tuple with the Transient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeAllOf) GetTransientOk() ([]string, bool) {
	if o == nil || IsNil(o.Transient) {
		return nil, false
	}
	return o.Transient, true
}

// HasTransient returns a boolean if a field has been set.
func (o *ScopeAllOf) HasTransient() bool {
	if o != nil && !IsNil(o.Transient) {
		return true
	}

	return false
}

// SetTransient gets a reference to the given []string and assigns it to the Transient field.
func (o *ScopeAllOf) SetTransient(v []string) {
	o.Transient = v
}

func (o ScopeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScopeAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Transient) {
		toSerialize["transient"] = o.Transient
	}
	return toSerialize, nil
}

type NullableScopeAllOf struct {
	value *ScopeAllOf
	isSet bool
}

func (v NullableScopeAllOf) Get() *ScopeAllOf {
	return v.value
}

func (v *NullableScopeAllOf) Set(val *ScopeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableScopeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableScopeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopeAllOf(val *ScopeAllOf) *NullableScopeAllOf {
	return &NullableScopeAllOf{value: val, isSet: true}
}

func (v NullableScopeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

