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

// checks if the UpdateEmployeeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEmployeeRequest{}

// UpdateEmployeeRequest struct for UpdateEmployeeRequest
type UpdateEmployeeRequest struct {
	// 所属コード
	DivisionCode *string `json:"divisionCode,omitempty"`
	// 性別（male： 男性　female： 女性）
	Gender *string `json:"gender,omitempty"`
	// 雇用区分コード
	TypeCode *string `json:"typeCode,omitempty"`
	// 従業員コード
	Code *string `json:"code,omitempty"`
	// 姓
	LastName *string `json:"lastName,omitempty"`
	// 名
	FirstName *string `json:"firstName,omitempty"`
	// 姓（カナ）
	LastNamePhonetics *string `json:"lastNamePhonetics,omitempty"`
	// 名（カナ）
	FirstNamePhonetics *string `json:"firstNamePhonetics,omitempty"`
	// 入社年月日
	HiredDate *string `json:"hiredDate,omitempty"`
	// 生年月日
	BirthDate *string `json:"birthDate,omitempty"`
	// 退職年月日
	ResignationDate *string `json:"resignationDate,omitempty"`
	// メールアドレス
	EmailAddresses []string `json:"emailAddresses,omitempty"`
	// 日の契約労働時間
	AllDayRegardingWorkInMinute *int32 `json:"allDayRegardingWorkInMinute,omitempty"`
}

// NewUpdateEmployeeRequest instantiates a new UpdateEmployeeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEmployeeRequest() *UpdateEmployeeRequest {
	this := UpdateEmployeeRequest{}
	return &this
}

// NewUpdateEmployeeRequestWithDefaults instantiates a new UpdateEmployeeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEmployeeRequestWithDefaults() *UpdateEmployeeRequest {
	this := UpdateEmployeeRequest{}
	return &this
}

// GetDivisionCode returns the DivisionCode field value if set, zero value otherwise.
func (o *UpdateEmployeeRequest) GetDivisionCode() string {
	if o == nil || IsNil(o.DivisionCode) {
		var ret string
		return ret
	}
	return *o.DivisionCode
}

// GetDivisionCodeOk returns a tuple with the DivisionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmployeeRequest) GetDivisionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DivisionCode) {
		return nil, false
	}
	return o.DivisionCode, true
}

// HasDivisionCode returns a boolean if a field has been set.
func (o *UpdateEmployeeRequest) HasDivisionCode() bool {
	if o != nil && !IsNil(o.DivisionCode) {
		return true
	}

	return false
}

// SetDivisionCode gets a reference to the given string and assigns it to the DivisionCode field.
func (o *UpdateEmployeeRequest) SetDivisionCode(v string) {
	o.DivisionCode = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *UpdateEmployeeRequest) GetGender() string {
	if o == nil || IsNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmployeeRequest) GetGenderOk() (*string, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *UpdateEmployeeRequest) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *UpdateEmployeeRequest) SetGender(v string) {
	o.Gender = &v
}

// GetTypeCode returns the TypeCode field value if set, zero value otherwise.
func (o *UpdateEmployeeRequest) GetTypeCode() string {
	if o == nil || IsNil(o.TypeCode) {
		var ret string
		return ret
	}
	return *o.TypeCode
}

// GetTypeCodeOk returns a tuple with the TypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmployeeRequest) GetTypeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TypeCode) {
		return nil, false
	}
	return o.TypeCode, true
}

// HasTypeCode returns a boolean if a field has been set.
func (o *UpdateEmployeeRequest) HasTypeCode() bool {
	if o != nil && !IsNil(o.TypeCode) {
		return true
	}

	return false
}

// SetTypeCode gets a reference to the given string and assigns it to the TypeCode field.
func (o *UpdateEmployeeRequest) SetTypeCode(v string) {
	o.TypeCode = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UpdateEmployeeRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmployeeRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UpdateEmployeeRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UpdateEmployeeRequest) SetCode(v string) {
	o.Code = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UpdateEmployeeRequest) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmployeeRequest) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UpdateEmployeeRequest) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UpdateEmployeeRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UpdateEmployeeRequest) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmployeeRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UpdateEmployeeRequest) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UpdateEmployeeRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastNamePhonetics returns the LastNamePhonetics field value if set, zero value otherwise.
