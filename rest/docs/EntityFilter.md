# EntityFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEntityFilter

`func NewEntityFilter() *EntityFilter`

NewEntityFilter instantiates a new EntityFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityFilterWithDefaults

`func NewEntityFilterWithDefaults() *EntityFilter`

NewEntityFilterWithDefaults instantiates a new EntityFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EntityFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityFilter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntityFilter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIds

`func (o *EntityFilter) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *EntityFilter) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *EntityFilter) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *EntityFilter) HasIds() bool`

HasIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


