# Attribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to [**Entity**](Entity.md) |  | [optional] 
**Attribute** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**Any**](Any.md) |  | [optional] 

## Methods

### NewAttribute

`func NewAttribute() *Attribute`

NewAttribute instantiates a new Attribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeWithDefaults

`func NewAttributeWithDefaults() *Attribute`

NewAttributeWithDefaults instantiates a new Attribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *Attribute) GetEntity() Entity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *Attribute) GetEntityOk() (*Entity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *Attribute) SetEntity(v Entity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *Attribute) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetAttribute

`func (o *Attribute) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *Attribute) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *Attribute) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *Attribute) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetValue

`func (o *Attribute) GetValue() Any`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Attribute) GetValueOk() (*Any, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Attribute) SetValue(v Any)`

SetValue sets Value field to given value.

### HasValue

`func (o *Attribute) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


