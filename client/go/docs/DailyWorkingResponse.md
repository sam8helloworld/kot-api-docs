# DailyWorkingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | 日時 | 
**EmployeeKey** | **string** | 従業員識別キー（従業員コードが変更されても不変） | 
**WorkPlaceDivisionCode** | **string** | 出勤先所属コード | 
**WorkPlaceDivisionName** | Pointer to **string** | 出勤先所属名 | [optional] 
**IsClosing** | **bool** | 締め状況 | 
**IsHelp** | **bool** | ヘルプ勤務状況 | 
**IsError** | **bool** | エラー勤務状況 | 
**WorkdayTypeName** | **string** | 勤務日種別名 | 
**Assigned** | **int32** | 所定時間（分） | 
**Unassigned** | **int32** | 所定外時間（分） | 
**Overtime** | **int32** | 残業時間（分） | 
**LateNight** | **int32** | 深夜時間（分） | 
**LateNightUnassigned** | **int32** | 深夜所定外時間（分） | 
**LateNightOvertime** | **int32** | 深夜残業時間（分） | 
**BreakTime** | **int32** | 休憩時間（分） | 
**Late** | **int32** | 遅刻時間（分） | 
**EarlyLeave** | **int32** | 早退時間（分） | 
**TotalWork** | **int32** | 労働合計時間（分） | 
**HolidaysObtained** | [**DailyWorkingResponseHolidaysObtained**](DailyWorkingResponseHolidaysObtained.md) |  | 
**AutoBreakOff** | **NullableInt32** | 自動休憩無効（null： 休憩を無効化しない 1：　雇用区分休憩無効　2： スケジュール休憩無効　3： 全ての自動休憩無効） | 
**DiscretionaryVacation** | **int32** | 休暇みなし時間（分） | 
**CustomDailyWorkings** | [**DailyWorkingResponseCustomDailyWorkings**](DailyWorkingResponseCustomDailyWorkings.md) |  | 
**CurrentDateEmployee** | Pointer to [**DailyWorkingResponseCurrentDateEmployee**](DailyWorkingResponseCurrentDateEmployee.md) |  | [optional] 

## Methods

### NewDailyWorkingResponse

`func NewDailyWorkingResponse(date string, employeeKey string, workPlaceDivisionCode string, isClosing bool, isHelp bool, isError bool, workdayTypeName string, assigned int32, unassigned int32, overtime int32, lateNight int32, lateNightUnassigned int32, lateNightOvertime int32, breakTime int32, late int32, earlyLeave int32, totalWork int32, holidaysObtained DailyWorkingResponseHolidaysObtained, autoBreakOff NullableInt32, discretionaryVacation int32, customDailyWorkings DailyWorkingResponseCustomDailyWorkings, ) *DailyWorkingResponse`

NewDailyWorkingResponse instantiates a new DailyWorkingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyWorkingResponseWithDefaults

`func NewDailyWorkingResponseWithDefaults() *DailyWorkingResponse`

