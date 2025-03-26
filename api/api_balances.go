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
	"strings"
)


type BalancesAPI interface {

	/*
	BalancesControllerAddBalanceV1 Method for BalancesControllerAddBalanceV1

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiBalancesControllerAddBalanceV1Request
	*/
	BalancesControllerAddBalanceV1(ctx context.Context, id string) ApiBalancesControllerAddBalanceV1Request

	// BalancesControllerAddBalanceV1Execute executes the request
	BalancesControllerAddBalanceV1Execute(r ApiBalancesControllerAddBalanceV1Request) (*http.Response, error)

	/*
	BalancesControllerGetBalanceForSelfV1 Method for BalancesControllerGetBalanceForSelfV1

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiBalancesControllerGetBalanceForSelfV1Request
	*/
	BalancesControllerGetBalanceForSelfV1(ctx context.Context) ApiBalancesControllerGetBalanceForSelfV1Request

	// BalancesControllerGetBalanceForSelfV1Execute executes the request
	//  @return BalanceResultDto
	BalancesControllerGetBalanceForSelfV1Execute(r ApiBalancesControllerGetBalanceForSelfV1Request) (*BalanceResultDto, *http.Response, error)

	/*
	BalancesControllerGetBalanceV1 Method for BalancesControllerGetBalanceV1

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiBalancesControllerGetBalanceV1Request
	*/
	BalancesControllerGetBalanceV1(ctx context.Context, id string) ApiBalancesControllerGetBalanceV1Request

	// BalancesControllerGetBalanceV1Execute executes the request
	//  @return BalanceResultDto
	BalancesControllerGetBalanceV1Execute(r ApiBalancesControllerGetBalanceV1Request) (*BalanceResultDto, *http.Response, error)
}

// BalancesAPIService BalancesAPI service
type BalancesAPIService service

type ApiBalancesControllerAddBalanceV1Request struct {
	ctx context.Context
	ApiService BalancesAPI
	id string
	addBalanceControllerParamsDto *AddBalanceControllerParamsDto
}

func (r ApiBalancesControllerAddBalanceV1Request) AddBalanceControllerParamsDto(addBalanceControllerParamsDto AddBalanceControllerParamsDto) ApiBalancesControllerAddBalanceV1Request {
	r.addBalanceControllerParamsDto = &addBalanceControllerParamsDto
	return r
}

func (r ApiBalancesControllerAddBalanceV1Request) Execute() (*http.Response, error) {
	return r.ApiService.BalancesControllerAddBalanceV1Execute(r)
}

/*
BalancesControllerAddBalanceV1 Method for BalancesControllerAddBalanceV1

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiBalancesControllerAddBalanceV1Request
*/
func (a *BalancesAPIService) BalancesControllerAddBalanceV1(ctx context.Context, id string) ApiBalancesControllerAddBalanceV1Request {
	return ApiBalancesControllerAddBalanceV1Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *BalancesAPIService) BalancesControllerAddBalanceV1Execute(r ApiBalancesControllerAddBalanceV1Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BalancesAPIService.BalancesControllerAddBalanceV1")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/accounts/{id}/balance"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addBalanceControllerParamsDto == nil {
		return nil, reportError("addBalanceControllerParamsDto is required and must be specified")
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
	localVarPostBody = r.addBalanceControllerParamsDto
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

type ApiBalancesControllerGetBalanceForSelfV1Request struct {
	ctx context.Context
	ApiService BalancesAPI
}

func (r ApiBalancesControllerGetBalanceForSelfV1Request) Execute() (*BalanceResultDto, *http.Response, error) {
	return r.ApiService.BalancesControllerGetBalanceForSelfV1Execute(r)
}

/*
BalancesControllerGetBalanceForSelfV1 Method for BalancesControllerGetBalanceForSelfV1

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBalancesControllerGetBalanceForSelfV1Request
*/
func (a *BalancesAPIService) BalancesControllerGetBalanceForSelfV1(ctx context.Context) ApiBalancesControllerGetBalanceForSelfV1Request {
	return ApiBalancesControllerGetBalanceForSelfV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BalanceResultDto
func (a *BalancesAPIService) BalancesControllerGetBalanceForSelfV1Execute(r ApiBalancesControllerGetBalanceForSelfV1Request) (*BalanceResultDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BalanceResultDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BalancesAPIService.BalancesControllerGetBalanceForSelfV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/accounts/self/balance"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiBalancesControllerGetBalanceV1Request struct {
	ctx context.Context
	ApiService BalancesAPI
	id string
}

func (r ApiBalancesControllerGetBalanceV1Request) Execute() (*BalanceResultDto, *http.Response, error) {
	return r.ApiService.BalancesControllerGetBalanceV1Execute(r)
}

/*
BalancesControllerGetBalanceV1 Method for BalancesControllerGetBalanceV1

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiBalancesControllerGetBalanceV1Request
*/
func (a *BalancesAPIService) BalancesControllerGetBalanceV1(ctx context.Context, id string) ApiBalancesControllerGetBalanceV1Request {
	return ApiBalancesControllerGetBalanceV1Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BalanceResultDto
func (a *BalancesAPIService) BalancesControllerGetBalanceV1Execute(r ApiBalancesControllerGetBalanceV1Request) (*BalanceResultDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BalanceResultDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BalancesAPIService.BalancesControllerGetBalanceV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/accounts/{id}/balance"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
