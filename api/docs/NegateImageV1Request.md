# NegateImageV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | [**NegateImageV1Input**](NegateImageV1Input.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewNegateImageV1Request

`func NewNegateImageV1Request(input NegateImageV1Input, ) *NegateImageV1Request`

NewNegateImageV1Request instantiates a new NegateImageV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegateImageV1RequestWithDefaults

`func NewNegateImageV1RequestWithDefaults() *NegateImageV1Request`

NewNegateImageV1RequestWithDefaults instantiates a new NegateImageV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *NegateImageV1Request) GetInput() NegateImageV1Input`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *NegateImageV1Request) GetInputOk() (*NegateImageV1Input, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *NegateImageV1Request) SetInput(v NegateImageV1Input)`

SetInput sets Input field to given value.


### GetMetadata

`func (o *NegateImageV1Request) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NegateImageV1Request) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NegateImageV1Request) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NegateImageV1Request) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


