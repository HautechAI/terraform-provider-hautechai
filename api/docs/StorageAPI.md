# \StorageAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageControllerCreateRecordV1**](StorageAPI.md#StorageControllerCreateRecordV1) | **Post** /v1/storage | 
[**StorageControllerDeleteRecordV1**](StorageAPI.md#StorageControllerDeleteRecordV1) | **Post** /v1/storage/delete | 
[**StorageControllerGetRecordsV1**](StorageAPI.md#StorageControllerGetRecordsV1) | **Post** /v1/storage/many | 
[**StorageControllerUpdateRecordV1**](StorageAPI.md#StorageControllerUpdateRecordV1) | **Post** /v1/storage/write | 



## StorageControllerCreateRecordV1

> StorageEntity StorageControllerCreateRecordV1(ctx).CreateStorageRecordParamsDto(createStorageRecordParamsDto).Execute()



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
	createStorageRecordParamsDto := *openapiclient.NewCreateStorageRecordParamsDto("Key_example", map[string]interface{}(123)) // CreateStorageRecordParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageControllerCreateRecordV1(context.Background()).CreateStorageRecordParamsDto(createStorageRecordParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageControllerCreateRecordV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageControllerCreateRecordV1`: StorageEntity
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageControllerCreateRecordV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorageControllerCreateRecordV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStorageRecordParamsDto** | [**CreateStorageRecordParamsDto**](CreateStorageRecordParamsDto.md) |  | 

### Return type

[**StorageEntity**](StorageEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageControllerDeleteRecordV1

> StorageControllerDeleteRecordV1(ctx).DeleteStorageParamsDto(deleteStorageParamsDto).Execute()



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
	deleteStorageParamsDto := *openapiclient.NewDeleteStorageParamsDto("Key_example") // DeleteStorageParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageAPI.StorageControllerDeleteRecordV1(context.Background()).DeleteStorageParamsDto(deleteStorageParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageControllerDeleteRecordV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorageControllerDeleteRecordV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteStorageParamsDto** | [**DeleteStorageParamsDto**](DeleteStorageParamsDto.md) |  | 

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


## StorageControllerGetRecordsV1

> []StorageRecordsResultDto StorageControllerGetRecordsV1(ctx).GetStorageRecordParamsDto(getStorageRecordParamsDto).Execute()



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
	getStorageRecordParamsDto := *openapiclient.NewGetStorageRecordParamsDto([]string{"Keys_example"}) // GetStorageRecordParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageControllerGetRecordsV1(context.Background()).GetStorageRecordParamsDto(getStorageRecordParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageControllerGetRecordsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageControllerGetRecordsV1`: []StorageRecordsResultDto
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageControllerGetRecordsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorageControllerGetRecordsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getStorageRecordParamsDto** | [**GetStorageRecordParamsDto**](GetStorageRecordParamsDto.md) |  | 

### Return type

[**[]StorageRecordsResultDto**](StorageRecordsResultDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageControllerUpdateRecordV1

> StorageEntity StorageControllerUpdateRecordV1(ctx).UpdateStorageRecordParamsDto(updateStorageRecordParamsDto).Execute()



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
	updateStorageRecordParamsDto := *openapiclient.NewUpdateStorageRecordParamsDto("Key_example", map[string]interface{}(123)) // UpdateStorageRecordParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageControllerUpdateRecordV1(context.Background()).UpdateStorageRecordParamsDto(updateStorageRecordParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageControllerUpdateRecordV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageControllerUpdateRecordV1`: StorageEntity
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageControllerUpdateRecordV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorageControllerUpdateRecordV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateStorageRecordParamsDto** | [**UpdateStorageRecordParamsDto**](UpdateStorageRecordParamsDto.md) |  | 

### Return type

[**StorageEntity**](StorageEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

