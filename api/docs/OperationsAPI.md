# \OperationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OperationsControllerGetOperationV1**](OperationsAPI.md#OperationsControllerGetOperationV1) | **Get** /v1/operations/{id} | 
[**OperationsControllerGetOperationsV1**](OperationsAPI.md#OperationsControllerGetOperationsV1) | **Post** /v1/operations/many | 
[**OperationsControllerListOperationsV1**](OperationsAPI.md#OperationsControllerListOperationsV1) | **Get** /v1/operations | 
[**OperationsControllerRunCompositeV1V1**](OperationsAPI.md#OperationsControllerRunCompositeV1V1) | **Post** /v1/operations/run/composite.v1 | 
[**OperationsControllerRunCropV1V1**](OperationsAPI.md#OperationsControllerRunCropV1V1) | **Post** /v1/operations/run/crop.v1 | 
[**OperationsControllerRunCutV1V1**](OperationsAPI.md#OperationsControllerRunCutV1V1) | **Post** /v1/operations/run/cut.v1 | 
[**OperationsControllerRunGptV1V1**](OperationsAPI.md#OperationsControllerRunGptV1V1) | **Post** /v1/operations/run/gpt.v1 | 
[**OperationsControllerRunHauteLindaV1V1**](OperationsAPI.md#OperationsControllerRunHauteLindaV1V1) | **Post** /v1/operations/run/haute.linda.v1 | 
[**OperationsControllerRunHauteNaomiV1V1**](OperationsAPI.md#OperationsControllerRunHauteNaomiV1V1) | **Post** /v1/operations/run/haute.naomi.v1 | 
[**OperationsControllerRunImagineKateV1V1**](OperationsAPI.md#OperationsControllerRunImagineKateV1V1) | **Post** /v1/operations/run/imagine.kate.v1 | 
[**OperationsControllerRunInpaintKateV1V1**](OperationsAPI.md#OperationsControllerRunInpaintKateV1V1) | **Post** /v1/operations/run/inpaint.kate.v1 | 
[**OperationsControllerRunNegateImageV1V1**](OperationsAPI.md#OperationsControllerRunNegateImageV1V1) | **Post** /v1/operations/run/negateImage.v1 | 
[**OperationsControllerRunNoiseV1V1**](OperationsAPI.md#OperationsControllerRunNoiseV1V1) | **Post** /v1/operations/run/noise.v1 | 
[**OperationsControllerRunObjectDetectionV1V1**](OperationsAPI.md#OperationsControllerRunObjectDetectionV1V1) | **Post** /v1/operations/run/objectDetection.v1 | 
[**OperationsControllerRunPoseEstimationV1V1**](OperationsAPI.md#OperationsControllerRunPoseEstimationV1V1) | **Post** /v1/operations/run/poseEstimation.v1 | 
[**OperationsControllerRunSegmentAnythingEmbeddingsV1V1**](OperationsAPI.md#OperationsControllerRunSegmentAnythingEmbeddingsV1V1) | **Post** /v1/operations/run/segmentAnything.embeddings.v1 | 
[**OperationsControllerRunSegmentAnythingMaskV1V1**](OperationsAPI.md#OperationsControllerRunSegmentAnythingMaskV1V1) | **Post** /v1/operations/run/segmentAnything.mask.v1 | 
[**OperationsControllerRunUpscaleV1V1**](OperationsAPI.md#OperationsControllerRunUpscaleV1V1) | **Post** /v1/operations/run/upscale.v1 | 
[**OperationsControllerRunVtonGiseleV1V1**](OperationsAPI.md#OperationsControllerRunVtonGiseleV1V1) | **Post** /v1/operations/run/vton.gisele.v1 | 
[**OperationsControllerUpdateMetadataV1**](OperationsAPI.md#OperationsControllerUpdateMetadataV1) | **Put** /v1/operations/{id}/metadata | 



## OperationsControllerGetOperationV1

> OperationEntity OperationsControllerGetOperationV1(ctx, id).Execute()



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
	resp, r, err := apiClient.OperationsAPI.OperationsControllerGetOperationV1(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerGetOperationV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerGetOperationV1`: OperationEntity
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerGetOperationV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerGetOperationV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationEntity**](OperationEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerGetOperationsV1

> []OperationEntity OperationsControllerGetOperationsV1(ctx).GetOperationsParamsDto(getOperationsParamsDto).Execute()



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
	getOperationsParamsDto := *openapiclient.NewGetOperationsParamsDto([]string{"Ids_example"}) // GetOperationsParamsDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.OperationsControllerGetOperationsV1(context.Background()).GetOperationsParamsDto(getOperationsParamsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerGetOperationsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerGetOperationsV1`: []OperationEntity
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerGetOperationsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerGetOperationsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getOperationsParamsDto** | [**GetOperationsParamsDto**](GetOperationsParamsDto.md) |  | 

### Return type

[**[]OperationEntity**](OperationEntity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerListOperationsV1

> ListOperationsDto OperationsControllerListOperationsV1(ctx).OrderBy(orderBy).Limit(limit).Cursor(cursor).Execute()



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
	resp, r, err := apiClient.OperationsAPI.OperationsControllerListOperationsV1(context.Background()).OrderBy(orderBy).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerListOperationsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerListOperationsV1`: ListOperationsDto
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerListOperationsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerListOperationsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **string** |  | [default to &quot;createdAt_DESC&quot;]
 **limit** | **float32** |  | [default to 50]
 **cursor** | **string** |  | 

### Return type

[**ListOperationsDto**](ListOperationsDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerRunCompositeV1V1

> CompositeV1Response OperationsControllerRunCompositeV1V1(ctx).CompositeV1Request(compositeV1Request).Execute()



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
	compositeV1Request := *openapiclient.NewCompositeV1Request(*openapiclient.NewCompositeV1Input(float32(123), float32(123), "Background_example", []openapiclient.CompositeElement{*openapiclient.NewCompositeElement("ImageId_example", float32(123), float32(123), float32(123), float32(123), "Fit_example")})) // CompositeV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.OperationsControllerRunCompositeV1V1(context.Background()).CompositeV1Request(compositeV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerRunCompositeV1V1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerRunCompositeV1V1`: CompositeV1Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerRunCompositeV1V1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerRunCompositeV1V1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compositeV1Request** | [**CompositeV1Request**](CompositeV1Request.md) |  | 

### Return type

[**CompositeV1Response**](CompositeV1Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerRunCropV1V1

> CropV1Response OperationsControllerRunCropV1V1(ctx).CropV1Request(cropV1Request).Execute()



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
	cropV1Request := *openapiclient.NewCropV1Request(*openapiclient.NewCropV1Input("ImageId_example", float32(123), float32(123), float32(123), float32(123))) // CropV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.OperationsControllerRunCropV1V1(context.Background()).CropV1Request(cropV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerRunCropV1V1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerRunCropV1V1`: CropV1Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerRunCropV1V1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerRunCropV1V1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cropV1Request** | [**CropV1Request**](CropV1Request.md) |  | 

### Return type

[**CropV1Response**](CropV1Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerRunCutV1V1

> CutV1Response OperationsControllerRunCutV1V1(ctx).CutV1Request(cutV1Request).Execute()



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
	cutV1Request := *openapiclient.NewCutV1Request(*openapiclient.NewCutV1Input("ImageId_example", "MaskImageId_example")) // CutV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.OperationsControllerRunCutV1V1(context.Background()).CutV1Request(cutV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerRunCutV1V1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerRunCutV1V1`: CutV1Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerRunCutV1V1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerRunCutV1V1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cutV1Request** | [**CutV1Request**](CutV1Request.md) |  | 

### Return type

[**CutV1Response**](CutV1Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerRunGptV1V1

> GptV1Response OperationsControllerRunGptV1V1(ctx).GptV1Request(gptV1Request).Execute()



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
	gptV1Request := *openapiclient.NewGptV1Request(*openapiclient.NewGPTV1Input("Prompt_example")) // GptV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.OperationsControllerRunGptV1V1(context.Background()).GptV1Request(gptV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerRunGptV1V1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerRunGptV1V1`: GptV1Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerRunGptV1V1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerRunGptV1V1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gptV1Request** | [**GptV1Request**](GptV1Request.md) |  | 

### Return type

[**GptV1Response**](GptV1Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerRunHauteLindaV1V1

> HauteLindaV1Response OperationsControllerRunHauteLindaV1V1(ctx).HauteLindaV1Request(hauteLindaV1Request).Execute()



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
	hauteLindaV1Request := *openapiclient.NewHauteLindaV1Request(*openapiclient.NewLindaHauteV1Input("AspectRatio_example", "ProductImageId_example", "Prompt_example", float32(123))) // HauteLindaV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.OperationsControllerRunHauteLindaV1V1(context.Background()).HauteLindaV1Request(hauteLindaV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerRunHauteLindaV1V1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerRunHauteLindaV1V1`: HauteLindaV1Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerRunHauteLindaV1V1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerRunHauteLindaV1V1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hauteLindaV1Request** | [**HauteLindaV1Request**](HauteLindaV1Request.md) |  | 

### Return type

[**HauteLindaV1Response**](HauteLindaV1Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerRunHauteNaomiV1V1

> HauteNaomiV1Response OperationsControllerRunHauteNaomiV1V1(ctx).HauteNaomiV1Request(hauteNaomiV1Request).Execute()



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
	hauteNaomiV1Request := *openapiclient.NewHauteNaomiV1Request(*openapiclient.NewNaomiHauteV1Input("Prompt_example", "Category_example", "GarmentImageId_example", "PoseId_example", float32(123), float32(123), float32(123))) // HauteNaomiV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.OperationsControllerRunHauteNaomiV1V1(context.Background()).HauteNaomiV1Request(hauteNaomiV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerRunHauteNaomiV1V1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerRunHauteNaomiV1V1`: HauteNaomiV1Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerRunHauteNaomiV1V1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerRunHauteNaomiV1V1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hauteNaomiV1Request** | [**HauteNaomiV1Request**](HauteNaomiV1Request.md) |  | 

### Return type

[**HauteNaomiV1Response**](HauteNaomiV1Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerRunImagineKateV1V1

> ImagineKateV1Response OperationsControllerRunImagineKateV1V1(ctx).ImagineKateV1Request(imagineKateV1Request).Execute()



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
	imagineKateV1Request := *openapiclient.NewImagineKateV1Request(*openapiclient.NewKateImagineV1Input("AspectRatio_example", "Seed_example", "Prompt_example")) // ImagineKateV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.OperationsControllerRunImagineKateV1V1(context.Background()).ImagineKateV1Request(imagineKateV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerRunImagineKateV1V1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerRunImagineKateV1V1`: ImagineKateV1Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerRunImagineKateV1V1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerRunImagineKateV1V1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imagineKateV1Request** | [**ImagineKateV1Request**](ImagineKateV1Request.md) |  | 

### Return type

[**ImagineKateV1Response**](ImagineKateV1Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerRunInpaintKateV1V1

> InpaintKateV1Response OperationsControllerRunInpaintKateV1V1(ctx).InpaintKateV1Request(inpaintKateV1Request).Execute()



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
	inpaintKateV1Request := *openapiclient.NewInpaintKateV1Request(*openapiclient.NewKateInpaintV1Input("ImageId_example", "Prompt_example", float32(123), float32(123))) // InpaintKateV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.OperationsControllerRunInpaintKateV1V1(context.Background()).InpaintKateV1Request(inpaintKateV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerRunInpaintKateV1V1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerRunInpaintKateV1V1`: InpaintKateV1Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerRunInpaintKateV1V1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerRunInpaintKateV1V1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inpaintKateV1Request** | [**InpaintKateV1Request**](InpaintKateV1Request.md) |  | 

### Return type

[**InpaintKateV1Response**](InpaintKateV1Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerRunNegateImageV1V1

> NegateImageV1Response OperationsControllerRunNegateImageV1V1(ctx).NegateImageV1Request(negateImageV1Request).Execute()



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
	negateImageV1Request := *openapiclient.NewNegateImageV1Request(*openapiclient.NewNegateImageV1Input("ImageId_example")) // NegateImageV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.OperationsControllerRunNegateImageV1V1(context.Background()).NegateImageV1Request(negateImageV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerRunNegateImageV1V1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerRunNegateImageV1V1`: NegateImageV1Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerRunNegateImageV1V1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerRunNegateImageV1V1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **negateImageV1Request** | [**NegateImageV1Request**](NegateImageV1Request.md) |  | 

### Return type

[**NegateImageV1Response**](NegateImageV1Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerRunNoiseV1V1

> NoiseV1Response OperationsControllerRunNoiseV1V1(ctx).NoiseV1Request(noiseV1Request).Execute()



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
	noiseV1Request := *openapiclient.NewNoiseV1Request(*openapiclient.NewNoiseV1Input("ImageId_example")) // NoiseV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.OperationsControllerRunNoiseV1V1(context.Background()).NoiseV1Request(noiseV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerRunNoiseV1V1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerRunNoiseV1V1`: NoiseV1Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerRunNoiseV1V1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerRunNoiseV1V1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noiseV1Request** | [**NoiseV1Request**](NoiseV1Request.md) |  | 

### Return type

[**NoiseV1Response**](NoiseV1Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerRunObjectDetectionV1V1

> ObjectDetectionV1Response OperationsControllerRunObjectDetectionV1V1(ctx).ObjectDetectionV1Request(objectDetectionV1Request).Execute()



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
	objectDetectionV1Request := *openapiclient.NewObjectDetectionV1Request(*openapiclient.NewObjectDetectionV1Input([]string{"Labels_example"}, "ImageId_example")) // ObjectDetectionV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.OperationsControllerRunObjectDetectionV1V1(context.Background()).ObjectDetectionV1Request(objectDetectionV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerRunObjectDetectionV1V1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerRunObjectDetectionV1V1`: ObjectDetectionV1Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerRunObjectDetectionV1V1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerRunObjectDetectionV1V1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectDetectionV1Request** | [**ObjectDetectionV1Request**](ObjectDetectionV1Request.md) |  | 

### Return type

[**ObjectDetectionV1Response**](ObjectDetectionV1Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerRunPoseEstimationV1V1

> PoseEstimationV1Response OperationsControllerRunPoseEstimationV1V1(ctx).PoseEstimationV1Request(poseEstimationV1Request).Execute()



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
	poseEstimationV1Request := *openapiclient.NewPoseEstimationV1Request(*openapiclient.NewPoseEstimationV1Input("ImageId_example")) // PoseEstimationV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.OperationsControllerRunPoseEstimationV1V1(context.Background()).PoseEstimationV1Request(poseEstimationV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerRunPoseEstimationV1V1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerRunPoseEstimationV1V1`: PoseEstimationV1Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerRunPoseEstimationV1V1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerRunPoseEstimationV1V1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poseEstimationV1Request** | [**PoseEstimationV1Request**](PoseEstimationV1Request.md) |  | 

### Return type

[**PoseEstimationV1Response**](PoseEstimationV1Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerRunSegmentAnythingEmbeddingsV1V1

> SegmentAnythingEmbeddingsV1Response OperationsControllerRunSegmentAnythingEmbeddingsV1V1(ctx).SegmentAnythingEmbeddingsV1Request(segmentAnythingEmbeddingsV1Request).Execute()



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
	segmentAnythingEmbeddingsV1Request := *openapiclient.NewSegmentAnythingEmbeddingsV1Request(*openapiclient.NewSegmentAnythingEmbeddingsV1Input("ImageId_example")) // SegmentAnythingEmbeddingsV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.OperationsControllerRunSegmentAnythingEmbeddingsV1V1(context.Background()).SegmentAnythingEmbeddingsV1Request(segmentAnythingEmbeddingsV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerRunSegmentAnythingEmbeddingsV1V1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerRunSegmentAnythingEmbeddingsV1V1`: SegmentAnythingEmbeddingsV1Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerRunSegmentAnythingEmbeddingsV1V1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerRunSegmentAnythingEmbeddingsV1V1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **segmentAnythingEmbeddingsV1Request** | [**SegmentAnythingEmbeddingsV1Request**](SegmentAnythingEmbeddingsV1Request.md) |  | 

### Return type

[**SegmentAnythingEmbeddingsV1Response**](SegmentAnythingEmbeddingsV1Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerRunSegmentAnythingMaskV1V1

> SegmentAnythingMaskV1Response OperationsControllerRunSegmentAnythingMaskV1V1(ctx).SegmentAnythingMaskV1Request(segmentAnythingMaskV1Request).Execute()



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
	segmentAnythingMaskV1Request := *openapiclient.NewSegmentAnythingMaskV1Request(*openapiclient.NewSegmentAnythingMaskV1Input("ImageId_example", []float32{float32(123)})) // SegmentAnythingMaskV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.OperationsControllerRunSegmentAnythingMaskV1V1(context.Background()).SegmentAnythingMaskV1Request(segmentAnythingMaskV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerRunSegmentAnythingMaskV1V1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerRunSegmentAnythingMaskV1V1`: SegmentAnythingMaskV1Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerRunSegmentAnythingMaskV1V1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerRunSegmentAnythingMaskV1V1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **segmentAnythingMaskV1Request** | [**SegmentAnythingMaskV1Request**](SegmentAnythingMaskV1Request.md) |  | 

### Return type

[**SegmentAnythingMaskV1Response**](SegmentAnythingMaskV1Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerRunUpscaleV1V1

> UpscaleV1Response OperationsControllerRunUpscaleV1V1(ctx).UpscaleV1Request(upscaleV1Request).Execute()



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
	upscaleV1Request := *openapiclient.NewUpscaleV1Request(*openapiclient.NewUpscaleV1Input("ImageId_example")) // UpscaleV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.OperationsControllerRunUpscaleV1V1(context.Background()).UpscaleV1Request(upscaleV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerRunUpscaleV1V1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerRunUpscaleV1V1`: UpscaleV1Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerRunUpscaleV1V1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerRunUpscaleV1V1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upscaleV1Request** | [**UpscaleV1Request**](UpscaleV1Request.md) |  | 

### Return type

[**UpscaleV1Response**](UpscaleV1Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerRunVtonGiseleV1V1

> VtonGiseleV1Response OperationsControllerRunVtonGiseleV1V1(ctx).VtonGiseleV1Request(vtonGiseleV1Request).Execute()



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
	vtonGiseleV1Request := *openapiclient.NewVtonGiseleV1Request(*openapiclient.NewGiseleVtonV1Input("Category_example", "ImageId_example", "ProductDescription_example", "ProductImageId_example", float32(123))) // VtonGiseleV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.OperationsControllerRunVtonGiseleV1V1(context.Background()).VtonGiseleV1Request(vtonGiseleV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerRunVtonGiseleV1V1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerRunVtonGiseleV1V1`: VtonGiseleV1Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerRunVtonGiseleV1V1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerRunVtonGiseleV1V1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vtonGiseleV1Request** | [**VtonGiseleV1Request**](VtonGiseleV1Request.md) |  | 

### Return type

[**VtonGiseleV1Response**](VtonGiseleV1Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsControllerUpdateMetadataV1

> ResourceEntity OperationsControllerUpdateMetadataV1(ctx, id).UpdateMetadataDto(updateMetadataDto).Execute()



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
	resp, r, err := apiClient.OperationsAPI.OperationsControllerUpdateMetadataV1(context.Background(), id).UpdateMetadataDto(updateMetadataDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.OperationsControllerUpdateMetadataV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperationsControllerUpdateMetadataV1`: ResourceEntity
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.OperationsControllerUpdateMetadataV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperationsControllerUpdateMetadataV1Request struct via the builder pattern


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

