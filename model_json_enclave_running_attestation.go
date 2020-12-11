/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonEnclaveRunningAttestation 
type JsonEnclaveRunningAttestation struct {
	NodeInfo *JsonNodeInfo `json:"nodeInfo,omitempty"`
	EnclaveProcess *JsonEnclaveProcess `json:"enclaveProcess,omitempty"`
}

// NewJsonEnclaveRunningAttestation instantiates a new JsonEnclaveRunningAttestation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonEnclaveRunningAttestation() *JsonEnclaveRunningAttestation {
	this := JsonEnclaveRunningAttestation{}
	return &this
}

// NewJsonEnclaveRunningAttestationWithDefaults instantiates a new JsonEnclaveRunningAttestation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonEnclaveRunningAttestationWithDefaults() *JsonEnclaveRunningAttestation {
	this := JsonEnclaveRunningAttestation{}
	return &this
}

// GetNodeInfo returns the NodeInfo field value if set, zero value otherwise.
func (o *JsonEnclaveRunningAttestation) GetNodeInfo() JsonNodeInfo {
	if o == nil || o.NodeInfo == nil {
		var ret JsonNodeInfo
		return ret
	}
	return *o.NodeInfo
}

// GetNodeInfoOk returns a tuple with the NodeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveRunningAttestation) GetNodeInfoOk() (*JsonNodeInfo, bool) {
	if o == nil || o.NodeInfo == nil {
		return nil, false
	}
	return o.NodeInfo, true
}

// HasNodeInfo returns a boolean if a field has been set.
func (o *JsonEnclaveRunningAttestation) HasNodeInfo() bool {
	if o != nil && o.NodeInfo != nil {
		return true
	}

	return false
}

// SetNodeInfo gets a reference to the given JsonNodeInfo and assigns it to the NodeInfo field.
func (o *JsonEnclaveRunningAttestation) SetNodeInfo(v JsonNodeInfo) {
	o.NodeInfo = &v
}

// GetEnclaveProcess returns the EnclaveProcess field value if set, zero value otherwise.
func (o *JsonEnclaveRunningAttestation) GetEnclaveProcess() JsonEnclaveProcess {
	if o == nil || o.EnclaveProcess == nil {
		var ret JsonEnclaveProcess
		return ret
	}
	return *o.EnclaveProcess
}

// GetEnclaveProcessOk returns a tuple with the EnclaveProcess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonEnclaveRunningAttestation) GetEnclaveProcessOk() (*JsonEnclaveProcess, bool) {
	if o == nil || o.EnclaveProcess == nil {
		return nil, false
	}
	return o.EnclaveProcess, true
}

// HasEnclaveProcess returns a boolean if a field has been set.
func (o *JsonEnclaveRunningAttestation) HasEnclaveProcess() bool {
	if o != nil && o.EnclaveProcess != nil {
		return true
	}

	return false
}

// SetEnclaveProcess gets a reference to the given JsonEnclaveProcess and assigns it to the EnclaveProcess field.
func (o *JsonEnclaveRunningAttestation) SetEnclaveProcess(v JsonEnclaveProcess) {
	o.EnclaveProcess = &v
}

func (o JsonEnclaveRunningAttestation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NodeInfo != nil {
		toSerialize["nodeInfo"] = o.NodeInfo
	}
	if o.EnclaveProcess != nil {
		toSerialize["enclaveProcess"] = o.EnclaveProcess
	}
	return json.Marshal(toSerialize)
}

type NullableJsonEnclaveRunningAttestation struct {
	value *JsonEnclaveRunningAttestation
	isSet bool
}

func (v NullableJsonEnclaveRunningAttestation) Get() *JsonEnclaveRunningAttestation {
	return v.value
}

func (v *NullableJsonEnclaveRunningAttestation) Set(val *JsonEnclaveRunningAttestation) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonEnclaveRunningAttestation) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonEnclaveRunningAttestation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonEnclaveRunningAttestation(val *JsonEnclaveRunningAttestation) *NullableJsonEnclaveRunningAttestation {
	return &NullableJsonEnclaveRunningAttestation{value: val, isSet: true}
}

func (v NullableJsonEnclaveRunningAttestation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonEnclaveRunningAttestation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


