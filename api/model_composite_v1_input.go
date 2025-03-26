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

// checks if the CompositeV1Input type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompositeV1Input{}

// CompositeV1Input struct for CompositeV1Input
type CompositeV1Input struct {
	Width float32 `json:"width"`
	Height float32 `json:"height"`
	Background string `json:"background" validate:"regexp=^#?([0-9A-F]{3}|[0-9A-F]{4}|[0-9A-F]{6}|[0-9A-F]{8})$"`
	Elements []CompositeElement `json:"elements"`
}

type _CompositeV1Input CompositeV1Input

// NewCompositeV1Input instantiates a new CompositeV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompositeV1Input(width float32, height float32, background string, elements []CompositeElement) *CompositeV1Input {
	this := CompositeV1Input{}
	this.Width = width
	this.Height = height
	this.Background = background
	this.Elements = elements
	return &this
}

// NewCompositeV1InputWithDefaults instantiates a new CompositeV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompositeV1InputWithDefaults() *CompositeV1Input {
	this := CompositeV1Input{}
	return &this
}

// GetWidth returns the Width field value
func (o *CompositeV1Input) GetWidth() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *CompositeV1Input) GetWidthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *CompositeV1Input) SetWidth(v float32) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *CompositeV1Input) GetHeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *CompositeV1Input) GetHeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *CompositeV1Input) SetHeight(v float32) {
	o.Height = v
}

// GetBackground returns the Background field value
func (o *CompositeV1Input) GetBackground() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Background
}

// GetBackgroundOk returns a tuple with the Background field value
// and a boolean to check if the value has been set.
func (o *CompositeV1Input) GetBackgroundOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Background, true
}

// SetBackground sets field value
func (o *CompositeV1Input) SetBackground(v string) {
	o.Background = v
}

// GetElements returns the Elements field value
func (o *CompositeV1Input) GetElements() []CompositeElement {
	if o == nil {
		var ret []CompositeElement
		return ret
	}

	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value
// and a boolean to check if the value has been set.
func (o *CompositeV1Input) GetElementsOk() ([]CompositeElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Elements, true
}

// SetElements sets field value
func (o *CompositeV1Input) SetElements(v []CompositeElement) {
	o.Elements = v
}

func (o CompositeV1Input) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompositeV1Input) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["width"] = o.Width
	toSerialize["height"] = o.Height
	toSerialize["background"] = o.Background
	toSerialize["elements"] = o.Elements
	return toSerialize, nil
}

func (o *CompositeV1Input) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"width",
		"height",
		"background",
		"elements",
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

	varCompositeV1Input := _CompositeV1Input{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompositeV1Input)

	if err != nil {
		return err
	}

	*o = CompositeV1Input(varCompositeV1Input)

	return err
}

type NullableCompositeV1Input struct {
	value *CompositeV1Input
	isSet bool
}

func (v NullableCompositeV1Input) Get() *CompositeV1Input {
	return v.value
}

func (v *NullableCompositeV1Input) Set(val *CompositeV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableCompositeV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableCompositeV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompositeV1Input(val *CompositeV1Input) *NullableCompositeV1Input {
	return &NullableCompositeV1Input{value: val, isSet: true}
}

func (v NullableCompositeV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompositeV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


