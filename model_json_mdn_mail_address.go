/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.9
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonMDNMailAddress 
type JsonMDNMailAddress struct {
	// 
	Mail *string `json:"mail,omitempty"`
}

// NewJsonMDNMailAddress instantiates a new JsonMDNMailAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonMDNMailAddress() *JsonMDNMailAddress {
	this := JsonMDNMailAddress{}
	return &this
}

// NewJsonMDNMailAddressWithDefaults instantiates a new JsonMDNMailAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonMDNMailAddressWithDefaults() *JsonMDNMailAddress {
	this := JsonMDNMailAddress{}
	return &this
}

// GetMail returns the Mail field value if set, zero value otherwise.
func (o *JsonMDNMailAddress) GetMail() string {
	if o == nil || o.Mail == nil {
		var ret string
		return ret
	}
	return *o.Mail
}

// GetMailOk returns a tuple with the Mail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNMailAddress) GetMailOk() (*string, bool) {
	if o == nil || o.Mail == nil {
		return nil, false
	}
	return o.Mail, true
}

// HasMail returns a boolean if a field has been set.
func (o *JsonMDNMailAddress) HasMail() bool {
	if o != nil && o.Mail != nil {
		return true
	}

	return false
}

// SetMail gets a reference to the given string and assigns it to the Mail field.
func (o *JsonMDNMailAddress) SetMail(v string) {
	o.Mail = &v
}

func (o JsonMDNMailAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mail != nil {
		toSerialize["mail"] = o.Mail
	}
	return json.Marshal(toSerialize)
}

type NullableJsonMDNMailAddress struct {
	value *JsonMDNMailAddress
	isSet bool
}

func (v NullableJsonMDNMailAddress) Get() *JsonMDNMailAddress {
	return v.value
}

func (v *NullableJsonMDNMailAddress) Set(val *JsonMDNMailAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonMDNMailAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonMDNMailAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonMDNMailAddress(val *JsonMDNMailAddress) *NullableJsonMDNMailAddress {
	return &NullableJsonMDNMailAddress{value: val, isSet: true}
}

func (v NullableJsonMDNMailAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonMDNMailAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


