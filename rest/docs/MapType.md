# MapType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyType** | Pointer to [**V1alpha1Type**](V1alpha1Type.md) |  | [optional] 
**ValueType** | Pointer to [**V1alpha1Type**](V1alpha1Type.md) |  | [optional] 

## Methods

### NewMapType

`func NewMapType() *MapType`

NewMapType instantiates a new MapType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapTypeWithDefaults

`func NewMapTypeWithDefaults() *MapType`

NewMapTypeWithDefaults instantiates a new MapType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyType

`func (o *MapType) GetKeyType() V1alpha1Type`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *MapType) GetKeyTypeOk() (*V1alpha1Type, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *MapType) SetKeyType(v V1alpha1Type)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *MapType) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetValueType

`func (o *MapType) GetValueType() V1alpha1Type`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *MapType) GetValueTypeOk() (*V1alpha1Type, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *MapType) SetValueType(v V1alpha1Type)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *MapType) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


