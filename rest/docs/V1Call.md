# V1Call

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleName** | Pointer to **string** |  | [optional] 
**Arguments** | Pointer to [**[]Argument**](Argument.md) |  | [optional] 

## Methods

### NewV1Call

`func NewV1Call() *V1Call`

NewV1Call instantiates a new V1Call object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1CallWithDefaults

`func NewV1CallWithDefaults() *V1Call`

NewV1CallWithDefaults instantiates a new V1Call object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleName

`func (o *V1Call) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *V1Call) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *V1Call) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *V1Call) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetArguments

`func (o *V1Call) GetArguments() []Argument`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *V1Call) GetArgumentsOk() (*[]Argument, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *V1Call) SetArguments(v []Argument)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *V1Call) HasArguments() bool`

HasArguments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


