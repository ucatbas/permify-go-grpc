# DataDeleteBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TupleFilter** | Pointer to [**TupleFilter**](TupleFilter.md) |  | [optional] 
**AttributeFilter** | Pointer to [**AttributeFilter**](AttributeFilter.md) |  | [optional] 

## Methods

### NewDataDeleteBody

`func NewDataDeleteBody() *DataDeleteBody`

NewDataDeleteBody instantiates a new DataDeleteBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataDeleteBodyWithDefaults

`func NewDataDeleteBodyWithDefaults() *DataDeleteBody`

NewDataDeleteBodyWithDefaults instantiates a new DataDeleteBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTupleFilter

`func (o *DataDeleteBody) GetTupleFilter() TupleFilter`

GetTupleFilter returns the TupleFilter field if non-nil, zero value otherwise.

### GetTupleFilterOk

`func (o *DataDeleteBody) GetTupleFilterOk() (*TupleFilter, bool)`

GetTupleFilterOk returns a tuple with the TupleFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleFilter

`func (o *DataDeleteBody) SetTupleFilter(v TupleFilter)`

SetTupleFilter sets TupleFilter field to given value.

### HasTupleFilter

`func (o *DataDeleteBody) HasTupleFilter() bool`

HasTupleFilter returns a boolean if a field has been set.

### GetAttributeFilter

`func (o *DataDeleteBody) GetAttributeFilter() AttributeFilter`

GetAttributeFilter returns the AttributeFilter field if non-nil, zero value otherwise.

### GetAttributeFilterOk

`func (o *DataDeleteBody) GetAttributeFilterOk() (*AttributeFilter, bool)`

GetAttributeFilterOk returns a tuple with the AttributeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeFilter

`func (o *DataDeleteBody) SetAttributeFilter(v AttributeFilter)`

SetAttributeFilter sets AttributeFilter field to given value.

### HasAttributeFilter

`func (o *DataDeleteBody) HasAttributeFilter() bool`

HasAttributeFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


