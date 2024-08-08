# V1alpha1Reference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The fully qualified name of the declaration. | [optional] 
**OverloadId** | Pointer to **[]string** | For references to functions, this is a list of &#x60;Overload.overload_id&#x60; values which match according to typing rules.  If the list has more than one element, overload resolution among the presented candidates must happen at runtime because of dynamic types. The type checker attempts to narrow down this list as much as possible.  Empty if this is not a reference to a [Decl.FunctionDecl][google.api.expr.v1alpha1.Decl.FunctionDecl]. | [optional] 
**Value** | Pointer to [**Constant**](Constant.md) |  | [optional] 

## Methods

### NewV1alpha1Reference

`func NewV1alpha1Reference() *V1alpha1Reference`

NewV1alpha1Reference instantiates a new V1alpha1Reference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ReferenceWithDefaults

`func NewV1alpha1ReferenceWithDefaults() *V1alpha1Reference`

NewV1alpha1ReferenceWithDefaults instantiates a new V1alpha1Reference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1alpha1Reference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1alpha1Reference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1alpha1Reference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1alpha1Reference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverloadId

`func (o *V1alpha1Reference) GetOverloadId() []string`

GetOverloadId returns the OverloadId field if non-nil, zero value otherwise.

### GetOverloadIdOk

`func (o *V1alpha1Reference) GetOverloadIdOk() (*[]string, bool)`

GetOverloadIdOk returns a tuple with the OverloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverloadId

`func (o *V1alpha1Reference) SetOverloadId(v []string)`

SetOverloadId sets OverloadId field to given value.

### HasOverloadId

`func (o *V1alpha1Reference) HasOverloadId() bool`

HasOverloadId returns a boolean if a field has been set.

### GetValue

`func (o *V1alpha1Reference) GetValue() Constant`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V1alpha1Reference) GetValueOk() (*Constant, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V1alpha1Reference) SetValue(v Constant)`

SetValue sets Value field to given value.

### HasValue

`func (o *V1alpha1Reference) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


