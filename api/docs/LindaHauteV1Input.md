# LindaHauteV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AspectRatio** | **string** |  | 
**ProductImageId** | **string** |  | 
**Prompt** | **string** |  | 
**Seed** | **float32** |  | 
**ImageWeight** | Pointer to **float32** |  | [optional] [default to 0.5]
**NegativePrompt** | Pointer to **string** |  | [optional] [default to "(too many fingers, face asymmetry, eyes asymmetry, deformed eyes, open mouth), worst quality, low quality, illustration, 3d, 2d, painting, cartoons, sketch"]
**InferenceSteps** | Pointer to **float32** |  | [optional] [default to 33]
**GuidanceScale** | Pointer to **float32** |  | [optional] [default to 6]
**Strength** | Pointer to **float32** |  | [optional] [default to 0.3]

## Methods

### NewLindaHauteV1Input

`func NewLindaHauteV1Input(aspectRatio string, productImageId string, prompt string, seed float32, ) *LindaHauteV1Input`

NewLindaHauteV1Input instantiates a new LindaHauteV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLindaHauteV1InputWithDefaults

`func NewLindaHauteV1InputWithDefaults() *LindaHauteV1Input`

NewLindaHauteV1InputWithDefaults instantiates a new LindaHauteV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAspectRatio

`func (o *LindaHauteV1Input) GetAspectRatio() string`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *LindaHauteV1Input) GetAspectRatioOk() (*string, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *LindaHauteV1Input) SetAspectRatio(v string)`

SetAspectRatio sets AspectRatio field to given value.


### GetProductImageId

`func (o *LindaHauteV1Input) GetProductImageId() string`

GetProductImageId returns the ProductImageId field if non-nil, zero value otherwise.

### GetProductImageIdOk

`func (o *LindaHauteV1Input) GetProductImageIdOk() (*string, bool)`

GetProductImageIdOk returns a tuple with the ProductImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductImageId

`func (o *LindaHauteV1Input) SetProductImageId(v string)`

SetProductImageId sets ProductImageId field to given value.


### GetPrompt

`func (o *LindaHauteV1Input) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *LindaHauteV1Input) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *LindaHauteV1Input) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetSeed

`func (o *LindaHauteV1Input) GetSeed() float32`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *LindaHauteV1Input) GetSeedOk() (*float32, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *LindaHauteV1Input) SetSeed(v float32)`

SetSeed sets Seed field to given value.


### GetImageWeight

`func (o *LindaHauteV1Input) GetImageWeight() float32`

GetImageWeight returns the ImageWeight field if non-nil, zero value otherwise.

### GetImageWeightOk

`func (o *LindaHauteV1Input) GetImageWeightOk() (*float32, bool)`

GetImageWeightOk returns a tuple with the ImageWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageWeight

`func (o *LindaHauteV1Input) SetImageWeight(v float32)`

SetImageWeight sets ImageWeight field to given value.

### HasImageWeight

`func (o *LindaHauteV1Input) HasImageWeight() bool`

HasImageWeight returns a boolean if a field has been set.

### GetNegativePrompt

`func (o *LindaHauteV1Input) GetNegativePrompt() string`

GetNegativePrompt returns the NegativePrompt field if non-nil, zero value otherwise.

### GetNegativePromptOk

`func (o *LindaHauteV1Input) GetNegativePromptOk() (*string, bool)`

GetNegativePromptOk returns a tuple with the NegativePrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativePrompt

`func (o *LindaHauteV1Input) SetNegativePrompt(v string)`

SetNegativePrompt sets NegativePrompt field to given value.

### HasNegativePrompt

`func (o *LindaHauteV1Input) HasNegativePrompt() bool`

HasNegativePrompt returns a boolean if a field has been set.

### GetInferenceSteps

`func (o *LindaHauteV1Input) GetInferenceSteps() float32`

GetInferenceSteps returns the InferenceSteps field if non-nil, zero value otherwise.

### GetInferenceStepsOk

`func (o *LindaHauteV1Input) GetInferenceStepsOk() (*float32, bool)`

GetInferenceStepsOk returns a tuple with the InferenceSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceSteps

`func (o *LindaHauteV1Input) SetInferenceSteps(v float32)`

SetInferenceSteps sets InferenceSteps field to given value.

### HasInferenceSteps

`func (o *LindaHauteV1Input) HasInferenceSteps() bool`

HasInferenceSteps returns a boolean if a field has been set.

### GetGuidanceScale

`func (o *LindaHauteV1Input) GetGuidanceScale() float32`

GetGuidanceScale returns the GuidanceScale field if non-nil, zero value otherwise.

### GetGuidanceScaleOk

`func (o *LindaHauteV1Input) GetGuidanceScaleOk() (*float32, bool)`

GetGuidanceScaleOk returns a tuple with the GuidanceScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidanceScale

`func (o *LindaHauteV1Input) SetGuidanceScale(v float32)`

SetGuidanceScale sets GuidanceScale field to given value.

### HasGuidanceScale

`func (o *LindaHauteV1Input) HasGuidanceScale() bool`

HasGuidanceScale returns a boolean if a field has been set.

### GetStrength

`func (o *LindaHauteV1Input) GetStrength() float32`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *LindaHauteV1Input) GetStrengthOk() (*float32, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *LindaHauteV1Input) SetStrength(v float32)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *LindaHauteV1Input) HasStrength() bool`

HasStrength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


