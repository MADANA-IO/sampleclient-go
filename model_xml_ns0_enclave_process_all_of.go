/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.30
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// XmlNs0EnclaveProcessAllOf struct for XmlNs0EnclaveProcessAllOf
type XmlNs0EnclaveProcessAllOf struct {
	// 
	AttestationServer *string `json:"attestationServer,omitempty"`
	// 
	ConsoleOutput *string `json:"consoleOutput,omitempty"`
	// 
	EnclaveIdent *string `json:"enclaveIdent,omitempty"`
	EnclaveInputstream *XmlNs0InputStream `json:"enclaveInputstream,omitempty"`
	// 
	EndingTime *string `json:"endingTime,omitempty"`
	Environment *XmlNs0Environment `json:"environment,omitempty"`
	// 
	InternalAttesationServer *string `json:"internalAttesationServer,omitempty"`
	// 
	InternalIdent *string `json:"internalIdent,omitempty"`
	// 
	InternalRemoteControlServer *string `json:"internalRemoteControlServer,omitempty"`
	// 
	InternalWireguardServer *string `json:"internalWireguardServer,omitempty"`
	KubernetesEnclave *XmlNs0KubernetesEnclave `json:"kubernetesEnclave,omitempty"`
	// 
	PortMapping *map[string]interface{} `json:"portMapping,omitempty"`
	// 
	Ports *[]XmlNs0EnclavePort `json:"ports,omitempty"`
	Process *XmlNs0Process `json:"process,omitempty"`
	// 
	PublicIdent *string `json:"publicIdent,omitempty"`
	// 
	RemoteControlServer *string `json:"remoteControlServer,omitempty"`
	// 
	SignerIdent *string `json:"signerIdent,omitempty"`
	// 
	StartupCMD *string `json:"startupCMD,omitempty"`
	// 
	StartupTime *string `json:"startupTime,omitempty"`
	// 
	Status *string `json:"status,omitempty"`
	WgInterface *XmlNs0WireguardInterface `json:"wgInterface,omitempty"`
	// 
	WireguardPublicKey *string `json:"wireguardPublicKey,omitempty"`
	// 
	WireguardServer *string `json:"wireguardServer,omitempty"`
}

// NewXmlNs0EnclaveProcessAllOf instantiates a new XmlNs0EnclaveProcessAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXmlNs0EnclaveProcessAllOf() *XmlNs0EnclaveProcessAllOf {
	this := XmlNs0EnclaveProcessAllOf{}
	return &this
}

// NewXmlNs0EnclaveProcessAllOfWithDefaults instantiates a new XmlNs0EnclaveProcessAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXmlNs0EnclaveProcessAllOfWithDefaults() *XmlNs0EnclaveProcessAllOf {
	this := XmlNs0EnclaveProcessAllOf{}
	return &this
}

// GetAttestationServer returns the AttestationServer field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetAttestationServer() string {
	if o == nil || o.AttestationServer == nil {
		var ret string
		return ret
	}
	return *o.AttestationServer
}

// GetAttestationServerOk returns a tuple with the AttestationServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetAttestationServerOk() (*string, bool) {
	if o == nil || o.AttestationServer == nil {
		return nil, false
	}
	return o.AttestationServer, true
}

// HasAttestationServer returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasAttestationServer() bool {
	if o != nil && o.AttestationServer != nil {
		return true
	}

	return false
}

// SetAttestationServer gets a reference to the given string and assigns it to the AttestationServer field.
func (o *XmlNs0EnclaveProcessAllOf) SetAttestationServer(v string) {
	o.AttestationServer = &v
}

// GetConsoleOutput returns the ConsoleOutput field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetConsoleOutput() string {
	if o == nil || o.ConsoleOutput == nil {
		var ret string
		return ret
	}
	return *o.ConsoleOutput
}

// GetConsoleOutputOk returns a tuple with the ConsoleOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetConsoleOutputOk() (*string, bool) {
	if o == nil || o.ConsoleOutput == nil {
		return nil, false
	}
	return o.ConsoleOutput, true
}

