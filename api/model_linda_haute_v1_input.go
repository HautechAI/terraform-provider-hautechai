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

// checks if the LindaHauteV1Input type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LindaHauteV1Input{}

// LindaHauteV1Input struct for LindaHauteV1Input
type LindaHauteV1Input struct {
	AspectRatio string `json:"aspectRatio"`
	ProductImageId string `json:"productImageId"`
	Prompt string `json:"prompt"`
	Seed float32 `json:"seed"`
	ImageWeight *float32 `json:"imageWeight,omitempty"`
	NegativePrompt *string `json:"negativePrompt,omitempty"`
	InferenceSteps *float32 `json:"inferenceSteps,omitempty"`
	GuidanceScale *float32 `json:"guidanceScale,omitempty"`
	Strength *float32 `json:"strength,omitempty"`
}

type _LindaHauteV1Input LindaHauteV1Input

// NewLindaHauteV1Input instantiates a new LindaHauteV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLindaHauteV1Input(aspectRatio string, productImageId string, prompt string, seed float32) *LindaHauteV1Input {
	this := LindaHauteV1Input{}
	this.AspectRatio = aspectRatio
	this.ProductImageId = productImageId
	this.Prompt = prompt
	this.Seed = seed
	var imageWeight float32 = 0.5
	this.ImageWeight = &imageWeight
	var negativePrompt string = "(too many fingers, face asymmetry, eyes asymmetry, deformed eyes, open mouth), worst quality, low quality, illustration, 3d, 2d, painting, cartoons, sketch"
	this.NegativePrompt = &negativePrompt
	var inferenceSteps float32 = 33
	this.InferenceSteps = &inferenceSteps
	var guidanceScale float32 = 6
	this.GuidanceScale = &guidanceScale
	var strength float32 = 0.3
	this.Strength = &strength
	return &this
}

// NewLindaHauteV1InputWithDefaults instantiates a new LindaHauteV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLindaHauteV1InputWithDefaults() *LindaHauteV1Input {
	this := LindaHauteV1Input{}
	var imageWeight float32 = 0.5
	this.ImageWeight = &imageWeight
	var negativePrompt string = "(too many fingers, face asymmetry, eyes asymmetry, deformed eyes, open mouth), worst quality, low quality, illustration, 3d, 2d, painting, cartoons, sketch"
	this.NegativePrompt = &negativePrompt
	var inferenceSteps float32 = 33
	this.InferenceSteps = &inferenceSteps
	var guidanceScale float32 = 6
	this.GuidanceScale = &guidanceScale
	var strength float32 = 0.3
	this.Strength = &strength
	return &this
}

// GetAspectRatio returns the AspectRatio field value
func (o *LindaHauteV1Input) GetAspectRatio() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AspectRatio
}

// GetAspectRatioOk returns a tuple with the AspectRatio field value
// and a boolean to check if the value has been set.
func (o *LindaHauteV1Input) GetAspectRatioOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AspectRatio, true
}

// SetAspectRatio sets field value
func (o *LindaHauteV1Input) SetAspectRatio(v string) {
	o.AspectRatio = v
}

// GetProductImageId returns the ProductImageId field value
func (o *LindaHauteV1Input) GetProductImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductImageId
}

// GetProductImageIdOk returns a tuple with the ProductImageId field value
// and a boolean to check if the value has been set.
func (o *LindaHauteV1Input) GetProductImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductImageId, true
}

// SetProductImageId sets field value
func (o *LindaHauteV1Input) SetProductImageId(v string) {
	o.ProductImageId = v
}

// GetPrompt returns the Prompt field value
func (o *LindaHauteV1Input) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *LindaHauteV1Input) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *LindaHauteV1Input) SetPrompt(v string) {
	o.Prompt = v
}

// GetSeed returns the Seed field value
func (o *LindaHauteV1Input) GetSeed() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Seed
}

// GetSeedOk returns a tuple with the Seed field value
// and a boolean to check if the value has been set.
func (o *LindaHauteV1Input) GetSeedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seed, true
}

// SetSeed sets field value
func (o *LindaHauteV1Input) SetSeed(v float32) {
	o.Seed = v
}

// GetImageWeight returns the ImageWeight field value if set, zero value otherwise.
func (o *LindaHauteV1Input) GetImageWeight() float32 {
	if o == nil || IsNil(o.ImageWeight) {
		var ret float32
		return ret
	}
	return *o.ImageWeight
}

// GetImageWeightOk returns a tuple with the ImageWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LindaHauteV1Input) GetImageWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.ImageWeight) {
		return nil, false
	}
	return o.ImageWeight, true
}

// HasImageWeight returns a boolean if a field has been set.
func (o *LindaHauteV1Input) HasImageWeight() bool {
	if o != nil && !IsNil(o.ImageWeight) {
		return true
	}

	return false
}

// SetImageWeight gets a reference to the given float32 and assigns it to the ImageWeight field.
func (o *LindaHauteV1Input) SetImageWeight(v float32) {
	o.ImageWeight = &v
}

// GetNegativePrompt returns the NegativePrompt field value if set, zero value otherwise.
func (o *LindaHauteV1Input) GetNegativePrompt() string {
	if o == nil || IsNil(o.NegativePrompt) {
		var ret string
		return ret
	}
	return *o.NegativePrompt
}

