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

// checks if the SchemaPartialWriteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaPartialWriteResponse{}

// SchemaPartialWriteResponse SchemaPartialWriteResponse is the response message for the Parietal Write method in the Schema service. It returns the requested schema.
type SchemaPartialWriteResponse struct {
	// schema_version is the string that identifies the version of the written schema.
	SchemaVersion *string `json:"schema_version,omitempty"`
}

// NewSchemaPartialWriteResponse instantiates a new SchemaPartialWriteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaPartialWriteResponse() *SchemaPartialWriteResponse {
	this := SchemaPartialWriteResponse{}
	return &this
}

// NewSchemaPartialWriteResponseWithDefaults instantiates a new SchemaPartialWriteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaPartialWriteResponseWithDefaults() *SchemaPartialWriteResponse {
	this := SchemaPartialWriteResponse{}
	return &this
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *SchemaPartialWriteResponse) GetSchemaVersion() string {
	if o == nil || IsNil(o.SchemaVersion) {
		var ret string
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaPartialWriteResponse) GetSchemaVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaVersion) {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *SchemaPartialWriteResponse) HasSchemaVersion() bool {
	if o != nil && !IsNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given string and assigns it to the SchemaVersion field.
func (o *SchemaPartialWriteResponse) SetSchemaVersion(v string) {
	o.SchemaVersion = &v
}

func (o SchemaPartialWriteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaPartialWriteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SchemaVersion) {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	return toSerialize, nil
}

type NullableSchemaPartialWriteResponse struct {
	value *SchemaPartialWriteResponse
	isSet bool
}

func (v NullableSchemaPartialWriteResponse) Get() *SchemaPartialWriteResponse {
	return v.value
}

func (v *NullableSchemaPartialWriteResponse) Set(val *SchemaPartialWriteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaPartialWriteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaPartialWriteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaPartialWriteResponse(val *SchemaPartialWriteResponse) *NullableSchemaPartialWriteResponse {
	return &NullableSchemaPartialWriteResponse{value: val, isSet: true}
}

func (v NullableSchemaPartialWriteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaPartialWriteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


