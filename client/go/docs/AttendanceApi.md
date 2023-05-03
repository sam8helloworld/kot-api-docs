# \AttendanceApi

All URIs are relative to *https://api.kingtime.jp/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDailyWorking**](AttendanceApi.md#GetDailyWorking) | **Get** /daily-workings/{date} | 
[**GetDailyWorkings**](AttendanceApi.md#GetDailyWorkings) | **Get** /daily-workings | 
[**RegisterDailyWorkingTimerecord**](AttendanceApi.md#RegisterDailyWorkingTimerecord) | **Post** /daily-workings/timerecord/{employeeKey} | 



## GetDailyWorking

> DailyWorkingResponse GetDailyWorking(ctx, date).Division(division).Ondivision(ondivision).AdditionalFields(additionalFields).Execute()





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
    date := time.Now() // string | 詳細を取得したい年月日 ・過去日は最大3年前まで ・未来日は最大1年後まで
    division := "1000" // string | 所属コード (optional)
    ondivision := true // bool | ・true:所属に基づいた勤務データ ・false:出勤先に基づいた勤務データ ・divisionが指定されている場合のみ使用可能 (optional) (default to true)
    additionalFields := []string{"AdditionalFields_example"} // []string | 指定されたプロパティをレスポンスに追加 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttendanceApi.GetDailyWorking(context.Background(), date).Division(division).Ondivision(ondivision).AdditionalFields(additionalFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttendanceApi.GetDailyWorking``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDailyWorking`: DailyWorkingResponse
    fmt.Fprintf(os.Stdout, "Response from `AttendanceApi.GetDailyWorking`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**date** | **string** | 詳細を取得したい年月日 ・過去日は最大3年前まで ・未来日は最大1年後まで | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDailyWorkingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **division** | **string** | 所属コード | 
 **ondivision** | **bool** | ・true:所属に基づいた勤務データ ・false:出勤先に基づいた勤務データ ・divisionが指定されている場合のみ使用可能 | [default to true]
 **additionalFields** | **[]string** | 指定されたプロパティをレスポンスに追加 | 

### Return type

[**DailyWorkingResponse**](DailyWorkingResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDailyWorkings

> []DailyWorkingResponse GetDailyWorkings(ctx).Division(division).Ondivision(ondivision).Start(start).End(end).AdditionalFields(additionalFields).Execute()





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
    division := "1000" // string | 所属コード (optional)
    ondivision := true // bool | ・true:所属に基づいた勤務データ ・false:出勤先に基づいた勤務データ ・divisionが指定されている場合のみ使用可能 (optional) (default to true)
    start := time.Now() // string | 取得したい期間の開始年月日 ・過去日は最大3年前まで (optional)
    end := time.Now() // string | 取得したい期間の終了年月日 ・startとendは同時に指定 ・期間は最大62日 ・未来日は最大1年後まで (optional)
    additionalFields := []string{"AdditionalFields_example"} // []string | 指定されたプロパティをレスポンスに追加 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttendanceApi.GetDailyWorkings(context.Background()).Division(division).Ondivision(ondivision).Start(start).End(end).AdditionalFields(additionalFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttendanceApi.GetDailyWorkings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDailyWorkings`: []DailyWorkingResponse
    fmt.Fprintf(os.Stdout, "Response from `AttendanceApi.GetDailyWorkings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDailyWorkingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **division** | **string** | 所属コード | 
 **ondivision** | **bool** | ・true:所属に基づいた勤務データ ・false:出勤先に基づいた勤務データ ・divisionが指定されている場合のみ使用可能 | [default to true]
 **start** | **string** | 取得したい期間の開始年月日 ・過去日は最大3年前まで | 
 **end** | **string** | 取得したい期間の終了年月日 ・startとendは同時に指定 ・期間は最大62日 ・未来日は最大1年後まで | 
 **additionalFields** | **[]string** | 指定されたプロパティをレスポンスに追加 | 

### Return type

[**[]DailyWorkingResponse**](DailyWorkingResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterDailyWorkingTimerecord

> DailyWorkingTimerecordResponse RegisterDailyWorkingTimerecord(ctx, employeeKey).DailyWorkingTimerecordRequest(dailyWorkingTimerecordRequest).Execute()





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
    dailyWorkingTimerecordRequest := *openapiclient.NewDailyWorkingTimerecordRequest(time.Now(), time.Now()) // DailyWorkingTimerecordRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttendanceApi.RegisterDailyWorkingTimerecord(context.Background(), employeeKey).DailyWorkingTimerecordRequest(dailyWorkingTimerecordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttendanceApi.RegisterDailyWorkingTimerecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterDailyWorkingTimerecord`: DailyWorkingTimerecordResponse
    fmt.Fprintf(os.Stdout, "Response from `AttendanceApi.RegisterDailyWorkingTimerecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employeeKey** | **string** | 従業員識別キー（従業員コードが変更されても不変） | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterDailyWorkingTimerecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dailyWorkingTimerecordRequest** | [**DailyWorkingTimerecordRequest**](DailyWorkingTimerecordRequest.md) |  | 

### Return type

[**DailyWorkingTimerecordResponse**](DailyWorkingTimerecordResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

