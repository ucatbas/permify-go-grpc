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
	"fmt"
)

// SchemaDefinitionReference The Reference enum helps distinguish whether a name corresponds to an entity or a rule.   - REFERENCE_ENTITY: Indicates that the name refers to an entity.  - REFERENCE_RULE: Indicates that the name refers to a rule.
type SchemaDefinitionReference string

// List of SchemaDefinition.Reference
const (
	SCHEMADEFINITIONREFERENCE_ENTITY SchemaDefinitionReference = "REFERENCE_ENTITY"
	SCHEMADEFINITIONREFERENCE_RULE SchemaDefinitionReference = "REFERENCE_RULE"
)

// All allowed values of SchemaDefinitionReference enum
var AllowedSchemaDefinitionReferenceEnumValues = []SchemaDefinitionReference{
	"REFERENCE_ENTITY",
	"REFERENCE_RULE",
}

func (v *SchemaDefinitionReference) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SchemaDefinitionReference(value)
	for _, existing := range AllowedSchemaDefinitionReferenceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SchemaDefinitionReference", value)
}

// NewSchemaDefinitionReferenceFromValue returns a pointer to a valid SchemaDefinitionReference
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSchemaDefinitionReferenceFromValue(v string) (*SchemaDefinitionReference, error) {
	ev := SchemaDefinitionReference(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SchemaDefinitionReference: valid values are %v", v, AllowedSchemaDefinitionReferenceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SchemaDefinitionReference) IsValid() bool {
	for _, existing := range AllowedSchemaDefinitionReferenceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SchemaDefinition.Reference value
func (v SchemaDefinitionReference) Ptr() *SchemaDefinitionReference {
	return &v
}

type NullableSchemaDefinitionReference struct {
	value *SchemaDefinitionReference
	isSet bool
}

func (v NullableSchemaDefinitionReference) Get() *SchemaDefinitionReference {
	return v.value
}

func (v *NullableSchemaDefinitionReference) Set(val *SchemaDefinitionReference) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaDefinitionReference) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaDefinitionReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaDefinitionReference(val *SchemaDefinitionReference) *NullableSchemaDefinitionReference {
	return &NullableSchemaDefinitionReference{value: val, isSet: true}
}

func (v NullableSchemaDefinitionReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaDefinitionReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

