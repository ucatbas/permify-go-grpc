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

// checks if the Context type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Context{}

// Context Context encapsulates the information related to a single operation, including the tuples involved and the associated attributes.
type Context struct {
	// A repeated field of tuples involved in the operation.
	Tuples []Tuple `json:"tuples,omitempty"`
	// A repeated field of attributes associated with the operation.
	Attributes []Attribute `json:"attributes,omitempty"`
	// Additional data associated with the context.
	Data map[string]interface{} `json:"data,omitempty"`
}

// NewContext instantiates a new Context object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContext() *Context {
	this := Context{}
	return &this
}

// NewContextWithDefaults instantiates a new Context object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextWithDefaults() *Context {
	this := Context{}
	return &this
}

// GetTuples returns the Tuples field value if set, zero value otherwise.
func (o *Context) GetTuples() []Tuple {
	if o == nil || IsNil(o.Tuples) {
		var ret []Tuple
		return ret
	}
	return o.Tuples
}

// GetTuplesOk returns a tuple with the Tuples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetTuplesOk() ([]Tuple, bool) {
	if o == nil || IsNil(o.Tuples) {
		return nil, false
	}
	return o.Tuples, true
}

// HasTuples returns a boolean if a field has been set.
func (o *Context) HasTuples() bool {
	if o != nil && !IsNil(o.Tuples) {
		return true
	}

	return false
}

// SetTuples gets a reference to the given []Tuple and assigns it to the Tuples field.
func (o *Context) SetTuples(v []Tuple) {
	o.Tuples = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Context) GetAttributes() []Attribute {
	if o == nil || IsNil(o.Attributes) {
		var ret []Attribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetAttributesOk() ([]Attribute, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Context) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []Attribute and assigns it to the Attributes field.
func (o *Context) SetAttributes(v []Attribute) {
	o.Attributes = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Context) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Context) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *Context) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o Context) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Context) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tuples) {
		toSerialize["tuples"] = o.Tuples
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableContext struct {
	value *Context
	isSet bool
}

func (v NullableContext) Get() *Context {
	return v.value
}

func (v *NullableContext) Set(val *Context) {
	v.value = val
	v.isSet = true
}

func (v NullableContext) IsSet() bool {
	return v.isSet
}

func (v *NullableContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContext(val *Context) *NullableContext {
	return &NullableContext{value: val, isSet: true}
}

func (v NullableContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


