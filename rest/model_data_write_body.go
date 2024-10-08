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

// checks if the DataWriteBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataWriteBody{}

// DataWriteBody DataWriteRequest defines the structure of a request for writing data. It contains the necessary information such as tenant_id, metadata, tuples and attributes for the write operation.
type DataWriteBody struct {
	Metadata *DataWriteRequestMetadata `json:"metadata,omitempty"`
	// tuples contains the list of tuples (entity-relation-entity triples) that need to be written.
	Tuples []Tuple `json:"tuples,omitempty"`
	// attributes contains the list of attributes (entity-attribute-value triples) that need to be written.
	Attributes []Attribute `json:"attributes,omitempty"`
}

// NewDataWriteBody instantiates a new DataWriteBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataWriteBody() *DataWriteBody {
	this := DataWriteBody{}
	return &this
}

// NewDataWriteBodyWithDefaults instantiates a new DataWriteBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataWriteBodyWithDefaults() *DataWriteBody {
	this := DataWriteBody{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DataWriteBody) GetMetadata() DataWriteRequestMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret DataWriteRequestMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWriteBody) GetMetadataOk() (*DataWriteRequestMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DataWriteBody) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given DataWriteRequestMetadata and assigns it to the Metadata field.
func (o *DataWriteBody) SetMetadata(v DataWriteRequestMetadata) {
	o.Metadata = &v
}

// GetTuples returns the Tuples field value if set, zero value otherwise.
func (o *DataWriteBody) GetTuples() []Tuple {
	if o == nil || IsNil(o.Tuples) {
		var ret []Tuple
		return ret
	}
	return o.Tuples
}

// GetTuplesOk returns a tuple with the Tuples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWriteBody) GetTuplesOk() ([]Tuple, bool) {
	if o == nil || IsNil(o.Tuples) {
		return nil, false
	}
	return o.Tuples, true
}

// HasTuples returns a boolean if a field has been set.
func (o *DataWriteBody) HasTuples() bool {
	if o != nil && !IsNil(o.Tuples) {
		return true
	}

	return false
}

// SetTuples gets a reference to the given []Tuple and assigns it to the Tuples field.
func (o *DataWriteBody) SetTuples(v []Tuple) {
	o.Tuples = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DataWriteBody) GetAttributes() []Attribute {
	if o == nil || IsNil(o.Attributes) {
		var ret []Attribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataWriteBody) GetAttributesOk() ([]Attribute, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DataWriteBody) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []Attribute and assigns it to the Attributes field.
func (o *DataWriteBody) SetAttributes(v []Attribute) {
	o.Attributes = v
}

func (o DataWriteBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataWriteBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Tuples) {
		toSerialize["tuples"] = o.Tuples
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableDataWriteBody struct {
	value *DataWriteBody
	isSet bool
}

func (v NullableDataWriteBody) Get() *DataWriteBody {
	return v.value
}

func (v *NullableDataWriteBody) Set(val *DataWriteBody) {
	v.value = val
	v.isSet = true
}

func (v NullableDataWriteBody) IsSet() bool {
	return v.isSet
}

func (v *NullableDataWriteBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataWriteBody(val *DataWriteBody) *NullableDataWriteBody {
	return &NullableDataWriteBody{value: val, isSet: true}
}

func (v NullableDataWriteBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataWriteBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


