# InpaintKateV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | [**KateInpaintV1Input**](KateInpaintV1Input.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewInpaintKateV1Request

`func NewInpaintKateV1Request(input KateInpaintV1Input, ) *InpaintKateV1Request`

NewInpaintKateV1Request instantiates a new InpaintKateV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInpaintKateV1RequestWithDefaults

`func NewInpaintKateV1RequestWithDefaults() *InpaintKateV1Request`

NewInpaintKateV1RequestWithDefaults instantiates a new InpaintKateV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *InpaintKateV1Request) GetInput() KateInpaintV1Input`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *InpaintKateV1Request) GetInputOk() (*KateInpaintV1Input, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *InpaintKateV1Request) SetInput(v KateInpaintV1Input)`

SetInput sets Input field to given value.


### GetMetadata

`func (o *InpaintKateV1Request) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InpaintKateV1Request) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InpaintKateV1Request) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InpaintKateV1Request) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


