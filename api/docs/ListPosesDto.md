# ListPosesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PoseEntity**](PoseEntity.md) |  | 
**PageInfo** | [**ListAccountsDtoPageInfo**](ListAccountsDtoPageInfo.md) |  | 

## Methods

### NewListPosesDto

`func NewListPosesDto(data []PoseEntity, pageInfo ListAccountsDtoPageInfo, ) *ListPosesDto`

NewListPosesDto instantiates a new ListPosesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPosesDtoWithDefaults

`func NewListPosesDtoWithDefaults() *ListPosesDto`

NewListPosesDtoWithDefaults instantiates a new ListPosesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListPosesDto) GetData() []PoseEntity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPosesDto) GetDataOk() (*[]PoseEntity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPosesDto) SetData(v []PoseEntity)`

SetData sets Data field to given value.


### GetPageInfo

`func (o *ListPosesDto) GetPageInfo() ListAccountsDtoPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *ListPosesDto) GetPageInfoOk() (*ListAccountsDtoPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *ListPosesDto) SetPageInfo(v ListAccountsDtoPageInfo)`

SetPageInfo sets PageInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


