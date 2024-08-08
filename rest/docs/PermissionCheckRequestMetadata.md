# PermissionCheckRequestMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** | Version of the schema. | [optional] 
**SnapToken** | Pointer to **string** | The snap token to avoid stale cache, see more details on [Snap Tokens](../../operations/snap-tokens) | [optional] 
**Depth** | Pointer to **int32** | Query limit when if recursive database queries got in loop | [optional] 

## Methods

### NewPermissionCheckRequestMetadata

`func NewPermissionCheckRequestMetadata() *PermissionCheckRequestMetadata`

NewPermissionCheckRequestMetadata instantiates a new PermissionCheckRequestMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionCheckRequestMetadataWithDefaults

`func NewPermissionCheckRequestMetadataWithDefaults() *PermissionCheckRequestMetadata`

NewPermissionCheckRequestMetadataWithDefaults instantiates a new PermissionCheckRequestMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *PermissionCheckRequestMetadata) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *PermissionCheckRequestMetadata) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *PermissionCheckRequestMetadata) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *PermissionCheckRequestMetadata) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSnapToken

`func (o *PermissionCheckRequestMetadata) GetSnapToken() string`

GetSnapToken returns the SnapToken field if non-nil, zero value otherwise.

### GetSnapTokenOk

`func (o *PermissionCheckRequestMetadata) GetSnapTokenOk() (*string, bool)`

GetSnapTokenOk returns a tuple with the SnapToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapToken

`func (o *PermissionCheckRequestMetadata) SetSnapToken(v string)`

SetSnapToken sets SnapToken field to given value.

### HasSnapToken

`func (o *PermissionCheckRequestMetadata) HasSnapToken() bool`

HasSnapToken returns a boolean if a field has been set.

### GetDepth

`func (o *PermissionCheckRequestMetadata) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *PermissionCheckRequestMetadata) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *PermissionCheckRequestMetadata) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *PermissionCheckRequestMetadata) HasDepth() bool`

HasDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


