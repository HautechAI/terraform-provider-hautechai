# UpdateResourceMetadataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Update** | [**UpdateMetadataDto**](UpdateMetadataDto.md) |  | 

## Methods

### NewUpdateResourceMetadataDto

`func NewUpdateResourceMetadataDto(id string, update UpdateMetadataDto, ) *UpdateResourceMetadataDto`

NewUpdateResourceMetadataDto instantiates a new UpdateResourceMetadataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResourceMetadataDtoWithDefaults

`func NewUpdateResourceMetadataDtoWithDefaults() *UpdateResourceMetadataDto`

NewUpdateResourceMetadataDtoWithDefaults instantiates a new UpdateResourceMetadataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateResourceMetadataDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateResourceMetadataDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateResourceMetadataDto) SetId(v string)`

SetId sets Id field to given value.


### GetUpdate

`func (o *UpdateResourceMetadataDto) GetUpdate() UpdateMetadataDto`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *UpdateResourceMetadataDto) GetUpdateOk() (*UpdateMetadataDto, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *UpdateResourceMetadataDto) SetUpdate(v UpdateMetadataDto)`

SetUpdate sets Update field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


