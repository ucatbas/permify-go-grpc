# ExprCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | Pointer to [**Expr**](Expr.md) |  | [optional] 
**Function** | Pointer to **string** | Required. The name of the function or method being called. | [optional] 
**Args** | Pointer to [**[]Expr**](Expr.md) | The arguments. | [optional] 

## Methods

### NewExprCall

`func NewExprCall() *ExprCall`

NewExprCall instantiates a new ExprCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExprCallWithDefaults

`func NewExprCallWithDefaults() *ExprCall`

NewExprCallWithDefaults instantiates a new ExprCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *ExprCall) GetTarget() Expr`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ExprCall) GetTargetOk() (*Expr, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ExprCall) SetTarget(v Expr)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ExprCall) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetFunction

`func (o *ExprCall) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ExprCall) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ExprCall) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *ExprCall) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetArgs

`func (o *ExprCall) GetArgs() []Expr`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ExprCall) GetArgsOk() (*[]Expr, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ExprCall) SetArgs(v []Expr)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *ExprCall) HasArgs() bool`

HasArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


