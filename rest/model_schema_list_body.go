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

// checks if the SchemaListBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaListBody{}

// SchemaListBody SchemaListRequest is the request message for the List method in the Schema service. It contains tenant_id for which the schemas are to be listed.
type SchemaListBody struct {
	// page_size is the number of tenants to be returned in the response. The value should be between 1 and 100.
	PageSize *int64 `json:"page_size,omitempty"`
	// continuous_token is an optional parameter used for pagination. It should be the value received in the previous response.
	ContinuousToken *string `json:"continuous_token,omitempty"`
}

// NewSchemaListBody instantiates a new SchemaListBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaListBody() *SchemaListBody {
	this := SchemaListBody{}
	return &this
}

// NewSchemaListBodyWithDefaults instantiates a new SchemaListBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaListBodyWithDefaults() *SchemaListBody {
	this := SchemaListBody{}
	return &this
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *SchemaListBody) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaListBody) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *SchemaListBody) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *SchemaListBody) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetContinuousToken returns the ContinuousToken field value if set, zero value otherwise.
func (o *SchemaListBody) GetContinuousToken() string {
	if o == nil || IsNil(o.ContinuousToken) {
		var ret string
		return ret
	}
	return *o.ContinuousToken
}

// GetContinuousTokenOk returns a tuple with the ContinuousToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaListBody) GetContinuousTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ContinuousToken) {
		return nil, false
	}
	return o.ContinuousToken, true
}

// HasContinuousToken returns a boolean if a field has been set.
func (o *SchemaListBody) HasContinuousToken() bool {
	if o != nil && !IsNil(o.ContinuousToken) {
		return true
	}

	return false
}

// SetContinuousToken gets a reference to the given string and assigns it to the ContinuousToken field.
func (o *SchemaListBody) SetContinuousToken(v string) {
	o.ContinuousToken = &v
}

func (o SchemaListBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaListBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.ContinuousToken) {
		toSerialize["continuous_token"] = o.ContinuousToken
	}
	return toSerialize, nil
}

type NullableSchemaListBody struct {
	value *SchemaListBody
	isSet bool
}

func (v NullableSchemaListBody) Get() *SchemaListBody {
	return v.value
}

func (v *NullableSchemaListBody) Set(val *SchemaListBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaListBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaListBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaListBody(val *SchemaListBody) *NullableSchemaListBody {
	return &NullableSchemaListBody{value: val, isSet: true}
}

func (v NullableSchemaListBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaListBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


