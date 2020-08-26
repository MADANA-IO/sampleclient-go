/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.4.14-master.23
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package madanaapiclient
// JsonEnclaveProcess 
type JsonEnclaveProcess struct {
	// 
	EnclaveInputstream map[string]interface{} `json:"enclaveInputstream,omitempty"`
	// 
	InternalAttesationServer string `json:"internalAttesationServer,omitempty"`
	// 
	StartupTime string `json:"startupTime,omitempty"`
	// 
	WireguardPublicKey string `json:"wireguardPublicKey,omitempty"`
	// 
	SignerIdent string `json:"signerIdent,omitempty"`
	// 
	StartupCMD string `json:"startupCMD,omitempty"`
	// 
	InternalIdent string `json:"internalIdent,omitempty"`
	// 
	AttestationServer string `json:"attestationServer,omitempty"`
	// 
	RemoteControlServer string `json:"remoteControlServer,omitempty"`
	// 
	InternalRemoteControlServer string `json:"internalRemoteControlServer,omitempty"`
	// 
	EnclaveIdent string `json:"enclaveIdent,omitempty"`
	// 
	Status string `json:"status,omitempty"`
	// 
	EndingTime string `json:"endingTime,omitempty"`
	WgInterface JsonWireguardInterface `json:"wgInterface,omitempty"`
	Process JsonProcess `json:"process,omitempty"`
	// 
	PublicIdent string `json:"publicIdent,omitempty"`
	// 
	ConsoleOutput string `json:"consoleOutput,omitempty"`
	Environment JsonEnvironment `json:"environment,omitempty"`
}
