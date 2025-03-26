# ListAccountsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AccountEntity**](AccountEntity.md) |  | 
**PageInfo** | [**ListAccountsDtoPageInfo**](ListAccountsDtoPageInfo.md) |  | 

## Methods

### NewListAccountsDto

`func NewListAccountsDto(data []AccountEntity, pageInfo ListAccountsDtoPageInfo, ) *ListAccountsDto`

NewListAccountsDto instantiates a new ListAccountsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccountsDtoWithDefaults

`func NewListAccountsDtoWithDefaults() *ListAccountsDto`

NewListAccountsDtoWithDefaults instantiates a new ListAccountsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAccountsDto) GetData() []AccountEntity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAccountsDto) GetDataOk() (*[]AccountEntity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAccountsDto) SetData(v []AccountEntity)`

SetData sets Data field to given value.


### GetPageInfo

`func (o *ListAccountsDto) GetPageInfo() ListAccountsDtoPageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *ListAccountsDto) GetPageInfoOk() (*ListAccountsDtoPageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *ListAccountsDto) SetPageInfo(v ListAccountsDtoPageInfo)`

SetPageInfo sets PageInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


