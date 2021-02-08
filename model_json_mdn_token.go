/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.51
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonMDNToken 
type JsonMDNToken struct {
	// 
	Token *string `json:"token,omitempty"`
}

// NewJsonMDNToken instantiates a new JsonMDNToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonMDNToken() *JsonMDNToken {
	this := JsonMDNToken{}
	return &this
}

// NewJsonMDNTokenWithDefaults instantiates a new JsonMDNToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonMDNTokenWithDefaults() *JsonMDNToken {
	this := JsonMDNToken{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *JsonMDNToken) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNToken) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *JsonMDNToken) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *JsonMDNToken) SetToken(v string) {
	o.Token = &v
}

func (o JsonMDNToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableJsonMDNToken struct {
	value *JsonMDNToken
	isSet bool
}

func (v NullableJsonMDNToken) Get() *JsonMDNToken {
	return v.value
}

func (v *NullableJsonMDNToken) Set(val *JsonMDNToken) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonMDNToken) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonMDNToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonMDNToken(val *JsonMDNToken) *NullableJsonMDNToken {
	return &NullableJsonMDNToken{value: val, isSet: true}
}

func (v NullableJsonMDNToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonMDNToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


