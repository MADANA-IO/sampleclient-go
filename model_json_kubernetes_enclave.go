/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.44
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonKubernetesEnclave 
type JsonKubernetesEnclave struct {
	Environment *JsonEnvironment `json:"environment,omitempty"`
	// 
	SignerIdent *string `json:"signerIdent,omitempty"`
	// 
	Status *string `json:"status,omitempty"`
	// 
	ConsoleOutput *string `json:"consoleOutput,omitempty"`
	// 
	AttestationServer *string `json:"attestationServer,omitempty"`
	// 
	EnclaveInputstream *map[string]interface{} `json:"enclaveInputstream,omitempty"`
	// 
	EndingTime *string `json:"endingTime,omitempty"`
	// 
	PublicIdent *string `json:"publicIdent,omitempty"`
	// 
	InternalWireguardServer *string `json:"internalWireguardServer,omitempty"`
	// 
	InternalRemoteControlServer *string `json:"internalRemoteControlServer,omitempty"`
	Process *JsonProcess `json:"process,omitempty"`
	WgInterface *JsonWireguardInterface `json:"wgInterface,omitempty"`
	// 
	WireguardServer *string `json:"wireguardServer,omitempty"`
	// 
	WireguardPublicKey *string `json:"wireguardPublicKey,omitempty"`
	// 
	InternalIdent *string `json:"internalIdent,omitempty"`
	// 
	RemoteControlServer *string `json:"remoteControlServer,omitempty"`
	// 
	StartupTime *string `json:"startupTime,omitempty"`
	// 
	StartupCMD *string `json:"startupCMD,omitempty"`
	// 
	InternalAttesationServer *string `json:"internalAttesationServer,omitempty"`
	// 
	PortMapping *map[string]string `json:"portMapping,omitempty"`
	KubernetesEnclave *JsonKubernetesEnclave `json:"kubernetesEnclave,omitempty"`
	// 
	Ports *[]JsonEnclavePort `json:"ports,omitempty"`
	// 
	EnclaveIdent *string `json:"enclaveIdent,omitempty"`
	// 
	RemoteControlIP *string `json:"remoteControlIP,omitempty"`
	EnclaveReplicaSetEvents *JsonV1EventList `json:"enclaveReplicaSetEvents,omitempty"`
	// 
	PodPhase *string `json:"podPhase,omitempty"`
	EnclaveDeploymentEvents *JsonV1EventList `json:"enclaveDeploymentEvents,omitempty"`
	EnclavePodEvents *JsonV1EventList `json:"enclavePodEvents,omitempty"`
	// 
	DebugInfo *string `json:"debugInfo,omitempty"`
	// 
	IsUsingInitContainer *bool `json:"isUsingInitContainer,omitempty"`
	// 
	AttestationPort *int32 `json:"attestationPort,omitempty"`
	// 
	WireguardPort *int32 `json:"wireguardPort,omitempty"`
}

// NewJsonKubernetesEnclave instantiates a new JsonKubernetesEnclave object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonKubernetesEnclave() *JsonKubernetesEnclave {
	this := JsonKubernetesEnclave{}
	return &this
}

// NewJsonKubernetesEnclaveWithDefaults instantiates a new JsonKubernetesEnclave object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonKubernetesEnclaveWithDefaults() *JsonKubernetesEnclave {
	this := JsonKubernetesEnclave{}
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetEnvironment() JsonEnvironment {
	if o == nil || o.Environment == nil {
		var ret JsonEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetEnvironmentOk() (*JsonEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given JsonEnvironment and assigns it to the Environment field.
func (o *JsonKubernetesEnclave) SetEnvironment(v JsonEnvironment) {
	o.Environment = &v
}

// GetSignerIdent returns the SignerIdent field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetSignerIdent() string {
	if o == nil || o.SignerIdent == nil {
		var ret string
		return ret
	}
	return *o.SignerIdent
}

// GetSignerIdentOk returns a tuple with the SignerIdent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetSignerIdentOk() (*string, bool) {
	if o == nil || o.SignerIdent == nil {
		return nil, false
	}
	return o.SignerIdent, true
}

// HasSignerIdent returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasSignerIdent() bool {
	if o != nil && o.SignerIdent != nil {
		return true
	}

	return false
}