// HasConsoleOutput returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasConsoleOutput() bool {
	if o != nil && o.ConsoleOutput != nil {
		return true
	}

	return false
}

// SetConsoleOutput gets a reference to the given string and assigns it to the ConsoleOutput field.
func (o *XmlNs0EnclaveProcessAllOf) SetConsoleOutput(v string) {
	o.ConsoleOutput = &v
}

// GetEnclaveIdent returns the EnclaveIdent field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetEnclaveIdent() string {
	if o == nil || o.EnclaveIdent == nil {
		var ret string
		return ret
	}
	return *o.EnclaveIdent
}

// GetEnclaveIdentOk returns a tuple with the EnclaveIdent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetEnclaveIdentOk() (*string, bool) {
	if o == nil || o.EnclaveIdent == nil {
		return nil, false
	}
	return o.EnclaveIdent, true
}

// HasEnclaveIdent returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasEnclaveIdent() bool {
	if o != nil && o.EnclaveIdent != nil {
		return true
	}

	return false
}

// SetEnclaveIdent gets a reference to the given string and assigns it to the EnclaveIdent field.
func (o *XmlNs0EnclaveProcessAllOf) SetEnclaveIdent(v string) {
	o.EnclaveIdent = &v
}

// GetEnclaveInputstream returns the EnclaveInputstream field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetEnclaveInputstream() XmlNs0InputStream {
	if o == nil || o.EnclaveInputstream == nil {
		var ret XmlNs0InputStream
		return ret
	}
	return *o.EnclaveInputstream
}

// GetEnclaveInputstreamOk returns a tuple with the EnclaveInputstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetEnclaveInputstreamOk() (*XmlNs0InputStream, bool) {
	if o == nil || o.EnclaveInputstream == nil {
		return nil, false
	}
	return o.EnclaveInputstream, true
}

// HasEnclaveInputstream returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasEnclaveInputstream() bool {
	if o != nil && o.EnclaveInputstream != nil {
		return true
	}

	return false
}

// SetEnclaveInputstream gets a reference to the given XmlNs0InputStream and assigns it to the EnclaveInputstream field.
func (o *XmlNs0EnclaveProcessAllOf) SetEnclaveInputstream(v XmlNs0InputStream) {
	o.EnclaveInputstream = &v
}

// GetEndingTime returns the EndingTime field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetEndingTime() string {
	if o == nil || o.EndingTime == nil {
		var ret string
		return ret
	}
	return *o.EndingTime
}

// GetEndingTimeOk returns a tuple with the EndingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetEndingTimeOk() (*string, bool) {
	if o == nil || o.EndingTime == nil {
		return nil, false
	}
	return o.EndingTime, true
}

// HasEndingTime returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasEndingTime() bool {
	if o != nil && o.EndingTime != nil {
		return true
	}

	return false
}

// SetEndingTime gets a reference to the given string and assigns it to the EndingTime field.
func (o *XmlNs0EnclaveProcessAllOf) SetEndingTime(v string) {
	o.EndingTime = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetEnvironment() XmlNs0Environment {
	if o == nil || o.Environment == nil {
		var ret XmlNs0Environment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetEnvironmentOk() (*XmlNs0Environment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given XmlNs0Environment and assigns it to the Environment field.
func (o *XmlNs0EnclaveProcessAllOf) SetEnvironment(v XmlNs0Environment) {
	o.Environment = &v
}

// GetInternalAttesationServer returns the InternalAttesationServer field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetInternalAttesationServer() string {
	if o == nil || o.InternalAttesationServer == nil {
		var ret string
		return ret
	}
	return *o.InternalAttesationServer
}

// GetInternalAttesationServerOk returns a tuple with the InternalAttesationServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetInternalAttesationServerOk() (*string, bool) {
	if o == nil || o.InternalAttesationServer == nil {
		return nil, false
	}
	return o.InternalAttesationServer, true
}

// HasInternalAttesationServer returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasInternalAttesationServer() bool {
	if o != nil && o.InternalAttesationServer != nil {
		return true
	}

	return false
}

// SetInternalAttesationServer gets a reference to the given string and assigns it to the InternalAttesationServer field.
func (o *XmlNs0EnclaveProcessAllOf) SetInternalAttesationServer(v string) {
	o.InternalAttesationServer = &v
}

// GetInternalIdent returns the InternalIdent field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetInternalIdent() string {
	if o == nil || o.InternalIdent == nil {
		var ret string
		return ret
	}
	return *o.InternalIdent
}

