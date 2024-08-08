# Child

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Leaf** | Pointer to [**Leaf**](Leaf.md) |  | [optional] 
**Rewrite** | Pointer to [**Rewrite**](Rewrite.md) |  | [optional] 

## Methods

### NewChild

`func NewChild() *Child`

NewChild instantiates a new Child object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChildWithDefaults

`func NewChildWithDefaults() *Child`

NewChildWithDefaults instantiates a new Child object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeaf

`func (o *Child) GetLeaf() Leaf`

GetLeaf returns the Leaf field if non-nil, zero value otherwise.

### GetLeafOk

`func (o *Child) GetLeafOk() (*Leaf, bool)`

GetLeafOk returns a tuple with the Leaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaf

`func (o *Child) SetLeaf(v Leaf)`

SetLeaf sets Leaf field to given value.

### HasLeaf

`func (o *Child) HasLeaf() bool`

HasLeaf returns a boolean if a field has been set.

### GetRewrite

`func (o *Child) GetRewrite() Rewrite`

GetRewrite returns the Rewrite field if non-nil, zero value otherwise.

### GetRewriteOk

`func (o *Child) GetRewriteOk() (*Rewrite, bool)`

GetRewriteOk returns a tuple with the Rewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewrite

`func (o *Child) SetRewrite(v Rewrite)`

SetRewrite sets Rewrite field to given value.

### HasRewrite

`func (o *Child) HasRewrite() bool`

HasRewrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


