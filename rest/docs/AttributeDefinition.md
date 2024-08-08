# AttributeDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the attribute, which follows a specific string pattern and has a maximum byte size. | [optional] 
**Type** | Pointer to [**AttributeType**](AttributeType.md) |  | [optional] 

## Methods

### NewAttributeDefinition

`func NewAttributeDefinition() *AttributeDefinition`

NewAttributeDefinition instantiates a new AttributeDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeDefinitionWithDefaults

`func NewAttributeDefinitionWithDefaults() *AttributeDefinition`

NewAttributeDefinitionWithDefaults instantiates a new AttributeDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AttributeDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttributeDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttributeDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AttributeDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AttributeDefinition) GetType() AttributeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttributeDefinition) GetTypeOk() (*AttributeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttributeDefinition) SetType(v AttributeType)`

SetType sets Type field to given value.

### HasType

`func (o *AttributeDefinition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


