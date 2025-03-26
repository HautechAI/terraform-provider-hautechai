/*
Hautech API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hautechapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the VtonGiseleV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VtonGiseleV1Response{}

// VtonGiseleV1Response struct for VtonGiseleV1Response
type VtonGiseleV1Response struct {
	Kind string `json:"kind"`
	Output NullableOperationOutputImageSingle `json:"output"`
	Input map[string]interface{} `json:"input"`
	Status string `json:"status"`
	Type string `json:"type"`
	Id string `json:"id"`
	CreatorId string `json:"creatorId"`
	Metadata map[string]interface{} `json:"metadata"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type _VtonGiseleV1Response VtonGiseleV1Response

// NewVtonGiseleV1Response instantiates a new VtonGiseleV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVtonGiseleV1Response(kind string, output NullableOperationOutputImageSingle, input map[string]interface{}, status string, type_ string, id string, creatorId string, metadata map[string]interface{}, createdAt time.Time, updatedAt time.Time) *VtonGiseleV1Response {
	this := VtonGiseleV1Response{}
	this.Kind = kind
	this.Output = output
	this.Input = input
	this.Status = status
	this.Type = type_
	this.Id = id
	this.CreatorId = creatorId
	this.Metadata = metadata
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewVtonGiseleV1ResponseWithDefaults instantiates a new VtonGiseleV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVtonGiseleV1ResponseWithDefaults() *VtonGiseleV1Response {
	this := VtonGiseleV1Response{}
	var kind string = "operation"
	this.Kind = kind
	return &this
}

// GetKind returns the Kind field value
func (o *VtonGiseleV1Response) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *VtonGiseleV1Response) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *VtonGiseleV1Response) SetKind(v string) {
	o.Kind = v
}

// GetOutput returns the Output field value
// If the value is explicit nil, the zero value for OperationOutputImageSingle will be returned
func (o *VtonGiseleV1Response) GetOutput() OperationOutputImageSingle {
	if o == nil || o.Output.Get() == nil {
		var ret OperationOutputImageSingle
		return ret
	}

	return *o.Output.Get()
}

// GetOutputOk returns a tuple with the Output field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VtonGiseleV1Response) GetOutputOk() (*OperationOutputImageSingle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Output.Get(), o.Output.IsSet()
}

// SetOutput sets field value
func (o *VtonGiseleV1Response) SetOutput(v OperationOutputImageSingle) {
	o.Output.Set(&v)
}

// GetInput returns the Input field value
func (o *VtonGiseleV1Response) GetInput() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *VtonGiseleV1Response) GetInputOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Input, true
}

// SetInput sets field value
func (o *VtonGiseleV1Response) SetInput(v map[string]interface{}) {
	o.Input = v
}

// GetStatus returns the Status field value
func (o *VtonGiseleV1Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VtonGiseleV1Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VtonGiseleV1Response) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *VtonGiseleV1Response) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VtonGiseleV1Response) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VtonGiseleV1Response) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *VtonGiseleV1Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VtonGiseleV1Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VtonGiseleV1Response) SetId(v string) {
	o.Id = v
}

// GetCreatorId returns the CreatorId field value
func (o *VtonGiseleV1Response) GetCreatorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *VtonGiseleV1Response) GetCreatorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *VtonGiseleV1Response) SetCreatorId(v string) {
	o.CreatorId = v
}

// GetMetadata returns the Metadata field value
func (o *VtonGiseleV1Response) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *VtonGiseleV1Response) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *VtonGiseleV1Response) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *VtonGiseleV1Response) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *VtonGiseleV1Response) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *VtonGiseleV1Response) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *VtonGiseleV1Response) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *VtonGiseleV1Response) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *VtonGiseleV1Response) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o VtonGiseleV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VtonGiseleV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["output"] = o.Output.Get()
	toSerialize["input"] = o.Input
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["creatorId"] = o.CreatorId
	toSerialize["metadata"] = o.Metadata
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *VtonGiseleV1Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kind",
		"output",
		"input",
		"status",
		"type",
		"id",
		"creatorId",
		"metadata",
		"createdAt",
		"updatedAt",
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

	varVtonGiseleV1Response := _VtonGiseleV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVtonGiseleV1Response)

	if err != nil {
		return err
	}

	*o = VtonGiseleV1Response(varVtonGiseleV1Response)

	return err
}

type NullableVtonGiseleV1Response struct {
	value *VtonGiseleV1Response
	isSet bool
}

func (v NullableVtonGiseleV1Response) Get() *VtonGiseleV1Response {
	return v.value
}

func (v *NullableVtonGiseleV1Response) Set(val *VtonGiseleV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableVtonGiseleV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableVtonGiseleV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVtonGiseleV1Response(val *VtonGiseleV1Response) *NullableVtonGiseleV1Response {
	return &NullableVtonGiseleV1Response{value: val, isSet: true}
}

func (v NullableVtonGiseleV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVtonGiseleV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


