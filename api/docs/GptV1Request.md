# GptV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | [**GPTV1Input**](GPTV1Input.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGptV1Request

`func NewGptV1Request(input GPTV1Input, ) *GptV1Request`

NewGptV1Request instantiates a new GptV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGptV1RequestWithDefaults

`func NewGptV1RequestWithDefaults() *GptV1Request`

NewGptV1RequestWithDefaults instantiates a new GptV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *GptV1Request) GetInput() GPTV1Input`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *GptV1Request) GetInputOk() (*GPTV1Input, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *GptV1Request) SetInput(v GPTV1Input)`

SetInput sets Input field to given value.


### GetMetadata

`func (o *GptV1Request) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GptV1Request) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GptV1Request) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GptV1Request) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


