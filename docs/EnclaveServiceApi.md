# \EnclaveServiceApi

All URIs are relative to *http://api.madana.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddHistory**](EnclaveServiceApi.md#AddHistory) | **Post** /enclaves/{uuid}/history | 
[**ApproveEnclave**](EnclaveServiceApi.md#ApproveEnclave) | **Post** /enclaves/{uuid}/approval | 
[**AssignEnclaveAgent**](EnclaveServiceApi.md#AssignEnclaveAgent) | **Post** /enclaves/{uuid}/assign | 
[**AttestateEnclave**](EnclaveServiceApi.md#AttestateEnclave) | **Post** /enclaves/{uuid}/attestation | 
[**CreateEnclaveRunRequest**](EnclaveServiceApi.md#CreateEnclaveRunRequest) | **Post** /enclaves | 
[**GetEnclave**](EnclaveServiceApi.md#GetEnclave) | **Get** /enclaves/{uuid} | 
[**GetEnclaveTypes**](EnclaveServiceApi.md#GetEnclaveTypes) | **Get** /enclaves/types | 
[**GetEnclaves**](EnclaveServiceApi.md#GetEnclaves) | **Get** /enclaves | Returns UUIDs of existing analyses.
[**GetStats**](EnclaveServiceApi.md#GetStats) | **Get** /enclaves/stats | 
[**KillEnclave**](EnclaveServiceApi.md#KillEnclave) | **Post** /enclaves/{uuid}/kill | 



## AddHistory

> *os.File AddHistory(ctx, uuid).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uuid := "uuid_example" // string | 
    body := *openapiclient.NewJsonSignedData() // JsonSignedData |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnclaveServiceApi.AddHistory(context.Background(), uuid).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.AddHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddHistory`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.AddHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**JsonSignedData**](JsonSignedData.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApproveEnclave

> *os.File ApproveEnclave(ctx, uuid).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uuid := "uuid_example" // string | 
    body := *openapiclient.NewJsonEnclaveRunningAttestationApproval() // JsonEnclaveRunningAttestationApproval |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnclaveServiceApi.ApproveEnclave(context.Background(), uuid).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.ApproveEnclave``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApproveEnclave`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.ApproveEnclave`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveEnclaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**JsonEnclaveRunningAttestationApproval**](JsonEnclaveRunningAttestationApproval.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignEnclaveAgent

> *os.File AssignEnclaveAgent(ctx, uuid).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uuid := "uuid_example" // string | 
    body := *openapiclient.NewJsonNodeInfo() // JsonNodeInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnclaveServiceApi.AssignEnclaveAgent(context.Background(), uuid).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.AssignEnclaveAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignEnclaveAgent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.AssignEnclaveAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignEnclaveAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**JsonNodeInfo**](JsonNodeInfo.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttestateEnclave

> *os.File AttestateEnclave(ctx, uuid).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uuid := "uuid_example" // string | 
    body := *openapiclient.NewJsonEnclaveRunningAttestation() // JsonEnclaveRunningAttestation |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnclaveServiceApi.AttestateEnclave(context.Background(), uuid).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.AttestateEnclave``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttestateEnclave`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.AttestateEnclave`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttestateEnclaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**JsonEnclaveRunningAttestation**](JsonEnclaveRunningAttestation.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnclaveRunRequest

> *os.File CreateEnclaveRunRequest(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewJsonEnclaveRunRequest() // JsonEnclaveRunRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnclaveServiceApi.CreateEnclaveRunRequest(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.CreateEnclaveRunRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnclaveRunRequest`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.CreateEnclaveRunRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnclaveRunRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**JsonEnclaveRunRequest**](JsonEnclaveRunRequest.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnclave

> *os.File GetEnclave(ctx, uuid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnclaveServiceApi.GetEnclave(context.Background(), uuid).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.GetEnclave``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnclave`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.GetEnclave`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnclaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnclaveTypes

> *os.File GetEnclaveTypes(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnclaveServiceApi.GetEnclaveTypes(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.GetEnclaveTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnclaveTypes`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.GetEnclaveTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnclaveTypesRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnclaves

> *os.File GetEnclaves(ctx).Authorization(authorization).Created(created).Limit(limit).Offset(offset).Status(status).Execute()

Returns UUIDs of existing analyses.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c (optional)
    created := "created_example" // string | - if Queryparam \"created=true\" only the UUIDs of own Requests are shown (optional) (default to "true")
    limit := "limit_example" // string | Used for offset pagination. Limit/Offset Paging would look like GET /request?limit=20&offset=100. This query would return the 20 rows starting with the 100th row (optional) (default to "30")
    offset := "offset_example" // string | Used for offset pagination. Limit/Offset Paging would look like GET /request?limit=20&offset=100. This query would return the 20 rows starting with the 100th row (optional) (default to "0")
    status := "status_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnclaveServiceApi.GetEnclaves(context.Background()).Authorization(authorization).Created(created).Limit(limit).Offset(offset).Status(status).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.GetEnclaves``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnclaves`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.GetEnclaves`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnclavesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c | 
 **created** | **string** | - if Queryparam \&quot;created&#x3D;true\&quot; only the UUIDs of own Requests are shown | [default to &quot;true&quot;]
 **limit** | **string** | Used for offset pagination. Limit/Offset Paging would look like GET /request?limit&#x3D;20&amp;offset&#x3D;100. This query would return the 20 rows starting with the 100th row | [default to &quot;30&quot;]
 **offset** | **string** | Used for offset pagination. Limit/Offset Paging would look like GET /request?limit&#x3D;20&amp;offset&#x3D;100. This query would return the 20 rows starting with the 100th row | [default to &quot;0&quot;]
 **status** | **string** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStats

> *os.File GetStats(ctx).Dayssince(dayssince).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dayssince := "dayssince_example" // string |  (optional) (default to "30")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnclaveServiceApi.GetStats(context.Background()).Dayssince(dayssince).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.GetStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStats`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.GetStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dayssince** | **string** |  | [default to &quot;30&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KillEnclave

> *os.File KillEnclave(ctx, uuid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnclaveServiceApi.KillEnclave(context.Background(), uuid).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `EnclaveServiceApi.KillEnclave``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KillEnclave`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `EnclaveServiceApi.KillEnclave`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKillEnclaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

