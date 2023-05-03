# RegisterDailyWorkingTimerecord201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | 日時 | 
**EmployeeKey** | **string** | 従業員識別キー（従業員コードが変更されても不変） | 
**TimeRecord** | [**RegisterDailyWorkingTimerecord201ResponseTimeRecord**](RegisterDailyWorkingTimerecord201ResponseTimeRecord.md) |  | 

## Methods

### NewRegisterDailyWorkingTimerecord201Response

`func NewRegisterDailyWorkingTimerecord201Response(date string, employeeKey string, timeRecord RegisterDailyWorkingTimerecord201ResponseTimeRecord, ) *RegisterDailyWorkingTimerecord201Response`

NewRegisterDailyWorkingTimerecord201Response instantiates a new RegisterDailyWorkingTimerecord201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterDailyWorkingTimerecord201ResponseWithDefaults

`func NewRegisterDailyWorkingTimerecord201ResponseWithDefaults() *RegisterDailyWorkingTimerecord201Response`

NewRegisterDailyWorkingTimerecord201ResponseWithDefaults instantiates a new RegisterDailyWorkingTimerecord201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *RegisterDailyWorkingTimerecord201Response) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *RegisterDailyWorkingTimerecord201Response) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *RegisterDailyWorkingTimerecord201Response) SetDate(v string)`

SetDate sets Date field to given value.


### GetEmployeeKey

`func (o *RegisterDailyWorkingTimerecord201Response) GetEmployeeKey() string`

GetEmployeeKey returns the EmployeeKey field if non-nil, zero value otherwise.

### GetEmployeeKeyOk

`func (o *RegisterDailyWorkingTimerecord201Response) GetEmployeeKeyOk() (*string, bool)`

GetEmployeeKeyOk returns a tuple with the EmployeeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeKey

`func (o *RegisterDailyWorkingTimerecord201Response) SetEmployeeKey(v string)`

SetEmployeeKey sets EmployeeKey field to given value.


### GetTimeRecord

`func (o *RegisterDailyWorkingTimerecord201Response) GetTimeRecord() RegisterDailyWorkingTimerecord201ResponseTimeRecord`

GetTimeRecord returns the TimeRecord field if non-nil, zero value otherwise.

### GetTimeRecordOk

`func (o *RegisterDailyWorkingTimerecord201Response) GetTimeRecordOk() (*RegisterDailyWorkingTimerecord201ResponseTimeRecord, bool)`

GetTimeRecordOk returns a tuple with the TimeRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRecord

`func (o *RegisterDailyWorkingTimerecord201Response) SetTimeRecord(v RegisterDailyWorkingTimerecord201ResponseTimeRecord)`

SetTimeRecord sets TimeRecord field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


