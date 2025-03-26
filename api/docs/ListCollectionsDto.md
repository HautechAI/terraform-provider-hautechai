# ListCollectionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CollectionEntity**](CollectionEntity.md) |  | 
**PageInfo** | [**ListAccountsDtoPageInfo**](ListAccountsDtoPageInfo.md) |  | 

## Methods

### NewListCollectionsDto

`func NewListCollectionsDto(data []CollectionEntity, pageInfo ListAccountsDtoPageInfo, ) *ListCollectionsDto`

NewListCollectionsDto instantiates a new ListCollectionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCollectionsDtoWithDefaults

`func NewListCollectionsDtoWithDefaults() *ListCollectionsDto`

NewListCollectionsDtoWithDefaults instantiates a new ListCollectionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCollectionsDto) GetData() []CollectionEntity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCollectionsDto) GetDataOk() (*[]CollectionEntity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCollectionsDto) SetData(v []CollectionEntity)`

SetData sets Data field to given value.


### GetPageInfo

`func (o *ListCollectionsDto) GetPageInfo() ListAccountsDtoPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *ListCollectionsDto) GetPageInfoOk() (*ListAccountsDtoPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *ListCollectionsDto) SetPageInfo(v ListAccountsDtoPageInfo)`

SetPageInfo sets PageInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


