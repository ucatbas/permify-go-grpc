# RuleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the rule, which follows a specific string pattern and has a maximum byte size. | [optional] 
**Arguments** | Pointer to [**map[string]AttributeType**](AttributeType.md) | Map of arguments for this rule. The key is the attribute name, and the value is the AttributeType. | [optional] 
**Expression** | Pointer to [**CheckedExpr**](CheckedExpr.md) |  | [optional] 

## Methods

### NewRuleDefinition

`func NewRuleDefinition() *RuleDefinition`

NewRuleDefinition instantiates a new RuleDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleDefinitionWithDefaults

`func NewRuleDefinitionWithDefaults() *RuleDefinition`

NewRuleDefinitionWithDefaults instantiates a new RuleDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RuleDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArguments

`func (o *RuleDefinition) GetArguments() map[string]AttributeType`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *RuleDefinition) GetArgumentsOk() (*map[string]AttributeType, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *RuleDefinition) SetArguments(v map[string]AttributeType)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *RuleDefinition) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetExpression

`func (o *RuleDefinition) GetExpression() CheckedExpr`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *RuleDefinition) GetExpressionOk() (*CheckedExpr, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *RuleDefinition) SetExpression(v CheckedExpr)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *RuleDefinition) HasExpression() bool`

HasExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


