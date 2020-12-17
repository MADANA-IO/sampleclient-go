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

// XmlNs0MDNUserProfileImage 
type XmlNs0MDNUserProfileImage struct {
	XmlNs0MDNUserProfileImageAllOf
}

// NewXmlNs0MDNUserProfileImage instantiates a new XmlNs0MDNUserProfileImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0MDNUserProfileImage() *XmlNs0MDNUserProfileImage {
	this := XmlNs0MDNUserProfileImage{}
	return &this
}

// NewXmlNs0MDNUserProfileImageWithDefaults instantiates a new XmlNs0MDNUserProfileImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0MDNUserProfileImageWithDefaults() *XmlNs0MDNUserProfileImage {
	this := XmlNs0MDNUserProfileImage{}
	return &this
}

func (o XmlNs0MDNUserProfileImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedXmlNs0MDNUserProfileImageAllOf, errXmlNs0MDNUserProfileImageAllOf := json.Marshal(o.XmlNs0MDNUserProfileImageAllOf)
	if errXmlNs0MDNUserProfileImageAllOf != nil {
		return []byte{}, errXmlNs0MDNUserProfileImageAllOf
	}
	errXmlNs0MDNUserProfileImageAllOf = json.Unmarshal([]byte(serializedXmlNs0MDNUserProfileImageAllOf), &toSerialize)
	if errXmlNs0MDNUserProfileImageAllOf != nil {
		return []byte{}, errXmlNs0MDNUserProfileImageAllOf
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0MDNUserProfileImage struct {
	value *XmlNs0MDNUserProfileImage
	isSet bool
}

func (v NullableXmlNs0MDNUserProfileImage) Get() *XmlNs0MDNUserProfileImage {
	return v.value
}

func (v *NullableXmlNs0MDNUserProfileImage) Set(val *XmlNs0MDNUserProfileImage) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0MDNUserProfileImage) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0MDNUserProfileImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0MDNUserProfileImage(val *XmlNs0MDNUserProfileImage) *NullableXmlNs0MDNUserProfileImage {
	return &NullableXmlNs0MDNUserProfileImage{value: val, isSet: true}
}

func (v NullableXmlNs0MDNUserProfileImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0MDNUserProfileImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


