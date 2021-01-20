/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.41
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonV1EventSeries 
type JsonV1EventSeries struct {
	// 
	Count *float32 `json:"count,omitempty"`
	// 
	LastObservedTime *float32 `json:"lastObservedTime,omitempty"`
	// 
	State *string `json:"state,omitempty"`
}

// NewJsonV1EventSeries instantiates a new JsonV1EventSeries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonV1EventSeries() *JsonV1EventSeries {
	this := JsonV1EventSeries{}
	return &this
}

// NewJsonV1EventSeriesWithDefaults instantiates a new JsonV1EventSeries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonV1EventSeriesWithDefaults() *JsonV1EventSeries {
	this := JsonV1EventSeries{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *JsonV1EventSeries) GetCount() float32 {
	if o == nil || o.Count == nil {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1EventSeries) GetCountOk() (*float32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *JsonV1EventSeries) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *JsonV1EventSeries) SetCount(v float32) {
	o.Count = &v
}

// GetLastObservedTime returns the LastObservedTime field value if set, zero value otherwise.
func (o *JsonV1EventSeries) GetLastObservedTime() float32 {
	if o == nil || o.LastObservedTime == nil {
		var ret float32
		return ret
	}
	return *o.LastObservedTime
}

// GetLastObservedTimeOk returns a tuple with the LastObservedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1EventSeries) GetLastObservedTimeOk() (*float32, bool) {
	if o == nil || o.LastObservedTime == nil {
		return nil, false
	}
	return o.LastObservedTime, true
}

// HasLastObservedTime returns a boolean if a field has been set.
func (o *JsonV1EventSeries) HasLastObservedTime() bool {
	if o != nil && o.LastObservedTime != nil {
		return true
	}

	return false
}

// SetLastObservedTime gets a reference to the given float32 and assigns it to the LastObservedTime field.
func (o *JsonV1EventSeries) SetLastObservedTime(v float32) {
	o.LastObservedTime = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *JsonV1EventSeries) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1EventSeries) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *JsonV1EventSeries) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *JsonV1EventSeries) SetState(v string) {
	o.State = &v
}

func (o JsonV1EventSeries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.LastObservedTime != nil {
		toSerialize["lastObservedTime"] = o.LastObservedTime
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableJsonV1EventSeries struct {
	value *JsonV1EventSeries
	isSet bool
}

func (v NullableJsonV1EventSeries) Get() *JsonV1EventSeries {
	return v.value
}

func (v *NullableJsonV1EventSeries) Set(val *JsonV1EventSeries) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonV1EventSeries) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonV1EventSeries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonV1EventSeries(val *JsonV1EventSeries) *NullableJsonV1EventSeries {
	return &NullableJsonV1EventSeries{value: val, isSet: true}
}

func (v NullableJsonV1EventSeries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonV1EventSeries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


