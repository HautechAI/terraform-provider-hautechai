# NaomiHauteV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoraIds** | Pointer to **[]string** | UNSTABLE | [optional] [default to ["6a463153-464d-42f2-867e-76f6fef76281"]]
**Prompt** | **string** |  | 
**Category** | **string** |  | 
**GarmentImageId** | **string** |  | 
**PoseId** | **string** |  | 
**Seed** | **float32** |  | 
**Width** | **float32** |  | 
**Height** | **float32** |  | 
**TextGuidanceScale** | Pointer to **float32** |  | [optional] [default to 2.5]
**ImageGuidanceScale** | Pointer to **float32** |  | [optional] [default to 1.6]
**NumInferenceSteps** | Pointer to **float32** |  | [optional] [default to 50]
**Mode** | Pointer to **string** |  | [optional] [default to "apparel_to_model"]

## Methods

### NewNaomiHauteV1Input

`func NewNaomiHauteV1Input(prompt string, category string, garmentImageId string, poseId string, seed float32, width float32, height float32, ) *NaomiHauteV1Input`

NewNaomiHauteV1Input instantiates a new NaomiHauteV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNaomiHauteV1InputWithDefaults

`func NewNaomiHauteV1InputWithDefaults() *NaomiHauteV1Input`

NewNaomiHauteV1InputWithDefaults instantiates a new NaomiHauteV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoraIds

`func (o *NaomiHauteV1Input) GetLoraIds() []string`

GetLoraIds returns the LoraIds field if non-nil, zero value otherwise.

### GetLoraIdsOk

`func (o *NaomiHauteV1Input) GetLoraIdsOk() (*[]string, bool)`

GetLoraIdsOk returns a tuple with the LoraIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoraIds

`func (o *NaomiHauteV1Input) SetLoraIds(v []string)`

SetLoraIds sets LoraIds field to given value.

### HasLoraIds

`func (o *NaomiHauteV1Input) HasLoraIds() bool`

HasLoraIds returns a boolean if a field has been set.

### GetPrompt

`func (o *NaomiHauteV1Input) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *NaomiHauteV1Input) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *NaomiHauteV1Input) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetCategory

`func (o *NaomiHauteV1Input) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NaomiHauteV1Input) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NaomiHauteV1Input) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetGarmentImageId

`func (o *NaomiHauteV1Input) GetGarmentImageId() string`

GetGarmentImageId returns the GarmentImageId field if non-nil, zero value otherwise.

### GetGarmentImageIdOk

`func (o *NaomiHauteV1Input) GetGarmentImageIdOk() (*string, bool)`

GetGarmentImageIdOk returns a tuple with the GarmentImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarmentImageId

`func (o *NaomiHauteV1Input) SetGarmentImageId(v string)`

SetGarmentImageId sets GarmentImageId field to given value.


### GetPoseId

`func (o *NaomiHauteV1Input) GetPoseId() string`

GetPoseId returns the PoseId field if non-nil, zero value otherwise.

### GetPoseIdOk

`func (o *NaomiHauteV1Input) GetPoseIdOk() (*string, bool)`

GetPoseIdOk returns a tuple with the PoseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoseId

`func (o *NaomiHauteV1Input) SetPoseId(v string)`

SetPoseId sets PoseId field to given value.


### GetSeed

`func (o *NaomiHauteV1Input) GetSeed() float32`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *NaomiHauteV1Input) GetSeedOk() (*float32, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *NaomiHauteV1Input) SetSeed(v float32)`

SetSeed sets Seed field to given value.


### GetWidth

`func (o *NaomiHauteV1Input) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *NaomiHauteV1Input) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *NaomiHauteV1Input) SetWidth(v float32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *NaomiHauteV1Input) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *NaomiHauteV1Input) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *NaomiHauteV1Input) SetHeight(v float32)`

SetHeight sets Height field to given value.


### GetTextGuidanceScale

`func (o *NaomiHauteV1Input) GetTextGuidanceScale() float32`

GetTextGuidanceScale returns the TextGuidanceScale field if non-nil, zero value otherwise.

### GetTextGuidanceScaleOk

`func (o *NaomiHauteV1Input) GetTextGuidanceScaleOk() (*float32, bool)`

GetTextGuidanceScaleOk returns a tuple with the TextGuidanceScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextGuidanceScale

`func (o *NaomiHauteV1Input) SetTextGuidanceScale(v float32)`

SetTextGuidanceScale sets TextGuidanceScale field to given value.

### HasTextGuidanceScale

`func (o *NaomiHauteV1Input) HasTextGuidanceScale() bool`

HasTextGuidanceScale returns a boolean if a field has been set.

### GetImageGuidanceScale

`func (o *NaomiHauteV1Input) GetImageGuidanceScale() float32`

GetImageGuidanceScale returns the ImageGuidanceScale field if non-nil, zero value otherwise.

### GetImageGuidanceScaleOk

`func (o *NaomiHauteV1Input) GetImageGuidanceScaleOk() (*float32, bool)`

GetImageGuidanceScaleOk returns a tuple with the ImageGuidanceScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageGuidanceScale

`func (o *NaomiHauteV1Input) SetImageGuidanceScale(v float32)`

SetImageGuidanceScale sets ImageGuidanceScale field to given value.

### HasImageGuidanceScale

`func (o *NaomiHauteV1Input) HasImageGuidanceScale() bool`

HasImageGuidanceScale returns a boolean if a field has been set.

### GetNumInferenceSteps

`func (o *NaomiHauteV1Input) GetNumInferenceSteps() float32`

GetNumInferenceSteps returns the NumInferenceSteps field if non-nil, zero value otherwise.

### GetNumInferenceStepsOk

`func (o *NaomiHauteV1Input) GetNumInferenceStepsOk() (*float32, bool)`

GetNumInferenceStepsOk returns a tuple with the NumInferenceSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInferenceSteps

`func (o *NaomiHauteV1Input) SetNumInferenceSteps(v float32)`

SetNumInferenceSteps sets NumInferenceSteps field to given value.

### HasNumInferenceSteps

`func (o *NaomiHauteV1Input) HasNumInferenceSteps() bool`

HasNumInferenceSteps returns a boolean if a field has been set.

### GetMode

`func (o *NaomiHauteV1Input) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *NaomiHauteV1Input) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *NaomiHauteV1Input) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *NaomiHauteV1Input) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


