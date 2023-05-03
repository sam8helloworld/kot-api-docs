# DailyWorkingTimerecordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **time.Time** | 打刻時間 | 
**Code** | Pointer to **string** | 打刻種別コード（1： 出勤　2： 退勤　3： 休憩開始　4： 休憩終了　7： 外出入　8： 外出戻） 省略時は、KING OF TIMEの処理に従って打刻種別を決定します。 | [optional] 
**DivisionCode** | Pointer to **string** | 打刻先所属コード 省略時は、該当従業員が所属している所属を打刻先所属として扱います。 | [optional] 
**Date** | **string** | 勤務日 指定された勤務日に打刻データが紐づきます。 | 
**Latitude** | Pointer to **float64** | 緯度 | [optional] 
**Longitude** | Pointer to **float64** | 経度 | [optional] 

## Methods

### NewDailyWorkingTimerecordRequest

`func NewDailyWorkingTimerecordRequest(time time.Time, date string, ) *DailyWorkingTimerecordRequest`

NewDailyWorkingTimerecordRequest instantiates a new DailyWorkingTimerecordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyWorkingTimerecordRequestWithDefaults

`func NewDailyWorkingTimerecordRequestWithDefaults() *DailyWorkingTimerecordRequest`

NewDailyWorkingTimerecordRequestWithDefaults instantiates a new DailyWorkingTimerecordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *DailyWorkingTimerecordRequest) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DailyWorkingTimerecordRequest) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DailyWorkingTimerecordRequest) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetCode

`func (o *DailyWorkingTimerecordRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DailyWorkingTimerecordRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DailyWorkingTimerecordRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *DailyWorkingTimerecordRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDivisionCode

`func (o *DailyWorkingTimerecordRequest) GetDivisionCode() string`

GetDivisionCode returns the DivisionCode field if non-nil, zero value otherwise.

### GetDivisionCodeOk

`func (o *DailyWorkingTimerecordRequest) GetDivisionCodeOk() (*string, bool)`

GetDivisionCodeOk returns a tuple with the DivisionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionCode

`func (o *DailyWorkingTimerecordRequest) SetDivisionCode(v string)`

SetDivisionCode sets DivisionCode field to given value.

### HasDivisionCode

`func (o *DailyWorkingTimerecordRequest) HasDivisionCode() bool`

HasDivisionCode returns a boolean if a field has been set.

### GetDate

`func (o *DailyWorkingTimerecordRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DailyWorkingTimerecordRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DailyWorkingTimerecordRequest) SetDate(v string)`

SetDate sets Date field to given value.


### GetLatitude

`func (o *DailyWorkingTimerecordRequest) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *DailyWorkingTimerecordRequest) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *DailyWorkingTimerecordRequest) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *DailyWorkingTimerecordRequest) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *DailyWorkingTimerecordRequest) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *DailyWorkingTimerecordRequest) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *DailyWorkingTimerecordRequest) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *DailyWorkingTimerecordRequest) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


