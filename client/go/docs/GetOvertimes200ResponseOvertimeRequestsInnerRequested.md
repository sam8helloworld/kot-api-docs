# GetOvertimes200ResponseOvertimeRequestsInnerRequested

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsBeforeSchedule** | **bool** | 勤務予定前の時間外申請か（true 予定前　false： 予定後） | 
**Start** | **time.Time** | 開始時刻 | 
**End** | **time.Time** | 終了時刻 | 

## Methods

### NewGetOvertimes200ResponseOvertimeRequestsInnerRequested

`func NewGetOvertimes200ResponseOvertimeRequestsInnerRequested(isBeforeSchedule bool, start time.Time, end time.Time, ) *GetOvertimes200ResponseOvertimeRequestsInnerRequested`

NewGetOvertimes200ResponseOvertimeRequestsInnerRequested instantiates a new GetOvertimes200ResponseOvertimeRequestsInnerRequested object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOvertimes200ResponseOvertimeRequestsInnerRequestedWithDefaults

`func NewGetOvertimes200ResponseOvertimeRequestsInnerRequestedWithDefaults() *GetOvertimes200ResponseOvertimeRequestsInnerRequested`

NewGetOvertimes200ResponseOvertimeRequestsInnerRequestedWithDefaults instantiates a new GetOvertimes200ResponseOvertimeRequestsInnerRequested object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsBeforeSchedule

`func (o *GetOvertimes200ResponseOvertimeRequestsInnerRequested) GetIsBeforeSchedule() bool`

GetIsBeforeSchedule returns the IsBeforeSchedule field if non-nil, zero value otherwise.

### GetIsBeforeScheduleOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInnerRequested) GetIsBeforeScheduleOk() (*bool, bool)`

GetIsBeforeScheduleOk returns a tuple with the IsBeforeSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBeforeSchedule

`func (o *GetOvertimes200ResponseOvertimeRequestsInnerRequested) SetIsBeforeSchedule(v bool)`

SetIsBeforeSchedule sets IsBeforeSchedule field to given value.


### GetStart

`func (o *GetOvertimes200ResponseOvertimeRequestsInnerRequested) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInnerRequested) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *GetOvertimes200ResponseOvertimeRequestsInnerRequested) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *GetOvertimes200ResponseOvertimeRequestsInnerRequested) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInnerRequested) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *GetOvertimes200ResponseOvertimeRequestsInnerRequested) SetEnd(v time.Time)`

SetEnd sets End field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


