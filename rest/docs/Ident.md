# Ident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Required. Holds a single, unqualified identifier, possibly preceded by a &#39;.&#39;.  Qualified names are represented by the [Expr.Select][google.api.expr.v1alpha1.Expr.Select] expression. | [optional] 

## Methods

### NewIdent

`func NewIdent() *Ident`

NewIdent instantiates a new Ident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentWithDefaults

`func NewIdentWithDefaults() *Ident`

NewIdentWithDefaults instantiates a new Ident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Ident) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ident) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ident) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Ident) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


