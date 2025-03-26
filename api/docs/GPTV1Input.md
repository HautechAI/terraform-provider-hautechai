# GPTV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** |  | [optional] [default to "gpt-4o"]
**Prompt** | **string** |  | 
**ImageId** | Pointer to **string** |  | [optional] 

## Methods

### NewGPTV1Input

`func NewGPTV1Input(prompt string, ) *GPTV1Input`

NewGPTV1Input instantiates a new GPTV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGPTV1InputWithDefaults

`func NewGPTV1InputWithDefaults() *GPTV1Input`

NewGPTV1InputWithDefaults instantiates a new GPTV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *GPTV1Input) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GPTV1Input) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GPTV1Input) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GPTV1Input) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPrompt

`func (o *GPTV1Input) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *GPTV1Input) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *GPTV1Input) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetImageId

`func (o *GPTV1Input) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *GPTV1Input) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *GPTV1Input) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *GPTV1Input) HasImageId() bool`

HasImageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