// SetSignerIdent gets a reference to the given string and assigns it to the SignerIdent field.
func (o *JsonKubernetesEnclave) SetSignerIdent(v string) {
	o.SignerIdent = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *JsonKubernetesEnclave) SetStatus(v string) {
	o.Status = &v
}

// GetConsoleOutput returns the ConsoleOutput field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetConsoleOutput() string {
	if o == nil || o.ConsoleOutput == nil {
		var ret string
		return ret
	}
	return *o.ConsoleOutput
}

// GetConsoleOutputOk returns a tuple with the ConsoleOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetConsoleOutputOk() (*string, bool) {
	if o == nil || o.ConsoleOutput == nil {
		return nil, false
	}
	return o.ConsoleOutput, true
}

// HasConsoleOutput returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasConsoleOutput() bool {
	if o != nil && o.ConsoleOutput != nil {
		return true
	}

	return false
}

// SetConsoleOutput gets a reference to the given string and assigns it to the ConsoleOutput field.
func (o *JsonKubernetesEnclave) SetConsoleOutput(v string) {
	o.ConsoleOutput = &v
}

// GetAttestationServer returns the AttestationServer field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetAttestationServer() string {
	if o == nil || o.AttestationServer == nil {
		var ret string
		return ret
	}
	return *o.AttestationServer
}

// GetAttestationServerOk returns a tuple with the AttestationServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetAttestationServerOk() (*string, bool) {
	if o == nil || o.AttestationServer == nil {
		return nil, false
	}
	return o.AttestationServer, true
}

// HasAttestationServer returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasAttestationServer() bool {
	if o != nil && o.AttestationServer != nil {
		return true
	}

	return false
}

// SetAttestationServer gets a reference to the given string and assigns it to the AttestationServer field.
func (o *JsonKubernetesEnclave) SetAttestationServer(v string) {
	o.AttestationServer = &v
}

// GetEnclaveInputstream returns the EnclaveInputstream field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetEnclaveInputstream() map[string]interface{} {
	if o == nil || o.EnclaveInputstream == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.EnclaveInputstream
}

// GetEnclaveInputstreamOk returns a tuple with the EnclaveInputstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetEnclaveInputstreamOk() (*map[string]interface{}, bool) {
	if o == nil || o.EnclaveInputstream == nil {
		return nil, false
	}
	return o.EnclaveInputstream, true
}

// HasEnclaveInputstream returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasEnclaveInputstream() bool {
	if o != nil && o.EnclaveInputstream != nil {
		return true
	}

	return false
}

// SetEnclaveInputstream gets a reference to the given map[string]interface{} and assigns it to the EnclaveInputstream field.
func (o *JsonKubernetesEnclave) SetEnclaveInputstream(v map[string]interface{}) {
	o.EnclaveInputstream = &v
}

// GetEndingTime returns the EndingTime field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetEndingTime() string {
	if o == nil || o.EndingTime == nil {
		var ret string
		return ret
	}
	return *o.EndingTime
}

// GetEndingTimeOk returns a tuple with the EndingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetEndingTimeOk() (*string, bool) {
	if o == nil || o.EndingTime == nil {
		return nil, false
	}
	return o.EndingTime, true
}

// HasEndingTime returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasEndingTime() bool {
	if o != nil && o.EndingTime != nil {
		return true
	}

	return false
}

// SetEndingTime gets a reference to the given string and assigns it to the EndingTime field.
func (o *JsonKubernetesEnclave) SetEndingTime(v string) {
	o.EndingTime = &v
}

// GetPublicIdent returns the PublicIdent field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetPublicIdent() string {
	if o == nil || o.PublicIdent == nil {
		var ret string
		return ret
	}
	return *o.PublicIdent
}

// GetPublicIdentOk returns a tuple with the PublicIdent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetPublicIdentOk() (*string, bool) {
	if o == nil || o.PublicIdent == nil {
		return nil, false
	}
	return o.PublicIdent, true
}

// HasPublicIdent returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasPublicIdent() bool {
	if o != nil && o.PublicIdent != nil {
		return true
	}

	return false
}

