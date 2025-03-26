# \PosesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PosesControllerGetPoseV1**](PosesAPI.md#PosesControllerGetPoseV1) | **Get** /v1/poses/{id} | 
[**PosesControllerListPosesV1**](PosesAPI.md#PosesControllerListPosesV1) | **Get** /v1/poses | 
[**PosesControllerSetPosePreviewV1**](PosesAPI.md#PosesControllerSetPosePreviewV1) | **Put** /v1/poses/{id}/preview | 
[**PosesControllerUpdateMetadataV1**](PosesAPI.md#PosesControllerUpdateMetadataV1) | **Put** /v1/poses/{id}/metadata | 



## PosesControllerGetPoseV1

> PoseEntity PosesControllerGetPoseV1(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/hautechapi"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PosesAPI.PosesControllerGetPoseV1(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PosesAPI.PosesControllerGetPoseV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PosesControllerGetPoseV1`: PoseEntity
	fmt.Fprintf(os.Stdout, "Response from `PosesAPI.PosesControllerGetPoseV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPosesControllerGetPoseV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PoseEntity**](PoseEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PosesControllerListPosesV1

> ListPosesDto PosesControllerListPosesV1(ctx).OrderBy(orderBy).Limit(limit).Cursor(cursor).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/hautechapi"
)

func main() {
	orderBy := "orderBy_example" // string |  (optional) (default to "createdAt_DESC")
	limit := float32(8.14) // float32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PosesAPI.PosesControllerListPosesV1(context.Background()).OrderBy(orderBy).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PosesAPI.PosesControllerListPosesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PosesControllerListPosesV1`: ListPosesDto
	fmt.Fprintf(os.Stdout, "Response from `PosesAPI.PosesControllerListPosesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPosesControllerListPosesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **string** |  | [default to &quot;createdAt_DESC&quot;]
 **limit** | **float32** |  | [default to 50]
 **cursor** | **string** |  | 

### Return type

[**ListPosesDto**](ListPosesDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PosesControllerSetPosePreviewV1

> PosesControllerSetPosePreviewV1(ctx, id).SetPosePreviewControllerParamsDto(setPosePreviewControllerParamsDto).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/hautechapi"
)

func main() {
	id := "id_example" // string | 
	setPosePreviewControllerParamsDto := *openapiclient.NewSetPosePreviewControllerParamsDto("PreviewImageId_example") // SetPosePreviewControllerParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PosesAPI.PosesControllerSetPosePreviewV1(context.Background(), id).SetPosePreviewControllerParamsDto(setPosePreviewControllerParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PosesAPI.PosesControllerSetPosePreviewV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPosesControllerSetPosePreviewV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setPosePreviewControllerParamsDto** | [**SetPosePreviewControllerParamsDto**](SetPosePreviewControllerParamsDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PosesControllerUpdateMetadataV1

> ResourceEntity PosesControllerUpdateMetadataV1(ctx, id).UpdateMetadataDto(updateMetadataDto).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/hautechapi"
)

func main() {
	id := "id_example" // string | 
	updateMetadataDto := *openapiclient.NewUpdateMetadataDto(map[string]interface{}(123)) // UpdateMetadataDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PosesAPI.PosesControllerUpdateMetadataV1(context.Background(), id).UpdateMetadataDto(updateMetadataDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PosesAPI.PosesControllerUpdateMetadataV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PosesControllerUpdateMetadataV1`: ResourceEntity
	fmt.Fprintf(os.Stdout, "Response from `PosesAPI.PosesControllerUpdateMetadataV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPosesControllerUpdateMetadataV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMetadataDto** | [**UpdateMetadataDto**](UpdateMetadataDto.md) |  | 

### Return type

[**ResourceEntity**](ResourceEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

