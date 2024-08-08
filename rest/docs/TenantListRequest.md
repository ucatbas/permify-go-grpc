# TenantListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageSize** | Pointer to **int64** | page_size is the number of tenants to be returned in the response. The value should be between 1 and 100. | [optional] 
**ContinuousToken** | Pointer to **string** | continuous_token is an optional parameter used for pagination. It should be the value received in the previous response. | [optional] 

## Methods

### NewTenantListRequest

`func NewTenantListRequest() *TenantListRequest`

NewTenantListRequest instantiates a new TenantListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantListRequestWithDefaults

`func NewTenantListRequestWithDefaults() *TenantListRequest`

NewTenantListRequestWithDefaults instantiates a new TenantListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageSize

`func (o *TenantListRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *TenantListRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *TenantListRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *TenantListRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetContinuousToken

`func (o *TenantListRequest) GetContinuousToken() string`

GetContinuousToken returns the ContinuousToken field if non-nil, zero value otherwise.

### GetContinuousTokenOk

`func (o *TenantListRequest) GetContinuousTokenOk() (*string, bool)`

GetContinuousTokenOk returns a tuple with the ContinuousToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousToken

`func (o *TenantListRequest) SetContinuousToken(v string)`

SetContinuousToken sets ContinuousToken field to given value.

### HasContinuousToken

`func (o *TenantListRequest) HasContinuousToken() bool`

HasContinuousToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


