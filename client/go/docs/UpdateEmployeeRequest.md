# UpdateEmployeeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DivisionCode** | Pointer to **string** | 所属コード | [optional] 
**Gender** | Pointer to **string** | 性別（male： 男性　female： 女性） | [optional] 
**TypeCode** | Pointer to **string** | 雇用区分コード | [optional] 
**Code** | Pointer to **string** | 従業員コード | [optional] 
**LastName** | Pointer to **string** | 姓 | [optional] 
**FirstName** | Pointer to **string** | 名 | [optional] 
**LastNamePhonetics** | Pointer to **string** | 姓（カナ） | [optional] 
**FirstNamePhonetics** | Pointer to **string** | 名（カナ） | [optional] 
**HiredDate** | Pointer to **string** | 入社年月日 | [optional] 
**BirthDate** | Pointer to **string** | 生年月日 | [optional] 
**ResignationDate** | Pointer to **string** | 退職年月日 | [optional] 
**EmailAddresses** | Pointer to **[]string** | メールアドレス | [optional] 
**AllDayRegardingWorkInMinute** | Pointer to **int32** | 日の契約労働時間 | [optional] 

## Methods

### NewUpdateEmployeeRequest

`func NewUpdateEmployeeRequest() *UpdateEmployeeRequest`

NewUpdateEmployeeRequest instantiates a new UpdateEmployeeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEmployeeRequestWithDefaults

`func NewUpdateEmployeeRequestWithDefaults() *UpdateEmployeeRequest`

NewUpdateEmployeeRequestWithDefaults instantiates a new UpdateEmployeeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDivisionCode

`func (o *UpdateEmployeeRequest) GetDivisionCode() string`

GetDivisionCode returns the DivisionCode field if non-nil, zero value otherwise.

### GetDivisionCodeOk

`func (o *UpdateEmployeeRequest) GetDivisionCodeOk() (*string, bool)`

GetDivisionCodeOk returns a tuple with the DivisionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionCode

`func (o *UpdateEmployeeRequest) SetDivisionCode(v string)`

SetDivisionCode sets DivisionCode field to given value.

### HasDivisionCode

`func (o *UpdateEmployeeRequest) HasDivisionCode() bool`

HasDivisionCode returns a boolean if a field has been set.

### GetGender

`func (o *UpdateEmployeeRequest) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UpdateEmployeeRequest) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UpdateEmployeeRequest) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *UpdateEmployeeRequest) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetTypeCode

`func (o *UpdateEmployeeRequest) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *UpdateEmployeeRequest) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *UpdateEmployeeRequest) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *UpdateEmployeeRequest) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetCode

`func (o *UpdateEmployeeRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateEmployeeRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateEmployeeRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateEmployeeRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLastName

`func (o *UpdateEmployeeRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateEmployeeRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateEmployeeRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UpdateEmployeeRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetFirstName

`func (o *UpdateEmployeeRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateEmployeeRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateEmployeeRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UpdateEmployeeRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastNamePhonetics

`func (o *UpdateEmployeeRequest) GetLastNamePhonetics() string`

GetLastNamePhonetics returns the LastNamePhonetics field if non-nil, zero value otherwise.

### GetLastNamePhoneticsOk

`func (o *UpdateEmployeeRequest) GetLastNamePhoneticsOk() (*string, bool)`

GetLastNamePhoneticsOk returns a tuple with the LastNamePhonetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNamePhonetics

`func (o *UpdateEmployeeRequest) SetLastNamePhonetics(v string)`

SetLastNamePhonetics sets LastNamePhonetics field to given value.

### HasLastNamePhonetics

`func (o *UpdateEmployeeRequest) HasLastNamePhonetics() bool`

HasLastNamePhonetics returns a boolean if a field has been set.

### GetFirstNamePhonetics

`func (o *UpdateEmployeeRequest) GetFirstNamePhonetics() string`

GetFirstNamePhonetics returns the FirstNamePhonetics field if non-nil, zero value otherwise.

### GetFirstNamePhoneticsOk

`func (o *UpdateEmployeeRequest) GetFirstNamePhoneticsOk() (*string, bool)`

GetFirstNamePhoneticsOk returns a tuple with the FirstNamePhonetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNamePhonetics

`func (o *UpdateEmployeeRequest) SetFirstNamePhonetics(v string)`

SetFirstNamePhonetics sets FirstNamePhonetics field to given value.

### HasFirstNamePhonetics

`func (o *UpdateEmployeeRequest) HasFirstNamePhonetics() bool`

HasFirstNamePhonetics returns a boolean if a field has been set.

### GetHiredDate

`func (o *UpdateEmployeeRequest) GetHiredDate() string`

GetHiredDate returns the HiredDate field if non-nil, zero value otherwise.

### GetHiredDateOk

`func (o *UpdateEmployeeRequest) GetHiredDateOk() (*string, bool)`

GetHiredDateOk returns a tuple with the HiredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiredDate

`func (o *UpdateEmployeeRequest) SetHiredDate(v string)`

SetHiredDate sets HiredDate field to given value.

### HasHiredDate

`func (o *UpdateEmployeeRequest) HasHiredDate() bool`

HasHiredDate returns a boolean if a field has been set.

### GetBirthDate

`func (o *UpdateEmployeeRequest) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *UpdateEmployeeRequest) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *UpdateEmployeeRequest) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *UpdateEmployeeRequest) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetResignationDate

`func (o *UpdateEmployeeRequest) GetResignationDate() string`

GetResignationDate returns the ResignationDate field if non-nil, zero value otherwise.

### GetResignationDateOk

`func (o *UpdateEmployeeRequest) GetResignationDateOk() (*string, bool)`

GetResignationDateOk returns a tuple with the ResignationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResignationDate

`func (o *UpdateEmployeeRequest) SetResignationDate(v string)`

SetResignationDate sets ResignationDate field to given value.

### HasResignationDate

`func (o *UpdateEmployeeRequest) HasResignationDate() bool`

HasResignationDate returns a boolean if a field has been set.

### GetEmailAddresses

`func (o *UpdateEmployeeRequest) GetEmailAddresses() []string`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *UpdateEmployeeRequest) GetEmailAddressesOk() (*[]string, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *UpdateEmployeeRequest) SetEmailAddresses(v []string)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *UpdateEmployeeRequest) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### GetAllDayRegardingWorkInMinute

`func (o *UpdateEmployeeRequest) GetAllDayRegardingWorkInMinute() int32`

GetAllDayRegardingWorkInMinute returns the AllDayRegardingWorkInMinute field if non-nil, zero value otherwise.

### GetAllDayRegardingWorkInMinuteOk

`func (o *UpdateEmployeeRequest) GetAllDayRegardingWorkInMinuteOk() (*int32, bool)`

GetAllDayRegardingWorkInMinuteOk returns a tuple with the AllDayRegardingWorkInMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDayRegardingWorkInMinute

`func (o *UpdateEmployeeRequest) SetAllDayRegardingWorkInMinute(v int32)`

SetAllDayRegardingWorkInMinute sets AllDayRegardingWorkInMinute field to given value.

### HasAllDayRegardingWorkInMinute

`func (o *UpdateEmployeeRequest) HasAllDayRegardingWorkInMinute() bool`

HasAllDayRegardingWorkInMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


