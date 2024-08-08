# V1Expand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to [**Entity**](Entity.md) |  | [optional] 
**Permission** | Pointer to **string** | permission is the permission applied to the entity. | [optional] 
**Arguments** | Pointer to [**[]Argument**](Argument.md) | arguments are the additional information or context used to evaluate permissions. | [optional] 
**Expand** | Pointer to [**ExpandTreeNode**](ExpandTreeNode.md) |  | [optional] 
**Leaf** | Pointer to [**ExpandLeaf**](ExpandLeaf.md) |  | [optional] 

## Methods

### NewV1Expand

`func NewV1Expand() *V1Expand`

NewV1Expand instantiates a new V1Expand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExpandWithDefaults

`func NewV1ExpandWithDefaults() *V1Expand`

NewV1ExpandWithDefaults instantiates a new V1Expand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *V1Expand) GetEntity() Entity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *V1Expand) GetEntityOk() (*Entity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *V1Expand) SetEntity(v Entity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *V1Expand) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetPermission

`func (o *V1Expand) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *V1Expand) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *V1Expand) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *V1Expand) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetArguments

`func (o *V1Expand) GetArguments() []Argument`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *V1Expand) GetArgumentsOk() (*[]Argument, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *V1Expand) SetArguments(v []Argument)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *V1Expand) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetExpand

`func (o *V1Expand) GetExpand() ExpandTreeNode`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *V1Expand) GetExpandOk() (*ExpandTreeNode, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *V1Expand) SetExpand(v ExpandTreeNode)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *V1Expand) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetLeaf

`func (o *V1Expand) GetLeaf() ExpandLeaf`

GetLeaf returns the Leaf field if non-nil, zero value otherwise.

### GetLeafOk

`func (o *V1Expand) GetLeafOk() (*ExpandLeaf, bool)`

GetLeafOk returns a tuple with the Leaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaf

`func (o *V1Expand) SetLeaf(v ExpandLeaf)`

SetLeaf sets Leaf field to given value.

### HasLeaf

`func (o *V1Expand) HasLeaf() bool`

HasLeaf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


