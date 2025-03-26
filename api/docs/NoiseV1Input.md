# NoiseV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | **string** |  | 
**Mean** | Pointer to **float32** |  | [optional] 
**Sigma** | Pointer to **float32** |  | [optional] 

## Methods

### NewNoiseV1Input

`func NewNoiseV1Input(imageId string, ) *NoiseV1Input`

NewNoiseV1Input instantiates a new NoiseV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoiseV1InputWithDefaults

`func NewNoiseV1InputWithDefaults() *NoiseV1Input`

NewNoiseV1InputWithDefaults instantiates a new NoiseV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *NoiseV1Input) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *NoiseV1Input) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *NoiseV1Input) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetMean

`func (o *NoiseV1Input) GetMean() float32`

GetMean returns the Mean field if non-nil, zero value otherwise.

### GetMeanOk

`func (o *NoiseV1Input) GetMeanOk() (*float32, bool)`

GetMeanOk returns a tuple with the Mean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMean

`func (o *NoiseV1Input) SetMean(v float32)`

SetMean sets Mean field to given value.

### HasMean

`func (o *NoiseV1Input) HasMean() bool`

HasMean returns a boolean if a field has been set.

### GetSigma

`func (o *NoiseV1Input) GetSigma() float32`

GetSigma returns the Sigma field if non-nil, zero value otherwise.

### GetSigmaOk

`func (o *NoiseV1Input) GetSigmaOk() (*float32, bool)`

GetSigmaOk returns a tuple with the Sigma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigma

`func (o *NoiseV1Input) SetSigma(v float32)`

SetSigma sets Sigma field to given value.

### HasSigma

`func (o *NoiseV1Input) HasSigma() bool`

HasSigma returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