// GetInternalIdentOk returns a tuple with the InternalIdent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetInternalIdentOk() (*string, bool) {
	if o == nil || o.InternalIdent == nil {
		return nil, false
	}
	return o.InternalIdent, true
}

// HasInternalIdent returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasInternalIdent() bool {
	if o != nil && o.InternalIdent != nil {
		return true
	}

	return false
}

// SetInternalIdent gets a reference to the given string and assigns it to the InternalIdent field.
func (o *XmlNs0EnclaveProcessAllOf) SetInternalIdent(v string) {
	o.InternalIdent = &v
}

// GetInternalRemoteControlServer returns the InternalRemoteControlServer field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetInternalRemoteControlServer() string {
	if o == nil || o.InternalRemoteControlServer == nil {
		var ret string
		return ret
	}
	return *o.InternalRemoteControlServer
}

// GetInternalRemoteControlServerOk returns a tuple with the InternalRemoteControlServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetInternalRemoteControlServerOk() (*string, bool) {
	if o == nil || o.InternalRemoteControlServer == nil {
		return nil, false
	}
	return o.InternalRemoteControlServer, true
}

// HasInternalRemoteControlServer returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasInternalRemoteControlServer() bool {
	if o != nil && o.InternalRemoteControlServer != nil {
		return true
	}

	return false
}

// SetInternalRemoteControlServer gets a reference to the given string and assigns it to the InternalRemoteControlServer field.
func (o *XmlNs0EnclaveProcessAllOf) SetInternalRemoteControlServer(v string) {
	o.InternalRemoteControlServer = &v
}

// GetInternalWireguardServer returns the InternalWireguardServer field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetInternalWireguardServer() string {
	if o == nil || o.InternalWireguardServer == nil {
		var ret string
		return ret
	}
	return *o.InternalWireguardServer
}

// GetInternalWireguardServerOk returns a tuple with the InternalWireguardServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetInternalWireguardServerOk() (*string, bool) {
	if o == nil || o.InternalWireguardServer == nil {
		return nil, false
	}
	return o.InternalWireguardServer, true
}

// HasInternalWireguardServer returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasInternalWireguardServer() bool {
	if o != nil && o.InternalWireguardServer != nil {
		return true
	}

	return false
}

// SetInternalWireguardServer gets a reference to the given string and assigns it to the InternalWireguardServer field.
func (o *XmlNs0EnclaveProcessAllOf) SetInternalWireguardServer(v string) {
	o.InternalWireguardServer = &v
}

// GetKubernetesEnclave returns the KubernetesEnclave field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetKubernetesEnclave() XmlNs0KubernetesEnclave {
	if o == nil || o.KubernetesEnclave == nil {
		var ret XmlNs0KubernetesEnclave
		return ret
	}
	return *o.KubernetesEnclave
}

// GetKubernetesEnclaveOk returns a tuple with the KubernetesEnclave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetKubernetesEnclaveOk() (*XmlNs0KubernetesEnclave, bool) {
	if o == nil || o.KubernetesEnclave == nil {
		return nil, false
	}
	return o.KubernetesEnclave, true
}

// HasKubernetesEnclave returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasKubernetesEnclave() bool {
	if o != nil && o.KubernetesEnclave != nil {
		return true
	}

	return false
}

// SetKubernetesEnclave gets a reference to the given XmlNs0KubernetesEnclave and assigns it to the KubernetesEnclave field.
func (o *XmlNs0EnclaveProcessAllOf) SetKubernetesEnclave(v XmlNs0KubernetesEnclave) {
	o.KubernetesEnclave = &v
}

