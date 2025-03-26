# AddItemsToStackParamsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | **string** |  | 
**ItemIds** | **[]string** |  | 

## Methods

### NewAddItemsToStackParamsDto

`func NewAddItemsToStackParamsDto(stackId string, itemIds []string, ) *AddItemsToStackParamsDto`

NewAddItemsToStackParamsDto instantiates a new AddItemsToStackParamsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddItemsToStackParamsDtoWithDefaults

`func NewAddItemsToStackParamsDtoWithDefaults() *AddItemsToStackParamsDto`

NewAddItemsToStackParamsDtoWithDefaults instantiates a new AddItemsToStackParamsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackId

`func (o *AddItemsToStackParamsDto) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *AddItemsToStackParamsDto) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *AddItemsToStackParamsDto) SetStackId(v string)`

SetStackId sets StackId field to given value.


### GetItemIds

`func (o *AddItemsToStackParamsDto) GetItemIds() []string`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *AddItemsToStackParamsDto) GetItemIdsOk() (*[]string, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *AddItemsToStackParamsDto) SetItemIds(v []string)`

SetItemIds sets ItemIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


