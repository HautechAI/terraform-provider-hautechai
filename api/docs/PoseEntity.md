# PoseEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | [default to "pose"]
**SourceImage** | [**ImageEntity**](ImageEntity.md) |  | 
**PreviewImage** | Pointer to [**ImageEntity**](ImageEntity.md) |  | [optional] 
**PreviewImageId** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**CreatorId** | **string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewPoseEntity

`func NewPoseEntity(kind string, sourceImage ImageEntity, id string, creatorId string, metadata map[string]interface{}, createdAt time.Time, updatedAt time.Time, ) *PoseEntity`

NewPoseEntity instantiates a new PoseEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoseEntityWithDefaults

`func NewPoseEntityWithDefaults() *PoseEntity`

NewPoseEntityWithDefaults instantiates a new PoseEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *PoseEntity) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PoseEntity) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PoseEntity) SetKind(v string)`

SetKind sets Kind field to given value.


### GetSourceImage

`func (o *PoseEntity) GetSourceImage() ImageEntity`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *PoseEntity) GetSourceImageOk() (*ImageEntity, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *PoseEntity) SetSourceImage(v ImageEntity)`

SetSourceImage sets SourceImage field to given value.


### GetPreviewImage

`func (o *PoseEntity) GetPreviewImage() ImageEntity`

GetPreviewImage returns the PreviewImage field if non-nil, zero value otherwise.

### GetPreviewImageOk

`func (o *PoseEntity) GetPreviewImageOk() (*ImageEntity, bool)`

GetPreviewImageOk returns a tuple with the PreviewImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewImage

`func (o *PoseEntity) SetPreviewImage(v ImageEntity)`

SetPreviewImage sets PreviewImage field to given value.

### HasPreviewImage

`func (o *PoseEntity) HasPreviewImage() bool`

HasPreviewImage returns a boolean if a field has been set.

### GetPreviewImageId

`func (o *PoseEntity) GetPreviewImageId() string`

GetPreviewImageId returns the PreviewImageId field if non-nil, zero value otherwise.

### GetPreviewImageIdOk

`func (o *PoseEntity) GetPreviewImageIdOk() (*string, bool)`

GetPreviewImageIdOk returns a tuple with the PreviewImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewImageId

`func (o *PoseEntity) SetPreviewImageId(v string)`

SetPreviewImageId sets PreviewImageId field to given value.

### HasPreviewImageId

`func (o *PoseEntity) HasPreviewImageId() bool`

HasPreviewImageId returns a boolean if a field has been set.

### GetId

`func (o *PoseEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoseEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoseEntity) SetId(v string)`

SetId sets Id field to given value.


### GetCreatorId

`func (o *PoseEntity) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *PoseEntity) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *PoseEntity) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.


### GetMetadata

`func (o *PoseEntity) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PoseEntity) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PoseEntity) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetCreatedAt

`func (o *PoseEntity) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PoseEntity) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PoseEntity) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PoseEntity) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PoseEntity) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PoseEntity) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


