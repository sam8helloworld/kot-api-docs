# EmployeeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DivisionCode** | **string** | 所属コード | 
**DivisionName** | **string** | 所属名 | 
**Gender** | **string** | 性別（male： 男性　female： 女性） | 
**TypeCode** | **string** | 雇用区分コード | 
**TypeName** | **string** | 雇用区分名 | 
**Code** | **string** | 従業員コード | 
**Key** | **string** | 従業員識別キー（従業員コードが変更されても不変） | 
**LastName** | **string** | 姓 | 
**FirstName** | **string** | 名 | 
**EmployeeGroups** | [**[]EmployeeResponseEmployeeGroupsInner**](EmployeeResponseEmployeeGroupsInner.md) | 従業員グループ情報 | 
**LastNamePhonetics** | Pointer to **string** | 姓（カナ） | [optional] 
**FirstNamePhonetics** | Pointer to **string** | 名（カナ） | [optional] 
**HiredDate** | Pointer to **string** | 入社年月日 | [optional] 
**BirthDate** | Pointer to **string** | 生年月日 | [optional] 
**ResignationDate** | Pointer to **string** | 退職年月日 | [optional] 
**EmailAddresses** | Pointer to **[]string** | メールアドレス | [optional] 
**AllDayRegardingWorkInMinute** | Pointer to **int32** | 日の契約労働時間 | [optional] 

## Methods

### NewEmployeeResponse

`func NewEmployeeResponse(divisionCode string, divisionName string, gender string, typeCode string, typeName string, code string, key string, lastName string, firstName string, employeeGroups []EmployeeResponseEmployeeGroupsInner, ) *EmployeeResponse`

NewEmployeeResponse instantiates a new EmployeeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeResponseWithDefaults

`func NewEmployeeResponseWithDefaults() *EmployeeResponse`

NewEmployeeResponseWithDefaults instantiates a new EmployeeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDivisionCode

`func (o *EmployeeResponse) GetDivisionCode() string`

GetDivisionCode returns the DivisionCode field if non-nil, zero value otherwise.

### GetDivisionCodeOk

`func (o *EmployeeResponse) GetDivisionCodeOk() (*string, bool)`

GetDivisionCodeOk returns a tuple with the DivisionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionCode

`func (o *EmployeeResponse) SetDivisionCode(v string)`

SetDivisionCode sets DivisionCode field to given value.


### GetDivisionName

`func (o *EmployeeResponse) GetDivisionName() string`

GetDivisionName returns the DivisionName field if non-nil, zero value otherwise.

### GetDivisionNameOk

`func (o *EmployeeResponse) GetDivisionNameOk() (*string, bool)`

GetDivisionNameOk returns a tuple with the DivisionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionName

`func (o *EmployeeResponse) SetDivisionName(v string)`

SetDivisionName sets DivisionName field to given value.


### GetGender

`func (o *EmployeeResponse) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *EmployeeResponse) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *EmployeeResponse) SetGender(v string)`

SetGender sets Gender field to given value.


### GetTypeCode

`func (o *EmployeeResponse) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *EmployeeResponse) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *EmployeeResponse) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.


### GetTypeName

`func (o *EmployeeResponse) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *EmployeeResponse) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *EmployeeResponse) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetCode

