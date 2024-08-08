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

// checks if the WriteRelationshipsBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WriteRelationshipsBody{}

// WriteRelationshipsBody Represents a request to write relationship data.
type WriteRelationshipsBody struct {
	Metadata *RelationshipWriteRequestMetadata `json:"metadata,omitempty"`
	// List of tuples for the request. Must have between 1 and 100 items.
	Tuples []Tuple `json:"tuples,omitempty"`
}

// NewWriteRelationshipsBody instantiates a new WriteRelationshipsBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWriteRelationshipsBody() *WriteRelationshipsBody {
	this := WriteRelationshipsBody{}
	return &this
}

// NewWriteRelationshipsBodyWithDefaults instantiates a new WriteRelationshipsBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteRelationshipsBodyWithDefaults() *WriteRelationshipsBody {
	this := WriteRelationshipsBody{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *WriteRelationshipsBody) GetMetadata() RelationshipWriteRequestMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret RelationshipWriteRequestMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteRelationshipsBody) GetMetadataOk() (*RelationshipWriteRequestMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *WriteRelationshipsBody) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given RelationshipWriteRequestMetadata and assigns it to the Metadata field.
func (o *WriteRelationshipsBody) SetMetadata(v RelationshipWriteRequestMetadata) {
	o.Metadata = &v
}

// GetTuples returns the Tuples field value if set, zero value otherwise.
func (o *WriteRelationshipsBody) GetTuples() []Tuple {
	if o == nil || IsNil(o.Tuples) {
		var ret []Tuple
		return ret
	}
	return o.Tuples
}

// GetTuplesOk returns a tuple with the Tuples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteRelationshipsBody) GetTuplesOk() ([]Tuple, bool) {
	if o == nil || IsNil(o.Tuples) {
		return nil, false
	}
	return o.Tuples, true
}

// HasTuples returns a boolean if a field has been set.
func (o *WriteRelationshipsBody) HasTuples() bool {
	if o != nil && !IsNil(o.Tuples) {
		return true
	}

	return false
}

// SetTuples gets a reference to the given []Tuple and assigns it to the Tuples field.
func (o *WriteRelationshipsBody) SetTuples(v []Tuple) {
	o.Tuples = v
}

func (o WriteRelationshipsBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WriteRelationshipsBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Tuples) {
		toSerialize["tuples"] = o.Tuples
	}
	return toSerialize, nil
}

type NullableWriteRelationshipsBody struct {
	value *WriteRelationshipsBody
	isSet bool
}

func (v NullableWriteRelationshipsBody) Get() *WriteRelationshipsBody {
	return v.value
}

func (v *NullableWriteRelationshipsBody) Set(val *WriteRelationshipsBody) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteRelationshipsBody) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteRelationshipsBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteRelationshipsBody(val *WriteRelationshipsBody) *NullableWriteRelationshipsBody {
	return &NullableWriteRelationshipsBody{value: val, isSet: true}
}

func (v NullableWriteRelationshipsBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteRelationshipsBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


