# WatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | Pointer to [**DataChanges**](DataChanges.md) |  | [optional] 

## Methods

### NewWatchResponse

`func NewWatchResponse() *WatchResponse`

NewWatchResponse instantiates a new WatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatchResponseWithDefaults

`func NewWatchResponseWithDefaults() *WatchResponse`

NewWatchResponseWithDefaults instantiates a new WatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanges

`func (o *WatchResponse) GetChanges() DataChanges`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *WatchResponse) GetChangesOk() (*DataChanges, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *WatchResponse) SetChanges(v DataChanges)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *WatchResponse) HasChanges() bool`

HasChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


