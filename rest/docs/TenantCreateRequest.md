# TenantCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | id is a unique identifier for the tenant. | [optional] 
**Name** | Pointer to **string** | name is the name of the tenant. | [optional] 

## Methods

### NewTenantCreateRequest

`func NewTenantCreateRequest() *TenantCreateRequest`

NewTenantCreateRequest instantiates a new TenantCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantCreateRequestWithDefaults

`func NewTenantCreateRequestWithDefaults() *TenantCreateRequest`

NewTenantCreateRequestWithDefaults instantiates a new TenantCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantCreateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantCreateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantCreateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantCreateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TenantCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


