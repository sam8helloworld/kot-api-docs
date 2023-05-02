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

// checks if the GetCompany200ResponseSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCompany200ResponseSettings{}

// GetCompany200ResponseSettings struct for GetCompany200ResponseSettings
type GetCompany200ResponseSettings struct {
	TimeDisplayFormat string `json:"timeDisplayFormat"`
	DecimalTreatType string `json:"decimalTreatType"`
}

// NewGetCompany200ResponseSettings instantiates a new GetCompany200ResponseSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCompany200ResponseSettings(timeDisplayFormat string, decimalTreatType string) *GetCompany200ResponseSettings {
	this := GetCompany200ResponseSettings{}
	this.TimeDisplayFormat = timeDisplayFormat
	this.DecimalTreatType = decimalTreatType
	return &this
}

// NewGetCompany200ResponseSettingsWithDefaults instantiates a new GetCompany200ResponseSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCompany200ResponseSettingsWithDefaults() *GetCompany200ResponseSettings {
	this := GetCompany200ResponseSettings{}
	return &this
}

// GetTimeDisplayFormat returns the TimeDisplayFormat field value
func (o *GetCompany200ResponseSettings) GetTimeDisplayFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeDisplayFormat
}

// GetTimeDisplayFormatOk returns a tuple with the TimeDisplayFormat field value
// and a boolean to check if the value has been set.
func (o *GetCompany200ResponseSettings) GetTimeDisplayFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeDisplayFormat, true
}

// SetTimeDisplayFormat sets field value
func (o *GetCompany200ResponseSettings) SetTimeDisplayFormat(v string) {
	o.TimeDisplayFormat = v
}

// GetDecimalTreatType returns the DecimalTreatType field value
func (o *GetCompany200ResponseSettings) GetDecimalTreatType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DecimalTreatType
}

// GetDecimalTreatTypeOk returns a tuple with the DecimalTreatType field value
// and a boolean to check if the value has been set.
func (o *GetCompany200ResponseSettings) GetDecimalTreatTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DecimalTreatType, true
}

// SetDecimalTreatType sets field value
func (o *GetCompany200ResponseSettings) SetDecimalTreatType(v string) {
	o.DecimalTreatType = v
}

func (o GetCompany200ResponseSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCompany200ResponseSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timeDisplayFormat"] = o.TimeDisplayFormat
	toSerialize["decimalTreatType"] = o.DecimalTreatType
	return toSerialize, nil
}

type NullableGetCompany200ResponseSettings struct {
	value *GetCompany200ResponseSettings
	isSet bool
}

func (v NullableGetCompany200ResponseSettings) Get() *GetCompany200ResponseSettings {
	return v.value
}

func (v *NullableGetCompany200ResponseSettings) Set(val *GetCompany200ResponseSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCompany200ResponseSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCompany200ResponseSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCompany200ResponseSettings(val *GetCompany200ResponseSettings) *NullableGetCompany200ResponseSettings {
	return &NullableGetCompany200ResponseSettings{value: val, isSet: true}
}

func (v NullableGetCompany200ResponseSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCompany200ResponseSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


