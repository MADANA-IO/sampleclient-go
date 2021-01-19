/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.38
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonNetworkInterface 
type JsonNetworkInterface struct {
	// 
	Name *string `json:"name,omitempty"`
	// 
	Address *string `json:"address,omitempty"`
}

// NewJsonNetworkInterface instantiates a new JsonNetworkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonNetworkInterface() *JsonNetworkInterface {
	this := JsonNetworkInterface{}
	return &this
}

// NewJsonNetworkInterfaceWithDefaults instantiates a new JsonNetworkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonNetworkInterfaceWithDefaults() *JsonNetworkInterface {
	this := JsonNetworkInterface{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *JsonNetworkInterface) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNetworkInterface) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *JsonNetworkInterface) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *JsonNetworkInterface) SetName(v string) {
	o.Name = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *JsonNetworkInterface) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNetworkInterface) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *JsonNetworkInterface) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *JsonNetworkInterface) SetAddress(v string) {
	o.Address = &v
}

func (o JsonNetworkInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableJsonNetworkInterface struct {
	value *JsonNetworkInterface
	isSet bool
}

func (v NullableJsonNetworkInterface) Get() *JsonNetworkInterface {
	return v.value
}

func (v *NullableJsonNetworkInterface) Set(val *JsonNetworkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonNetworkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonNetworkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonNetworkInterface(val *JsonNetworkInterface) *NullableJsonNetworkInterface {
	return &NullableJsonNetworkInterface{value: val, isSet: true}
}

func (v NullableJsonNetworkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonNetworkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


