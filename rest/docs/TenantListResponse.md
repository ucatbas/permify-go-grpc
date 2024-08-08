# TenantListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenants** | Pointer to [**[]Tenant**](Tenant.md) | tenants is a list of tenants. | [optional] 
**ContinuousToken** | Pointer to **string** | continuous_token is a string that can be used to paginate and retrieve the next set of results. | [optional] 

## Methods

### NewTenantListResponse

`func NewTenantListResponse() *TenantListResponse`

NewTenantListResponse instantiates a new TenantListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantListResponseWithDefaults

`func NewTenantListResponseWithDefaults() *TenantListResponse`

NewTenantListResponseWithDefaults instantiates a new TenantListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenants

`func (o *TenantListResponse) GetTenants() []Tenant`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *TenantListResponse) GetTenantsOk() (*[]Tenant, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *TenantListResponse) SetTenants(v []Tenant)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *TenantListResponse) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetContinuousToken

`func (o *TenantListResponse) GetContinuousToken() string`

GetContinuousToken returns the ContinuousToken field if non-nil, zero value otherwise.

### GetContinuousTokenOk

`func (o *TenantListResponse) GetContinuousTokenOk() (*string, bool)`

GetContinuousTokenOk returns a tuple with the ContinuousToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousToken

`func (o *TenantListResponse) SetContinuousToken(v string)`

SetContinuousToken sets ContinuousToken field to given value.

### HasContinuousToken

`func (o *TenantListResponse) HasContinuousToken() bool`

HasContinuousToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


