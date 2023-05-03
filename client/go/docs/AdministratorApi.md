# \AdministratorApi

All URIs are relative to *https://api.kingtime.jp/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdministrators**](AdministratorApi.md#GetAdministrators) | **Get** /administrators | /administrators{?additionalFields}



## GetAdministrators

> []GetAdministrators200ResponseInner GetAdministrators(ctx).AdditionalFields(additionalFields).Execute()

/administrators{?additionalFields}



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sam8helloworld/kot-api-docs"
)

func main() {
    additionalFields := []string{"AdditionalFields_example"} // []string | 指定されたプロパティをレスポンスに追加 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministratorApi.GetAdministrators(context.Background()).AdditionalFields(additionalFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministratorApi.GetAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdministrators`: []GetAdministrators200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `AdministratorApi.GetAdministrators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdministratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **additionalFields** | **[]string** | 指定されたプロパティをレスポンスに追加 | 

### Return type

[**[]GetAdministrators200ResponseInner**](GetAdministrators200ResponseInner.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

