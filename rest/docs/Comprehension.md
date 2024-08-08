# Comprehension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IterVar** | Pointer to **string** | The name of the iteration variable. | [optional] 
**IterRange** | Pointer to [**Expr**](Expr.md) |  | [optional] 
**AccuVar** | Pointer to **string** | The name of the variable used for accumulation of the result. | [optional] 
**AccuInit** | Pointer to [**Expr**](Expr.md) |  | [optional] 
**LoopCondition** | Pointer to [**Expr**](Expr.md) |  | [optional] 
**LoopStep** | Pointer to [**Expr**](Expr.md) |  | [optional] 
**Result** | Pointer to [**Expr**](Expr.md) |  | [optional] 

## Methods

### NewComprehension

`func NewComprehension() *Comprehension`

NewComprehension instantiates a new Comprehension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComprehensionWithDefaults

`func NewComprehensionWithDefaults() *Comprehension`

NewComprehensionWithDefaults instantiates a new Comprehension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIterVar

`func (o *Comprehension) GetIterVar() string`

GetIterVar returns the IterVar field if non-nil, zero value otherwise.

### GetIterVarOk

`func (o *Comprehension) GetIterVarOk() (*string, bool)`

GetIterVarOk returns a tuple with the IterVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterVar

`func (o *Comprehension) SetIterVar(v string)`

SetIterVar sets IterVar field to given value.

### HasIterVar

`func (o *Comprehension) HasIterVar() bool`

HasIterVar returns a boolean if a field has been set.

### GetIterRange

`func (o *Comprehension) GetIterRange() Expr`

GetIterRange returns the IterRange field if non-nil, zero value otherwise.

### GetIterRangeOk

`func (o *Comprehension) GetIterRangeOk() (*Expr, bool)`

GetIterRangeOk returns a tuple with the IterRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterRange

`func (o *Comprehension) SetIterRange(v Expr)`

SetIterRange sets IterRange field to given value.

### HasIterRange

`func (o *Comprehension) HasIterRange() bool`

HasIterRange returns a boolean if a field has been set.

### GetAccuVar

`func (o *Comprehension) GetAccuVar() string`

GetAccuVar returns the AccuVar field if non-nil, zero value otherwise.

### GetAccuVarOk

`func (o *Comprehension) GetAccuVarOk() (*string, bool)`

GetAccuVarOk returns a tuple with the AccuVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuVar

`func (o *Comprehension) SetAccuVar(v string)`

SetAccuVar sets AccuVar field to given value.

### HasAccuVar

`func (o *Comprehension) HasAccuVar() bool`

HasAccuVar returns a boolean if a field has been set.

### GetAccuInit

`func (o *Comprehension) GetAccuInit() Expr`

GetAccuInit returns the AccuInit field if non-nil, zero value otherwise.

### GetAccuInitOk

`func (o *Comprehension) GetAccuInitOk() (*Expr, bool)`

GetAccuInitOk returns a tuple with the AccuInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuInit

`func (o *Comprehension) SetAccuInit(v Expr)`

SetAccuInit sets AccuInit field to given value.

### HasAccuInit

`func (o *Comprehension) HasAccuInit() bool`

HasAccuInit returns a boolean if a field has been set.

### GetLoopCondition

`func (o *Comprehension) GetLoopCondition() Expr`

GetLoopCondition returns the LoopCondition field if non-nil, zero value otherwise.

### GetLoopConditionOk

`func (o *Comprehension) GetLoopConditionOk() (*Expr, bool)`

GetLoopConditionOk returns a tuple with the LoopCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopCondition

`func (o *Comprehension) SetLoopCondition(v Expr)`

SetLoopCondition sets LoopCondition field to given value.

### HasLoopCondition

`func (o *Comprehension) HasLoopCondition() bool`

HasLoopCondition returns a boolean if a field has been set.

### GetLoopStep

`func (o *Comprehension) GetLoopStep() Expr`

GetLoopStep returns the LoopStep field if non-nil, zero value otherwise.

### GetLoopStepOk

`func (o *Comprehension) GetLoopStepOk() (*Expr, bool)`

GetLoopStepOk returns a tuple with the LoopStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopStep

`func (o *Comprehension) SetLoopStep(v Expr)`

SetLoopStep sets LoopStep field to given value.

### HasLoopStep

`func (o *Comprehension) HasLoopStep() bool`

HasLoopStep returns a boolean if a field has been set.

### GetResult

`func (o *Comprehension) GetResult() Expr`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Comprehension) GetResultOk() (*Expr, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Comprehension) SetResult(v Expr)`

SetResult sets Result field to given value.

### HasResult

`func (o *Comprehension) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


