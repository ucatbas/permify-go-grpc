# Tuple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to [**Entity**](Entity.md) |  | [optional] 
**Relation** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to [**Subject**](Subject.md) |  | [optional] 

## Methods

### NewTuple

`func NewTuple() *Tuple`

NewTuple instantiates a new Tuple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTupleWithDefaults

`func NewTupleWithDefaults() *Tuple`

NewTupleWithDefaults instantiates a new Tuple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *Tuple) GetEntity() Entity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *Tuple) GetEntityOk() (*Entity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *Tuple) SetEntity(v Entity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *Tuple) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetRelation

`func (o *Tuple) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *Tuple) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *Tuple) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *Tuple) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetSubject

`func (o *Tuple) GetSubject() Subject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Tuple) GetSubjectOk() (*Subject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Tuple) SetSubject(v Subject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Tuple) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


