# GetOvertimes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Year** | **int32** |  | 
**Month** | **int32** |  | 
**OvertimeRequests** | [**[]GetOvertimes200ResponseOvertimeRequestsInner**](GetOvertimes200ResponseOvertimeRequestsInner.md) |  | 

## Methods

### NewGetOvertimes200Response

`func NewGetOvertimes200Response(year int32, month int32, overtimeRequests []GetOvertimes200ResponseOvertimeRequestsInner, ) *GetOvertimes200Response`

NewGetOvertimes200Response instantiates a new GetOvertimes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOvertimes200ResponseWithDefaults

`func NewGetOvertimes200ResponseWithDefaults() *GetOvertimes200Response`

NewGetOvertimes200ResponseWithDefaults instantiates a new GetOvertimes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYear

`func (o *GetOvertimes200Response) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *GetOvertimes200Response) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *GetOvertimes200Response) SetYear(v int32)`

SetYear sets Year field to given value.


### GetMonth

`func (o *GetOvertimes200Response) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *GetOvertimes200Response) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *GetOvertimes200Response) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetOvertimeRequests

`func (o *GetOvertimes200Response) GetOvertimeRequests() []GetOvertimes200ResponseOvertimeRequestsInner`

GetOvertimeRequests returns the OvertimeRequests field if non-nil, zero value otherwise.

### GetOvertimeRequestsOk

`func (o *GetOvertimes200Response) GetOvertimeRequestsOk() (*[]GetOvertimes200ResponseOvertimeRequestsInner, bool)`

GetOvertimeRequestsOk returns a tuple with the OvertimeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvertimeRequests

`func (o *GetOvertimes200Response) SetOvertimeRequests(v []GetOvertimes200ResponseOvertimeRequestsInner)`

SetOvertimeRequests sets OvertimeRequests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


