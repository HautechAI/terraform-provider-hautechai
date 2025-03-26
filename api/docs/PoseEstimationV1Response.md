# PoseEstimationV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | [default to "operation"]
**Output** | [**NullableOperationOutputJSON**](OperationOutputJSON.md) |  | 
**Input** | **map[string]interface{}** |  | 
**Status** | **string** |  | 
**Type** | **string** |  | 
**Id** | **string** |  | 
**CreatorId** | **string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewPoseEstimationV1Response

`func NewPoseEstimationV1Response(kind string, output NullableOperationOutputJSON, input map[string]interface{}, status string, type_ string, id string, creatorId string, metadata map[string]interface{}, createdAt time.Time, updatedAt time.Time, ) *PoseEstimationV1Response`

NewPoseEstimationV1Response instantiates a new PoseEstimationV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoseEstimationV1ResponseWithDefaults

`func NewPoseEstimationV1ResponseWithDefaults() *PoseEstimationV1Response`

NewPoseEstimationV1ResponseWithDefaults instantiates a new PoseEstimationV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *PoseEstimationV1Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PoseEstimationV1Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PoseEstimationV1Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetOutput

`func (o *PoseEstimationV1Response) GetOutput() OperationOutputJSON`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *PoseEstimationV1Response) GetOutputOk() (*OperationOutputJSON, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *PoseEstimationV1Response) SetOutput(v OperationOutputJSON)`

SetOutput sets Output field to given value.


### SetOutputNil

`func (o *PoseEstimationV1Response) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *PoseEstimationV1Response) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetInput

`func (o *PoseEstimationV1Response) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *PoseEstimationV1Response) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *PoseEstimationV1Response) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.


### GetStatus

`func (o *PoseEstimationV1Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PoseEstimationV1Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PoseEstimationV1Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *PoseEstimationV1Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PoseEstimationV1Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PoseEstimationV1Response) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PoseEstimationV1Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoseEstimationV1Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoseEstimationV1Response) SetId(v string)`

SetId sets Id field to given value.


### GetCreatorId

`func (o *PoseEstimationV1Response) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *PoseEstimationV1Response) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *PoseEstimationV1Response) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.


### GetMetadata

`func (o *PoseEstimationV1Response) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PoseEstimationV1Response) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PoseEstimationV1Response) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetCreatedAt

`func (o *PoseEstimationV1Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PoseEstimationV1Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PoseEstimationV1Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PoseEstimationV1Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PoseEstimationV1Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PoseEstimationV1Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


