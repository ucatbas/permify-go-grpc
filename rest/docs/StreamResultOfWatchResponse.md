# StreamResultOfWatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**WatchResponse**](WatchResponse.md) |  | [optional] 
**Error** | Pointer to [**Status**](Status.md) |  | [optional] 

## Methods

### NewStreamResultOfWatchResponse

`func NewStreamResultOfWatchResponse() *StreamResultOfWatchResponse`

NewStreamResultOfWatchResponse instantiates a new StreamResultOfWatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamResultOfWatchResponseWithDefaults

`func NewStreamResultOfWatchResponseWithDefaults() *StreamResultOfWatchResponse`

NewStreamResultOfWatchResponseWithDefaults instantiates a new StreamResultOfWatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *StreamResultOfWatchResponse) GetResult() WatchResponse`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *StreamResultOfWatchResponse) GetResultOk() (*WatchResponse, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *StreamResultOfWatchResponse) SetResult(v WatchResponse)`

SetResult sets Result field to given value.

### HasResult

`func (o *StreamResultOfWatchResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetError

`func (o *StreamResultOfWatchResponse) GetError() Status`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StreamResultOfWatchResponse) GetErrorOk() (*Status, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StreamResultOfWatchResponse) SetError(v Status)`

SetError sets Error field to given value.

### HasError

`func (o *StreamResultOfWatchResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


