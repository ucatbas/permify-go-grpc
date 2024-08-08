# AttributeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to [**EntityFilter**](EntityFilter.md) |  | [optional] 
**Attributes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAttributeFilter

`func NewAttributeFilter() *AttributeFilter`

NewAttributeFilter instantiates a new AttributeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeFilterWithDefaults

`func NewAttributeFilterWithDefaults() *AttributeFilter`

NewAttributeFilterWithDefaults instantiates a new AttributeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *AttributeFilter) GetEntity() EntityFilter`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *AttributeFilter) GetEntityOk() (*EntityFilter, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *AttributeFilter) SetEntity(v EntityFilter)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *AttributeFilter) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetAttributes

`func (o *AttributeFilter) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AttributeFilter) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AttributeFilter) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AttributeFilter) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


