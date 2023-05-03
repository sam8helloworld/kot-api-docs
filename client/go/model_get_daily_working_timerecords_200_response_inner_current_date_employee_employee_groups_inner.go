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

// checks if the GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner{}

// GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner struct for GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner
type GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner struct {
	// 従業員グループコード
	Code string `json:"code"`
	// 従業員グループ名
	Name string `json:"name"`
}

// NewGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner instantiates a new GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner(code string, name string) *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner {
	this := GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner{}
	this.Code = code
	this.Name = name
	return &this
}

// NewGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInnerWithDefaults instantiates a new GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInnerWithDefaults() *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner {
	this := GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner{}
	return &this
}

// GetCode returns the Code field value
func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner) SetCode(v string) {
	o.Code = v
}

// GetName returns the Name field value
func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner) SetName(v string) {
	o.Name = v
}

func (o GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner struct {
	value *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner
	isSet bool
}

func (v NullableGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner) Get() *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner {
	return v.value
}

func (v *NullableGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner) Set(val *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner(val *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner) *NullableGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner {
	return &NullableGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner{value: val, isSet: true}
}

func (v NullableGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


