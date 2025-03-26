/*
Hautech API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hautechapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetBalanceParamsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBalanceParamsDto{}

// GetBalanceParamsDto struct for GetBalanceParamsDto
type GetBalanceParamsDto struct {
	AccountId string `json:"accountId"`
}

type _GetBalanceParamsDto GetBalanceParamsDto

// NewGetBalanceParamsDto instantiates a new GetBalanceParamsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalanceParamsDto(accountId string) *GetBalanceParamsDto {
	this := GetBalanceParamsDto{}
	this.AccountId = accountId
	return &this
}

// NewGetBalanceParamsDtoWithDefaults instantiates a new GetBalanceParamsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalanceParamsDtoWithDefaults() *GetBalanceParamsDto {
	this := GetBalanceParamsDto{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *GetBalanceParamsDto) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *GetBalanceParamsDto) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *GetBalanceParamsDto) SetAccountId(v string) {
	o.AccountId = v
}

func (o GetBalanceParamsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBalanceParamsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountId"] = o.AccountId
	return toSerialize, nil
}

func (o *GetBalanceParamsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accountId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetBalanceParamsDto := _GetBalanceParamsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetBalanceParamsDto)

	if err != nil {
		return err
	}

	*o = GetBalanceParamsDto(varGetBalanceParamsDto)

	return err
}

type NullableGetBalanceParamsDto struct {
	value *GetBalanceParamsDto
	isSet bool
}

func (v NullableGetBalanceParamsDto) Get() *GetBalanceParamsDto {
	return v.value
}

func (v *NullableGetBalanceParamsDto) Set(val *GetBalanceParamsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalanceParamsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalanceParamsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalanceParamsDto(val *GetBalanceParamsDto) *NullableGetBalanceParamsDto {
	return &NullableGetBalanceParamsDto{value: val, isSet: true}
}

func (v NullableGetBalanceParamsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalanceParamsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


