# LookupEntityBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**PermissionLookupEntityRequestMetadata**](PermissionLookupEntityRequestMetadata.md) |  | [optional] 
**EntityType** | Pointer to **string** | Type of the entity to lookup, required, must start with a letter and can include alphanumeric and underscore, max 64 bytes. | [optional] 
**Permission** | Pointer to **string** | Name of the permission to check, required, must start with a letter and can include alphanumeric and underscore, max 64 bytes. | [optional] 
**Subject** | Pointer to [**Subject**](Subject.md) |  | [optional] 
**Context** | Pointer to [**Context**](Context.md) |  | [optional] 

## Methods

### NewLookupEntityBody

`func NewLookupEntityBody() *LookupEntityBody`

NewLookupEntityBody instantiates a new LookupEntityBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupEntityBodyWithDefaults

`func NewLookupEntityBodyWithDefaults() *LookupEntityBody`

NewLookupEntityBodyWithDefaults instantiates a new LookupEntityBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *LookupEntityBody) GetMetadata() PermissionLookupEntityRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LookupEntityBody) GetMetadataOk() (*PermissionLookupEntityRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LookupEntityBody) SetMetadata(v PermissionLookupEntityRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LookupEntityBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEntityType

`func (o *LookupEntityBody) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *LookupEntityBody) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *LookupEntityBody) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *LookupEntityBody) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetPermission

`func (o *LookupEntityBody) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *LookupEntityBody) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *LookupEntityBody) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *LookupEntityBody) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetSubject

`func (o *LookupEntityBody) GetSubject() Subject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *LookupEntityBody) GetSubjectOk() (*Subject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *LookupEntityBody) SetSubject(v Subject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *LookupEntityBody) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetContext

`func (o *LookupEntityBody) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *LookupEntityBody) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *LookupEntityBody) SetContext(v Context)`

SetContext sets Context field to given value.

### HasContext

`func (o *LookupEntityBody) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


