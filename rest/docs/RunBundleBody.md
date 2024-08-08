# RunBundleBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the bundle to be executed. | [optional] 
**Arguments** | Pointer to **map[string]string** | Additional key-value pairs for execution arguments. | [optional] 

## Methods

### NewRunBundleBody

`func NewRunBundleBody() *RunBundleBody`

NewRunBundleBody instantiates a new RunBundleBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunBundleBodyWithDefaults

`func NewRunBundleBodyWithDefaults() *RunBundleBody`

NewRunBundleBodyWithDefaults instantiates a new RunBundleBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RunBundleBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RunBundleBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RunBundleBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RunBundleBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArguments

`func (o *RunBundleBody) GetArguments() map[string]string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *RunBundleBody) GetArgumentsOk() (*map[string]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *RunBundleBody) SetArguments(v map[string]string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *RunBundleBody) HasArguments() bool`

HasArguments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


