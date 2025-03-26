# ListStacksDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]StackEntity**](StackEntity.md) |  | 
**PageInfo** | [**ListAccountsDtoPageInfo**](ListAccountsDtoPageInfo.md) |  | 

## Methods

### NewListStacksDto

`func NewListStacksDto(data []StackEntity, pageInfo ListAccountsDtoPageInfo, ) *ListStacksDto`

NewListStacksDto instantiates a new ListStacksDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStacksDtoWithDefaults

`func NewListStacksDtoWithDefaults() *ListStacksDto`

NewListStacksDtoWithDefaults instantiates a new ListStacksDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListStacksDto) GetData() []StackEntity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListStacksDto) GetDataOk() (*[]StackEntity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListStacksDto) SetData(v []StackEntity)`

SetData sets Data field to given value.


### GetPageInfo

`func (o *ListStacksDto) GetPageInfo() ListAccountsDtoPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *ListStacksDto) GetPageInfoOk() (*ListAccountsDtoPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *ListStacksDto) SetPageInfo(v ListAccountsDtoPageInfo)`

SetPageInfo sets PageInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


