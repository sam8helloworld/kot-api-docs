# GetCompany200ResponseSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeDisplayFormat** | **string** | 表示形式（decimal： 10進法　sexagesimal： 60進法） | 
**DecimalTreatType** | **string** | 10進表示の小数第3位の扱い（roundDown： 切下げ　roundUp： 切上げ　round： 四捨五入） | 

## Methods

### NewGetCompany200ResponseSettings

`func NewGetCompany200ResponseSettings(timeDisplayFormat string, decimalTreatType string, ) *GetCompany200ResponseSettings`

NewGetCompany200ResponseSettings instantiates a new GetCompany200ResponseSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCompany200ResponseSettingsWithDefaults

`func NewGetCompany200ResponseSettingsWithDefaults() *GetCompany200ResponseSettings`

NewGetCompany200ResponseSettingsWithDefaults instantiates a new GetCompany200ResponseSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeDisplayFormat

`func (o *GetCompany200ResponseSettings) GetTimeDisplayFormat() string`

GetTimeDisplayFormat returns the TimeDisplayFormat field if non-nil, zero value otherwise.

### GetTimeDisplayFormatOk

`func (o *GetCompany200ResponseSettings) GetTimeDisplayFormatOk() (*string, bool)`

GetTimeDisplayFormatOk returns a tuple with the TimeDisplayFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeDisplayFormat

`func (o *GetCompany200ResponseSettings) SetTimeDisplayFormat(v string)`

SetTimeDisplayFormat sets TimeDisplayFormat field to given value.


### GetDecimalTreatType

`func (o *GetCompany200ResponseSettings) GetDecimalTreatType() string`

GetDecimalTreatType returns the DecimalTreatType field if non-nil, zero value otherwise.

### GetDecimalTreatTypeOk

`func (o *GetCompany200ResponseSettings) GetDecimalTreatTypeOk() (*string, bool)`

GetDecimalTreatTypeOk returns a tuple with the DecimalTreatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalTreatType

`func (o *GetCompany200ResponseSettings) SetDecimalTreatType(v string)`

SetDecimalTreatType sets DecimalTreatType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


