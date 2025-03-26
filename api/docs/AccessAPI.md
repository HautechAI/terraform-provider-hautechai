# \AccessAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessControllerAccessV1**](AccessAPI.md#AccessControllerAccessV1) | **Get** /v1/resources/{id}/access | UNSTABLE
[**AccessControllerAttachAccessV1**](AccessAPI.md#AccessControllerAttachAccessV1) | **Post** /v1/resources/{id}/access/attach | 
[**AccessControllerDetachAccessV1**](AccessAPI.md#AccessControllerDetachAccessV1) | **Post** /v1/resources/{id}/access/detach | 
[**AccessControllerGrantAccessV1**](AccessAPI.md#AccessControllerGrantAccessV1) | **Post** /v1/resources/{id}/access/grant | 
[**AccessControllerRevokeAccessV1**](AccessAPI.md#AccessControllerRevokeAccessV1) | **Post** /v1/resources/{id}/access/revoke | 



## AccessControllerAccessV1

> ListAccessControllerDto AccessControllerAccessV1(ctx, id).Execute()

UNSTABLE

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
	resp, r, err := apiClient.AccessAPI.AccessControllerAccessV1(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessAPI.AccessControllerAccessV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessControllerAccessV1`: ListAccessControllerDto
	fmt.Fprintf(os.Stdout, "Response from `AccessAPI.AccessControllerAccessV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessControllerAccessV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListAccessControllerDto**](ListAccessControllerDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessControllerAttachAccessV1

> AccessControllerAttachAccessV1(ctx, id).AttachAccessControllerParamsDto(attachAccessControllerParamsDto).Execute()



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
	attachAccessControllerParamsDto := *openapiclient.NewAttachAccessControllerParamsDto("ParentResourceId_example") // AttachAccessControllerParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessAPI.AccessControllerAttachAccessV1(context.Background(), id).AttachAccessControllerParamsDto(attachAccessControllerParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessAPI.AccessControllerAttachAccessV1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAccessControllerAttachAccessV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachAccessControllerParamsDto** | [**AttachAccessControllerParamsDto**](AttachAccessControllerParamsDto.md) |  | 

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


## AccessControllerDetachAccessV1

> AccessControllerDetachAccessV1(ctx, id).DetachAccessControllerParamsDto(detachAccessControllerParamsDto).Execute()



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
	detachAccessControllerParamsDto := *openapiclient.NewDetachAccessControllerParamsDto("ParentResourceId_example") // DetachAccessControllerParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessAPI.AccessControllerDetachAccessV1(context.Background(), id).DetachAccessControllerParamsDto(detachAccessControllerParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessAPI.AccessControllerDetachAccessV1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAccessControllerDetachAccessV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detachAccessControllerParamsDto** | [**DetachAccessControllerParamsDto**](DetachAccessControllerParamsDto.md) |  | 

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


## AccessControllerGrantAccessV1

> AccessControllerGrantAccessV1(ctx, id).GrantAccessControllerParams(grantAccessControllerParams).Execute()



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
	grantAccessControllerParams := *openapiclient.NewGrantAccessControllerParams("PrincipalType_example", "PrincipalId_example", "Access_example") // GrantAccessControllerParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessAPI.AccessControllerGrantAccessV1(context.Background(), id).GrantAccessControllerParams(grantAccessControllerParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessAPI.AccessControllerGrantAccessV1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAccessControllerGrantAccessV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **grantAccessControllerParams** | [**GrantAccessControllerParams**](GrantAccessControllerParams.md) |  | 

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


## AccessControllerRevokeAccessV1

> AccessControllerRevokeAccessV1(ctx, id).RevokeAccessControllerParamsDto(revokeAccessControllerParamsDto).Execute()



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
	revokeAccessControllerParamsDto := *openapiclient.NewRevokeAccessControllerParamsDto("PrincipalType_example", "PrincipalId_example", "Access_example") // RevokeAccessControllerParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessAPI.AccessControllerRevokeAccessV1(context.Background(), id).RevokeAccessControllerParamsDto(revokeAccessControllerParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessAPI.AccessControllerRevokeAccessV1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAccessControllerRevokeAccessV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revokeAccessControllerParamsDto** | [**RevokeAccessControllerParamsDto**](RevokeAccessControllerParamsDto.md) |  | 

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

