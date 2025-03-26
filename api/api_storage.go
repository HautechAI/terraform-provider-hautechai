/*
Hautech API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hautechapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


type StorageAPI interface {

	/*
	StorageControllerCreateRecordV1 Method for StorageControllerCreateRecordV1

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiStorageControllerCreateRecordV1Request
	*/
	StorageControllerCreateRecordV1(ctx context.Context) ApiStorageControllerCreateRecordV1Request

	// StorageControllerCreateRecordV1Execute executes the request
	//  @return StorageEntity
	StorageControllerCreateRecordV1Execute(r ApiStorageControllerCreateRecordV1Request) (*StorageEntity, *http.Response, error)

	/*
	StorageControllerDeleteRecordV1 Method for StorageControllerDeleteRecordV1

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiStorageControllerDeleteRecordV1Request
	*/
	StorageControllerDeleteRecordV1(ctx context.Context) ApiStorageControllerDeleteRecordV1Request

	// StorageControllerDeleteRecordV1Execute executes the request
	StorageControllerDeleteRecordV1Execute(r ApiStorageControllerDeleteRecordV1Request) (*http.Response, error)

	/*
	StorageControllerGetRecordsV1 Method for StorageControllerGetRecordsV1

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiStorageControllerGetRecordsV1Request
	*/
	StorageControllerGetRecordsV1(ctx context.Context) ApiStorageControllerGetRecordsV1Request

	// StorageControllerGetRecordsV1Execute executes the request
	//  @return []StorageRecordsResultDto
	StorageControllerGetRecordsV1Execute(r ApiStorageControllerGetRecordsV1Request) ([]StorageRecordsResultDto, *http.Response, error)

	/*
	StorageControllerUpdateRecordV1 Method for StorageControllerUpdateRecordV1

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiStorageControllerUpdateRecordV1Request
	*/
	StorageControllerUpdateRecordV1(ctx context.Context) ApiStorageControllerUpdateRecordV1Request

	// StorageControllerUpdateRecordV1Execute executes the request
	//  @return StorageEntity
	StorageControllerUpdateRecordV1Execute(r ApiStorageControllerUpdateRecordV1Request) (*StorageEntity, *http.Response, error)
}

// StorageAPIService StorageAPI service
type StorageAPIService service

type ApiStorageControllerCreateRecordV1Request struct {
	ctx context.Context
	ApiService StorageAPI
	createStorageRecordParamsDto *CreateStorageRecordParamsDto
}

func (r ApiStorageControllerCreateRecordV1Request) CreateStorageRecordParamsDto(createStorageRecordParamsDto CreateStorageRecordParamsDto) ApiStorageControllerCreateRecordV1Request {
	r.createStorageRecordParamsDto = &createStorageRecordParamsDto
	return r
}

func (r ApiStorageControllerCreateRecordV1Request) Execute() (*StorageEntity, *http.Response, error) {
	return r.ApiService.StorageControllerCreateRecordV1Execute(r)
}

