/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.11
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0Process 
type XmlNs0Process struct {
}

// NewXmlNs0Process instantiates a new XmlNs0Process object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0Process() *XmlNs0Process {
	this := XmlNs0Process{}
	return &this
}

// NewXmlNs0ProcessWithDefaults instantiates a new XmlNs0Process object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0ProcessWithDefaults() *XmlNs0Process {
	this := XmlNs0Process{}
	return &this
}

func (o XmlNs0Process) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0Process struct {
	value *XmlNs0Process
	isSet bool
}

func (v NullableXmlNs0Process) Get() *XmlNs0Process {
	return v.value
}

func (v *NullableXmlNs0Process) Set(val *XmlNs0Process) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0Process) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0Process) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0Process(val *XmlNs0Process) *NullableXmlNs0Process {
	return &NullableXmlNs0Process{value: val, isSet: true}
}

func (v NullableXmlNs0Process) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0Process) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


