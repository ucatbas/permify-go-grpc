# AbstractType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The fully qualified name of this abstract type. | [optional] 
**ParameterTypes** | Pointer to [**[]V1alpha1Type**](V1alpha1Type.md) | Parameter types for this abstract type. | [optional] 

## Methods

### NewAbstractType

`func NewAbstractType() *AbstractType`

NewAbstractType instantiates a new AbstractType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractTypeWithDefaults

`func NewAbstractTypeWithDefaults() *AbstractType`

NewAbstractTypeWithDefaults instantiates a new AbstractType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AbstractType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AbstractType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AbstractType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AbstractType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameterTypes

`func (o *AbstractType) GetParameterTypes() []V1alpha1Type`

GetParameterTypes returns the ParameterTypes field if non-nil, zero value otherwise.

### GetParameterTypesOk

`func (o *AbstractType) GetParameterTypesOk() (*[]V1alpha1Type, bool)`

GetParameterTypesOk returns a tuple with the ParameterTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterTypes

`func (o *AbstractType) SetParameterTypes(v []V1alpha1Type)`

SetParameterTypes sets ParameterTypes field to given value.

### HasParameterTypes

`func (o *AbstractType) HasParameterTypes() bool`

HasParameterTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


