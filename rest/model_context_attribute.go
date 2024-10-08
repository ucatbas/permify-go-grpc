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

// checks if the ContextAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextAttribute{}

// ContextAttribute ContextAttribute defines a context attribute which includes its name.
type ContextAttribute struct {
	Name *string `json:"name,omitempty"`
}

// NewContextAttribute instantiates a new ContextAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextAttribute() *ContextAttribute {
	this := ContextAttribute{}
	return &this
}

// NewContextAttributeWithDefaults instantiates a new ContextAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextAttributeWithDefaults() *ContextAttribute {
	this := ContextAttribute{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContextAttribute) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextAttribute) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContextAttribute) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContextAttribute) SetName(v string) {
	o.Name = &v
}

func (o ContextAttribute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableContextAttribute struct {
	value *ContextAttribute
	isSet bool
}

func (v NullableContextAttribute) Get() *ContextAttribute {
	return v.value
}

func (v *NullableContextAttribute) Set(val *ContextAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableContextAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableContextAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextAttribute(val *ContextAttribute) *NullableContextAttribute {
	return &NullableContextAttribute{value: val, isSet: true}
}

func (v NullableContextAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


