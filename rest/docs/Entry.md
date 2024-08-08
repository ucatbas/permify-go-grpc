# Entry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Required. An id assigned to this node by the parser which is unique in a given expression tree. This is used to associate type information and other attributes to the node. | [optional] 
**FieldKey** | Pointer to **string** | The field key for a message creator statement. | [optional] 
**MapKey** | Pointer to [**Expr**](Expr.md) |  | [optional] 
**Value** | Pointer to [**Expr**](Expr.md) |  | [optional] 
**OptionalEntry** | Pointer to **bool** | Whether the key-value pair is optional. | [optional] 

## Methods

### NewEntry

`func NewEntry() *Entry`

NewEntry instantiates a new Entry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryWithDefaults

`func NewEntryWithDefaults() *Entry`

NewEntryWithDefaults instantiates a new Entry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Entry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Entry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFieldKey

`func (o *Entry) GetFieldKey() string`

GetFieldKey returns the FieldKey field if non-nil, zero value otherwise.

### GetFieldKeyOk

`func (o *Entry) GetFieldKeyOk() (*string, bool)`

GetFieldKeyOk returns a tuple with the FieldKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldKey

`func (o *Entry) SetFieldKey(v string)`

SetFieldKey sets FieldKey field to given value.

### HasFieldKey

`func (o *Entry) HasFieldKey() bool`

HasFieldKey returns a boolean if a field has been set.

### GetMapKey

`func (o *Entry) GetMapKey() Expr`

GetMapKey returns the MapKey field if non-nil, zero value otherwise.

### GetMapKeyOk

`func (o *Entry) GetMapKeyOk() (*Expr, bool)`

GetMapKeyOk returns a tuple with the MapKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapKey

`func (o *Entry) SetMapKey(v Expr)`

SetMapKey sets MapKey field to given value.

### HasMapKey

`func (o *Entry) HasMapKey() bool`

HasMapKey returns a boolean if a field has been set.

### GetValue

`func (o *Entry) GetValue() Expr`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Entry) GetValueOk() (*Expr, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Entry) SetValue(v Expr)`

SetValue sets Value field to given value.

### HasValue

`func (o *Entry) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOptionalEntry

`func (o *Entry) GetOptionalEntry() bool`

GetOptionalEntry returns the OptionalEntry field if non-nil, zero value otherwise.

### GetOptionalEntryOk

`func (o *Entry) GetOptionalEntryOk() (*bool, bool)`

GetOptionalEntryOk returns a tuple with the OptionalEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalEntry

`func (o *Entry) SetOptionalEntry(v bool)`

SetOptionalEntry sets OptionalEntry field to given value.

### HasOptionalEntry

`func (o *Entry) HasOptionalEntry() bool`

HasOptionalEntry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


