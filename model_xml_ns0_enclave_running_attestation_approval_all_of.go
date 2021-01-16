/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.36
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0EnclaveRunningAttestationApprovalAllOf struct for XmlNs0EnclaveRunningAttestationApprovalAllOf
type XmlNs0EnclaveRunningAttestationApprovalAllOf struct {
	// 
	Approved *string `json:"approved,omitempty"`
}

// NewXmlNs0EnclaveRunningAttestationApprovalAllOf instantiates a new XmlNs0EnclaveRunningAttestationApprovalAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0EnclaveRunningAttestationApprovalAllOf() *XmlNs0EnclaveRunningAttestationApprovalAllOf {
	this := XmlNs0EnclaveRunningAttestationApprovalAllOf{}
	return &this
}

// NewXmlNs0EnclaveRunningAttestationApprovalAllOfWithDefaults instantiates a new XmlNs0EnclaveRunningAttestationApprovalAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0EnclaveRunningAttestationApprovalAllOfWithDefaults() *XmlNs0EnclaveRunningAttestationApprovalAllOf {
	this := XmlNs0EnclaveRunningAttestationApprovalAllOf{}
	return &this
}

// GetApproved returns the Approved field value if set, zero value otherwise.
func (o *XmlNs0EnclaveRunningAttestationApprovalAllOf) GetApproved() string {
	if o == nil || o.Approved == nil {
		var ret string
		return ret
	}
	return *o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveRunningAttestationApprovalAllOf) GetApprovedOk() (*string, bool) {
	if o == nil || o.Approved == nil {
		return nil, false
	}
	return o.Approved, true
}

// HasApproved returns a boolean if a field has been set.
func (o *XmlNs0EnclaveRunningAttestationApprovalAllOf) HasApproved() bool {
	if o != nil && o.Approved != nil {
		return true
	}

	return false
}

// SetApproved gets a reference to the given string and assigns it to the Approved field.
func (o *XmlNs0EnclaveRunningAttestationApprovalAllOf) SetApproved(v string) {
	o.Approved = &v
}

func (o XmlNs0EnclaveRunningAttestationApprovalAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Approved != nil {
		toSerialize["approved"] = o.Approved
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0EnclaveRunningAttestationApprovalAllOf struct {
	value *XmlNs0EnclaveRunningAttestationApprovalAllOf
	isSet bool
}

func (v NullableXmlNs0EnclaveRunningAttestationApprovalAllOf) Get() *XmlNs0EnclaveRunningAttestationApprovalAllOf {
	return v.value
}

func (v *NullableXmlNs0EnclaveRunningAttestationApprovalAllOf) Set(val *XmlNs0EnclaveRunningAttestationApprovalAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0EnclaveRunningAttestationApprovalAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0EnclaveRunningAttestationApprovalAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0EnclaveRunningAttestationApprovalAllOf(val *XmlNs0EnclaveRunningAttestationApprovalAllOf) *NullableXmlNs0EnclaveRunningAttestationApprovalAllOf {
	return &NullableXmlNs0EnclaveRunningAttestationApprovalAllOf{value: val, isSet: true}
}

func (v NullableXmlNs0EnclaveRunningAttestationApprovalAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0EnclaveRunningAttestationApprovalAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


