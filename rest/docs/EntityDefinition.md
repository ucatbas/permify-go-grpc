# EntityDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the entity, which follows a specific string pattern and has a maximum byte size. | [optional] 
**Relations** | Pointer to [**map[string]RelationDefinition**](RelationDefinition.md) | Map of relation definitions within this entity. The key is the relation name, and the value is the RelationDefinition. | [optional] 
**Permissions** | Pointer to [**map[string]PermissionDefinition**](PermissionDefinition.md) | Map of permission definitions within this entity. The key is the permission name, and the value is the PermissionDefinition. | [optional] 
**Attributes** | Pointer to [**map[string]AttributeDefinition**](AttributeDefinition.md) | Map of attribute definitions within this entity. The key is the attribute name, and the value is the AttributeDefinition. | [optional] 
**References** | Pointer to [**map[string]EntityDefinitionReference**](EntityDefinitionReference.md) | Map of references indicating whether a string pertains to a relation, permission, or attribute. | [optional] 

## Methods

### NewEntityDefinition

`func NewEntityDefinition() *EntityDefinition`

NewEntityDefinition instantiates a new EntityDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityDefinitionWithDefaults

`func NewEntityDefinitionWithDefaults() *EntityDefinition`

NewEntityDefinitionWithDefaults instantiates a new EntityDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EntityDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntityDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelations

`func (o *EntityDefinition) GetRelations() map[string]RelationDefinition`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *EntityDefinition) GetRelationsOk() (*map[string]RelationDefinition, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *EntityDefinition) SetRelations(v map[string]RelationDefinition)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *EntityDefinition) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### GetPermissions

`func (o *EntityDefinition) GetPermissions() map[string]PermissionDefinition`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *EntityDefinition) GetPermissionsOk() (*map[string]PermissionDefinition, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *EntityDefinition) SetPermissions(v map[string]PermissionDefinition)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *EntityDefinition) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetAttributes

`func (o *EntityDefinition) GetAttributes() map[string]AttributeDefinition`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EntityDefinition) GetAttributesOk() (*map[string]AttributeDefinition, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EntityDefinition) SetAttributes(v map[string]AttributeDefinition)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *EntityDefinition) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetReferences

`func (o *EntityDefinition) GetReferences() map[string]EntityDefinitionReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *EntityDefinition) GetReferencesOk() (*map[string]EntityDefinitionReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *EntityDefinition) SetReferences(v map[string]EntityDefinitionReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *EntityDefinition) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


