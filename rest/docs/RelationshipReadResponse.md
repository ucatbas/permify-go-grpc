# RelationshipReadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tuples** | Pointer to [**[]Tuple**](Tuple.md) | tuples is a list of the relationships retrieved in the read operation, represented as entity-relation-entity triples. | [optional] 
**ContinuousToken** | Pointer to **string** | continuous_token is used in the case of paginated reads to retrieve the next page of results. | [optional] 

## Methods

### NewRelationshipReadResponse

`func NewRelationshipReadResponse() *RelationshipReadResponse`

NewRelationshipReadResponse instantiates a new RelationshipReadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipReadResponseWithDefaults

`func NewRelationshipReadResponseWithDefaults() *RelationshipReadResponse`

NewRelationshipReadResponseWithDefaults instantiates a new RelationshipReadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTuples

`func (o *RelationshipReadResponse) GetTuples() []Tuple`

GetTuples returns the Tuples field if non-nil, zero value otherwise.

### GetTuplesOk

`func (o *RelationshipReadResponse) GetTuplesOk() (*[]Tuple, bool)`

GetTuplesOk returns a tuple with the Tuples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuples

`func (o *RelationshipReadResponse) SetTuples(v []Tuple)`

SetTuples sets Tuples field to given value.

### HasTuples

`func (o *RelationshipReadResponse) HasTuples() bool`

HasTuples returns a boolean if a field has been set.

### GetContinuousToken

`func (o *RelationshipReadResponse) GetContinuousToken() string`

GetContinuousToken returns the ContinuousToken field if non-nil, zero value otherwise.

### GetContinuousTokenOk

`func (o *RelationshipReadResponse) GetContinuousTokenOk() (*string, bool)`

GetContinuousTokenOk returns a tuple with the ContinuousToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousToken

`func (o *RelationshipReadResponse) SetContinuousToken(v string)`

SetContinuousToken sets ContinuousToken field to given value.

### HasContinuousToken

`func (o *RelationshipReadResponse) HasContinuousToken() bool`

HasContinuousToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


