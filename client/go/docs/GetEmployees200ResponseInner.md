# GetEmployees200ResponseInner

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
**EmployeeGroups** | [**[]GetEmployees200ResponseInnerEmployeeGroupsInner**](GetEmployees200ResponseInnerEmployeeGroupsInner.md) | 従業員グループ情報 | 
**LastNamePhonetics** | Pointer to **string** | 姓（カナ） | [optional] 
**FirstNamePhonetics** | Pointer to **string** | 名（カナ） | [optional] 
**HiredDate** | Pointer to **string** | 入社年月日 | [optional] 
**BirthDate** | Pointer to **string** | 生年月日 | [optional] 
**ResignationDate** | Pointer to **string** | 退職年月日 | [optional] 
**EmailAddresses** | Pointer to **[]string** | メールアドレス | [optional] 
**AllDayRegardingWorkInMinute** | Pointer to **int32** | 日の契約労働時間 | [optional] 

## Methods

### NewGetEmployees200ResponseInner

`func NewGetEmployees200ResponseInner(divisionCode string, divisionName string, gender string, typeCode string, typeName string, code string, key string, lastName string, firstName string, employeeGroups []GetEmployees200ResponseInnerEmployeeGroupsInner, ) *GetEmployees200ResponseInner`

NewGetEmployees200ResponseInner instantiates a new GetEmployees200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEmployees200ResponseInnerWithDefaults

`func NewGetEmployees200ResponseInnerWithDefaults() *GetEmployees200ResponseInner`

NewGetEmployees200ResponseInnerWithDefaults instantiates a new GetEmployees200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDivisionCode

`func (o *GetEmployees200ResponseInner) GetDivisionCode() string`

GetDivisionCode returns the DivisionCode field if non-nil, zero value otherwise.

### GetDivisionCodeOk

`func (o *GetEmployees200ResponseInner) GetDivisionCodeOk() (*string, bool)`

GetDivisionCodeOk returns a tuple with the DivisionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionCode

`func (o *GetEmployees200ResponseInner) SetDivisionCode(v string)`

SetDivisionCode sets DivisionCode field to given value.


### GetDivisionName

`func (o *GetEmployees200ResponseInner) GetDivisionName() string`

GetDivisionName returns the DivisionName field if non-nil, zero value otherwise.

### GetDivisionNameOk

`func (o *GetEmployees200ResponseInner) GetDivisionNameOk() (*string, bool)`

GetDivisionNameOk returns a tuple with the DivisionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionName

`func (o *GetEmployees200ResponseInner) SetDivisionName(v string)`

SetDivisionName sets DivisionName field to given value.


### GetGender

`func (o *GetEmployees200ResponseInner) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *GetEmployees200ResponseInner) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *GetEmployees200ResponseInner) SetGender(v string)`

SetGender sets Gender field to given value.


### GetTypeCode

`func (o *GetEmployees200ResponseInner) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *GetEmployees200ResponseInner) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *GetEmployees200ResponseInner) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.


### GetTypeName

`func (o *GetEmployees200ResponseInner) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *GetEmployees200ResponseInner) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *GetEmployees200ResponseInner) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetCode

`func (o *GetEmployees200ResponseInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetEmployees200ResponseInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetEmployees200ResponseInner) SetCode(v string)`

SetCode sets Code field to given value.


### GetKey

`func (o *GetEmployees200ResponseInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetEmployees200ResponseInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetEmployees200ResponseInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetLastName

`func (o *GetEmployees200ResponseInner) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *GetEmployees200ResponseInner) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *GetEmployees200ResponseInner) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetFirstName

