/*
Permify API

Permify is an open source authorization service for creating fine-grained and scalable authorization systems.

API version: v0.10.1
Contact: hello@permify.co
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package permify

import (
	"encoding/json"
)

// checks if the BundleRunResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleRunResponse{}

// BundleRunResponse BundleRunResponse is the response for a BundleRunRequest. It includes a snap_token, which may be used for tracking the execution or its results.
type BundleRunResponse struct {
	// The snap token to avoid stale cache, see more details on [Snap Tokens](../../operations/snap-tokens)
	SnapToken *string `json:"snap_token,omitempty"`
}

// NewBundleRunResponse instantiates a new BundleRunResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleRunResponse() *BundleRunResponse {
	this := BundleRunResponse{}
	return &this
}

// NewBundleRunResponseWithDefaults instantiates a new BundleRunResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleRunResponseWithDefaults() *BundleRunResponse {
	this := BundleRunResponse{}
	return &this
}

// GetSnapToken returns the SnapToken field value if set, zero value otherwise.
func (o *BundleRunResponse) GetSnapToken() string {
	if o == nil || IsNil(o.SnapToken) {
		var ret string
		return ret
	}
	return *o.SnapToken
}

// GetSnapTokenOk returns a tuple with the SnapToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleRunResponse) GetSnapTokenOk() (*string, bool) {
	if o == nil || IsNil(o.SnapToken) {
		return nil, false
	}
	return o.SnapToken, true
}

// HasSnapToken returns a boolean if a field has been set.
func (o *BundleRunResponse) HasSnapToken() bool {
	if o != nil && !IsNil(o.SnapToken) {
		return true
	}

	return false
}

// SetSnapToken gets a reference to the given string and assigns it to the SnapToken field.
func (o *BundleRunResponse) SetSnapToken(v string) {
	o.SnapToken = &v
}

func (o BundleRunResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleRunResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SnapToken) {
		toSerialize["snap_token"] = o.SnapToken
	}
	return toSerialize, nil
}

type NullableBundleRunResponse struct {
	value *BundleRunResponse
	isSet bool
}

func (v NullableBundleRunResponse) Get() *BundleRunResponse {
	return v.value
}

func (v *NullableBundleRunResponse) Set(val *BundleRunResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleRunResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleRunResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleRunResponse(val *BundleRunResponse) *NullableBundleRunResponse {
	return &NullableBundleRunResponse{value: val, isSet: true}
}

func (v NullableBundleRunResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleRunResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


