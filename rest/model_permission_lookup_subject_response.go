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

// checks if the PermissionLookupSubjectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionLookupSubjectResponse{}

// PermissionLookupSubjectResponse PermissionLookupSubjectResponse is the response message for the LookupSubject method in the Permission service.
type PermissionLookupSubjectResponse struct {
	// List of identifiers for subjects that match the lookup.
	SubjectIds []string `json:"subject_ids,omitempty"`
}

// NewPermissionLookupSubjectResponse instantiates a new PermissionLookupSubjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionLookupSubjectResponse() *PermissionLookupSubjectResponse {
	this := PermissionLookupSubjectResponse{}
	return &this
}

// NewPermissionLookupSubjectResponseWithDefaults instantiates a new PermissionLookupSubjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionLookupSubjectResponseWithDefaults() *PermissionLookupSubjectResponse {
	this := PermissionLookupSubjectResponse{}
	return &this
}

// GetSubjectIds returns the SubjectIds field value if set, zero value otherwise.
func (o *PermissionLookupSubjectResponse) GetSubjectIds() []string {
	if o == nil || IsNil(o.SubjectIds) {
		var ret []string
		return ret
	}
	return o.SubjectIds
}

// GetSubjectIdsOk returns a tuple with the SubjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionLookupSubjectResponse) GetSubjectIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SubjectIds) {
		return nil, false
	}
	return o.SubjectIds, true
}

// HasSubjectIds returns a boolean if a field has been set.
func (o *PermissionLookupSubjectResponse) HasSubjectIds() bool {
	if o != nil && !IsNil(o.SubjectIds) {
		return true
	}

	return false
}

// SetSubjectIds gets a reference to the given []string and assigns it to the SubjectIds field.
func (o *PermissionLookupSubjectResponse) SetSubjectIds(v []string) {
	o.SubjectIds = v
}

func (o PermissionLookupSubjectResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionLookupSubjectResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubjectIds) {
		toSerialize["subject_ids"] = o.SubjectIds
	}
	return toSerialize, nil
}

type NullablePermissionLookupSubjectResponse struct {
	value *PermissionLookupSubjectResponse
	isSet bool
}

func (v NullablePermissionLookupSubjectResponse) Get() *PermissionLookupSubjectResponse {
	return v.value
}

func (v *NullablePermissionLookupSubjectResponse) Set(val *PermissionLookupSubjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionLookupSubjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionLookupSubjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionLookupSubjectResponse(val *PermissionLookupSubjectResponse) *NullablePermissionLookupSubjectResponse {
	return &NullablePermissionLookupSubjectResponse{value: val, isSet: true}
}

func (v NullablePermissionLookupSubjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionLookupSubjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


