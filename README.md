# Go API client for madanasampleclientgo

<h1>API Quickstart Guide</h1>
       <p>This documentation contains a Quickstart Guide, a few <a href=\"downloads.html\">sample clients</a>  for download and information about the available  <a href=\"resources.html\">endpoints</a>  and  <a href=\"data.html\">DataTypes</a>  </p>

   <p>The <a target=\"_blank\" href=\"http://madana-explorer-staging.eu-central-1.elasticbeanstalk.com/login\">  MADANA Explorer</a> can be used to verify the interactions with the API</p>
   
      <p>Internal use only. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a></p>
     

 <br> <br>

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.4.12
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./madanasampleclientgo"
```

## Documentation for API Endpoints

All URIs are relative to *http://api.madana.io/rest*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountServiceApi* | [**ActivateUser**](docs/AccountServiceApi.md#activateuser) | **Get** /account/activation/{token} | 
*AccountServiceApi* | [**CreateObject**](docs/AccountServiceApi.md#createobject) | **Post** /account/password | Sends an Password reset mail to the given MailAddress.
*AccountServiceApi* | [**RequestVerificationMail**](docs/AccountServiceApi.md#requestverificationmail) | **Get** /account/verifymail | Used to request a new  activation-mail for the user.
*AccountServiceApi* | [**UpdateObject**](docs/AccountServiceApi.md#updateobject) | **Put** /account/password | Receives the Password reset and tries to set the provided password for the user.
*AuthenticationServiceApi* | [**AuthenticateApplication**](docs/AuthenticationServiceApi.md#authenticateapplication) | **Post** /authentication/application | Authenticates a new application and returns the token.
*AuthenticationServiceApi* | [**AuthenticateEthereumWallet**](docs/AuthenticationServiceApi.md#authenticateethereumwallet) | **Post** /authentication/ethereum/{wallet} | 
*AuthenticationServiceApi* | [**AuthenticateUser**](docs/AuthenticationServiceApi.md#authenticateuser) | **Post** /authentication | Authenticates a new user and returns the token (  forbidden if the credentials cannot be validated ).
*AuthenticationServiceApi* | [**AuthenticateWithEthereumChallenge**](docs/AuthenticationServiceApi.md#authenticatewithethereumchallenge) | **Post** /authentication/ethereum/{wallet}/challenge | 
*AuthenticationServiceApi* | [**GetFractalAuthenticationURL**](docs/AuthenticationServiceApi.md#getfractalauthenticationurl) | **Get** /authentication/fractal | Returns the AUthorization URL to verify a Twitter Accounts.
*AuthenticationServiceApi* | [**GetNonceForEthereumWallet**](docs/AuthenticationServiceApi.md#getnonceforethereumwallet) | **Get** /authentication/ethereum/{wallet} | Returns a nonce for the client which is used as content for the to be created signature.
*AuthenticationServiceApi* | [**GetObject**](docs/AuthenticationServiceApi.md#getobject) | **Get** /authentication | Used to validate the active connection with the API.
*AuthenticationServiceApi* | [**GetTwitterAuthenticationURL**](docs/AuthenticationServiceApi.md#gettwitterauthenticationurl) | **Get** /authentication/twitter | Returns the AUthorization URL to verify a Twitter Accounts.
*AuthenticationServiceApi* | [**SetFacebookUID**](docs/AuthenticationServiceApi.md#setfacebookuid) | **Post** /authentication/facebook | Used as Callback URL when users have successfully authorized their facbeook account.
*AuthenticationServiceApi* | [**SetFractalUID**](docs/AuthenticationServiceApi.md#setfractaluid) | **Post** /authentication/fractal | 
*AuthenticationServiceApi* | [**SetTwitterUID**](docs/AuthenticationServiceApi.md#settwitteruid) | **Post** /authentication/twitter | 
*CertificateServiceApi* | [**AuthenticateCertificate**](docs/CertificateServiceApi.md#authenticatecertificate) | **Post** /certificates | Issues certificates for logged-in users.
*CertificateServiceApi* | [**GetCertificate**](docs/CertificateServiceApi.md#getcertificate) | **Get** /certificates/root | 
*CertificateServiceApi* | [**GetCertificate_0**](docs/CertificateServiceApi.md#getcertificate_0) | **Get** /certificates/{fingerprint} | 
*DataCollectionServiceApi* | [**GetMethodsForType**](docs/DataCollectionServiceApi.md#getmethodsfortype) | **Get** /datacollection/types/{name}/methods | 
*DataCollectionServiceApi* | [**GetNodes**](docs/DataCollectionServiceApi.md#getnodes) | **Get** /datacollection/methods | 
*DataCollectionServiceApi* | [**GetTypes**](docs/DataCollectionServiceApi.md#gettypes) | **Get** /datacollection/types | 
*NodeServiceApi* | [**GetNodes2**](docs/NodeServiceApi.md#getnodes2) | **Get** /nodes | 
*NodeServiceApi* | [**PostNodeInfo**](docs/NodeServiceApi.md#postnodeinfo) | **Post** /nodes | 
*OrganizationServiceApi* | [**GetNodes3**](docs/OrganizationServiceApi.md#getnodes3) | **Get** /organizations | 
*RequestServiceApi* | [**AddData**](docs/RequestServiceApi.md#adddata) | **Post** /requests/{uuid}/data | Is used to upload and park the data till the AnalysisRequest gets processed.
*RequestServiceApi* | [**CancelProcessing**](docs/RequestServiceApi.md#cancelprocessing) | **Post** /requests/{uuid}/cancel | Endpoint is called from the Analysis Processing entity to submit the result.
*RequestServiceApi* | [**CreateNewRequest**](docs/RequestServiceApi.md#createnewrequest) | **Post** /requests | Endpoint used to create a new Analysis Request.
*RequestServiceApi* | [**GetActions**](docs/RequestServiceApi.md#getactions) | **Get** /requests/actions | 
*RequestServiceApi* | [**GetAgent**](docs/RequestServiceApi.md#getagent) | **Get** /requests/{uuid}/agent | Is called from the APE to request all parked datasets.
*RequestServiceApi* | [**GetAllRequests**](docs/RequestServiceApi.md#getallrequests) | **Get** /requests | Returns UUIDs of existing analyses.
*RequestServiceApi* | [**GetData**](docs/RequestServiceApi.md#getdata) | **Get** /requests/{uuid}/data | Is called from the APE to request all parked datasets.
*RequestServiceApi* | [**GetRequest**](docs/RequestServiceApi.md#getrequest) | **Get** /requests/{uuid} | Returns the details for certain Request.
*RequestServiceApi* | [**GetResult**](docs/RequestServiceApi.md#getresult) | **Get** /requests/{uuid}/result | Can be called from creator to request the AnalysisResult.
*RequestServiceApi* | [**GiveConsent**](docs/RequestServiceApi.md#giveconsent) | **Post** /requests/{uuid}/consent | Used to give consent for request.
*RequestServiceApi* | [**InitRequestParameters**](docs/RequestServiceApi.md#initrequestparameters) | **Post** /requests/{uuid} | Endpoint used initialized addition datacollection parameters for requester.
*RequestServiceApi* | [**SetAgent**](docs/RequestServiceApi.md#setagent) | **Post** /requests/{uuid}/agent | Is called from the APE to request all parked datasets.
*RequestServiceApi* | [**SetResult**](docs/RequestServiceApi.md#setresult) | **Post** /requests/{uuid}/result | Endpoint is called from the Analysis Processing entity to submit the result.
*SocialPlatformServiceApi* | [**GetPlatforms**](docs/SocialPlatformServiceApi.md#getplatforms) | **Get** /platforms | Used to Handle Incoming Webhooks from Facebook.
*SocialPlatformServiceApi* | [**ListenTwitterWebhook**](docs/SocialPlatformServiceApi.md#listentwitterwebhook) | **Post** /platforms/twitter | Used to Handle Incoming Webhooks from Facebook.
*SocialPlatformServiceApi* | [**RegisterTwitterWebhook**](docs/SocialPlatformServiceApi.md#registertwitterwebhook) | **Get** /platforms/twitter | Used to Handle Incoming Webhooks from Twitter.
*SocialServiceApi* | [**GetMyProfile**](docs/SocialServiceApi.md#getmyprofile) | **Get** /social/profiles/me | 
*SocialServiceApi* | [**GetPlatforms2**](docs/SocialServiceApi.md#getplatforms2) | **Get** /social | Returns all Platforms / Systems that can be Connected to the MADANA Service.
*SocialServiceApi* | [**GetRanking**](docs/SocialServiceApi.md#getranking) | **Get** /social/ranking | Returns the Ranking by PTS within the System.
*SocialServiceApi* | [**GetSocialPlatformFeed**](docs/SocialServiceApi.md#getsocialplatformfeed) | **Get** /social/feed/{platform} | 
*SocialServiceApi* | [**GetUserProfile**](docs/SocialServiceApi.md#getuserprofile) | **Get** /social/profiles/{username} | 
*SocialServiceApi* | [**GetUserProfile_0**](docs/SocialServiceApi.md#getuserprofile_0) | **Get** /social/profiles/{username}/simple | 
*SystemServiceApi* | [**GetAllObjects**](docs/SystemServiceApi.md#getallobjects) | **Get** /system/health | 
*SystemServiceApi* | [**GetApplication**](docs/SystemServiceApi.md#getapplication) | **Get** /system/usage | Return the current application usage.
*UserServiceApi* | [**CreateObject2**](docs/UserServiceApi.md#createobject2) | **Post** /users | Creates a new user object.
*UserServiceApi* | [**DeleteObject**](docs/UserServiceApi.md#deleteobject) | **Delete** /users/{username} | Deletes an User based on the provided id and securitycontext.
*UserServiceApi* | [**DeleteObject_0**](docs/UserServiceApi.md#deleteobject_0) | **Delete** /users/{username}/social/{platform}/{ident} | Deletes linked account from the user and securitycontext.
*UserServiceApi* | [**GetAvatars**](docs/UserServiceApi.md#getavatars) | **Get** /users/{username}/avatars | 
*UserServiceApi* | [**GetCertificates**](docs/UserServiceApi.md#getcertificates) | **Get** /users/{username}/certificates | 
*UserServiceApi* | [**GetObject2**](docs/UserServiceApi.md#getobject2) | **Get** /users/{username} | 
*UserServiceApi* | [**SetAvatar**](docs/UserServiceApi.md#setavatar) | **Post** /users/{username}/avatars | 
*UserServiceApi* | [**SetSettings**](docs/UserServiceApi.md#setsettings) | **Post** /users/{username}/settings | 
*UserServiceApi* | [**UpdateObject2**](docs/UserServiceApi.md#updateobject2) | **Put** /users/{username} | Updates Userproperties based on the provided user object.


## Documentation For Models

 - [JsonAnalysis](docs/JsonAnalysis.md)
 - [JsonAnalysisAllOf](docs/JsonAnalysisAllOf.md)
 - [JsonAnalysisRequest](docs/JsonAnalysisRequest.md)
 - [JsonAnalysisRequestAction](docs/JsonAnalysisRequestAction.md)
 - [JsonAnalysisResult](docs/JsonAnalysisResult.md)
 - [JsonAnalysisResultSubGroup](docs/JsonAnalysisResultSubGroup.md)
 - [JsonAnalysisResultSubGroupAllOf](docs/JsonAnalysisResultSubGroupAllOf.md)
 - [JsonAnalysisVisualization](docs/JsonAnalysisVisualization.md)
 - [JsonDatasetInfo](docs/JsonDatasetInfo.md)
 - [JsonDatasetInfoAllOf](docs/JsonDatasetInfoAllOf.md)
 - [JsonMdnAUserObject](docs/JsonMdnAUserObject.md)
 - [JsonMdnCertificate](docs/JsonMdnCertificate.md)
 - [JsonMdnData](docs/JsonMdnData.md)
 - [JsonMdnMailAddress](docs/JsonMdnMailAddress.md)
 - [JsonMdnOAuthToken](docs/JsonMdnOAuthToken.md)
 - [JsonMdnPasswordReset](docs/JsonMdnPasswordReset.md)
 - [JsonMdnSetting](docs/JsonMdnSetting.md)
 - [JsonMdnSocialUserObject](docs/JsonMdnSocialUserObject.md)
 - [JsonMdnToken](docs/JsonMdnToken.md)
 - [JsonMdnUser](docs/JsonMdnUser.md)
 - [JsonMdnUserAllOf](docs/JsonMdnUserAllOf.md)
 - [JsonMdnUserCredentials](docs/JsonMdnUserCredentials.md)
 - [JsonMdnUserProfileImage](docs/JsonMdnUserProfileImage.md)
 - [JsonMdnUserSetting](docs/JsonMdnUserSetting.md)
 - [JsonMdnUserSettingAllOf](docs/JsonMdnUserSettingAllOf.md)
 - [JsonNodeInfo](docs/JsonNodeInfo.md)
 - [JsonSignedData](docs/JsonSignedData.md)
 - [XmlNs0MdnSetting](docs/XmlNs0MdnSetting.md)
 - [XmlNs0MdnSettingAllOf](docs/XmlNs0MdnSettingAllOf.md)
 - [XmlNs0MdnUserProfileImage](docs/XmlNs0MdnUserProfileImage.md)
 - [XmlNs0MdnUserProfileImageAllOf](docs/XmlNs0MdnUserProfileImageAllOf.md)
 - [XmlNs0MdnUserSetting](docs/XmlNs0MdnUserSetting.md)
 - [XmlNs0MdnUserSettingAllOf](docs/XmlNs0MdnUserSettingAllOf.md)
 - [XmlNs0SignedData](docs/XmlNs0SignedData.md)
 - [XmlNs0SignedDataAllOf](docs/XmlNs0SignedDataAllOf.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author



