# GetOvertimes200ResponseOvertimeRequestsInnerCurrent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsBeforeSchedule** | **bool** | 勤務予定前の時間外申請か（true 予定前　false： 予定後） | 
**Start** | **time.Time** | 開始時刻 | 
**End** | **time.Time** | 終了時刻 | 

## Methods

### NewGetOvertimes200ResponseOvertimeRequestsInnerCurrent

`func NewGetOvertimes200ResponseOvertimeRequestsInnerCurrent(isBeforeSchedule bool, start time.Time, end time.Time, ) *GetOvertimes200ResponseOvertimeRequestsInnerCurrent`

NewGetOvertimes200ResponseOvertimeRequestsInnerCurrent instantiates a new GetOvertimes200ResponseOvertimeRequestsInnerCurrent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOvertimes200ResponseOvertimeRequestsInnerCurrentWithDefaults

`func NewGetOvertimes200ResponseOvertimeRequestsInnerCurrentWithDefaults() *GetOvertimes200ResponseOvertimeRequestsInnerCurrent`

NewGetOvertimes200ResponseOvertimeRequestsInnerCurrentWithDefaults instantiates a new GetOvertimes200ResponseOvertimeRequestsInnerCurrent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsBeforeSchedule

`func (o *GetOvertimes200ResponseOvertimeRequestsInnerCurrent) GetIsBeforeSchedule() bool`

GetIsBeforeSchedule returns the IsBeforeSchedule field if non-nil, zero value otherwise.

### GetIsBeforeScheduleOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInnerCurrent) GetIsBeforeScheduleOk() (*bool, bool)`

GetIsBeforeScheduleOk returns a tuple with the IsBeforeSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBeforeSchedule

`func (o *GetOvertimes200ResponseOvertimeRequestsInnerCurrent) SetIsBeforeSchedule(v bool)`

SetIsBeforeSchedule sets IsBeforeSchedule field to given value.


### GetStart

`func (o *GetOvertimes200ResponseOvertimeRequestsInnerCurrent) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInnerCurrent) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *GetOvertimes200ResponseOvertimeRequestsInnerCurrent) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *GetOvertimes200ResponseOvertimeRequestsInnerCurrent) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInnerCurrent) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *GetOvertimes200ResponseOvertimeRequestsInnerCurrent) SetEnd(v time.Time)`

SetEnd sets End field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


