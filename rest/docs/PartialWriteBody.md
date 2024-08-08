# PartialWriteBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**SchemaPartialWriteRequestMetadata**](SchemaPartialWriteRequestMetadata.md) |  | [optional] 
**Partials** | Pointer to [**map[string]Partials**](Partials.md) |  | [optional] 

## Methods

### NewPartialWriteBody

`func NewPartialWriteBody() *PartialWriteBody`

NewPartialWriteBody instantiates a new PartialWriteBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialWriteBodyWithDefaults

`func NewPartialWriteBodyWithDefaults() *PartialWriteBody`

NewPartialWriteBodyWithDefaults instantiates a new PartialWriteBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *PartialWriteBody) GetMetadata() SchemaPartialWriteRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PartialWriteBody) GetMetadataOk() (*SchemaPartialWriteRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PartialWriteBody) SetMetadata(v SchemaPartialWriteRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PartialWriteBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPartials

`func (o *PartialWriteBody) GetPartials() map[string]Partials`

GetPartials returns the Partials field if non-nil, zero value otherwise.

### GetPartialsOk

`func (o *PartialWriteBody) GetPartialsOk() (*map[string]Partials, bool)`

GetPartialsOk returns a tuple with the Partials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartials

`func (o *PartialWriteBody) SetPartials(v map[string]Partials)`

SetPartials sets Partials field to given value.

### HasPartials

`func (o *PartialWriteBody) HasPartials() bool`

HasPartials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


