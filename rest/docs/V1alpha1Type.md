# V1alpha1Type

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dyn** | Pointer to **map[string]interface{}** | Dynamic type. | [optional] 
**Null** | Pointer to **string** | Null value. | [optional] 
**Primitive** | Pointer to [**PrimitiveType**](PrimitiveType.md) |  | [optional] 
**Wrapper** | Pointer to [**PrimitiveType**](PrimitiveType.md) |  | [optional] 
**WellKnown** | Pointer to [**WellKnownType**](WellKnownType.md) |  | [optional] 
**ListType** | Pointer to [**ListType**](ListType.md) |  | [optional] 
**MapType** | Pointer to [**MapType**](MapType.md) |  | [optional] 
**Function** | Pointer to [**FunctionType**](FunctionType.md) |  | [optional] 
**MessageType** | Pointer to **string** | Protocol buffer message type.  The &#x60;message_type&#x60; string specifies the qualified message type name. For example, &#x60;google.plus.Profile&#x60;. | [optional] 
**TypeParam** | Pointer to **string** | Type param type.  The &#x60;type_param&#x60; string specifies the type parameter name, e.g. &#x60;list&lt;E&gt;&#x60; would be a &#x60;list_type&#x60; whose element type was a &#x60;type_param&#x60; type named &#x60;E&#x60;. | [optional] 
**Type** | Pointer to [**V1alpha1Type**](V1alpha1Type.md) |  | [optional] 
**Error** | Pointer to **map[string]interface{}** | Error type.  During type-checking if an expression is an error, its type is propagated as the &#x60;ERROR&#x60; type. This permits the type-checker to discover other errors present in the expression. | [optional] 
**AbstractType** | Pointer to [**AbstractType**](AbstractType.md) |  | [optional] 

## Methods

### NewV1alpha1Type

`func NewV1alpha1Type() *V1alpha1Type`

NewV1alpha1Type instantiates a new V1alpha1Type object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1TypeWithDefaults

`func NewV1alpha1TypeWithDefaults() *V1alpha1Type`

NewV1alpha1TypeWithDefaults instantiates a new V1alpha1Type object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDyn

`func (o *V1alpha1Type) GetDyn() map[string]interface{}`

GetDyn returns the Dyn field if non-nil, zero value otherwise.

### GetDynOk

`func (o *V1alpha1Type) GetDynOk() (*map[string]interface{}, bool)`

GetDynOk returns a tuple with the Dyn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDyn

`func (o *V1alpha1Type) SetDyn(v map[string]interface{})`

SetDyn sets Dyn field to given value.

### HasDyn

`func (o *V1alpha1Type) HasDyn() bool`

HasDyn returns a boolean if a field has been set.

### GetNull

`func (o *V1alpha1Type) GetNull() string`

GetNull returns the Null field if non-nil, zero value otherwise.

### GetNullOk

`func (o *V1alpha1Type) GetNullOk() (*string, bool)`

GetNullOk returns a tuple with the Null field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNull

`func (o *V1alpha1Type) SetNull(v string)`

SetNull sets Null field to given value.

### HasNull

`func (o *V1alpha1Type) HasNull() bool`

HasNull returns a boolean if a field has been set.

### GetPrimitive

`func (o *V1alpha1Type) GetPrimitive() PrimitiveType`

GetPrimitive returns the Primitive field if non-nil, zero value otherwise.

### GetPrimitiveOk

`func (o *V1alpha1Type) GetPrimitiveOk() (*PrimitiveType, bool)`

GetPrimitiveOk returns a tuple with the Primitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimitive

`func (o *V1alpha1Type) SetPrimitive(v PrimitiveType)`

SetPrimitive sets Primitive field to given value.

### HasPrimitive

`func (o *V1alpha1Type) HasPrimitive() bool`

HasPrimitive returns a boolean if a field has been set.

### GetWrapper

`func (o *V1alpha1Type) GetWrapper() PrimitiveType`

GetWrapper returns the Wrapper field if non-nil, zero value otherwise.

### GetWrapperOk

`func (o *V1alpha1Type) GetWrapperOk() (*PrimitiveType, bool)`

