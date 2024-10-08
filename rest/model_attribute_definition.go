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

// checks if the AttributeDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributeDefinition{}

// AttributeDefinition The AttributeDefinition message provides detailed information about a specific attribute.
type AttributeDefinition struct {
	// The name of the attribute, which follows a specific string pattern and has a maximum byte size.
	Name *string `json:"name,omitempty"`
	Type *AttributeType `json:"type,omitempty"`
}

// NewAttributeDefinition instantiates a new AttributeDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeDefinition() *AttributeDefinition {
	this := AttributeDefinition{}
	return &this
}

// NewAttributeDefinitionWithDefaults instantiates a new AttributeDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeDefinitionWithDefaults() *AttributeDefinition {
	this := AttributeDefinition{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AttributeDefinition) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeDefinition) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AttributeDefinition) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AttributeDefinition) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AttributeDefinition) GetType() AttributeType {
	if o == nil || IsNil(o.Type) {
		var ret AttributeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeDefinition) GetTypeOk() (*AttributeType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AttributeDefinition) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given AttributeType and assigns it to the Type field.
func (o *AttributeDefinition) SetType(v AttributeType) {
	o.Type = &v
}

func (o AttributeDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributeDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAttributeDefinition struct {
	value *AttributeDefinition
	isSet bool
}

func (v NullableAttributeDefinition) Get() *AttributeDefinition {
	return v.value
}

func (v *NullableAttributeDefinition) Set(val *AttributeDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeDefinition(val *AttributeDefinition) *NullableAttributeDefinition {
	return &NullableAttributeDefinition{value: val, isSet: true}
}

func (v NullableAttributeDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


