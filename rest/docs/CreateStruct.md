# CreateStruct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageName** | Pointer to **string** | The type name of the message to be created, empty when creating map literals. | [optional] 
**Entries** | Pointer to [**[]Entry**](Entry.md) | The entries in the creation expression. | [optional] 

## Methods

### NewCreateStruct

`func NewCreateStruct() *CreateStruct`

NewCreateStruct instantiates a new CreateStruct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStructWithDefaults

`func NewCreateStructWithDefaults() *CreateStruct`

NewCreateStructWithDefaults instantiates a new CreateStruct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageName

`func (o *CreateStruct) GetMessageName() string`

GetMessageName returns the MessageName field if non-nil, zero value otherwise.

### GetMessageNameOk

`func (o *CreateStruct) GetMessageNameOk() (*string, bool)`

GetMessageNameOk returns a tuple with the MessageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageName

`func (o *CreateStruct) SetMessageName(v string)`

SetMessageName sets MessageName field to given value.

### HasMessageName

`func (o *CreateStruct) HasMessageName() bool`

HasMessageName returns a boolean if a field has been set.

### GetEntries

`func (o *CreateStruct) GetEntries() []Entry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *CreateStruct) GetEntriesOk() (*[]Entry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *CreateStruct) SetEntries(v []Entry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *CreateStruct) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


