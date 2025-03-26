# ListCollectionItemsParamsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | **string** |  | 
**OrderBy** | Pointer to **string** |  | [optional] [default to "createdAt_DESC"]
**Limit** | Pointer to **float32** |  | [optional] [default to 50]
**Kind** | Pointer to **string** |  | [optional] 
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListCollectionItemsParamsDto

`func NewListCollectionItemsParamsDto(collectionId string, ) *ListCollectionItemsParamsDto`

NewListCollectionItemsParamsDto instantiates a new ListCollectionItemsParamsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCollectionItemsParamsDtoWithDefaults

`func NewListCollectionItemsParamsDtoWithDefaults() *ListCollectionItemsParamsDto`

NewListCollectionItemsParamsDtoWithDefaults instantiates a new ListCollectionItemsParamsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *ListCollectionItemsParamsDto) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *ListCollectionItemsParamsDto) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *ListCollectionItemsParamsDto) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.


### GetOrderBy

`func (o *ListCollectionItemsParamsDto) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ListCollectionItemsParamsDto) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ListCollectionItemsParamsDto) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ListCollectionItemsParamsDto) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetLimit

`func (o *ListCollectionItemsParamsDto) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListCollectionItemsParamsDto) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListCollectionItemsParamsDto) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListCollectionItemsParamsDto) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetKind

`func (o *ListCollectionItemsParamsDto) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListCollectionItemsParamsDto) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListCollectionItemsParamsDto) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ListCollectionItemsParamsDto) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCursor

`func (o *ListCollectionItemsParamsDto) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListCollectionItemsParamsDto) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListCollectionItemsParamsDto) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListCollectionItemsParamsDto) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