// SetPublicIdent gets a reference to the given string and assigns it to the PublicIdent field.
func (o *JsonKubernetesEnclave) SetPublicIdent(v string) {
	o.PublicIdent = &v
}

// GetInternalWireguardServer returns the InternalWireguardServer field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetInternalWireguardServer() string {
	if o == nil || o.InternalWireguardServer == nil {
		var ret string
		return ret
	}
	return *o.InternalWireguardServer
}

// GetInternalWireguardServerOk returns a tuple with the InternalWireguardServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetInternalWireguardServerOk() (*string, bool) {
	if o == nil || o.InternalWireguardServer == nil {
		return nil, false
	}
	return o.InternalWireguardServer, true
}

// HasInternalWireguardServer returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasInternalWireguardServer() bool {
	if o != nil && o.InternalWireguardServer != nil {
		return true
	}

	return false
}

// SetInternalWireguardServer gets a reference to the given string and assigns it to the InternalWireguardServer field.
func (o *JsonKubernetesEnclave) SetInternalWireguardServer(v string) {
	o.InternalWireguardServer = &v
}

// GetInternalRemoteControlServer returns the InternalRemoteControlServer field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetInternalRemoteControlServer() string {
	if o == nil || o.InternalRemoteControlServer == nil {
		var ret string
		return ret
	}
	return *o.InternalRemoteControlServer
}

// GetInternalRemoteControlServerOk returns a tuple with the InternalRemoteControlServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetInternalRemoteControlServerOk() (*string, bool) {
	if o == nil || o.InternalRemoteControlServer == nil {
		return nil, false
	}
	return o.InternalRemoteControlServer, true
}

// HasInternalRemoteControlServer returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasInternalRemoteControlServer() bool {
	if o != nil && o.InternalRemoteControlServer != nil {
		return true
	}

	return false
}

// SetInternalRemoteControlServer gets a reference to the given string and assigns it to the InternalRemoteControlServer field.
func (o *JsonKubernetesEnclave) SetInternalRemoteControlServer(v string) {
	o.InternalRemoteControlServer = &v
}

// GetProcess returns the Process field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetProcess() JsonProcess {
	if o == nil || o.Process == nil {
		var ret JsonProcess
		return ret
	}
	return *o.Process
}

// GetProcessOk returns a tuple with the Process field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetProcessOk() (*JsonProcess, bool) {
	if o == nil || o.Process == nil {
		return nil, false
	}
	return o.Process, true
}

// HasProcess returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasProcess() bool {
	if o != nil && o.Process != nil {
		return true
	}

	return false
}

// SetProcess gets a reference to the given JsonProcess and assigns it to the Process field.
func (o *JsonKubernetesEnclave) SetProcess(v JsonProcess) {
	o.Process = &v
}

// GetWgInterface returns the WgInterface field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetWgInterface() JsonWireguardInterface {
	if o == nil || o.WgInterface == nil {
		var ret JsonWireguardInterface
		return ret
	}
	return *o.WgInterface
}

// GetWgInterfaceOk returns a tuple with the WgInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetWgInterfaceOk() (*JsonWireguardInterface, bool) {
	if o == nil || o.WgInterface == nil {
		return nil, false
	}
	return o.WgInterface, true
}

// HasWgInterface returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasWgInterface() bool {
	if o != nil && o.WgInterface != nil {
		return true
	}

	return false
}

// SetWgInterface gets a reference to the given JsonWireguardInterface and assigns it to the WgInterface field.
func (o *JsonKubernetesEnclave) SetWgInterface(v JsonWireguardInterface) {
	o.WgInterface = &v
}

// GetWireguardServer returns the WireguardServer field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetWireguardServer() string {
	if o == nil || o.WireguardServer == nil {
		var ret string
		return ret
	}
	return *o.WireguardServer
}

// GetWireguardServerOk returns a tuple with the WireguardServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetWireguardServerOk() (*string, bool) {
	if o == nil || o.WireguardServer == nil {
		return nil, false
	}
	return o.WireguardServer, true
}

