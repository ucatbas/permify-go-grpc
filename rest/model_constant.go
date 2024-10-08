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
	"time"
)

// checks if the Constant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Constant{}

// Constant Represents a primitive literal.  Named 'Constant' here for backwards compatibility.  This is similar as the primitives supported in the well-known type `google.protobuf.Value`, but richer so it can represent CEL's full range of primitives.  Lists and structs are not included as constants as these aggregate types may contain [Expr][google.api.expr.v1alpha1.Expr] elements which require evaluation and are thus not constant.  Examples of literals include: `\"hello\"`, `b'bytes'`, `1u`, `4.2`, `-2`, `true`, `null`.
type Constant struct {
	// null value.
	NullValue *string `json:"nullValue,omitempty"`
	// boolean value.
	BoolValue *bool `json:"boolValue,omitempty"`
	// int64 value.
	Int64Value *string `json:"int64Value,omitempty"`
	// uint64 value.
	Uint64Value *string `json:"uint64Value,omitempty"`
	// double value.
	DoubleValue *float64 `json:"doubleValue,omitempty"`
	// string value.
	StringValue *string `json:"stringValue,omitempty"`
	// bytes value.
	BytesValue *string `json:"bytesValue,omitempty" validate:"regexp=^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=)?$"`
	// protobuf.Duration value.  Deprecated: duration is no longer considered a builtin cel type.
	DurationValue *string `json:"durationValue,omitempty"`
	// protobuf.Timestamp value.  Deprecated: timestamp is no longer considered a builtin cel type.
	TimestampValue *time.Time `json:"timestampValue,omitempty"`
}

// NewConstant instantiates a new Constant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstant() *Constant {
	this := Constant{}
	return &this
}

// NewConstantWithDefaults instantiates a new Constant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstantWithDefaults() *Constant {
	this := Constant{}
	return &this
}

// GetNullValue returns the NullValue field value if set, zero value otherwise.
func (o *Constant) GetNullValue() string {
	if o == nil || IsNil(o.NullValue) {
		var ret string
		return ret
	}
	return *o.NullValue
}

// GetNullValueOk returns a tuple with the NullValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constant) GetNullValueOk() (*string, bool) {
	if o == nil || IsNil(o.NullValue) {
		return nil, false
	}
	return o.NullValue, true
}

// HasNullValue returns a boolean if a field has been set.
func (o *Constant) HasNullValue() bool {
	if o != nil && !IsNil(o.NullValue) {
		return true
	}

	return false
}

// SetNullValue gets a reference to the given string and assigns it to the NullValue field.
func (o *Constant) SetNullValue(v string) {
	o.NullValue = &v
}

// GetBoolValue returns the BoolValue field value if set, zero value otherwise.
func (o *Constant) GetBoolValue() bool {
	if o == nil || IsNil(o.BoolValue) {
		var ret bool
		return ret
	}
	return *o.BoolValue
}

// GetBoolValueOk returns a tuple with the BoolValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constant) GetBoolValueOk() (*bool, bool) {
	if o == nil || IsNil(o.BoolValue) {
		return nil, false
	}
	return o.BoolValue, true
}

// HasBoolValue returns a boolean if a field has been set.
func (o *Constant) HasBoolValue() bool {
	if o != nil && !IsNil(o.BoolValue) {
		return true
	}

	return false
}

// SetBoolValue gets a reference to the given bool and assigns it to the BoolValue field.
func (o *Constant) SetBoolValue(v bool) {
	o.BoolValue = &v
}

// GetInt64Value returns the Int64Value field value if set, zero value otherwise.
func (o *Constant) GetInt64Value() string {
	if o == nil || IsNil(o.Int64Value) {
		var ret string
		return ret
	}
	return *o.Int64Value
}

// GetInt64ValueOk returns a tuple with the Int64Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constant) GetInt64ValueOk() (*string, bool) {
	if o == nil || IsNil(o.Int64Value) {
		return nil, false
	}
	return o.Int64Value, true
}

// HasInt64Value returns a boolean if a field has been set.
func (o *Constant) HasInt64Value() bool {
	if o != nil && !IsNil(o.Int64Value) {
		return true
	}

	return false
}

// SetInt64Value gets a reference to the given string and assigns it to the Int64Value field.
func (o *Constant) SetInt64Value(v string) {
	o.Int64Value = &v
}

// GetUint64Value returns the Uint64Value field value if set, zero value otherwise.
func (o *Constant) GetUint64Value() string {
	if o == nil || IsNil(o.Uint64Value) {
		var ret string
		return ret
	}
	return *o.Uint64Value
}

// GetUint64ValueOk returns a tuple with the Uint64Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constant) GetUint64ValueOk() (*string, bool) {
	if o == nil || IsNil(o.Uint64Value) {
		return nil, false
	}
	return o.Uint64Value, true
}

// HasUint64Value returns a boolean if a field has been set.
func (o *Constant) HasUint64Value() bool {
	if o != nil && !IsNil(o.Uint64Value) {
		return true
	}

	return false
}

// SetUint64Value gets a reference to the given string and assigns it to the Uint64Value field.
func (o *Constant) SetUint64Value(v string) {
	o.Uint64Value = &v
}

// GetDoubleValue returns the DoubleValue field value if set, zero value otherwise.
func (o *Constant) GetDoubleValue() float64 {
	if o == nil || IsNil(o.DoubleValue) {
		var ret float64
		return ret
	}
	return *o.DoubleValue
}

