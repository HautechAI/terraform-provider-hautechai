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


type PermissionsAPI interface {

	/*
	PermissionsControllerListAvailablePermissionsV1 Method for PermissionsControllerListAvailablePermissionsV1

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPermissionsControllerListAvailablePermissionsV1Request
	*/
	PermissionsControllerListAvailablePermissionsV1(ctx context.Context) ApiPermissionsControllerListAvailablePermissionsV1Request

	// PermissionsControllerListAvailablePermissionsV1Execute executes the request
	//  @return []string
	PermissionsControllerListAvailablePermissionsV1Execute(r ApiPermissionsControllerListAvailablePermissionsV1Request) ([]string, *http.Response, error)
}

// PermissionsAPIService PermissionsAPI service
type PermissionsAPIService service

type ApiPermissionsControllerListAvailablePermissionsV1Request struct {
	ctx context.Context
	ApiService PermissionsAPI
}

func (r ApiPermissionsControllerListAvailablePermissionsV1Request) Execute() ([]string, *http.Response, error) {
	return r.ApiService.PermissionsControllerListAvailablePermissionsV1Execute(r)
}

/*
PermissionsControllerListAvailablePermissionsV1 Method for PermissionsControllerListAvailablePermissionsV1

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPermissionsControllerListAvailablePermissionsV1Request
*/
func (a *PermissionsAPIService) PermissionsControllerListAvailablePermissionsV1(ctx context.Context) ApiPermissionsControllerListAvailablePermissionsV1Request {
	return ApiPermissionsControllerListAvailablePermissionsV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []string
func (a *PermissionsAPIService) PermissionsControllerListAvailablePermissionsV1Execute(r ApiPermissionsControllerListAvailablePermissionsV1Request) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionsAPIService.PermissionsControllerListAvailablePermissionsV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/permissions/available"

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