// HasWireguardServer returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasWireguardServer() bool {
	if o != nil && o.WireguardServer != nil {
		return true
	}

	return false
}

// SetWireguardServer gets a reference to the given string and assigns it to the WireguardServer field.
func (o *JsonKubernetesEnclave) SetWireguardServer(v string) {
	o.WireguardServer = &v
}

// GetWireguardPublicKey returns the WireguardPublicKey field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetWireguardPublicKey() string {
	if o == nil || o.WireguardPublicKey == nil {
		var ret string
		return ret
	}
	return *o.WireguardPublicKey
}

// GetWireguardPublicKeyOk returns a tuple with the WireguardPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetWireguardPublicKeyOk() (*string, bool) {
	if o == nil || o.WireguardPublicKey == nil {
		return nil, false
	}
	return o.WireguardPublicKey, true
}

// HasWireguardPublicKey returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasWireguardPublicKey() bool {
	if o != nil && o.WireguardPublicKey != nil {
		return true
	}

	return false
}

// SetWireguardPublicKey gets a reference to the given string and assigns it to the WireguardPublicKey field.
func (o *JsonKubernetesEnclave) SetWireguardPublicKey(v string) {
	o.WireguardPublicKey = &v
}

// GetInternalIdent returns the InternalIdent field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetInternalIdent() string {
	if o == nil || o.InternalIdent == nil {
		var ret string
		return ret
	}
	return *o.InternalIdent
}

// GetInternalIdentOk returns a tuple with the InternalIdent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetInternalIdentOk() (*string, bool) {
	if o == nil || o.InternalIdent == nil {
		return nil, false
	}
	return o.InternalIdent, true
}

// HasInternalIdent returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasInternalIdent() bool {
	if o != nil && o.InternalIdent != nil {
		return true
	}

	return false
}

// SetInternalIdent gets a reference to the given string and assigns it to the InternalIdent field.
func (o *JsonKubernetesEnclave) SetInternalIdent(v string) {
	o.InternalIdent = &v
}

// GetRemoteControlServer returns the RemoteControlServer field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetRemoteControlServer() string {
	if o == nil || o.RemoteControlServer == nil {
		var ret string
		return ret
	}
	return *o.RemoteControlServer
}

// GetRemoteControlServerOk returns a tuple with the RemoteControlServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetRemoteControlServerOk() (*string, bool) {
	if o == nil || o.RemoteControlServer == nil {
		return nil, false
	}
	return o.RemoteControlServer, true
}

// HasRemoteControlServer returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasRemoteControlServer() bool {
	if o != nil && o.RemoteControlServer != nil {
		return true
	}

	return false
}

// SetRemoteControlServer gets a reference to the given string and assigns it to the RemoteControlServer field.
func (o *JsonKubernetesEnclave) SetRemoteControlServer(v string) {
	o.RemoteControlServer = &v
}

// GetStartupTime returns the StartupTime field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetStartupTime() string {
	if o == nil || o.StartupTime == nil {
		var ret string
		return ret
	}
	return *o.StartupTime
}

// GetStartupTimeOk returns a tuple with the StartupTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetStartupTimeOk() (*string, bool) {
	if o == nil || o.StartupTime == nil {
		return nil, false
	}
	return o.StartupTime, true
}

// HasStartupTime returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasStartupTime() bool {
	if o != nil && o.StartupTime != nil {
		return true
	}

	return false
}

// SetStartupTime gets a reference to the given string and assigns it to the StartupTime field.
func (o *JsonKubernetesEnclave) SetStartupTime(v string) {
	o.StartupTime = &v
}

// GetStartupCMD returns the StartupCMD field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetStartupCMD() string {
	if o == nil || o.StartupCMD == nil {
		var ret string
		return ret
	}
	return *o.StartupCMD
}

// GetStartupCMDOk returns a tuple with the StartupCMD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetStartupCMDOk() (*string, bool) {
	if o == nil || o.StartupCMD == nil {
		return nil, false
	}
	return o.StartupCMD, true
}

// HasStartupCMD returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasStartupCMD() bool {
	if o != nil && o.StartupCMD != nil {
		return true
	}

	return false
}

