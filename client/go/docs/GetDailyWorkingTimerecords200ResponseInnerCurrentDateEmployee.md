# GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee

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
**EmployeeGroups** | [**[]GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner**](GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner.md) | 従業員グループ情報 | 

## Methods

### NewGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee

`func NewGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee(divisionCode string, divisionName string, gender string, typeCode string, typeName string, code string, lastName string, firstName string, lastNamePhonetics string, firstNamePhonetics string, employeeGroups []GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner, ) *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee`

NewGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee instantiates a new GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeWithDefaults

`func NewGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeWithDefaults() *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee`

NewGetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeWithDefaults instantiates a new GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDivisionCode

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetDivisionCode() string`

GetDivisionCode returns the DivisionCode field if non-nil, zero value otherwise.

### GetDivisionCodeOk

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetDivisionCodeOk() (*string, bool)`

GetDivisionCodeOk returns a tuple with the DivisionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionCode

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) SetDivisionCode(v string)`

SetDivisionCode sets DivisionCode field to given value.


### GetDivisionName

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetDivisionName() string`

GetDivisionName returns the DivisionName field if non-nil, zero value otherwise.

### GetDivisionNameOk

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetDivisionNameOk() (*string, bool)`

GetDivisionNameOk returns a tuple with the DivisionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionName

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) SetDivisionName(v string)`

SetDivisionName sets DivisionName field to given value.


### GetGender

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) SetGender(v string)`

SetGender sets Gender field to given value.


### GetTypeCode

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.


### GetTypeName

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetCode

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) SetCode(v string)`

SetCode sets Code field to given value.


### GetLastName

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetFirstName

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastNamePhonetics

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetLastNamePhonetics() string`

GetLastNamePhonetics returns the LastNamePhonetics field if non-nil, zero value otherwise.

### GetLastNamePhoneticsOk

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetLastNamePhoneticsOk() (*string, bool)`

GetLastNamePhoneticsOk returns a tuple with the LastNamePhonetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNamePhonetics

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) SetLastNamePhonetics(v string)`

SetLastNamePhonetics sets LastNamePhonetics field to given value.


### GetFirstNamePhonetics

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetFirstNamePhonetics() string`

GetFirstNamePhonetics returns the FirstNamePhonetics field if non-nil, zero value otherwise.

### GetFirstNamePhoneticsOk

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetFirstNamePhoneticsOk() (*string, bool)`

GetFirstNamePhoneticsOk returns a tuple with the FirstNamePhonetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNamePhonetics

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) SetFirstNamePhonetics(v string)`

SetFirstNamePhonetics sets FirstNamePhonetics field to given value.


### GetEmployeeGroups

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetEmployeeGroups() []GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner`

GetEmployeeGroups returns the EmployeeGroups field if non-nil, zero value otherwise.

### GetEmployeeGroupsOk

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) GetEmployeeGroupsOk() (*[]GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner, bool)`

GetEmployeeGroupsOk returns a tuple with the EmployeeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeGroups

`func (o *GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployee) SetEmployeeGroups(v []GetDailyWorkingTimerecords200ResponseInnerCurrentDateEmployeeEmployeeGroupsInner)`

SetEmployeeGroups sets EmployeeGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


