# \PipelinesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PipelinesControllerCreatePipelineV1**](PipelinesAPI.md#PipelinesControllerCreatePipelineV1) | **Post** /v1/pipelines | 
[**PipelinesControllerGetPipelineV1**](PipelinesAPI.md#PipelinesControllerGetPipelineV1) | **Get** /v1/pipelines/{id} | 
[**PipelinesControllerListPipelinesV1**](PipelinesAPI.md#PipelinesControllerListPipelinesV1) | **Get** /v1/pipelines | 



## PipelinesControllerCreatePipelineV1

> PipelineEntity PipelinesControllerCreatePipelineV1(ctx).CreatePipelineParamsDto(createPipelineParamsDto).Execute()



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
	createPipelineParamsDto := *openapiclient.NewCreatePipelineParamsDto(map[string]interface{}(123)) // CreatePipelineParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelinesAPI.PipelinesControllerCreatePipelineV1(context.Background()).CreatePipelineParamsDto(createPipelineParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.PipelinesControllerCreatePipelineV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PipelinesControllerCreatePipelineV1`: PipelineEntity
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.PipelinesControllerCreatePipelineV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPipelinesControllerCreatePipelineV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPipelineParamsDto** | [**CreatePipelineParamsDto**](CreatePipelineParamsDto.md) |  | 

### Return type

[**PipelineEntity**](PipelineEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PipelinesControllerGetPipelineV1

> PipelineEntity PipelinesControllerGetPipelineV1(ctx, id).Execute()



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
	resp, r, err := apiClient.PipelinesAPI.PipelinesControllerGetPipelineV1(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.PipelinesControllerGetPipelineV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PipelinesControllerGetPipelineV1`: PipelineEntity
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.PipelinesControllerGetPipelineV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPipelinesControllerGetPipelineV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PipelineEntity**](PipelineEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PipelinesControllerListPipelinesV1

> ListPipelinesDto PipelinesControllerListPipelinesV1(ctx).OrderBy(orderBy).Limit(limit).Cursor(cursor).Execute()



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
	resp, r, err := apiClient.PipelinesAPI.PipelinesControllerListPipelinesV1(context.Background()).OrderBy(orderBy).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.PipelinesControllerListPipelinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PipelinesControllerListPipelinesV1`: ListPipelinesDto
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.PipelinesControllerListPipelinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPipelinesControllerListPipelinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **string** |  | [default to &quot;createdAt_DESC&quot;]
 **limit** | **float32** |  | [default to 50]
 **cursor** | **string** |  | 

### Return type

[**ListPipelinesDto**](ListPipelinesDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

