/*
 * madana-api
 *
 * <h1>API Quickstart Guide</h1>        <p>This documentation contains a Quickstart Guide, a few <a href=\"downloads.html\">sample clients</a>  for download and information about the available  <a href=\"resources.html\">endpoints</a>  and  <a href=\"data.html\">DataTypes</a>  </p>     <p>The <a target=\"_blank\" href=\"http://madana-explorer-staging.eu-central-1.elasticbeanstalk.com/login\">  MADANA Explorer</a> can be used to verify the interactions with the API</p>           <p>Internal use only. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a></p>         <br> <br>
 *
 * API version: 0.4.12
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package madanasampleclientgo
// JsonDatasetInfo 
type JsonDatasetInfo struct {
	// 
	Fingerpint string `json:"fingerpint,omitempty"`
	// 
	Signature string `json:"signature,omitempty"`
	// 
	Data string `json:"data,omitempty"`
	// 
	Size string `json:"size,omitempty"`
	// 
	Hash string `json:"hash,omitempty"`
	// 
	Creationdate string `json:"creationdate,omitempty"`
	// 
	Fingerprint string `json:"fingerprint,omitempty"`
}
