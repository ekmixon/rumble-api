/*
 * Rumble API
 *
 * Rumble Network Discovery API
 *
 * API version: 1.0.4
 * Contact: support@rumble.run
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// AssetsWithCheckpoint struct for AssetsWithCheckpoint
type AssetsWithCheckpoint struct {
	Since  int64   `json:"since"`
	Assets []Asset `json:"assets"`
}

// NewAssetsWithCheckpoint instantiates a new AssetsWithCheckpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsWithCheckpoint(since int64, assets []Asset) *AssetsWithCheckpoint {
	this := AssetsWithCheckpoint{}
	this.Since = since
	this.Assets = assets
	return &this
}

// NewAssetsWithCheckpointWithDefaults instantiates a new AssetsWithCheckpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsWithCheckpointWithDefaults() *AssetsWithCheckpoint {
	this := AssetsWithCheckpoint{}
	return &this
}

// GetSince returns the Since field value
func (o *AssetsWithCheckpoint) GetSince() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Since
}

// GetSinceOk returns a tuple with the Since field value
// and a boolean to check if the value has been set.
func (o *AssetsWithCheckpoint) GetSinceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Since, true
}

// SetSince sets field value
func (o *AssetsWithCheckpoint) SetSince(v int64) {
	o.Since = v
}

// GetAssets returns the Assets field value
func (o *AssetsWithCheckpoint) GetAssets() []Asset {
	if o == nil {
		var ret []Asset
		return ret
	}

	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value
// and a boolean to check if the value has been set.
func (o *AssetsWithCheckpoint) GetAssetsOk() (*[]Asset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assets, true
}

// SetAssets sets field value
func (o *AssetsWithCheckpoint) SetAssets(v []Asset) {
	o.Assets = v
}

func (o AssetsWithCheckpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["since"] = o.Since
	}
	if true {
		toSerialize["assets"] = o.Assets
	}
	return json.Marshal(toSerialize)
}

type NullableAssetsWithCheckpoint struct {
	value *AssetsWithCheckpoint
	isSet bool
}

func (v NullableAssetsWithCheckpoint) Get() *AssetsWithCheckpoint {
	return v.value
}

func (v *NullableAssetsWithCheckpoint) Set(val *AssetsWithCheckpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsWithCheckpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsWithCheckpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsWithCheckpoint(val *AssetsWithCheckpoint) *NullableAssetsWithCheckpoint {
	return &NullableAssetsWithCheckpoint{value: val, isSet: true}
}

func (v NullableAssetsWithCheckpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsWithCheckpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
