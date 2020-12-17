/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.6
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0SignedData 
type XmlNs0SignedData struct {
	XmlNs0SignedDataAllOf
}

// NewXmlNs0SignedData instantiates a new XmlNs0SignedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0SignedData() *XmlNs0SignedData {
	this := XmlNs0SignedData{}
	return &this
}

// NewXmlNs0SignedDataWithDefaults instantiates a new XmlNs0SignedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0SignedDataWithDefaults() *XmlNs0SignedData {
	this := XmlNs0SignedData{}
	return &this
}

func (o XmlNs0SignedData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedXmlNs0SignedDataAllOf, errXmlNs0SignedDataAllOf := json.Marshal(o.XmlNs0SignedDataAllOf)
	if errXmlNs0SignedDataAllOf != nil {
		return []byte{}, errXmlNs0SignedDataAllOf
	}
	errXmlNs0SignedDataAllOf = json.Unmarshal([]byte(serializedXmlNs0SignedDataAllOf), &toSerialize)
	if errXmlNs0SignedDataAllOf != nil {
		return []byte{}, errXmlNs0SignedDataAllOf
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0SignedData struct {
	value *XmlNs0SignedData
	isSet bool
}

func (v NullableXmlNs0SignedData) Get() *XmlNs0SignedData {
	return v.value
}

func (v *NullableXmlNs0SignedData) Set(val *XmlNs0SignedData) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0SignedData) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0SignedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0SignedData(val *XmlNs0SignedData) *NullableXmlNs0SignedData {
	return &NullableXmlNs0SignedData{value: val, isSet: true}
}

func (v NullableXmlNs0SignedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0SignedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


