# CheckedExpr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceMap** | Pointer to [**map[string]V1alpha1Reference**](V1alpha1Reference.md) | A map from expression ids to resolved references.  The following entries are in this table:  - An Ident or Select expression is represented here if it resolves to a   declaration. For instance, if &#x60;a.b.c&#x60; is represented by   &#x60;select(select(id(a), b), c)&#x60;, and &#x60;a.b&#x60; resolves to a declaration,   while &#x60;c&#x60; is a field selection, then the reference is attached to the   nested select expression (but not to the id or or the outer select).   In turn, if &#x60;a&#x60; resolves to a declaration and &#x60;b.c&#x60; are field selections,   the reference is attached to the ident expression. - Every Call expression has an entry here, identifying the function being   called. - Every CreateStruct expression for a message has an entry, identifying   the message. | [optional] 
**TypeMap** | Pointer to [**map[string]V1alpha1Type**](V1alpha1Type.md) | A map from expression ids to types.  Every expression node which has a type different than DYN has a mapping here. If an expression has type DYN, it is omitted from this map to save space. | [optional] 
**SourceInfo** | Pointer to [**SourceInfo**](SourceInfo.md) |  | [optional] 
**ExprVersion** | Pointer to **string** | The expr version indicates the major / minor version number of the &#x60;expr&#x60; representation.  The most common reason for a version change will be to indicate to the CEL runtimes that transformations have been performed on the expr during static analysis. In some cases, this will save the runtime the work of applying the same or similar transformations prior to evaluation. | [optional] 
**Expr** | Pointer to [**Expr**](Expr.md) |  | [optional] 

## Methods

### NewCheckedExpr

`func NewCheckedExpr() *CheckedExpr`

NewCheckedExpr instantiates a new CheckedExpr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckedExprWithDefaults

`func NewCheckedExprWithDefaults() *CheckedExpr`

NewCheckedExprWithDefaults instantiates a new CheckedExpr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceMap

`func (o *CheckedExpr) GetReferenceMap() map[string]V1alpha1Reference`

GetReferenceMap returns the ReferenceMap field if non-nil, zero value otherwise.

### GetReferenceMapOk

`func (o *CheckedExpr) GetReferenceMapOk() (*map[string]V1alpha1Reference, bool)`

GetReferenceMapOk returns a tuple with the ReferenceMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMap

`func (o *CheckedExpr) SetReferenceMap(v map[string]V1alpha1Reference)`

SetReferenceMap sets ReferenceMap field to given value.

### HasReferenceMap

`func (o *CheckedExpr) HasReferenceMap() bool`

HasReferenceMap returns a boolean if a field has been set.

### GetTypeMap

`func (o *CheckedExpr) GetTypeMap() map[string]V1alpha1Type`

GetTypeMap returns the TypeMap field if non-nil, zero value otherwise.

### GetTypeMapOk

`func (o *CheckedExpr) GetTypeMapOk() (*map[string]V1alpha1Type, bool)`

GetTypeMapOk returns a tuple with the TypeMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeMap

`func (o *CheckedExpr) SetTypeMap(v map[string]V1alpha1Type)`

SetTypeMap sets TypeMap field to given value.

### HasTypeMap

`func (o *CheckedExpr) HasTypeMap() bool`

HasTypeMap returns a boolean if a field has been set.

### GetSourceInfo

`func (o *CheckedExpr) GetSourceInfo() SourceInfo`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *CheckedExpr) GetSourceInfoOk() (*SourceInfo, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *CheckedExpr) SetSourceInfo(v SourceInfo)`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *CheckedExpr) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.

### GetExprVersion

`func (o *CheckedExpr) GetExprVersion() string`

GetExprVersion returns the ExprVersion field if non-nil, zero value otherwise.

### GetExprVersionOk

`func (o *CheckedExpr) GetExprVersionOk() (*string, bool)`

GetExprVersionOk returns a tuple with the ExprVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExprVersion

`func (o *CheckedExpr) SetExprVersion(v string)`

SetExprVersion sets ExprVersion field to given value.

### HasExprVersion

`func (o *CheckedExpr) HasExprVersion() bool`

HasExprVersion returns a boolean if a field has been set.

### GetExpr

`func (o *CheckedExpr) GetExpr() Expr`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *CheckedExpr) GetExprOk() (*Expr, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpr

`func (o *CheckedExpr) SetExpr(v Expr)`

SetExpr sets Expr field to given value.

### HasExpr

`func (o *CheckedExpr) HasExpr() bool`

HasExpr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


