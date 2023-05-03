# GetDailyWorkingTimerecord200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**DailyWorkings** | [**[]DailyWorkingTimerecord**](DailyWorkingTimerecord.md) |  | 

## Methods

### NewGetDailyWorkingTimerecord200Response

`func NewGetDailyWorkingTimerecord200Response(date string, dailyWorkings []DailyWorkingTimerecord, ) *GetDailyWorkingTimerecord200Response`

NewGetDailyWorkingTimerecord200Response instantiates a new GetDailyWorkingTimerecord200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDailyWorkingTimerecord200ResponseWithDefaults

`func NewGetDailyWorkingTimerecord200ResponseWithDefaults() *GetDailyWorkingTimerecord200Response`

NewGetDailyWorkingTimerecord200ResponseWithDefaults instantiates a new GetDailyWorkingTimerecord200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetDailyWorkingTimerecord200Response) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetDailyWorkingTimerecord200Response) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetDailyWorkingTimerecord200Response) SetDate(v string)`

SetDate sets Date field to given value.


### GetDailyWorkings

`func (o *GetDailyWorkingTimerecord200Response) GetDailyWorkings() []DailyWorkingTimerecord`

GetDailyWorkings returns the DailyWorkings field if non-nil, zero value otherwise.

### GetDailyWorkingsOk

`func (o *GetDailyWorkingTimerecord200Response) GetDailyWorkingsOk() (*[]DailyWorkingTimerecord, bool)`

GetDailyWorkingsOk returns a tuple with the DailyWorkings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyWorkings

`func (o *GetDailyWorkingTimerecord200Response) SetDailyWorkings(v []DailyWorkingTimerecord)`

SetDailyWorkings sets DailyWorkings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


