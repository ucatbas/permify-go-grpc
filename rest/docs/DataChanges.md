# DataChanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapToken** | Pointer to **string** | The snapshot token. | [optional] 
**DataChanges** | Pointer to [**[]DataChange**](DataChange.md) | The list of data changes. | [optional] 

## Methods

### NewDataChanges

`func NewDataChanges() *DataChanges`

NewDataChanges instantiates a new DataChanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataChangesWithDefaults

`func NewDataChangesWithDefaults() *DataChanges`

NewDataChangesWithDefaults instantiates a new DataChanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapToken

`func (o *DataChanges) GetSnapToken() string`

GetSnapToken returns the SnapToken field if non-nil, zero value otherwise.

### GetSnapTokenOk

`func (o *DataChanges) GetSnapTokenOk() (*string, bool)`

GetSnapTokenOk returns a tuple with the SnapToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapToken

`func (o *DataChanges) SetSnapToken(v string)`

SetSnapToken sets SnapToken field to given value.

### HasSnapToken

`func (o *DataChanges) HasSnapToken() bool`

HasSnapToken returns a boolean if a field has been set.

### GetDataChanges

`func (o *DataChanges) GetDataChanges() []DataChange`

GetDataChanges returns the DataChanges field if non-nil, zero value otherwise.

### GetDataChangesOk

`func (o *DataChanges) GetDataChangesOk() (*[]DataChange, bool)`

GetDataChangesOk returns a tuple with the DataChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataChanges

`func (o *DataChanges) SetDataChanges(v []DataChange)`

SetDataChanges sets DataChanges field to given value.

### HasDataChanges

`func (o *DataChanges) HasDataChanges() bool`

HasDataChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


