# GetSchedules200ResponseScheduleRequestsInnerRequested

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClockInSchedule** | **time.Time** | 出勤予定時間 | 
**ClockOutSchedule** | **time.Time** | 退勤予定時間 | 
**BreakSchedule** | **int32** | 休憩予定時間（分） | 
**StartEndBreakSchedule** | [**[]GetSchedules200ResponseScheduleRequestsInnerRequestedStartEndBreakScheduleInner**](GetSchedules200ResponseScheduleRequestsInnerRequestedStartEndBreakScheduleInner.md) | 休憩開始終了予定 | 
**HolidaysObtained** | [**GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtained**](GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtained.md) |  | 
**WorkDayTypeName** | **string** | 勤務日種別名 | 
**OvertimeUpperLimit** | Pointer to **int32** | 残業上限時間（分） | [optional] 
**WorkPlaceDivisionCode** | **string** | 出勤先所属コード | 
**WorkPlaceDivisionName** | **string** | 出勤先所属名 | 
**SubstitutionClockInName** | **string** | 振替出勤 | 
**ScheduleTypeName** | **string** | スケジュール種別 | 

## Methods

### NewGetSchedules200ResponseScheduleRequestsInnerRequested

`func NewGetSchedules200ResponseScheduleRequestsInnerRequested(clockInSchedule time.Time, clockOutSchedule time.Time, breakSchedule int32, startEndBreakSchedule []GetSchedules200ResponseScheduleRequestsInnerRequestedStartEndBreakScheduleInner, holidaysObtained GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtained, workDayTypeName string, workPlaceDivisionCode string, workPlaceDivisionName string, substitutionClockInName string, scheduleTypeName string, ) *GetSchedules200ResponseScheduleRequestsInnerRequested`

NewGetSchedules200ResponseScheduleRequestsInnerRequested instantiates a new GetSchedules200ResponseScheduleRequestsInnerRequested object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSchedules200ResponseScheduleRequestsInnerRequestedWithDefaults

`func NewGetSchedules200ResponseScheduleRequestsInnerRequestedWithDefaults() *GetSchedules200ResponseScheduleRequestsInnerRequested`

NewGetSchedules200ResponseScheduleRequestsInnerRequestedWithDefaults instantiates a new GetSchedules200ResponseScheduleRequestsInnerRequested object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClockInSchedule

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetClockInSchedule() time.Time`

GetClockInSchedule returns the ClockInSchedule field if non-nil, zero value otherwise.

### GetClockInScheduleOk

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetClockInScheduleOk() (*time.Time, bool)`

GetClockInScheduleOk returns a tuple with the ClockInSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockInSchedule

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) SetClockInSchedule(v time.Time)`

SetClockInSchedule sets ClockInSchedule field to given value.


### GetClockOutSchedule

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetClockOutSchedule() time.Time`

GetClockOutSchedule returns the ClockOutSchedule field if non-nil, zero value otherwise.

### GetClockOutScheduleOk

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetClockOutScheduleOk() (*time.Time, bool)`

GetClockOutScheduleOk returns a tuple with the ClockOutSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockOutSchedule

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) SetClockOutSchedule(v time.Time)`

SetClockOutSchedule sets ClockOutSchedule field to given value.


### GetBreakSchedule

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetBreakSchedule() int32`

GetBreakSchedule returns the BreakSchedule field if non-nil, zero value otherwise.

### GetBreakScheduleOk

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetBreakScheduleOk() (*int32, bool)`

GetBreakScheduleOk returns a tuple with the BreakSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakSchedule

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) SetBreakSchedule(v int32)`

SetBreakSchedule sets BreakSchedule field to given value.


### GetStartEndBreakSchedule

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetStartEndBreakSchedule() []GetSchedules200ResponseScheduleRequestsInnerRequestedStartEndBreakScheduleInner`

GetStartEndBreakSchedule returns the StartEndBreakSchedule field if non-nil, zero value otherwise.

### GetStartEndBreakScheduleOk

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetStartEndBreakScheduleOk() (*[]GetSchedules200ResponseScheduleRequestsInnerRequestedStartEndBreakScheduleInner, bool)`

GetStartEndBreakScheduleOk returns a tuple with the StartEndBreakSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartEndBreakSchedule

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) SetStartEndBreakSchedule(v []GetSchedules200ResponseScheduleRequestsInnerRequestedStartEndBreakScheduleInner)`

