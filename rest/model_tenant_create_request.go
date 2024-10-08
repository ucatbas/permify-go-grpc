/*
Permify API

Permify is an open source authorization service for creating fine-grained and scalable authorization systems.

API version: v0.10.1
Contact: hello@permify.co
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package permify

import (
	"encoding/json"
)

// checks if the TenantCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantCreateRequest{}

// TenantCreateRequest TenantCreateRequest is the message used for the request to create a tenant.
type TenantCreateRequest struct {
	// id is a unique identifier for the tenant.
	Id *string `json:"id,omitempty"`
	// name is the name of the tenant.
	Name *string `json:"name,omitempty"`
}

// NewTenantCreateRequest instantiates a new TenantCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantCreateRequest() *TenantCreateRequest {
	this := TenantCreateRequest{}
	return &this
}

// NewTenantCreateRequestWithDefaults instantiates a new TenantCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantCreateRequestWithDefaults() *TenantCreateRequest {
	this := TenantCreateRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TenantCreateRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreateRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TenantCreateRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TenantCreateRequest) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TenantCreateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TenantCreateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TenantCreateRequest) SetName(v string) {
	o.Name = &v
}

func (o TenantCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableTenantCreateRequest struct {
	value *TenantCreateRequest
	isSet bool
}

func (v NullableTenantCreateRequest) Get() *TenantCreateRequest {
	return v.value
}

func (v *NullableTenantCreateRequest) Set(val *TenantCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantCreateRequest(val *TenantCreateRequest) *NullableTenantCreateRequest {
	return &NullableTenantCreateRequest{value: val, isSet: true}
}

func (v NullableTenantCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


