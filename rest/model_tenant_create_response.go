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

// checks if the TenantCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantCreateResponse{}

// TenantCreateResponse TenantCreateResponse is the message returned from the request to create a tenant.
type TenantCreateResponse struct {
	Tenant *Tenant `json:"tenant,omitempty"`
}

// NewTenantCreateResponse instantiates a new TenantCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantCreateResponse() *TenantCreateResponse {
	this := TenantCreateResponse{}
	return &this
}

// NewTenantCreateResponseWithDefaults instantiates a new TenantCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantCreateResponseWithDefaults() *TenantCreateResponse {
	this := TenantCreateResponse{}
	return &this
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *TenantCreateResponse) GetTenant() Tenant {
	if o == nil || IsNil(o.Tenant) {
		var ret Tenant
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreateResponse) GetTenantOk() (*Tenant, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *TenantCreateResponse) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given Tenant and assigns it to the Tenant field.
func (o *TenantCreateResponse) SetTenant(v Tenant) {
	o.Tenant = &v
}

func (o TenantCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	return toSerialize, nil
}

type NullableTenantCreateResponse struct {
	value *TenantCreateResponse
	isSet bool
}

func (v NullableTenantCreateResponse) Get() *TenantCreateResponse {
	return v.value
}

func (v *NullableTenantCreateResponse) Set(val *TenantCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantCreateResponse(val *TenantCreateResponse) *NullableTenantCreateResponse {
	return &NullableTenantCreateResponse{value: val, isSet: true}
}

func (v NullableTenantCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


