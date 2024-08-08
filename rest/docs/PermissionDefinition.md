# PermissionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the permission, which follows a specific string pattern and has a maximum byte size. | [optional] 
**Child** | Pointer to [**Child**](Child.md) |  | [optional] 

## Methods

### NewPermissionDefinition

`func NewPermissionDefinition() *PermissionDefinition`

NewPermissionDefinition instantiates a new PermissionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionDefinitionWithDefaults

`func NewPermissionDefinitionWithDefaults() *PermissionDefinition`

NewPermissionDefinitionWithDefaults instantiates a new PermissionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PermissionDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PermissionDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetChild

`func (o *PermissionDefinition) GetChild() Child`

GetChild returns the Child field if non-nil, zero value otherwise.

### GetChildOk

`func (o *PermissionDefinition) GetChildOk() (*Child, bool)`

GetChildOk returns a tuple with the Child field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChild

`func (o *PermissionDefinition) SetChild(v Child)`

SetChild sets Child field to given value.

### HasChild

`func (o *PermissionDefinition) HasChild() bool`

HasChild returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


