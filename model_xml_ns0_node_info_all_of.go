/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.40
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0NodeInfoAllOf struct for XmlNs0NodeInfoAllOf
type XmlNs0NodeInfoAllOf struct {
	// 
	ConnectionURL *string `json:"connectionURL,omitempty"`
	// 
	CpuFamily *string `json:"cpuFamily,omitempty"`
	// 
	CpuFrequency *string `json:"cpuFrequency,omitempty"`
	// 
	CpuLogicalCount *int32 `json:"cpuLogicalCount,omitempty"`
	// 
	CpuModel *string `json:"cpuModel,omitempty"`
	// 
	CpuPhysicalCores *int32 `json:"cpuPhysicalCores,omitempty"`
	// 
	HardwareBaseboard *string `json:"hardwareBaseboard,omitempty"`
	// 
	HardwareFirmware *string `json:"hardwareFirmware,omitempty"`
	IpfsInfo *XmlNs0IPFSSystemInfo `json:"ipfsInfo,omitempty"`
	// 
	Memory *string `json:"memory,omitempty"`
	// 
	OperatingSystem *string `json:"operatingSystem,omitempty"`
	// 
	OperatingSystemUptime *int64 `json:"operatingSystemUptime,omitempty"`
	// 
	Owner *string `json:"owner,omitempty"`
	// 
	Processors *[]string `json:"processors,omitempty"`
	// 
	PublicKey *string `json:"publicKey,omitempty"`
	SgxInfo *XmlNs0SGXInfo `json:"sgxInfo,omitempty"`
	// 
	Status *string `json:"status,omitempty"`
}

// NewXmlNs0NodeInfoAllOf instantiates a new XmlNs0NodeInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0NodeInfoAllOf() *XmlNs0NodeInfoAllOf {
	this := XmlNs0NodeInfoAllOf{}
	return &this
}

// NewXmlNs0NodeInfoAllOfWithDefaults instantiates a new XmlNs0NodeInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0NodeInfoAllOfWithDefaults() *XmlNs0NodeInfoAllOf {
	this := XmlNs0NodeInfoAllOf{}
	return &this
}

// GetConnectionURL returns the ConnectionURL field value if set, zero value otherwise.
func (o *XmlNs0NodeInfoAllOf) GetConnectionURL() string {
	if o == nil || o.ConnectionURL == nil {
		var ret string
		return ret
	}
	return *o.ConnectionURL
}

// GetConnectionURLOk returns a tuple with the ConnectionURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0NodeInfoAllOf) GetConnectionURLOk() (*string, bool) {
	if o == nil || o.ConnectionURL == nil {
		return nil, false
	}
	return o.ConnectionURL, true
}

// HasConnectionURL returns a boolean if a field has been set.
func (o *XmlNs0NodeInfoAllOf) HasConnectionURL() bool {
	if o != nil && o.ConnectionURL != nil {
		return true
	}

	return false
}

// SetConnectionURL gets a reference to the given string and assigns it to the ConnectionURL field.
func (o *XmlNs0NodeInfoAllOf) SetConnectionURL(v string) {
	o.ConnectionURL = &v
}

// GetCpuFamily returns the CpuFamily field value if set, zero value otherwise.
func (o *XmlNs0NodeInfoAllOf) GetCpuFamily() string {
	if o == nil || o.CpuFamily == nil {
		var ret string
		return ret
	}
	return *o.CpuFamily
}

// GetCpuFamilyOk returns a tuple with the CpuFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0NodeInfoAllOf) GetCpuFamilyOk() (*string, bool) {
	if o == nil || o.CpuFamily == nil {
		return nil, false
	}
	return o.CpuFamily, true
}

// HasCpuFamily returns a boolean if a field has been set.
func (o *XmlNs0NodeInfoAllOf) HasCpuFamily() bool {
	if o != nil && o.CpuFamily != nil {
		return true
	}

	return false
}

// SetCpuFamily gets a reference to the given string and assigns it to the CpuFamily field.
func (o *XmlNs0NodeInfoAllOf) SetCpuFamily(v string) {
	o.CpuFamily = &v
}

// GetCpuFrequency returns the CpuFrequency field value if set, zero value otherwise.
func (o *XmlNs0NodeInfoAllOf) GetCpuFrequency() string {
	if o == nil || o.CpuFrequency == nil {
		var ret string
		return ret
	}
	return *o.CpuFrequency
}

// GetCpuFrequencyOk returns a tuple with the CpuFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0NodeInfoAllOf) GetCpuFrequencyOk() (*string, bool) {
	if o == nil || o.CpuFrequency == nil {
		return nil, false
	}
	return o.CpuFrequency, true
}