/*
StorageControllerCreateRecordV1 Method for StorageControllerCreateRecordV1

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiStorageControllerCreateRecordV1Request
*/
func (a *StorageAPIService) StorageControllerCreateRecordV1(ctx context.Context) ApiStorageControllerCreateRecordV1Request {
	return ApiStorageControllerCreateRecordV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StorageEntity
func (a *StorageAPIService) StorageControllerCreateRecordV1Execute(r ApiStorageControllerCreateRecordV1Request) (*StorageEntity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StorageEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StorageAPIService.StorageControllerCreateRecordV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/storage"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createStorageRecordParamsDto == nil {
		return localVarReturnValue, nil, reportError("createStorageRecordParamsDto is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createStorageRecordParamsDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiStorageControllerDeleteRecordV1Request struct {
	ctx context.Context
	ApiService StorageAPI
	deleteStorageParamsDto *DeleteStorageParamsDto
}

func (r ApiStorageControllerDeleteRecordV1Request) DeleteStorageParamsDto(deleteStorageParamsDto DeleteStorageParamsDto) ApiStorageControllerDeleteRecordV1Request {
	r.deleteStorageParamsDto = &deleteStorageParamsDto
	return r
}

func (r ApiStorageControllerDeleteRecordV1Request) Execute() (*http.Response, error) {
	return r.ApiService.StorageControllerDeleteRecordV1Execute(r)
}

/*
StorageControllerDeleteRecordV1 Method for StorageControllerDeleteRecordV1

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiStorageControllerDeleteRecordV1Request
*/
func (a *StorageAPIService) StorageControllerDeleteRecordV1(ctx context.Context) ApiStorageControllerDeleteRecordV1Request {
	return ApiStorageControllerDeleteRecordV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *StorageAPIService) StorageControllerDeleteRecordV1Execute(r ApiStorageControllerDeleteRecordV1Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StorageAPIService.StorageControllerDeleteRecordV1")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/storage/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deleteStorageParamsDto == nil {
		return nil, reportError("deleteStorageParamsDto is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.deleteStorageParamsDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiStorageControllerGetRecordsV1Request struct {
	ctx context.Context
	ApiService StorageAPI
	getStorageRecordParamsDto *GetStorageRecordParamsDto
}

func (r ApiStorageControllerGetRecordsV1Request) GetStorageRecordParamsDto(getStorageRecordParamsDto GetStorageRecordParamsDto) ApiStorageControllerGetRecordsV1Request {
	r.getStorageRecordParamsDto = &getStorageRecordParamsDto
	return r
}

func (r ApiStorageControllerGetRecordsV1Request) Execute() ([]StorageRecordsResultDto, *http.Response, error) {
	return r.ApiService.StorageControllerGetRecordsV1Execute(r)
}

/*
StorageControllerGetRecordsV1 Method for StorageControllerGetRecordsV1

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiStorageControllerGetRecordsV1Request
*/
func (a *StorageAPIService) StorageControllerGetRecordsV1(ctx context.Context) ApiStorageControllerGetRecordsV1Request {
	return ApiStorageControllerGetRecordsV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []StorageRecordsResultDto
func (a *StorageAPIService) StorageControllerGetRecordsV1Execute(r ApiStorageControllerGetRecordsV1Request) ([]StorageRecordsResultDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []StorageRecordsResultDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StorageAPIService.StorageControllerGetRecordsV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/storage/many"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getStorageRecordParamsDto == nil {
		return localVarReturnValue, nil, reportError("getStorageRecordParamsDto is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.getStorageRecordParamsDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiStorageControllerUpdateRecordV1Request struct {
	ctx context.Context
	ApiService StorageAPI
	updateStorageRecordParamsDto *UpdateStorageRecordParamsDto
}

func (r ApiStorageControllerUpdateRecordV1Request) UpdateStorageRecordParamsDto(updateStorageRecordParamsDto UpdateStorageRecordParamsDto) ApiStorageControllerUpdateRecordV1Request {
	r.updateStorageRecordParamsDto = &updateStorageRecordParamsDto
	return r
}

func (r ApiStorageControllerUpdateRecordV1Request) Execute() (*StorageEntity, *http.Response, error) {
	return r.ApiService.StorageControllerUpdateRecordV1Execute(r)
}

/*
StorageControllerUpdateRecordV1 Method for StorageControllerUpdateRecordV1

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiStorageControllerUpdateRecordV1Request
*/
func (a *StorageAPIService) StorageControllerUpdateRecordV1(ctx context.Context) ApiStorageControllerUpdateRecordV1Request {
	return ApiStorageControllerUpdateRecordV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StorageEntity
func (a *StorageAPIService) StorageControllerUpdateRecordV1Execute(r ApiStorageControllerUpdateRecordV1Request) (*StorageEntity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StorageEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StorageAPIService.StorageControllerUpdateRecordV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/storage/write"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateStorageRecordParamsDto == nil {
		return localVarReturnValue, nil, reportError("updateStorageRecordParamsDto is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateStorageRecordParamsDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
