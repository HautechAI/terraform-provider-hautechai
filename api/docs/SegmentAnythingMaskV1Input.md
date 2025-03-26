# SegmentAnythingMaskV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | **string** |  | 
**Box** | **[]float32** |  | 
**Smoothness** | Pointer to **float32** |  | [optional] [default to 0.1]
**MaskThreshold** | Pointer to **float32** |  | [optional] [default to 0]

## Methods

### NewSegmentAnythingMaskV1Input

`func NewSegmentAnythingMaskV1Input(imageId string, box []float32, ) *SegmentAnythingMaskV1Input`

NewSegmentAnythingMaskV1Input instantiates a new SegmentAnythingMaskV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentAnythingMaskV1InputWithDefaults

`func NewSegmentAnythingMaskV1InputWithDefaults() *SegmentAnythingMaskV1Input`

NewSegmentAnythingMaskV1InputWithDefaults instantiates a new SegmentAnythingMaskV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *SegmentAnythingMaskV1Input) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *SegmentAnythingMaskV1Input) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *SegmentAnythingMaskV1Input) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetBox

`func (o *SegmentAnythingMaskV1Input) GetBox() []float32`

GetBox returns the Box field if non-nil, zero value otherwise.

### GetBoxOk

`func (o *SegmentAnythingMaskV1Input) GetBoxOk() (*[]float32, bool)`

GetBoxOk returns a tuple with the Box field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBox

`func (o *SegmentAnythingMaskV1Input) SetBox(v []float32)`

SetBox sets Box field to given value.


### GetSmoothness

`func (o *SegmentAnythingMaskV1Input) GetSmoothness() float32`

GetSmoothness returns the Smoothness field if non-nil, zero value otherwise.

### GetSmoothnessOk

`func (o *SegmentAnythingMaskV1Input) GetSmoothnessOk() (*float32, bool)`

GetSmoothnessOk returns a tuple with the Smoothness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmoothness

`func (o *SegmentAnythingMaskV1Input) SetSmoothness(v float32)`

SetSmoothness sets Smoothness field to given value.

### HasSmoothness

`func (o *SegmentAnythingMaskV1Input) HasSmoothness() bool`

HasSmoothness returns a boolean if a field has been set.

### GetMaskThreshold

`func (o *SegmentAnythingMaskV1Input) GetMaskThreshold() float32`

GetMaskThreshold returns the MaskThreshold field if non-nil, zero value otherwise.

### GetMaskThresholdOk

`func (o *SegmentAnythingMaskV1Input) GetMaskThresholdOk() (*float32, bool)`

GetMaskThresholdOk returns a tuple with the MaskThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskThreshold

`func (o *SegmentAnythingMaskV1Input) SetMaskThreshold(v float32)`

SetMaskThreshold sets MaskThreshold field to given value.

### HasMaskThreshold

`func (o *SegmentAnythingMaskV1Input) HasMaskThreshold() bool`

HasMaskThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


