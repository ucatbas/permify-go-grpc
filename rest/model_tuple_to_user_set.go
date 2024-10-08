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

// checks if the TupleToUserSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TupleToUserSet{}

// TupleToUserSet TupleToUserSet defines a mapping from tuple sets to computed user sets.
type TupleToUserSet struct {
	TupleSet *TupleSet `json:"tupleSet,omitempty"`
	Computed *ComputedUserSet `json:"computed,omitempty"`
}

// NewTupleToUserSet instantiates a new TupleToUserSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTupleToUserSet() *TupleToUserSet {
	this := TupleToUserSet{}
	return &this
}

// NewTupleToUserSetWithDefaults instantiates a new TupleToUserSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTupleToUserSetWithDefaults() *TupleToUserSet {
	this := TupleToUserSet{}
	return &this
}

// GetTupleSet returns the TupleSet field value if set, zero value otherwise.
func (o *TupleToUserSet) GetTupleSet() TupleSet {
	if o == nil || IsNil(o.TupleSet) {
		var ret TupleSet
		return ret
	}
	return *o.TupleSet
}

// GetTupleSetOk returns a tuple with the TupleSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TupleToUserSet) GetTupleSetOk() (*TupleSet, bool) {
	if o == nil || IsNil(o.TupleSet) {
		return nil, false
	}
	return o.TupleSet, true
}

// HasTupleSet returns a boolean if a field has been set.
func (o *TupleToUserSet) HasTupleSet() bool {
	if o != nil && !IsNil(o.TupleSet) {
		return true
	}

	return false
}

// SetTupleSet gets a reference to the given TupleSet and assigns it to the TupleSet field.
func (o *TupleToUserSet) SetTupleSet(v TupleSet) {
	o.TupleSet = &v
}

// GetComputed returns the Computed field value if set, zero value otherwise.
func (o *TupleToUserSet) GetComputed() ComputedUserSet {
	if o == nil || IsNil(o.Computed) {
		var ret ComputedUserSet
		return ret
	}
	return *o.Computed
}

// GetComputedOk returns a tuple with the Computed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TupleToUserSet) GetComputedOk() (*ComputedUserSet, bool) {
	if o == nil || IsNil(o.Computed) {
		return nil, false
	}
	return o.Computed, true
}

// HasComputed returns a boolean if a field has been set.
func (o *TupleToUserSet) HasComputed() bool {
	if o != nil && !IsNil(o.Computed) {
		return true
	}

	return false
}

// SetComputed gets a reference to the given ComputedUserSet and assigns it to the Computed field.
func (o *TupleToUserSet) SetComputed(v ComputedUserSet) {
	o.Computed = &v
}

func (o TupleToUserSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TupleToUserSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TupleSet) {
		toSerialize["tupleSet"] = o.TupleSet
	}
	if !IsNil(o.Computed) {
		toSerialize["computed"] = o.Computed
	}
	return toSerialize, nil
}

type NullableTupleToUserSet struct {
	value *TupleToUserSet
	isSet bool
}

func (v NullableTupleToUserSet) Get() *TupleToUserSet {
	return v.value
}

func (v *NullableTupleToUserSet) Set(val *TupleToUserSet) {
	v.value = val
	v.isSet = true
}

func (v NullableTupleToUserSet) IsSet() bool {
	return v.isSet
}

func (v *NullableTupleToUserSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTupleToUserSet(val *TupleToUserSet) *NullableTupleToUserSet {
	return &NullableTupleToUserSet{value: val, isSet: true}
}

func (v NullableTupleToUserSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTupleToUserSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