func (o *UpdateEmployeeRequest) GetLastNamePhonetics() string {
	if o == nil || IsNil(o.LastNamePhonetics) {
		var ret string
		return ret
	}
	return *o.LastNamePhonetics
}

// GetLastNamePhoneticsOk returns a tuple with the LastNamePhonetics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmployeeRequest) GetLastNamePhoneticsOk() (*string, bool) {
	if o == nil || IsNil(o.LastNamePhonetics) {
		return nil, false
	}
	return o.LastNamePhonetics, true
}

// HasLastNamePhonetics returns a boolean if a field has been set.
func (o *UpdateEmployeeRequest) HasLastNamePhonetics() bool {
	if o != nil && !IsNil(o.LastNamePhonetics) {
		return true
	}

	return false
}

// SetLastNamePhonetics gets a reference to the given string and assigns it to the LastNamePhonetics field.
func (o *UpdateEmployeeRequest) SetLastNamePhonetics(v string) {
	o.LastNamePhonetics = &v
}

// GetFirstNamePhonetics returns the FirstNamePhonetics field value if set, zero value otherwise.
func (o *UpdateEmployeeRequest) GetFirstNamePhonetics() string {
	if o == nil || IsNil(o.FirstNamePhonetics) {
		var ret string
		return ret
	}
	return *o.FirstNamePhonetics
}

// GetFirstNamePhoneticsOk returns a tuple with the FirstNamePhonetics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmployeeRequest) GetFirstNamePhoneticsOk() (*string, bool) {
	if o == nil || IsNil(o.FirstNamePhonetics) {
		return nil, false
	}
	return o.FirstNamePhonetics, true
}

// HasFirstNamePhonetics returns a boolean if a field has been set.
func (o *UpdateEmployeeRequest) HasFirstNamePhonetics() bool {
	if o != nil && !IsNil(o.FirstNamePhonetics) {
		return true
	}

	return false
}

// SetFirstNamePhonetics gets a reference to the given string and assigns it to the FirstNamePhonetics field.
func (o *UpdateEmployeeRequest) SetFirstNamePhonetics(v string) {
	o.FirstNamePhonetics = &v
}

// GetHiredDate returns the HiredDate field value if set, zero value otherwise.
func (o *UpdateEmployeeRequest) GetHiredDate() string {
	if o == nil || IsNil(o.HiredDate) {
		var ret string
		return ret
	}
	return *o.HiredDate
}

// GetHiredDateOk returns a tuple with the HiredDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmployeeRequest) GetHiredDateOk() (*string, bool) {
	if o == nil || IsNil(o.HiredDate) {
		return nil, false
	}
	return o.HiredDate, true
}

// HasHiredDate returns a boolean if a field has been set.
func (o *UpdateEmployeeRequest) HasHiredDate() bool {
	if o != nil && !IsNil(o.HiredDate) {
		return true
	}

	return false
}

// SetHiredDate gets a reference to the given string and assigns it to the HiredDate field.
func (o *UpdateEmployeeRequest) SetHiredDate(v string) {
	o.HiredDate = &v
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise.
func (o *UpdateEmployeeRequest) GetBirthDate() string {
	if o == nil || IsNil(o.BirthDate) {
		var ret string
		return ret
	}
	return *o.BirthDate
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmployeeRequest) GetBirthDateOk() (*string, bool) {
	if o == nil || IsNil(o.BirthDate) {
		return nil, false
	}
	return o.BirthDate, true
}

// HasBirthDate returns a boolean if a field has been set.
func (o *UpdateEmployeeRequest) HasBirthDate() bool {
	if o != nil && !IsNil(o.BirthDate) {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given string and assigns it to the BirthDate field.
func (o *UpdateEmployeeRequest) SetBirthDate(v string) {
	o.BirthDate = &v
}

// GetResignationDate returns the ResignationDate field value if set, zero value otherwise.
func (o *UpdateEmployeeRequest) GetResignationDate() string {
	if o == nil || IsNil(o.ResignationDate) {
		var ret string
		return ret
	}
	return *o.ResignationDate
}

// GetResignationDateOk returns a tuple with the ResignationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmployeeRequest) GetResignationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ResignationDate) {
		return nil, false
	}
	return o.ResignationDate, true
}

// HasResignationDate returns a boolean if a field has been set.
func (o *UpdateEmployeeRequest) HasResignationDate() bool {
	if o != nil && !IsNil(o.ResignationDate) {
		return true
	}

	return false
}

// SetResignationDate gets a reference to the given string and assigns it to the ResignationDate field.
func (o *UpdateEmployeeRequest) SetResignationDate(v string) {
	o.ResignationDate = &v
}

// GetEmailAddresses returns the EmailAddresses field value if set, zero value otherwise.
func (o *UpdateEmployeeRequest) GetEmailAddresses() []string {
	if o == nil || IsNil(o.EmailAddresses) {
		var ret []string
		return ret
	}
	return o.EmailAddresses
}

// GetEmailAddressesOk returns a tuple with the EmailAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmployeeRequest) GetEmailAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailAddresses) {
		return nil, false
	}
	return o.EmailAddresses, true
}

