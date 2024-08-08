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

// checks if the SchemaWriteBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaWriteBody{}

// SchemaWriteBody SchemaWriteRequest is the request message for the Write method in the Schema service. It contains tenant_id and the schema to be written.
type SchemaWriteBody struct {
	// schema is the string representation of the schema to be written.
	Schema *string `json:"schema,omitempty"`
}

// NewSchemaWriteBody instantiates a new SchemaWriteBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaWriteBody() *SchemaWriteBody {
	this := SchemaWriteBody{}
	return &this
}

// NewSchemaWriteBodyWithDefaults instantiates a new SchemaWriteBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaWriteBodyWithDefaults() *SchemaWriteBody {
	this := SchemaWriteBody{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *SchemaWriteBody) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaWriteBody) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *SchemaWriteBody) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *SchemaWriteBody) SetSchema(v string) {
	o.Schema = &v
}

func (o SchemaWriteBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaWriteBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	return toSerialize, nil
}

type NullableSchemaWriteBody struct {
	value *SchemaWriteBody
	isSet bool
}

func (v NullableSchemaWriteBody) Get() *SchemaWriteBody {
	return v.value
}

func (v *NullableSchemaWriteBody) Set(val *SchemaWriteBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaWriteBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaWriteBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaWriteBody(val *SchemaWriteBody) *NullableSchemaWriteBody {
	return &NullableSchemaWriteBody{value: val, isSet: true}
}

func (v NullableSchemaWriteBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaWriteBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


