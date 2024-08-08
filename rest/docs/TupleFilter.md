# TupleFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to [**EntityFilter**](EntityFilter.md) |  | [optional] 
**Relation** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to [**SubjectFilter**](SubjectFilter.md) |  | [optional] 

## Methods

### NewTupleFilter

`func NewTupleFilter() *TupleFilter`

NewTupleFilter instantiates a new TupleFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTupleFilterWithDefaults

`func NewTupleFilterWithDefaults() *TupleFilter`

NewTupleFilterWithDefaults instantiates a new TupleFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *TupleFilter) GetEntity() EntityFilter`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *TupleFilter) GetEntityOk() (*EntityFilter, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *TupleFilter) SetEntity(v EntityFilter)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *TupleFilter) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetRelation

`func (o *TupleFilter) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *TupleFilter) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *TupleFilter) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *TupleFilter) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetSubject

`func (o *TupleFilter) GetSubject() SubjectFilter`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TupleFilter) GetSubjectOk() (*SubjectFilter, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TupleFilter) SetSubject(v SubjectFilter)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TupleFilter) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


