# DataBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | &#39;name&#39; is a simple string field representing the name of the DataBundle. | [optional] 
**Arguments** | Pointer to **[]string** | &#39;arguments&#39; is a repeated field, which means it can contain multiple strings. These are used to store a list of arguments related to the DataBundle. | [optional] 
**Operations** | Pointer to [**[]V1Operation**](V1Operation.md) | &#39;operations&#39; is a repeated field containing multiple Operation messages. Each Operation represents a specific action or set of actions to be performed. | [optional] 

## Methods

### NewDataBundle

`func NewDataBundle() *DataBundle`

NewDataBundle instantiates a new DataBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataBundleWithDefaults

`func NewDataBundleWithDefaults() *DataBundle`

NewDataBundleWithDefaults instantiates a new DataBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DataBundle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataBundle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataBundle) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataBundle) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArguments

`func (o *DataBundle) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *DataBundle) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *DataBundle) SetArguments(v []string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *DataBundle) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetOperations

`func (o *DataBundle) GetOperations() []V1Operation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *DataBundle) GetOperationsOk() (*[]V1Operation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *DataBundle) SetOperations(v []V1Operation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *DataBundle) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


