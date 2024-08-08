# BundleReadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bundle** | Pointer to [**DataBundle**](DataBundle.md) |  | [optional] 

## Methods

### NewBundleReadResponse

`func NewBundleReadResponse() *BundleReadResponse`

NewBundleReadResponse instantiates a new BundleReadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleReadResponseWithDefaults

`func NewBundleReadResponseWithDefaults() *BundleReadResponse`

NewBundleReadResponseWithDefaults instantiates a new BundleReadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundle

`func (o *BundleReadResponse) GetBundle() DataBundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *BundleReadResponse) GetBundleOk() (*DataBundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *BundleReadResponse) SetBundle(v DataBundle)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *BundleReadResponse) HasBundle() bool`

HasBundle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


