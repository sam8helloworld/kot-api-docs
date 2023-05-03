# GetEmployeeGroups200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | 従業員グループコード | 
**Name** | **string** | 従業員グループ名 | 
**Category** | Pointer to [**GetEmployeeGroups200ResponseInnerCategory**](GetEmployeeGroups200ResponseInnerCategory.md) |  | [optional] 

## Methods

### NewGetEmployeeGroups200ResponseInner

`func NewGetEmployeeGroups200ResponseInner(code string, name string, ) *GetEmployeeGroups200ResponseInner`

NewGetEmployeeGroups200ResponseInner instantiates a new GetEmployeeGroups200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEmployeeGroups200ResponseInnerWithDefaults

`func NewGetEmployeeGroups200ResponseInnerWithDefaults() *GetEmployeeGroups200ResponseInner`

NewGetEmployeeGroups200ResponseInnerWithDefaults instantiates a new GetEmployeeGroups200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetEmployeeGroups200ResponseInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetEmployeeGroups200ResponseInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetEmployeeGroups200ResponseInner) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *GetEmployeeGroups200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetEmployeeGroups200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetEmployeeGroups200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *GetEmployeeGroups200ResponseInner) GetCategory() GetEmployeeGroups200ResponseInnerCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetEmployeeGroups200ResponseInner) GetCategoryOk() (*GetEmployeeGroups200ResponseInnerCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetEmployeeGroups200ResponseInner) SetCategory(v GetEmployeeGroups200ResponseInnerCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetEmployeeGroups200ResponseInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


