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

// checks if the ListStacksDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListStacksDto{}

// ListStacksDto struct for ListStacksDto
type ListStacksDto struct {
	Data []StackEntity `json:"data"`
	PageInfo ListAccountsDtoPageInfo `json:"pageInfo"`
}

type _ListStacksDto ListStacksDto

// NewListStacksDto instantiates a new ListStacksDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListStacksDto(data []StackEntity, pageInfo ListAccountsDtoPageInfo) *ListStacksDto {
	this := ListStacksDto{}
	this.Data = data
	this.PageInfo = pageInfo
	return &this
}

// NewListStacksDtoWithDefaults instantiates a new ListStacksDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListStacksDtoWithDefaults() *ListStacksDto {
	this := ListStacksDto{}
	return &this
}

// GetData returns the Data field value
func (o *ListStacksDto) GetData() []StackEntity {
	if o == nil {
		var ret []StackEntity
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListStacksDto) GetDataOk() ([]StackEntity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListStacksDto) SetData(v []StackEntity) {
	o.Data = v
}

// GetPageInfo returns the PageInfo field value
func (o *ListStacksDto) GetPageInfo() ListAccountsDtoPageInfo {
	if o == nil {
		var ret ListAccountsDtoPageInfo
		return ret
	}

	return o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value
// and a boolean to check if the value has been set.
func (o *ListStacksDto) GetPageInfoOk() (*ListAccountsDtoPageInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageInfo, true
}

// SetPageInfo sets field value
func (o *ListStacksDto) SetPageInfo(v ListAccountsDtoPageInfo) {
	o.PageInfo = v
}

func (o ListStacksDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListStacksDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["pageInfo"] = o.PageInfo
	return toSerialize, nil
}

func (o *ListStacksDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"pageInfo",
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

	varListStacksDto := _ListStacksDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListStacksDto)

	if err != nil {
		return err
	}

	*o = ListStacksDto(varListStacksDto)

	return err
}

type NullableListStacksDto struct {
	value *ListStacksDto
	isSet bool
}

func (v NullableListStacksDto) Get() *ListStacksDto {
	return v.value
}

func (v *NullableListStacksDto) Set(val *ListStacksDto) {
	v.value = val
	v.isSet = true
}

func (v NullableListStacksDto) IsSet() bool {
	return v.isSet
}

func (v *NullableListStacksDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListStacksDto(val *ListStacksDto) *NullableListStacksDto {
	return &NullableListStacksDto{value: val, isSet: true}
}

func (v NullableListStacksDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListStacksDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


