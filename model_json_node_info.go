/*
 * madana-api
 *
 * <h1>API Quickstart Guide</h1>        <p>This documentation contains a Quickstart Guide, a few <a href=\"downloads.html\">sample clients</a>  for download and information about the available  <a href=\"resources.html\">endpoints</a>  and  <a href=\"data.html\">DataTypes</a>  </p>     <p>The <a target=\"_blank\" href=\"http://madana-explorer-staging.eu-central-1.elasticbeanstalk.com/login\">  MADANA Explorer</a> can be used to verify the interactions with the API</p>           <p>Internal use only. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a></p>         <br> <br>
 *
 * API version: 0.4.12
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package madanasampleclientgo
// JsonNodeInfo 
type JsonNodeInfo struct {
	// 
	HardwareBaseboard string `json:"hardwareBaseboard,omitempty"`
	// 
	PublicKey string `json:"publicKey,omitempty"`
	// 
	Memory string `json:"memory,omitempty"`
	// 
	Processors []string `json:"processors,omitempty"`
	// 
	CpuLogicalCount int32 `json:"cpuLogicalCount,omitempty"`
	// 
	ConnectionURL string `json:"connectionURL,omitempty"`
	// 
	OperatingSystem string `json:"operatingSystem,omitempty"`
	// 
	CpuPhysicalCores int32 `json:"cpuPhysicalCores,omitempty"`
	// 
	Status string `json:"status,omitempty"`
	// 
	OperatingSystemUptime float32 `json:"operatingSystemUptime,omitempty"`
	// 
	CpuModel string `json:"cpuModel,omitempty"`
	// 
	Owner string `json:"owner,omitempty"`
	// 
	CpuFamily string `json:"cpuFamily,omitempty"`
	// 
	CpuFrequency string `json:"cpuFrequency,omitempty"`
	// 
	HardwareFirmware string `json:"hardwareFirmware,omitempty"`
}
