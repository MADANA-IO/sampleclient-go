/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.54
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0SignedData 
type XmlNs0SignedData struct {
	// 
	Data *string `json:"data,omitempty"`
	// 
	Fingerpint *string `json:"fingerpint,omitempty"`
	// 
	Signature *string `json:"signature,omitempty"`
}

// NewXmlNs0SignedData instantiates a new XmlNs0SignedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0SignedData() *XmlNs0SignedData {
	this := XmlNs0SignedData{}
	return &this
}

// NewXmlNs0SignedDataWithDefaults instantiates a new XmlNs0SignedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0SignedDataWithDefaults() *XmlNs0SignedData {
	this := XmlNs0SignedData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *XmlNs0SignedData) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0SignedData) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *XmlNs0SignedData) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *XmlNs0SignedData) SetData(v string) {
	o.Data = &v
}

// GetFingerpint returns the Fingerpint field value if set, zero value otherwise.
func (o *XmlNs0SignedData) GetFingerpint() string {
	if o == nil || o.Fingerpint == nil {
		var ret string
		return ret
	}
	return *o.Fingerpint
}

// GetFingerpintOk returns a tuple with the Fingerpint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0SignedData) GetFingerpintOk() (*string, bool) {
	if o == nil || o.Fingerpint == nil {
		return nil, false
	}
	return o.Fingerpint, true
}

// HasFingerpint returns a boolean if a field has been set.
func (o *XmlNs0SignedData) HasFingerpint() bool {
	if o != nil && o.Fingerpint != nil {
		return true
	}

	return false
}

// SetFingerpint gets a reference to the given string and assigns it to the Fingerpint field.
func (o *XmlNs0SignedData) SetFingerpint(v string) {
	o.Fingerpint = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *XmlNs0SignedData) GetSignature() string {
	if o == nil || o.Signature == nil {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0SignedData) GetSignatureOk() (*string, bool) {
	if o == nil || o.Signature == nil {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *XmlNs0SignedData) HasSignature() bool {
	if o != nil && o.Signature != nil {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *XmlNs0SignedData) SetSignature(v string) {
	o.Signature = &v
}

func (o XmlNs0SignedData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Fingerpint != nil {
		toSerialize["fingerpint"] = o.Fingerpint
	}
	if o.Signature != nil {
		toSerialize["signature"] = o.Signature
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0SignedData struct {
	value *XmlNs0SignedData
	isSet bool
}

func (v NullableXmlNs0SignedData) Get() *XmlNs0SignedData {
	return v.value
}

func (v *NullableXmlNs0SignedData) Set(val *XmlNs0SignedData) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0SignedData) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0SignedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0SignedData(val *XmlNs0SignedData) *NullableXmlNs0SignedData {
	return &NullableXmlNs0SignedData{value: val, isSet: true}
}

func (v NullableXmlNs0SignedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0SignedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


