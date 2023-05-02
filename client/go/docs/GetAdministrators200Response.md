# GetAdministrators200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | 管理者コード | 
**Key** | **string** | 管理者識別キー（管理者コードが変更されても不変） | 
**Name** | **string** | 管理者名 | 
**EmailAddresses** | Pointer to **string** | メールアドレス | [optional] 
**AssociatedEmployees** | Pointer to [**GetAdministrators200ResponseAssociatedEmployees**](GetAdministrators200ResponseAssociatedEmployees.md) |  | [optional] 

## Methods

### NewGetAdministrators200Response

`func NewGetAdministrators200Response(code string, key string, name string, ) *GetAdministrators200Response`

NewGetAdministrators200Response instantiates a new GetAdministrators200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdministrators200ResponseWithDefaults

`func NewGetAdministrators200ResponseWithDefaults() *GetAdministrators200Response`

NewGetAdministrators200ResponseWithDefaults instantiates a new GetAdministrators200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetAdministrators200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetAdministrators200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetAdministrators200Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetKey

`func (o *GetAdministrators200Response) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetAdministrators200Response) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetAdministrators200Response) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *GetAdministrators200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAdministrators200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAdministrators200Response) SetName(v string)`

SetName sets Name field to given value.


### GetEmailAddresses

`func (o *GetAdministrators200Response) GetEmailAddresses() string`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *GetAdministrators200Response) GetEmailAddressesOk() (*string, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *GetAdministrators200Response) SetEmailAddresses(v string)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *GetAdministrators200Response) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### GetAssociatedEmployees

`func (o *GetAdministrators200Response) GetAssociatedEmployees() GetAdministrators200ResponseAssociatedEmployees`

GetAssociatedEmployees returns the AssociatedEmployees field if non-nil, zero value otherwise.

### GetAssociatedEmployeesOk

`func (o *GetAdministrators200Response) GetAssociatedEmployeesOk() (*GetAdministrators200ResponseAssociatedEmployees, bool)`

GetAssociatedEmployeesOk returns a tuple with the AssociatedEmployees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedEmployees

`func (o *GetAdministrators200Response) SetAssociatedEmployees(v GetAdministrators200ResponseAssociatedEmployees)`

SetAssociatedEmployees sets AssociatedEmployees field to given value.

### HasAssociatedEmployees

`func (o *GetAdministrators200Response) HasAssociatedEmployees() bool`

HasAssociatedEmployees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


