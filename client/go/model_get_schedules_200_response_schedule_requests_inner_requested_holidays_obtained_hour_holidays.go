/*
KING OF TIME WebAPI

KING OF TIME WebAPI specification https://developer.kingtime.jp/

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays{}

// GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays struct for GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays
type GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays struct {
	// 休暇開始予定
	Start time.Time `json:"start"`
	// 休暇終了予定
	End time.Time `json:"end"`
	// 休暇取得時間
	Minutes int32 `json:"minutes"`
	// 休暇区分コード
	Code string `json:"code"`
	// 休暇区分名
	Name string `json:"name"`
}

// NewGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays instantiates a new GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays(start time.Time, end time.Time, minutes int32, code string, name string) *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays {
	this := GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays{}
	this.Start = start
	this.End = end
	this.Minutes = minutes
	this.Code = code
	this.Name = name
	return &this
}

// NewGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysWithDefaults instantiates a new GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysWithDefaults() *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays {
	this := GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays{}
	return &this
}

// GetStart returns the Start field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) SetStart(v time.Time) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) GetEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) GetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) SetEnd(v time.Time) {
	o.End = v
}

// GetMinutes returns the Minutes field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) GetMinutes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value
// and a boolean to check if the value has been set.
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) GetMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minutes, true
}

// SetMinutes sets field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) SetMinutes(v int32) {
	o.Minutes = v
}

// GetCode returns the Code field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) SetCode(v string) {
	o.Code = v
}

// GetName returns the Name field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) SetName(v string) {
	o.Name = v
}

func (o GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	toSerialize["end"] = o.End
	toSerialize["minutes"] = o.Minutes
	toSerialize["code"] = o.Code
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays struct {
	value *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays
	isSet bool
}

func (v NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) Get() *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays {
	return v.value
}

func (v *NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) Set(val *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays(val *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) *NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays {
	return &NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays{value: val, isSet: true}
}

func (v NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidays) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

