# \CompanyApi

All URIs are relative to *https://api.kingtime.jp/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompany**](CompanyApi.md#GetCompany) | **Get** /company | 



## GetCompany

> GetCompany200Response GetCompany(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompanyApi.GetCompany(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.GetCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompany`: GetCompany200Response
    fmt.Fprintf(os.Stdout, "Response from `CompanyApi.GetCompany`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyRequest struct via the builder pattern


### Return type

[**GetCompany200Response**](GetCompany200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

