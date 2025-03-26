# ListOperationsParamsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderBy** | Pointer to **string** |  | [optional] [default to "createdAt_DESC"]
**Limit** | Pointer to **float32** |  | [optional] [default to 50]
**Cursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListOperationsParamsDto

`func NewListOperationsParamsDto() *ListOperationsParamsDto`

NewListOperationsParamsDto instantiates a new ListOperationsParamsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOperationsParamsDtoWithDefaults

`func NewListOperationsParamsDtoWithDefaults() *ListOperationsParamsDto`

NewListOperationsParamsDtoWithDefaults instantiates a new ListOperationsParamsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderBy

`func (o *ListOperationsParamsDto) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ListOperationsParamsDto) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ListOperationsParamsDto) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ListOperationsParamsDto) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetLimit

`func (o *ListOperationsParamsDto) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListOperationsParamsDto) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListOperationsParamsDto) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListOperationsParamsDto) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetCursor

`func (o *ListOperationsParamsDto) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListOperationsParamsDto) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListOperationsParamsDto) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListOperationsParamsDto) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


