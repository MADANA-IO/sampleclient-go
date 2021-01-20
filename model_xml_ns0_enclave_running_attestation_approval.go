/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.39
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0EnclaveRunningAttestationApproval 
type XmlNs0EnclaveRunningAttestationApproval struct {
	EnclaveProcess *XmlNs0EnclaveProcess `json:"enclaveProcess,omitempty"`
	NodeInfo *XmlNs0NodeInfo `json:"nodeInfo,omitempty"`
	// 
	Approved *string `json:"approved,omitempty"`
}

// NewXmlNs0EnclaveRunningAttestationApproval instantiates a new XmlNs0EnclaveRunningAttestationApproval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0EnclaveRunningAttestationApproval() *XmlNs0EnclaveRunningAttestationApproval {
	this := XmlNs0EnclaveRunningAttestationApproval{}
	return &this
}

// NewXmlNs0EnclaveRunningAttestationApprovalWithDefaults instantiates a new XmlNs0EnclaveRunningAttestationApproval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0EnclaveRunningAttestationApprovalWithDefaults() *XmlNs0EnclaveRunningAttestationApproval {
	this := XmlNs0EnclaveRunningAttestationApproval{}
	return &this
}

// GetEnclaveProcess returns the EnclaveProcess field value if set, zero value otherwise.
func (o *XmlNs0EnclaveRunningAttestationApproval) GetEnclaveProcess() XmlNs0EnclaveProcess {
	if o == nil || o.EnclaveProcess == nil {
		var ret XmlNs0EnclaveProcess
		return ret
	}
	return *o.EnclaveProcess
}

// GetEnclaveProcessOk returns a tuple with the EnclaveProcess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveRunningAttestationApproval) GetEnclaveProcessOk() (*XmlNs0EnclaveProcess, bool) {
	if o == nil || o.EnclaveProcess == nil {
		return nil, false
	}
	return o.EnclaveProcess, true
}

// HasEnclaveProcess returns a boolean if a field has been set.
func (o *XmlNs0EnclaveRunningAttestationApproval) HasEnclaveProcess() bool {
	if o != nil && o.EnclaveProcess != nil {
		return true
	}

	return false
}

// SetEnclaveProcess gets a reference to the given XmlNs0EnclaveProcess and assigns it to the EnclaveProcess field.
func (o *XmlNs0EnclaveRunningAttestationApproval) SetEnclaveProcess(v XmlNs0EnclaveProcess) {
	o.EnclaveProcess = &v
}

// GetNodeInfo returns the NodeInfo field value if set, zero value otherwise.
func (o *XmlNs0EnclaveRunningAttestationApproval) GetNodeInfo() XmlNs0NodeInfo {
	if o == nil || o.NodeInfo == nil {
		var ret XmlNs0NodeInfo
		return ret
	}
	return *o.NodeInfo
}

// GetNodeInfoOk returns a tuple with the NodeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveRunningAttestationApproval) GetNodeInfoOk() (*XmlNs0NodeInfo, bool) {
	if o == nil || o.NodeInfo == nil {
		return nil, false
	}
	return o.NodeInfo, true
}

// HasNodeInfo returns a boolean if a field has been set.
func (o *XmlNs0EnclaveRunningAttestationApproval) HasNodeInfo() bool {
	if o != nil && o.NodeInfo != nil {
		return true
	}

	return false
}

// SetNodeInfo gets a reference to the given XmlNs0NodeInfo and assigns it to the NodeInfo field.
func (o *XmlNs0EnclaveRunningAttestationApproval) SetNodeInfo(v XmlNs0NodeInfo) {
	o.NodeInfo = &v
}

// GetApproved returns the Approved field value if set, zero value otherwise.
func (o *XmlNs0EnclaveRunningAttestationApproval) GetApproved() string {
	if o == nil || o.Approved == nil {
		var ret string
		return ret
	}
	return *o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveRunningAttestationApproval) GetApprovedOk() (*string, bool) {
	if o == nil || o.Approved == nil {
		return nil, false
	}
	return o.Approved, true
}

// HasApproved returns a boolean if a field has been set.
func (o *XmlNs0EnclaveRunningAttestationApproval) HasApproved() bool {
	if o != nil && o.Approved != nil {
		return true
	}

	return false
}

// SetApproved gets a reference to the given string and assigns it to the Approved field.
func (o *XmlNs0EnclaveRunningAttestationApproval) SetApproved(v string) {
	o.Approved = &v
}

func (o XmlNs0EnclaveRunningAttestationApproval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnclaveProcess != nil {
		toSerialize["enclaveProcess"] = o.EnclaveProcess
	}
	if o.NodeInfo != nil {
		toSerialize["nodeInfo"] = o.NodeInfo
	}
	if o.Approved != nil {
		toSerialize["approved"] = o.Approved
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0EnclaveRunningAttestationApproval struct {
	value *XmlNs0EnclaveRunningAttestationApproval
	isSet bool
}

func (v NullableXmlNs0EnclaveRunningAttestationApproval) Get() *XmlNs0EnclaveRunningAttestationApproval {
	return v.value
}

func (v *NullableXmlNs0EnclaveRunningAttestationApproval) Set(val *XmlNs0EnclaveRunningAttestationApproval) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0EnclaveRunningAttestationApproval) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0EnclaveRunningAttestationApproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0EnclaveRunningAttestationApproval(val *XmlNs0EnclaveRunningAttestationApproval) *NullableXmlNs0EnclaveRunningAttestationApproval {
	return &NullableXmlNs0EnclaveRunningAttestationApproval{value: val, isSet: true}
}

func (v NullableXmlNs0EnclaveRunningAttestationApproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0EnclaveRunningAttestationApproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


