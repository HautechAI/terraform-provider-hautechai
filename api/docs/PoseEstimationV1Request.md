# PoseEstimationV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | [**PoseEstimationV1Input**](PoseEstimationV1Input.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPoseEstimationV1Request

`func NewPoseEstimationV1Request(input PoseEstimationV1Input, ) *PoseEstimationV1Request`

NewPoseEstimationV1Request instantiates a new PoseEstimationV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoseEstimationV1RequestWithDefaults

`func NewPoseEstimationV1RequestWithDefaults() *PoseEstimationV1Request`

NewPoseEstimationV1RequestWithDefaults instantiates a new PoseEstimationV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *PoseEstimationV1Request) GetInput() PoseEstimationV1Input`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *PoseEstimationV1Request) GetInputOk() (*PoseEstimationV1Input, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *PoseEstimationV1Request) SetInput(v PoseEstimationV1Input)`

SetInput sets Input field to given value.


### GetMetadata

`func (o *PoseEstimationV1Request) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PoseEstimationV1Request) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PoseEstimationV1Request) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PoseEstimationV1Request) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


