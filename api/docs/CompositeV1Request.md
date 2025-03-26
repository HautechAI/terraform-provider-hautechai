# CompositeV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | [**CompositeV1Input**](CompositeV1Input.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCompositeV1Request

`func NewCompositeV1Request(input CompositeV1Input, ) *CompositeV1Request`

NewCompositeV1Request instantiates a new CompositeV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompositeV1RequestWithDefaults

`func NewCompositeV1RequestWithDefaults() *CompositeV1Request`

NewCompositeV1RequestWithDefaults instantiates a new CompositeV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *CompositeV1Request) GetInput() CompositeV1Input`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *CompositeV1Request) GetInputOk() (*CompositeV1Input, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *CompositeV1Request) SetInput(v CompositeV1Input)`

SetInput sets Input field to given value.


### GetMetadata

`func (o *CompositeV1Request) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CompositeV1Request) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CompositeV1Request) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CompositeV1Request) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