// GetPortMapping returns the PortMapping field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetPortMapping() map[string]interface{} {
	if o == nil || o.PortMapping == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.PortMapping
}

// GetPortMappingOk returns a tuple with the PortMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetPortMappingOk() (*map[string]interface{}, bool) {
	if o == nil || o.PortMapping == nil {
		return nil, false
	}
	return o.PortMapping, true
}

// HasPortMapping returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasPortMapping() bool {
	if o != nil && o.PortMapping != nil {
		return true
	}

	return false
}

// SetPortMapping gets a reference to the given map[string]interface{} and assigns it to the PortMapping field.
func (o *XmlNs0EnclaveProcessAllOf) SetPortMapping(v map[string]interface{}) {
	o.PortMapping = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetPorts() []XmlNs0EnclavePort {
	if o == nil || o.Ports == nil {
		var ret []XmlNs0EnclavePort
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetPortsOk() (*[]XmlNs0EnclavePort, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []XmlNs0EnclavePort and assigns it to the Ports field.
func (o *XmlNs0EnclaveProcessAllOf) SetPorts(v []XmlNs0EnclavePort) {
	o.Ports = &v
}

// GetProcess returns the Process field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetProcess() XmlNs0Process {
	if o == nil || o.Process == nil {
		var ret XmlNs0Process
		return ret
	}
	return *o.Process
}

// GetProcessOk returns a tuple with the Process field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetProcessOk() (*XmlNs0Process, bool) {
	if o == nil || o.Process == nil {
		return nil, false
	}
	return o.Process, true
}

// HasProcess returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasProcess() bool {
	if o != nil && o.Process != nil {
		return true
	}

	return false
}

// SetProcess gets a reference to the given XmlNs0Process and assigns it to the Process field.
func (o *XmlNs0EnclaveProcessAllOf) SetProcess(v XmlNs0Process) {
	o.Process = &v
}

// GetPublicIdent returns the PublicIdent field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetPublicIdent() string {
	if o == nil || o.PublicIdent == nil {
		var ret string
		return ret
	}
	return *o.PublicIdent
}

// GetPublicIdentOk returns a tuple with the PublicIdent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetPublicIdentOk() (*string, bool) {
	if o == nil || o.PublicIdent == nil {
		return nil, false
	}
	return o.PublicIdent, true
}

// HasPublicIdent returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasPublicIdent() bool {
	if o != nil && o.PublicIdent != nil {
		return true
	}

	return false
}

// SetPublicIdent gets a reference to the given string and assigns it to the PublicIdent field.
func (o *XmlNs0EnclaveProcessAllOf) SetPublicIdent(v string) {
	o.PublicIdent = &v
}

// GetRemoteControlServer returns the RemoteControlServer field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetRemoteControlServer() string {
	if o == nil || o.RemoteControlServer == nil {
		var ret string
		return ret
	}
	return *o.RemoteControlServer
}

// GetRemoteControlServerOk returns a tuple with the RemoteControlServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetRemoteControlServerOk() (*string, bool) {
	if o == nil || o.RemoteControlServer == nil {
		return nil, false
	}
	return o.RemoteControlServer, true
}

// HasRemoteControlServer returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasRemoteControlServer() bool {
	if o != nil && o.RemoteControlServer != nil {
		return true
	}

	return false
}

// SetRemoteControlServer gets a reference to the given string and assigns it to the RemoteControlServer field.
func (o *XmlNs0EnclaveProcessAllOf) SetRemoteControlServer(v string) {
	o.RemoteControlServer = &v
}

// GetSignerIdent returns the SignerIdent field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetSignerIdent() string {
	if o == nil || o.SignerIdent == nil {
		var ret string
		return ret
	}
	return *o.SignerIdent
}

// GetSignerIdentOk returns a tuple with the SignerIdent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetSignerIdentOk() (*string, bool) {
	if o == nil || o.SignerIdent == nil {
		return nil, false
	}
	return o.SignerIdent, true
}

