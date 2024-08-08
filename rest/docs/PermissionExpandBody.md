# PermissionExpandBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**PermissionExpandRequestMetadata**](PermissionExpandRequestMetadata.md) |  | [optional] 
**Entity** | Pointer to [**Entity**](Entity.md) |  | [optional] 
**Permission** | Pointer to **string** | Name of the permission to be expanded, not required, must start with a letter and can include alphanumeric and underscore, max 64 bytes. | [optional] 
**Context** | Pointer to [**Context**](Context.md) |  | [optional] 
**Arguments** | Pointer to [**[]Argument**](Argument.md) | Additional arguments associated with this request. | [optional] 

## Methods

### NewPermissionExpandBody

`func NewPermissionExpandBody() *PermissionExpandBody`

NewPermissionExpandBody instantiates a new PermissionExpandBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionExpandBodyWithDefaults

`func NewPermissionExpandBodyWithDefaults() *PermissionExpandBody`

NewPermissionExpandBodyWithDefaults instantiates a new PermissionExpandBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *PermissionExpandBody) GetMetadata() PermissionExpandRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PermissionExpandBody) GetMetadataOk() (*PermissionExpandRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PermissionExpandBody) SetMetadata(v PermissionExpandRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PermissionExpandBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEntity

`func (o *PermissionExpandBody) GetEntity() Entity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *PermissionExpandBody) GetEntityOk() (*Entity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *PermissionExpandBody) SetEntity(v Entity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *PermissionExpandBody) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetPermission

`func (o *PermissionExpandBody) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *PermissionExpandBody) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *PermissionExpandBody) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *PermissionExpandBody) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetContext

`func (o *PermissionExpandBody) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PermissionExpandBody) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PermissionExpandBody) SetContext(v Context)`

SetContext sets Context field to given value.

### HasContext

`func (o *PermissionExpandBody) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetArguments

`func (o *PermissionExpandBody) GetArguments() []Argument`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *PermissionExpandBody) GetArgumentsOk() (*[]Argument, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *PermissionExpandBody) SetArguments(v []Argument)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *PermissionExpandBody) HasArguments() bool`

HasArguments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


