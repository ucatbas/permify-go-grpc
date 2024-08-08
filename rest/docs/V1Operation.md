# V1Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelationshipsWrite** | Pointer to **[]string** | &#39;relationships_write&#39; is a repeated string field for storing relationship keys that are to be written or created. | [optional] 
**RelationshipsDelete** | Pointer to **[]string** | &#39;relationships_delete&#39; is a repeated string field for storing relationship keys that are to be deleted or removed. | [optional] 
**AttributesWrite** | Pointer to **[]string** | &#39;attributes_write&#39; is a repeated string field for storing attribute keys that are to be written or created. | [optional] 
**AttributesDelete** | Pointer to **[]string** | &#39;attributes_delete&#39; is a repeated string field for storing attribute keys that are to be deleted or removed. | [optional] 

## Methods

### NewV1Operation

`func NewV1Operation() *V1Operation`

NewV1Operation instantiates a new V1Operation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1OperationWithDefaults

`func NewV1OperationWithDefaults() *V1Operation`

NewV1OperationWithDefaults instantiates a new V1Operation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelationshipsWrite

`func (o *V1Operation) GetRelationshipsWrite() []string`

GetRelationshipsWrite returns the RelationshipsWrite field if non-nil, zero value otherwise.

### GetRelationshipsWriteOk

`func (o *V1Operation) GetRelationshipsWriteOk() (*[]string, bool)`

GetRelationshipsWriteOk returns a tuple with the RelationshipsWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipsWrite

`func (o *V1Operation) SetRelationshipsWrite(v []string)`

SetRelationshipsWrite sets RelationshipsWrite field to given value.

### HasRelationshipsWrite

`func (o *V1Operation) HasRelationshipsWrite() bool`

HasRelationshipsWrite returns a boolean if a field has been set.

### GetRelationshipsDelete

`func (o *V1Operation) GetRelationshipsDelete() []string`

GetRelationshipsDelete returns the RelationshipsDelete field if non-nil, zero value otherwise.

### GetRelationshipsDeleteOk

`func (o *V1Operation) GetRelationshipsDeleteOk() (*[]string, bool)`

GetRelationshipsDeleteOk returns a tuple with the RelationshipsDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipsDelete

`func (o *V1Operation) SetRelationshipsDelete(v []string)`

SetRelationshipsDelete sets RelationshipsDelete field to given value.

### HasRelationshipsDelete

`func (o *V1Operation) HasRelationshipsDelete() bool`

HasRelationshipsDelete returns a boolean if a field has been set.

### GetAttributesWrite

`func (o *V1Operation) GetAttributesWrite() []string`

GetAttributesWrite returns the AttributesWrite field if non-nil, zero value otherwise.

### GetAttributesWriteOk

`func (o *V1Operation) GetAttributesWriteOk() (*[]string, bool)`

GetAttributesWriteOk returns a tuple with the AttributesWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesWrite

`func (o *V1Operation) SetAttributesWrite(v []string)`

SetAttributesWrite sets AttributesWrite field to given value.

### HasAttributesWrite

`func (o *V1Operation) HasAttributesWrite() bool`

HasAttributesWrite returns a boolean if a field has been set.

### GetAttributesDelete

`func (o *V1Operation) GetAttributesDelete() []string`

GetAttributesDelete returns the AttributesDelete field if non-nil, zero value otherwise.

### GetAttributesDeleteOk

`func (o *V1Operation) GetAttributesDeleteOk() (*[]string, bool)`

GetAttributesDeleteOk returns a tuple with the AttributesDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesDelete

`func (o *V1Operation) SetAttributesDelete(v []string)`

SetAttributesDelete sets AttributesDelete field to given value.

### HasAttributesDelete

`func (o *V1Operation) HasAttributesDelete() bool`

HasAttributesDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