// HasEmailAddresses returns a boolean if a field has been set.
func (o *UpdateEmployeeRequest) HasEmailAddresses() bool {
	if o != nil && !IsNil(o.EmailAddresses) {
		return true
	}

	return false
}

// SetEmailAddresses gets a reference to the given []string and assigns it to the EmailAddresses field.
func (o *UpdateEmployeeRequest) SetEmailAddresses(v []string) {
	o.EmailAddresses = v
}

// GetAllDayRegardingWorkInMinute returns the AllDayRegardingWorkInMinute field value if set, zero value otherwise.
func (o *UpdateEmployeeRequest) GetAllDayRegardingWorkInMinute() int32 {
	if o == nil || IsNil(o.AllDayRegardingWorkInMinute) {
		var ret int32
		return ret
	}
	return *o.AllDayRegardingWorkInMinute
}

// GetAllDayRegardingWorkInMinuteOk returns a tuple with the AllDayRegardingWorkInMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmployeeRequest) GetAllDayRegardingWorkInMinuteOk() (*int32, bool) {
	if o == nil || IsNil(o.AllDayRegardingWorkInMinute) {
		return nil, false
	}
	return o.AllDayRegardingWorkInMinute, true
}

// HasAllDayRegardingWorkInMinute returns a boolean if a field has been set.
func (o *UpdateEmployeeRequest) HasAllDayRegardingWorkInMinute() bool {
	if o != nil && !IsNil(o.AllDayRegardingWorkInMinute) {
		return true
	}

	return false
}

// SetAllDayRegardingWorkInMinute gets a reference to the given int32 and assigns it to the AllDayRegardingWorkInMinute field.
func (o *UpdateEmployeeRequest) SetAllDayRegardingWorkInMinute(v int32) {
	o.AllDayRegardingWorkInMinute = &v
}

func (o UpdateEmployeeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEmployeeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DivisionCode) {
		toSerialize["divisionCode"] = o.DivisionCode
	}
	if !IsNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !IsNil(o.TypeCode) {
		toSerialize["typeCode"] = o.TypeCode
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastNamePhonetics) {
		toSerialize["lastNamePhonetics"] = o.LastNamePhonetics
	}
	if !IsNil(o.FirstNamePhonetics) {
		toSerialize["firstNamePhonetics"] = o.FirstNamePhonetics
	}
	if !IsNil(o.HiredDate) {
		toSerialize["hiredDate"] = o.HiredDate
	}
	if !IsNil(o.BirthDate) {
		toSerialize["birthDate"] = o.BirthDate
	}
	if !IsNil(o.ResignationDate) {
		toSerialize["resignationDate"] = o.ResignationDate
	}
	if !IsNil(o.EmailAddresses) {
		toSerialize["emailAddresses"] = o.EmailAddresses
	}
	if !IsNil(o.AllDayRegardingWorkInMinute) {
		toSerialize["allDayRegardingWorkInMinute"] = o.AllDayRegardingWorkInMinute
	}
	return toSerialize, nil
}

type NullableUpdateEmployeeRequest struct {
	value *UpdateEmployeeRequest
	isSet bool
}

func (v NullableUpdateEmployeeRequest) Get() *UpdateEmployeeRequest {
	return v.value
}

func (v *NullableUpdateEmployeeRequest) Set(val *UpdateEmployeeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEmployeeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEmployeeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEmployeeRequest(val *UpdateEmployeeRequest) *NullableUpdateEmployeeRequest {
	return &NullableUpdateEmployeeRequest{value: val, isSet: true}
}

func (v NullableUpdateEmployeeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEmployeeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