NewDailyWorkingResponseWithDefaults instantiates a new DailyWorkingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *DailyWorkingResponse) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DailyWorkingResponse) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DailyWorkingResponse) SetDate(v string)`

SetDate sets Date field to given value.


### GetEmployeeKey

`func (o *DailyWorkingResponse) GetEmployeeKey() string`

GetEmployeeKey returns the EmployeeKey field if non-nil, zero value otherwise.

### GetEmployeeKeyOk

`func (o *DailyWorkingResponse) GetEmployeeKeyOk() (*string, bool)`

GetEmployeeKeyOk returns a tuple with the EmployeeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeKey

`func (o *DailyWorkingResponse) SetEmployeeKey(v string)`

SetEmployeeKey sets EmployeeKey field to given value.


### GetWorkPlaceDivisionCode

`func (o *DailyWorkingResponse) GetWorkPlaceDivisionCode() string`

GetWorkPlaceDivisionCode returns the WorkPlaceDivisionCode field if non-nil, zero value otherwise.

### GetWorkPlaceDivisionCodeOk

`func (o *DailyWorkingResponse) GetWorkPlaceDivisionCodeOk() (*string, bool)`

GetWorkPlaceDivisionCodeOk returns a tuple with the WorkPlaceDivisionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkPlaceDivisionCode

`func (o *DailyWorkingResponse) SetWorkPlaceDivisionCode(v string)`

SetWorkPlaceDivisionCode sets WorkPlaceDivisionCode field to given value.


### GetWorkPlaceDivisionName

`func (o *DailyWorkingResponse) GetWorkPlaceDivisionName() string`

GetWorkPlaceDivisionName returns the WorkPlaceDivisionName field if non-nil, zero value otherwise.

### GetWorkPlaceDivisionNameOk

`func (o *DailyWorkingResponse) GetWorkPlaceDivisionNameOk() (*string, bool)`

GetWorkPlaceDivisionNameOk returns a tuple with the WorkPlaceDivisionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkPlaceDivisionName

`func (o *DailyWorkingResponse) SetWorkPlaceDivisionName(v string)`

SetWorkPlaceDivisionName sets WorkPlaceDivisionName field to given value.

### HasWorkPlaceDivisionName

`func (o *DailyWorkingResponse) HasWorkPlaceDivisionName() bool`

HasWorkPlaceDivisionName returns a boolean if a field has been set.

### GetIsClosing

`func (o *DailyWorkingResponse) GetIsClosing() bool`

GetIsClosing returns the IsClosing field if non-nil, zero value otherwise.

### GetIsClosingOk

`func (o *DailyWorkingResponse) GetIsClosingOk() (*bool, bool)`

GetIsClosingOk returns a tuple with the IsClosing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosing

`func (o *DailyWorkingResponse) SetIsClosing(v bool)`

SetIsClosing sets IsClosing field to given value.


### GetIsHelp

`func (o *DailyWorkingResponse) GetIsHelp() bool`

GetIsHelp returns the IsHelp field if non-nil, zero value otherwise.

### GetIsHelpOk

`func (o *DailyWorkingResponse) GetIsHelpOk() (*bool, bool)`

GetIsHelpOk returns a tuple with the IsHelp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHelp

`func (o *DailyWorkingResponse) SetIsHelp(v bool)`

SetIsHelp sets IsHelp field to given value.


### GetIsError

`func (o *DailyWorkingResponse) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *DailyWorkingResponse) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *DailyWorkingResponse) SetIsError(v bool)`

SetIsError sets IsError field to given value.


### GetWorkdayTypeName

`func (o *DailyWorkingResponse) GetWorkdayTypeName() string`

GetWorkdayTypeName returns the WorkdayTypeName field if non-nil, zero value otherwise.

### GetWorkdayTypeNameOk

`func (o *DailyWorkingResponse) GetWorkdayTypeNameOk() (*string, bool)`

GetWorkdayTypeNameOk returns a tuple with the WorkdayTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkdayTypeName

`func (o *DailyWorkingResponse) SetWorkdayTypeName(v string)`

SetWorkdayTypeName sets WorkdayTypeName field to given value.


### GetAssigned

`func (o *DailyWorkingResponse) GetAssigned() int32`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *DailyWorkingResponse) GetAssignedOk() (*int32, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *DailyWorkingResponse) SetAssigned(v int32)`

SetAssigned sets Assigned field to given value.


### GetUnassigned

`func (o *DailyWorkingResponse) GetUnassigned() int32`

GetUnassigned returns the Unassigned field if non-nil, zero value otherwise.

### GetUnassignedOk

`func (o *DailyWorkingResponse) GetUnassignedOk() (*int32, bool)`

GetUnassignedOk returns a tuple with the Unassigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassigned

`func (o *DailyWorkingResponse) SetUnassigned(v int32)`

SetUnassigned sets Unassigned field to given value.


### GetOvertime

`func (o *DailyWorkingResponse) GetOvertime() int32`

GetOvertime returns the Overtime field if non-nil, zero value otherwise.

