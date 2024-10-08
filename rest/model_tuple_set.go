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

// checks if the TupleSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TupleSet{}

// TupleSet TupleSet represents a set of tuples associated with a specific relation.
type TupleSet struct {
	Relation *string `json:"relation,omitempty"`
}

// NewTupleSet instantiates a new TupleSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTupleSet() *TupleSet {
	this := TupleSet{}
	return &this
}

// NewTupleSetWithDefaults instantiates a new TupleSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTupleSetWithDefaults() *TupleSet {
	this := TupleSet{}
	return &this
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *TupleSet) GetRelation() string {
	if o == nil || IsNil(o.Relation) {
		var ret string
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TupleSet) GetRelationOk() (*string, bool) {
	if o == nil || IsNil(o.Relation) {
		return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *TupleSet) HasRelation() bool {
	if o != nil && !IsNil(o.Relation) {
		return true
	}

	return false
}

// SetRelation gets a reference to the given string and assigns it to the Relation field.
func (o *TupleSet) SetRelation(v string) {
	o.Relation = &v
}

func (o TupleSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TupleSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Relation) {
		toSerialize["relation"] = o.Relation
	}
	return toSerialize, nil
}

type NullableTupleSet struct {
	value *TupleSet
	isSet bool
}

func (v NullableTupleSet) Get() *TupleSet {
	return v.value
}

func (v *NullableTupleSet) Set(val *TupleSet) {
	v.value = val
	v.isSet = true
}

func (v NullableTupleSet) IsSet() bool {
	return v.isSet
}

func (v *NullableTupleSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTupleSet(val *TupleSet) *NullableTupleSet {
	return &NullableTupleSet{value: val, isSet: true}
}

func (v NullableTupleSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTupleSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


