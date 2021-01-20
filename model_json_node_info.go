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

// JsonNodeInfo 
type JsonNodeInfo struct {
	// 
	Processors *[]string `json:"processors,omitempty"`
	// 
	HardwareBaseboard *string `json:"hardwareBaseboard,omitempty"`
	SgxInfo *JsonSGXInfo `json:"sgxInfo,omitempty"`
	// 
	ConnectionURL *string `json:"connectionURL,omitempty"`
	// 
	Owner *string `json:"owner,omitempty"`
	// 
	CpuModel *string `json:"cpuModel,omitempty"`
	// 
	HardwareFirmware *string `json:"hardwareFirmware,omitempty"`
	// 
	Memory *string `json:"memory,omitempty"`
	// 
	CpuLogicalCount *int32 `json:"cpuLogicalCount,omitempty"`
	// 
	OperatingSystem *string `json:"operatingSystem,omitempty"`
	// 
	CpuPhysicalCores *int32 `json:"cpuPhysicalCores,omitempty"`
	// 
	PublicKey *string `json:"publicKey,omitempty"`
	// 
	CpuFrequency *string `json:"cpuFrequency,omitempty"`
	// 
	OperatingSystemUptime *float32 `json:"operatingSystemUptime,omitempty"`
	// 
	CpuFamily *string `json:"cpuFamily,omitempty"`
	// 
	Status *string `json:"status,omitempty"`
	IpfsInfo *JsonIPFSSystemInfo `json:"ipfsInfo,omitempty"`
}

// NewJsonNodeInfo instantiates a new JsonNodeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonNodeInfo() *JsonNodeInfo {
	this := JsonNodeInfo{}
	return &this
}

// NewJsonNodeInfoWithDefaults instantiates a new JsonNodeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonNodeInfoWithDefaults() *JsonNodeInfo {
	this := JsonNodeInfo{}
	return &this
}

// GetProcessors returns the Processors field value if set, zero value otherwise.
func (o *JsonNodeInfo) GetProcessors() []string {
	if o == nil || o.Processors == nil {
		var ret []string
		return ret
	}
	return *o.Processors
}

// GetProcessorsOk returns a tuple with the Processors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeInfo) GetProcessorsOk() (*[]string, bool) {
	if o == nil || o.Processors == nil {
		return nil, false
	}
	return o.Processors, true
}

// HasProcessors returns a boolean if a field has been set.
func (o *JsonNodeInfo) HasProcessors() bool {
	if o != nil && o.Processors != nil {
		return true
	}

	return false
}

// SetProcessors gets a reference to the given []string and assigns it to the Processors field.
func (o *JsonNodeInfo) SetProcessors(v []string) {
	o.Processors = &v
}

// GetHardwareBaseboard returns the HardwareBaseboard field value if set, zero value otherwise.
func (o *JsonNodeInfo) GetHardwareBaseboard() string {
	if o == nil || o.HardwareBaseboard == nil {
		var ret string
		return ret
	}
	return *o.HardwareBaseboard
}

// GetHardwareBaseboardOk returns a tuple with the HardwareBaseboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeInfo) GetHardwareBaseboardOk() (*string, bool) {
	if o == nil || o.HardwareBaseboard == nil {
		return nil, false
	}
	return o.HardwareBaseboard, true
}

// HasHardwareBaseboard returns a boolean if a field has been set.
func (o *JsonNodeInfo) HasHardwareBaseboard() bool {
	if o != nil && o.HardwareBaseboard != nil {
		return true
	}

	return false
}

// SetHardwareBaseboard gets a reference to the given string and assigns it to the HardwareBaseboard field.
func (o *JsonNodeInfo) SetHardwareBaseboard(v string) {
	o.HardwareBaseboard = &v
}

// GetSgxInfo returns the SgxInfo field value if set, zero value otherwise.
func (o *JsonNodeInfo) GetSgxInfo() JsonSGXInfo {
	if o == nil || o.SgxInfo == nil {
		var ret JsonSGXInfo
		return ret
	}
	return *o.SgxInfo
}

// GetSgxInfoOk returns a tuple with the SgxInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeInfo) GetSgxInfoOk() (*JsonSGXInfo, bool) {
	if o == nil || o.SgxInfo == nil {
		return nil, false
	}
	return o.SgxInfo, true
}

