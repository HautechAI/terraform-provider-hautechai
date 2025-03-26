# \AccountsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsControllerCreateAccountV1**](AccountsAPI.md#AccountsControllerCreateAccountV1) | **Post** /v1/accounts | 
[**AccountsControllerGetAccountByAliasV1**](AccountsAPI.md#AccountsControllerGetAccountByAliasV1) | **Get** /v1/accounts/alias/{id} | 
[**AccountsControllerGetAccountV1**](AccountsAPI.md#AccountsControllerGetAccountV1) | **Get** /v1/accounts/{id} | 
[**AccountsControllerGetSelfV1**](AccountsAPI.md#AccountsControllerGetSelfV1) | **Get** /v1/accounts/self | 
[**AccountsControllerListAccountsV1**](AccountsAPI.md#AccountsControllerListAccountsV1) | **Get** /v1/accounts | 



## AccountsControllerCreateAccountV1

> AccountEntity AccountsControllerCreateAccountV1(ctx).CreateAccountParamsDto(createAccountParamsDto).Execute()



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
	createAccountParamsDto := *openapiclient.NewCreateAccountParamsDto() // CreateAccountParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.AccountsControllerCreateAccountV1(context.Background()).CreateAccountParamsDto(createAccountParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountsControllerCreateAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountsControllerCreateAccountV1`: AccountEntity
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountsControllerCreateAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsControllerCreateAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountParamsDto** | [**CreateAccountParamsDto**](CreateAccountParamsDto.md) |  | 

### Return type

[**AccountEntity**](AccountEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsControllerGetAccountByAliasV1

> AccountEntity AccountsControllerGetAccountByAliasV1(ctx, id).Execute()



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
	resp, r, err := apiClient.AccountsAPI.AccountsControllerGetAccountByAliasV1(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountsControllerGetAccountByAliasV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountsControllerGetAccountByAliasV1`: AccountEntity
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountsControllerGetAccountByAliasV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsControllerGetAccountByAliasV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountEntity**](AccountEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsControllerGetAccountV1

> AccountEntity AccountsControllerGetAccountV1(ctx, id).Execute()



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
	resp, r, err := apiClient.AccountsAPI.AccountsControllerGetAccountV1(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountsControllerGetAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountsControllerGetAccountV1`: AccountEntity
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountsControllerGetAccountV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsControllerGetAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountEntity**](AccountEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsControllerGetSelfV1

> AccountEntity AccountsControllerGetSelfV1(ctx).Execute()



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
	resp, r, err := apiClient.AccountsAPI.AccountsControllerGetSelfV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountsControllerGetSelfV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountsControllerGetSelfV1`: AccountEntity
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountsControllerGetSelfV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsControllerGetSelfV1Request struct via the builder pattern


### Return type

[**AccountEntity**](AccountEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsControllerListAccountsV1

> ListAccountsDto AccountsControllerListAccountsV1(ctx).OrderBy(orderBy).Limit(limit).Cursor(cursor).Execute()



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
	resp, r, err := apiClient.AccountsAPI.AccountsControllerListAccountsV1(context.Background()).OrderBy(orderBy).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountsControllerListAccountsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountsControllerListAccountsV1`: ListAccountsDto
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountsControllerListAccountsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsControllerListAccountsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **string** |  | [default to &quot;createdAt_DESC&quot;]
 **limit** | **float32** |  | [default to 50]
 **cursor** | **string** |  | 

### Return type

[**ListAccountsDto**](ListAccountsDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