GetWrapperOk returns a tuple with the Wrapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapper

`func (o *V1alpha1Type) SetWrapper(v PrimitiveType)`

SetWrapper sets Wrapper field to given value.

### HasWrapper

`func (o *V1alpha1Type) HasWrapper() bool`

HasWrapper returns a boolean if a field has been set.

### GetWellKnown

`func (o *V1alpha1Type) GetWellKnown() WellKnownType`

GetWellKnown returns the WellKnown field if non-nil, zero value otherwise.

### GetWellKnownOk

`func (o *V1alpha1Type) GetWellKnownOk() (*WellKnownType, bool)`

GetWellKnownOk returns a tuple with the WellKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWellKnown

`func (o *V1alpha1Type) SetWellKnown(v WellKnownType)`

SetWellKnown sets WellKnown field to given value.

### HasWellKnown

`func (o *V1alpha1Type) HasWellKnown() bool`

HasWellKnown returns a boolean if a field has been set.

### GetListType

`func (o *V1alpha1Type) GetListType() ListType`

GetListType returns the ListType field if non-nil, zero value otherwise.

### GetListTypeOk

`func (o *V1alpha1Type) GetListTypeOk() (*ListType, bool)`

GetListTypeOk returns a tuple with the ListType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListType

`func (o *V1alpha1Type) SetListType(v ListType)`

SetListType sets ListType field to given value.

### HasListType

`func (o *V1alpha1Type) HasListType() bool`

HasListType returns a boolean if a field has been set.

### GetMapType

`func (o *V1alpha1Type) GetMapType() MapType`

GetMapType returns the MapType field if non-nil, zero value otherwise.

### GetMapTypeOk

`func (o *V1alpha1Type) GetMapTypeOk() (*MapType, bool)`

GetMapTypeOk returns a tuple with the MapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapType

`func (o *V1alpha1Type) SetMapType(v MapType)`

SetMapType sets MapType field to given value.

### HasMapType

`func (o *V1alpha1Type) HasMapType() bool`

HasMapType returns a boolean if a field has been set.

### GetFunction

`func (o *V1alpha1Type) GetFunction() FunctionType`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *V1alpha1Type) GetFunctionOk() (*FunctionType, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *V1alpha1Type) SetFunction(v FunctionType)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *V1alpha1Type) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetMessageType

`func (o *V1alpha1Type) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *V1alpha1Type) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *V1alpha1Type) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *V1alpha1Type) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetTypeParam

`func (o *V1alpha1Type) GetTypeParam() string`

GetTypeParam returns the TypeParam field if non-nil, zero value otherwise.

### GetTypeParamOk

`func (o *V1alpha1Type) GetTypeParamOk() (*string, bool)`

GetTypeParamOk returns a tuple with the TypeParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeParam

`func (o *V1alpha1Type) SetTypeParam(v string)`

SetTypeParam sets TypeParam field to given value.

### HasTypeParam

`func (o *V1alpha1Type) HasTypeParam() bool`

HasTypeParam returns a boolean if a field has been set.

### GetType

`func (o *V1alpha1Type) GetType() V1alpha1Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1alpha1Type) GetTypeOk() (*V1alpha1Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1alpha1Type) SetType(v V1alpha1Type)`

SetType sets Type field to given value.

### HasType

`func (o *V1alpha1Type) HasType() bool`

HasType returns a boolean if a field has been set.

### GetError

`func (o *V1alpha1Type) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V1alpha1Type) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V1alpha1Type) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *V1alpha1Type) HasError() bool`

HasError returns a boolean if a field has been set.

### GetAbstractType

`func (o *V1alpha1Type) GetAbstractType() AbstractType`

GetAbstractType returns the AbstractType field if non-nil, zero value otherwise.

### GetAbstractTypeOk

`func (o *V1alpha1Type) GetAbstractTypeOk() (*AbstractType, bool)`

GetAbstractTypeOk returns a tuple with the AbstractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstractType

`func (o *V1alpha1Type) SetAbstractType(v AbstractType)`

SetAbstractType sets AbstractType field to given value.

### HasAbstractType

`func (o *V1alpha1Type) HasAbstractType() bool`

HasAbstractType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


