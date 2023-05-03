# GetDailyWorkingTimerecords200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | 日時 | 
**EmployeeKey** | **string** | 従業員識別キー（従業員コードが変更されても不変） | 
**TimeRecord** | [**[]DailyWorkingTimerecord**](DailyWorkingTimerecord.md) | 打刻 | 
**CurrentDateEmployee** | Pointer to [**GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee**](GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee.md) |  | [optional] 

## Methods

### NewGetDailyWorkingTimerecords200ResponseInner

`func NewGetDailyWorkingTimerecords200ResponseInner(date string, employeeKey string, timeRecord []DailyWorkingTimerecord, ) *GetDailyWorkingTimerecords200ResponseInner`

NewGetDailyWorkingTimerecords200ResponseInner instantiates a new GetDailyWorkingTimerecords200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDailyWorkingTimerecords200ResponseInnerWithDefaults

`func NewGetDailyWorkingTimerecords200ResponseInnerWithDefaults() *GetDailyWorkingTimerecords200ResponseInner`

NewGetDailyWorkingTimerecords200ResponseInnerWithDefaults instantiates a new GetDailyWorkingTimerecords200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetDailyWorkingTimerecords200ResponseInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetDailyWorkingTimerecords200ResponseInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetDailyWorkingTimerecords200ResponseInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetEmployeeKey

`func (o *GetDailyWorkingTimerecords200ResponseInner) GetEmployeeKey() string`

GetEmployeeKey returns the EmployeeKey field if non-nil, zero value otherwise.

### GetEmployeeKeyOk

`func (o *GetDailyWorkingTimerecords200ResponseInner) GetEmployeeKeyOk() (*string, bool)`

GetEmployeeKeyOk returns a tuple with the EmployeeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeKey

`func (o *GetDailyWorkingTimerecords200ResponseInner) SetEmployeeKey(v string)`

SetEmployeeKey sets EmployeeKey field to given value.


### GetTimeRecord

`func (o *GetDailyWorkingTimerecords200ResponseInner) GetTimeRecord() []DailyWorkingTimerecord`

GetTimeRecord returns the TimeRecord field if non-nil, zero value otherwise.

### GetTimeRecordOk

`func (o *GetDailyWorkingTimerecords200ResponseInner) GetTimeRecordOk() (*[]DailyWorkingTimerecord, bool)`

GetTimeRecordOk returns a tuple with the TimeRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRecord

`func (o *GetDailyWorkingTimerecords200ResponseInner) SetTimeRecord(v []DailyWorkingTimerecord)`

SetTimeRecord sets TimeRecord field to given value.


### GetCurrentDateEmployee

`func (o *GetDailyWorkingTimerecords200ResponseInner) GetCurrentDateEmployee() GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee`

GetCurrentDateEmployee returns the CurrentDateEmployee field if non-nil, zero value otherwise.

### GetCurrentDateEmployeeOk

`func (o *GetDailyWorkingTimerecords200ResponseInner) GetCurrentDateEmployeeOk() (*GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee, bool)`

GetCurrentDateEmployeeOk returns a tuple with the CurrentDateEmployee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDateEmployee

`func (o *GetDailyWorkingTimerecords200ResponseInner) SetCurrentDateEmployee(v GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee)`

SetCurrentDateEmployee sets CurrentDateEmployee field to given value.

### HasCurrentDateEmployee

`func (o *GetDailyWorkingTimerecords200ResponseInner) HasCurrentDateEmployee() bool`

HasCurrentDateEmployee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


