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

// checks if the EmployeeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmployeeResponse{}

// EmployeeResponse struct for EmployeeResponse
type EmployeeResponse struct {
	// 所属コード
	DivisionCode string `json:"divisionCode"`
	// 所属名
	DivisionName string `json:"divisionName"`
	// 性別（male： 男性　female： 女性）
	Gender string `json:"gender"`
	// 雇用区分コード
	TypeCode string `json:"typeCode"`
	// 雇用区分名
	TypeName string `json:"typeName"`
	// 従業員コード
	Code string `json:"code"`
	// 従業員識別キー（従業員コードが変更されても不変）
	Key string `json:"key"`
	// 姓
	LastName string `json:"lastName"`
	// 名
	FirstName string `json:"firstName"`
	// 従業員グループ情報
	EmployeeGroups []EmployeeResponseEmployeeGroupsInner `json:"employeeGroups"`
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

// NewEmployeeResponse instantiates a new EmployeeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployeeResponse(divisionCode string, divisionName string, gender string, typeCode string, typeName string, code string, key string, lastName string, firstName string, employeeGroups []EmployeeResponseEmployeeGroupsInner) *EmployeeResponse {
	this := EmployeeResponse{}
	this.DivisionCode = divisionCode
	this.DivisionName = divisionName
	this.Gender = gender
	this.TypeCode = typeCode
	this.TypeName = typeName
	this.Code = code
	this.Key = key
	this.LastName = lastName
	this.FirstName = firstName
	this.EmployeeGroups = employeeGroups
	return &this
}

// NewEmployeeResponseWithDefaults instantiates a new EmployeeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployeeResponseWithDefaults() *EmployeeResponse {
	this := EmployeeResponse{}
	return &this
}

// GetDivisionCode returns the DivisionCode field value
func (o *EmployeeResponse) GetDivisionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DivisionCode
}

// GetDivisionCodeOk returns a tuple with the DivisionCode field value
// and a boolean to check if the value has been set.
func (o *EmployeeResponse) GetDivisionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DivisionCode, true
}

// SetDivisionCode sets field value
func (o *EmployeeResponse) SetDivisionCode(v string) {
	o.DivisionCode = v
}

// GetDivisionName returns the DivisionName field value
func (o *EmployeeResponse) GetDivisionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DivisionName
}

// GetDivisionNameOk returns a tuple with the DivisionName field value
// and a boolean to check if the value has been set.
func (o *EmployeeResponse) GetDivisionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DivisionName, true
}

// SetDivisionName sets field value
func (o *EmployeeResponse) SetDivisionName(v string) {
	o.DivisionName = v
}

// GetGender returns the Gender field value
func (o *EmployeeResponse) GetGender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gender
}

// GetGenderOk returns a tuple with the Gender field value
// and a boolean to check if the value has been set.
func (o *EmployeeResponse) GetGenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gender, true
}

// SetGender sets field value
func (o *EmployeeResponse) SetGender(v string) {
	o.Gender = v
}

// GetTypeCode returns the TypeCode field value
func (o *EmployeeResponse) GetTypeCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeCode
}

// GetTypeCodeOk returns a tuple with the TypeCode field value
// and a boolean to check if the value has been set.
func (o *EmployeeResponse) GetTypeCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeCode, true
}

// SetTypeCode sets field value
func (o *EmployeeResponse) SetTypeCode(v string) {
	o.TypeCode = v
}

// GetTypeName returns the TypeName field value
func (o *EmployeeResponse) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *EmployeeResponse) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *EmployeeResponse) SetTypeName(v string) {
	o.TypeName = v
}

// GetCode returns the Code field value
func (o *EmployeeResponse) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *EmployeeResponse) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *EmployeeResponse) SetCode(v string) {
	o.Code = v
}

// GetKey returns the Key field value
func (o *EmployeeResponse) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EmployeeResponse) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *EmployeeResponse) SetKey(v string) {
	o.Key = v
}

// GetLastName returns the LastName field value
func (o *EmployeeResponse) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *EmployeeResponse) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *EmployeeResponse) SetLastName(v string) {
	o.LastName = v
}

// GetFirstName returns the FirstName field value
func (o *EmployeeResponse) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *EmployeeResponse) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *EmployeeResponse) SetFirstName(v string) {
	o.FirstName = v
}

