/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.40
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0EnclavePort 
type XmlNs0EnclavePort struct {
	// 
	Name *string `json:"name,omitempty"`
	// 
	Port *string `json:"port,omitempty"`
	// 
	Protocol *string `json:"protocol,omitempty"`
}

// NewXmlNs0EnclavePort instantiates a new XmlNs0EnclavePort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0EnclavePort() *XmlNs0EnclavePort {
	this := XmlNs0EnclavePort{}
	return &this
}

// NewXmlNs0EnclavePortWithDefaults instantiates a new XmlNs0EnclavePort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0EnclavePortWithDefaults() *XmlNs0EnclavePort {
	this := XmlNs0EnclavePort{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *XmlNs0EnclavePort) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclavePort) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *XmlNs0EnclavePort) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *XmlNs0EnclavePort) SetName(v string) {
	o.Name = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *XmlNs0EnclavePort) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclavePort) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *XmlNs0EnclavePort) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *XmlNs0EnclavePort) SetPort(v string) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *XmlNs0EnclavePort) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclavePort) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *XmlNs0EnclavePort) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *XmlNs0EnclavePort) SetProtocol(v string) {
	o.Protocol = &v
}

func (o XmlNs0EnclavePort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0EnclavePort struct {
	value *XmlNs0EnclavePort
	isSet bool
}

func (v NullableXmlNs0EnclavePort) Get() *XmlNs0EnclavePort {
	return v.value
}

func (v *NullableXmlNs0EnclavePort) Set(val *XmlNs0EnclavePort) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0EnclavePort) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0EnclavePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0EnclavePort(val *XmlNs0EnclavePort) *NullableXmlNs0EnclavePort {
	return &NullableXmlNs0EnclavePort{value: val, isSet: true}
}

func (v NullableXmlNs0EnclavePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0EnclavePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


