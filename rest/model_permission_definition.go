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

// checks if the PermissionDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionDefinition{}

// PermissionDefinition The PermissionDefinition message provides detailed information about a specific permission.
type PermissionDefinition struct {
	// The name of the permission, which follows a specific string pattern and has a maximum byte size.
	Name *string `json:"name,omitempty"`
	Child *Child `json:"child,omitempty"`
}

// NewPermissionDefinition instantiates a new PermissionDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionDefinition() *PermissionDefinition {
	this := PermissionDefinition{}
	return &this
}

// NewPermissionDefinitionWithDefaults instantiates a new PermissionDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionDefinitionWithDefaults() *PermissionDefinition {
	this := PermissionDefinition{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PermissionDefinition) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionDefinition) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PermissionDefinition) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PermissionDefinition) SetName(v string) {
	o.Name = &v
}

// GetChild returns the Child field value if set, zero value otherwise.
func (o *PermissionDefinition) GetChild() Child {
	if o == nil || IsNil(o.Child) {
		var ret Child
		return ret
	}
	return *o.Child
}

// GetChildOk returns a tuple with the Child field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionDefinition) GetChildOk() (*Child, bool) {
	if o == nil || IsNil(o.Child) {
		return nil, false
	}
	return o.Child, true
}

// HasChild returns a boolean if a field has been set.
func (o *PermissionDefinition) HasChild() bool {
	if o != nil && !IsNil(o.Child) {
		return true
	}

	return false
}

// SetChild gets a reference to the given Child and assigns it to the Child field.
func (o *PermissionDefinition) SetChild(v Child) {
	o.Child = &v
}

func (o PermissionDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Child) {
		toSerialize["child"] = o.Child
	}
	return toSerialize, nil
}

type NullablePermissionDefinition struct {
	value *PermissionDefinition
	isSet bool
}

func (v NullablePermissionDefinition) Get() *PermissionDefinition {
	return v.value
}

func (v *NullablePermissionDefinition) Set(val *PermissionDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionDefinition(val *PermissionDefinition) *NullablePermissionDefinition {
	return &NullablePermissionDefinition{value: val, isSet: true}
}

func (v NullablePermissionDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


