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

// JsonV1EventList 
type JsonV1EventList struct {
	// 
	ApiVersion *string `json:"apiVersion,omitempty"`
	// 
	Items *[]JsonV1Event `json:"items,omitempty"`
	Metadata *JsonV1ListMeta `json:"metadata,omitempty"`
	// 
	Kind *string `json:"kind,omitempty"`
}

// NewJsonV1EventList instantiates a new JsonV1EventList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonV1EventList() *JsonV1EventList {
	this := JsonV1EventList{}
	return &this
}

// NewJsonV1EventListWithDefaults instantiates a new JsonV1EventList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonV1EventListWithDefaults() *JsonV1EventList {
	this := JsonV1EventList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *JsonV1EventList) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1EventList) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *JsonV1EventList) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *JsonV1EventList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *JsonV1EventList) GetItems() []JsonV1Event {
	if o == nil || o.Items == nil {
		var ret []JsonV1Event
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1EventList) GetItemsOk() (*[]JsonV1Event, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *JsonV1EventList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []JsonV1Event and assigns it to the Items field.
func (o *JsonV1EventList) SetItems(v []JsonV1Event) {
	o.Items = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *JsonV1EventList) GetMetadata() JsonV1ListMeta {
	if o == nil || o.Metadata == nil {
		var ret JsonV1ListMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1EventList) GetMetadataOk() (*JsonV1ListMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *JsonV1EventList) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given JsonV1ListMeta and assigns it to the Metadata field.
func (o *JsonV1EventList) SetMetadata(v JsonV1ListMeta) {
	o.Metadata = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *JsonV1EventList) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1EventList) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *JsonV1EventList) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *JsonV1EventList) SetKind(v string) {
	o.Kind = &v
}

func (o JsonV1EventList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	return json.Marshal(toSerialize)
}

type NullableJsonV1EventList struct {
	value *JsonV1EventList
	isSet bool
}

func (v NullableJsonV1EventList) Get() *JsonV1EventList {
	return v.value
}

func (v *NullableJsonV1EventList) Set(val *JsonV1EventList) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonV1EventList) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonV1EventList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonV1EventList(val *JsonV1EventList) *NullableJsonV1EventList {
	return &NullableJsonV1EventList{value: val, isSet: true}
}

func (v NullableJsonV1EventList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonV1EventList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


