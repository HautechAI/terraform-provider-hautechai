# \CollectionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CollectionsControllerAddItemsV1**](CollectionsAPI.md#CollectionsControllerAddItemsV1) | **Post** /v1/collections/{id}/items/add | 
[**CollectionsControllerCreateCollectionV1**](CollectionsAPI.md#CollectionsControllerCreateCollectionV1) | **Post** /v1/collections | 
[**CollectionsControllerGetCollectionV1**](CollectionsAPI.md#CollectionsControllerGetCollectionV1) | **Get** /v1/collections/{id} | 
[**CollectionsControllerListCollectionsV1**](CollectionsAPI.md#CollectionsControllerListCollectionsV1) | **Get** /v1/collections | 
[**CollectionsControllerListItemsV1**](CollectionsAPI.md#CollectionsControllerListItemsV1) | **Get** /v1/collections/{id}/items | 
[**CollectionsControllerRemoveItemsV1**](CollectionsAPI.md#CollectionsControllerRemoveItemsV1) | **Post** /v1/collections/{id}/items/remove | 
[**CollectionsControllerUpdateMetadataV1**](CollectionsAPI.md#CollectionsControllerUpdateMetadataV1) | **Put** /v1/collections/{id}/metadata | 



## CollectionsControllerAddItemsV1

> CollectionsControllerAddItemsV1(ctx, id).AddItemsToCollectionControllerParamsDto(addItemsToCollectionControllerParamsDto).Execute()



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
	addItemsToCollectionControllerParamsDto := *openapiclient.NewAddItemsToCollectionControllerParamsDto([]string{"ItemIds_example"}) // AddItemsToCollectionControllerParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsAPI.CollectionsControllerAddItemsV1(context.Background(), id).AddItemsToCollectionControllerParamsDto(addItemsToCollectionControllerParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionsControllerAddItemsV1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCollectionsControllerAddItemsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addItemsToCollectionControllerParamsDto** | [**AddItemsToCollectionControllerParamsDto**](AddItemsToCollectionControllerParamsDto.md) |  | 

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


## CollectionsControllerCreateCollectionV1

> CollectionEntity CollectionsControllerCreateCollectionV1(ctx).CreateCollectionParamsDto(createCollectionParamsDto).Execute()



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
	createCollectionParamsDto := *openapiclient.NewCreateCollectionParamsDto() // CreateCollectionParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.CollectionsControllerCreateCollectionV1(context.Background()).CreateCollectionParamsDto(createCollectionParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionsControllerCreateCollectionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CollectionsControllerCreateCollectionV1`: CollectionEntity
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.CollectionsControllerCreateCollectionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsControllerCreateCollectionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCollectionParamsDto** | [**CreateCollectionParamsDto**](CreateCollectionParamsDto.md) |  | 

### Return type

[**CollectionEntity**](CollectionEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectionsControllerGetCollectionV1

> CollectionEntity CollectionsControllerGetCollectionV1(ctx, id).Execute()



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
	resp, r, err := apiClient.CollectionsAPI.CollectionsControllerGetCollectionV1(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionsControllerGetCollectionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CollectionsControllerGetCollectionV1`: CollectionEntity
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.CollectionsControllerGetCollectionV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsControllerGetCollectionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionEntity**](CollectionEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectionsControllerListCollectionsV1

> ListCollectionsDto CollectionsControllerListCollectionsV1(ctx).OrderBy(orderBy).Limit(limit).Cursor(cursor).Execute()



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
	resp, r, err := apiClient.CollectionsAPI.CollectionsControllerListCollectionsV1(context.Background()).OrderBy(orderBy).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionsControllerListCollectionsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CollectionsControllerListCollectionsV1`: ListCollectionsDto
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.CollectionsControllerListCollectionsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsControllerListCollectionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **string** |  | [default to &quot;createdAt_DESC&quot;]
 **limit** | **float32** |  | [default to 50]
 **cursor** | **string** |  | 

### Return type

[**ListCollectionsDto**](ListCollectionsDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectionsControllerListItemsV1

> ListCollectionItemsDto CollectionsControllerListItemsV1(ctx, id).Cursor(cursor).OrderBy(orderBy).Limit(limit).Kind(kind).Execute()



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
	cursor := "cursor_example" // string |  (optional)
	orderBy := "orderBy_example" // string |  (optional) (default to "createdAt_DESC")
	limit := float32(8.14) // float32 |  (optional) (default to 50)
	kind := "kind_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.CollectionsControllerListItemsV1(context.Background(), id).Cursor(cursor).OrderBy(orderBy).Limit(limit).Kind(kind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionsControllerListItemsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CollectionsControllerListItemsV1`: ListCollectionItemsDto
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.CollectionsControllerListItemsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsControllerListItemsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** |  | 
 **orderBy** | **string** |  | [default to &quot;createdAt_DESC&quot;]
 **limit** | **float32** |  | [default to 50]
 **kind** | **string** |  | 

### Return type

[**ListCollectionItemsDto**](ListCollectionItemsDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectionsControllerRemoveItemsV1

> CollectionsControllerRemoveItemsV1(ctx, id).RemoveItemsFromCollectionControllerParamsDto(removeItemsFromCollectionControllerParamsDto).Execute()



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
	r, err := apiClient.CollectionsAPI.CollectionsControllerRemoveItemsV1(context.Background(), id).RemoveItemsFromCollectionControllerParamsDto(removeItemsFromCollectionControllerParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionsControllerRemoveItemsV1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCollectionsControllerRemoveItemsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeItemsFromCollectionControllerParamsDto** | [**RemoveItemsFromCollectionControllerParamsDto**](RemoveItemsFromCollectionControllerParamsDto.md) |  | 

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


## CollectionsControllerUpdateMetadataV1

> ResourceEntity CollectionsControllerUpdateMetadataV1(ctx, id).UpdateMetadataDto(updateMetadataDto).Execute()



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
	resp, r, err := apiClient.CollectionsAPI.CollectionsControllerUpdateMetadataV1(context.Background(), id).UpdateMetadataDto(updateMetadataDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionsControllerUpdateMetadataV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CollectionsControllerUpdateMetadataV1`: ResourceEntity
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.CollectionsControllerUpdateMetadataV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsControllerUpdateMetadataV1Request struct via the builder pattern


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

