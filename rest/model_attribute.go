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

// checks if the Attribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Attribute{}

// Attribute Attribute represents an attribute of an entity with a specific type and value.
type Attribute struct {
	Entity *Entity `json:"entity,omitempty"`
	Attribute *string `json:"attribute,omitempty"`
	Value *Any `json:"value,omitempty"`
}

// NewAttribute instantiates a new Attribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttribute() *Attribute {
	this := Attribute{}
	return &this
}

// NewAttributeWithDefaults instantiates a new Attribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeWithDefaults() *Attribute {
	this := Attribute{}
	return &this
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *Attribute) GetEntity() Entity {
	if o == nil || IsNil(o.Entity) {
		var ret Entity
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetEntityOk() (*Entity, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *Attribute) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given Entity and assigns it to the Entity field.
func (o *Attribute) SetEntity(v Entity) {
	o.Entity = &v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *Attribute) GetAttribute() string {
	if o == nil || IsNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *Attribute) HasAttribute() bool {
	if o != nil && !IsNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *Attribute) SetAttribute(v string) {
	o.Attribute = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Attribute) GetValue() Any {
	if o == nil || IsNil(o.Value) {
		var ret Any
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetValueOk() (*Any, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Attribute) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given Any and assigns it to the Value field.
func (o *Attribute) SetValue(v Any) {
	o.Value = &v
}

func (o Attribute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Attribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableAttribute struct {
	value *Attribute
	isSet bool
}

func (v NullableAttribute) Get() *Attribute {
	return v.value
}

func (v *NullableAttribute) Set(val *Attribute) {
	v.value = val
	v.isSet = true
}

func (v NullableAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttribute(val *Attribute) *NullableAttribute {
	return &NullableAttribute{value: val, isSet: true}
}

func (v NullableAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