// HasCpuFrequency returns a boolean if a field has been set.
func (o *XmlNs0NodeInfoAllOf) HasCpuFrequency() bool {
	if o != nil && o.CpuFrequency != nil {
		return true
	}

	return false
}

// SetCpuFrequency gets a reference to the given string and assigns it to the CpuFrequency field.
func (o *XmlNs0NodeInfoAllOf) SetCpuFrequency(v string) {
	o.CpuFrequency = &v
}

// GetCpuLogicalCount returns the CpuLogicalCount field value if set, zero value otherwise.
func (o *XmlNs0NodeInfoAllOf) GetCpuLogicalCount() int32 {
	if o == nil || o.CpuLogicalCount == nil {
		var ret int32
		return ret
	}
	return *o.CpuLogicalCount
}

// GetCpuLogicalCountOk returns a tuple with the CpuLogicalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0NodeInfoAllOf) GetCpuLogicalCountOk() (*int32, bool) {
	if o == nil || o.CpuLogicalCount == nil {
		return nil, false
	}
	return o.CpuLogicalCount, true
}

// HasCpuLogicalCount returns a boolean if a field has been set.
func (o *XmlNs0NodeInfoAllOf) HasCpuLogicalCount() bool {
	if o != nil && o.CpuLogicalCount != nil {
		return true
	}

	return false
}

// SetCpuLogicalCount gets a reference to the given int32 and assigns it to the CpuLogicalCount field.
func (o *XmlNs0NodeInfoAllOf) SetCpuLogicalCount(v int32) {
	o.CpuLogicalCount = &v
}

// GetCpuModel returns the CpuModel field value if set, zero value otherwise.
func (o *XmlNs0NodeInfoAllOf) GetCpuModel() string {
	if o == nil || o.CpuModel == nil {
		var ret string
		return ret
	}
	return *o.CpuModel
}

// GetCpuModelOk returns a tuple with the CpuModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0NodeInfoAllOf) GetCpuModelOk() (*string, bool) {
	if o == nil || o.CpuModel == nil {
		return nil, false
	}
	return o.CpuModel, true
}

// HasCpuModel returns a boolean if a field has been set.
func (o *XmlNs0NodeInfoAllOf) HasCpuModel() bool {
	if o != nil && o.CpuModel != nil {
		return true
	}

	return false
}

// SetCpuModel gets a reference to the given string and assigns it to the CpuModel field.
func (o *XmlNs0NodeInfoAllOf) SetCpuModel(v string) {
	o.CpuModel = &v
}

// GetCpuPhysicalCores returns the CpuPhysicalCores field value if set, zero value otherwise.
func (o *XmlNs0NodeInfoAllOf) GetCpuPhysicalCores() int32 {
	if o == nil || o.CpuPhysicalCores == nil {
		var ret int32
		return ret
	}
	return *o.CpuPhysicalCores
}

// GetCpuPhysicalCoresOk returns a tuple with the CpuPhysicalCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0NodeInfoAllOf) GetCpuPhysicalCoresOk() (*int32, bool) {
	if o == nil || o.CpuPhysicalCores == nil {
		return nil, false
	}
	return o.CpuPhysicalCores, true
}

// HasCpuPhysicalCores returns a boolean if a field has been set.
func (o *XmlNs0NodeInfoAllOf) HasCpuPhysicalCores() bool {
	if o != nil && o.CpuPhysicalCores != nil {
		return true
	}

	return false
}

// SetCpuPhysicalCores gets a reference to the given int32 and assigns it to the CpuPhysicalCores field.
func (o *XmlNs0NodeInfoAllOf) SetCpuPhysicalCores(v int32) {
	o.CpuPhysicalCores = &v
}

// GetHardwareBaseboard returns the HardwareBaseboard field value if set, zero value otherwise.
func (o *XmlNs0NodeInfoAllOf) GetHardwareBaseboard() string {
	if o == nil || o.HardwareBaseboard == nil {
		var ret string
		return ret
	}
	return *o.HardwareBaseboard
}

// GetHardwareBaseboardOk returns a tuple with the HardwareBaseboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0NodeInfoAllOf) GetHardwareBaseboardOk() (*string, bool) {
	if o == nil || o.HardwareBaseboard == nil {
		return nil, false
	}
	return o.HardwareBaseboard, true
}

// HasHardwareBaseboard returns a boolean if a field has been set.
func (o *XmlNs0NodeInfoAllOf) HasHardwareBaseboard() bool {
	if o != nil && o.HardwareBaseboard != nil {
		return true
	}

	return false
}

