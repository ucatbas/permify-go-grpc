# CreateList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elements** | Pointer to [**[]Expr**](Expr.md) | The elements part of the list. | [optional] 
**OptionalIndices** | Pointer to **[]int32** | The indices within the elements list which are marked as optional elements.  When an optional-typed value is present, the value it contains is included in the list. If the optional-typed value is absent, the list element is omitted from the CreateList result. | [optional] 

## Methods

### NewCreateList

`func NewCreateList() *CreateList`

NewCreateList instantiates a new CreateList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateListWithDefaults

`func NewCreateListWithDefaults() *CreateList`

NewCreateListWithDefaults instantiates a new CreateList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElements

`func (o *CreateList) GetElements() []Expr`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *CreateList) GetElementsOk() (*[]Expr, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *CreateList) SetElements(v []Expr)`

SetElements sets Elements field to given value.

### HasElements

`func (o *CreateList) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetOptionalIndices

`func (o *CreateList) GetOptionalIndices() []int32`

GetOptionalIndices returns the OptionalIndices field if non-nil, zero value otherwise.

### GetOptionalIndicesOk

`func (o *CreateList) GetOptionalIndicesOk() (*[]int32, bool)`

GetOptionalIndicesOk returns a tuple with the OptionalIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalIndices

`func (o *CreateList) SetOptionalIndices(v []int32)`

SetOptionalIndices sets OptionalIndices field to given value.

### HasOptionalIndices

`func (o *CreateList) HasOptionalIndices() bool`

HasOptionalIndices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


