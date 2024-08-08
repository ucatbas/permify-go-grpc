# Expr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Required. An id assigned to this node by the parser which is unique in a given expression tree. This is used to associate type information and other attributes to a node in the parse tree. | [optional] 
**ConstExpr** | Pointer to [**Constant**](Constant.md) |  | [optional] 
**IdentExpr** | Pointer to [**Ident**](Ident.md) |  | [optional] 
**SelectExpr** | Pointer to [**Select**](Select.md) |  | [optional] 
**CallExpr** | Pointer to [**ExprCall**](ExprCall.md) |  | [optional] 
**ListExpr** | Pointer to [**CreateList**](CreateList.md) |  | [optional] 
**StructExpr** | Pointer to [**CreateStruct**](CreateStruct.md) |  | [optional] 
**ComprehensionExpr** | Pointer to [**Comprehension**](Comprehension.md) |  | [optional] 

## Methods

### NewExpr

`func NewExpr() *Expr`

NewExpr instantiates a new Expr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExprWithDefaults

`func NewExprWithDefaults() *Expr`

NewExprWithDefaults instantiates a new Expr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Expr) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Expr) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Expr) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Expr) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConstExpr

`func (o *Expr) GetConstExpr() Constant`

GetConstExpr returns the ConstExpr field if non-nil, zero value otherwise.

### GetConstExprOk

`func (o *Expr) GetConstExprOk() (*Constant, bool)`

GetConstExprOk returns a tuple with the ConstExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstExpr

`func (o *Expr) SetConstExpr(v Constant)`

SetConstExpr sets ConstExpr field to given value.

### HasConstExpr

`func (o *Expr) HasConstExpr() bool`

HasConstExpr returns a boolean if a field has been set.

### GetIdentExpr

`func (o *Expr) GetIdentExpr() Ident`

GetIdentExpr returns the IdentExpr field if non-nil, zero value otherwise.

### GetIdentExprOk

`func (o *Expr) GetIdentExprOk() (*Ident, bool)`

GetIdentExprOk returns a tuple with the IdentExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentExpr

`func (o *Expr) SetIdentExpr(v Ident)`

SetIdentExpr sets IdentExpr field to given value.

### HasIdentExpr

`func (o *Expr) HasIdentExpr() bool`

HasIdentExpr returns a boolean if a field has been set.

### GetSelectExpr

`func (o *Expr) GetSelectExpr() Select`

GetSelectExpr returns the SelectExpr field if non-nil, zero value otherwise.

### GetSelectExprOk

`func (o *Expr) GetSelectExprOk() (*Select, bool)`

GetSelectExprOk returns a tuple with the SelectExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectExpr

`func (o *Expr) SetSelectExpr(v Select)`

SetSelectExpr sets SelectExpr field to given value.

### HasSelectExpr

`func (o *Expr) HasSelectExpr() bool`

HasSelectExpr returns a boolean if a field has been set.

### GetCallExpr

`func (o *Expr) GetCallExpr() ExprCall`

GetCallExpr returns the CallExpr field if non-nil, zero value otherwise.

### GetCallExprOk

`func (o *Expr) GetCallExprOk() (*ExprCall, bool)`

GetCallExprOk returns a tuple with the CallExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallExpr

`func (o *Expr) SetCallExpr(v ExprCall)`

SetCallExpr sets CallExpr field to given value.

### HasCallExpr

`func (o *Expr) HasCallExpr() bool`

HasCallExpr returns a boolean if a field has been set.

### GetListExpr

`func (o *Expr) GetListExpr() CreateList`

GetListExpr returns the ListExpr field if non-nil, zero value otherwise.

### GetListExprOk

`func (o *Expr) GetListExprOk() (*CreateList, bool)`

GetListExprOk returns a tuple with the ListExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListExpr

`func (o *Expr) SetListExpr(v CreateList)`

SetListExpr sets ListExpr field to given value.

### HasListExpr

`func (o *Expr) HasListExpr() bool`

HasListExpr returns a boolean if a field has been set.

### GetStructExpr

`func (o *Expr) GetStructExpr() CreateStruct`

GetStructExpr returns the StructExpr field if non-nil, zero value otherwise.

### GetStructExprOk

`func (o *Expr) GetStructExprOk() (*CreateStruct, bool)`

GetStructExprOk returns a tuple with the StructExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructExpr

`func (o *Expr) SetStructExpr(v CreateStruct)`

SetStructExpr sets StructExpr field to given value.

### HasStructExpr

`func (o *Expr) HasStructExpr() bool`

HasStructExpr returns a boolean if a field has been set.

### GetComprehensionExpr

`func (o *Expr) GetComprehensionExpr() Comprehension`

GetComprehensionExpr returns the ComprehensionExpr field if non-nil, zero value otherwise.

### GetComprehensionExprOk

`func (o *Expr) GetComprehensionExprOk() (*Comprehension, bool)`

GetComprehensionExprOk returns a tuple with the ComprehensionExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComprehensionExpr

`func (o *Expr) SetComprehensionExpr(v Comprehension)`

SetComprehensionExpr sets ComprehensionExpr field to given value.

### HasComprehensionExpr

`func (o *Expr) HasComprehensionExpr() bool`

HasComprehensionExpr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


