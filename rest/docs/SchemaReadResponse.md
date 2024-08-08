# SchemaReadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to [**SchemaDefinition**](SchemaDefinition.md) |  | [optional] 

## Methods

### NewSchemaReadResponse

`func NewSchemaReadResponse() *SchemaReadResponse`

NewSchemaReadResponse instantiates a new SchemaReadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaReadResponseWithDefaults

`func NewSchemaReadResponseWithDefaults() *SchemaReadResponse`

NewSchemaReadResponseWithDefaults instantiates a new SchemaReadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *SchemaReadResponse) GetSchema() SchemaDefinition`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SchemaReadResponse) GetSchemaOk() (*SchemaDefinition, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SchemaReadResponse) SetSchema(v SchemaDefinition)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *SchemaReadResponse) HasSchema() bool`

HasSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


