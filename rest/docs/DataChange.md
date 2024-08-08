# DataChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to [**DataChangeOperation**](DataChangeOperation.md) |  | [optional] 
**Tuple** | Pointer to [**Tuple**](Tuple.md) |  | [optional] 
**Attribute** | Pointer to [**Attribute**](Attribute.md) |  | [optional] 

## Methods

### NewDataChange

`func NewDataChange() *DataChange`

NewDataChange instantiates a new DataChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataChangeWithDefaults

`func NewDataChangeWithDefaults() *DataChange`

NewDataChangeWithDefaults instantiates a new DataChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *DataChange) GetOperation() DataChangeOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *DataChange) GetOperationOk() (*DataChangeOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *DataChange) SetOperation(v DataChangeOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *DataChange) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetTuple

`func (o *DataChange) GetTuple() Tuple`

GetTuple returns the Tuple field if non-nil, zero value otherwise.

### GetTupleOk

`func (o *DataChange) GetTupleOk() (*Tuple, bool)`

GetTupleOk returns a tuple with the Tuple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuple

`func (o *DataChange) SetTuple(v Tuple)`

SetTuple sets Tuple field to given value.

### HasTuple

`func (o *DataChange) HasTuple() bool`

HasTuple returns a boolean if a field has been set.

### GetAttribute

`func (o *DataChange) GetAttribute() Attribute`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *DataChange) GetAttributeOk() (*Attribute, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *DataChange) SetAttribute(v Attribute)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *DataChange) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


