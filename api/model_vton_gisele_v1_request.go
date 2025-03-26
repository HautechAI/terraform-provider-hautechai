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

// checks if the VtonGiseleV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VtonGiseleV1Request{}

// VtonGiseleV1Request struct for VtonGiseleV1Request
type VtonGiseleV1Request struct {
	Input GiseleVtonV1Input `json:"input"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type _VtonGiseleV1Request VtonGiseleV1Request

// NewVtonGiseleV1Request instantiates a new VtonGiseleV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVtonGiseleV1Request(input GiseleVtonV1Input) *VtonGiseleV1Request {
	this := VtonGiseleV1Request{}
	this.Input = input
	return &this
}

// NewVtonGiseleV1RequestWithDefaults instantiates a new VtonGiseleV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVtonGiseleV1RequestWithDefaults() *VtonGiseleV1Request {
	this := VtonGiseleV1Request{}
	return &this
}

// GetInput returns the Input field value
func (o *VtonGiseleV1Request) GetInput() GiseleVtonV1Input {
	if o == nil {
		var ret GiseleVtonV1Input
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *VtonGiseleV1Request) GetInputOk() (*GiseleVtonV1Input, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *VtonGiseleV1Request) SetInput(v GiseleVtonV1Input) {
	o.Input = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *VtonGiseleV1Request) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VtonGiseleV1Request) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *VtonGiseleV1Request) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *VtonGiseleV1Request) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o VtonGiseleV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VtonGiseleV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["input"] = o.Input
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *VtonGiseleV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"input",
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

	varVtonGiseleV1Request := _VtonGiseleV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVtonGiseleV1Request)

	if err != nil {
		return err
	}

	*o = VtonGiseleV1Request(varVtonGiseleV1Request)

	return err
}

type NullableVtonGiseleV1Request struct {
	value *VtonGiseleV1Request
	isSet bool
}

func (v NullableVtonGiseleV1Request) Get() *VtonGiseleV1Request {
	return v.value
}

func (v *NullableVtonGiseleV1Request) Set(val *VtonGiseleV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableVtonGiseleV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableVtonGiseleV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVtonGiseleV1Request(val *VtonGiseleV1Request) *NullableVtonGiseleV1Request {
	return &NullableVtonGiseleV1Request{value: val, isSet: true}
}

func (v NullableVtonGiseleV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVtonGiseleV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


