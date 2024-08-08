# CheckBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**PermissionCheckRequestMetadata**](PermissionCheckRequestMetadata.md) |  | [optional] 
**Entity** | Pointer to [**Entity**](Entity.md) |  | [optional] 
**Permission** | Pointer to **string** | The action the user wants to perform on the resource | [optional] 
**Subject** | Pointer to [**Subject**](Subject.md) |  | [optional] 
**Context** | Pointer to [**Context**](Context.md) |  | [optional] 
**Arguments** | Pointer to [**[]Argument**](Argument.md) | Additional arguments associated with this request. | [optional] 

## Methods

### NewCheckBody

`func NewCheckBody() *CheckBody`

NewCheckBody instantiates a new CheckBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckBodyWithDefaults

`func NewCheckBodyWithDefaults() *CheckBody`

NewCheckBodyWithDefaults instantiates a new CheckBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *CheckBody) GetMetadata() PermissionCheckRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CheckBody) GetMetadataOk() (*PermissionCheckRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CheckBody) SetMetadata(v PermissionCheckRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CheckBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEntity

`func (o *CheckBody) GetEntity() Entity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *CheckBody) GetEntityOk() (*Entity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *CheckBody) SetEntity(v Entity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *CheckBody) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetPermission

`func (o *CheckBody) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *CheckBody) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *CheckBody) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *CheckBody) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetSubject

`func (o *CheckBody) GetSubject() Subject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CheckBody) GetSubjectOk() (*Subject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CheckBody) SetSubject(v Subject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CheckBody) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetContext

`func (o *CheckBody) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CheckBody) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CheckBody) SetContext(v Context)`

SetContext sets Context field to given value.

### HasContext

`func (o *CheckBody) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetArguments

`func (o *CheckBody) GetArguments() []Argument`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *CheckBody) GetArgumentsOk() (*[]Argument, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *CheckBody) SetArguments(v []Argument)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *CheckBody) HasArguments() bool`

HasArguments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


