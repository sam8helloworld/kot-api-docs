# DailyWorkingTimerecordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | 日時 | 
**EmployeeKey** | **string** | 従業員識別キー（従業員コードが変更されても不変） | 
**TimeRecord** | [**[]DailyWorkingTimerecordResponseTimeRecordInner**](DailyWorkingTimerecordResponseTimeRecordInner.md) | 打刻 | 
**CurrentDateEmployee** | Pointer to [**DailyWorkingResponseCurrentDateEmployee**](DailyWorkingResponseCurrentDateEmployee.md) |  | [optional] 

## Methods

### NewDailyWorkingTimerecordResponse

`func NewDailyWorkingTimerecordResponse(date string, employeeKey string, timeRecord []DailyWorkingTimerecordResponseTimeRecordInner, ) *DailyWorkingTimerecordResponse`

NewDailyWorkingTimerecordResponse instantiates a new DailyWorkingTimerecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyWorkingTimerecordResponseWithDefaults

`func NewDailyWorkingTimerecordResponseWithDefaults() *DailyWorkingTimerecordResponse`

NewDailyWorkingTimerecordResponseWithDefaults instantiates a new DailyWorkingTimerecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *DailyWorkingTimerecordResponse) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DailyWorkingTimerecordResponse) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DailyWorkingTimerecordResponse) SetDate(v string)`

SetDate sets Date field to given value.


### GetEmployeeKey

`func (o *DailyWorkingTimerecordResponse) GetEmployeeKey() string`

GetEmployeeKey returns the EmployeeKey field if non-nil, zero value otherwise.

### GetEmployeeKeyOk

`func (o *DailyWorkingTimerecordResponse) GetEmployeeKeyOk() (*string, bool)`

GetEmployeeKeyOk returns a tuple with the EmployeeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeKey

`func (o *DailyWorkingTimerecordResponse) SetEmployeeKey(v string)`

SetEmployeeKey sets EmployeeKey field to given value.


### GetTimeRecord

`func (o *DailyWorkingTimerecordResponse) GetTimeRecord() []DailyWorkingTimerecordResponseTimeRecordInner`

GetTimeRecord returns the TimeRecord field if non-nil, zero value otherwise.

### GetTimeRecordOk

`func (o *DailyWorkingTimerecordResponse) GetTimeRecordOk() (*[]DailyWorkingTimerecordResponseTimeRecordInner, bool)`

GetTimeRecordOk returns a tuple with the TimeRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRecord

`func (o *DailyWorkingTimerecordResponse) SetTimeRecord(v []DailyWorkingTimerecordResponseTimeRecordInner)`

SetTimeRecord sets TimeRecord field to given value.


### GetCurrentDateEmployee

`func (o *DailyWorkingTimerecordResponse) GetCurrentDateEmployee() DailyWorkingResponseCurrentDateEmployee`

GetCurrentDateEmployee returns the CurrentDateEmployee field if non-nil, zero value otherwise.

### GetCurrentDateEmployeeOk

`func (o *DailyWorkingTimerecordResponse) GetCurrentDateEmployeeOk() (*DailyWorkingResponseCurrentDateEmployee, bool)`

GetCurrentDateEmployeeOk returns a tuple with the CurrentDateEmployee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDateEmployee

`func (o *DailyWorkingTimerecordResponse) SetCurrentDateEmployee(v DailyWorkingResponseCurrentDateEmployee)`

SetCurrentDateEmployee sets CurrentDateEmployee field to given value.

### HasCurrentDateEmployee

`func (o *DailyWorkingTimerecordResponse) HasCurrentDateEmployee() bool`

HasCurrentDateEmployee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


