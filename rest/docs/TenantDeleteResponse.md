# TenantDeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenant** | Pointer to [**Tenant**](Tenant.md) |  | [optional] 

## Methods

### NewTenantDeleteResponse

`func NewTenantDeleteResponse() *TenantDeleteResponse`

NewTenantDeleteResponse instantiates a new TenantDeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantDeleteResponseWithDefaults

`func NewTenantDeleteResponseWithDefaults() *TenantDeleteResponse`

NewTenantDeleteResponseWithDefaults instantiates a new TenantDeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenant

`func (o *TenantDeleteResponse) GetTenant() Tenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TenantDeleteResponse) GetTenantOk() (*Tenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TenantDeleteResponse) SetTenant(v Tenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *TenantDeleteResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


