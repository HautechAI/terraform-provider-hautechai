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

// checks if the NegateImageV1Input type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NegateImageV1Input{}

// NegateImageV1Input struct for NegateImageV1Input
type NegateImageV1Input struct {
	ImageId string `json:"imageId"`
}

type _NegateImageV1Input NegateImageV1Input

// NewNegateImageV1Input instantiates a new NegateImageV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNegateImageV1Input(imageId string) *NegateImageV1Input {
	this := NegateImageV1Input{}
	this.ImageId = imageId
	return &this
}

// NewNegateImageV1InputWithDefaults instantiates a new NegateImageV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNegateImageV1InputWithDefaults() *NegateImageV1Input {
	this := NegateImageV1Input{}
	return &this
}

// GetImageId returns the ImageId field value
func (o *NegateImageV1Input) GetImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value
// and a boolean to check if the value has been set.
func (o *NegateImageV1Input) GetImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageId, true
}

// SetImageId sets field value
func (o *NegateImageV1Input) SetImageId(v string) {
	o.ImageId = v
}

func (o NegateImageV1Input) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NegateImageV1Input) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["imageId"] = o.ImageId
	return toSerialize, nil
}

func (o *NegateImageV1Input) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"imageId",
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

	varNegateImageV1Input := _NegateImageV1Input{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNegateImageV1Input)

	if err != nil {
		return err
	}

	*o = NegateImageV1Input(varNegateImageV1Input)

	return err
}

type NullableNegateImageV1Input struct {
	value *NegateImageV1Input
	isSet bool
}

func (v NullableNegateImageV1Input) Get() *NegateImageV1Input {
	return v.value
}

func (v *NullableNegateImageV1Input) Set(val *NegateImageV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableNegateImageV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableNegateImageV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNegateImageV1Input(val *NegateImageV1Input) *NullableNegateImageV1Input {
	return &NullableNegateImageV1Input{value: val, isSet: true}
}

func (v NullableNegateImageV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNegateImageV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


