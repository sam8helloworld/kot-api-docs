/*
KING OF TIME WebAPI

KING OF TIME WebAPI specification https://developer.kingtime.jp/

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DailyWorkingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DailyWorkingResponse{}

// DailyWorkingResponse struct for DailyWorkingResponse
type DailyWorkingResponse struct {
	// 日時
	Date string `json:"date"`
	// 従業員識別キー（従業員コードが変更されても不変）
	EmployeeKey string `json:"employeeKey"`
	// 出勤先所属コード
	WorkPlaceDivisionCode string `json:"workPlaceDivisionCode"`
	// 出勤先所属名
	WorkPlaceDivisionName *string `json:"workPlaceDivisionName,omitempty"`
	// 締め状況
	IsClosing bool `json:"isClosing"`
	// ヘルプ勤務状況
	IsHelp bool `json:"isHelp"`
	// エラー勤務状況
	IsError bool `json:"isError"`
	// 勤務日種別名
	WorkdayTypeName string `json:"workdayTypeName"`
	// 所定時間（分）
	Assigned int32 `json:"assigned"`
	// 所定外時間（分）
	Unassigned int32 `json:"unassigned"`
	// 残業時間（分）
	Overtime int32 `json:"overtime"`
	// 深夜時間（分）
	LateNight int32 `json:"lateNight"`
	// 深夜所定外時間（分）
	LateNightUnassigned int32 `json:"lateNightUnassigned"`
	// 深夜残業時間（分）
	LateNightOvertime int32 `json:"lateNightOvertime"`
	// 休憩時間（分）
	BreakTime int32 `json:"breakTime"`
	// 遅刻時間（分）
	Late int32 `json:"late"`
	// 早退時間（分）
	EarlyLeave int32 `json:"earlyLeave"`
	// 労働合計時間（分）
	TotalWork int32 `json:"totalWork"`
	HolidaysObtained DailyWorkingResponseHolidaysObtained `json:"holidaysObtained"`
	// 自動休憩無効（null： 休憩を無効化しない 1：　雇用区分休憩無効　2： スケジュール休憩無効　3： 全ての自動休憩無効）
	AutoBreakOff NullableInt32 `json:"autoBreakOff"`
	// 休暇みなし時間（分）
	DiscretionaryVacation int32 `json:"discretionaryVacation"`
	CustomDailyWorkings DailyWorkingResponseCustomDailyWorkings `json:"customDailyWorkings"`
	CurrentDateEmployee *DailyWorkingResponseCurrentDateEmployee `json:"currentDateEmployee,omitempty"`
}

// NewDailyWorkingResponse instantiates a new DailyWorkingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailyWorkingResponse(date string, employeeKey string, workPlaceDivisionCode string, isClosing bool, isHelp bool, isError bool, workdayTypeName string, assigned int32, unassigned int32, overtime int32, lateNight int32, lateNightUnassigned int32, lateNightOvertime int32, breakTime int32, late int32, earlyLeave int32, totalWork int32, holidaysObtained DailyWorkingResponseHolidaysObtained, autoBreakOff NullableInt32, discretionaryVacation int32, customDailyWorkings DailyWorkingResponseCustomDailyWorkings) *DailyWorkingResponse {
	this := DailyWorkingResponse{}
	this.Date = date
	this.EmployeeKey = employeeKey
	this.WorkPlaceDivisionCode = workPlaceDivisionCode
	this.IsClosing = isClosing
	this.IsHelp = isHelp
	this.IsError = isError
	this.WorkdayTypeName = workdayTypeName
	this.Assigned = assigned
	this.Unassigned = unassigned
	this.Overtime = overtime
	this.LateNight = lateNight
	this.LateNightUnassigned = lateNightUnassigned
	this.LateNightOvertime = lateNightOvertime
	this.BreakTime = breakTime
	this.Late = late
	this.EarlyLeave = earlyLeave
	this.TotalWork = totalWork
	this.HolidaysObtained = holidaysObtained
	this.AutoBreakOff = autoBreakOff
	this.DiscretionaryVacation = discretionaryVacation
	this.CustomDailyWorkings = customDailyWorkings
	return &this
}

// NewDailyWorkingResponseWithDefaults instantiates a new DailyWorkingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyWorkingResponseWithDefaults() *DailyWorkingResponse {
	this := DailyWorkingResponse{}
	return &this
}

// GetDate returns the Date field value
func (o *DailyWorkingResponse) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *DailyWorkingResponse) SetDate(v string) {
	o.Date = v
}

// GetEmployeeKey returns the EmployeeKey field value
func (o *DailyWorkingResponse) GetEmployeeKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmployeeKey
}

// GetEmployeeKeyOk returns a tuple with the EmployeeKey field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetEmployeeKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmployeeKey, true
}

// SetEmployeeKey sets field value
func (o *DailyWorkingResponse) SetEmployeeKey(v string) {
	o.EmployeeKey = v
}

// GetWorkPlaceDivisionCode returns the WorkPlaceDivisionCode field value
func (o *DailyWorkingResponse) GetWorkPlaceDivisionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkPlaceDivisionCode
}

// GetWorkPlaceDivisionCodeOk returns a tuple with the WorkPlaceDivisionCode field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetWorkPlaceDivisionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkPlaceDivisionCode, true
}

// SetWorkPlaceDivisionCode sets field value
func (o *DailyWorkingResponse) SetWorkPlaceDivisionCode(v string) {
	o.WorkPlaceDivisionCode = v
}

// GetWorkPlaceDivisionName returns the WorkPlaceDivisionName field value if set, zero value otherwise.
func (o *DailyWorkingResponse) GetWorkPlaceDivisionName() string {
	if o == nil || IsNil(o.WorkPlaceDivisionName) {
		var ret string
		return ret
	}
	return *o.WorkPlaceDivisionName
}

// GetWorkPlaceDivisionNameOk returns a tuple with the WorkPlaceDivisionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetWorkPlaceDivisionNameOk() (*string, bool) {
	if o == nil || IsNil(o.WorkPlaceDivisionName) {
		return nil, false
	}
	return o.WorkPlaceDivisionName, true
}

// HasWorkPlaceDivisionName returns a boolean if a field has been set.
func (o *DailyWorkingResponse) HasWorkPlaceDivisionName() bool {
	if o != nil && !IsNil(o.WorkPlaceDivisionName) {
		return true
	}

	return false
}

// SetWorkPlaceDivisionName gets a reference to the given string and assigns it to the WorkPlaceDivisionName field.
func (o *DailyWorkingResponse) SetWorkPlaceDivisionName(v string) {
	o.WorkPlaceDivisionName = &v
}

// GetIsClosing returns the IsClosing field value
func (o *DailyWorkingResponse) GetIsClosing() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsClosing
}

// GetIsClosingOk returns a tuple with the IsClosing field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetIsClosingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsClosing, true
}

// SetIsClosing sets field value
func (o *DailyWorkingResponse) SetIsClosing(v bool) {
	o.IsClosing = v
}

// GetIsHelp returns the IsHelp field value
func (o *DailyWorkingResponse) GetIsHelp() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsHelp
}

// GetIsHelpOk returns a tuple with the IsHelp field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetIsHelpOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsHelp, true
}

// SetIsHelp sets field value
func (o *DailyWorkingResponse) SetIsHelp(v bool) {
	o.IsHelp = v
}

// GetIsError returns the IsError field value
func (o *DailyWorkingResponse) GetIsError() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsError
}

// GetIsErrorOk returns a tuple with the IsError field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetIsErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsError, true
}

// SetIsError sets field value
func (o *DailyWorkingResponse) SetIsError(v bool) {
	o.IsError = v
}

// GetWorkdayTypeName returns the WorkdayTypeName field value
func (o *DailyWorkingResponse) GetWorkdayTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkdayTypeName
}

// GetWorkdayTypeNameOk returns a tuple with the WorkdayTypeName field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetWorkdayTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkdayTypeName, true
}

// SetWorkdayTypeName sets field value
func (o *DailyWorkingResponse) SetWorkdayTypeName(v string) {
	o.WorkdayTypeName = v
}

// GetAssigned returns the Assigned field value
func (o *DailyWorkingResponse) GetAssigned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetAssignedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assigned, true
}

// SetAssigned sets field value
func (o *DailyWorkingResponse) SetAssigned(v int32) {
	o.Assigned = v
}

// GetUnassigned returns the Unassigned field value
func (o *DailyWorkingResponse) GetUnassigned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Unassigned
}

// GetUnassignedOk returns a tuple with the Unassigned field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetUnassignedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unassigned, true
}

// SetUnassigned sets field value
func (o *DailyWorkingResponse) SetUnassigned(v int32) {
	o.Unassigned = v
}

// GetOvertime returns the Overtime field value
func (o *DailyWorkingResponse) GetOvertime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Overtime
}

// GetOvertimeOk returns a tuple with the Overtime field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetOvertimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Overtime, true
}

// SetOvertime sets field value
func (o *DailyWorkingResponse) SetOvertime(v int32) {
	o.Overtime = v
}

// GetLateNight returns the LateNight field value
func (o *DailyWorkingResponse) GetLateNight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LateNight
}

// GetLateNightOk returns a tuple with the LateNight field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetLateNightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LateNight, true
}

// SetLateNight sets field value
func (o *DailyWorkingResponse) SetLateNight(v int32) {
	o.LateNight = v
}

// GetLateNightUnassigned returns the LateNightUnassigned field value
func (o *DailyWorkingResponse) GetLateNightUnassigned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LateNightUnassigned
}

// GetLateNightUnassignedOk returns a tuple with the LateNightUnassigned field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetLateNightUnassignedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LateNightUnassigned, true
}

// SetLateNightUnassigned sets field value
func (o *DailyWorkingResponse) SetLateNightUnassigned(v int32) {
	o.LateNightUnassigned = v
}

// GetLateNightOvertime returns the LateNightOvertime field value
func (o *DailyWorkingResponse) GetLateNightOvertime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LateNightOvertime
}

// GetLateNightOvertimeOk returns a tuple with the LateNightOvertime field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetLateNightOvertimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LateNightOvertime, true
}

// SetLateNightOvertime sets field value
func (o *DailyWorkingResponse) SetLateNightOvertime(v int32) {
	o.LateNightOvertime = v
}

// GetBreakTime returns the BreakTime field value
func (o *DailyWorkingResponse) GetBreakTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BreakTime
}

// GetBreakTimeOk returns a tuple with the BreakTime field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetBreakTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BreakTime, true
}

// SetBreakTime sets field value
func (o *DailyWorkingResponse) SetBreakTime(v int32) {
	o.BreakTime = v
}

// GetLate returns the Late field value
func (o *DailyWorkingResponse) GetLate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Late
}

// GetLateOk returns a tuple with the Late field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetLateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Late, true
}

// SetLate sets field value
func (o *DailyWorkingResponse) SetLate(v int32) {
	o.Late = v
}

// GetEarlyLeave returns the EarlyLeave field value
func (o *DailyWorkingResponse) GetEarlyLeave() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EarlyLeave
}

// GetEarlyLeaveOk returns a tuple with the EarlyLeave field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetEarlyLeaveOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EarlyLeave, true
}

// SetEarlyLeave sets field value
func (o *DailyWorkingResponse) SetEarlyLeave(v int32) {
	o.EarlyLeave = v
}

// GetTotalWork returns the TotalWork field value
func (o *DailyWorkingResponse) GetTotalWork() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalWork
}

// GetTotalWorkOk returns a tuple with the TotalWork field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetTotalWorkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalWork, true
}

// SetTotalWork sets field value
func (o *DailyWorkingResponse) SetTotalWork(v int32) {
	o.TotalWork = v
}

// GetHolidaysObtained returns the HolidaysObtained field value
func (o *DailyWorkingResponse) GetHolidaysObtained() DailyWorkingResponseHolidaysObtained {
	if o == nil {
		var ret DailyWorkingResponseHolidaysObtained
		return ret
	}

	return o.HolidaysObtained
}

// GetHolidaysObtainedOk returns a tuple with the HolidaysObtained field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetHolidaysObtainedOk() (*DailyWorkingResponseHolidaysObtained, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HolidaysObtained, true
}

// SetHolidaysObtained sets field value
func (o *DailyWorkingResponse) SetHolidaysObtained(v DailyWorkingResponseHolidaysObtained) {
	o.HolidaysObtained = v
}

// GetAutoBreakOff returns the AutoBreakOff field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *DailyWorkingResponse) GetAutoBreakOff() int32 {
	if o == nil || o.AutoBreakOff.Get() == nil {
		var ret int32
		return ret
	}

	return *o.AutoBreakOff.Get()
}

// GetAutoBreakOffOk returns a tuple with the AutoBreakOff field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DailyWorkingResponse) GetAutoBreakOffOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoBreakOff.Get(), o.AutoBreakOff.IsSet()
}

// SetAutoBreakOff sets field value
func (o *DailyWorkingResponse) SetAutoBreakOff(v int32) {
	o.AutoBreakOff.Set(&v)
}

// GetDiscretionaryVacation returns the DiscretionaryVacation field value
func (o *DailyWorkingResponse) GetDiscretionaryVacation() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DiscretionaryVacation
}

// GetDiscretionaryVacationOk returns a tuple with the DiscretionaryVacation field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetDiscretionaryVacationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscretionaryVacation, true
}

// SetDiscretionaryVacation sets field value
func (o *DailyWorkingResponse) SetDiscretionaryVacation(v int32) {
	o.DiscretionaryVacation = v
}

// GetCustomDailyWorkings returns the CustomDailyWorkings field value
func (o *DailyWorkingResponse) GetCustomDailyWorkings() DailyWorkingResponseCustomDailyWorkings {
	if o == nil {
		var ret DailyWorkingResponseCustomDailyWorkings
		return ret
	}

	return o.CustomDailyWorkings
}

// GetCustomDailyWorkingsOk returns a tuple with the CustomDailyWorkings field value
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetCustomDailyWorkingsOk() (*DailyWorkingResponseCustomDailyWorkings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomDailyWorkings, true
}

// SetCustomDailyWorkings sets field value
func (o *DailyWorkingResponse) SetCustomDailyWorkings(v DailyWorkingResponseCustomDailyWorkings) {
	o.CustomDailyWorkings = v
}

// GetCurrentDateEmployee returns the CurrentDateEmployee field value if set, zero value otherwise.
func (o *DailyWorkingResponse) GetCurrentDateEmployee() DailyWorkingResponseCurrentDateEmployee {
	if o == nil || IsNil(o.CurrentDateEmployee) {
		var ret DailyWorkingResponseCurrentDateEmployee
		return ret
	}
	return *o.CurrentDateEmployee
}

// GetCurrentDateEmployeeOk returns a tuple with the CurrentDateEmployee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyWorkingResponse) GetCurrentDateEmployeeOk() (*DailyWorkingResponseCurrentDateEmployee, bool) {
	if o == nil || IsNil(o.CurrentDateEmployee) {
		return nil, false
	}
	return o.CurrentDateEmployee, true
}

// HasCurrentDateEmployee returns a boolean if a field has been set.
func (o *DailyWorkingResponse) HasCurrentDateEmployee() bool {
	if o != nil && !IsNil(o.CurrentDateEmployee) {
		return true
	}

	return false
}

// SetCurrentDateEmployee gets a reference to the given DailyWorkingResponseCurrentDateEmployee and assigns it to the CurrentDateEmployee field.
func (o *DailyWorkingResponse) SetCurrentDateEmployee(v DailyWorkingResponseCurrentDateEmployee) {
	o.CurrentDateEmployee = &v
}

func (o DailyWorkingResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DailyWorkingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["employeeKey"] = o.EmployeeKey
	toSerialize["workPlaceDivisionCode"] = o.WorkPlaceDivisionCode
	if !IsNil(o.WorkPlaceDivisionName) {
		toSerialize["workPlaceDivisionName"] = o.WorkPlaceDivisionName
	}
	toSerialize["isClosing"] = o.IsClosing
	toSerialize["isHelp"] = o.IsHelp
	toSerialize["isError"] = o.IsError
	toSerialize["workdayTypeName"] = o.WorkdayTypeName
	toSerialize["assigned"] = o.Assigned
	toSerialize["unassigned"] = o.Unassigned
	toSerialize["overtime"] = o.Overtime
	toSerialize["lateNight"] = o.LateNight
	toSerialize["lateNightUnassigned"] = o.LateNightUnassigned
	toSerialize["lateNightOvertime"] = o.LateNightOvertime
	toSerialize["breakTime"] = o.BreakTime
	toSerialize["late"] = o.Late
	toSerialize["earlyLeave"] = o.EarlyLeave
	toSerialize["totalWork"] = o.TotalWork
	toSerialize["holidaysObtained"] = o.HolidaysObtained
	toSerialize["autoBreakOff"] = o.AutoBreakOff.Get()
	toSerialize["discretionaryVacation"] = o.DiscretionaryVacation
	toSerialize["customDailyWorkings"] = o.CustomDailyWorkings
	if !IsNil(o.CurrentDateEmployee) {
		toSerialize["currentDateEmployee"] = o.CurrentDateEmployee
	}
	return toSerialize, nil
}

type NullableDailyWorkingResponse struct {
	value *DailyWorkingResponse
	isSet bool
}

func (v NullableDailyWorkingResponse) Get() *DailyWorkingResponse {
	return v.value
}

func (v *NullableDailyWorkingResponse) Set(val *DailyWorkingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyWorkingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyWorkingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyWorkingResponse(val *DailyWorkingResponse) *NullableDailyWorkingResponse {
	return &NullableDailyWorkingResponse{value: val, isSet: true}
}

func (v NullableDailyWorkingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyWorkingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


