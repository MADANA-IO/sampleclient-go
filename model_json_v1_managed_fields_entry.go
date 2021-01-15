/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.29
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonV1ManagedFieldsEntry 
type JsonV1ManagedFieldsEntry struct {
	// 
	ApiVersion *string `json:"apiVersion,omitempty"`
	// 
	FieldsType *string `json:"fieldsType,omitempty"`
	// 
	Time *float32 `json:"time,omitempty"`
	// 
	Operation *string `json:"operation,omitempty"`
	// 
	Manager *string `json:"manager,omitempty"`
	// 
	FieldsV1 *map[string]interface{} `json:"fieldsV1,omitempty"`
}

// NewJsonV1ManagedFieldsEntry instantiates a new JsonV1ManagedFieldsEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonV1ManagedFieldsEntry() *JsonV1ManagedFieldsEntry {
	this := JsonV1ManagedFieldsEntry{}
	return &this
}

// NewJsonV1ManagedFieldsEntryWithDefaults instantiates a new JsonV1ManagedFieldsEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonV1ManagedFieldsEntryWithDefaults() *JsonV1ManagedFieldsEntry {
	this := JsonV1ManagedFieldsEntry{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *JsonV1ManagedFieldsEntry) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ManagedFieldsEntry) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *JsonV1ManagedFieldsEntry) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *JsonV1ManagedFieldsEntry) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetFieldsType returns the FieldsType field value if set, zero value otherwise.
func (o *JsonV1ManagedFieldsEntry) GetFieldsType() string {
	if o == nil || o.FieldsType == nil {
		var ret string
		return ret
	}
	return *o.FieldsType
}

// GetFieldsTypeOk returns a tuple with the FieldsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ManagedFieldsEntry) GetFieldsTypeOk() (*string, bool) {
	if o == nil || o.FieldsType == nil {
		return nil, false
	}
	return o.FieldsType, true
}

// HasFieldsType returns a boolean if a field has been set.
func (o *JsonV1ManagedFieldsEntry) HasFieldsType() bool {
	if o != nil && o.FieldsType != nil {
		return true
	}

	return false
}

// SetFieldsType gets a reference to the given string and assigns it to the FieldsType field.
func (o *JsonV1ManagedFieldsEntry) SetFieldsType(v string) {
	o.FieldsType = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *JsonV1ManagedFieldsEntry) GetTime() float32 {
	if o == nil || o.Time == nil {
		var ret float32
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ManagedFieldsEntry) GetTimeOk() (*float32, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *JsonV1ManagedFieldsEntry) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given float32 and assigns it to the Time field.
func (o *JsonV1ManagedFieldsEntry) SetTime(v float32) {
	o.Time = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *JsonV1ManagedFieldsEntry) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ManagedFieldsEntry) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *JsonV1ManagedFieldsEntry) HasOperation() bool {
	if o != nil && o.Operation != nil {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *JsonV1ManagedFieldsEntry) SetOperation(v string) {
	o.Operation = &v
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *JsonV1ManagedFieldsEntry) GetManager() string {
	if o == nil || o.Manager == nil {
		var ret string
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ManagedFieldsEntry) GetManagerOk() (*string, bool) {
	if o == nil || o.Manager == nil {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *JsonV1ManagedFieldsEntry) HasManager() bool {
	if o != nil && o.Manager != nil {
		return true
	}

	return false
}

// SetManager gets a reference to the given string and assigns it to the Manager field.
func (o *JsonV1ManagedFieldsEntry) SetManager(v string) {
	o.Manager = &v
}

// GetFieldsV1 returns the FieldsV1 field value if set, zero value otherwise.
func (o *JsonV1ManagedFieldsEntry) GetFieldsV1() map[string]interface{} {
	if o == nil || o.FieldsV1 == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.FieldsV1
}

// GetFieldsV1Ok returns a tuple with the FieldsV1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ManagedFieldsEntry) GetFieldsV1Ok() (*map[string]interface{}, bool) {
	if o == nil || o.FieldsV1 == nil {
		return nil, false
	}
	return o.FieldsV1, true
}

// HasFieldsV1 returns a boolean if a field has been set.
func (o *JsonV1ManagedFieldsEntry) HasFieldsV1() bool {
	if o != nil && o.FieldsV1 != nil {
		return true
	}

	return false
}

// SetFieldsV1 gets a reference to the given map[string]interface{} and assigns it to the FieldsV1 field.
func (o *JsonV1ManagedFieldsEntry) SetFieldsV1(v map[string]interface{}) {
	o.FieldsV1 = &v
}

func (o JsonV1ManagedFieldsEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if o.FieldsType != nil {
		toSerialize["fieldsType"] = o.FieldsType
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}
	if o.Manager != nil {
		toSerialize["manager"] = o.Manager
	}
	if o.FieldsV1 != nil {
		toSerialize["fieldsV1"] = o.FieldsV1
	}
	return json.Marshal(toSerialize)
}

type NullableJsonV1ManagedFieldsEntry struct {
	value *JsonV1ManagedFieldsEntry
	isSet bool
}

func (v NullableJsonV1ManagedFieldsEntry) Get() *JsonV1ManagedFieldsEntry {
	return v.value
}

func (v *NullableJsonV1ManagedFieldsEntry) Set(val *JsonV1ManagedFieldsEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonV1ManagedFieldsEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonV1ManagedFieldsEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonV1ManagedFieldsEntry(val *JsonV1ManagedFieldsEntry) *NullableJsonV1ManagedFieldsEntry {
	return &NullableJsonV1ManagedFieldsEntry{value: val, isSet: true}
}

func (v NullableJsonV1ManagedFieldsEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonV1ManagedFieldsEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


