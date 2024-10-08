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

// checks if the RelationshipWriteRequestMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelationshipWriteRequestMetadata{}

// RelationshipWriteRequestMetadata struct for RelationshipWriteRequestMetadata
type RelationshipWriteRequestMetadata struct {
	SchemaVersion *string `json:"schema_version,omitempty"`
}

// NewRelationshipWriteRequestMetadata instantiates a new RelationshipWriteRequestMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipWriteRequestMetadata() *RelationshipWriteRequestMetadata {
	this := RelationshipWriteRequestMetadata{}
	return &this
}

// NewRelationshipWriteRequestMetadataWithDefaults instantiates a new RelationshipWriteRequestMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipWriteRequestMetadataWithDefaults() *RelationshipWriteRequestMetadata {
	this := RelationshipWriteRequestMetadata{}
	return &this
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *RelationshipWriteRequestMetadata) GetSchemaVersion() string {
	if o == nil || IsNil(o.SchemaVersion) {
		var ret string
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipWriteRequestMetadata) GetSchemaVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaVersion) {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *RelationshipWriteRequestMetadata) HasSchemaVersion() bool {
	if o != nil && !IsNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given string and assigns it to the SchemaVersion field.
func (o *RelationshipWriteRequestMetadata) SetSchemaVersion(v string) {
	o.SchemaVersion = &v
}

func (o RelationshipWriteRequestMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelationshipWriteRequestMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SchemaVersion) {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	return toSerialize, nil
}

type NullableRelationshipWriteRequestMetadata struct {
	value *RelationshipWriteRequestMetadata
	isSet bool
}

func (v NullableRelationshipWriteRequestMetadata) Get() *RelationshipWriteRequestMetadata {
	return v.value
}

func (v *NullableRelationshipWriteRequestMetadata) Set(val *RelationshipWriteRequestMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipWriteRequestMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipWriteRequestMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipWriteRequestMetadata(val *RelationshipWriteRequestMetadata) *NullableRelationshipWriteRequestMetadata {
	return &NullableRelationshipWriteRequestMetadata{value: val, isSet: true}
}

func (v NullableRelationshipWriteRequestMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipWriteRequestMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