`func (o *GetEmployees200ResponseInner) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *GetEmployees200ResponseInner) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *GetEmployees200ResponseInner) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetEmployeeGroups

`func (o *GetEmployees200ResponseInner) GetEmployeeGroups() []GetEmployees200ResponseInnerEmployeeGroupsInner`

GetEmployeeGroups returns the EmployeeGroups field if non-nil, zero value otherwise.

### GetEmployeeGroupsOk

`func (o *GetEmployees200ResponseInner) GetEmployeeGroupsOk() (*[]GetEmployees200ResponseInnerEmployeeGroupsInner, bool)`

GetEmployeeGroupsOk returns a tuple with the EmployeeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeGroups

`func (o *GetEmployees200ResponseInner) SetEmployeeGroups(v []GetEmployees200ResponseInnerEmployeeGroupsInner)`

SetEmployeeGroups sets EmployeeGroups field to given value.


### GetLastNamePhonetics

`func (o *GetEmployees200ResponseInner) GetLastNamePhonetics() string`

GetLastNamePhonetics returns the LastNamePhonetics field if non-nil, zero value otherwise.

### GetLastNamePhoneticsOk

`func (o *GetEmployees200ResponseInner) GetLastNamePhoneticsOk() (*string, bool)`

GetLastNamePhoneticsOk returns a tuple with the LastNamePhonetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNamePhonetics

`func (o *GetEmployees200ResponseInner) SetLastNamePhonetics(v string)`

SetLastNamePhonetics sets LastNamePhonetics field to given value.

### HasLastNamePhonetics

`func (o *GetEmployees200ResponseInner) HasLastNamePhonetics() bool`

HasLastNamePhonetics returns a boolean if a field has been set.

### GetFirstNamePhonetics

`func (o *GetEmployees200ResponseInner) GetFirstNamePhonetics() string`

GetFirstNamePhonetics returns the FirstNamePhonetics field if non-nil, zero value otherwise.

### GetFirstNamePhoneticsOk

`func (o *GetEmployees200ResponseInner) GetFirstNamePhoneticsOk() (*string, bool)`

GetFirstNamePhoneticsOk returns a tuple with the FirstNamePhonetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNamePhonetics

`func (o *GetEmployees200ResponseInner) SetFirstNamePhonetics(v string)`

SetFirstNamePhonetics sets FirstNamePhonetics field to given value.

### HasFirstNamePhonetics

`func (o *GetEmployees200ResponseInner) HasFirstNamePhonetics() bool`

HasFirstNamePhonetics returns a boolean if a field has been set.

### GetHiredDate

`func (o *GetEmployees200ResponseInner) GetHiredDate() string`

GetHiredDate returns the HiredDate field if non-nil, zero value otherwise.

### GetHiredDateOk

`func (o *GetEmployees200ResponseInner) GetHiredDateOk() (*string, bool)`

GetHiredDateOk returns a tuple with the HiredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiredDate

`func (o *GetEmployees200ResponseInner) SetHiredDate(v string)`

SetHiredDate sets HiredDate field to given value.

### HasHiredDate

`func (o *GetEmployees200ResponseInner) HasHiredDate() bool`

HasHiredDate returns a boolean if a field has been set.

### GetBirthDate

`func (o *GetEmployees200ResponseInner) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *GetEmployees200ResponseInner) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *GetEmployees200ResponseInner) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *GetEmployees200ResponseInner) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetResignationDate

`func (o *GetEmployees200ResponseInner) GetResignationDate() string`

GetResignationDate returns the ResignationDate field if non-nil, zero value otherwise.

### GetResignationDateOk

`func (o *GetEmployees200ResponseInner) GetResignationDateOk() (*string, bool)`

GetResignationDateOk returns a tuple with the ResignationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResignationDate

`func (o *GetEmployees200ResponseInner) SetResignationDate(v string)`

SetResignationDate sets ResignationDate field to given value.

### HasResignationDate

`func (o *GetEmployees200ResponseInner) HasResignationDate() bool`

HasResignationDate returns a boolean if a field has been set.

### GetEmailAddresses

`func (o *GetEmployees200ResponseInner) GetEmailAddresses() []string`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *GetEmployees200ResponseInner) GetEmailAddressesOk() (*[]string, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *GetEmployees200ResponseInner) SetEmailAddresses(v []string)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *GetEmployees200ResponseInner) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### GetAllDayRegardingWorkInMinute

`func (o *GetEmployees200ResponseInner) GetAllDayRegardingWorkInMinute() int32`

GetAllDayRegardingWorkInMinute returns the AllDayRegardingWorkInMinute field if non-nil, zero value otherwise.

### GetAllDayRegardingWorkInMinuteOk

`func (o *GetEmployees200ResponseInner) GetAllDayRegardingWorkInMinuteOk() (*int32, bool)`

GetAllDayRegardingWorkInMinuteOk returns a tuple with the AllDayRegardingWorkInMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDayRegardingWorkInMinute

`func (o *GetEmployees200ResponseInner) SetAllDayRegardingWorkInMinute(v int32)`

SetAllDayRegardingWorkInMinute sets AllDayRegardingWorkInMinute field to given value.

### HasAllDayRegardingWorkInMinute

`func (o *GetEmployees200ResponseInner) HasAllDayRegardingWorkInMinute() bool`

HasAllDayRegardingWorkInMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


