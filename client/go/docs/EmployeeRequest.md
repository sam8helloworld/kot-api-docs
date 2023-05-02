# EmployeeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DivisionCode** | **string** | 所属コード | 
**Gender** | **string** | 性別（male： 男性　female： 女性） | 
**TypeCode** | **string** | 雇用区分コード | 
**Code** | **string** | 従業員コード | 
**LastName** | **string** | 姓 | 
**FirstName** | **string** | 名 | 
**LastNamePhonetics** | Pointer to **string** | 姓（カナ） | [optional] 
**FirstNamePhonetics** | Pointer to **string** | 名（カナ） | [optional] 
**HiredDate** | Pointer to **string** | 入社年月日 | [optional] 
**BirthDate** | Pointer to **string** | 生年月日 | [optional] 
**ResignationDate** | Pointer to **string** | 退職年月日 | [optional] 
**EmailAddresses** | Pointer to **[]string** | メールアドレス | [optional] 
**AllDayRegardingWorkInMinute** | Pointer to **int32** | 日の契約労働時間 | [optional] 

## Methods

### NewEmployeeRequest

`func NewEmployeeRequest(divisionCode string, gender string, typeCode string, code string, lastName string, firstName string, ) *EmployeeRequest`

NewEmployeeRequest instantiates a new EmployeeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeRequestWithDefaults

`func NewEmployeeRequestWithDefaults() *EmployeeRequest`

NewEmployeeRequestWithDefaults instantiates a new EmployeeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDivisionCode

`func (o *EmployeeRequest) GetDivisionCode() string`

GetDivisionCode returns the DivisionCode field if non-nil, zero value otherwise.

### GetDivisionCodeOk

`func (o *EmployeeRequest) GetDivisionCodeOk() (*string, bool)`

GetDivisionCodeOk returns a tuple with the DivisionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionCode

`func (o *EmployeeRequest) SetDivisionCode(v string)`

SetDivisionCode sets DivisionCode field to given value.


### GetGender

`func (o *EmployeeRequest) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *EmployeeRequest) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *EmployeeRequest) SetGender(v string)`

SetGender sets Gender field to given value.


### GetTypeCode

`func (o *EmployeeRequest) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *EmployeeRequest) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *EmployeeRequest) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.


### GetCode

`func (o *EmployeeRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EmployeeRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EmployeeRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetLastName

`func (o *EmployeeRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *EmployeeRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *EmployeeRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetFirstName

`func (o *EmployeeRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *EmployeeRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *EmployeeRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastNamePhonetics

`func (o *EmployeeRequest) GetLastNamePhonetics() string`

GetLastNamePhonetics returns the LastNamePhonetics field if non-nil, zero value otherwise.

### GetLastNamePhoneticsOk

`func (o *EmployeeRequest) GetLastNamePhoneticsOk() (*string, bool)`

GetLastNamePhoneticsOk returns a tuple with the LastNamePhonetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNamePhonetics

`func (o *EmployeeRequest) SetLastNamePhonetics(v string)`

SetLastNamePhonetics sets LastNamePhonetics field to given value.

### HasLastNamePhonetics

`func (o *EmployeeRequest) HasLastNamePhonetics() bool`

HasLastNamePhonetics returns a boolean if a field has been set.

### GetFirstNamePhonetics

`func (o *EmployeeRequest) GetFirstNamePhonetics() string`

GetFirstNamePhonetics returns the FirstNamePhonetics field if non-nil, zero value otherwise.

### GetFirstNamePhoneticsOk

`func (o *EmployeeRequest) GetFirstNamePhoneticsOk() (*string, bool)`

GetFirstNamePhoneticsOk returns a tuple with the FirstNamePhonetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNamePhonetics

`func (o *EmployeeRequest) SetFirstNamePhonetics(v string)`

SetFirstNamePhonetics sets FirstNamePhonetics field to given value.

### HasFirstNamePhonetics

`func (o *EmployeeRequest) HasFirstNamePhonetics() bool`

HasFirstNamePhonetics returns a boolean if a field has been set.

### GetHiredDate

`func (o *EmployeeRequest) GetHiredDate() string`

GetHiredDate returns the HiredDate field if non-nil, zero value otherwise.

### GetHiredDateOk

`func (o *EmployeeRequest) GetHiredDateOk() (*string, bool)`

GetHiredDateOk returns a tuple with the HiredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiredDate

`func (o *EmployeeRequest) SetHiredDate(v string)`

SetHiredDate sets HiredDate field to given value.

### HasHiredDate

`func (o *EmployeeRequest) HasHiredDate() bool`

HasHiredDate returns a boolean if a field has been set.

### GetBirthDate

`func (o *EmployeeRequest) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *EmployeeRequest) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *EmployeeRequest) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *EmployeeRequest) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetResignationDate

`func (o *EmployeeRequest) GetResignationDate() string`

GetResignationDate returns the ResignationDate field if non-nil, zero value otherwise.

### GetResignationDateOk

`func (o *EmployeeRequest) GetResignationDateOk() (*string, bool)`

GetResignationDateOk returns a tuple with the ResignationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResignationDate

`func (o *EmployeeRequest) SetResignationDate(v string)`

SetResignationDate sets ResignationDate field to given value.

### HasResignationDate

`func (o *EmployeeRequest) HasResignationDate() bool`

HasResignationDate returns a boolean if a field has been set.

### GetEmailAddresses

`func (o *EmployeeRequest) GetEmailAddresses() []string`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *EmployeeRequest) GetEmailAddressesOk() (*[]string, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *EmployeeRequest) SetEmailAddresses(v []string)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *EmployeeRequest) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### GetAllDayRegardingWorkInMinute

`func (o *EmployeeRequest) GetAllDayRegardingWorkInMinute() int32`

GetAllDayRegardingWorkInMinute returns the AllDayRegardingWorkInMinute field if non-nil, zero value otherwise.

### GetAllDayRegardingWorkInMinuteOk

`func (o *EmployeeRequest) GetAllDayRegardingWorkInMinuteOk() (*int32, bool)`

GetAllDayRegardingWorkInMinuteOk returns a tuple with the AllDayRegardingWorkInMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDayRegardingWorkInMinute

`func (o *EmployeeRequest) SetAllDayRegardingWorkInMinute(v int32)`

SetAllDayRegardingWorkInMinute sets AllDayRegardingWorkInMinute field to given value.

### HasAllDayRegardingWorkInMinute

`func (o *EmployeeRequest) HasAllDayRegardingWorkInMinute() bool`

HasAllDayRegardingWorkInMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