`func (o *EmployeeResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EmployeeResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EmployeeResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetKey

`func (o *EmployeeResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EmployeeResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EmployeeResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetLastName

`func (o *EmployeeResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *EmployeeResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *EmployeeResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetFirstName

`func (o *EmployeeResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *EmployeeResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *EmployeeResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetEmployeeGroups

`func (o *EmployeeResponse) GetEmployeeGroups() []EmployeeResponseEmployeeGroupsInner`

GetEmployeeGroups returns the EmployeeGroups field if non-nil, zero value otherwise.

### GetEmployeeGroupsOk

`func (o *EmployeeResponse) GetEmployeeGroupsOk() (*[]EmployeeResponseEmployeeGroupsInner, bool)`

GetEmployeeGroupsOk returns a tuple with the EmployeeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeGroups

`func (o *EmployeeResponse) SetEmployeeGroups(v []EmployeeResponseEmployeeGroupsInner)`

SetEmployeeGroups sets EmployeeGroups field to given value.


### GetLastNamePhonetics

`func (o *EmployeeResponse) GetLastNamePhonetics() string`

GetLastNamePhonetics returns the LastNamePhonetics field if non-nil, zero value otherwise.

### GetLastNamePhoneticsOk

`func (o *EmployeeResponse) GetLastNamePhoneticsOk() (*string, bool)`

GetLastNamePhoneticsOk returns a tuple with the LastNamePhonetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNamePhonetics

`func (o *EmployeeResponse) SetLastNamePhonetics(v string)`

SetLastNamePhonetics sets LastNamePhonetics field to given value.

### HasLastNamePhonetics

`func (o *EmployeeResponse) HasLastNamePhonetics() bool`

HasLastNamePhonetics returns a boolean if a field has been set.

### GetFirstNamePhonetics

`func (o *EmployeeResponse) GetFirstNamePhonetics() string`

GetFirstNamePhonetics returns the FirstNamePhonetics field if non-nil, zero value otherwise.

### GetFirstNamePhoneticsOk

`func (o *EmployeeResponse) GetFirstNamePhoneticsOk() (*string, bool)`

GetFirstNamePhoneticsOk returns a tuple with the FirstNamePhonetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNamePhonetics

`func (o *EmployeeResponse) SetFirstNamePhonetics(v string)`

SetFirstNamePhonetics sets FirstNamePhonetics field to given value.

### HasFirstNamePhonetics

`func (o *EmployeeResponse) HasFirstNamePhonetics() bool`

HasFirstNamePhonetics returns a boolean if a field has been set.

### GetHiredDate

`func (o *EmployeeResponse) GetHiredDate() string`

GetHiredDate returns the HiredDate field if non-nil, zero value otherwise.

### GetHiredDateOk

`func (o *EmployeeResponse) GetHiredDateOk() (*string, bool)`

GetHiredDateOk returns a tuple with the HiredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiredDate

`func (o *EmployeeResponse) SetHiredDate(v string)`

SetHiredDate sets HiredDate field to given value.

### HasHiredDate

`func (o *EmployeeResponse) HasHiredDate() bool`

HasHiredDate returns a boolean if a field has been set.

### GetBirthDate

`func (o *EmployeeResponse) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *EmployeeResponse) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *EmployeeResponse) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *EmployeeResponse) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetResignationDate

`func (o *EmployeeResponse) GetResignationDate() string`

GetResignationDate returns the ResignationDate field if non-nil, zero value otherwise.

### GetResignationDateOk

`func (o *EmployeeResponse) GetResignationDateOk() (*string, bool)`

GetResignationDateOk returns a tuple with the ResignationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResignationDate

`func (o *EmployeeResponse) SetResignationDate(v string)`

SetResignationDate sets ResignationDate field to given value.

### HasResignationDate

`func (o *EmployeeResponse) HasResignationDate() bool`

HasResignationDate returns a boolean if a field has been set.

### GetEmailAddresses

`func (o *EmployeeResponse) GetEmailAddresses() []string`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *EmployeeResponse) GetEmailAddressesOk() (*[]string, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *EmployeeResponse) SetEmailAddresses(v []string)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *EmployeeResponse) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### GetAllDayRegardingWorkInMinute

`func (o *EmployeeResponse) GetAllDayRegardingWorkInMinute() int32`

GetAllDayRegardingWorkInMinute returns the AllDayRegardingWorkInMinute field if non-nil, zero value otherwise.

### GetAllDayRegardingWorkInMinuteOk

`func (o *EmployeeResponse) GetAllDayRegardingWorkInMinuteOk() (*int32, bool)`

GetAllDayRegardingWorkInMinuteOk returns a tuple with the AllDayRegardingWorkInMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDayRegardingWorkInMinute

`func (o *EmployeeResponse) SetAllDayRegardingWorkInMinute(v int32)`

SetAllDayRegardingWorkInMinute sets AllDayRegardingWorkInMinute field to given value.

### HasAllDayRegardingWorkInMinute

`func (o *EmployeeResponse) HasAllDayRegardingWorkInMinute() bool`

HasAllDayRegardingWorkInMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


