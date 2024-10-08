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

// checks if the CheckBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckBody{}

// CheckBody PermissionCheckRequest is the request message for the Check method in the Permission service.
type CheckBody struct {
	Metadata *PermissionCheckRequestMetadata `json:"metadata,omitempty"`
	Entity *Entity `json:"entity,omitempty"`
	// The action the user wants to perform on the resource
	Permission *string `json:"permission,omitempty"`
	Subject *Subject `json:"subject,omitempty"`
	Context *Context `json:"context,omitempty"`
	// Additional arguments associated with this request.
	Arguments []Argument `json:"arguments,omitempty"`
}

// NewCheckBody instantiates a new CheckBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckBody() *CheckBody {
	this := CheckBody{}
	return &this
}

// NewCheckBodyWithDefaults instantiates a new CheckBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckBodyWithDefaults() *CheckBody {
	this := CheckBody{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CheckBody) GetMetadata() PermissionCheckRequestMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret PermissionCheckRequestMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckBody) GetMetadataOk() (*PermissionCheckRequestMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CheckBody) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given PermissionCheckRequestMetadata and assigns it to the Metadata field.
func (o *CheckBody) SetMetadata(v PermissionCheckRequestMetadata) {
	o.Metadata = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *CheckBody) GetEntity() Entity {
	if o == nil || IsNil(o.Entity) {
		var ret Entity
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckBody) GetEntityOk() (*Entity, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *CheckBody) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given Entity and assigns it to the Entity field.
func (o *CheckBody) SetEntity(v Entity) {
	o.Entity = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *CheckBody) GetPermission() string {
	if o == nil || IsNil(o.Permission) {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckBody) GetPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *CheckBody) HasPermission() bool {
	if o != nil && !IsNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *CheckBody) SetPermission(v string) {
	o.Permission = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CheckBody) GetSubject() Subject {
	if o == nil || IsNil(o.Subject) {
		var ret Subject
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckBody) GetSubjectOk() (*Subject, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CheckBody) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given Subject and assigns it to the Subject field.
func (o *CheckBody) SetSubject(v Subject) {
	o.Subject = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *CheckBody) GetContext() Context {
	if o == nil || IsNil(o.Context) {
		var ret Context
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckBody) GetContextOk() (*Context, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *CheckBody) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given Context and assigns it to the Context field.
func (o *CheckBody) SetContext(v Context) {
	o.Context = &v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *CheckBody) GetArguments() []Argument {
	if o == nil || IsNil(o.Arguments) {
		var ret []Argument
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckBody) GetArgumentsOk() ([]Argument, bool) {
	if o == nil || IsNil(o.Arguments) {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *CheckBody) HasArguments() bool {
	if o != nil && !IsNil(o.Arguments) {
		return true
	}

	return false
}

// SetArguments gets a reference to the given []Argument and assigns it to the Arguments field.
func (o *CheckBody) SetArguments(v []Argument) {
	o.Arguments = v
}

func (o CheckBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.Permission) {
		toSerialize["permission"] = o.Permission
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.Arguments) {
		toSerialize["arguments"] = o.Arguments
	}
	return toSerialize, nil
}

type NullableCheckBody struct {
	value *CheckBody
	isSet bool
}

func (v NullableCheckBody) Get() *CheckBody {
	return v.value
}

func (v *NullableCheckBody) Set(val *CheckBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckBody(val *CheckBody) *NullableCheckBody {
	return &NullableCheckBody{value: val, isSet: true}
}

func (v NullableCheckBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