### GetOvertimeOk

`func (o *DailyWorkingResponse) GetOvertimeOk() (*int32, bool)`

GetOvertimeOk returns a tuple with the Overtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvertime

`func (o *DailyWorkingResponse) SetOvertime(v int32)`

SetOvertime sets Overtime field to given value.


### GetLateNight

`func (o *DailyWorkingResponse) GetLateNight() int32`

GetLateNight returns the LateNight field if non-nil, zero value otherwise.

### GetLateNightOk

`func (o *DailyWorkingResponse) GetLateNightOk() (*int32, bool)`

GetLateNightOk returns a tuple with the LateNight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLateNight

`func (o *DailyWorkingResponse) SetLateNight(v int32)`

SetLateNight sets LateNight field to given value.


### GetLateNightUnassigned

`func (o *DailyWorkingResponse) GetLateNightUnassigned() int32`

GetLateNightUnassigned returns the LateNightUnassigned field if non-nil, zero value otherwise.

### GetLateNightUnassignedOk

`func (o *DailyWorkingResponse) GetLateNightUnassignedOk() (*int32, bool)`

GetLateNightUnassignedOk returns a tuple with the LateNightUnassigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLateNightUnassigned

`func (o *DailyWorkingResponse) SetLateNightUnassigned(v int32)`

SetLateNightUnassigned sets LateNightUnassigned field to given value.


### GetLateNightOvertime

`func (o *DailyWorkingResponse) GetLateNightOvertime() int32`

GetLateNightOvertime returns the LateNightOvertime field if non-nil, zero value otherwise.

### GetLateNightOvertimeOk

`func (o *DailyWorkingResponse) GetLateNightOvertimeOk() (*int32, bool)`

GetLateNightOvertimeOk returns a tuple with the LateNightOvertime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLateNightOvertime

`func (o *DailyWorkingResponse) SetLateNightOvertime(v int32)`

SetLateNightOvertime sets LateNightOvertime field to given value.


### GetBreakTime

`func (o *DailyWorkingResponse) GetBreakTime() int32`

GetBreakTime returns the BreakTime field if non-nil, zero value otherwise.

### GetBreakTimeOk

`func (o *DailyWorkingResponse) GetBreakTimeOk() (*int32, bool)`

GetBreakTimeOk returns a tuple with the BreakTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakTime

`func (o *DailyWorkingResponse) SetBreakTime(v int32)`

SetBreakTime sets BreakTime field to given value.


### GetLate

`func (o *DailyWorkingResponse) GetLate() int32`

GetLate returns the Late field if non-nil, zero value otherwise.

### GetLateOk

`func (o *DailyWorkingResponse) GetLateOk() (*int32, bool)`

GetLateOk returns a tuple with the Late field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLate

`func (o *DailyWorkingResponse) SetLate(v int32)`

SetLate sets Late field to given value.


### GetEarlyLeave

`func (o *DailyWorkingResponse) GetEarlyLeave() int32`

GetEarlyLeave returns the EarlyLeave field if non-nil, zero value otherwise.

### GetEarlyLeaveOk

`func (o *DailyWorkingResponse) GetEarlyLeaveOk() (*int32, bool)`

GetEarlyLeaveOk returns a tuple with the EarlyLeave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyLeave

`func (o *DailyWorkingResponse) SetEarlyLeave(v int32)`

SetEarlyLeave sets EarlyLeave field to given value.


### GetTotalWork

`func (o *DailyWorkingResponse) GetTotalWork() int32`

GetTotalWork returns the TotalWork field if non-nil, zero value otherwise.

### GetTotalWorkOk

`func (o *DailyWorkingResponse) GetTotalWorkOk() (*int32, bool)`

GetTotalWorkOk returns a tuple with the TotalWork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWork

`func (o *DailyWorkingResponse) SetTotalWork(v int32)`

SetTotalWork sets TotalWork field to given value.


### GetHolidaysObtained

