/*
People info

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// DefaultAPIService DefaultAPI service
type DefaultAPIService service

type ApiInfoGetRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
	passportSerie *int32
	passportNumber *int32
}

func (r ApiInfoGetRequest) PassportSerie(passportSerie int32) ApiInfoGetRequest {
	r.passportSerie = &passportSerie
	return r
}

func (r ApiInfoGetRequest) PassportNumber(passportNumber int32) ApiInfoGetRequest {
	r.passportNumber = &passportNumber
	return r
}

func (r ApiInfoGetRequest) Execute() (*People, *http.Response, error) {
	return r.ApiService.InfoGetExecute(r)
}

/*
InfoGet Method for InfoGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInfoGetRequest
*/
func (a *DefaultAPIService) InfoGet(ctx context.Context) ApiInfoGetRequest {
	return ApiInfoGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return People
func (a *DefaultAPIService) InfoGetExecute(r ApiInfoGetRequest) (*People, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *People
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.InfoGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.passportSerie == nil {
		return localVarReturnValue, nil, reportError("passportSerie is required and must be specified")
	}
	if r.passportNumber == nil {
		return localVarReturnValue, nil, reportError("passportNumber is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "passportSerie", r.passportSerie, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "passportNumber", r.passportNumber, "")
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
