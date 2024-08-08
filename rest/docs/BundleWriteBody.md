# BundleWriteBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bundles** | Pointer to [**[]DataBundle**](DataBundle.md) | Contains the bundle data to be written. | [optional] 

## Methods

### NewBundleWriteBody

`func NewBundleWriteBody() *BundleWriteBody`

NewBundleWriteBody instantiates a new BundleWriteBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleWriteBodyWithDefaults

`func NewBundleWriteBodyWithDefaults() *BundleWriteBody`

NewBundleWriteBodyWithDefaults instantiates a new BundleWriteBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundles

`func (o *BundleWriteBody) GetBundles() []DataBundle`

GetBundles returns the Bundles field if non-nil, zero value otherwise.

### GetBundlesOk

`func (o *BundleWriteBody) GetBundlesOk() (*[]DataBundle, bool)`

GetBundlesOk returns a tuple with the Bundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundles

`func (o *BundleWriteBody) SetBundles(v []DataBundle)`

SetBundles sets Bundles field to given value.

### HasBundles

`func (o *BundleWriteBody) HasBundles() bool`

HasBundles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


