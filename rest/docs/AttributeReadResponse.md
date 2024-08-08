# AttributeReadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]Attribute**](Attribute.md) | attributes is a list of the attributes retrieved in the read operation. | [optional] 
**ContinuousToken** | Pointer to **string** | continuous_token is used in the case of paginated reads to retrieve the next page of results. | [optional] 

## Methods

### NewAttributeReadResponse

`func NewAttributeReadResponse() *AttributeReadResponse`

NewAttributeReadResponse instantiates a new AttributeReadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeReadResponseWithDefaults

`func NewAttributeReadResponseWithDefaults() *AttributeReadResponse`

NewAttributeReadResponseWithDefaults instantiates a new AttributeReadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *AttributeReadResponse) GetAttributes() []Attribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AttributeReadResponse) GetAttributesOk() (*[]Attribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AttributeReadResponse) SetAttributes(v []Attribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AttributeReadResponse) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetContinuousToken

`func (o *AttributeReadResponse) GetContinuousToken() string`

GetContinuousToken returns the ContinuousToken field if non-nil, zero value otherwise.

### GetContinuousTokenOk

`func (o *AttributeReadResponse) GetContinuousTokenOk() (*string, bool)`

GetContinuousTokenOk returns a tuple with the ContinuousToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousToken

`func (o *AttributeReadResponse) SetContinuousToken(v string)`

SetContinuousToken sets ContinuousToken field to given value.

### HasContinuousToken

`func (o *AttributeReadResponse) HasContinuousToken() bool`

HasContinuousToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