// HasSgxInfo returns a boolean if a field has been set.
func (o *JsonNodeInfo) HasSgxInfo() bool {
	if o != nil && o.SgxInfo != nil {
		return true
	}

	return false
}

// SetSgxInfo gets a reference to the given JsonSGXInfo and assigns it to the SgxInfo field.
func (o *JsonNodeInfo) SetSgxInfo(v JsonSGXInfo) {
	o.SgxInfo = &v
}

// GetConnectionURL returns the ConnectionURL field value if set, zero value otherwise.
func (o *JsonNodeInfo) GetConnectionURL() string {
	if o == nil || o.ConnectionURL == nil {
		var ret string
		return ret
	}
	return *o.ConnectionURL
}

// GetConnectionURLOk returns a tuple with the ConnectionURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeInfo) GetConnectionURLOk() (*string, bool) {
	if o == nil || o.ConnectionURL == nil {
		return nil, false
	}
	return o.ConnectionURL, true
}

// HasConnectionURL returns a boolean if a field has been set.
func (o *JsonNodeInfo) HasConnectionURL() bool {
	if o != nil && o.ConnectionURL != nil {
		return true
	}

	return false
}

// SetConnectionURL gets a reference to the given string and assigns it to the ConnectionURL field.
func (o *JsonNodeInfo) SetConnectionURL(v string) {
	o.ConnectionURL = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *JsonNodeInfo) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeInfo) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *JsonNodeInfo) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *JsonNodeInfo) SetOwner(v string) {
	o.Owner = &v
}

// GetCpuModel returns the CpuModel field value if set, zero value otherwise.
func (o *JsonNodeInfo) GetCpuModel() string {
	if o == nil || o.CpuModel == nil {
		var ret string
		return ret
	}
	return *o.CpuModel
}

// GetCpuModelOk returns a tuple with the CpuModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeInfo) GetCpuModelOk() (*string, bool) {
	if o == nil || o.CpuModel == nil {
		return nil, false
	}
	return o.CpuModel, true
}

// HasCpuModel returns a boolean if a field has been set.
func (o *JsonNodeInfo) HasCpuModel() bool {
	if o != nil && o.CpuModel != nil {
		return true
	}

	return false
}

// SetCpuModel gets a reference to the given string and assigns it to the CpuModel field.
func (o *JsonNodeInfo) SetCpuModel(v string) {
	o.CpuModel = &v
}

// GetHardwareFirmware returns the HardwareFirmware field value if set, zero value otherwise.
func (o *JsonNodeInfo) GetHardwareFirmware() string {
	if o == nil || o.HardwareFirmware == nil {
		var ret string
		return ret
	}
	return *o.HardwareFirmware
}

// GetHardwareFirmwareOk returns a tuple with the HardwareFirmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeInfo) GetHardwareFirmwareOk() (*string, bool) {
	if o == nil || o.HardwareFirmware == nil {
		return nil, false
	}
	return o.HardwareFirmware, true
}

// HasHardwareFirmware returns a boolean if a field has been set.
func (o *JsonNodeInfo) HasHardwareFirmware() bool {
	if o != nil && o.HardwareFirmware != nil {
		return true
	}

	return false
}

// SetHardwareFirmware gets a reference to the given string and assigns it to the HardwareFirmware field.
func (o *JsonNodeInfo) SetHardwareFirmware(v string) {
	o.HardwareFirmware = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *JsonNodeInfo) GetMemory() string {
	if o == nil || o.Memory == nil {
		var ret string
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeInfo) GetMemoryOk() (*string, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *JsonNodeInfo) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given string and assigns it to the Memory field.
func (o *JsonNodeInfo) SetMemory(v string) {
	o.Memory = &v
}

// GetCpuLogicalCount returns the CpuLogicalCount field value if set, zero value otherwise.
func (o *JsonNodeInfo) GetCpuLogicalCount() int32 {
	if o == nil || o.CpuLogicalCount == nil {
		var ret int32
		return ret
	}
	return *o.CpuLogicalCount
}

// GetCpuLogicalCountOk returns a tuple with the CpuLogicalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeInfo) GetCpuLogicalCountOk() (*int32, bool) {
	if o == nil || o.CpuLogicalCount == nil {
		return nil, false
	}
	return o.CpuLogicalCount, true
}

