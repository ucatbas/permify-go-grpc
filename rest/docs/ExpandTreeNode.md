# ExpandTreeNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to [**ExpandTreeNodeOperation**](ExpandTreeNodeOperation.md) |  | [optional] 
**Children** | Pointer to [**[]V1Expand**](V1Expand.md) |  | [optional] 

## Methods

### NewExpandTreeNode

`func NewExpandTreeNode() *ExpandTreeNode`

NewExpandTreeNode instantiates a new ExpandTreeNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandTreeNodeWithDefaults

`func NewExpandTreeNodeWithDefaults() *ExpandTreeNode`

NewExpandTreeNodeWithDefaults instantiates a new ExpandTreeNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *ExpandTreeNode) GetOperation() ExpandTreeNodeOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ExpandTreeNode) GetOperationOk() (*ExpandTreeNodeOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ExpandTreeNode) SetOperation(v ExpandTreeNodeOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *ExpandTreeNode) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetChildren

`func (o *ExpandTreeNode) GetChildren() []V1Expand`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ExpandTreeNode) GetChildrenOk() (*[]V1Expand, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ExpandTreeNode) SetChildren(v []V1Expand)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *ExpandTreeNode) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