// SetHardwareBaseboard gets a reference to the given string and assigns it to the HardwareBaseboard field.
func (o *XmlNs0NodeInfoAllOf) SetHardwareBaseboard(v string) {
	o.HardwareBaseboard = &v
}

// GetHardwareFirmware returns the HardwareFirmware field value if set, zero value otherwise.
func (o *XmlNs0NodeInfoAllOf) GetHardwareFirmware() string {
	if o == nil || o.HardwareFirmware == nil {
		var ret string
		return ret
	}
	return *o.HardwareFirmware
}

// GetHardwareFirmwareOk returns a tuple with the HardwareFirmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0NodeInfoAllOf) GetHardwareFirmwareOk() (*string, bool) {
	if o == nil || o.HardwareFirmware == nil {
		return nil, false
	}
	return o.HardwareFirmware, true
}

// HasHardwareFirmware returns a boolean if a field has been set.
func (o *XmlNs0NodeInfoAllOf) HasHardwareFirmware() bool {
	if o != nil && o.HardwareFirmware != nil {
		return true
	}

	return false
}

// SetHardwareFirmware gets a reference to the given string and assigns it to the HardwareFirmware field.
func (o *XmlNs0NodeInfoAllOf) SetHardwareFirmware(v string) {
	o.HardwareFirmware = &v
}

// GetIpfsInfo returns the IpfsInfo field value if set, zero value otherwise.
func (o *XmlNs0NodeInfoAllOf) GetIpfsInfo() XmlNs0IPFSSystemInfo {
	if o == nil || o.IpfsInfo == nil {
		var ret XmlNs0IPFSSystemInfo
		return ret
	}
	return *o.IpfsInfo
}

// GetIpfsInfoOk returns a tuple with the IpfsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0NodeInfoAllOf) GetIpfsInfoOk() (*XmlNs0IPFSSystemInfo, bool) {
	if o == nil || o.IpfsInfo == nil {
		return nil, false
	}
	return o.IpfsInfo, true
}

// HasIpfsInfo returns a boolean if a field has been set.
func (o *XmlNs0NodeInfoAllOf) HasIpfsInfo() bool {
	if o != nil && o.IpfsInfo != nil {
		return true
	}

	return false
}

// SetIpfsInfo gets a reference to the given XmlNs0IPFSSystemInfo and assigns it to the IpfsInfo field.
func (o *XmlNs0NodeInfoAllOf) SetIpfsInfo(v XmlNs0IPFSSystemInfo) {
	o.IpfsInfo = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *XmlNs0NodeInfoAllOf) GetMemory() string {
	if o == nil || o.Memory == nil {
		var ret string
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0NodeInfoAllOf) GetMemoryOk() (*string, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *XmlNs0NodeInfoAllOf) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given string and assigns it to the Memory field.
func (o *XmlNs0NodeInfoAllOf) SetMemory(v string) {
	o.Memory = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *XmlNs0NodeInfoAllOf) GetOperatingSystem() string {
	if o == nil || o.OperatingSystem == nil {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0NodeInfoAllOf) GetOperatingSystemOk() (*string, bool) {
	if o == nil || o.OperatingSystem == nil {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *XmlNs0NodeInfoAllOf) HasOperatingSystem() bool {
	if o != nil && o.OperatingSystem != nil {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *XmlNs0NodeInfoAllOf) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

// GetOperatingSystemUptime returns the OperatingSystemUptime field value if set, zero value otherwise.
func (o *XmlNs0NodeInfoAllOf) GetOperatingSystemUptime() int64 {
	if o == nil || o.OperatingSystemUptime == nil {
		var ret int64
		return ret
	}
	return *o.OperatingSystemUptime
}

// GetOperatingSystemUptimeOk returns a tuple with the OperatingSystemUptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0NodeInfoAllOf) GetOperatingSystemUptimeOk() (*int64, bool) {
	if o == nil || o.OperatingSystemUptime == nil {
		return nil, false
	}
	return o.OperatingSystemUptime, true
}

// HasOperatingSystemUptime returns a boolean if a field has been set.
func (o *XmlNs0NodeInfoAllOf) HasOperatingSystemUptime() bool {
	if o != nil && o.OperatingSystemUptime != nil {
		return true
	}

	return false
}

// SetOperatingSystemUptime gets a reference to the given int64 and assigns it to the OperatingSystemUptime field.
func (o *XmlNs0NodeInfoAllOf) SetOperatingSystemUptime(v int64) {
	o.OperatingSystemUptime = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *XmlNs0NodeInfoAllOf) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0NodeInfoAllOf) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *XmlNs0NodeInfoAllOf) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *XmlNs0NodeInfoAllOf) SetOwner(v string) {
	o.Owner = &v
}

