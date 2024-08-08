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

// checks if the AttributeReadRequestMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributeReadRequestMetadata{}

// AttributeReadRequestMetadata AttributeReadRequestMetadata defines the structure for the metadata of an attribute read request. It includes the snap_token associated with a particular state of the database.
type AttributeReadRequestMetadata struct {
	// The snap token to avoid stale cache, see more details on [Snap Tokens](../../operations/snap-tokens)
	SnapToken *string `json:"snap_token,omitempty"`
}

// NewAttributeReadRequestMetadata instantiates a new AttributeReadRequestMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeReadRequestMetadata() *AttributeReadRequestMetadata {
	this := AttributeReadRequestMetadata{}
	return &this
}

// NewAttributeReadRequestMetadataWithDefaults instantiates a new AttributeReadRequestMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeReadRequestMetadataWithDefaults() *AttributeReadRequestMetadata {
	this := AttributeReadRequestMetadata{}
	return &this
}

// GetSnapToken returns the SnapToken field value if set, zero value otherwise.
func (o *AttributeReadRequestMetadata) GetSnapToken() string {
	if o == nil || IsNil(o.SnapToken) {
		var ret string
		return ret
	}
	return *o.SnapToken
}

// GetSnapTokenOk returns a tuple with the SnapToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeReadRequestMetadata) GetSnapTokenOk() (*string, bool) {
	if o == nil || IsNil(o.SnapToken) {
		return nil, false
	}
	return o.SnapToken, true
}

// HasSnapToken returns a boolean if a field has been set.
func (o *AttributeReadRequestMetadata) HasSnapToken() bool {
	if o != nil && !IsNil(o.SnapToken) {
		return true
	}

	return false
}

// SetSnapToken gets a reference to the given string and assigns it to the SnapToken field.
func (o *AttributeReadRequestMetadata) SetSnapToken(v string) {
	o.SnapToken = &v
}

func (o AttributeReadRequestMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributeReadRequestMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SnapToken) {
		toSerialize["snap_token"] = o.SnapToken
	}
	return toSerialize, nil
}

type NullableAttributeReadRequestMetadata struct {
	value *AttributeReadRequestMetadata
	isSet bool
}

func (v NullableAttributeReadRequestMetadata) Get() *AttributeReadRequestMetadata {
	return v.value
}

func (v *NullableAttributeReadRequestMetadata) Set(val *AttributeReadRequestMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeReadRequestMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeReadRequestMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeReadRequestMetadata(val *AttributeReadRequestMetadata) *NullableAttributeReadRequestMetadata {
	return &NullableAttributeReadRequestMetadata{value: val, isSet: true}
}

func (v NullableAttributeReadRequestMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeReadRequestMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