SetStartEndBreakSchedule sets StartEndBreakSchedule field to given value.


### GetHolidaysObtained

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetHolidaysObtained() GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtained`

GetHolidaysObtained returns the HolidaysObtained field if non-nil, zero value otherwise.

### GetHolidaysObtainedOk

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetHolidaysObtainedOk() (*GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtained, bool)`

GetHolidaysObtainedOk returns a tuple with the HolidaysObtained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidaysObtained

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) SetHolidaysObtained(v GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtained)`

SetHolidaysObtained sets HolidaysObtained field to given value.


### GetWorkDayTypeName

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetWorkDayTypeName() string`

GetWorkDayTypeName returns the WorkDayTypeName field if non-nil, zero value otherwise.

### GetWorkDayTypeNameOk

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetWorkDayTypeNameOk() (*string, bool)`

GetWorkDayTypeNameOk returns a tuple with the WorkDayTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkDayTypeName

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) SetWorkDayTypeName(v string)`

SetWorkDayTypeName sets WorkDayTypeName field to given value.


### GetOvertimeUpperLimit

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetOvertimeUpperLimit() int32`

GetOvertimeUpperLimit returns the OvertimeUpperLimit field if non-nil, zero value otherwise.

### GetOvertimeUpperLimitOk

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetOvertimeUpperLimitOk() (*int32, bool)`

GetOvertimeUpperLimitOk returns a tuple with the OvertimeUpperLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvertimeUpperLimit

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) SetOvertimeUpperLimit(v int32)`

SetOvertimeUpperLimit sets OvertimeUpperLimit field to given value.

### HasOvertimeUpperLimit

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) HasOvertimeUpperLimit() bool`

HasOvertimeUpperLimit returns a boolean if a field has been set.

### GetWorkPlaceDivisionCode

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetWorkPlaceDivisionCode() string`

GetWorkPlaceDivisionCode returns the WorkPlaceDivisionCode field if non-nil, zero value otherwise.

### GetWorkPlaceDivisionCodeOk

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetWorkPlaceDivisionCodeOk() (*string, bool)`

GetWorkPlaceDivisionCodeOk returns a tuple with the WorkPlaceDivisionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkPlaceDivisionCode

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) SetWorkPlaceDivisionCode(v string)`

SetWorkPlaceDivisionCode sets WorkPlaceDivisionCode field to given value.


### GetWorkPlaceDivisionName

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetWorkPlaceDivisionName() string`

GetWorkPlaceDivisionName returns the WorkPlaceDivisionName field if non-nil, zero value otherwise.

### GetWorkPlaceDivisionNameOk

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetWorkPlaceDivisionNameOk() (*string, bool)`

GetWorkPlaceDivisionNameOk returns a tuple with the WorkPlaceDivisionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkPlaceDivisionName

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) SetWorkPlaceDivisionName(v string)`

SetWorkPlaceDivisionName sets WorkPlaceDivisionName field to given value.


### GetSubstitutionClockInName

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetSubstitutionClockInName() string`

GetSubstitutionClockInName returns the SubstitutionClockInName field if non-nil, zero value otherwise.

### GetSubstitutionClockInNameOk

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetSubstitutionClockInNameOk() (*string, bool)`

GetSubstitutionClockInNameOk returns a tuple with the SubstitutionClockInName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstitutionClockInName

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) SetSubstitutionClockInName(v string)`

SetSubstitutionClockInName sets SubstitutionClockInName field to given value.


### GetScheduleTypeName

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetScheduleTypeName() string`

GetScheduleTypeName returns the ScheduleTypeName field if non-nil, zero value otherwise.

### GetScheduleTypeNameOk

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) GetScheduleTypeNameOk() (*string, bool)`

GetScheduleTypeNameOk returns a tuple with the ScheduleTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTypeName

`func (o *GetSchedules200ResponseScheduleRequestsInnerRequested) SetScheduleTypeName(v string)`

SetScheduleTypeName sets ScheduleTypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


