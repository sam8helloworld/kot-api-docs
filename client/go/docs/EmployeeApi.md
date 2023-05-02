# \EmployeeApi

All URIs are relative to *https://api.kingtime.jp/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEmployees**](EmployeeApi.md#GetEmployees) | **Get** /employees | 



## GetEmployees

> []GetEmployees200ResponseInner GetEmployees(ctx).Date(date).Division(division).IncludeResigner(includeResigner).AdditionalFields(additionalFields).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/sam8helloworld/kot-api-docs"
)

func main() {
    date := time.Now() // string | 指定された年月日時点での従業員のデータを表示 ・過去日は最大3年前まで ・未来日は最大1年後まで (optional)
    division := "1000" // string | 所属コード (optional)
    includeResigner := true // bool | 指定された年月日時点で退職済みの従業員を含む場合 True (optional) (default to false)
    additionalFields := []string{"AdditionalFields_example"} // []string | 指定されたプロパティをレスポンスに追加 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmployeeApi.GetEmployees(context.Background()).Date(date).Division(division).IncludeResigner(includeResigner).AdditionalFields(additionalFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmployeeApi.GetEmployees``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmployees`: []GetEmployees200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `EmployeeApi.GetEmployees`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployeesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **string** | 指定された年月日時点での従業員のデータを表示 ・過去日は最大3年前まで ・未来日は最大1年後まで | 
 **division** | **string** | 所属コード | 
 **includeResigner** | **bool** | 指定された年月日時点で退職済みの従業員を含む場合 True | [default to false]
 **additionalFields** | **[]string** | 指定されたプロパティをレスポンスに追加 | 

### Return type

[**[]GetEmployees200ResponseInner**](GetEmployees200ResponseInner.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

