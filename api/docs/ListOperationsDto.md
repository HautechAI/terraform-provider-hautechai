# ListOperationsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]OperationEntity**](OperationEntity.md) |  | 
**PageInfo** | [**ListAccountsDtoPageInfo**](ListAccountsDtoPageInfo.md) |  | 

## Methods

### NewListOperationsDto

`func NewListOperationsDto(data []OperationEntity, pageInfo ListAccountsDtoPageInfo, ) *ListOperationsDto`

NewListOperationsDto instantiates a new ListOperationsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOperationsDtoWithDefaults

`func NewListOperationsDtoWithDefaults() *ListOperationsDto`

NewListOperationsDtoWithDefaults instantiates a new ListOperationsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListOperationsDto) GetData() []OperationEntity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListOperationsDto) GetDataOk() (*[]OperationEntity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListOperationsDto) SetData(v []OperationEntity)`

SetData sets Data field to given value.


### GetPageInfo

`func (o *ListOperationsDto) GetPageInfo() ListAccountsDtoPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *ListOperationsDto) GetPageInfoOk() (*ListAccountsDtoPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *ListOperationsDto) SetPageInfo(v ListAccountsDtoPageInfo)`

SetPageInfo sets PageInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