// SetStartupCMD gets a reference to the given string and assigns it to the StartupCMD field.
func (o *JsonKubernetesEnclave) SetStartupCMD(v string) {
	o.StartupCMD = &v
}

// GetInternalAttesationServer returns the InternalAttesationServer field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetInternalAttesationServer() string {
	if o == nil || o.InternalAttesationServer == nil {
		var ret string
		return ret
	}
	return *o.InternalAttesationServer
}

// GetInternalAttesationServerOk returns a tuple with the InternalAttesationServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetInternalAttesationServerOk() (*string, bool) {
	if o == nil || o.InternalAttesationServer == nil {
		return nil, false
	}
	return o.InternalAttesationServer, true
}

// HasInternalAttesationServer returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasInternalAttesationServer() bool {
	if o != nil && o.InternalAttesationServer != nil {
		return true
	}

	return false
}

// SetInternalAttesationServer gets a reference to the given string and assigns it to the InternalAttesationServer field.
func (o *JsonKubernetesEnclave) SetInternalAttesationServer(v string) {
	o.InternalAttesationServer = &v
}

// GetPortMapping returns the PortMapping field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetPortMapping() map[string]string {
	if o == nil || o.PortMapping == nil {
		var ret map[string]string
		return ret
	}
	return *o.PortMapping
}

// GetPortMappingOk returns a tuple with the PortMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetPortMappingOk() (*map[string]string, bool) {
	if o == nil || o.PortMapping == nil {
		return nil, false
	}
	return o.PortMapping, true
}

// HasPortMapping returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasPortMapping() bool {
	if o != nil && o.PortMapping != nil {
		return true
	}

	return false
}

// SetPortMapping gets a reference to the given map[string]string and assigns it to the PortMapping field.
func (o *JsonKubernetesEnclave) SetPortMapping(v map[string]string) {
	o.PortMapping = &v
}

// GetKubernetesEnclave returns the KubernetesEnclave field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetKubernetesEnclave() JsonKubernetesEnclave {
	if o == nil || o.KubernetesEnclave == nil {
		var ret JsonKubernetesEnclave
		return ret
	}
	return *o.KubernetesEnclave
}

// GetKubernetesEnclaveOk returns a tuple with the KubernetesEnclave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetKubernetesEnclaveOk() (*JsonKubernetesEnclave, bool) {
	if o == nil || o.KubernetesEnclave == nil {
		return nil, false
	}
	return o.KubernetesEnclave, true
}

// HasKubernetesEnclave returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasKubernetesEnclave() bool {
	if o != nil && o.KubernetesEnclave != nil {
		return true
	}

	return false
}

// SetKubernetesEnclave gets a reference to the given JsonKubernetesEnclave and assigns it to the KubernetesEnclave field.
func (o *JsonKubernetesEnclave) SetKubernetesEnclave(v JsonKubernetesEnclave) {
	o.KubernetesEnclave = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetPorts() []JsonEnclavePort {
	if o == nil || o.Ports == nil {
		var ret []JsonEnclavePort
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetPortsOk() (*[]JsonEnclavePort, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []JsonEnclavePort and assigns it to the Ports field.
func (o *JsonKubernetesEnclave) SetPorts(v []JsonEnclavePort) {
	o.Ports = &v
}

// GetEnclaveIdent returns the EnclaveIdent field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetEnclaveIdent() string {
	if o == nil || o.EnclaveIdent == nil {
		var ret string
		return ret
	}
	return *o.EnclaveIdent
}

// GetEnclaveIdentOk returns a tuple with the EnclaveIdent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetEnclaveIdentOk() (*string, bool) {
	if o == nil || o.EnclaveIdent == nil {
		return nil, false
	}
	return o.EnclaveIdent, true
}

// HasEnclaveIdent returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasEnclaveIdent() bool {
	if o != nil && o.EnclaveIdent != nil {
		return true
	}

	return false
}

// SetEnclaveIdent gets a reference to the given string and assigns it to the EnclaveIdent field.
func (o *JsonKubernetesEnclave) SetEnclaveIdent(v string) {
	o.EnclaveIdent = &v
}

// GetRemoteControlIP returns the RemoteControlIP field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetRemoteControlIP() string {
	if o == nil || o.RemoteControlIP == nil {
		var ret string
		return ret
	}
	return *o.RemoteControlIP
}

