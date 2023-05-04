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

// checks if the GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner{}

// GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner struct for GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner
type GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner struct {
	// 休暇開始予定
	Start time.Time `json:"start"`
	// 休暇終了予定
	End time.Time `json:"end"`
	// 休暇取得時間
	Minutes int32 `json:"minutes"`
	// 休暇区分コード
	Code int32 `json:"code"`
	// 休暇区分名
	Name string `json:"name"`
}

// NewGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner instantiates a new GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner(start time.Time, end time.Time, minutes int32, code int32, name string) *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner {
	this := GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner{}
	this.Start = start
	this.End = end
	this.Minutes = minutes
	this.Code = code
	this.Name = name
	return &this
}

// NewGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInnerWithDefaults instantiates a new GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInnerWithDefaults() *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner {
	this := GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner{}
	return &this
}

// GetStart returns the Start field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) SetStart(v time.Time) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) GetEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) GetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) SetEnd(v time.Time) {
	o.End = v
}

// GetMinutes returns the Minutes field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) GetMinutes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value
// and a boolean to check if the value has been set.
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) GetMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minutes, true
}

// SetMinutes sets field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) SetMinutes(v int32) {
	o.Minutes = v
}

// GetCode returns the Code field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) SetCode(v int32) {
	o.Code = v
}

// GetName returns the Name field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) SetName(v string) {
	o.Name = v
}

func (o GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	toSerialize["end"] = o.End
	toSerialize["minutes"] = o.Minutes
	toSerialize["code"] = o.Code
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner struct {
	value *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner
	isSet bool
}

func (v NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) Get() *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner {
	return v.value
}

func (v *NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) Set(val *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner(val *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) *NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner {
	return &NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner{value: val, isSet: true}
}

func (v NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHourHolidaysInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


