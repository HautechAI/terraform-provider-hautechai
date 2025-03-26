# \ImagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImagesControllerFinalizeUploadV1**](ImagesAPI.md#ImagesControllerFinalizeUploadV1) | **Post** /v1/images/upload/finalize | 
[**ImagesControllerGetImageV1**](ImagesAPI.md#ImagesControllerGetImageV1) | **Get** /v1/images/{id} | 
[**ImagesControllerGetRepresentationV1**](ImagesAPI.md#ImagesControllerGetRepresentationV1) | **Get** /v1/images/{id}/representation/{type} | 
[**ImagesControllerGetUrlsV1**](ImagesAPI.md#ImagesControllerGetUrlsV1) | **Post** /v1/images/many | 
[**ImagesControllerStartUploadV1**](ImagesAPI.md#ImagesControllerStartUploadV1) | **Post** /v1/images/upload/initialize | 



## ImagesControllerFinalizeUploadV1

> ImageEntity ImagesControllerFinalizeUploadV1(ctx).CreateImageParamsDto(createImageParamsDto).Execute()



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
	createImageParamsDto := *openapiclient.NewCreateImageParamsDto() // CreateImageParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.ImagesControllerFinalizeUploadV1(context.Background()).CreateImageParamsDto(createImageParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesControllerFinalizeUploadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImagesControllerFinalizeUploadV1`: ImageEntity
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesControllerFinalizeUploadV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagesControllerFinalizeUploadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createImageParamsDto** | [**CreateImageParamsDto**](CreateImageParamsDto.md) |  | 

### Return type

[**ImageEntity**](ImageEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesControllerGetImageV1

> ImageEntity ImagesControllerGetImageV1(ctx, id).Execute()



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
	resp, r, err := apiClient.ImagesAPI.ImagesControllerGetImageV1(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesControllerGetImageV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImagesControllerGetImageV1`: ImageEntity
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesControllerGetImageV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesControllerGetImageV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImageEntity**](ImageEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesControllerGetRepresentationV1

> ImageRepresentationResponseDto ImagesControllerGetRepresentationV1(ctx, id, type_).Execute()



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
	type_ := "type__example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.ImagesControllerGetRepresentationV1(context.Background(), id, type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesControllerGetRepresentationV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImagesControllerGetRepresentationV1`: ImageRepresentationResponseDto
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesControllerGetRepresentationV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**type_** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImagesControllerGetRepresentationV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ImageRepresentationResponseDto**](ImageRepresentationResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesControllerGetUrlsV1

> []ImageEntity ImagesControllerGetUrlsV1(ctx).GetUrlsForImagesParamsDto(getUrlsForImagesParamsDto).Execute()



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
	getUrlsForImagesParamsDto := *openapiclient.NewGetUrlsForImagesParamsDto([]string{"Ids_example"}) // GetUrlsForImagesParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.ImagesControllerGetUrlsV1(context.Background()).GetUrlsForImagesParamsDto(getUrlsForImagesParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesControllerGetUrlsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImagesControllerGetUrlsV1`: []ImageEntity
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesControllerGetUrlsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagesControllerGetUrlsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getUrlsForImagesParamsDto** | [**GetUrlsForImagesParamsDto**](GetUrlsForImagesParamsDto.md) |  | 

### Return type

[**[]ImageEntity**](ImageEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagesControllerStartUploadV1

> InitializeImageUploadResultDto ImagesControllerStartUploadV1(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.ImagesControllerStartUploadV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesControllerStartUploadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImagesControllerStartUploadV1`: InitializeImageUploadResultDto
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesControllerStartUploadV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiImagesControllerStartUploadV1Request struct via the builder pattern


### Return type

[**InitializeImageUploadResultDto**](InitializeImageUploadResultDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

