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

// checks if the GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner{}

// GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner struct for GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner
type GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner struct {
	// 半休種別名
	TypeName string `json:"typeName"`
	// 休暇区分コード
	Code int32 `json:"code"`
	// 休暇区分名
	Name string `json:"name"`
}

// NewGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner instantiates a new GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner(typeName string, code int32, name string) *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner {
	this := GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner{}
	this.TypeName = typeName
	this.Code = code
	this.Name = name
	return &this
}

// NewGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInnerWithDefaults instantiates a new GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInnerWithDefaults() *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner {
	this := GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner{}
	return &this
}

// GetTypeName returns the TypeName field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) SetTypeName(v string) {
	o.TypeName = v
}

// GetCode returns the Code field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) SetCode(v int32) {
	o.Code = v
}

// GetName returns the Name field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) SetName(v string) {
	o.Name = v
}

func (o GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["typeName"] = o.TypeName
	toSerialize["code"] = o.Code
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner struct {
	value *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner
	isSet bool
}

func (v NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) Get() *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner {
	return v.value
}

func (v *NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) Set(val *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner(val *GetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) *NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner {
	return &NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner{value: val, isSet: true}
}

func (v NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSchedules200ResponseScheduleRequestsInnerRequestedHolidaysObtainedHalfdayHolidaysInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


