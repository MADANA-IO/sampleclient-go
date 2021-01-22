/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.44
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0EnclaveRunningAttestationAllOf struct for XmlNs0EnclaveRunningAttestationAllOf
type XmlNs0EnclaveRunningAttestationAllOf struct {
	EnclaveProcess *XmlNs0EnclaveProcess `json:"enclaveProcess,omitempty"`
	NodeInfo *XmlNs0NodeInfo `json:"nodeInfo,omitempty"`
}

// NewXmlNs0EnclaveRunningAttestationAllOf instantiates a new XmlNs0EnclaveRunningAttestationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0EnclaveRunningAttestationAllOf() *XmlNs0EnclaveRunningAttestationAllOf {
	this := XmlNs0EnclaveRunningAttestationAllOf{}
	return &this
}

// NewXmlNs0EnclaveRunningAttestationAllOfWithDefaults instantiates a new XmlNs0EnclaveRunningAttestationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0EnclaveRunningAttestationAllOfWithDefaults() *XmlNs0EnclaveRunningAttestationAllOf {
	this := XmlNs0EnclaveRunningAttestationAllOf{}
	return &this
}

// GetEnclaveProcess returns the EnclaveProcess field value if set, zero value otherwise.
func (o *XmlNs0EnclaveRunningAttestationAllOf) GetEnclaveProcess() XmlNs0EnclaveProcess {
	if o == nil || o.EnclaveProcess == nil {
		var ret XmlNs0EnclaveProcess
		return ret
	}
	return *o.EnclaveProcess
}

// GetEnclaveProcessOk returns a tuple with the EnclaveProcess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveRunningAttestationAllOf) GetEnclaveProcessOk() (*XmlNs0EnclaveProcess, bool) {
	if o == nil || o.EnclaveProcess == nil {
		return nil, false
	}
	return o.EnclaveProcess, true
}

// HasEnclaveProcess returns a boolean if a field has been set.
func (o *XmlNs0EnclaveRunningAttestationAllOf) HasEnclaveProcess() bool {
	if o != nil && o.EnclaveProcess != nil {
		return true
	}

	return false
}

// SetEnclaveProcess gets a reference to the given XmlNs0EnclaveProcess and assigns it to the EnclaveProcess field.
func (o *XmlNs0EnclaveRunningAttestationAllOf) SetEnclaveProcess(v XmlNs0EnclaveProcess) {
	o.EnclaveProcess = &v
}

// GetNodeInfo returns the NodeInfo field value if set, zero value otherwise.
func (o *XmlNs0EnclaveRunningAttestationAllOf) GetNodeInfo() XmlNs0NodeInfo {
	if o == nil || o.NodeInfo == nil {
		var ret XmlNs0NodeInfo
		return ret
	}
	return *o.NodeInfo
}

// GetNodeInfoOk returns a tuple with the NodeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveRunningAttestationAllOf) GetNodeInfoOk() (*XmlNs0NodeInfo, bool) {
	if o == nil || o.NodeInfo == nil {
		return nil, false
	}
	return o.NodeInfo, true
}

// HasNodeInfo returns a boolean if a field has been set.
func (o *XmlNs0EnclaveRunningAttestationAllOf) HasNodeInfo() bool {
	if o != nil && o.NodeInfo != nil {
		return true
	}

	return false
}

// SetNodeInfo gets a reference to the given XmlNs0NodeInfo and assigns it to the NodeInfo field.
func (o *XmlNs0EnclaveRunningAttestationAllOf) SetNodeInfo(v XmlNs0NodeInfo) {
	o.NodeInfo = &v
}

func (o XmlNs0EnclaveRunningAttestationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnclaveProcess != nil {
		toSerialize["enclaveProcess"] = o.EnclaveProcess
	}
	if o.NodeInfo != nil {
		toSerialize["nodeInfo"] = o.NodeInfo
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0EnclaveRunningAttestationAllOf struct {
	value *XmlNs0EnclaveRunningAttestationAllOf
	isSet bool
}

func (v NullableXmlNs0EnclaveRunningAttestationAllOf) Get() *XmlNs0EnclaveRunningAttestationAllOf {
	return v.value
}

func (v *NullableXmlNs0EnclaveRunningAttestationAllOf) Set(val *XmlNs0EnclaveRunningAttestationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0EnclaveRunningAttestationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0EnclaveRunningAttestationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0EnclaveRunningAttestationAllOf(val *XmlNs0EnclaveRunningAttestationAllOf) *NullableXmlNs0EnclaveRunningAttestationAllOf {
	return &NullableXmlNs0EnclaveRunningAttestationAllOf{value: val, isSet: true}
}

func (v NullableXmlNs0EnclaveRunningAttestationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0EnclaveRunningAttestationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