// GetRemoteControlIPOk returns a tuple with the RemoteControlIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetRemoteControlIPOk() (*string, bool) {
	if o == nil || o.RemoteControlIP == nil {
		return nil, false
	}
	return o.RemoteControlIP, true
}

// HasRemoteControlIP returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasRemoteControlIP() bool {
	if o != nil && o.RemoteControlIP != nil {
		return true
	}

	return false
}

// SetRemoteControlIP gets a reference to the given string and assigns it to the RemoteControlIP field.
func (o *JsonKubernetesEnclave) SetRemoteControlIP(v string) {
	o.RemoteControlIP = &v
}

// GetEnclaveReplicaSetEvents returns the EnclaveReplicaSetEvents field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetEnclaveReplicaSetEvents() JsonV1EventList {
	if o == nil || o.EnclaveReplicaSetEvents == nil {
		var ret JsonV1EventList
		return ret
	}
	return *o.EnclaveReplicaSetEvents
}

// GetEnclaveReplicaSetEventsOk returns a tuple with the EnclaveReplicaSetEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetEnclaveReplicaSetEventsOk() (*JsonV1EventList, bool) {
	if o == nil || o.EnclaveReplicaSetEvents == nil {
		return nil, false
	}
	return o.EnclaveReplicaSetEvents, true
}

// HasEnclaveReplicaSetEvents returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasEnclaveReplicaSetEvents() bool {
	if o != nil && o.EnclaveReplicaSetEvents != nil {
		return true
	}

	return false
}

// SetEnclaveReplicaSetEvents gets a reference to the given JsonV1EventList and assigns it to the EnclaveReplicaSetEvents field.
func (o *JsonKubernetesEnclave) SetEnclaveReplicaSetEvents(v JsonV1EventList) {
	o.EnclaveReplicaSetEvents = &v
}

// GetPodPhase returns the PodPhase field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetPodPhase() string {
	if o == nil || o.PodPhase == nil {
		var ret string
		return ret
	}
	return *o.PodPhase
}

// GetPodPhaseOk returns a tuple with the PodPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetPodPhaseOk() (*string, bool) {
	if o == nil || o.PodPhase == nil {
		return nil, false
	}
	return o.PodPhase, true
}

// HasPodPhase returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasPodPhase() bool {
	if o != nil && o.PodPhase != nil {
		return true
	}

	return false
}

// SetPodPhase gets a reference to the given string and assigns it to the PodPhase field.
func (o *JsonKubernetesEnclave) SetPodPhase(v string) {
	o.PodPhase = &v
}

// GetEnclaveDeploymentEvents returns the EnclaveDeploymentEvents field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetEnclaveDeploymentEvents() JsonV1EventList {
	if o == nil || o.EnclaveDeploymentEvents == nil {
		var ret JsonV1EventList
		return ret
	}
	return *o.EnclaveDeploymentEvents
}

// GetEnclaveDeploymentEventsOk returns a tuple with the EnclaveDeploymentEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetEnclaveDeploymentEventsOk() (*JsonV1EventList, bool) {
	if o == nil || o.EnclaveDeploymentEvents == nil {
		return nil, false
	}
	return o.EnclaveDeploymentEvents, true
}

// HasEnclaveDeploymentEvents returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasEnclaveDeploymentEvents() bool {
	if o != nil && o.EnclaveDeploymentEvents != nil {
		return true
	}

	return false
}

// SetEnclaveDeploymentEvents gets a reference to the given JsonV1EventList and assigns it to the EnclaveDeploymentEvents field.
func (o *JsonKubernetesEnclave) SetEnclaveDeploymentEvents(v JsonV1EventList) {
	o.EnclaveDeploymentEvents = &v
}

// GetEnclavePodEvents returns the EnclavePodEvents field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetEnclavePodEvents() JsonV1EventList {
	if o == nil || o.EnclavePodEvents == nil {
		var ret JsonV1EventList
		return ret
	}
	return *o.EnclavePodEvents
}

