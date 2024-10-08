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

// checks if the ReadAttributesBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadAttributesBody{}

// ReadAttributesBody AttributeReadRequest defines the structure of a request for reading attributes. It includes the tenant_id, metadata, attribute filter, page size for pagination, and a continuous token for multi-page results.
type ReadAttributesBody struct {
	Metadata *AttributeReadRequestMetadata `json:"metadata,omitempty"`
	Filter *AttributeFilter `json:"filter,omitempty"`
	// page_size specifies the number of results to return in a single page. If more results are available, a continuous_token is included in the response.
	PageSize *int64 `json:"page_size,omitempty"`
	// continuous_token is used in case of paginated reads to get the next page of results.
	ContinuousToken *string `json:"continuous_token,omitempty"`
}

// NewReadAttributesBody instantiates a new ReadAttributesBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadAttributesBody() *ReadAttributesBody {
	this := ReadAttributesBody{}
	return &this
}

// NewReadAttributesBodyWithDefaults instantiates a new ReadAttributesBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadAttributesBodyWithDefaults() *ReadAttributesBody {
	this := ReadAttributesBody{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ReadAttributesBody) GetMetadata() AttributeReadRequestMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret AttributeReadRequestMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAttributesBody) GetMetadataOk() (*AttributeReadRequestMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ReadAttributesBody) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given AttributeReadRequestMetadata and assigns it to the Metadata field.
func (o *ReadAttributesBody) SetMetadata(v AttributeReadRequestMetadata) {
	o.Metadata = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ReadAttributesBody) GetFilter() AttributeFilter {
	if o == nil || IsNil(o.Filter) {
		var ret AttributeFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAttributesBody) GetFilterOk() (*AttributeFilter, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ReadAttributesBody) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given AttributeFilter and assigns it to the Filter field.
func (o *ReadAttributesBody) SetFilter(v AttributeFilter) {
	o.Filter = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ReadAttributesBody) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAttributesBody) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ReadAttributesBody) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *ReadAttributesBody) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetContinuousToken returns the ContinuousToken field value if set, zero value otherwise.
func (o *ReadAttributesBody) GetContinuousToken() string {
	if o == nil || IsNil(o.ContinuousToken) {
		var ret string
		return ret
	}
	return *o.ContinuousToken
}

// GetContinuousTokenOk returns a tuple with the ContinuousToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAttributesBody) GetContinuousTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ContinuousToken) {
		return nil, false
	}
	return o.ContinuousToken, true
}

// HasContinuousToken returns a boolean if a field has been set.
func (o *ReadAttributesBody) HasContinuousToken() bool {
	if o != nil && !IsNil(o.ContinuousToken) {
		return true
	}

	return false
}

// SetContinuousToken gets a reference to the given string and assigns it to the ContinuousToken field.
func (o *ReadAttributesBody) SetContinuousToken(v string) {
	o.ContinuousToken = &v
}

func (o ReadAttributesBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadAttributesBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.ContinuousToken) {
		toSerialize["continuous_token"] = o.ContinuousToken
	}
	return toSerialize, nil
}

type NullableReadAttributesBody struct {
	value *ReadAttributesBody
	isSet bool
}

func (v NullableReadAttributesBody) Get() *ReadAttributesBody {
	return v.value
}

func (v *NullableReadAttributesBody) Set(val *ReadAttributesBody) {
	v.value = val
	v.isSet = true
}

func (v NullableReadAttributesBody) IsSet() bool {
	return v.isSet
}

func (v *NullableReadAttributesBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadAttributesBody(val *ReadAttributesBody) *NullableReadAttributesBody {
	return &NullableReadAttributesBody{value: val, isSet: true}
}

func (v NullableReadAttributesBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadAttributesBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