`func (o *DailyWorkingResponse) GetHolidaysObtained() DailyWorkingResponseHolidaysObtained`

GetHolidaysObtained returns the HolidaysObtained field if non-nil, zero value otherwise.

### GetHolidaysObtainedOk

`func (o *DailyWorkingResponse) GetHolidaysObtainedOk() (*DailyWorkingResponseHolidaysObtained, bool)`

GetHolidaysObtainedOk returns a tuple with the HolidaysObtained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidaysObtained

`func (o *DailyWorkingResponse) SetHolidaysObtained(v DailyWorkingResponseHolidaysObtained)`

SetHolidaysObtained sets HolidaysObtained field to given value.


### GetAutoBreakOff

`func (o *DailyWorkingResponse) GetAutoBreakOff() int32`

GetAutoBreakOff returns the AutoBreakOff field if non-nil, zero value otherwise.

### GetAutoBreakOffOk

`func (o *DailyWorkingResponse) GetAutoBreakOffOk() (*int32, bool)`

GetAutoBreakOffOk returns a tuple with the AutoBreakOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBreakOff

`func (o *DailyWorkingResponse) SetAutoBreakOff(v int32)`

SetAutoBreakOff sets AutoBreakOff field to given value.


### SetAutoBreakOffNil

`func (o *DailyWorkingResponse) SetAutoBreakOffNil(b bool)`

 SetAutoBreakOffNil sets the value for AutoBreakOff to be an explicit nil

### UnsetAutoBreakOff
`func (o *DailyWorkingResponse) UnsetAutoBreakOff()`

UnsetAutoBreakOff ensures that no value is present for AutoBreakOff, not even an explicit nil
### GetDiscretionaryVacation

`func (o *DailyWorkingResponse) GetDiscretionaryVacation() int32`

GetDiscretionaryVacation returns the DiscretionaryVacation field if non-nil, zero value otherwise.

### GetDiscretionaryVacationOk

`func (o *DailyWorkingResponse) GetDiscretionaryVacationOk() (*int32, bool)`

GetDiscretionaryVacationOk returns a tuple with the DiscretionaryVacation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscretionaryVacation

`func (o *DailyWorkingResponse) SetDiscretionaryVacation(v int32)`

SetDiscretionaryVacation sets DiscretionaryVacation field to given value.


### GetCustomDailyWorkings

`func (o *DailyWorkingResponse) GetCustomDailyWorkings() DailyWorkingResponseCustomDailyWorkings`

GetCustomDailyWorkings returns the CustomDailyWorkings field if non-nil, zero value otherwise.

### GetCustomDailyWorkingsOk

`func (o *DailyWorkingResponse) GetCustomDailyWorkingsOk() (*DailyWorkingResponseCustomDailyWorkings, bool)`

GetCustomDailyWorkingsOk returns a tuple with the CustomDailyWorkings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDailyWorkings

`func (o *DailyWorkingResponse) SetCustomDailyWorkings(v DailyWorkingResponseCustomDailyWorkings)`

SetCustomDailyWorkings sets CustomDailyWorkings field to given value.


### GetCurrentDateEmployee

`func (o *DailyWorkingResponse) GetCurrentDateEmployee() DailyWorkingResponseCurrentDateEmployee`

GetCurrentDateEmployee returns the CurrentDateEmployee field if non-nil, zero value otherwise.

### GetCurrentDateEmployeeOk

`func (o *DailyWorkingResponse) GetCurrentDateEmployeeOk() (*DailyWorkingResponseCurrentDateEmployee, bool)`

GetCurrentDateEmployeeOk returns a tuple with the CurrentDateEmployee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDateEmployee

`func (o *DailyWorkingResponse) SetCurrentDateEmployee(v DailyWorkingResponseCurrentDateEmployee)`

SetCurrentDateEmployee sets CurrentDateEmployee field to given value.

### HasCurrentDateEmployee

`func (o *DailyWorkingResponse) HasCurrentDateEmployee() bool`

HasCurrentDateEmployee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