// GetEnclavePodEventsOk returns a tuple with the EnclavePodEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetEnclavePodEventsOk() (*JsonV1EventList, bool) {
	if o == nil || o.EnclavePodEvents == nil {
		return nil, false
	}
	return o.EnclavePodEvents, true
}

// HasEnclavePodEvents returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasEnclavePodEvents() bool {
	if o != nil && o.EnclavePodEvents != nil {
		return true
	}

	return false
}

// SetEnclavePodEvents gets a reference to the given JsonV1EventList and assigns it to the EnclavePodEvents field.
func (o *JsonKubernetesEnclave) SetEnclavePodEvents(v JsonV1EventList) {
	o.EnclavePodEvents = &v
}

// GetDebugInfo returns the DebugInfo field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetDebugInfo() string {
	if o == nil || o.DebugInfo == nil {
		var ret string
		return ret
	}
	return *o.DebugInfo
}

// GetDebugInfoOk returns a tuple with the DebugInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetDebugInfoOk() (*string, bool) {
	if o == nil || o.DebugInfo == nil {
		return nil, false
	}
	return o.DebugInfo, true
}

// HasDebugInfo returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasDebugInfo() bool {
	if o != nil && o.DebugInfo != nil {
		return true
	}

	return false
}

// SetDebugInfo gets a reference to the given string and assigns it to the DebugInfo field.
func (o *JsonKubernetesEnclave) SetDebugInfo(v string) {
	o.DebugInfo = &v
}

// GetIsUsingInitContainer returns the IsUsingInitContainer field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetIsUsingInitContainer() bool {
	if o == nil || o.IsUsingInitContainer == nil {
		var ret bool
		return ret
	}
	return *o.IsUsingInitContainer
}

// GetIsUsingInitContainerOk returns a tuple with the IsUsingInitContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetIsUsingInitContainerOk() (*bool, bool) {
	if o == nil || o.IsUsingInitContainer == nil {
		return nil, false
	}
	return o.IsUsingInitContainer, true
}

// HasIsUsingInitContainer returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasIsUsingInitContainer() bool {
	if o != nil && o.IsUsingInitContainer != nil {
		return true
	}

	return false
}

// SetIsUsingInitContainer gets a reference to the given bool and assigns it to the IsUsingInitContainer field.
func (o *JsonKubernetesEnclave) SetIsUsingInitContainer(v bool) {
	o.IsUsingInitContainer = &v
}

// GetAttestationPort returns the AttestationPort field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetAttestationPort() int32 {
	if o == nil || o.AttestationPort == nil {
		var ret int32
		return ret
	}
	return *o.AttestationPort
}

// GetAttestationPortOk returns a tuple with the AttestationPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetAttestationPortOk() (*int32, bool) {
	if o == nil || o.AttestationPort == nil {
		return nil, false
	}
	return o.AttestationPort, true
}

// HasAttestationPort returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasAttestationPort() bool {
	if o != nil && o.AttestationPort != nil {
		return true
	}

	return false
}

// SetAttestationPort gets a reference to the given int32 and assigns it to the AttestationPort field.
func (o *JsonKubernetesEnclave) SetAttestationPort(v int32) {
	o.AttestationPort = &v
}

// GetWireguardPort returns the WireguardPort field value if set, zero value otherwise.
func (o *JsonKubernetesEnclave) GetWireguardPort() int32 {
	if o == nil || o.WireguardPort == nil {
		var ret int32
		return ret
	}
	return *o.WireguardPort
}

// GetWireguardPortOk returns a tuple with the WireguardPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonKubernetesEnclave) GetWireguardPortOk() (*int32, bool) {
	if o == nil || o.WireguardPort == nil {
		return nil, false
	}
	return o.WireguardPort, true
}

// HasWireguardPort returns a boolean if a field has been set.
func (o *JsonKubernetesEnclave) HasWireguardPort() bool {
	if o != nil && o.WireguardPort != nil {
		return true
	}

	return false
}

// SetWireguardPort gets a reference to the given int32 and assigns it to the WireguardPort field.
func (o *JsonKubernetesEnclave) SetWireguardPort(v int32) {
	o.WireguardPort = &v
}