// GetEmployeeGroups returns the EmployeeGroups field value
func (o *EmployeeResponse) GetEmployeeGroups() []EmployeeResponseEmployeeGroupsInner {
	if o == nil {
		var ret []EmployeeResponseEmployeeGroupsInner
		return ret
	}

	return o.EmployeeGroups
}

// GetEmployeeGroupsOk returns a tuple with the EmployeeGroups field value
// and a boolean to check if the value has been set.
func (o *EmployeeResponse) GetEmployeeGroupsOk() ([]EmployeeResponseEmployeeGroupsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmployeeGroups, true
}

// SetEmployeeGroups sets field value
func (o *EmployeeResponse) SetEmployeeGroups(v []EmployeeResponseEmployeeGroupsInner) {
	o.EmployeeGroups = v
}

// GetLastNamePhonetics returns the LastNamePhonetics field value if set, zero value otherwise.
func (o *EmployeeResponse) GetLastNamePhonetics() string {
	if o == nil || IsNil(o.LastNamePhonetics) {
		var ret string
		return ret
	}
	return *o.LastNamePhonetics
}

// GetLastNamePhoneticsOk returns a tuple with the LastNamePhonetics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeResponse) GetLastNamePhoneticsOk() (*string, bool) {
	if o == nil || IsNil(o.LastNamePhonetics) {
		return nil, false
	}
	return o.LastNamePhonetics, true
}

// HasLastNamePhonetics returns a boolean if a field has been set.
func (o *EmployeeResponse) HasLastNamePhonetics() bool {
	if o != nil && !IsNil(o.LastNamePhonetics) {
		return true
	}

	return false
}

// SetLastNamePhonetics gets a reference to the given string and assigns it to the LastNamePhonetics field.
func (o *EmployeeResponse) SetLastNamePhonetics(v string) {
	o.LastNamePhonetics = &v
}

// GetFirstNamePhonetics returns the FirstNamePhonetics field value if set, zero value otherwise.
func (o *EmployeeResponse) GetFirstNamePhonetics() string {
	if o == nil || IsNil(o.FirstNamePhonetics) {
		var ret string
		return ret
	}
	return *o.FirstNamePhonetics
}

// GetFirstNamePhoneticsOk returns a tuple with the FirstNamePhonetics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeResponse) GetFirstNamePhoneticsOk() (*string, bool) {
	if o == nil || IsNil(o.FirstNamePhonetics) {
		return nil, false
	}
	return o.FirstNamePhonetics, true
}

// HasFirstNamePhonetics returns a boolean if a field has been set.
func (o *EmployeeResponse) HasFirstNamePhonetics() bool {
	if o != nil && !IsNil(o.FirstNamePhonetics) {
		return true
	}

	return false
}

// SetFirstNamePhonetics gets a reference to the given string and assigns it to the FirstNamePhonetics field.
func (o *EmployeeResponse) SetFirstNamePhonetics(v string) {
	o.FirstNamePhonetics = &v
}

// GetHiredDate returns the HiredDate field value if set, zero value otherwise.
func (o *EmployeeResponse) GetHiredDate() string {
	if o == nil || IsNil(o.HiredDate) {
		var ret string
		return ret
	}
	return *o.HiredDate
}

// GetHiredDateOk returns a tuple with the HiredDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeResponse) GetHiredDateOk() (*string, bool) {
	if o == nil || IsNil(o.HiredDate) {
		return nil, false
	}
	return o.HiredDate, true
}

// HasHiredDate returns a boolean if a field has been set.
func (o *EmployeeResponse) HasHiredDate() bool {
	if o != nil && !IsNil(o.HiredDate) {
		return true
	}

	return false
}

