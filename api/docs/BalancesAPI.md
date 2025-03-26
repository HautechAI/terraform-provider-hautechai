# \BalancesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BalancesControllerAddBalanceV1**](BalancesAPI.md#BalancesControllerAddBalanceV1) | **Put** /v1/accounts/{id}/balance | 
[**BalancesControllerGetBalanceForSelfV1**](BalancesAPI.md#BalancesControllerGetBalanceForSelfV1) | **Get** /v1/accounts/self/balance | 
[**BalancesControllerGetBalanceV1**](BalancesAPI.md#BalancesControllerGetBalanceV1) | **Get** /v1/accounts/{id}/balance | 



## BalancesControllerAddBalanceV1

> BalancesControllerAddBalanceV1(ctx, id).AddBalanceControllerParamsDto(addBalanceControllerParamsDto).Execute()



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
	addBalanceControllerParamsDto := *openapiclient.NewAddBalanceControllerParamsDto("Amount_example") // AddBalanceControllerParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BalancesAPI.BalancesControllerAddBalanceV1(context.Background(), id).AddBalanceControllerParamsDto(addBalanceControllerParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BalancesAPI.BalancesControllerAddBalanceV1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBalancesControllerAddBalanceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addBalanceControllerParamsDto** | [**AddBalanceControllerParamsDto**](AddBalanceControllerParamsDto.md) |  | 

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


## BalancesControllerGetBalanceForSelfV1

> BalanceResultDto BalancesControllerGetBalanceForSelfV1(ctx).Execute()



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
	resp, r, err := apiClient.BalancesAPI.BalancesControllerGetBalanceForSelfV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BalancesAPI.BalancesControllerGetBalanceForSelfV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BalancesControllerGetBalanceForSelfV1`: BalanceResultDto
	fmt.Fprintf(os.Stdout, "Response from `BalancesAPI.BalancesControllerGetBalanceForSelfV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBalancesControllerGetBalanceForSelfV1Request struct via the builder pattern


### Return type

[**BalanceResultDto**](BalanceResultDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BalancesControllerGetBalanceV1

> BalanceResultDto BalancesControllerGetBalanceV1(ctx, id).Execute()



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
	resp, r, err := apiClient.BalancesAPI.BalancesControllerGetBalanceV1(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BalancesAPI.BalancesControllerGetBalanceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BalancesControllerGetBalanceV1`: BalanceResultDto
	fmt.Fprintf(os.Stdout, "Response from `BalancesAPI.BalancesControllerGetBalanceV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBalancesControllerGetBalanceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BalanceResultDto**](BalanceResultDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