// HasCpuLogicalCount returns a boolean if a field has been set.
func (o *JsonNodeInfo) HasCpuLogicalCount() bool {
	if o != nil && o.CpuLogicalCount != nil {
		return true
	}

	return false
}

// SetCpuLogicalCount gets a reference to the given int32 and assigns it to the CpuLogicalCount field.
func (o *JsonNodeInfo) SetCpuLogicalCount(v int32) {
	o.CpuLogicalCount = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *JsonNodeInfo) GetOperatingSystem() string {
	if o == nil || o.OperatingSystem == nil {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeInfo) GetOperatingSystemOk() (*string, bool) {
	if o == nil || o.OperatingSystem == nil {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *JsonNodeInfo) HasOperatingSystem() bool {
	if o != nil && o.OperatingSystem != nil {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *JsonNodeInfo) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

// GetCpuPhysicalCores returns the CpuPhysicalCores field value if set, zero value otherwise.
func (o *JsonNodeInfo) GetCpuPhysicalCores() int32 {
	if o == nil || o.CpuPhysicalCores == nil {
		var ret int32
		return ret
	}
	return *o.CpuPhysicalCores
}

// GetCpuPhysicalCoresOk returns a tuple with the CpuPhysicalCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeInfo) GetCpuPhysicalCoresOk() (*int32, bool) {
	if o == nil || o.CpuPhysicalCores == nil {
		return nil, false
	}
	return o.CpuPhysicalCores, true
}

// HasCpuPhysicalCores returns a boolean if a field has been set.
func (o *JsonNodeInfo) HasCpuPhysicalCores() bool {
	if o != nil && o.CpuPhysicalCores != nil {
		return true
	}

	return false
}

// SetCpuPhysicalCores gets a reference to the given int32 and assigns it to the CpuPhysicalCores field.
func (o *JsonNodeInfo) SetCpuPhysicalCores(v int32) {
	o.CpuPhysicalCores = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *JsonNodeInfo) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeInfo) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *JsonNodeInfo) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *JsonNodeInfo) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetCpuFrequency returns the CpuFrequency field value if set, zero value otherwise.
func (o *JsonNodeInfo) GetCpuFrequency() string {
	if o == nil || o.CpuFrequency == nil {
		var ret string
		return ret
	}
	return *o.CpuFrequency
}

// GetCpuFrequencyOk returns a tuple with the CpuFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeInfo) GetCpuFrequencyOk() (*string, bool) {
	if o == nil || o.CpuFrequency == nil {
		return nil, false
	}
	return o.CpuFrequency, true
}

// HasCpuFrequency returns a boolean if a field has been set.
func (o *JsonNodeInfo) HasCpuFrequency() bool {
	if o != nil && o.CpuFrequency != nil {
		return true
	}

	return false
}

// SetCpuFrequency gets a reference to the given string and assigns it to the CpuFrequency field.
func (o *JsonNodeInfo) SetCpuFrequency(v string) {
	o.CpuFrequency = &v
}

// GetOperatingSystemUptime returns the OperatingSystemUptime field value if set, zero value otherwise.
func (o *JsonNodeInfo) GetOperatingSystemUptime() float32 {
	if o == nil || o.OperatingSystemUptime == nil {
		var ret float32
		return ret
	}
	return *o.OperatingSystemUptime
}

// GetOperatingSystemUptimeOk returns a tuple with the OperatingSystemUptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeInfo) GetOperatingSystemUptimeOk() (*float32, bool) {
	if o == nil || o.OperatingSystemUptime == nil {
		return nil, false
	}
	return o.OperatingSystemUptime, true
}

// HasOperatingSystemUptime returns a boolean if a field has been set.
func (o *JsonNodeInfo) HasOperatingSystemUptime() bool {
	if o != nil && o.OperatingSystemUptime != nil {
		return true
	}

	return false
}

// SetOperatingSystemUptime gets a reference to the given float32 and assigns it to the OperatingSystemUptime field.
func (o *JsonNodeInfo) SetOperatingSystemUptime(v float32) {
	o.OperatingSystemUptime = &v
}

