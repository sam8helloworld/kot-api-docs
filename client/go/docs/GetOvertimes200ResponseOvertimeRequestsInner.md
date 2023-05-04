# GetOvertimes200ResponseOvertimeRequestsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | 対象日 | 
**RequestedDate** | **string** | 申請日 | 
**EmployeeKey** | **string** | 従業員識別キー（従業員コードが変更されても不変） | 
**RequestKey** | **string** | 申請識別キー | 
**Applicant** | [**GetOvertimes200ResponseOvertimeRequestsInnerApplicant**](GetOvertimes200ResponseOvertimeRequestsInnerApplicant.md) |  | 
**Status** | **string** | 申請ステータス（applying 申請中　rejected： 棄却　approved： 承認） | 
**CurrentFlow** | **int32** | 現在の承認フローレベル（１～５） | 
**LastModifiedAdministratorKey** | **string** | 最終更新管理者識別キー | 
**Message** | **string** | 申請メッセージ | 
**AdminComment** | **string** | 管理者コメント | 
**Requested** | [**GetOvertimes200ResponseOvertimeRequestsInnerRequested**](GetOvertimes200ResponseOvertimeRequestsInnerRequested.md) |  | 
**Current** | [**GetOvertimes200ResponseOvertimeRequestsInnerCurrent**](GetOvertimes200ResponseOvertimeRequestsInnerCurrent.md) |  | 
**Flow** | Pointer to [**[]GetSchedules200ResponseScheduleRequestsInnerFlowInner**](GetSchedules200ResponseScheduleRequestsInnerFlowInner.md) | 承認フロー | [optional] 

## Methods

### NewGetOvertimes200ResponseOvertimeRequestsInner

`func NewGetOvertimes200ResponseOvertimeRequestsInner(date string, requestedDate string, employeeKey string, requestKey string, applicant GetOvertimes200ResponseOvertimeRequestsInnerApplicant, status string, currentFlow int32, lastModifiedAdministratorKey string, message string, adminComment string, requested GetOvertimes200ResponseOvertimeRequestsInnerRequested, current GetOvertimes200ResponseOvertimeRequestsInnerCurrent, ) *GetOvertimes200ResponseOvertimeRequestsInner`

NewGetOvertimes200ResponseOvertimeRequestsInner instantiates a new GetOvertimes200ResponseOvertimeRequestsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOvertimes200ResponseOvertimeRequestsInnerWithDefaults

`func NewGetOvertimes200ResponseOvertimeRequestsInnerWithDefaults() *GetOvertimes200ResponseOvertimeRequestsInner`

NewGetOvertimes200ResponseOvertimeRequestsInnerWithDefaults instantiates a new GetOvertimes200ResponseOvertimeRequestsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetRequestedDate

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetRequestedDate() string`

GetRequestedDate returns the RequestedDate field if non-nil, zero value otherwise.

### GetRequestedDateOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetRequestedDateOk() (*string, bool)`

GetRequestedDateOk returns a tuple with the RequestedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDate

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) SetRequestedDate(v string)`

SetRequestedDate sets RequestedDate field to given value.


### GetEmployeeKey

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetEmployeeKey() string`

GetEmployeeKey returns the EmployeeKey field if non-nil, zero value otherwise.

### GetEmployeeKeyOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetEmployeeKeyOk() (*string, bool)`

GetEmployeeKeyOk returns a tuple with the EmployeeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeKey

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) SetEmployeeKey(v string)`

SetEmployeeKey sets EmployeeKey field to given value.


### GetRequestKey

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetRequestKey() string`

GetRequestKey returns the RequestKey field if non-nil, zero value otherwise.

### GetRequestKeyOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetRequestKeyOk() (*string, bool)`

GetRequestKeyOk returns a tuple with the RequestKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestKey

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) SetRequestKey(v string)`

SetRequestKey sets RequestKey field to given value.


### GetApplicant

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetApplicant() GetOvertimes200ResponseOvertimeRequestsInnerApplicant`

GetApplicant returns the Applicant field if non-nil, zero value otherwise.

### GetApplicantOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetApplicantOk() (*GetOvertimes200ResponseOvertimeRequestsInnerApplicant, bool)`

GetApplicantOk returns a tuple with the Applicant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicant

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) SetApplicant(v GetOvertimes200ResponseOvertimeRequestsInnerApplicant)`

SetApplicant sets Applicant field to given value.


### GetStatus

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCurrentFlow

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetCurrentFlow() int32`

GetCurrentFlow returns the CurrentFlow field if non-nil, zero value otherwise.

### GetCurrentFlowOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetCurrentFlowOk() (*int32, bool)`

GetCurrentFlowOk returns a tuple with the CurrentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentFlow

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) SetCurrentFlow(v int32)`

SetCurrentFlow sets CurrentFlow field to given value.


### GetLastModifiedAdministratorKey

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetLastModifiedAdministratorKey() string`

GetLastModifiedAdministratorKey returns the LastModifiedAdministratorKey field if non-nil, zero value otherwise.

### GetLastModifiedAdministratorKeyOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetLastModifiedAdministratorKeyOk() (*string, bool)`

GetLastModifiedAdministratorKeyOk returns a tuple with the LastModifiedAdministratorKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAdministratorKey

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) SetLastModifiedAdministratorKey(v string)`

SetLastModifiedAdministratorKey sets LastModifiedAdministratorKey field to given value.


### GetMessage

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAdminComment

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetAdminComment() string`

GetAdminComment returns the AdminComment field if non-nil, zero value otherwise.

### GetAdminCommentOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetAdminCommentOk() (*string, bool)`

GetAdminCommentOk returns a tuple with the AdminComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminComment

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) SetAdminComment(v string)`

SetAdminComment sets AdminComment field to given value.


### GetRequested

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetRequested() GetOvertimes200ResponseOvertimeRequestsInnerRequested`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetRequestedOk() (*GetOvertimes200ResponseOvertimeRequestsInnerRequested, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) SetRequested(v GetOvertimes200ResponseOvertimeRequestsInnerRequested)`

SetRequested sets Requested field to given value.


### GetCurrent

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetCurrent() GetOvertimes200ResponseOvertimeRequestsInnerCurrent`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetCurrentOk() (*GetOvertimes200ResponseOvertimeRequestsInnerCurrent, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) SetCurrent(v GetOvertimes200ResponseOvertimeRequestsInnerCurrent)`

SetCurrent sets Current field to given value.


### GetFlow

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetFlow() []GetSchedules200ResponseScheduleRequestsInnerFlowInner`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) GetFlowOk() (*[]GetSchedules200ResponseScheduleRequestsInnerFlowInner, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) SetFlow(v []GetSchedules200ResponseScheduleRequestsInnerFlowInner)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *GetOvertimes200ResponseOvertimeRequestsInner) HasFlow() bool`

HasFlow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


