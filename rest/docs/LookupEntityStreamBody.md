# LookupEntityStreamBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**PermissionLookupEntityRequestMetadata**](PermissionLookupEntityRequestMetadata.md) |  | [optional] 
**EntityType** | Pointer to **string** | Type of the entity to lookup, required, must start with a letter and can include alphanumeric and underscore, max 64 bytes. | [optional] 
**Permission** | Pointer to **string** | Name of the permission to check, required, must start with a letter and can include alphanumeric and underscore, max 64 bytes. | [optional] 
**Subject** | Pointer to [**Subject**](Subject.md) |  | [optional] 
**Context** | Pointer to [**Context**](Context.md) |  | [optional] 

## Methods

### NewLookupEntityStreamBody

`func NewLookupEntityStreamBody() *LookupEntityStreamBody`

NewLookupEntityStreamBody instantiates a new LookupEntityStreamBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupEntityStreamBodyWithDefaults

`func NewLookupEntityStreamBodyWithDefaults() *LookupEntityStreamBody`

NewLookupEntityStreamBodyWithDefaults instantiates a new LookupEntityStreamBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *LookupEntityStreamBody) GetMetadata() PermissionLookupEntityRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LookupEntityStreamBody) GetMetadataOk() (*PermissionLookupEntityRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LookupEntityStreamBody) SetMetadata(v PermissionLookupEntityRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LookupEntityStreamBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEntityType

`func (o *LookupEntityStreamBody) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *LookupEntityStreamBody) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *LookupEntityStreamBody) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *LookupEntityStreamBody) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetPermission

`func (o *LookupEntityStreamBody) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *LookupEntityStreamBody) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *LookupEntityStreamBody) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *LookupEntityStreamBody) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetSubject

`func (o *LookupEntityStreamBody) GetSubject() Subject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *LookupEntityStreamBody) GetSubjectOk() (*Subject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *LookupEntityStreamBody) SetSubject(v Subject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *LookupEntityStreamBody) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetContext

`func (o *LookupEntityStreamBody) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *LookupEntityStreamBody) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *LookupEntityStreamBody) SetContext(v Context)`

SetContext sets Context field to given value.

### HasContext

`func (o *LookupEntityStreamBody) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