// GetCpuFamily returns the CpuFamily field value if set, zero value otherwise.
func (o *JsonNodeInfo) GetCpuFamily() string {
	if o == nil || o.CpuFamily == nil {
		var ret string
		return ret
	}
	return *o.CpuFamily
}

// GetCpuFamilyOk returns a tuple with the CpuFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeInfo) GetCpuFamilyOk() (*string, bool) {
	if o == nil || o.CpuFamily == nil {
		return nil, false
	}
	return o.CpuFamily, true
}

// HasCpuFamily returns a boolean if a field has been set.
func (o *JsonNodeInfo) HasCpuFamily() bool {
	if o != nil && o.CpuFamily != nil {
		return true
	}

	return false
}

// SetCpuFamily gets a reference to the given string and assigns it to the CpuFamily field.
func (o *JsonNodeInfo) SetCpuFamily(v string) {
	o.CpuFamily = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *JsonNodeInfo) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeInfo) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *JsonNodeInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *JsonNodeInfo) SetStatus(v string) {
	o.Status = &v
}

// GetIpfsInfo returns the IpfsInfo field value if set, zero value otherwise.
func (o *JsonNodeInfo) GetIpfsInfo() JsonIPFSSystemInfo {
	if o == nil || o.IpfsInfo == nil {
		var ret JsonIPFSSystemInfo
		return ret
	}
	return *o.IpfsInfo
}

// GetIpfsInfoOk returns a tuple with the IpfsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonNodeInfo) GetIpfsInfoOk() (*JsonIPFSSystemInfo, bool) {
	if o == nil || o.IpfsInfo == nil {
		return nil, false
	}
	return o.IpfsInfo, true
}

// HasIpfsInfo returns a boolean if a field has been set.
func (o *JsonNodeInfo) HasIpfsInfo() bool {
	if o != nil && o.IpfsInfo != nil {
		return true
	}

	return false
}

// SetIpfsInfo gets a reference to the given JsonIPFSSystemInfo and assigns it to the IpfsInfo field.
func (o *JsonNodeInfo) SetIpfsInfo(v JsonIPFSSystemInfo) {
	o.IpfsInfo = &v
}

func (o JsonNodeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Processors != nil {
		toSerialize["processors"] = o.Processors
	}
	if o.HardwareBaseboard != nil {
		toSerialize["hardwareBaseboard"] = o.HardwareBaseboard
	}
	if o.SgxInfo != nil {
		toSerialize["sgxInfo"] = o.SgxInfo
	}
	if o.ConnectionURL != nil {
		toSerialize["connectionURL"] = o.ConnectionURL
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.CpuModel != nil {
		toSerialize["cpuModel"] = o.CpuModel
	}
	if o.HardwareFirmware != nil {
		toSerialize["hardwareFirmware"] = o.HardwareFirmware
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.CpuLogicalCount != nil {
		toSerialize["cpuLogicalCount"] = o.CpuLogicalCount
	}
	if o.OperatingSystem != nil {
		toSerialize["operatingSystem"] = o.OperatingSystem
	}
	if o.CpuPhysicalCores != nil {
		toSerialize["cpuPhysicalCores"] = o.CpuPhysicalCores
	}
	if o.PublicKey != nil {
		toSerialize["publicKey"] = o.PublicKey
	}
	if o.CpuFrequency != nil {
		toSerialize["cpuFrequency"] = o.CpuFrequency
	}
	if o.OperatingSystemUptime != nil {
		toSerialize["operatingSystemUptime"] = o.OperatingSystemUptime
	}
	if o.CpuFamily != nil {
		toSerialize["cpuFamily"] = o.CpuFamily
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.IpfsInfo != nil {
		toSerialize["ipfsInfo"] = o.IpfsInfo
	}
	return json.Marshal(toSerialize)
}

type NullableJsonNodeInfo struct {
	value *JsonNodeInfo
	isSet bool
}

func (v NullableJsonNodeInfo) Get() *JsonNodeInfo {
	return v.value
}

func (v *NullableJsonNodeInfo) Set(val *JsonNodeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonNodeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonNodeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonNodeInfo(val *JsonNodeInfo) *NullableJsonNodeInfo {
	return &NullableJsonNodeInfo{value: val, isSet: true}
}

func (v NullableJsonNodeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonNodeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


