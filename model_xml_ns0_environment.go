/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.46
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0Environment 
type XmlNs0Environment struct {
	// 
	Content *[]string `json:"content,omitempty"`
	DefaultRunConfiguration *XmlNs0RunConfig `json:"defaultRunConfiguration,omitempty"`
	// 
	Description *string `json:"description,omitempty"`
	// 
	IpfsHash *string `json:"ipfsHash,omitempty"`
	// 
	Name *string `json:"name,omitempty"`
	// 
	Packages *[]string `json:"packages,omitempty"`
	// 
	Published *bool `json:"published,omitempty"`
	// 
	RootHashOffset *string `json:"rootHashOffset,omitempty"`
	// 
	Roothash *string `json:"roothash,omitempty"`
	// 
	Size *string `json:"size,omitempty"`
	// 
	Uuid *string `json:"uuid,omitempty"`
}

// NewXmlNs0Environment instantiates a new XmlNs0Environment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0Environment() *XmlNs0Environment {
	this := XmlNs0Environment{}
	return &this
}

// NewXmlNs0EnvironmentWithDefaults instantiates a new XmlNs0Environment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0EnvironmentWithDefaults() *XmlNs0Environment {
	this := XmlNs0Environment{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *XmlNs0Environment) GetContent() []string {
	if o == nil || o.Content == nil {
		var ret []string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0Environment) GetContentOk() (*[]string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *XmlNs0Environment) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []string and assigns it to the Content field.
func (o *XmlNs0Environment) SetContent(v []string) {
	o.Content = &v
}

// GetDefaultRunConfiguration returns the DefaultRunConfiguration field value if set, zero value otherwise.
func (o *XmlNs0Environment) GetDefaultRunConfiguration() XmlNs0RunConfig {
	if o == nil || o.DefaultRunConfiguration == nil {
		var ret XmlNs0RunConfig
		return ret
	}
	return *o.DefaultRunConfiguration
}

// GetDefaultRunConfigurationOk returns a tuple with the DefaultRunConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0Environment) GetDefaultRunConfigurationOk() (*XmlNs0RunConfig, bool) {
	if o == nil || o.DefaultRunConfiguration == nil {
		return nil, false
	}
	return o.DefaultRunConfiguration, true
}

// HasDefaultRunConfiguration returns a boolean if a field has been set.
func (o *XmlNs0Environment) HasDefaultRunConfiguration() bool {
	if o != nil && o.DefaultRunConfiguration != nil {
		return true
	}

	return false
}

// SetDefaultRunConfiguration gets a reference to the given XmlNs0RunConfig and assigns it to the DefaultRunConfiguration field.
func (o *XmlNs0Environment) SetDefaultRunConfiguration(v XmlNs0RunConfig) {
	o.DefaultRunConfiguration = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *XmlNs0Environment) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0Environment) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *XmlNs0Environment) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *XmlNs0Environment) SetDescription(v string) {
	o.Description = &v
}

// GetIpfsHash returns the IpfsHash field value if set, zero value otherwise.
func (o *XmlNs0Environment) GetIpfsHash() string {
	if o == nil || o.IpfsHash == nil {
		var ret string
		return ret
	}
	return *o.IpfsHash
}

// GetIpfsHashOk returns a tuple with the IpfsHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0Environment) GetIpfsHashOk() (*string, bool) {
	if o == nil || o.IpfsHash == nil {
		return nil, false
	}
	return o.IpfsHash, true
}

// HasIpfsHash returns a boolean if a field has been set.
func (o *XmlNs0Environment) HasIpfsHash() bool {
	if o != nil && o.IpfsHash != nil {
		return true
	}

	return false
}