func (o JsonKubernetesEnclave) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.SignerIdent != nil {
		toSerialize["signerIdent"] = o.SignerIdent
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ConsoleOutput != nil {
		toSerialize["consoleOutput"] = o.ConsoleOutput
	}
	if o.AttestationServer != nil {
		toSerialize["attestationServer"] = o.AttestationServer
	}
	if o.EnclaveInputstream != nil {
		toSerialize["enclaveInputstream"] = o.EnclaveInputstream
	}
	if o.EndingTime != nil {
		toSerialize["endingTime"] = o.EndingTime
	}
	if o.PublicIdent != nil {
		toSerialize["publicIdent"] = o.PublicIdent
	}
	if o.InternalWireguardServer != nil {
		toSerialize["internalWireguardServer"] = o.InternalWireguardServer
	}
	if o.InternalRemoteControlServer != nil {
		toSerialize["internalRemoteControlServer"] = o.InternalRemoteControlServer
	}
	if o.Process != nil {
		toSerialize["process"] = o.Process
	}
	if o.WgInterface != nil {
		toSerialize["wgInterface"] = o.WgInterface
	}
	if o.WireguardServer != nil {
		toSerialize["wireguardServer"] = o.WireguardServer
	}
	if o.WireguardPublicKey != nil {
		toSerialize["wireguardPublicKey"] = o.WireguardPublicKey
	}
	if o.InternalIdent != nil {
		toSerialize["internalIdent"] = o.InternalIdent
	}
	if o.RemoteControlServer != nil {
		toSerialize["remoteControlServer"] = o.RemoteControlServer
	}
	if o.StartupTime != nil {
		toSerialize["startupTime"] = o.StartupTime
	}
	if o.StartupCMD != nil {
		toSerialize["startupCMD"] = o.StartupCMD
	}
	if o.InternalAttesationServer != nil {
		toSerialize["internalAttesationServer"] = o.InternalAttesationServer
	}
	if o.PortMapping != nil {
		toSerialize["portMapping"] = o.PortMapping
	}
	if o.KubernetesEnclave != nil {
		toSerialize["kubernetesEnclave"] = o.KubernetesEnclave
	}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	if o.EnclaveIdent != nil {
		toSerialize["enclaveIdent"] = o.EnclaveIdent
	}
	if o.RemoteControlIP != nil {
		toSerialize["remoteControlIP"] = o.RemoteControlIP
	}
	if o.EnclaveReplicaSetEvents != nil {
		toSerialize["enclaveReplicaSetEvents"] = o.EnclaveReplicaSetEvents
	}
	if o.PodPhase != nil {
		toSerialize["podPhase"] = o.PodPhase
	}
	if o.EnclaveDeploymentEvents != nil {
		toSerialize["enclaveDeploymentEvents"] = o.EnclaveDeploymentEvents
	}
	if o.EnclavePodEvents != nil {
		toSerialize["enclavePodEvents"] = o.EnclavePodEvents
	}
	if o.DebugInfo != nil {
		toSerialize["debugInfo"] = o.DebugInfo
	}
	if o.IsUsingInitContainer != nil {
		toSerialize["isUsingInitContainer"] = o.IsUsingInitContainer
	}
	if o.AttestationPort != nil {
		toSerialize["attestationPort"] = o.AttestationPort
	}
	if o.WireguardPort != nil {
		toSerialize["wireguardPort"] = o.WireguardPort
	}
	return json.Marshal(toSerialize)
}

type NullableJsonKubernetesEnclave struct {
	value *JsonKubernetesEnclave
	isSet bool
}

func (v NullableJsonKubernetesEnclave) Get() *JsonKubernetesEnclave {
	return v.value
}

func (v *NullableJsonKubernetesEnclave) Set(val *JsonKubernetesEnclave) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonKubernetesEnclave) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonKubernetesEnclave) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonKubernetesEnclave(val *JsonKubernetesEnclave) *NullableJsonKubernetesEnclave {
	return &NullableJsonKubernetesEnclave{value: val, isSet: true}
}

func (v NullableJsonKubernetesEnclave) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonKubernetesEnclave) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