// SetHiredDate gets a reference to the given string and assigns it to the HiredDate field.
func (o *EmployeeResponse) SetHiredDate(v string) {
	o.HiredDate = &v
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise.
func (o *EmployeeResponse) GetBirthDate() string {
	if o == nil || IsNil(o.BirthDate) {
		var ret string
		return ret
	}
	return *o.BirthDate
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeResponse) GetBirthDateOk() (*string, bool) {
	if o == nil || IsNil(o.BirthDate) {
		return nil, false
	}
	return o.BirthDate, true
}

// HasBirthDate returns a boolean if a field has been set.
func (o *EmployeeResponse) HasBirthDate() bool {
	if o != nil && !IsNil(o.BirthDate) {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given string and assigns it to the BirthDate field.
func (o *EmployeeResponse) SetBirthDate(v string) {
	o.BirthDate = &v
}

// GetResignationDate returns the ResignationDate field value if set, zero value otherwise.
func (o *EmployeeResponse) GetResignationDate() string {
	if o == nil || IsNil(o.ResignationDate) {
		var ret string
		return ret
	}
	return *o.ResignationDate
}

// GetResignationDateOk returns a tuple with the ResignationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeResponse) GetResignationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ResignationDate) {
		return nil, false
	}
	return o.ResignationDate, true
}

// HasResignationDate returns a boolean if a field has been set.
func (o *EmployeeResponse) HasResignationDate() bool {
	if o != nil && !IsNil(o.ResignationDate) {
		return true
	}

	return false
}

// SetResignationDate gets a reference to the given string and assigns it to the ResignationDate field.
func (o *EmployeeResponse) SetResignationDate(v string) {
	o.ResignationDate = &v
}

// GetEmailAddresses returns the EmailAddresses field value if set, zero value otherwise.
func (o *EmployeeResponse) GetEmailAddresses() []string {
	if o == nil || IsNil(o.EmailAddresses) {
		var ret []string
		return ret
	}
	return o.EmailAddresses
}

// GetEmailAddressesOk returns a tuple with the EmailAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeResponse) GetEmailAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailAddresses) {
		return nil, false
	}
	return o.EmailAddresses, true
}

// HasEmailAddresses returns a boolean if a field has been set.
func (o *EmployeeResponse) HasEmailAddresses() bool {
	if o != nil && !IsNil(o.EmailAddresses) {
		return true
	}

	return false
}

// SetEmailAddresses gets a reference to the given []string and assigns it to the EmailAddresses field.
func (o *EmployeeResponse) SetEmailAddresses(v []string) {
	o.EmailAddresses = v
}

// GetAllDayRegardingWorkInMinute returns the AllDayRegardingWorkInMinute field value if set, zero value otherwise.
func (o *EmployeeResponse) GetAllDayRegardingWorkInMinute() int32 {
	if o == nil || IsNil(o.AllDayRegardingWorkInMinute) {
		var ret int32
		return ret
	}
	return *o.AllDayRegardingWorkInMinute
}

// GetAllDayRegardingWorkInMinuteOk returns a tuple with the AllDayRegardingWorkInMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeResponse) GetAllDayRegardingWorkInMinuteOk() (*int32, bool) {
	if o == nil || IsNil(o.AllDayRegardingWorkInMinute) {
		return nil, false
	}
	return o.AllDayRegardingWorkInMinute, true
}

// HasAllDayRegardingWorkInMinute returns a boolean if a field has been set.
func (o *EmployeeResponse) HasAllDayRegardingWorkInMinute() bool {
	if o != nil && !IsNil(o.AllDayRegardingWorkInMinute) {
		return true
	}

	return false
}

// SetAllDayRegardingWorkInMinute gets a reference to the given int32 and assigns it to the AllDayRegardingWorkInMinute field.
func (o *EmployeeResponse) SetAllDayRegardingWorkInMinute(v int32) {
	o.AllDayRegardingWorkInMinute = &v
}

func (o EmployeeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmployeeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["divisionCode"] = o.DivisionCode
	toSerialize["divisionName"] = o.DivisionName
	toSerialize["gender"] = o.Gender
	toSerialize["typeCode"] = o.TypeCode
	toSerialize["typeName"] = o.TypeName
	toSerialize["code"] = o.Code
	toSerialize["key"] = o.Key
	toSerialize["lastName"] = o.LastName
	toSerialize["firstName"] = o.FirstName
	toSerialize["employeeGroups"] = o.EmployeeGroups
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

type NullableEmployeeResponse struct {
	value *EmployeeResponse
	isSet bool
}

func (v NullableEmployeeResponse) Get() *EmployeeResponse {
	return v.value
}

func (v *NullableEmployeeResponse) Set(val *EmployeeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployeeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployeeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployeeResponse(val *EmployeeResponse) *NullableEmployeeResponse {
	return &NullableEmployeeResponse{value: val, isSet: true}
}

func (v NullableEmployeeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployeeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

