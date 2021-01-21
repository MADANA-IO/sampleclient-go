/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.42
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0SGXInfo 
type XmlNs0SGXInfo struct {
	// 
	Status *string `json:"status,omitempty"`
	// 
	Version *string `json:"version,omitempty"`
}

// NewXmlNs0SGXInfo instantiates a new XmlNs0SGXInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0SGXInfo() *XmlNs0SGXInfo {
	this := XmlNs0SGXInfo{}
	return &this
}

// NewXmlNs0SGXInfoWithDefaults instantiates a new XmlNs0SGXInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0SGXInfoWithDefaults() *XmlNs0SGXInfo {
	this := XmlNs0SGXInfo{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *XmlNs0SGXInfo) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0SGXInfo) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *XmlNs0SGXInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *XmlNs0SGXInfo) SetStatus(v string) {
	o.Status = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *XmlNs0SGXInfo) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0SGXInfo) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *XmlNs0SGXInfo) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *XmlNs0SGXInfo) SetVersion(v string) {
	o.Version = &v
}

func (o XmlNs0SGXInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0SGXInfo struct {
	value *XmlNs0SGXInfo
	isSet bool
}

func (v NullableXmlNs0SGXInfo) Get() *XmlNs0SGXInfo {
	return v.value
}

func (v *NullableXmlNs0SGXInfo) Set(val *XmlNs0SGXInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0SGXInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0SGXInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0SGXInfo(val *XmlNs0SGXInfo) *NullableXmlNs0SGXInfo {
	return &NullableXmlNs0SGXInfo{value: val, isSet: true}
}

func (v NullableXmlNs0SGXInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0SGXInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