// HasSignerIdent returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasSignerIdent() bool {
	if o != nil && o.SignerIdent != nil {
		return true
	}

	return false
}

// SetSignerIdent gets a reference to the given string and assigns it to the SignerIdent field.
func (o *XmlNs0EnclaveProcessAllOf) SetSignerIdent(v string) {
	o.SignerIdent = &v
}

// GetStartupCMD returns the StartupCMD field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetStartupCMD() string {
	if o == nil || o.StartupCMD == nil {
		var ret string
		return ret
	}
	return *o.StartupCMD
}

// GetStartupCMDOk returns a tuple with the StartupCMD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetStartupCMDOk() (*string, bool) {
	if o == nil || o.StartupCMD == nil {
		return nil, false
	}
	return o.StartupCMD, true
}

// HasStartupCMD returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasStartupCMD() bool {
	if o != nil && o.StartupCMD != nil {
		return true
	}

	return false
}

// SetStartupCMD gets a reference to the given string and assigns it to the StartupCMD field.
func (o *XmlNs0EnclaveProcessAllOf) SetStartupCMD(v string) {
	o.StartupCMD = &v
}

// GetStartupTime returns the StartupTime field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetStartupTime() string {
	if o == nil || o.StartupTime == nil {
		var ret string
		return ret
	}
	return *o.StartupTime
}

// GetStartupTimeOk returns a tuple with the StartupTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetStartupTimeOk() (*string, bool) {
	if o == nil || o.StartupTime == nil {
		return nil, false
	}
	return o.StartupTime, true
}

// HasStartupTime returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasStartupTime() bool {
	if o != nil && o.StartupTime != nil {
		return true
	}

	return false
}

// SetStartupTime gets a reference to the given string and assigns it to the StartupTime field.
func (o *XmlNs0EnclaveProcessAllOf) SetStartupTime(v string) {
	o.StartupTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *XmlNs0EnclaveProcessAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetWgInterface returns the WgInterface field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetWgInterface() XmlNs0WireguardInterface {
	if o == nil || o.WgInterface == nil {
		var ret XmlNs0WireguardInterface
		return ret
	}
	return *o.WgInterface
}

// GetWgInterfaceOk returns a tuple with the WgInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetWgInterfaceOk() (*XmlNs0WireguardInterface, bool) {
	if o == nil || o.WgInterface == nil {
		return nil, false
	}
	return o.WgInterface, true
}

// HasWgInterface returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasWgInterface() bool {
	if o != nil && o.WgInterface != nil {
		return true
	}

	return false
}

// SetWgInterface gets a reference to the given XmlNs0WireguardInterface and assigns it to the WgInterface field.
func (o *XmlNs0EnclaveProcessAllOf) SetWgInterface(v XmlNs0WireguardInterface) {
	o.WgInterface = &v
}

// GetWireguardPublicKey returns the WireguardPublicKey field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetWireguardPublicKey() string {
	if o == nil || o.WireguardPublicKey == nil {
		var ret string
		return ret
	}
	return *o.WireguardPublicKey
}

// GetWireguardPublicKeyOk returns a tuple with the WireguardPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetWireguardPublicKeyOk() (*string, bool) {
	if o == nil || o.WireguardPublicKey == nil {
		return nil, false
	}
	return o.WireguardPublicKey, true
}

// HasWireguardPublicKey returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasWireguardPublicKey() bool {
	if o != nil && o.WireguardPublicKey != nil {
		return true
	}

	return false
}

// SetWireguardPublicKey gets a reference to the given string and assigns it to the WireguardPublicKey field.
func (o *XmlNs0EnclaveProcessAllOf) SetWireguardPublicKey(v string) {
	o.WireguardPublicKey = &v
}

// GetWireguardServer returns the WireguardServer field value if set, zero value otherwise.
func (o *XmlNs0EnclaveProcessAllOf) GetWireguardServer() string {
	if o == nil || o.WireguardServer == nil {
		var ret string
		return ret
	}
	return *o.WireguardServer
}

