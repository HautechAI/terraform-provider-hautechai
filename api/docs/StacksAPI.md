# \StacksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StacksControllerAddItemsV1**](StacksAPI.md#StacksControllerAddItemsV1) | **Post** /v1/stacks/{id}/items/add | 
[**StacksControllerCreateStackV1**](StacksAPI.md#StacksControllerCreateStackV1) | **Post** /v1/stacks | 
[**StacksControllerGetStackV1**](StacksAPI.md#StacksControllerGetStackV1) | **Get** /v1/stacks/{id} | 
[**StacksControllerListStacksV1**](StacksAPI.md#StacksControllerListStacksV1) | **Get** /v1/stacks | 
[**StacksControllerRemoveItemsV1**](StacksAPI.md#StacksControllerRemoveItemsV1) | **Post** /v1/stacks/{id}/items/remove | 
[**StacksControllerUpdateMetadataV1**](StacksAPI.md#StacksControllerUpdateMetadataV1) | **Put** /v1/stacks/{id}/metadata | 



## StacksControllerAddItemsV1

> StackEntity StacksControllerAddItemsV1(ctx, id).AddItemsToStackControllerParamsDto(addItemsToStackControllerParamsDto).Execute()



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
	addItemsToStackControllerParamsDto := *openapiclient.NewAddItemsToStackControllerParamsDto([]string{"ItemIds_example"}) // AddItemsToStackControllerParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.StacksControllerAddItemsV1(context.Background(), id).AddItemsToStackControllerParamsDto(addItemsToStackControllerParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.StacksControllerAddItemsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StacksControllerAddItemsV1`: StackEntity
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.StacksControllerAddItemsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStacksControllerAddItemsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addItemsToStackControllerParamsDto** | [**AddItemsToStackControllerParamsDto**](AddItemsToStackControllerParamsDto.md) |  | 

### Return type

[**StackEntity**](StackEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StacksControllerCreateStackV1

> StackEntity StacksControllerCreateStackV1(ctx).CreateStackParamsDto(createStackParamsDto).Execute()



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
	createStackParamsDto := *openapiclient.NewCreateStackParamsDto() // CreateStackParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.StacksControllerCreateStackV1(context.Background()).CreateStackParamsDto(createStackParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.StacksControllerCreateStackV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StacksControllerCreateStackV1`: StackEntity
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.StacksControllerCreateStackV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStacksControllerCreateStackV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStackParamsDto** | [**CreateStackParamsDto**](CreateStackParamsDto.md) |  | 

### Return type

[**StackEntity**](StackEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StacksControllerGetStackV1

> StackEntity StacksControllerGetStackV1(ctx, id).Execute()



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
	resp, r, err := apiClient.StacksAPI.StacksControllerGetStackV1(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.StacksControllerGetStackV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StacksControllerGetStackV1`: StackEntity
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.StacksControllerGetStackV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStacksControllerGetStackV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StackEntity**](StackEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StacksControllerListStacksV1

> ListStacksDto StacksControllerListStacksV1(ctx).OrderBy(orderBy).Limit(limit).Cursor(cursor).Execute()



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
	resp, r, err := apiClient.StacksAPI.StacksControllerListStacksV1(context.Background()).OrderBy(orderBy).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.StacksControllerListStacksV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StacksControllerListStacksV1`: ListStacksDto
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.StacksControllerListStacksV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStacksControllerListStacksV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **string** |  | [default to &quot;createdAt_DESC&quot;]
 **limit** | **float32** |  | [default to 50]
 **cursor** | **string** |  | 

### Return type

[**ListStacksDto**](ListStacksDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StacksControllerRemoveItemsV1

> StackEntity StacksControllerRemoveItemsV1(ctx, id).RemoveItemsFromCollectionControllerParamsDto(removeItemsFromCollectionControllerParamsDto).Execute()



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
	removeItemsFromCollectionControllerParamsDto := *openapiclient.NewRemoveItemsFromCollectionControllerParamsDto([]string{"ItemIds_example"}) // RemoveItemsFromCollectionControllerParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.StacksControllerRemoveItemsV1(context.Background(), id).RemoveItemsFromCollectionControllerParamsDto(removeItemsFromCollectionControllerParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.StacksControllerRemoveItemsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StacksControllerRemoveItemsV1`: StackEntity
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.StacksControllerRemoveItemsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStacksControllerRemoveItemsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeItemsFromCollectionControllerParamsDto** | [**RemoveItemsFromCollectionControllerParamsDto**](RemoveItemsFromCollectionControllerParamsDto.md) |  | 

### Return type

[**StackEntity**](StackEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StacksControllerUpdateMetadataV1

> ResourceEntity StacksControllerUpdateMetadataV1(ctx, id).UpdateMetadataDto(updateMetadataDto).Execute()



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
	resp, r, err := apiClient.StacksAPI.StacksControllerUpdateMetadataV1(context.Background(), id).UpdateMetadataDto(updateMetadataDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.StacksControllerUpdateMetadataV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StacksControllerUpdateMetadataV1`: ResourceEntity
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.StacksControllerUpdateMetadataV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStacksControllerUpdateMetadataV1Request struct via the builder pattern


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

