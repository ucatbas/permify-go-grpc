# Rewrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RewriteOperation** | Pointer to [**RewriteOperation**](RewriteOperation.md) |  | [optional] 
**Children** | Pointer to [**[]Child**](Child.md) | A list of children that are operated upon by the rewrite operation. | [optional] 

## Methods

### NewRewrite

`func NewRewrite() *Rewrite`

NewRewrite instantiates a new Rewrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewriteWithDefaults

`func NewRewriteWithDefaults() *Rewrite`

NewRewriteWithDefaults instantiates a new Rewrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRewriteOperation

`func (o *Rewrite) GetRewriteOperation() RewriteOperation`

GetRewriteOperation returns the RewriteOperation field if non-nil, zero value otherwise.

### GetRewriteOperationOk

`func (o *Rewrite) GetRewriteOperationOk() (*RewriteOperation, bool)`

GetRewriteOperationOk returns a tuple with the RewriteOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewriteOperation

`func (o *Rewrite) SetRewriteOperation(v RewriteOperation)`

SetRewriteOperation sets RewriteOperation field to given value.

### HasRewriteOperation

`func (o *Rewrite) HasRewriteOperation() bool`

HasRewriteOperation returns a boolean if a field has been set.

### GetChildren

`func (o *Rewrite) GetChildren() []Child`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Rewrite) GetChildrenOk() (*[]Child, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Rewrite) SetChildren(v []Child)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Rewrite) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


