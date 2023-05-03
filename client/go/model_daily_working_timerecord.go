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

// checks if the DailyWorkingTimerecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DailyWorkingTimerecord{}

// DailyWorkingTimerecord 打刻
type DailyWorkingTimerecord struct {
	// 打刻時間
	Time time.Time `json:"time"`
	// 打刻種別コード
	Code string `json:"code"`
	// 打刻種別名
	Name string `json:"name"`
	// 打刻所属コード
	DivisionCode string `json:"divisionCode"`
	// 打刻所属名
	DivisionName string `json:"divisionName"`
	// 緯度
	Latitude float64 `json:"latitude"`
	// 経度
	Longitude float64 `json:"longitude"`
}

// NewDailyWorkingTimerecord instantiates a new DailyWorkingTimerecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailyWorkingTimerecord(time time.Time, code string, name string, divisionCode string, divisionName string, latitude float64, longitude float64) *DailyWorkingTimerecord {
	this := DailyWorkingTimerecord{}
	this.Time = time
	this.Code = code
	this.Name = name
	this.DivisionCode = divisionCode
	this.DivisionName = divisionName
	this.Latitude = latitude
	this.Longitude = longitude
	return &this
}

// NewDailyWorkingTimerecordWithDefaults instantiates a new DailyWorkingTimerecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyWorkingTimerecordWithDefaults() *DailyWorkingTimerecord {
	this := DailyWorkingTimerecord{}
	return &this
}

// GetTime returns the Time field value
func (o *DailyWorkingTimerecord) GetTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingTimerecord) GetTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *DailyWorkingTimerecord) SetTime(v time.Time) {
	o.Time = v
}

// GetCode returns the Code field value
func (o *DailyWorkingTimerecord) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingTimerecord) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *DailyWorkingTimerecord) SetCode(v string) {
	o.Code = v
}

// GetName returns the Name field value
func (o *DailyWorkingTimerecord) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingTimerecord) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DailyWorkingTimerecord) SetName(v string) {
	o.Name = v
}

// GetDivisionCode returns the DivisionCode field value
func (o *DailyWorkingTimerecord) GetDivisionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DivisionCode
}

// GetDivisionCodeOk returns a tuple with the DivisionCode field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingTimerecord) GetDivisionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DivisionCode, true
}

// SetDivisionCode sets field value
func (o *DailyWorkingTimerecord) SetDivisionCode(v string) {
	o.DivisionCode = v
}

// GetDivisionName returns the DivisionName field value
func (o *DailyWorkingTimerecord) GetDivisionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DivisionName
}

// GetDivisionNameOk returns a tuple with the DivisionName field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingTimerecord) GetDivisionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DivisionName, true
}

// SetDivisionName sets field value
func (o *DailyWorkingTimerecord) SetDivisionName(v string) {
	o.DivisionName = v
}

// GetLatitude returns the Latitude field value
func (o *DailyWorkingTimerecord) GetLatitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingTimerecord) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *DailyWorkingTimerecord) SetLatitude(v float64) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *DailyWorkingTimerecord) GetLongitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingTimerecord) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *DailyWorkingTimerecord) SetLongitude(v float64) {
	o.Longitude = v
}

func (o DailyWorkingTimerecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DailyWorkingTimerecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["time"] = o.Time
	toSerialize["code"] = o.Code
	toSerialize["name"] = o.Name
	toSerialize["divisionCode"] = o.DivisionCode
	toSerialize["divisionName"] = o.DivisionName
	toSerialize["latitude"] = o.Latitude
	toSerialize["longitude"] = o.Longitude
	return toSerialize, nil
}

type NullableDailyWorkingTimerecord struct {
	value *DailyWorkingTimerecord
	isSet bool
}

func (v NullableDailyWorkingTimerecord) Get() *DailyWorkingTimerecord {
	return v.value
}

func (v *NullableDailyWorkingTimerecord) Set(val *DailyWorkingTimerecord) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyWorkingTimerecord) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyWorkingTimerecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyWorkingTimerecord(val *DailyWorkingTimerecord) *NullableDailyWorkingTimerecord {
	return &NullableDailyWorkingTimerecord{value: val, isSet: true}
}

func (v NullableDailyWorkingTimerecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyWorkingTimerecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

