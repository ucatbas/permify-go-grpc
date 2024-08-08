# WriteRelationshipsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**RelationshipWriteRequestMetadata**](RelationshipWriteRequestMetadata.md) |  | [optional] 
**Tuples** | Pointer to [**[]Tuple**](Tuple.md) | List of tuples for the request. Must have between 1 and 100 items. | [optional] 

## Methods

### NewWriteRelationshipsBody

`func NewWriteRelationshipsBody() *WriteRelationshipsBody`

NewWriteRelationshipsBody instantiates a new WriteRelationshipsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteRelationshipsBodyWithDefaults

`func NewWriteRelationshipsBodyWithDefaults() *WriteRelationshipsBody`

NewWriteRelationshipsBodyWithDefaults instantiates a new WriteRelationshipsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *WriteRelationshipsBody) GetMetadata() RelationshipWriteRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WriteRelationshipsBody) GetMetadataOk() (*RelationshipWriteRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WriteRelationshipsBody) SetMetadata(v RelationshipWriteRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WriteRelationshipsBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTuples

`func (o *WriteRelationshipsBody) GetTuples() []Tuple`

GetTuples returns the Tuples field if non-nil, zero value otherwise.

### GetTuplesOk

`func (o *WriteRelationshipsBody) GetTuplesOk() (*[]Tuple, bool)`

GetTuplesOk returns a tuple with the Tuples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuples

`func (o *WriteRelationshipsBody) SetTuples(v []Tuple)`

SetTuples sets Tuples field to given value.

### HasTuples

`func (o *WriteRelationshipsBody) HasTuples() bool`

HasTuples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


