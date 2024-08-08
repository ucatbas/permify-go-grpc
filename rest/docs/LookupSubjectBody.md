# LookupSubjectBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**PermissionLookupSubjectRequestMetadata**](PermissionLookupSubjectRequestMetadata.md) |  | [optional] 
**Entity** | Pointer to [**Entity**](Entity.md) |  | [optional] 
**Permission** | Pointer to **string** | Permission to be checked, can be a permission or relation. Required, and must match the pattern \&quot;^([a-zA-Z][a-zA-Z0-9_]{1,62}[a-zA-Z0-9])$\&quot;, max 64 bytes. | [optional] 
**SubjectReference** | Pointer to [**RelationReference**](RelationReference.md) |  | [optional] 
**Context** | Pointer to [**Context**](Context.md) |  | [optional] 

## Methods

### NewLookupSubjectBody

`func NewLookupSubjectBody() *LookupSubjectBody`

NewLookupSubjectBody instantiates a new LookupSubjectBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupSubjectBodyWithDefaults

`func NewLookupSubjectBodyWithDefaults() *LookupSubjectBody`

NewLookupSubjectBodyWithDefaults instantiates a new LookupSubjectBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *LookupSubjectBody) GetMetadata() PermissionLookupSubjectRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LookupSubjectBody) GetMetadataOk() (*PermissionLookupSubjectRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LookupSubjectBody) SetMetadata(v PermissionLookupSubjectRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LookupSubjectBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEntity

`func (o *LookupSubjectBody) GetEntity() Entity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *LookupSubjectBody) GetEntityOk() (*Entity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *LookupSubjectBody) SetEntity(v Entity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *LookupSubjectBody) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetPermission

`func (o *LookupSubjectBody) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *LookupSubjectBody) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *LookupSubjectBody) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *LookupSubjectBody) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetSubjectReference

`func (o *LookupSubjectBody) GetSubjectReference() RelationReference`

GetSubjectReference returns the SubjectReference field if non-nil, zero value otherwise.

### GetSubjectReferenceOk

`func (o *LookupSubjectBody) GetSubjectReferenceOk() (*RelationReference, bool)`

GetSubjectReferenceOk returns a tuple with the SubjectReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectReference

`func (o *LookupSubjectBody) SetSubjectReference(v RelationReference)`

SetSubjectReference sets SubjectReference field to given value.

### HasSubjectReference

`func (o *LookupSubjectBody) HasSubjectReference() bool`

HasSubjectReference returns a boolean if a field has been set.

### GetContext

`func (o *LookupSubjectBody) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *LookupSubjectBody) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *LookupSubjectBody) SetContext(v Context)`

SetContext sets Context field to given value.

### HasContext

`func (o *LookupSubjectBody) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


