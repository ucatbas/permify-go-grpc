# DataWriteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapToken** | Pointer to **string** | The snap token to avoid stale cache, see more details on [Snap Tokens](../../operations/snap-tokens). | [optional] 

## Methods

### NewDataWriteResponse

`func NewDataWriteResponse() *DataWriteResponse`

NewDataWriteResponse instantiates a new DataWriteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataWriteResponseWithDefaults

`func NewDataWriteResponseWithDefaults() *DataWriteResponse`

NewDataWriteResponseWithDefaults instantiates a new DataWriteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapToken

`func (o *DataWriteResponse) GetSnapToken() string`

GetSnapToken returns the SnapToken field if non-nil, zero value otherwise.

### GetSnapTokenOk

`func (o *DataWriteResponse) GetSnapTokenOk() (*string, bool)`

GetSnapTokenOk returns a tuple with the SnapToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapToken

`func (o *DataWriteResponse) SetSnapToken(v string)`

SetSnapToken sets SnapToken field to given value.

### HasSnapToken

`func (o *DataWriteResponse) HasSnapToken() bool`

HasSnapToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


