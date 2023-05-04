# GetSchedules200ResponseScheduleRequestsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | 対象日 | 
**RequestedDate** | **string** | 申請日 | 
**EmployeeKey** | **string** | 従業員識別キー（従業員コードが変更されても不変） | 
**RequestKey** | **string** | 申請識別キー | 
**Applicant** | [**GetSchedules200ResponseScheduleRequestsInnerApplicant**](GetSchedules200ResponseScheduleRequestsInnerApplicant.md) |  | 
**Status** | **string** | 申請ステータス（applying 申請中　rejected： 棄却　approved： 承認） | 
**CurrentFlow** | **int32** | 現在の承認フローレベル（１～５） | 
**LastModifiedAdministratorKey** | **string** | 最終更新管理者識別キー | 
**Message** | **string** | 申請メッセージ | 
**AdminComment** | **string** | 管理者コメント | 
**Requested** | [**GetSchedules200ResponseScheduleRequestsInnerRequested**](GetSchedules200ResponseScheduleRequestsInnerRequested.md) |  | 
**Current** | [**GetSchedules200ResponseScheduleRequestsInnerRequested**](GetSchedules200ResponseScheduleRequestsInnerRequested.md) |  | 
**SchedulePatternCode** | Pointer to **string** | スケジュールパターンコード | [optional] 
**SchedulePatternName** | Pointer to **string** | スケジュールパターン名 | [optional] 
**WorkFixedStart** | Pointer to **time.Time** | 勤務開始刻限 | [optional] 
**WorkFixedEnd** | Pointer to **time.Time** | 勤務終了刻限 | [optional] 
**Note** | Pointer to **string** | 備考 | [optional] 
**Flow** | Pointer to [**[]GetSchedules200ResponseScheduleRequestsInnerFlowInner**](GetSchedules200ResponseScheduleRequestsInnerFlowInner.md) | 承認フロー | [optional] 

## Methods

### NewGetSchedules200ResponseScheduleRequestsInner

`func NewGetSchedules200ResponseScheduleRequestsInner(date string, requestedDate string, employeeKey string, requestKey string, applicant GetSchedules200ResponseScheduleRequestsInnerApplicant, status string, currentFlow int32, lastModifiedAdministratorKey string, message string, adminComment string, requested GetSchedules200ResponseScheduleRequestsInnerRequested, current GetSchedules200ResponseScheduleRequestsInnerRequested, ) *GetSchedules200ResponseScheduleRequestsInner`

NewGetSchedules200ResponseScheduleRequestsInner instantiates a new GetSchedules200ResponseScheduleRequestsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSchedules200ResponseScheduleRequestsInnerWithDefaults

`func NewGetSchedules200ResponseScheduleRequestsInnerWithDefaults() *GetSchedules200ResponseScheduleRequestsInner`

NewGetSchedules200ResponseScheduleRequestsInnerWithDefaults instantiates a new GetSchedules200ResponseScheduleRequestsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetSchedules200ResponseScheduleRequestsInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetRequestedDate

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetRequestedDate() string`

GetRequestedDate returns the RequestedDate field if non-nil, zero value otherwise.

### GetRequestedDateOk

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetRequestedDateOk() (*string, bool)`

GetRequestedDateOk returns a tuple with the RequestedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDate

`func (o *GetSchedules200ResponseScheduleRequestsInner) SetRequestedDate(v string)`

SetRequestedDate sets RequestedDate field to given value.


### GetEmployeeKey

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetEmployeeKey() string`

GetEmployeeKey returns the EmployeeKey field if non-nil, zero value otherwise.

### GetEmployeeKeyOk

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetEmployeeKeyOk() (*string, bool)`

GetEmployeeKeyOk returns a tuple with the EmployeeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeKey

`func (o *GetSchedules200ResponseScheduleRequestsInner) SetEmployeeKey(v string)`

SetEmployeeKey sets EmployeeKey field to given value.


### GetRequestKey

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetRequestKey() string`

GetRequestKey returns the RequestKey field if non-nil, zero value otherwise.

### GetRequestKeyOk

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetRequestKeyOk() (*string, bool)`

GetRequestKeyOk returns a tuple with the RequestKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestKey

`func (o *GetSchedules200ResponseScheduleRequestsInner) SetRequestKey(v string)`

SetRequestKey sets RequestKey field to given value.


### GetApplicant

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetApplicant() GetSchedules200ResponseScheduleRequestsInnerApplicant`

GetApplicant returns the Applicant field if non-nil, zero value otherwise.

### GetApplicantOk

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetApplicantOk() (*GetSchedules200ResponseScheduleRequestsInnerApplicant, bool)`

GetApplicantOk returns a tuple with the Applicant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicant

`func (o *GetSchedules200ResponseScheduleRequestsInner) SetApplicant(v GetSchedules200ResponseScheduleRequestsInnerApplicant)`

SetApplicant sets Applicant field to given value.


### GetStatus

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSchedules200ResponseScheduleRequestsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCurrentFlow

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetCurrentFlow() int32`

GetCurrentFlow returns the CurrentFlow field if non-nil, zero value otherwise.

### GetCurrentFlowOk

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetCurrentFlowOk() (*int32, bool)`

GetCurrentFlowOk returns a tuple with the CurrentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentFlow

`func (o *GetSchedules200ResponseScheduleRequestsInner) SetCurrentFlow(v int32)`

SetCurrentFlow sets CurrentFlow field to given value.


