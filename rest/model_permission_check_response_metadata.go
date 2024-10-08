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

// checks if the PermissionCheckResponseMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionCheckResponseMetadata{}

// PermissionCheckResponseMetadata PermissionCheckResponseMetadata metadata for the PermissionCheckResponse.
type PermissionCheckResponseMetadata struct {
	// The count of the checks performed.
	CheckCount *int32 `json:"check_count,omitempty"`
}

// NewPermissionCheckResponseMetadata instantiates a new PermissionCheckResponseMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionCheckResponseMetadata() *PermissionCheckResponseMetadata {
	this := PermissionCheckResponseMetadata{}
	return &this
}

// NewPermissionCheckResponseMetadataWithDefaults instantiates a new PermissionCheckResponseMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionCheckResponseMetadataWithDefaults() *PermissionCheckResponseMetadata {
	this := PermissionCheckResponseMetadata{}
	return &this
}

// GetCheckCount returns the CheckCount field value if set, zero value otherwise.
func (o *PermissionCheckResponseMetadata) GetCheckCount() int32 {
	if o == nil || IsNil(o.CheckCount) {
		var ret int32
		return ret
	}
	return *o.CheckCount
}

// GetCheckCountOk returns a tuple with the CheckCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionCheckResponseMetadata) GetCheckCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CheckCount) {
		return nil, false
	}
	return o.CheckCount, true
}

// HasCheckCount returns a boolean if a field has been set.
func (o *PermissionCheckResponseMetadata) HasCheckCount() bool {
	if o != nil && !IsNil(o.CheckCount) {
		return true
	}

	return false
}

// SetCheckCount gets a reference to the given int32 and assigns it to the CheckCount field.
func (o *PermissionCheckResponseMetadata) SetCheckCount(v int32) {
	o.CheckCount = &v
}

func (o PermissionCheckResponseMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionCheckResponseMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CheckCount) {
		toSerialize["check_count"] = o.CheckCount
	}
	return toSerialize, nil
}

type NullablePermissionCheckResponseMetadata struct {
	value *PermissionCheckResponseMetadata
	isSet bool
}

func (v NullablePermissionCheckResponseMetadata) Get() *PermissionCheckResponseMetadata {
	return v.value
}

func (v *NullablePermissionCheckResponseMetadata) Set(val *PermissionCheckResponseMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionCheckResponseMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionCheckResponseMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionCheckResponseMetadata(val *PermissionCheckResponseMetadata) *NullablePermissionCheckResponseMetadata {
	return &NullablePermissionCheckResponseMetadata{value: val, isSet: true}
}

func (v NullablePermissionCheckResponseMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionCheckResponseMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


