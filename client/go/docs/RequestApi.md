# \RequestApi

All URIs are relative to *https://api.kingtime.jp/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSchedules**](RequestApi.md#GetSchedules) | **Get** /requests/schedules/{date} | /requests/schedules/{date}{?administratorKey,additionalFields}



## GetSchedules

> GetSchedules200Response GetSchedules(ctx, date).AdministratorKey(administratorKey).AdditionalFields(additionalFields).Execute()

/requests/schedules/{date}{?administratorKey,additionalFields}



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
    date := "2018-08" // string | 詳細を取得したい年月 ・過去日は最大3年前まで ・未来日は最大1年後まで
    administratorKey := "c77a34b32f5de30b6335d141ad714baf6713cd21ca98689efec9fe273526fac222" // string | 管理者識別キー（管理者コードが変更されても不変） (optional)
    additionalFields := []string{"flow"} // []string | 指定されたプロパティをレスポンスに追加 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestApi.GetSchedules(context.Background(), date).AdministratorKey(administratorKey).AdditionalFields(additionalFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestApi.GetSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchedules`: GetSchedules200Response
    fmt.Fprintf(os.Stdout, "Response from `RequestApi.GetSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**date** | **string** | 詳細を取得したい年月 ・過去日は最大3年前まで ・未来日は最大1年後まで | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **administratorKey** | **string** | 管理者識別キー（管理者コードが変更されても不変） | 
 **additionalFields** | **[]string** | 指定されたプロパティをレスポンスに追加 | 

### Return type

[**GetSchedules200Response**](GetSchedules200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

