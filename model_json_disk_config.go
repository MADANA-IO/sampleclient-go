/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.41
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonDiskConfig 
type JsonDiskConfig struct {
	// 
	RoothashOffset *int32 `json:"roothash_offset,omitempty"`
	// 
	Readonly *bool `json:"readonly,omitempty"`
	// 
	Disk *string `json:"disk,omitempty"`
	// 
	Roothash *string `json:"roothash,omitempty"`
}

// NewJsonDiskConfig instantiates a new JsonDiskConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonDiskConfig() *JsonDiskConfig {
	this := JsonDiskConfig{}
	return &this
}

// NewJsonDiskConfigWithDefaults instantiates a new JsonDiskConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonDiskConfigWithDefaults() *JsonDiskConfig {
	this := JsonDiskConfig{}
	return &this
}

// GetRoothashOffset returns the RoothashOffset field value if set, zero value otherwise.
func (o *JsonDiskConfig) GetRoothashOffset() int32 {
	if o == nil || o.RoothashOffset == nil {
		var ret int32
		return ret
	}
	return *o.RoothashOffset
}

// GetRoothashOffsetOk returns a tuple with the RoothashOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonDiskConfig) GetRoothashOffsetOk() (*int32, bool) {
	if o == nil || o.RoothashOffset == nil {
		return nil, false
	}
	return o.RoothashOffset, true
}

// HasRoothashOffset returns a boolean if a field has been set.
func (o *JsonDiskConfig) HasRoothashOffset() bool {
	if o != nil && o.RoothashOffset != nil {
		return true
	}

	return false
}

// SetRoothashOffset gets a reference to the given int32 and assigns it to the RoothashOffset field.
func (o *JsonDiskConfig) SetRoothashOffset(v int32) {
	o.RoothashOffset = &v
}

// GetReadonly returns the Readonly field value if set, zero value otherwise.
func (o *JsonDiskConfig) GetReadonly() bool {
	if o == nil || o.Readonly == nil {
		var ret bool
		return ret
	}
	return *o.Readonly
}

// GetReadonlyOk returns a tuple with the Readonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonDiskConfig) GetReadonlyOk() (*bool, bool) {
	if o == nil || o.Readonly == nil {
		return nil, false
	}
	return o.Readonly, true
}

// HasReadonly returns a boolean if a field has been set.
func (o *JsonDiskConfig) HasReadonly() bool {
	if o != nil && o.Readonly != nil {
		return true
	}

	return false
}

// SetReadonly gets a reference to the given bool and assigns it to the Readonly field.
func (o *JsonDiskConfig) SetReadonly(v bool) {
	o.Readonly = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *JsonDiskConfig) GetDisk() string {
	if o == nil || o.Disk == nil {
		var ret string
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonDiskConfig) GetDiskOk() (*string, bool) {
	if o == nil || o.Disk == nil {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *JsonDiskConfig) HasDisk() bool {
	if o != nil && o.Disk != nil {
		return true
	}

	return false
}

// SetDisk gets a reference to the given string and assigns it to the Disk field.
func (o *JsonDiskConfig) SetDisk(v string) {
	o.Disk = &v
}

// GetRoothash returns the Roothash field value if set, zero value otherwise.
func (o *JsonDiskConfig) GetRoothash() string {
	if o == nil || o.Roothash == nil {
		var ret string
		return ret
	}
	return *o.Roothash
}

// GetRoothashOk returns a tuple with the Roothash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonDiskConfig) GetRoothashOk() (*string, bool) {
	if o == nil || o.Roothash == nil {
		return nil, false
	}
	return o.Roothash, true
}

// HasRoothash returns a boolean if a field has been set.
func (o *JsonDiskConfig) HasRoothash() bool {
	if o != nil && o.Roothash != nil {
		return true
	}

	return false
}

// SetRoothash gets a reference to the given string and assigns it to the Roothash field.
func (o *JsonDiskConfig) SetRoothash(v string) {
	o.Roothash = &v
}

func (o JsonDiskConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RoothashOffset != nil {
		toSerialize["roothash_offset"] = o.RoothashOffset
	}
	if o.Readonly != nil {
		toSerialize["readonly"] = o.Readonly
	}
	if o.Disk != nil {
		toSerialize["disk"] = o.Disk
	}
	if o.Roothash != nil {
		toSerialize["roothash"] = o.Roothash
	}
	return json.Marshal(toSerialize)
}

type NullableJsonDiskConfig struct {
	value *JsonDiskConfig
	isSet bool
}

func (v NullableJsonDiskConfig) Get() *JsonDiskConfig {
	return v.value
}

func (v *NullableJsonDiskConfig) Set(val *JsonDiskConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonDiskConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonDiskConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonDiskConfig(val *JsonDiskConfig) *NullableJsonDiskConfig {
	return &NullableJsonDiskConfig{value: val, isSet: true}
}

func (v NullableJsonDiskConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonDiskConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