### GetLastModifiedAdministratorKey

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetLastModifiedAdministratorKey() string`

GetLastModifiedAdministratorKey returns the LastModifiedAdministratorKey field if non-nil, zero value otherwise.

### GetLastModifiedAdministratorKeyOk

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetLastModifiedAdministratorKeyOk() (*string, bool)`

GetLastModifiedAdministratorKeyOk returns a tuple with the LastModifiedAdministratorKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAdministratorKey

`func (o *GetSchedules200ResponseScheduleRequestsInner) SetLastModifiedAdministratorKey(v string)`

SetLastModifiedAdministratorKey sets LastModifiedAdministratorKey field to given value.


### GetMessage

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetSchedules200ResponseScheduleRequestsInner) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAdminComment

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetAdminComment() string`

GetAdminComment returns the AdminComment field if non-nil, zero value otherwise.

### GetAdminCommentOk

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetAdminCommentOk() (*string, bool)`

GetAdminCommentOk returns a tuple with the AdminComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminComment

`func (o *GetSchedules200ResponseScheduleRequestsInner) SetAdminComment(v string)`

SetAdminComment sets AdminComment field to given value.


### GetRequested

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetRequested() GetSchedules200ResponseScheduleRequestsInnerRequested`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetRequestedOk() (*GetSchedules200ResponseScheduleRequestsInnerRequested, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *GetSchedules200ResponseScheduleRequestsInner) SetRequested(v GetSchedules200ResponseScheduleRequestsInnerRequested)`

SetRequested sets Requested field to given value.


### GetCurrent

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetCurrent() GetSchedules200ResponseScheduleRequestsInnerRequested`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetCurrentOk() (*GetSchedules200ResponseScheduleRequestsInnerRequested, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *GetSchedules200ResponseScheduleRequestsInner) SetCurrent(v GetSchedules200ResponseScheduleRequestsInnerRequested)`

SetCurrent sets Current field to given value.


### GetSchedulePatternCode

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetSchedulePatternCode() string`

GetSchedulePatternCode returns the SchedulePatternCode field if non-nil, zero value otherwise.

### GetSchedulePatternCodeOk

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetSchedulePatternCodeOk() (*string, bool)`

GetSchedulePatternCodeOk returns a tuple with the SchedulePatternCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulePatternCode

`func (o *GetSchedules200ResponseScheduleRequestsInner) SetSchedulePatternCode(v string)`

SetSchedulePatternCode sets SchedulePatternCode field to given value.

### HasSchedulePatternCode

`func (o *GetSchedules200ResponseScheduleRequestsInner) HasSchedulePatternCode() bool`

HasSchedulePatternCode returns a boolean if a field has been set.

### GetSchedulePatternName

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetSchedulePatternName() string`

GetSchedulePatternName returns the SchedulePatternName field if non-nil, zero value otherwise.

### GetSchedulePatternNameOk

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetSchedulePatternNameOk() (*string, bool)`

GetSchedulePatternNameOk returns a tuple with the SchedulePatternName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulePatternName

`func (o *GetSchedules200ResponseScheduleRequestsInner) SetSchedulePatternName(v string)`

SetSchedulePatternName sets SchedulePatternName field to given value.

### HasSchedulePatternName

`func (o *GetSchedules200ResponseScheduleRequestsInner) HasSchedulePatternName() bool`

HasSchedulePatternName returns a boolean if a field has been set.

### GetWorkFixedStart

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetWorkFixedStart() time.Time`

GetWorkFixedStart returns the WorkFixedStart field if non-nil, zero value otherwise.

### GetWorkFixedStartOk

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetWorkFixedStartOk() (*time.Time, bool)`

GetWorkFixedStartOk returns a tuple with the WorkFixedStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkFixedStart

`func (o *GetSchedules200ResponseScheduleRequestsInner) SetWorkFixedStart(v time.Time)`

SetWorkFixedStart sets WorkFixedStart field to given value.

### HasWorkFixedStart

`func (o *GetSchedules200ResponseScheduleRequestsInner) HasWorkFixedStart() bool`

HasWorkFixedStart returns a boolean if a field has been set.

### GetWorkFixedEnd

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetWorkFixedEnd() time.Time`

GetWorkFixedEnd returns the WorkFixedEnd field if non-nil, zero value otherwise.

### GetWorkFixedEndOk

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetWorkFixedEndOk() (*time.Time, bool)`

GetWorkFixedEndOk returns a tuple with the WorkFixedEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkFixedEnd

`func (o *GetSchedules200ResponseScheduleRequestsInner) SetWorkFixedEnd(v time.Time)`

SetWorkFixedEnd sets WorkFixedEnd field to given value.

### HasWorkFixedEnd

`func (o *GetSchedules200ResponseScheduleRequestsInner) HasWorkFixedEnd() bool`

HasWorkFixedEnd returns a boolean if a field has been set.

### GetNote

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *GetSchedules200ResponseScheduleRequestsInner) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *GetSchedules200ResponseScheduleRequestsInner) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetFlow

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetFlow() []GetSchedules200ResponseScheduleRequestsInnerFlowInner`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *GetSchedules200ResponseScheduleRequestsInner) GetFlowOk() (*[]GetSchedules200ResponseScheduleRequestsInnerFlowInner, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *GetSchedules200ResponseScheduleRequestsInner) SetFlow(v []GetSchedules200ResponseScheduleRequestsInnerFlowInner)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *GetSchedules200ResponseScheduleRequestsInner) HasFlow() bool`

HasFlow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


