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

// checks if the AddItemsToStackParamsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddItemsToStackParamsDto{}

// AddItemsToStackParamsDto struct for AddItemsToStackParamsDto
type AddItemsToStackParamsDto struct {
	StackId string `json:"stackId"`
	ItemIds []string `json:"itemIds"`
}

type _AddItemsToStackParamsDto AddItemsToStackParamsDto

// NewAddItemsToStackParamsDto instantiates a new AddItemsToStackParamsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddItemsToStackParamsDto(stackId string, itemIds []string) *AddItemsToStackParamsDto {
	this := AddItemsToStackParamsDto{}
	this.StackId = stackId
	this.ItemIds = itemIds
	return &this
}

// NewAddItemsToStackParamsDtoWithDefaults instantiates a new AddItemsToStackParamsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddItemsToStackParamsDtoWithDefaults() *AddItemsToStackParamsDto {
	this := AddItemsToStackParamsDto{}
	return &this
}

// GetStackId returns the StackId field value
func (o *AddItemsToStackParamsDto) GetStackId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StackId
}

// GetStackIdOk returns a tuple with the StackId field value
// and a boolean to check if the value has been set.
func (o *AddItemsToStackParamsDto) GetStackIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StackId, true
}

// SetStackId sets field value
func (o *AddItemsToStackParamsDto) SetStackId(v string) {
	o.StackId = v
}

// GetItemIds returns the ItemIds field value
func (o *AddItemsToStackParamsDto) GetItemIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ItemIds
}

// GetItemIdsOk returns a tuple with the ItemIds field value
// and a boolean to check if the value has been set.
func (o *AddItemsToStackParamsDto) GetItemIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemIds, true
}

// SetItemIds sets field value
func (o *AddItemsToStackParamsDto) SetItemIds(v []string) {
	o.ItemIds = v
}

func (o AddItemsToStackParamsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddItemsToStackParamsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stackId"] = o.StackId
	toSerialize["itemIds"] = o.ItemIds
	return toSerialize, nil
}

func (o *AddItemsToStackParamsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"stackId",
		"itemIds",
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

	varAddItemsToStackParamsDto := _AddItemsToStackParamsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddItemsToStackParamsDto)

	if err != nil {
		return err
	}

	*o = AddItemsToStackParamsDto(varAddItemsToStackParamsDto)

	return err
}

type NullableAddItemsToStackParamsDto struct {
	value *AddItemsToStackParamsDto
	isSet bool
}

func (v NullableAddItemsToStackParamsDto) Get() *AddItemsToStackParamsDto {
	return v.value
}

func (v *NullableAddItemsToStackParamsDto) Set(val *AddItemsToStackParamsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAddItemsToStackParamsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAddItemsToStackParamsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddItemsToStackParamsDto(val *AddItemsToStackParamsDto) *NullableAddItemsToStackParamsDto {
	return &NullableAddItemsToStackParamsDto{value: val, isSet: true}
}

func (v NullableAddItemsToStackParamsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddItemsToStackParamsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


