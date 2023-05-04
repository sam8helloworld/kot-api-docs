# GetSchedules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Year** | **int32** |  | 
**Month** | **int32** |  | 
**ScheduleRequests** | [**[]GetSchedules200ResponseScheduleRequestsInner**](GetSchedules200ResponseScheduleRequestsInner.md) |  | 

## Methods

### NewGetSchedules200Response

`func NewGetSchedules200Response(year int32, month int32, scheduleRequests []GetSchedules200ResponseScheduleRequestsInner, ) *GetSchedules200Response`

NewGetSchedules200Response instantiates a new GetSchedules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSchedules200ResponseWithDefaults

`func NewGetSchedules200ResponseWithDefaults() *GetSchedules200Response`

NewGetSchedules200ResponseWithDefaults instantiates a new GetSchedules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYear

`func (o *GetSchedules200Response) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *GetSchedules200Response) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *GetSchedules200Response) SetYear(v int32)`

SetYear sets Year field to given value.


### GetMonth

`func (o *GetSchedules200Response) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *GetSchedules200Response) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *GetSchedules200Response) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetScheduleRequests

`func (o *GetSchedules200Response) GetScheduleRequests() []GetSchedules200ResponseScheduleRequestsInner`

GetScheduleRequests returns the ScheduleRequests field if non-nil, zero value otherwise.

### GetScheduleRequestsOk

`func (o *GetSchedules200Response) GetScheduleRequestsOk() (*[]GetSchedules200ResponseScheduleRequestsInner, bool)`

GetScheduleRequestsOk returns a tuple with the ScheduleRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleRequests

`func (o *GetSchedules200Response) SetScheduleRequests(v []GetSchedules200ResponseScheduleRequestsInner)`

SetScheduleRequests sets ScheduleRequests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