// GetProcessors returns the Processors field value if set, zero value otherwise.
func (o *XmlNs0NodeInfoAllOf) GetProcessors() []string {
	if o == nil || o.Processors == nil {
		var ret []string
		return ret
	}
	return *o.Processors
}

// GetProcessorsOk returns a tuple with the Processors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0NodeInfoAllOf) GetProcessorsOk() (*[]string, bool) {
	if o == nil || o.Processors == nil {
		return nil, false
	}
	return o.Processors, true
}

// HasProcessors returns a boolean if a field has been set.
func (o *XmlNs0NodeInfoAllOf) HasProcessors() bool {
	if o != nil && o.Processors != nil {
		return true
	}

	return false
}

// SetProcessors gets a reference to the given []string and assigns it to the Processors field.
func (o *XmlNs0NodeInfoAllOf) SetProcessors(v []string) {
	o.Processors = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *XmlNs0NodeInfoAllOf) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0NodeInfoAllOf) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *XmlNs0NodeInfoAllOf) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *XmlNs0NodeInfoAllOf) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetSgxInfo returns the SgxInfo field value if set, zero value otherwise.
func (o *XmlNs0NodeInfoAllOf) GetSgxInfo() XmlNs0SGXInfo {
	if o == nil || o.SgxInfo == nil {
		var ret XmlNs0SGXInfo
		return ret
	}
	return *o.SgxInfo
}

// GetSgxInfoOk returns a tuple with the SgxInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0NodeInfoAllOf) GetSgxInfoOk() (*XmlNs0SGXInfo, bool) {
	if o == nil || o.SgxInfo == nil {
		return nil, false
	}
	return o.SgxInfo, true
}

// HasSgxInfo returns a boolean if a field has been set.
func (o *XmlNs0NodeInfoAllOf) HasSgxInfo() bool {
	if o != nil && o.SgxInfo != nil {
		return true
	}

	return false
}

// SetSgxInfo gets a reference to the given XmlNs0SGXInfo and assigns it to the SgxInfo field.
func (o *XmlNs0NodeInfoAllOf) SetSgxInfo(v XmlNs0SGXInfo) {
	o.SgxInfo = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *XmlNs0NodeInfoAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0NodeInfoAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *XmlNs0NodeInfoAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *XmlNs0NodeInfoAllOf) SetStatus(v string) {
	o.Status = &v
}

func (o XmlNs0NodeInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectionURL != nil {
		toSerialize["connectionURL"] = o.ConnectionURL
	}
	if o.CpuFamily != nil {
		toSerialize["cpuFamily"] = o.CpuFamily
	}
	if o.CpuFrequency != nil {
		toSerialize["cpuFrequency"] = o.CpuFrequency
	}
	if o.CpuLogicalCount != nil {
		toSerialize["cpuLogicalCount"] = o.CpuLogicalCount
	}
	if o.CpuModel != nil {
		toSerialize["cpuModel"] = o.CpuModel
	}
	if o.CpuPhysicalCores != nil {
		toSerialize["cpuPhysicalCores"] = o.CpuPhysicalCores
	}
	if o.HardwareBaseboard != nil {
		toSerialize["hardwareBaseboard"] = o.HardwareBaseboard
	}
	if o.HardwareFirmware != nil {
		toSerialize["hardwareFirmware"] = o.HardwareFirmware
	}
	if o.IpfsInfo != nil {
		toSerialize["ipfsInfo"] = o.IpfsInfo
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.OperatingSystem != nil {
		toSerialize["operatingSystem"] = o.OperatingSystem
	}
	if o.OperatingSystemUptime != nil {
		toSerialize["operatingSystemUptime"] = o.OperatingSystemUptime
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Processors != nil {
		toSerialize["processors"] = o.Processors
	}
	if o.PublicKey != nil {
		toSerialize["publicKey"] = o.PublicKey
	}
	if o.SgxInfo != nil {
		toSerialize["sgxInfo"] = o.SgxInfo
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0NodeInfoAllOf struct {
	value *XmlNs0NodeInfoAllOf
	isSet bool
}

func (v NullableXmlNs0NodeInfoAllOf) Get() *XmlNs0NodeInfoAllOf {
	return v.value
}

func (v *NullableXmlNs0NodeInfoAllOf) Set(val *XmlNs0NodeInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0NodeInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0NodeInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0NodeInfoAllOf(val *XmlNs0NodeInfoAllOf) *NullableXmlNs0NodeInfoAllOf {
	return &NullableXmlNs0NodeInfoAllOf{value: val, isSet: true}
}

func (v NullableXmlNs0NodeInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0NodeInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


