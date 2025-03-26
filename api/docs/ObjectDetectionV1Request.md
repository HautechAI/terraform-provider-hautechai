# ObjectDetectionV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | [**ObjectDetectionV1Input**](ObjectDetectionV1Input.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewObjectDetectionV1Request

`func NewObjectDetectionV1Request(input ObjectDetectionV1Input, ) *ObjectDetectionV1Request`

NewObjectDetectionV1Request instantiates a new ObjectDetectionV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectDetectionV1RequestWithDefaults

`func NewObjectDetectionV1RequestWithDefaults() *ObjectDetectionV1Request`

NewObjectDetectionV1RequestWithDefaults instantiates a new ObjectDetectionV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *ObjectDetectionV1Request) GetInput() ObjectDetectionV1Input`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ObjectDetectionV1Request) GetInputOk() (*ObjectDetectionV1Input, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ObjectDetectionV1Request) SetInput(v ObjectDetectionV1Input)`

SetInput sets Input field to given value.


### GetMetadata

`func (o *ObjectDetectionV1Request) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ObjectDetectionV1Request) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ObjectDetectionV1Request) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ObjectDetectionV1Request) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


