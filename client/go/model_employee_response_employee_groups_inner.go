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

// checks if the EmployeeResponseEmployeeGroupsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmployeeResponseEmployeeGroupsInner{}

// EmployeeResponseEmployeeGroupsInner struct for EmployeeResponseEmployeeGroupsInner
type EmployeeResponseEmployeeGroupsInner struct {
	// 従業員グループコード
	Code string `json:"code"`
	// 従業員グループ名
	Name string `json:"name"`
}

// NewEmployeeResponseEmployeeGroupsInner instantiates a new EmployeeResponseEmployeeGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployeeResponseEmployeeGroupsInner(code string, name string) *EmployeeResponseEmployeeGroupsInner {
	this := EmployeeResponseEmployeeGroupsInner{}
	this.Code = code
	this.Name = name
	return &this
}

// NewEmployeeResponseEmployeeGroupsInnerWithDefaults instantiates a new EmployeeResponseEmployeeGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployeeResponseEmployeeGroupsInnerWithDefaults() *EmployeeResponseEmployeeGroupsInner {
	this := EmployeeResponseEmployeeGroupsInner{}
	return &this
}

// GetCode returns the Code field value
func (o *EmployeeResponseEmployeeGroupsInner) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *EmployeeResponseEmployeeGroupsInner) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *EmployeeResponseEmployeeGroupsInner) SetCode(v string) {
	o.Code = v
}

// GetName returns the Name field value
func (o *EmployeeResponseEmployeeGroupsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EmployeeResponseEmployeeGroupsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EmployeeResponseEmployeeGroupsInner) SetName(v string) {
	o.Name = v
}

func (o EmployeeResponseEmployeeGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmployeeResponseEmployeeGroupsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableEmployeeResponseEmployeeGroupsInner struct {
	value *EmployeeResponseEmployeeGroupsInner
	isSet bool
}

func (v NullableEmployeeResponseEmployeeGroupsInner) Get() *EmployeeResponseEmployeeGroupsInner {
	return v.value
}

func (v *NullableEmployeeResponseEmployeeGroupsInner) Set(val *EmployeeResponseEmployeeGroupsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployeeResponseEmployeeGroupsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployeeResponseEmployeeGroupsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployeeResponseEmployeeGroupsInner(val *EmployeeResponseEmployeeGroupsInner) *NullableEmployeeResponseEmployeeGroupsInner {
	return &NullableEmployeeResponseEmployeeGroupsInner{value: val, isSet: true}
}

func (v NullableEmployeeResponseEmployeeGroupsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployeeResponseEmployeeGroupsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

