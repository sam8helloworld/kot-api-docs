# DailyWorkingResponseCurrentDateEmployee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DivisionCode** | **string** | 所属コード | 
**DivisionName** | **string** | 所属名 | 
**Gender** | **string** | 性別（no_selected： 選択しない　male： 男性　female： 女性） | 
**TypeCode** | **string** | 雇用区分コード | 
**TypeName** | **string** | 雇用区分名 | 
**Code** | **string** | 従業員コード | 
**LastName** | **string** | 姓 | 
**FirstName** | **string** | 名 | 
**LastNamePhonetics** | **string** | 姓（カナ） | 
**FirstNamePhonetics** | **string** | 名（カナ） | 
**EmployeeGroups** | [**[]EmployeeResponseEmployeeGroupsInner**](EmployeeResponseEmployeeGroupsInner.md) | 従業員グループ情報 | 

## Methods

### NewDailyWorkingResponseCurrentDateEmployee

`func NewDailyWorkingResponseCurrentDateEmployee(divisionCode string, divisionName string, gender string, typeCode string, typeName string, code string, lastName string, firstName string, lastNamePhonetics string, firstNamePhonetics string, employeeGroups []EmployeeResponseEmployeeGroupsInner, ) *DailyWorkingResponseCurrentDateEmployee`

NewDailyWorkingResponseCurrentDateEmployee instantiates a new DailyWorkingResponseCurrentDateEmployee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyWorkingResponseCurrentDateEmployeeWithDefaults

`func NewDailyWorkingResponseCurrentDateEmployeeWithDefaults() *DailyWorkingResponseCurrentDateEmployee`

NewDailyWorkingResponseCurrentDateEmployeeWithDefaults instantiates a new DailyWorkingResponseCurrentDateEmployee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDivisionCode

`func (o *DailyWorkingResponseCurrentDateEmployee) GetDivisionCode() string`

GetDivisionCode returns the DivisionCode field if non-nil, zero value otherwise.

### GetDivisionCodeOk

`func (o *DailyWorkingResponseCurrentDateEmployee) GetDivisionCodeOk() (*string, bool)`

GetDivisionCodeOk returns a tuple with the DivisionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionCode

`func (o *DailyWorkingResponseCurrentDateEmployee) SetDivisionCode(v string)`

SetDivisionCode sets DivisionCode field to given value.


### GetDivisionName

`func (o *DailyWorkingResponseCurrentDateEmployee) GetDivisionName() string`

GetDivisionName returns the DivisionName field if non-nil, zero value otherwise.

### GetDivisionNameOk

`func (o *DailyWorkingResponseCurrentDateEmployee) GetDivisionNameOk() (*string, bool)`

GetDivisionNameOk returns a tuple with the DivisionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionName

`func (o *DailyWorkingResponseCurrentDateEmployee) SetDivisionName(v string)`

SetDivisionName sets DivisionName field to given value.


### GetGender

`func (o *DailyWorkingResponseCurrentDateEmployee) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *DailyWorkingResponseCurrentDateEmployee) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *DailyWorkingResponseCurrentDateEmployee) SetGender(v string)`

SetGender sets Gender field to given value.


### GetTypeCode

`func (o *DailyWorkingResponseCurrentDateEmployee) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *DailyWorkingResponseCurrentDateEmployee) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *DailyWorkingResponseCurrentDateEmployee) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.


### GetTypeName

`func (o *DailyWorkingResponseCurrentDateEmployee) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *DailyWorkingResponseCurrentDateEmployee) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *DailyWorkingResponseCurrentDateEmployee) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetCode

`func (o *DailyWorkingResponseCurrentDateEmployee) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DailyWorkingResponseCurrentDateEmployee) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DailyWorkingResponseCurrentDateEmployee) SetCode(v string)`

SetCode sets Code field to given value.


### GetLastName

`func (o *DailyWorkingResponseCurrentDateEmployee) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *DailyWorkingResponseCurrentDateEmployee) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *DailyWorkingResponseCurrentDateEmployee) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetFirstName

`func (o *DailyWorkingResponseCurrentDateEmployee) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *DailyWorkingResponseCurrentDateEmployee) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *DailyWorkingResponseCurrentDateEmployee) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastNamePhonetics

`func (o *DailyWorkingResponseCurrentDateEmployee) GetLastNamePhonetics() string`

GetLastNamePhonetics returns the LastNamePhonetics field if non-nil, zero value otherwise.

### GetLastNamePhoneticsOk

`func (o *DailyWorkingResponseCurrentDateEmployee) GetLastNamePhoneticsOk() (*string, bool)`

GetLastNamePhoneticsOk returns a tuple with the LastNamePhonetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNamePhonetics

`func (o *DailyWorkingResponseCurrentDateEmployee) SetLastNamePhonetics(v string)`

SetLastNamePhonetics sets LastNamePhonetics field to given value.


### GetFirstNamePhonetics

`func (o *DailyWorkingResponseCurrentDateEmployee) GetFirstNamePhonetics() string`

GetFirstNamePhonetics returns the FirstNamePhonetics field if non-nil, zero value otherwise.

### GetFirstNamePhoneticsOk

`func (o *DailyWorkingResponseCurrentDateEmployee) GetFirstNamePhoneticsOk() (*string, bool)`

GetFirstNamePhoneticsOk returns a tuple with the FirstNamePhonetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNamePhonetics

`func (o *DailyWorkingResponseCurrentDateEmployee) SetFirstNamePhonetics(v string)`

SetFirstNamePhonetics sets FirstNamePhonetics field to given value.


### GetEmployeeGroups

`func (o *DailyWorkingResponseCurrentDateEmployee) GetEmployeeGroups() []EmployeeResponseEmployeeGroupsInner`

GetEmployeeGroups returns the EmployeeGroups field if non-nil, zero value otherwise.

### GetEmployeeGroupsOk

`func (o *DailyWorkingResponseCurrentDateEmployee) GetEmployeeGroupsOk() (*[]EmployeeResponseEmployeeGroupsInner, bool)`

GetEmployeeGroupsOk returns a tuple with the EmployeeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeGroups

`func (o *DailyWorkingResponseCurrentDateEmployee) SetEmployeeGroups(v []EmployeeResponseEmployeeGroupsInner)`

SetEmployeeGroups sets EmployeeGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


