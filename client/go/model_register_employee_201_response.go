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

// checks if the RegisterEmployee201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterEmployee201Response{}

// RegisterEmployee201Response struct for RegisterEmployee201Response
type RegisterEmployee201Response struct {
	// 所属コード
	DivisionCode string `json:"divisionCode"`
	// 性別（male： 男性　female： 女性）
	Gender string `json:"gender"`
	// 雇用区分コード
	TypeCode string `json:"typeCode"`
	// 従業員コード
	Code string `json:"code"`
	// 姓
	LastName string `json:"lastName"`
	// 名
	FirstName string `json:"firstName"`
	// 従業員識別キー（従業員コードが変更されても不変）
	Key string `json:"key"`
}

// NewRegisterEmployee201Response instantiates a new RegisterEmployee201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterEmployee201Response(divisionCode string, gender string, typeCode string, code string, lastName string, firstName string, key string) *RegisterEmployee201Response {
	this := RegisterEmployee201Response{}
	this.DivisionCode = divisionCode
	this.Gender = gender
	this.TypeCode = typeCode
	this.Code = code
	this.LastName = lastName
	this.FirstName = firstName
	this.Key = key
	return &this
}

// NewRegisterEmployee201ResponseWithDefaults instantiates a new RegisterEmployee201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterEmployee201ResponseWithDefaults() *RegisterEmployee201Response {
	this := RegisterEmployee201Response{}
	return &this
}

// GetDivisionCode returns the DivisionCode field value
func (o *RegisterEmployee201Response) GetDivisionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DivisionCode
}

// GetDivisionCodeOk returns a tuple with the DivisionCode field value
// and a boolean to check if the value has been set.
func (o *RegisterEmployee201Response) GetDivisionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DivisionCode, true
}

// SetDivisionCode sets field value
func (o *RegisterEmployee201Response) SetDivisionCode(v string) {
	o.DivisionCode = v
}

// GetGender returns the Gender field value
func (o *RegisterEmployee201Response) GetGender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gender
}

// GetGenderOk returns a tuple with the Gender field value
// and a boolean to check if the value has been set.
func (o *RegisterEmployee201Response) GetGenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gender, true
}

// SetGender sets field value
func (o *RegisterEmployee201Response) SetGender(v string) {
	o.Gender = v
}

// GetTypeCode returns the TypeCode field value
func (o *RegisterEmployee201Response) GetTypeCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeCode
}

// GetTypeCodeOk returns a tuple with the TypeCode field value
// and a boolean to check if the value has been set.
func (o *RegisterEmployee201Response) GetTypeCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeCode, true
}

// SetTypeCode sets field value
func (o *RegisterEmployee201Response) SetTypeCode(v string) {
	o.TypeCode = v
}

// GetCode returns the Code field value
func (o *RegisterEmployee201Response) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *RegisterEmployee201Response) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *RegisterEmployee201Response) SetCode(v string) {
	o.Code = v
}

// GetLastName returns the LastName field value
func (o *RegisterEmployee201Response) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *RegisterEmployee201Response) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *RegisterEmployee201Response) SetLastName(v string) {
	o.LastName = v
}

// GetFirstName returns the FirstName field value
func (o *RegisterEmployee201Response) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *RegisterEmployee201Response) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *RegisterEmployee201Response) SetFirstName(v string) {
	o.FirstName = v
}

// GetKey returns the Key field value
func (o *RegisterEmployee201Response) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *RegisterEmployee201Response) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *RegisterEmployee201Response) SetKey(v string) {
	o.Key = v
}

func (o RegisterEmployee201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterEmployee201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["divisionCode"] = o.DivisionCode
	toSerialize["gender"] = o.Gender
	toSerialize["typeCode"] = o.TypeCode
	toSerialize["code"] = o.Code
	toSerialize["lastName"] = o.LastName
	toSerialize["firstName"] = o.FirstName
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

type NullableRegisterEmployee201Response struct {
	value *RegisterEmployee201Response
	isSet bool
}

func (v NullableRegisterEmployee201Response) Get() *RegisterEmployee201Response {
	return v.value
}

func (v *NullableRegisterEmployee201Response) Set(val *RegisterEmployee201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterEmployee201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterEmployee201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterEmployee201Response(val *RegisterEmployee201Response) *NullableRegisterEmployee201Response {
	return &NullableRegisterEmployee201Response{value: val, isSet: true}
}

func (v NullableRegisterEmployee201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterEmployee201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


