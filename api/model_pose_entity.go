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

// checks if the PoseEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoseEntity{}

// PoseEntity struct for PoseEntity
type PoseEntity struct {
	Kind string `json:"kind"`
	SourceImage ImageEntity `json:"sourceImage"`
	PreviewImage *ImageEntity `json:"previewImage,omitempty"`
	PreviewImageId *string `json:"previewImageId,omitempty"`
	Id string `json:"id"`
	CreatorId string `json:"creatorId"`
	Metadata map[string]interface{} `json:"metadata"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type _PoseEntity PoseEntity

// NewPoseEntity instantiates a new PoseEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoseEntity(kind string, sourceImage ImageEntity, id string, creatorId string, metadata map[string]interface{}, createdAt time.Time, updatedAt time.Time) *PoseEntity {
	this := PoseEntity{}
	this.Kind = kind
	this.SourceImage = sourceImage
	this.Id = id
	this.CreatorId = creatorId
	this.Metadata = metadata
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewPoseEntityWithDefaults instantiates a new PoseEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoseEntityWithDefaults() *PoseEntity {
	this := PoseEntity{}
	var kind string = "pose"
	this.Kind = kind
	return &this
}

// GetKind returns the Kind field value
func (o *PoseEntity) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *PoseEntity) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *PoseEntity) SetKind(v string) {
	o.Kind = v
}

// GetSourceImage returns the SourceImage field value
func (o *PoseEntity) GetSourceImage() ImageEntity {
	if o == nil {
		var ret ImageEntity
		return ret
	}

	return o.SourceImage
}

// GetSourceImageOk returns a tuple with the SourceImage field value
// and a boolean to check if the value has been set.
func (o *PoseEntity) GetSourceImageOk() (*ImageEntity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceImage, true
}

// SetSourceImage sets field value
func (o *PoseEntity) SetSourceImage(v ImageEntity) {
	o.SourceImage = v
}

// GetPreviewImage returns the PreviewImage field value if set, zero value otherwise.
func (o *PoseEntity) GetPreviewImage() ImageEntity {
	if o == nil || IsNil(o.PreviewImage) {
		var ret ImageEntity
		return ret
	}
	return *o.PreviewImage
}

// GetPreviewImageOk returns a tuple with the PreviewImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoseEntity) GetPreviewImageOk() (*ImageEntity, bool) {
	if o == nil || IsNil(o.PreviewImage) {
		return nil, false
	}
	return o.PreviewImage, true
}

// HasPreviewImage returns a boolean if a field has been set.
func (o *PoseEntity) HasPreviewImage() bool {
	if o != nil && !IsNil(o.PreviewImage) {
		return true
	}

	return false
}

// SetPreviewImage gets a reference to the given ImageEntity and assigns it to the PreviewImage field.
func (o *PoseEntity) SetPreviewImage(v ImageEntity) {
	o.PreviewImage = &v
}

// GetPreviewImageId returns the PreviewImageId field value if set, zero value otherwise.
func (o *PoseEntity) GetPreviewImageId() string {
	if o == nil || IsNil(o.PreviewImageId) {
		var ret string
		return ret
	}
	return *o.PreviewImageId
}

// GetPreviewImageIdOk returns a tuple with the PreviewImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoseEntity) GetPreviewImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.PreviewImageId) {
		return nil, false
	}
	return o.PreviewImageId, true
}

// HasPreviewImageId returns a boolean if a field has been set.
func (o *PoseEntity) HasPreviewImageId() bool {
	if o != nil && !IsNil(o.PreviewImageId) {
		return true
	}

	return false
}

// SetPreviewImageId gets a reference to the given string and assigns it to the PreviewImageId field.
func (o *PoseEntity) SetPreviewImageId(v string) {
	o.PreviewImageId = &v
}

// GetId returns the Id field value
func (o *PoseEntity) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PoseEntity) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PoseEntity) SetId(v string) {
	o.Id = v
}

// GetCreatorId returns the CreatorId field value
func (o *PoseEntity) GetCreatorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *PoseEntity) GetCreatorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *PoseEntity) SetCreatorId(v string) {
	o.CreatorId = v
}

// GetMetadata returns the Metadata field value
func (o *PoseEntity) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *PoseEntity) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *PoseEntity) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PoseEntity) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PoseEntity) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PoseEntity) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PoseEntity) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PoseEntity) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PoseEntity) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o PoseEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoseEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["sourceImage"] = o.SourceImage
	if !IsNil(o.PreviewImage) {
		toSerialize["previewImage"] = o.PreviewImage
	}
	if !IsNil(o.PreviewImageId) {
		toSerialize["previewImageId"] = o.PreviewImageId
	}
	toSerialize["id"] = o.Id
	toSerialize["creatorId"] = o.CreatorId
	toSerialize["metadata"] = o.Metadata
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *PoseEntity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kind",
		"sourceImage",
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

	varPoseEntity := _PoseEntity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPoseEntity)

	if err != nil {
		return err
	}

	*o = PoseEntity(varPoseEntity)

	return err
}

type NullablePoseEntity struct {
	value *PoseEntity
	isSet bool
}

func (v NullablePoseEntity) Get() *PoseEntity {
	return v.value
}

func (v *NullablePoseEntity) Set(val *PoseEntity) {
	v.value = val
	v.isSet = true
}

func (v NullablePoseEntity) IsSet() bool {
	return v.isSet
}

func (v *NullablePoseEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoseEntity(val *PoseEntity) *NullablePoseEntity {
	return &NullablePoseEntity{value: val, isSet: true}
}

func (v NullablePoseEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoseEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


