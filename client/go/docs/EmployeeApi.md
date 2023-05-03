# \EmployeeApi

All URIs are relative to *https://api.kingtime.jp/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEmployee**](EmployeeApi.md#DeleteEmployee) | **Delete** /employees/{employeeKey} | 
[**GetDivisions**](EmployeeApi.md#GetDivisions) | **Get** /divisions | 
[**GetEmployee**](EmployeeApi.md#GetEmployee) | **Get** /employees/{employeeCode}{?date,includeResigner,additionalFields} | 
[**GetEmployeeGroups**](EmployeeApi.md#GetEmployeeGroups) | **Get** /employee-groups | 
[**GetEmployees**](EmployeeApi.md#GetEmployees) | **Get** /employees{?date,division,includeResigner,additionalFields} | 
[**GetWorkingTypes**](EmployeeApi.md#GetWorkingTypes) | **Get** /working-types | 
[**RegisterEmployee**](EmployeeApi.md#RegisterEmployee) | **Post** /employees | 
[**UpdateEmployee**](EmployeeApi.md#UpdateEmployee) | **Put** /employees/{employeeKey}{?updateDate} | 



## DeleteEmployee

> DeleteEmployee(ctx, employeeKey).Execute()





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
    employeeKey := "8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3" // string | 従業員識別キー（従業員コードが変更されても不変）

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EmployeeApi.DeleteEmployee(context.Background(), employeeKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmployeeApi.DeleteEmployee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employeeKey** | **string** | 従業員識別キー（従業員コードが変更されても不変） | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmployeeRequest struct via the builder pattern


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


## GetDivisions

> []DivisionResponse GetDivisions(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmployeeApi.GetDivisions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmployeeApi.GetDivisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDivisions`: []DivisionResponse
    fmt.Fprintf(os.Stdout, "Response from `EmployeeApi.GetDivisions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDivisionsRequest struct via the builder pattern


### Return type

[**[]DivisionResponse**](DivisionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployee

> EmployeeResponse GetEmployee(ctx, employeeCode).Date(date).IncludeResigner(includeResigner).AdditionalFields(additionalFields).Execute()





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
    employeeCode := "1000" // string | 詳細を取得したい従業員コード
    date := time.Now() // string | 指定された年月日時点での従業員のデータを表示 ・過去日は最大3年前まで ・未来日は最大1年後まで (optional)
    includeResigner := true // bool | 指定された年月日時点で退職済みの従業員を含む場合 True (optional) (default to false)
    additionalFields := []string{"AdditionalFields_example"} // []string | 指定されたプロパティをレスポンスに追加 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmployeeApi.GetEmployee(context.Background(), employeeCode).Date(date).IncludeResigner(includeResigner).AdditionalFields(additionalFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmployeeApi.GetEmployee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmployee`: EmployeeResponse
    fmt.Fprintf(os.Stdout, "Response from `EmployeeApi.GetEmployee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employeeCode** | **string** | 詳細を取得したい従業員コード | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **string** | 指定された年月日時点での従業員のデータを表示 ・過去日は最大3年前まで ・未来日は最大1年後まで | 
 **includeResigner** | **bool** | 指定された年月日時点で退職済みの従業員を含む場合 True | [default to false]
 **additionalFields** | **[]string** | 指定されたプロパティをレスポンスに追加 | 

### Return type

[**EmployeeResponse**](EmployeeResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployeeGroups

> []GetEmployeeGroups200ResponseInner GetEmployeeGroups(ctx).AdditionalFields(additionalFields).Execute()





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
    resp, r, err := apiClient.EmployeeApi.GetEmployeeGroups(context.Background()).AdditionalFields(additionalFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmployeeApi.GetEmployeeGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmployeeGroups`: []GetEmployeeGroups200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `EmployeeApi.GetEmployeeGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployeeGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **additionalFields** | **[]string** | 指定されたプロパティをレスポンスに追加 | 

### Return type

[**[]GetEmployeeGroups200ResponseInner**](GetEmployeeGroups200ResponseInner.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployees

> []EmployeeResponse GetEmployees(ctx).Date(date).Division(division).IncludeResigner(includeResigner).AdditionalFields(additionalFields).Execute()





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
    // response from `GetEmployees`: []EmployeeResponse
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

[**[]EmployeeResponse**](EmployeeResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkingTypes

> []WorkingTypeResponse GetWorkingTypes(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmployeeApi.GetWorkingTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmployeeApi.GetWorkingTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkingTypes`: []WorkingTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `EmployeeApi.GetWorkingTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkingTypesRequest struct via the builder pattern


### Return type

[**[]WorkingTypeResponse**](WorkingTypeResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterEmployee

> RegisterEmployee201Response RegisterEmployee(ctx).EmployeeRequest(employeeRequest).Execute()





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
    employeeRequest := *openapiclient.NewEmployeeRequest("DivisionCode_example", "Gender_example", "TypeCode_example", "Code_example", "LastName_example", "FirstName_example") // EmployeeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmployeeApi.RegisterEmployee(context.Background()).EmployeeRequest(employeeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmployeeApi.RegisterEmployee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterEmployee`: RegisterEmployee201Response
    fmt.Fprintf(os.Stdout, "Response from `EmployeeApi.RegisterEmployee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterEmployeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **employeeRequest** | [**EmployeeRequest**](EmployeeRequest.md) |  | 

### Return type

[**RegisterEmployee201Response**](RegisterEmployee201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmployee

> EmployeeResponse UpdateEmployee(ctx, employeeKey).UpdateDate(updateDate).UpdateEmployeeRequest(updateEmployeeRequest).Execute()





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
    employeeKey := "8b6ee646a9620b286499c3df6918c4888a97dd7bbc6a26a18743f4697a1de4b3" // string | 従業員識別キー（従業員コードが変更されても不変）
    updateDate := time.Now() // string | 所属、雇用区分を更新したい年月日 (optional)
    updateEmployeeRequest := *openapiclient.NewUpdateEmployeeRequest() // UpdateEmployeeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmployeeApi.UpdateEmployee(context.Background(), employeeKey).UpdateDate(updateDate).UpdateEmployeeRequest(updateEmployeeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmployeeApi.UpdateEmployee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEmployee`: EmployeeResponse
    fmt.Fprintf(os.Stdout, "Response from `EmployeeApi.UpdateEmployee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employeeKey** | **string** | 従業員識別キー（従業員コードが変更されても不変） | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmployeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDate** | **string** | 所属、雇用区分を更新したい年月日 | 
 **updateEmployeeRequest** | [**UpdateEmployeeRequest**](UpdateEmployeeRequest.md) |  | 

### Return type

[**EmployeeResponse**](EmployeeResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