// GetWireguardServerOk returns a tuple with the WireguardServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XmlNs0EnclaveProcessAllOf) GetWireguardServerOk() (*string, bool) {
	if o == nil || o.WireguardServer == nil {
		return nil, false
	}
	return o.WireguardServer, true
}

// HasWireguardServer returns a boolean if a field has been set.
func (o *XmlNs0EnclaveProcessAllOf) HasWireguardServer() bool {
	if o != nil && o.WireguardServer != nil {
		return true
	}

	return false
}

// SetWireguardServer gets a reference to the given string and assigns it to the WireguardServer field.
func (o *XmlNs0EnclaveProcessAllOf) SetWireguardServer(v string) {
	o.WireguardServer = &v
}

func (o XmlNs0EnclaveProcessAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttestationServer != nil {
		toSerialize["attestationServer"] = o.AttestationServer
	}
	if o.ConsoleOutput != nil {
		toSerialize["consoleOutput"] = o.ConsoleOutput
	}
	if o.EnclaveIdent != nil {
		toSerialize["enclaveIdent"] = o.EnclaveIdent
	}
	if o.EnclaveInputstream != nil {
		toSerialize["enclaveInputstream"] = o.EnclaveInputstream
	}
	if o.EndingTime != nil {
		toSerialize["endingTime"] = o.EndingTime
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.InternalAttesationServer != nil {
		toSerialize["internalAttesationServer"] = o.InternalAttesationServer
	}
	if o.InternalIdent != nil {
		toSerialize["internalIdent"] = o.InternalIdent
	}
	if o.InternalRemoteControlServer != nil {
		toSerialize["internalRemoteControlServer"] = o.InternalRemoteControlServer
	}
	if o.InternalWireguardServer != nil {
		toSerialize["internalWireguardServer"] = o.InternalWireguardServer
	}
	if o.KubernetesEnclave != nil {
		toSerialize["kubernetesEnclave"] = o.KubernetesEnclave
	}
	if o.PortMapping != nil {
		toSerialize["portMapping"] = o.PortMapping
	}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	if o.Process != nil {
		toSerialize["process"] = o.Process
	}
	if o.PublicIdent != nil {
		toSerialize["publicIdent"] = o.PublicIdent
	}
	if o.RemoteControlServer != nil {
		toSerialize["remoteControlServer"] = o.RemoteControlServer
	}
	if o.SignerIdent != nil {
		toSerialize["signerIdent"] = o.SignerIdent
	}
	if o.StartupCMD != nil {
		toSerialize["startupCMD"] = o.StartupCMD
	}
	if o.StartupTime != nil {
		toSerialize["startupTime"] = o.StartupTime
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.WgInterface != nil {
		toSerialize["wgInterface"] = o.WgInterface
	}
	if o.WireguardPublicKey != nil {
		toSerialize["wireguardPublicKey"] = o.WireguardPublicKey
	}
	if o.WireguardServer != nil {
		toSerialize["wireguardServer"] = o.WireguardServer
	}
	return json.Marshal(toSerialize)
}

type NullableXmlNs0EnclaveProcessAllOf struct {
	value *XmlNs0EnclaveProcessAllOf
	isSet bool
}

func (v NullableXmlNs0EnclaveProcessAllOf) Get() *XmlNs0EnclaveProcessAllOf {
	return v.value
}

func (v *NullableXmlNs0EnclaveProcessAllOf) Set(val *XmlNs0EnclaveProcessAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableXmlNs0EnclaveProcessAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableXmlNs0EnclaveProcessAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXmlNs0EnclaveProcessAllOf(val *XmlNs0EnclaveProcessAllOf) *NullableXmlNs0EnclaveProcessAllOf {
	return &NullableXmlNs0EnclaveProcessAllOf{value: val, isSet: true}
}

func (v NullableXmlNs0EnclaveProcessAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXmlNs0EnclaveProcessAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