// SetIpfsHash gets a reference to the given string and assigns it to the IpfsHash field.
func (o *XmlNs0Environment) SetIpfsHash(v string) {
	o.IpfsHash = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *XmlNs0Environment) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0Environment) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *XmlNs0Environment) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *XmlNs0Environment) SetName(v string) {
	o.Name = &v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *XmlNs0Environment) GetPackages() []string {
	if o == nil || o.Packages == nil {
		var ret []string
		return ret
	}
	return *o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0Environment) GetPackagesOk() (*[]string, bool) {
	if o == nil || o.Packages == nil {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *XmlNs0Environment) HasPackages() bool {
	if o != nil && o.Packages != nil {
		return true
	}

	return false
}

// SetPackages gets a reference to the given []string and assigns it to the Packages field.
func (o *XmlNs0Environment) SetPackages(v []string) {
	o.Packages = &v
}

// GetPublished returns the Published field value if set, zero value otherwise.
func (o *XmlNs0Environment) GetPublished() bool {
	if o == nil || o.Published == nil {
		var ret bool
		return ret
	}
	return *o.Published
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0Environment) GetPublishedOk() (*bool, bool) {
	if o == nil || o.Published == nil {
		return nil, false
	}
	return o.Published, true
}

// HasPublished returns a boolean if a field has been set.
func (o *XmlNs0Environment) HasPublished() bool {
	if o != nil && o.Published != nil {
		return true
	}

	return false
}

// SetPublished gets a reference to the given bool and assigns it to the Published field.
func (o *XmlNs0Environment) SetPublished(v bool) {
	o.Published = &v
}

// GetRootHashOffset returns the RootHashOffset field value if set, zero value otherwise.
func (o *XmlNs0Environment) GetRootHashOffset() string {
	if o == nil || o.RootHashOffset == nil {
		var ret string
		return ret
	}
	return *o.RootHashOffset
}

// GetRootHashOffsetOk returns a tuple with the RootHashOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0Environment) GetRootHashOffsetOk() (*string, bool) {
	if o == nil || o.RootHashOffset == nil {
		return nil, false
	}
	return o.RootHashOffset, true
}

// HasRootHashOffset returns a boolean if a field has been set.
func (o *XmlNs0Environment) HasRootHashOffset() bool {
	if o != nil && o.RootHashOffset != nil {
		return true
	}

	return false
}

// SetRootHashOffset gets a reference to the given string and assigns it to the RootHashOffset field.
func (o *XmlNs0Environment) SetRootHashOffset(v string) {
	o.RootHashOffset = &v
}

// GetRoothash returns the Roothash field value if set, zero value otherwise.
func (o *XmlNs0Environment) GetRoothash() string {
	if o == nil || o.Roothash == nil {
		var ret string
		return ret
	}
	return *o.Roothash
}

// GetRoothashOk returns a tuple with the Roothash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0Environment) GetRoothashOk() (*string, bool) {
	if o == nil || o.Roothash == nil {
		return nil, false
	}
	return o.Roothash, true
}

// HasRoothash returns a boolean if a field has been set.
func (o *XmlNs0Environment) HasRoothash() bool {
	if o != nil && o.Roothash != nil {
		return true
	}

	return false
}

// SetRoothash gets a reference to the given string and assigns it to the Roothash field.
func (o *XmlNs0Environment) SetRoothash(v string) {
	o.Roothash = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *XmlNs0Environment) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0Environment) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *XmlNs0Environment) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *XmlNs0Environment) SetSize(v string) {
	o.Size = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *XmlNs0Environment) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0Environment) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *XmlNs0Environment) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *XmlNs0Environment) SetUuid(v string) {
	o.Uuid = &v
}

func (o XmlNs0Environment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.DefaultRunConfiguration != nil {
		toSerialize["defaultRunConfiguration"] = o.DefaultRunConfiguration
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.IpfsHash != nil {
		toSerialize["ipfsHash"] = o.IpfsHash
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Packages != nil {
		toSerialize["packages"] = o.Packages
	}
	if o.Published != nil {
		toSerialize["published"] = o.Published
	}
	if o.RootHashOffset != nil {
		toSerialize["rootHashOffset"] = o.RootHashOffset
	}
	if o.Roothash != nil {
		toSerialize["roothash"] = o.Roothash
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0Environment struct {
	value *XmlNs0Environment
	isSet bool
}

func (v NullableXmlNs0Environment) Get() *XmlNs0Environment {
	return v.value
}

func (v *NullableXmlNs0Environment) Set(val *XmlNs0Environment) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0Environment) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0Environment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0Environment(val *XmlNs0Environment) *NullableXmlNs0Environment {
	return &NullableXmlNs0Environment{value: val, isSet: true}
}

func (v NullableXmlNs0Environment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0Environment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


