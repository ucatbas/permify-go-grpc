# SubjectPermissionBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**PermissionSubjectPermissionRequestMetadata**](PermissionSubjectPermissionRequestMetadata.md) |  | [optional] 
**Entity** | Pointer to [**Entity**](Entity.md) |  | [optional] 
**Subject** | Pointer to [**Subject**](Subject.md) |  | [optional] 
**Context** | Pointer to [**Context**](Context.md) |  | [optional] 

## Methods

### NewSubjectPermissionBody

`func NewSubjectPermissionBody() *SubjectPermissionBody`

NewSubjectPermissionBody instantiates a new SubjectPermissionBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectPermissionBodyWithDefaults

`func NewSubjectPermissionBodyWithDefaults() *SubjectPermissionBody`

NewSubjectPermissionBodyWithDefaults instantiates a new SubjectPermissionBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *SubjectPermissionBody) GetMetadata() PermissionSubjectPermissionRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubjectPermissionBody) GetMetadataOk() (*PermissionSubjectPermissionRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubjectPermissionBody) SetMetadata(v PermissionSubjectPermissionRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubjectPermissionBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEntity

`func (o *SubjectPermissionBody) GetEntity() Entity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *SubjectPermissionBody) GetEntityOk() (*Entity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *SubjectPermissionBody) SetEntity(v Entity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *SubjectPermissionBody) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetSubject

`func (o *SubjectPermissionBody) GetSubject() Subject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SubjectPermissionBody) GetSubjectOk() (*Subject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SubjectPermissionBody) SetSubject(v Subject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SubjectPermissionBody) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetContext

`func (o *SubjectPermissionBody) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SubjectPermissionBody) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SubjectPermissionBody) SetContext(v Context)`

SetContext sets Context field to given value.

### HasContext

`func (o *SubjectPermissionBody) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


