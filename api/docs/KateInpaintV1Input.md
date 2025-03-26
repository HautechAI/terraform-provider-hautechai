# KateInpaintV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | **string** |  | 
**MaskImageId** | Pointer to **string** |  | [optional] 
**MaskSpread** | Pointer to **float32** |  | [optional] [default to 0]
**Prompt** | **string** |  | 
**Seed** | **float32** |  | 
**Strength** | **float32** |  | 
**Height** | Pointer to **float32** |  | [optional] 
**Width** | Pointer to **float32** |  | [optional] 
**NumInferenceSteps** | Pointer to **float32** |  | [optional] [default to 30]
**GuidanceScale** | Pointer to **float32** |  | [optional] [default to 7]
**Branch** | Pointer to **string** |  | [optional] [default to "stable"]

## Methods

### NewKateInpaintV1Input

`func NewKateInpaintV1Input(imageId string, prompt string, seed float32, strength float32, ) *KateInpaintV1Input`

NewKateInpaintV1Input instantiates a new KateInpaintV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKateInpaintV1InputWithDefaults

`func NewKateInpaintV1InputWithDefaults() *KateInpaintV1Input`

NewKateInpaintV1InputWithDefaults instantiates a new KateInpaintV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *KateInpaintV1Input) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *KateInpaintV1Input) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *KateInpaintV1Input) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetMaskImageId

`func (o *KateInpaintV1Input) GetMaskImageId() string`

GetMaskImageId returns the MaskImageId field if non-nil, zero value otherwise.

### GetMaskImageIdOk

`func (o *KateInpaintV1Input) GetMaskImageIdOk() (*string, bool)`

GetMaskImageIdOk returns a tuple with the MaskImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskImageId

`func (o *KateInpaintV1Input) SetMaskImageId(v string)`

SetMaskImageId sets MaskImageId field to given value.

### HasMaskImageId

`func (o *KateInpaintV1Input) HasMaskImageId() bool`

HasMaskImageId returns a boolean if a field has been set.

### GetMaskSpread

`func (o *KateInpaintV1Input) GetMaskSpread() float32`

GetMaskSpread returns the MaskSpread field if non-nil, zero value otherwise.

### GetMaskSpreadOk

`func (o *KateInpaintV1Input) GetMaskSpreadOk() (*float32, bool)`

GetMaskSpreadOk returns a tuple with the MaskSpread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskSpread

`func (o *KateInpaintV1Input) SetMaskSpread(v float32)`

SetMaskSpread sets MaskSpread field to given value.

### HasMaskSpread

`func (o *KateInpaintV1Input) HasMaskSpread() bool`

HasMaskSpread returns a boolean if a field has been set.

### GetPrompt

`func (o *KateInpaintV1Input) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *KateInpaintV1Input) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *KateInpaintV1Input) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetSeed

`func (o *KateInpaintV1Input) GetSeed() float32`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *KateInpaintV1Input) GetSeedOk() (*float32, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *KateInpaintV1Input) SetSeed(v float32)`

SetSeed sets Seed field to given value.


### GetStrength

`func (o *KateInpaintV1Input) GetStrength() float32`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *KateInpaintV1Input) GetStrengthOk() (*float32, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *KateInpaintV1Input) SetStrength(v float32)`

SetStrength sets Strength field to given value.


### GetHeight

`func (o *KateInpaintV1Input) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *KateInpaintV1Input) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *KateInpaintV1Input) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *KateInpaintV1Input) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *KateInpaintV1Input) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *KateInpaintV1Input) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *KateInpaintV1Input) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *KateInpaintV1Input) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetNumInferenceSteps

`func (o *KateInpaintV1Input) GetNumInferenceSteps() float32`

GetNumInferenceSteps returns the NumInferenceSteps field if non-nil, zero value otherwise.

### GetNumInferenceStepsOk

`func (o *KateInpaintV1Input) GetNumInferenceStepsOk() (*float32, bool)`

GetNumInferenceStepsOk returns a tuple with the NumInferenceSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInferenceSteps

`func (o *KateInpaintV1Input) SetNumInferenceSteps(v float32)`

SetNumInferenceSteps sets NumInferenceSteps field to given value.

### HasNumInferenceSteps

`func (o *KateInpaintV1Input) HasNumInferenceSteps() bool`

HasNumInferenceSteps returns a boolean if a field has been set.

### GetGuidanceScale

`func (o *KateInpaintV1Input) GetGuidanceScale() float32`

GetGuidanceScale returns the GuidanceScale field if non-nil, zero value otherwise.

### GetGuidanceScaleOk

`func (o *KateInpaintV1Input) GetGuidanceScaleOk() (*float32, bool)`

GetGuidanceScaleOk returns a tuple with the GuidanceScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidanceScale

`func (o *KateInpaintV1Input) SetGuidanceScale(v float32)`

SetGuidanceScale sets GuidanceScale field to given value.

### HasGuidanceScale

`func (o *KateInpaintV1Input) HasGuidanceScale() bool`

HasGuidanceScale returns a boolean if a field has been set.

### GetBranch

`func (o *KateInpaintV1Input) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *KateInpaintV1Input) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *KateInpaintV1Input) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *KateInpaintV1Input) HasBranch() bool`

HasBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


