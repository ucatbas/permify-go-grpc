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

// checks if the Tuple type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tuple{}

// Tuple Tuple is a structure that includes an entity, a relation, and a subject.
type Tuple struct {
	Entity *Entity `json:"entity,omitempty"`
	Relation *string `json:"relation,omitempty"`
	Subject *Subject `json:"subject,omitempty"`
}

// NewTuple instantiates a new Tuple object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTuple() *Tuple {
	this := Tuple{}
	return &this
}

// NewTupleWithDefaults instantiates a new Tuple object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTupleWithDefaults() *Tuple {
	this := Tuple{}
	return &this
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *Tuple) GetEntity() Entity {
	if o == nil || IsNil(o.Entity) {
		var ret Entity
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tuple) GetEntityOk() (*Entity, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *Tuple) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given Entity and assigns it to the Entity field.
func (o *Tuple) SetEntity(v Entity) {
	o.Entity = &v
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *Tuple) GetRelation() string {
	if o == nil || IsNil(o.Relation) {
		var ret string
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tuple) GetRelationOk() (*string, bool) {
	if o == nil || IsNil(o.Relation) {
		return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *Tuple) HasRelation() bool {
	if o != nil && !IsNil(o.Relation) {
		return true
	}

	return false
}

// SetRelation gets a reference to the given string and assigns it to the Relation field.
func (o *Tuple) SetRelation(v string) {
	o.Relation = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *Tuple) GetSubject() Subject {
	if o == nil || IsNil(o.Subject) {
		var ret Subject
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tuple) GetSubjectOk() (*Subject, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *Tuple) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given Subject and assigns it to the Subject field.
func (o *Tuple) SetSubject(v Subject) {
	o.Subject = &v
}

func (o Tuple) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tuple) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.Relation) {
		toSerialize["relation"] = o.Relation
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	return toSerialize, nil
}

type NullableTuple struct {
	value *Tuple
	isSet bool
}

func (v NullableTuple) Get() *Tuple {
	return v.value
}

func (v *NullableTuple) Set(val *Tuple) {
	v.value = val
	v.isSet = true
}

func (v NullableTuple) IsSet() bool {
	return v.isSet
}

func (v *NullableTuple) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTuple(val *Tuple) *NullableTuple {
	return &NullableTuple{value: val, isSet: true}
}

func (v NullableTuple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTuple) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


