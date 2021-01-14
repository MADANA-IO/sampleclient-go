/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.0.1-master.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonV1ListMeta 
type JsonV1ListMeta struct {
	// 
	Continue *string `json:"continue,omitempty"`
	// 
	RemainingItemCount *float32 `json:"remainingItemCount,omitempty"`
	// 
	SelfLink *string `json:"selfLink,omitempty"`
	// 
	ResourceVersion *string `json:"resourceVersion,omitempty"`
}

// NewJsonV1ListMeta instantiates a new JsonV1ListMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonV1ListMeta() *JsonV1ListMeta {
	this := JsonV1ListMeta{}
	return &this
}

// NewJsonV1ListMetaWithDefaults instantiates a new JsonV1ListMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonV1ListMetaWithDefaults() *JsonV1ListMeta {
	this := JsonV1ListMeta{}
	return &this
}

// GetContinue returns the Continue field value if set, zero value otherwise.
func (o *JsonV1ListMeta) GetContinue() string {
	if o == nil || o.Continue == nil {
		var ret string
		return ret
	}
	return *o.Continue
}

// GetContinueOk returns a tuple with the Continue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ListMeta) GetContinueOk() (*string, bool) {
	if o == nil || o.Continue == nil {
		return nil, false
	}
	return o.Continue, true
}

// HasContinue returns a boolean if a field has been set.
func (o *JsonV1ListMeta) HasContinue() bool {
	if o != nil && o.Continue != nil {
		return true
	}

	return false
}

// SetContinue gets a reference to the given string and assigns it to the Continue field.
func (o *JsonV1ListMeta) SetContinue(v string) {
	o.Continue = &v
}

// GetRemainingItemCount returns the RemainingItemCount field value if set, zero value otherwise.
func (o *JsonV1ListMeta) GetRemainingItemCount() float32 {
	if o == nil || o.RemainingItemCount == nil {
		var ret float32
		return ret
	}
	return *o.RemainingItemCount
}

// GetRemainingItemCountOk returns a tuple with the RemainingItemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ListMeta) GetRemainingItemCountOk() (*float32, bool) {
	if o == nil || o.RemainingItemCount == nil {
		return nil, false
	}
	return o.RemainingItemCount, true
}

// HasRemainingItemCount returns a boolean if a field has been set.
func (o *JsonV1ListMeta) HasRemainingItemCount() bool {
	if o != nil && o.RemainingItemCount != nil {
		return true
	}

	return false
}

// SetRemainingItemCount gets a reference to the given float32 and assigns it to the RemainingItemCount field.
func (o *JsonV1ListMeta) SetRemainingItemCount(v float32) {
	o.RemainingItemCount = &v
}

// GetSelfLink returns the SelfLink field value if set, zero value otherwise.
func (o *JsonV1ListMeta) GetSelfLink() string {
	if o == nil || o.SelfLink == nil {
		var ret string
		return ret
	}
	return *o.SelfLink
}

// GetSelfLinkOk returns a tuple with the SelfLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ListMeta) GetSelfLinkOk() (*string, bool) {
	if o == nil || o.SelfLink == nil {
		return nil, false
	}
	return o.SelfLink, true
}

// HasSelfLink returns a boolean if a field has been set.
func (o *JsonV1ListMeta) HasSelfLink() bool {
	if o != nil && o.SelfLink != nil {
		return true
	}

	return false
}

// SetSelfLink gets a reference to the given string and assigns it to the SelfLink field.
func (o *JsonV1ListMeta) SetSelfLink(v string) {
	o.SelfLink = &v
}

// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *JsonV1ListMeta) GetResourceVersion() string {
	if o == nil || o.ResourceVersion == nil {
		var ret string
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ListMeta) GetResourceVersionOk() (*string, bool) {
	if o == nil || o.ResourceVersion == nil {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *JsonV1ListMeta) HasResourceVersion() bool {
	if o != nil && o.ResourceVersion != nil {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given string and assigns it to the ResourceVersion field.
func (o *JsonV1ListMeta) SetResourceVersion(v string) {
	o.ResourceVersion = &v
}

func (o JsonV1ListMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Continue != nil {
		toSerialize["continue"] = o.Continue
	}
	if o.RemainingItemCount != nil {
		toSerialize["remainingItemCount"] = o.RemainingItemCount
	}
	if o.SelfLink != nil {
		toSerialize["selfLink"] = o.SelfLink
	}
	if o.ResourceVersion != nil {
		toSerialize["resourceVersion"] = o.ResourceVersion
	}
	return json.Marshal(toSerialize)
}

type NullableJsonV1ListMeta struct {
	value *JsonV1ListMeta
	isSet bool
}

func (v NullableJsonV1ListMeta) Get() *JsonV1ListMeta {
	return v.value
}

func (v *NullableJsonV1ListMeta) Set(val *JsonV1ListMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonV1ListMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonV1ListMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonV1ListMeta(val *JsonV1ListMeta) *NullableJsonV1ListMeta {
	return &NullableJsonV1ListMeta{value: val, isSet: true}
}

func (v NullableJsonV1ListMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonV1ListMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


