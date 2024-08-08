# ReadAttributesBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**AttributeReadRequestMetadata**](AttributeReadRequestMetadata.md) |  | [optional] 
**Filter** | Pointer to [**AttributeFilter**](AttributeFilter.md) |  | [optional] 
**PageSize** | Pointer to **int64** | page_size specifies the number of results to return in a single page. If more results are available, a continuous_token is included in the response. | [optional] 
**ContinuousToken** | Pointer to **string** | continuous_token is used in case of paginated reads to get the next page of results. | [optional] 

## Methods

### NewReadAttributesBody

`func NewReadAttributesBody() *ReadAttributesBody`

NewReadAttributesBody instantiates a new ReadAttributesBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadAttributesBodyWithDefaults

`func NewReadAttributesBodyWithDefaults() *ReadAttributesBody`

NewReadAttributesBodyWithDefaults instantiates a new ReadAttributesBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ReadAttributesBody) GetMetadata() AttributeReadRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReadAttributesBody) GetMetadataOk() (*AttributeReadRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReadAttributesBody) SetMetadata(v AttributeReadRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ReadAttributesBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFilter

`func (o *ReadAttributesBody) GetFilter() AttributeFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ReadAttributesBody) GetFilterOk() (*AttributeFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ReadAttributesBody) SetFilter(v AttributeFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ReadAttributesBody) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPageSize

`func (o *ReadAttributesBody) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ReadAttributesBody) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ReadAttributesBody) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ReadAttributesBody) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetContinuousToken

`func (o *ReadAttributesBody) GetContinuousToken() string`

GetContinuousToken returns the ContinuousToken field if non-nil, zero value otherwise.

### GetContinuousTokenOk

`func (o *ReadAttributesBody) GetContinuousTokenOk() (*string, bool)`

GetContinuousTokenOk returns a tuple with the ContinuousToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousToken

`func (o *ReadAttributesBody) SetContinuousToken(v string)`

SetContinuousToken sets ContinuousToken field to given value.

### HasContinuousToken

`func (o *ReadAttributesBody) HasContinuousToken() bool`

HasContinuousToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


