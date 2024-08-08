# SchemaDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityDefinitions** | Pointer to [**map[string]EntityDefinition**](EntityDefinition.md) | Map of entity definitions. The key is the entity name, and the value is the corresponding EntityDefinition. | [optional] 
**RuleDefinitions** | Pointer to [**map[string]RuleDefinition**](RuleDefinition.md) | Map of rule definitions. The key is the rule name, and the value is the corresponding RuleDefinition. | [optional] 
**References** | Pointer to [**map[string]SchemaDefinitionReference**](SchemaDefinitionReference.md) | Map of references to signify whether a string refers to an entity or a rule. | [optional] 

## Methods

### NewSchemaDefinition

`func NewSchemaDefinition() *SchemaDefinition`

NewSchemaDefinition instantiates a new SchemaDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaDefinitionWithDefaults

`func NewSchemaDefinitionWithDefaults() *SchemaDefinition`

NewSchemaDefinitionWithDefaults instantiates a new SchemaDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityDefinitions

`func (o *SchemaDefinition) GetEntityDefinitions() map[string]EntityDefinition`

GetEntityDefinitions returns the EntityDefinitions field if non-nil, zero value otherwise.

### GetEntityDefinitionsOk

`func (o *SchemaDefinition) GetEntityDefinitionsOk() (*map[string]EntityDefinition, bool)`

GetEntityDefinitionsOk returns a tuple with the EntityDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityDefinitions

`func (o *SchemaDefinition) SetEntityDefinitions(v map[string]EntityDefinition)`

SetEntityDefinitions sets EntityDefinitions field to given value.

### HasEntityDefinitions

`func (o *SchemaDefinition) HasEntityDefinitions() bool`

HasEntityDefinitions returns a boolean if a field has been set.

### GetRuleDefinitions

`func (o *SchemaDefinition) GetRuleDefinitions() map[string]RuleDefinition`

GetRuleDefinitions returns the RuleDefinitions field if non-nil, zero value otherwise.

### GetRuleDefinitionsOk

`func (o *SchemaDefinition) GetRuleDefinitionsOk() (*map[string]RuleDefinition, bool)`

GetRuleDefinitionsOk returns a tuple with the RuleDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDefinitions

`func (o *SchemaDefinition) SetRuleDefinitions(v map[string]RuleDefinition)`

SetRuleDefinitions sets RuleDefinitions field to given value.

### HasRuleDefinitions

`func (o *SchemaDefinition) HasRuleDefinitions() bool`

HasRuleDefinitions returns a boolean if a field has been set.

### GetReferences

`func (o *SchemaDefinition) GetReferences() map[string]SchemaDefinitionReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *SchemaDefinition) GetReferencesOk() (*map[string]SchemaDefinitionReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *SchemaDefinition) SetReferences(v map[string]SchemaDefinitionReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *SchemaDefinition) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


