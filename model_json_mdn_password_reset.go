/*
 * madana-api
 *
 * <h1>API Quickstart Guide</h1>        <p>This documentation contains a Quickstart Guide, a few <a href=\"downloads.html\">sample clients</a>  for download and information about the available  <a href=\"resources.html\">endpoints</a>  and  <a href=\"data.html\">DataTypes</a>  </p>     <p>The <a target=\"_blank\" href=\"http://madana-explorer-staging.eu-central-1.elasticbeanstalk.com/login\">  MADANA Explorer</a> can be used to verify the interactions with the API</p>           <p>Internal use only. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a></p>         <br> <br>
 *
 * API version: 0.4.12
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package madanasampleclientgo
// JsonMdnPasswordReset 
type JsonMdnPasswordReset struct {
	// 
	Password string `json:"password,omitempty"`
	// 
	Token string `json:"token,omitempty"`
	// 
	Mail string `json:"mail,omitempty"`
}
