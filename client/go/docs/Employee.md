# Employee

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
**EmployeeGroups** | [**[]EmployeeEmployeeGroupsInner**](EmployeeEmployeeGroupsInner.md) | 従業員グループ情報 | 
**LastNamePhonetics** | Pointer to **string** | 姓（カナ） | [optional] 
**FirstNamePhonetics** | Pointer to **string** | 名（カナ） | [optional] 
**HiredDate** | Pointer to **string** | 入社年月日 | [optional] 
**BirthDate** | Pointer to **string** | 生年月日 | [optional] 
**ResignationDate** | Pointer to **string** | 退職年月日 | [optional] 
**EmailAddresses** | Pointer to **[]string** | メールアドレス | [optional] 
**AllDayRegardingWorkInMinute** | Pointer to **int32** | 日の契約労働時間 | [optional] 

## Methods

### NewEmployee

`func NewEmployee(divisionCode string, divisionName string, gender string, typeCode string, typeName string, code string, key string, lastName string, firstName string, employeeGroups []EmployeeEmployeeGroupsInner, ) *Employee`

NewEmployee instantiates a new Employee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeWithDefaults

`func NewEmployeeWithDefaults() *Employee`

NewEmployeeWithDefaults instantiates a new Employee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDivisionCode

`func (o *Employee) GetDivisionCode() string`

GetDivisionCode returns the DivisionCode field if non-nil, zero value otherwise.

### GetDivisionCodeOk

`func (o *Employee) GetDivisionCodeOk() (*string, bool)`

GetDivisionCodeOk returns a tuple with the DivisionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionCode

`func (o *Employee) SetDivisionCode(v string)`

SetDivisionCode sets DivisionCode field to given value.


### GetDivisionName

`func (o *Employee) GetDivisionName() string`

GetDivisionName returns the DivisionName field if non-nil, zero value otherwise.

### GetDivisionNameOk

`func (o *Employee) GetDivisionNameOk() (*string, bool)`

GetDivisionNameOk returns a tuple with the DivisionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionName

`func (o *Employee) SetDivisionName(v string)`

SetDivisionName sets DivisionName field to given value.


### GetGender

`func (o *Employee) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *Employee) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *Employee) SetGender(v string)`

SetGender sets Gender field to given value.


### GetTypeCode

`func (o *Employee) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *Employee) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *Employee) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.


### GetTypeName

`func (o *Employee) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *Employee) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *Employee) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetCode

`func (o *Employee) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Employee) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Employee) SetCode(v string)`

SetCode sets Code field to given value.


### GetKey

`func (o *Employee) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Employee) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Employee) SetKey(v string)`

SetKey sets Key field to given value.


### GetLastName

`func (o *Employee) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Employee) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Employee) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetFirstName

`func (o *Employee) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Employee) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Employee) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetEmployeeGroups

`func (o *Employee) GetEmployeeGroups() []EmployeeEmployeeGroupsInner`

GetEmployeeGroups returns the EmployeeGroups field if non-nil, zero value otherwise.

### GetEmployeeGroupsOk

`func (o *Employee) GetEmployeeGroupsOk() (*[]EmployeeEmployeeGroupsInner, bool)`

GetEmployeeGroupsOk returns a tuple with the EmployeeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeGroups

`func (o *Employee) SetEmployeeGroups(v []EmployeeEmployeeGroupsInner)`

SetEmployeeGroups sets EmployeeGroups field to given value.


### GetLastNamePhonetics

`func (o *Employee) GetLastNamePhonetics() string`

GetLastNamePhonetics returns the LastNamePhonetics field if non-nil, zero value otherwise.

### GetLastNamePhoneticsOk

`func (o *Employee) GetLastNamePhoneticsOk() (*string, bool)`

GetLastNamePhoneticsOk returns a tuple with the LastNamePhonetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNamePhonetics

`func (o *Employee) SetLastNamePhonetics(v string)`

SetLastNamePhonetics sets LastNamePhonetics field to given value.

### HasLastNamePhonetics

`func (o *Employee) HasLastNamePhonetics() bool`

HasLastNamePhonetics returns a boolean if a field has been set.

### GetFirstNamePhonetics

`func (o *Employee) GetFirstNamePhonetics() string`

GetFirstNamePhonetics returns the FirstNamePhonetics field if non-nil, zero value otherwise.

### GetFirstNamePhoneticsOk

`func (o *Employee) GetFirstNamePhoneticsOk() (*string, bool)`

GetFirstNamePhoneticsOk returns a tuple with the FirstNamePhonetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNamePhonetics

`func (o *Employee) SetFirstNamePhonetics(v string)`

SetFirstNamePhonetics sets FirstNamePhonetics field to given value.

### HasFirstNamePhonetics

`func (o *Employee) HasFirstNamePhonetics() bool`

HasFirstNamePhonetics returns a boolean if a field has been set.

### GetHiredDate

`func (o *Employee) GetHiredDate() string`

GetHiredDate returns the HiredDate field if non-nil, zero value otherwise.

### GetHiredDateOk

`func (o *Employee) GetHiredDateOk() (*string, bool)`

GetHiredDateOk returns a tuple with the HiredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiredDate

`func (o *Employee) SetHiredDate(v string)`

SetHiredDate sets HiredDate field to given value.

### HasHiredDate

`func (o *Employee) HasHiredDate() bool`

HasHiredDate returns a boolean if a field has been set.

### GetBirthDate

`func (o *Employee) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *Employee) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *Employee) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *Employee) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetResignationDate

`func (o *Employee) GetResignationDate() string`

GetResignationDate returns the ResignationDate field if non-nil, zero value otherwise.

### GetResignationDateOk

`func (o *Employee) GetResignationDateOk() (*string, bool)`

GetResignationDateOk returns a tuple with the ResignationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResignationDate

`func (o *Employee) SetResignationDate(v string)`

SetResignationDate sets ResignationDate field to given value.

### HasResignationDate

`func (o *Employee) HasResignationDate() bool`

HasResignationDate returns a boolean if a field has been set.

### GetEmailAddresses

`func (o *Employee) GetEmailAddresses() []string`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *Employee) GetEmailAddressesOk() (*[]string, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *Employee) SetEmailAddresses(v []string)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *Employee) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### GetAllDayRegardingWorkInMinute

`func (o *Employee) GetAllDayRegardingWorkInMinute() int32`

GetAllDayRegardingWorkInMinute returns the AllDayRegardingWorkInMinute field if non-nil, zero value otherwise.

### GetAllDayRegardingWorkInMinuteOk

`func (o *Employee) GetAllDayRegardingWorkInMinuteOk() (*int32, bool)`

GetAllDayRegardingWorkInMinuteOk returns a tuple with the AllDayRegardingWorkInMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDayRegardingWorkInMinute

`func (o *Employee) SetAllDayRegardingWorkInMinute(v int32)`

SetAllDayRegardingWorkInMinute sets AllDayRegardingWorkInMinute field to given value.

### HasAllDayRegardingWorkInMinute

`func (o *Employee) HasAllDayRegardingWorkInMinute() bool`

HasAllDayRegardingWorkInMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


