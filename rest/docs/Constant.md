# Constant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NullValue** | Pointer to **string** | null value. | [optional] 
**BoolValue** | Pointer to **bool** | boolean value. | [optional] 
**Int64Value** | Pointer to **string** | int64 value. | [optional] 
**Uint64Value** | Pointer to **string** | uint64 value. | [optional] 
**DoubleValue** | Pointer to **float64** | double value. | [optional] 
**StringValue** | Pointer to **string** | string value. | [optional] 
**BytesValue** | Pointer to **string** | bytes value. | [optional] 
**DurationValue** | Pointer to **string** | protobuf.Duration value.  Deprecated: duration is no longer considered a builtin cel type. | [optional] 
**TimestampValue** | Pointer to **time.Time** | protobuf.Timestamp value.  Deprecated: timestamp is no longer considered a builtin cel type. | [optional] 

## Methods

### NewConstant

`func NewConstant() *Constant`

NewConstant instantiates a new Constant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstantWithDefaults

`func NewConstantWithDefaults() *Constant`

NewConstantWithDefaults instantiates a new Constant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNullValue

`func (o *Constant) GetNullValue() string`

GetNullValue returns the NullValue field if non-nil, zero value otherwise.

### GetNullValueOk

`func (o *Constant) GetNullValueOk() (*string, bool)`

GetNullValueOk returns a tuple with the NullValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullValue

`func (o *Constant) SetNullValue(v string)`

SetNullValue sets NullValue field to given value.

### HasNullValue

`func (o *Constant) HasNullValue() bool`

HasNullValue returns a boolean if a field has been set.

### GetBoolValue

`func (o *Constant) GetBoolValue() bool`

GetBoolValue returns the BoolValue field if non-nil, zero value otherwise.

### GetBoolValueOk

`func (o *Constant) GetBoolValueOk() (*bool, bool)`

GetBoolValueOk returns a tuple with the BoolValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolValue

`func (o *Constant) SetBoolValue(v bool)`

SetBoolValue sets BoolValue field to given value.

### HasBoolValue

`func (o *Constant) HasBoolValue() bool`

HasBoolValue returns a boolean if a field has been set.

### GetInt64Value

`func (o *Constant) GetInt64Value() string`

GetInt64Value returns the Int64Value field if non-nil, zero value otherwise.

### GetInt64ValueOk

`func (o *Constant) GetInt64ValueOk() (*string, bool)`

GetInt64ValueOk returns a tuple with the Int64Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInt64Value

`func (o *Constant) SetInt64Value(v string)`

SetInt64Value sets Int64Value field to given value.

### HasInt64Value

`func (o *Constant) HasInt64Value() bool`

HasInt64Value returns a boolean if a field has been set.

### GetUint64Value

`func (o *Constant) GetUint64Value() string`

GetUint64Value returns the Uint64Value field if non-nil, zero value otherwise.

### GetUint64ValueOk

`func (o *Constant) GetUint64ValueOk() (*string, bool)`

GetUint64ValueOk returns a tuple with the Uint64Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUint64Value

`func (o *Constant) SetUint64Value(v string)`

SetUint64Value sets Uint64Value field to given value.

### HasUint64Value

`func (o *Constant) HasUint64Value() bool`

HasUint64Value returns a boolean if a field has been set.

### GetDoubleValue

`func (o *Constant) GetDoubleValue() float64`

GetDoubleValue returns the DoubleValue field if non-nil, zero value otherwise.

### GetDoubleValueOk

`func (o *Constant) GetDoubleValueOk() (*float64, bool)`

GetDoubleValueOk returns a tuple with the DoubleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleValue

`func (o *Constant) SetDoubleValue(v float64)`

SetDoubleValue sets DoubleValue field to given value.

### HasDoubleValue

`func (o *Constant) HasDoubleValue() bool`

HasDoubleValue returns a boolean if a field has been set.

### GetStringValue

`func (o *Constant) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *Constant) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *Constant) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *Constant) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### GetBytesValue

`func (o *Constant) GetBytesValue() string`

GetBytesValue returns the BytesValue field if non-nil, zero value otherwise.

### GetBytesValueOk

`func (o *Constant) GetBytesValueOk() (*string, bool)`

GetBytesValueOk returns a tuple with the BytesValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesValue

`func (o *Constant) SetBytesValue(v string)`

SetBytesValue sets BytesValue field to given value.

### HasBytesValue

`func (o *Constant) HasBytesValue() bool`

HasBytesValue returns a boolean if a field has been set.

### GetDurationValue

`func (o *Constant) GetDurationValue() string`

GetDurationValue returns the DurationValue field if non-nil, zero value otherwise.

### GetDurationValueOk

`func (o *Constant) GetDurationValueOk() (*string, bool)`

GetDurationValueOk returns a tuple with the DurationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationValue

`func (o *Constant) SetDurationValue(v string)`

SetDurationValue sets DurationValue field to given value.

### HasDurationValue

`func (o *Constant) HasDurationValue() bool`

HasDurationValue returns a boolean if a field has been set.

### GetTimestampValue

`func (o *Constant) GetTimestampValue() time.Time`

GetTimestampValue returns the TimestampValue field if non-nil, zero value otherwise.

### GetTimestampValueOk

`func (o *Constant) GetTimestampValueOk() (*time.Time, bool)`

GetTimestampValueOk returns a tuple with the TimestampValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampValue

`func (o *Constant) SetTimestampValue(v time.Time)`

SetTimestampValue sets TimestampValue field to given value.

### HasTimestampValue

`func (o *Constant) HasTimestampValue() bool`

HasTimestampValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


