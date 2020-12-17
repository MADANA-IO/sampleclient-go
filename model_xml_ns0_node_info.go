/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.14
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0NodeInfo 
type XmlNs0NodeInfo struct {
	XmlNs0NodeInfoAllOf
}

// NewXmlNs0NodeInfo instantiates a new XmlNs0NodeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0NodeInfo() *XmlNs0NodeInfo {
	this := XmlNs0NodeInfo{}
	return &this
}

// NewXmlNs0NodeInfoWithDefaults instantiates a new XmlNs0NodeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0NodeInfoWithDefaults() *XmlNs0NodeInfo {
	this := XmlNs0NodeInfo{}
	return &this
}

func (o XmlNs0NodeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedXmlNs0NodeInfoAllOf, errXmlNs0NodeInfoAllOf := json.Marshal(o.XmlNs0NodeInfoAllOf)
	if errXmlNs0NodeInfoAllOf != nil {
		return []byte{}, errXmlNs0NodeInfoAllOf
	}
	errXmlNs0NodeInfoAllOf = json.Unmarshal([]byte(serializedXmlNs0NodeInfoAllOf), &toSerialize)
	if errXmlNs0NodeInfoAllOf != nil {
		return []byte{}, errXmlNs0NodeInfoAllOf
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0NodeInfo struct {
	value *XmlNs0NodeInfo
	isSet bool
}

func (v NullableXmlNs0NodeInfo) Get() *XmlNs0NodeInfo {
	return v.value
}

func (v *NullableXmlNs0NodeInfo) Set(val *XmlNs0NodeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0NodeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0NodeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0NodeInfo(val *XmlNs0NodeInfo) *NullableXmlNs0NodeInfo {
	return &NullableXmlNs0NodeInfo{value: val, isSet: true}
}

func (v NullableXmlNs0NodeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0NodeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


