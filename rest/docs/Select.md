# Select

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operand** | Pointer to [**Expr**](Expr.md) |  | [optional] 
**Field** | Pointer to **string** | Required. The name of the field to select.  For example, in the select expression &#x60;request.auth&#x60;, the &#x60;auth&#x60; portion of the expression would be the &#x60;field&#x60;. | [optional] 
**TestOnly** | Pointer to **bool** | Whether the select is to be interpreted as a field presence test.  This results from the macro &#x60;has(request.auth)&#x60;. | [optional] 

## Methods

### NewSelect

`func NewSelect() *Select`

NewSelect instantiates a new Select object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectWithDefaults

`func NewSelectWithDefaults() *Select`

NewSelectWithDefaults instantiates a new Select object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperand

`func (o *Select) GetOperand() Expr`

GetOperand returns the Operand field if non-nil, zero value otherwise.

### GetOperandOk

`func (o *Select) GetOperandOk() (*Expr, bool)`

GetOperandOk returns a tuple with the Operand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand

`func (o *Select) SetOperand(v Expr)`

SetOperand sets Operand field to given value.

### HasOperand

`func (o *Select) HasOperand() bool`

HasOperand returns a boolean if a field has been set.

### GetField

`func (o *Select) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *Select) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *Select) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *Select) HasField() bool`

HasField returns a boolean if a field has been set.

### GetTestOnly

`func (o *Select) GetTestOnly() bool`

GetTestOnly returns the TestOnly field if non-nil, zero value otherwise.

### GetTestOnlyOk

`func (o *Select) GetTestOnlyOk() (*bool, bool)`

GetTestOnlyOk returns a tuple with the TestOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestOnly

`func (o *Select) SetTestOnly(v bool)`

SetTestOnly sets TestOnly field to given value.

### HasTestOnly

`func (o *Select) HasTestOnly() bool`

HasTestOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