// GetNegativePromptOk returns a tuple with the NegativePrompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LindaHauteV1Input) GetNegativePromptOk() (*string, bool) {
	if o == nil || IsNil(o.NegativePrompt) {
		return nil, false
	}
	return o.NegativePrompt, true
}

// HasNegativePrompt returns a boolean if a field has been set.
func (o *LindaHauteV1Input) HasNegativePrompt() bool {
	if o != nil && !IsNil(o.NegativePrompt) {
		return true
	}

	return false
}

// SetNegativePrompt gets a reference to the given string and assigns it to the NegativePrompt field.
func (o *LindaHauteV1Input) SetNegativePrompt(v string) {
	o.NegativePrompt = &v
}

// GetInferenceSteps returns the InferenceSteps field value if set, zero value otherwise.
func (o *LindaHauteV1Input) GetInferenceSteps() float32 {
	if o == nil || IsNil(o.InferenceSteps) {
		var ret float32
		return ret
	}
	return *o.InferenceSteps
}

// GetInferenceStepsOk returns a tuple with the InferenceSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LindaHauteV1Input) GetInferenceStepsOk() (*float32, bool) {
	if o == nil || IsNil(o.InferenceSteps) {
		return nil, false
	}
	return o.InferenceSteps, true
}

// HasInferenceSteps returns a boolean if a field has been set.
func (o *LindaHauteV1Input) HasInferenceSteps() bool {
	if o != nil && !IsNil(o.InferenceSteps) {
		return true
	}

	return false
}

// SetInferenceSteps gets a reference to the given float32 and assigns it to the InferenceSteps field.
func (o *LindaHauteV1Input) SetInferenceSteps(v float32) {
	o.InferenceSteps = &v
}

// GetGuidanceScale returns the GuidanceScale field value if set, zero value otherwise.
func (o *LindaHauteV1Input) GetGuidanceScale() float32 {
	if o == nil || IsNil(o.GuidanceScale) {
		var ret float32
		return ret
	}
	return *o.GuidanceScale
}

// GetGuidanceScaleOk returns a tuple with the GuidanceScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LindaHauteV1Input) GetGuidanceScaleOk() (*float32, bool) {
	if o == nil || IsNil(o.GuidanceScale) {
		return nil, false
	}
	return o.GuidanceScale, true
}

// HasGuidanceScale returns a boolean if a field has been set.
func (o *LindaHauteV1Input) HasGuidanceScale() bool {
	if o != nil && !IsNil(o.GuidanceScale) {
		return true
	}

	return false
}

// SetGuidanceScale gets a reference to the given float32 and assigns it to the GuidanceScale field.
func (o *LindaHauteV1Input) SetGuidanceScale(v float32) {
	o.GuidanceScale = &v
}

// GetStrength returns the Strength field value if set, zero value otherwise.
func (o *LindaHauteV1Input) GetStrength() float32 {
	if o == nil || IsNil(o.Strength) {
		var ret float32
		return ret
	}
	return *o.Strength
}

// GetStrengthOk returns a tuple with the Strength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LindaHauteV1Input) GetStrengthOk() (*float32, bool) {
	if o == nil || IsNil(o.Strength) {
		return nil, false
	}
	return o.Strength, true
}

// HasStrength returns a boolean if a field has been set.
func (o *LindaHauteV1Input) HasStrength() bool {
	if o != nil && !IsNil(o.Strength) {
		return true
	}

	return false
}

// SetStrength gets a reference to the given float32 and assigns it to the Strength field.
func (o *LindaHauteV1Input) SetStrength(v float32) {
	o.Strength = &v
}

func (o LindaHauteV1Input) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LindaHauteV1Input) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aspectRatio"] = o.AspectRatio
	toSerialize["productImageId"] = o.ProductImageId
	toSerialize["prompt"] = o.Prompt
	toSerialize["seed"] = o.Seed
	if !IsNil(o.ImageWeight) {
		toSerialize["imageWeight"] = o.ImageWeight
	}
	if !IsNil(o.NegativePrompt) {
		toSerialize["negativePrompt"] = o.NegativePrompt
	}
	if !IsNil(o.InferenceSteps) {
		toSerialize["inferenceSteps"] = o.InferenceSteps
	}
	if !IsNil(o.GuidanceScale) {
		toSerialize["guidanceScale"] = o.GuidanceScale
	}
	if !IsNil(o.Strength) {
		toSerialize["strength"] = o.Strength
	}
	return toSerialize, nil
}

func (o *LindaHauteV1Input) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"aspectRatio",
		"productImageId",
		"prompt",
		"seed",
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

	varLindaHauteV1Input := _LindaHauteV1Input{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLindaHauteV1Input)

	if err != nil {
		return err
	}

	*o = LindaHauteV1Input(varLindaHauteV1Input)

	return err
}

type NullableLindaHauteV1Input struct {
	value *LindaHauteV1Input
	isSet bool
}

func (v NullableLindaHauteV1Input) Get() *LindaHauteV1Input {
	return v.value
}

func (v *NullableLindaHauteV1Input) Set(val *LindaHauteV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableLindaHauteV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableLindaHauteV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLindaHauteV1Input(val *LindaHauteV1Input) *NullableLindaHauteV1Input {
	return &NullableLindaHauteV1Input{value: val, isSet: true}
}

func (v NullableLindaHauteV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLindaHauteV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


