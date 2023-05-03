# DailyWorkingTimerecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **time.Time** | 打刻時間 | 
**Code** | **string** | 打刻種別コード | 
**Name** | **string** | 打刻種別名 | 
**DivisionCode** | **string** | 打刻所属コード | 
**DivisionName** | **string** | 打刻所属名 | 
**Latitude** | **float64** | 緯度 | 
**Longitude** | **float64** | 経度 | 

## Methods

### NewDailyWorkingTimerecord

`func NewDailyWorkingTimerecord(time time.Time, code string, name string, divisionCode string, divisionName string, latitude float64, longitude float64, ) *DailyWorkingTimerecord`

NewDailyWorkingTimerecord instantiates a new DailyWorkingTimerecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyWorkingTimerecordWithDefaults

`func NewDailyWorkingTimerecordWithDefaults() *DailyWorkingTimerecord`

NewDailyWorkingTimerecordWithDefaults instantiates a new DailyWorkingTimerecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *DailyWorkingTimerecord) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DailyWorkingTimerecord) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DailyWorkingTimerecord) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetCode

`func (o *DailyWorkingTimerecord) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DailyWorkingTimerecord) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DailyWorkingTimerecord) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *DailyWorkingTimerecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DailyWorkingTimerecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DailyWorkingTimerecord) SetName(v string)`

SetName sets Name field to given value.


### GetDivisionCode

`func (o *DailyWorkingTimerecord) GetDivisionCode() string`

GetDivisionCode returns the DivisionCode field if non-nil, zero value otherwise.

### GetDivisionCodeOk

`func (o *DailyWorkingTimerecord) GetDivisionCodeOk() (*string, bool)`

GetDivisionCodeOk returns a tuple with the DivisionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionCode

`func (o *DailyWorkingTimerecord) SetDivisionCode(v string)`

SetDivisionCode sets DivisionCode field to given value.


### GetDivisionName

`func (o *DailyWorkingTimerecord) GetDivisionName() string`

GetDivisionName returns the DivisionName field if non-nil, zero value otherwise.

### GetDivisionNameOk

`func (o *DailyWorkingTimerecord) GetDivisionNameOk() (*string, bool)`

GetDivisionNameOk returns a tuple with the DivisionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionName

`func (o *DailyWorkingTimerecord) SetDivisionName(v string)`

SetDivisionName sets DivisionName field to given value.


### GetLatitude

`func (o *DailyWorkingTimerecord) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *DailyWorkingTimerecord) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *DailyWorkingTimerecord) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *DailyWorkingTimerecord) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *DailyWorkingTimerecord) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *DailyWorkingTimerecord) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


