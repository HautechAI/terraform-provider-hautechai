# \PermissionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermissionsControllerListAvailablePermissionsV1**](PermissionsAPI.md#PermissionsControllerListAvailablePermissionsV1) | **Get** /v1/permissions/available | 



## PermissionsControllerListAvailablePermissionsV1

> []string PermissionsControllerListAvailablePermissionsV1(ctx).Execute()



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
	resp, r, err := apiClient.PermissionsAPI.PermissionsControllerListAvailablePermissionsV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.PermissionsControllerListAvailablePermissionsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionsControllerListAvailablePermissionsV1`: []string
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.PermissionsControllerListAvailablePermissionsV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsControllerListAvailablePermissionsV1Request struct via the builder pattern


### Return type

**[]string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