// GetDoubleValueOk returns a tuple with the DoubleValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constant) GetDoubleValueOk() (*float64, bool) {
	if o == nil || IsNil(o.DoubleValue) {
		return nil, false
	}
	return o.DoubleValue, true
}

// HasDoubleValue returns a boolean if a field has been set.
func (o *Constant) HasDoubleValue() bool {
	if o != nil && !IsNil(o.DoubleValue) {
		return true
	}

	return false
}

// SetDoubleValue gets a reference to the given float64 and assigns it to the DoubleValue field.
func (o *Constant) SetDoubleValue(v float64) {
	o.DoubleValue = &v
}

// GetStringValue returns the StringValue field value if set, zero value otherwise.
func (o *Constant) GetStringValue() string {
	if o == nil || IsNil(o.StringValue) {
		var ret string
		return ret
	}
	return *o.StringValue
}

// GetStringValueOk returns a tuple with the StringValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constant) GetStringValueOk() (*string, bool) {
	if o == nil || IsNil(o.StringValue) {
		return nil, false
	}
	return o.StringValue, true
}

// HasStringValue returns a boolean if a field has been set.
func (o *Constant) HasStringValue() bool {
	if o != nil && !IsNil(o.StringValue) {
		return true
	}

	return false
}

// SetStringValue gets a reference to the given string and assigns it to the StringValue field.
func (o *Constant) SetStringValue(v string) {
	o.StringValue = &v
}

// GetBytesValue returns the BytesValue field value if set, zero value otherwise.
func (o *Constant) GetBytesValue() string {
	if o == nil || IsNil(o.BytesValue) {
		var ret string
		return ret
	}
	return *o.BytesValue
}

// GetBytesValueOk returns a tuple with the BytesValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constant) GetBytesValueOk() (*string, bool) {
	if o == nil || IsNil(o.BytesValue) {
		return nil, false
	}
	return o.BytesValue, true
}

// HasBytesValue returns a boolean if a field has been set.
func (o *Constant) HasBytesValue() bool {
	if o != nil && !IsNil(o.BytesValue) {
		return true
	}

	return false
}

// SetBytesValue gets a reference to the given string and assigns it to the BytesValue field.
func (o *Constant) SetBytesValue(v string) {
	o.BytesValue = &v
}

// GetDurationValue returns the DurationValue field value if set, zero value otherwise.
func (o *Constant) GetDurationValue() string {
	if o == nil || IsNil(o.DurationValue) {
		var ret string
		return ret
	}
	return *o.DurationValue
}

// GetDurationValueOk returns a tuple with the DurationValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constant) GetDurationValueOk() (*string, bool) {
	if o == nil || IsNil(o.DurationValue) {
		return nil, false
	}
	return o.DurationValue, true
}

// HasDurationValue returns a boolean if a field has been set.
func (o *Constant) HasDurationValue() bool {
	if o != nil && !IsNil(o.DurationValue) {
		return true
	}

	return false
}

// SetDurationValue gets a reference to the given string and assigns it to the DurationValue field.
func (o *Constant) SetDurationValue(v string) {
	o.DurationValue = &v
}

// GetTimestampValue returns the TimestampValue field value if set, zero value otherwise.
func (o *Constant) GetTimestampValue() time.Time {
	if o == nil || IsNil(o.TimestampValue) {
		var ret time.Time
		return ret
	}
	return *o.TimestampValue
}

// GetTimestampValueOk returns a tuple with the TimestampValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constant) GetTimestampValueOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimestampValue) {
		return nil, false
	}
	return o.TimestampValue, true
}

// HasTimestampValue returns a boolean if a field has been set.
func (o *Constant) HasTimestampValue() bool {
	if o != nil && !IsNil(o.TimestampValue) {
		return true
	}

	return false
}

// SetTimestampValue gets a reference to the given time.Time and assigns it to the TimestampValue field.
func (o *Constant) SetTimestampValue(v time.Time) {
	o.TimestampValue = &v
}

func (o Constant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Constant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NullValue) {
		toSerialize["nullValue"] = o.NullValue
	}
	if !IsNil(o.BoolValue) {
		toSerialize["boolValue"] = o.BoolValue
	}
	if !IsNil(o.Int64Value) {
		toSerialize["int64Value"] = o.Int64Value
	}
	if !IsNil(o.Uint64Value) {
		toSerialize["uint64Value"] = o.Uint64Value
	}
	if !IsNil(o.DoubleValue) {
		toSerialize["doubleValue"] = o.DoubleValue
	}
	if !IsNil(o.StringValue) {
		toSerialize["stringValue"] = o.StringValue
	}
	if !IsNil(o.BytesValue) {
		toSerialize["bytesValue"] = o.BytesValue
	}
	if !IsNil(o.DurationValue) {
		toSerialize["durationValue"] = o.DurationValue
	}
	if !IsNil(o.TimestampValue) {
		toSerialize["timestampValue"] = o.TimestampValue
	}
	return toSerialize, nil
}

type NullableConstant struct {
	value *Constant
	isSet bool
}

func (v NullableConstant) Get() *Constant {
	return v.value
}

func (v *NullableConstant) Set(val *Constant) {
	v.value = val
	v.isSet = true
}

func (v NullableConstant) IsSet() bool {
	return v.isSet
}

func (v *NullableConstant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstant(val *Constant) *NullableConstant {
	return &NullableConstant{value: val, isSet: true}
}

func (v NullableConstant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


