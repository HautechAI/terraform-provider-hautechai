# NoiseV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | [**NoiseV1Input**](NoiseV1Input.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewNoiseV1Request

`func NewNoiseV1Request(input NoiseV1Input, ) *NoiseV1Request`

NewNoiseV1Request instantiates a new NoiseV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoiseV1RequestWithDefaults

`func NewNoiseV1RequestWithDefaults() *NoiseV1Request`

NewNoiseV1RequestWithDefaults instantiates a new NoiseV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *NoiseV1Request) GetInput() NoiseV1Input`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *NoiseV1Request) GetInputOk() (*NoiseV1Input, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *NoiseV1Request) SetInput(v NoiseV1Input)`

SetInput sets Input field to given value.


### GetMetadata

`func (o *NoiseV1Request) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NoiseV1Request) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NoiseV1Request) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NoiseV1Request) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


