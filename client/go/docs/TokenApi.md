# \TokenApi

All URIs are relative to *https://api.kingtime.jp/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessTokenAvailability**](TokenApi.md#GetAccessTokenAvailability) | **Get** /tokens/{token}/available | 
[**RefreshAccessToken**](TokenApi.md#RefreshAccessToken) | **Post** /tokens/{token} | 
[**SuspendAccessToken**](TokenApi.md#SuspendAccessToken) | **Delete** /tokens/{token} | 



## GetAccessTokenAvailability

> GetAccessTokenAvailability200Response GetAccessTokenAvailability(ctx, token).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    token := "8j9f7v4893y58rvt7nyfq2893n75tr78937n83" // string | 処理対象のトークン

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenApi.GetAccessTokenAvailability(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.GetAccessTokenAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessTokenAvailability`: GetAccessTokenAvailability200Response
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.GetAccessTokenAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | 処理対象のトークン | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessTokenAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAccessTokenAvailability200Response**](GetAccessTokenAvailability200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshAccessToken

> RefreshAccessToken201Response RefreshAccessToken(ctx, token).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    token := "8j9f7v4893y58rvt7nyfq2893n75tr78937n83" // string | 処理対象のトークン

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenApi.RefreshAccessToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.RefreshAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshAccessToken`: RefreshAccessToken201Response
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.RefreshAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | 処理対象のトークン | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RefreshAccessToken201Response**](RefreshAccessToken201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendAccessToken

> SuspendAccessToken(ctx, token).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    token := "8j9f7v4893y58rvt7nyfq2893n75tr78937n83" // string | 処理対象のトークン

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TokenApi.SuspendAccessToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.SuspendAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | 処理対象のトークン | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

