# PermissionCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | Pointer to [**CheckResult**](CheckResult.md) |  | [optional] 
**Metadata** | Pointer to [**PermissionCheckResponseMetadata**](PermissionCheckResponseMetadata.md) |  | [optional] 

## Methods

### NewPermissionCheckResponse

`func NewPermissionCheckResponse() *PermissionCheckResponse`

NewPermissionCheckResponse instantiates a new PermissionCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionCheckResponseWithDefaults

`func NewPermissionCheckResponseWithDefaults() *PermissionCheckResponse`

NewPermissionCheckResponseWithDefaults instantiates a new PermissionCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCan

`func (o *PermissionCheckResponse) GetCan() CheckResult`

GetCan returns the Can field if non-nil, zero value otherwise.

### GetCanOk

`func (o *PermissionCheckResponse) GetCanOk() (*CheckResult, bool)`

GetCanOk returns a tuple with the Can field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCan

`func (o *PermissionCheckResponse) SetCan(v CheckResult)`

SetCan sets Can field to given value.

### HasCan

`func (o *PermissionCheckResponse) HasCan() bool`

HasCan returns a boolean if a field has been set.

### GetMetadata

`func (o *PermissionCheckResponse) GetMetadata() PermissionCheckResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PermissionCheckResponse) GetMetadataOk() (*PermissionCheckResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PermissionCheckResponse) SetMetadata(v PermissionCheckResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PermissionCheckResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


