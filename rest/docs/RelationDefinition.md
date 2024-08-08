# RelationDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the relation, which follows a specific string pattern and has a maximum byte size. | [optional] 
**RelationReferences** | Pointer to [**[]RelationReference**](RelationReference.md) | A list of references to other relations. | [optional] 

## Methods

### NewRelationDefinition

`func NewRelationDefinition() *RelationDefinition`

NewRelationDefinition instantiates a new RelationDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationDefinitionWithDefaults

`func NewRelationDefinitionWithDefaults() *RelationDefinition`

NewRelationDefinitionWithDefaults instantiates a new RelationDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RelationDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RelationDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RelationDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RelationDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelationReferences

`func (o *RelationDefinition) GetRelationReferences() []RelationReference`

GetRelationReferences returns the RelationReferences field if non-nil, zero value otherwise.

### GetRelationReferencesOk

`func (o *RelationDefinition) GetRelationReferencesOk() (*[]RelationReference, bool)`

GetRelationReferencesOk returns a tuple with the RelationReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationReferences

`func (o *RelationDefinition) SetRelationReferences(v []RelationReference)`

SetRelationReferences sets RelationReferences field to given value.

### HasRelationReferences

`func (o *RelationDefinition) HasRelationReferences() bool`

HasRelationReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


