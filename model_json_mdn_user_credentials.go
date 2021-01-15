/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.32
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonMDNUserCredentials 
type JsonMDNUserCredentials struct {
	// 
	Username *string `json:"username,omitempty"`
	// 
	Password *string `json:"password,omitempty"`
}

// NewJsonMDNUserCredentials instantiates a new JsonMDNUserCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonMDNUserCredentials() *JsonMDNUserCredentials {
	this := JsonMDNUserCredentials{}
	return &this
}

// NewJsonMDNUserCredentialsWithDefaults instantiates a new JsonMDNUserCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonMDNUserCredentialsWithDefaults() *JsonMDNUserCredentials {
	this := JsonMDNUserCredentials{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *JsonMDNUserCredentials) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUserCredentials) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *JsonMDNUserCredentials) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *JsonMDNUserCredentials) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *JsonMDNUserCredentials) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonMDNUserCredentials) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *JsonMDNUserCredentials) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *JsonMDNUserCredentials) SetPassword(v string) {
	o.Password = &v
}

func (o JsonMDNUserCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableJsonMDNUserCredentials struct {
	value *JsonMDNUserCredentials
	isSet bool
}

func (v NullableJsonMDNUserCredentials) Get() *JsonMDNUserCredentials {
	return v.value
}

func (v *NullableJsonMDNUserCredentials) Set(val *JsonMDNUserCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonMDNUserCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonMDNUserCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonMDNUserCredentials(val *JsonMDNUserCredentials) *NullableJsonMDNUserCredentials {
	return &NullableJsonMDNUserCredentials{value: val, isSet: true}
}

func (v NullableJsonMDNUserCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonMDNUserCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


