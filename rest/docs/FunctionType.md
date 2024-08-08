# FunctionType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultType** | Pointer to [**V1alpha1Type**](V1alpha1Type.md) |  | [optional] 
**ArgTypes** | Pointer to [**[]V1alpha1Type**](V1alpha1Type.md) | Argument types of the function. | [optional] 

## Methods

### NewFunctionType

`func NewFunctionType() *FunctionType`

NewFunctionType instantiates a new FunctionType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionTypeWithDefaults

`func NewFunctionTypeWithDefaults() *FunctionType`

NewFunctionTypeWithDefaults instantiates a new FunctionType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultType

`func (o *FunctionType) GetResultType() V1alpha1Type`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *FunctionType) GetResultTypeOk() (*V1alpha1Type, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *FunctionType) SetResultType(v V1alpha1Type)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *FunctionType) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetArgTypes

`func (o *FunctionType) GetArgTypes() []V1alpha1Type`

GetArgTypes returns the ArgTypes field if non-nil, zero value otherwise.

### GetArgTypesOk

`func (o *FunctionType) GetArgTypesOk() (*[]V1alpha1Type, bool)`

GetArgTypesOk returns a tuple with the ArgTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgTypes

`func (o *FunctionType) SetArgTypes(v []V1alpha1Type)`

SetArgTypes sets ArgTypes field to given value.

### HasArgTypes

`func (o *FunctionType) HasArgTypes() bool`

HasArgTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


