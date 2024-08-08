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

// checks if the DataChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataChange{}

// DataChange DataChange represents a single change in data, with an operation type and the actual change which could be a tuple or an attribute.
type DataChange struct {
	Operation *DataChangeOperation `json:"operation,omitempty"`
	Tuple *Tuple `json:"tuple,omitempty"`
	Attribute *Attribute `json:"attribute,omitempty"`
}

// NewDataChange instantiates a new DataChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataChange() *DataChange {
	this := DataChange{}
	return &this
}

// NewDataChangeWithDefaults instantiates a new DataChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataChangeWithDefaults() *DataChange {
	this := DataChange{}
	return &this
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *DataChange) GetOperation() DataChangeOperation {
	if o == nil || IsNil(o.Operation) {
		var ret DataChangeOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChange) GetOperationOk() (*DataChangeOperation, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *DataChange) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given DataChangeOperation and assigns it to the Operation field.
func (o *DataChange) SetOperation(v DataChangeOperation) {
	o.Operation = &v
}

// GetTuple returns the Tuple field value if set, zero value otherwise.
func (o *DataChange) GetTuple() Tuple {
	if o == nil || IsNil(o.Tuple) {
		var ret Tuple
		return ret
	}
	return *o.Tuple
}

// GetTupleOk returns a tuple with the Tuple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChange) GetTupleOk() (*Tuple, bool) {
	if o == nil || IsNil(o.Tuple) {
		return nil, false
	}
	return o.Tuple, true
}

// HasTuple returns a boolean if a field has been set.
func (o *DataChange) HasTuple() bool {
	if o != nil && !IsNil(o.Tuple) {
		return true
	}

	return false
}

// SetTuple gets a reference to the given Tuple and assigns it to the Tuple field.
func (o *DataChange) SetTuple(v Tuple) {
	o.Tuple = &v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *DataChange) GetAttribute() Attribute {
	if o == nil || IsNil(o.Attribute) {
		var ret Attribute
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChange) GetAttributeOk() (*Attribute, bool) {
	if o == nil || IsNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *DataChange) HasAttribute() bool {
	if o != nil && !IsNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given Attribute and assigns it to the Attribute field.
func (o *DataChange) SetAttribute(v Attribute) {
	o.Attribute = &v
}

func (o DataChange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !IsNil(o.Tuple) {
		toSerialize["tuple"] = o.Tuple
	}
	if !IsNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	return toSerialize, nil
}

type NullableDataChange struct {
	value *DataChange
	isSet bool
}

func (v NullableDataChange) Get() *DataChange {
	return v.value
}

func (v *NullableDataChange) Set(val *DataChange) {
	v.value = val
	v.isSet = true
}

func (v NullableDataChange) IsSet() bool {
	return v.isSet
}

func (v *NullableDataChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataChange(val *DataChange) *NullableDataChange {
	return &NullableDataChange{value: val, isSet: true}
}

func (v NullableDataChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


