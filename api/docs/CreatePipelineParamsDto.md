# CreatePipelineParamsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Tasks** | **map[string]interface{}** |  | 

## Methods

### NewCreatePipelineParamsDto

`func NewCreatePipelineParamsDto(tasks map[string]interface{}, ) *CreatePipelineParamsDto`

NewCreatePipelineParamsDto instantiates a new CreatePipelineParamsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePipelineParamsDtoWithDefaults

`func NewCreatePipelineParamsDtoWithDefaults() *CreatePipelineParamsDto`

NewCreatePipelineParamsDtoWithDefaults instantiates a new CreatePipelineParamsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *CreatePipelineParamsDto) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreatePipelineParamsDto) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreatePipelineParamsDto) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreatePipelineParamsDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTasks

`func (o *CreatePipelineParamsDto) GetTasks() map[string]interface{}`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *CreatePipelineParamsDto) GetTasksOk() (*map[string]interface{}, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *CreatePipelineParamsDto) SetTasks(v map[string]interface{})`

SetTasks sets Tasks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


