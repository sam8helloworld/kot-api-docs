/*
KING OF TIME WebAPI

KING OF TIME WebAPI specification https://developer.kingtime.jp/

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GetDailyWorkingTimerecord200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDailyWorkingTimerecord200Response{}

// GetDailyWorkingTimerecord200Response struct for GetDailyWorkingTimerecord200Response
type GetDailyWorkingTimerecord200Response struct {
	Date string `json:"date"`
	DailyWorkings []DailyWorkingTimerecord `json:"dailyWorkings"`
}

// NewGetDailyWorkingTimerecord200Response instantiates a new GetDailyWorkingTimerecord200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDailyWorkingTimerecord200Response(date string, dailyWorkings []DailyWorkingTimerecord) *GetDailyWorkingTimerecord200Response {
	this := GetDailyWorkingTimerecord200Response{}
	this.Date = date
	this.DailyWorkings = dailyWorkings
	return &this
}

// NewGetDailyWorkingTimerecord200ResponseWithDefaults instantiates a new GetDailyWorkingTimerecord200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDailyWorkingTimerecord200ResponseWithDefaults() *GetDailyWorkingTimerecord200Response {
	this := GetDailyWorkingTimerecord200Response{}
	return &this
}

// GetDate returns the Date field value
func (o *GetDailyWorkingTimerecord200Response) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *GetDailyWorkingTimerecord200Response) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *GetDailyWorkingTimerecord200Response) SetDate(v string) {
	o.Date = v
}

// GetDailyWorkings returns the DailyWorkings field value
func (o *GetDailyWorkingTimerecord200Response) GetDailyWorkings() []DailyWorkingTimerecord {
	if o == nil {
		var ret []DailyWorkingTimerecord
		return ret
	}

	return o.DailyWorkings
}

// GetDailyWorkingsOk returns a tuple with the DailyWorkings field value
// and a boolean to check if the value has been set.
func (o *GetDailyWorkingTimerecord200Response) GetDailyWorkingsOk() ([]DailyWorkingTimerecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.DailyWorkings, true
}

// SetDailyWorkings sets field value
func (o *GetDailyWorkingTimerecord200Response) SetDailyWorkings(v []DailyWorkingTimerecord) {
	o.DailyWorkings = v
}

func (o GetDailyWorkingTimerecord200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDailyWorkingTimerecord200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["dailyWorkings"] = o.DailyWorkings
	return toSerialize, nil
}

type NullableGetDailyWorkingTimerecord200Response struct {
	value *GetDailyWorkingTimerecord200Response
	isSet bool
}

func (v NullableGetDailyWorkingTimerecord200Response) Get() *GetDailyWorkingTimerecord200Response {
	return v.value
}

func (v *NullableGetDailyWorkingTimerecord200Response) Set(val *GetDailyWorkingTimerecord200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDailyWorkingTimerecord200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDailyWorkingTimerecord200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDailyWorkingTimerecord200Response(val *GetDailyWorkingTimerecord200Response) *NullableGetDailyWorkingTimerecord200Response {
	return &NullableGetDailyWorkingTimerecord200Response{value: val, isSet: true}
}

func (v NullableGetDailyWorkingTimerecord200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDailyWorkingTimerecord200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


