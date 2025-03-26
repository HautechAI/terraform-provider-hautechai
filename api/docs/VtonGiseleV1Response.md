# VtonGiseleV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | [default to "operation"]
**Output** | [**NullableOperationOutputImageSingle**](OperationOutputImageSingle.md) |  | 
**Input** | **map[string]interface{}** |  | 
**Status** | **string** |  | 
**Type** | **string** |  | 
**Id** | **string** |  | 
**CreatorId** | **string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewVtonGiseleV1Response

`func NewVtonGiseleV1Response(kind string, output NullableOperationOutputImageSingle, input map[string]interface{}, status string, type_ string, id string, creatorId string, metadata map[string]interface{}, createdAt time.Time, updatedAt time.Time, ) *VtonGiseleV1Response`

NewVtonGiseleV1Response instantiates a new VtonGiseleV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVtonGiseleV1ResponseWithDefaults

`func NewVtonGiseleV1ResponseWithDefaults() *VtonGiseleV1Response`

NewVtonGiseleV1ResponseWithDefaults instantiates a new VtonGiseleV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *VtonGiseleV1Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *VtonGiseleV1Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *VtonGiseleV1Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetOutput

`func (o *VtonGiseleV1Response) GetOutput() OperationOutputImageSingle`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *VtonGiseleV1Response) GetOutputOk() (*OperationOutputImageSingle, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *VtonGiseleV1Response) SetOutput(v OperationOutputImageSingle)`

SetOutput sets Output field to given value.


### SetOutputNil

`func (o *VtonGiseleV1Response) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *VtonGiseleV1Response) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetInput

`func (o *VtonGiseleV1Response) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *VtonGiseleV1Response) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *VtonGiseleV1Response) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.


### GetStatus

`func (o *VtonGiseleV1Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VtonGiseleV1Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VtonGiseleV1Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *VtonGiseleV1Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VtonGiseleV1Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VtonGiseleV1Response) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *VtonGiseleV1Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VtonGiseleV1Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VtonGiseleV1Response) SetId(v string)`

SetId sets Id field to given value.


### GetCreatorId

`func (o *VtonGiseleV1Response) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *VtonGiseleV1Response) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *VtonGiseleV1Response) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.


### GetMetadata

`func (o *VtonGiseleV1Response) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VtonGiseleV1Response) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VtonGiseleV1Response) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetCreatedAt

`func (o *VtonGiseleV1Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VtonGiseleV1Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VtonGiseleV1Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *VtonGiseleV1Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VtonGiseleV1Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VtonGiseleV1Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


