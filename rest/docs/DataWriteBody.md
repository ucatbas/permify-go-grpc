# DataWriteBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**DataWriteRequestMetadata**](DataWriteRequestMetadata.md) |  | [optional] 
**Tuples** | Pointer to [**[]Tuple**](Tuple.md) | tuples contains the list of tuples (entity-relation-entity triples) that need to be written. | [optional] 
**Attributes** | Pointer to [**[]Attribute**](Attribute.md) | attributes contains the list of attributes (entity-attribute-value triples) that need to be written. | [optional] 

## Methods

### NewDataWriteBody

`func NewDataWriteBody() *DataWriteBody`

NewDataWriteBody instantiates a new DataWriteBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataWriteBodyWithDefaults

`func NewDataWriteBodyWithDefaults() *DataWriteBody`

NewDataWriteBodyWithDefaults instantiates a new DataWriteBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *DataWriteBody) GetMetadata() DataWriteRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DataWriteBody) GetMetadataOk() (*DataWriteRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DataWriteBody) SetMetadata(v DataWriteRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DataWriteBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTuples

`func (o *DataWriteBody) GetTuples() []Tuple`

GetTuples returns the Tuples field if non-nil, zero value otherwise.

### GetTuplesOk

`func (o *DataWriteBody) GetTuplesOk() (*[]Tuple, bool)`

GetTuplesOk returns a tuple with the Tuples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuples

`func (o *DataWriteBody) SetTuples(v []Tuple)`

SetTuples sets Tuples field to given value.

### HasTuples

`func (o *DataWriteBody) HasTuples() bool`

HasTuples returns a boolean if a field has been set.

### GetAttributes

`func (o *DataWriteBody) GetAttributes() []Attribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DataWriteBody) GetAttributesOk() (*[]Attribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DataWriteBody) SetAttributes(v []Attribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DataWriteBody) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


