# SchemaListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Head** | Pointer to **string** |  | [optional] 
**Schemas** | Pointer to [**[]SchemaList**](SchemaList.md) |  | [optional] 
**ContinuousToken** | Pointer to **string** | continuous_token is a string that can be used to paginate and retrieve the next set of results. | [optional] 

## Methods

### NewSchemaListResponse

`func NewSchemaListResponse() *SchemaListResponse`

NewSchemaListResponse instantiates a new SchemaListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaListResponseWithDefaults

`func NewSchemaListResponseWithDefaults() *SchemaListResponse`

NewSchemaListResponseWithDefaults instantiates a new SchemaListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHead

`func (o *SchemaListResponse) GetHead() string`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *SchemaListResponse) GetHeadOk() (*string, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *SchemaListResponse) SetHead(v string)`

SetHead sets Head field to given value.

### HasHead

`func (o *SchemaListResponse) HasHead() bool`

HasHead returns a boolean if a field has been set.

### GetSchemas

`func (o *SchemaListResponse) GetSchemas() []SchemaList`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SchemaListResponse) GetSchemasOk() (*[]SchemaList, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SchemaListResponse) SetSchemas(v []SchemaList)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *SchemaListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetContinuousToken

`func (o *SchemaListResponse) GetContinuousToken() string`

GetContinuousToken returns the ContinuousToken field if non-nil, zero value otherwise.

### GetContinuousTokenOk

`func (o *SchemaListResponse) GetContinuousTokenOk() (*string, bool)`

GetContinuousTokenOk returns a tuple with the ContinuousToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousToken

`func (o *SchemaListResponse) SetContinuousToken(v string)`

SetContinuousToken sets ContinuousToken field to given value.

### HasContinuousToken

`func (o *SchemaListResponse) HasContinuousToken() bool`

HasContinuousToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


