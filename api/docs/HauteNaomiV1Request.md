# HauteNaomiV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | [**NaomiHauteV1Input**](NaomiHauteV1Input.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewHauteNaomiV1Request

`func NewHauteNaomiV1Request(input NaomiHauteV1Input, ) *HauteNaomiV1Request`

NewHauteNaomiV1Request instantiates a new HauteNaomiV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHauteNaomiV1RequestWithDefaults

`func NewHauteNaomiV1RequestWithDefaults() *HauteNaomiV1Request`

NewHauteNaomiV1RequestWithDefaults instantiates a new HauteNaomiV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *HauteNaomiV1Request) GetInput() NaomiHauteV1Input`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *HauteNaomiV1Request) GetInputOk() (*NaomiHauteV1Input, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *HauteNaomiV1Request) SetInput(v NaomiHauteV1Input)`

SetInput sets Input field to given value.


### GetMetadata

`func (o *HauteNaomiV1Request) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HauteNaomiV1Request) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HauteNaomiV1Request) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HauteNaomiV1Request) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


