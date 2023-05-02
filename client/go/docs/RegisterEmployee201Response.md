# RegisterEmployee201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DivisionCode** | **string** | 所属コード | 
**Gender** | **string** | 性別（male： 男性　female： 女性） | 
**TypeCode** | **string** | 雇用区分コード | 
**Code** | **string** | 従業員コード | 
**LastName** | **string** | 姓 | 
**FirstName** | **string** | 名 | 
**Key** | **string** | 従業員識別キー（従業員コードが変更されても不変） | 

## Methods

### NewRegisterEmployee201Response

`func NewRegisterEmployee201Response(divisionCode string, gender string, typeCode string, code string, lastName string, firstName string, key string, ) *RegisterEmployee201Response`

NewRegisterEmployee201Response instantiates a new RegisterEmployee201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterEmployee201ResponseWithDefaults

`func NewRegisterEmployee201ResponseWithDefaults() *RegisterEmployee201Response`

NewRegisterEmployee201ResponseWithDefaults instantiates a new RegisterEmployee201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDivisionCode

`func (o *RegisterEmployee201Response) GetDivisionCode() string`

GetDivisionCode returns the DivisionCode field if non-nil, zero value otherwise.

### GetDivisionCodeOk

`func (o *RegisterEmployee201Response) GetDivisionCodeOk() (*string, bool)`

GetDivisionCodeOk returns a tuple with the DivisionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionCode

`func (o *RegisterEmployee201Response) SetDivisionCode(v string)`

SetDivisionCode sets DivisionCode field to given value.


### GetGender

`func (o *RegisterEmployee201Response) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *RegisterEmployee201Response) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *RegisterEmployee201Response) SetGender(v string)`

SetGender sets Gender field to given value.


### GetTypeCode

`func (o *RegisterEmployee201Response) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *RegisterEmployee201Response) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *RegisterEmployee201Response) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.


### GetCode

`func (o *RegisterEmployee201Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RegisterEmployee201Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RegisterEmployee201Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetLastName

`func (o *RegisterEmployee201Response) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *RegisterEmployee201Response) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *RegisterEmployee201Response) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetFirstName

`func (o *RegisterEmployee201Response) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *RegisterEmployee201Response) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *RegisterEmployee201Response) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetKey

`func (o *RegisterEmployee201Response) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RegisterEmployee201Response) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RegisterEmployee201Response) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


