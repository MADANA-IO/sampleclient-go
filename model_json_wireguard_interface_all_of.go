/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.56
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonWireguardInterfaceAllOf struct for JsonWireguardInterfaceAllOf
type JsonWireguardInterfaceAllOf struct {
	// 
	Port *string `json:"port,omitempty"`
}

// NewJsonWireguardInterfaceAllOf instantiates a new JsonWireguardInterfaceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonWireguardInterfaceAllOf() *JsonWireguardInterfaceAllOf {
	this := JsonWireguardInterfaceAllOf{}
	return &this
}

// NewJsonWireguardInterfaceAllOfWithDefaults instantiates a new JsonWireguardInterfaceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonWireguardInterfaceAllOfWithDefaults() *JsonWireguardInterfaceAllOf {
	this := JsonWireguardInterfaceAllOf{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *JsonWireguardInterfaceAllOf) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWireguardInterfaceAllOf) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *JsonWireguardInterfaceAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *JsonWireguardInterfaceAllOf) SetPort(v string) {
	o.Port = &v
}

func (o JsonWireguardInterfaceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableJsonWireguardInterfaceAllOf struct {
	value *JsonWireguardInterfaceAllOf
	isSet bool
}

func (v NullableJsonWireguardInterfaceAllOf) Get() *JsonWireguardInterfaceAllOf {
	return v.value
}

func (v *NullableJsonWireguardInterfaceAllOf) Set(val *JsonWireguardInterfaceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonWireguardInterfaceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonWireguardInterfaceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonWireguardInterfaceAllOf(val *JsonWireguardInterfaceAllOf) *NullableJsonWireguardInterfaceAllOf {
	return &NullableJsonWireguardInterfaceAllOf{value: val, isSet: true}
}

func (v NullableJsonWireguardInterfaceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonWireguardInterfaceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


