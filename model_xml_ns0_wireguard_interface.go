/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.22
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0WireguardInterface 
type XmlNs0WireguardInterface struct {
	// 
	Address *string `json:"address,omitempty"`
	// 
	Name *string `json:"name,omitempty"`
	// 
	Port *string `json:"port,omitempty"`
}

// NewXmlNs0WireguardInterface instantiates a new XmlNs0WireguardInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0WireguardInterface() *XmlNs0WireguardInterface {
	this := XmlNs0WireguardInterface{}
	return &this
}

// NewXmlNs0WireguardInterfaceWithDefaults instantiates a new XmlNs0WireguardInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0WireguardInterfaceWithDefaults() *XmlNs0WireguardInterface {
	this := XmlNs0WireguardInterface{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *XmlNs0WireguardInterface) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0WireguardInterface) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *XmlNs0WireguardInterface) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *XmlNs0WireguardInterface) SetAddress(v string) {
	o.Address = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *XmlNs0WireguardInterface) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0WireguardInterface) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *XmlNs0WireguardInterface) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *XmlNs0WireguardInterface) SetName(v string) {
	o.Name = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *XmlNs0WireguardInterface) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0WireguardInterface) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *XmlNs0WireguardInterface) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *XmlNs0WireguardInterface) SetPort(v string) {
	o.Port = &v
}

func (o XmlNs0WireguardInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0WireguardInterface struct {
	value *XmlNs0WireguardInterface
	isSet bool
}

func (v NullableXmlNs0WireguardInterface) Get() *XmlNs0WireguardInterface {
	return v.value
}

func (v *NullableXmlNs0WireguardInterface) Set(val *XmlNs0WireguardInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0WireguardInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0WireguardInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0WireguardInterface(val *XmlNs0WireguardInterface) *NullableXmlNs0WireguardInterface {
	return &NullableXmlNs0WireguardInterface{value: val, isSet: true}
}

func (v NullableXmlNs0WireguardInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0WireguardInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


