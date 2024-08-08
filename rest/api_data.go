/*
Permify API

Permify is an open source authorization service for creating fine-grained and scalable authorization systems.

API version: v0.10.1
Contact: hello@permify.co
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package permify

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type DataAPI interface {

	/*
	BundleRun run bundle

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tenantId Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
	@return ApiBundleRunRequest
	*/
	BundleRun(ctx context.Context, tenantId string) ApiBundleRunRequest

	// BundleRunExecute executes the request
	//  @return BundleRunResponse
	BundleRunExecute(r ApiBundleRunRequest) (*BundleRunResponse, *http.Response, error)

	/*
	DataAttributesRead read attributes

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tenantId Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
	@return ApiDataAttributesReadRequest
	*/
	DataAttributesRead(ctx context.Context, tenantId string) ApiDataAttributesReadRequest

	// DataAttributesReadExecute executes the request
	//  @return AttributeReadResponse
	DataAttributesReadExecute(r ApiDataAttributesReadRequest) (*AttributeReadResponse, *http.Response, error)

	/*
	DataDelete delete data

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tenantId Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
	@return ApiDataDeleteRequest
	*/
	DataDelete(ctx context.Context, tenantId string) ApiDataDeleteRequest

	// DataDeleteExecute executes the request
	//  @return DataDeleteResponse
	DataDeleteExecute(r ApiDataDeleteRequest) (*DataDeleteResponse, *http.Response, error)

	/*
	DataRelationshipsRead read relationships

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tenantId Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
	@return ApiDataRelationshipsReadRequest
	*/
	DataRelationshipsRead(ctx context.Context, tenantId string) ApiDataRelationshipsReadRequest

	// DataRelationshipsReadExecute executes the request
	//  @return RelationshipReadResponse
	DataRelationshipsReadExecute(r ApiDataRelationshipsReadRequest) (*RelationshipReadResponse, *http.Response, error)

	/*
	DataWrite write data

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tenantId Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
	@return ApiDataWriteRequest
	*/
	DataWrite(ctx context.Context, tenantId string) ApiDataWriteRequest

	// DataWriteExecute executes the request
	//  @return DataWriteResponse
	DataWriteExecute(r ApiDataWriteRequest) (*DataWriteResponse, *http.Response, error)

	/*
	RelationshipsDelete delete relationships

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tenantId Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
	@return ApiRelationshipsDeleteRequest
	*/
	RelationshipsDelete(ctx context.Context, tenantId string) ApiRelationshipsDeleteRequest

	// RelationshipsDeleteExecute executes the request
	//  @return RelationshipDeleteResponse
	RelationshipsDeleteExecute(r ApiRelationshipsDeleteRequest) (*RelationshipDeleteResponse, *http.Response, error)

	/*
	RelationshipsWrite write relationships

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tenantId Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
	@return ApiRelationshipsWriteRequest
	*/
	RelationshipsWrite(ctx context.Context, tenantId string) ApiRelationshipsWriteRequest

	// RelationshipsWriteExecute executes the request
	//  @return RelationshipWriteResponse
	RelationshipsWriteExecute(r ApiRelationshipsWriteRequest) (*RelationshipWriteResponse, *http.Response, error)
}

// DataAPIService DataAPI service
type DataAPIService service

type ApiBundleRunRequest struct {
	ctx context.Context
	ApiService DataAPI
	tenantId string
	body *RunBundleBody
}

func (r ApiBundleRunRequest) Body(body RunBundleBody) ApiBundleRunRequest {
	r.body = &body
	return r
}

func (r ApiBundleRunRequest) Execute() (*BundleRunResponse, *http.Response, error) {
	return r.ApiService.BundleRunExecute(r)
}

/*
BundleRun run bundle

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tenantId Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
 @return ApiBundleRunRequest
*/
func (a *DataAPIService) BundleRun(ctx context.Context, tenantId string) ApiBundleRunRequest {
	return ApiBundleRunRequest{
		ApiService: a,
		ctx: ctx,
		tenantId: tenantId,
	}
}

// Execute executes the request
//  @return BundleRunResponse
func (a *DataAPIService) BundleRunExecute(r ApiBundleRunRequest) (*BundleRunResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BundleRunResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataAPIService.BundleRun")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/tenants/{tenant_id}/data/run-bundle"
	localVarPath = strings.Replace(localVarPath, "{"+"tenant_id"+"}", url.PathEscape(parameterValueToString(r.tenantId, "tenantId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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
			var v Status
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiDataAttributesReadRequest struct {
	ctx context.Context
	ApiService DataAPI
	tenantId string
	body *ReadAttributesBody
}

func (r ApiDataAttributesReadRequest) Body(body ReadAttributesBody) ApiDataAttributesReadRequest {
	r.body = &body
	return r
}

func (r ApiDataAttributesReadRequest) Execute() (*AttributeReadResponse, *http.Response, error) {
	return r.ApiService.DataAttributesReadExecute(r)
}

/*
DataAttributesRead read attributes

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tenantId Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
 @return ApiDataAttributesReadRequest
*/
func (a *DataAPIService) DataAttributesRead(ctx context.Context, tenantId string) ApiDataAttributesReadRequest {
	return ApiDataAttributesReadRequest{
		ApiService: a,
		ctx: ctx,
		tenantId: tenantId,
	}
}

// Execute executes the request
//  @return AttributeReadResponse
func (a *DataAPIService) DataAttributesReadExecute(r ApiDataAttributesReadRequest) (*AttributeReadResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AttributeReadResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataAPIService.DataAttributesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/tenants/{tenant_id}/data/attributes/read"
	localVarPath = strings.Replace(localVarPath, "{"+"tenant_id"+"}", url.PathEscape(parameterValueToString(r.tenantId, "tenantId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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
			var v Status
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiDataDeleteRequest struct {
	ctx context.Context
	ApiService DataAPI
	tenantId string
	body *DataDeleteBody
}

func (r ApiDataDeleteRequest) Body(body DataDeleteBody) ApiDataDeleteRequest {
	r.body = &body
	return r
}

func (r ApiDataDeleteRequest) Execute() (*DataDeleteResponse, *http.Response, error) {
	return r.ApiService.DataDeleteExecute(r)
}

/*
DataDelete delete data

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tenantId Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
 @return ApiDataDeleteRequest
*/
func (a *DataAPIService) DataDelete(ctx context.Context, tenantId string) ApiDataDeleteRequest {
	return ApiDataDeleteRequest{
		ApiService: a,
		ctx: ctx,
		tenantId: tenantId,
	}
}

// Execute executes the request
//  @return DataDeleteResponse
func (a *DataAPIService) DataDeleteExecute(r ApiDataDeleteRequest) (*DataDeleteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DataDeleteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataAPIService.DataDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/tenants/{tenant_id}/data/delete"
	localVarPath = strings.Replace(localVarPath, "{"+"tenant_id"+"}", url.PathEscape(parameterValueToString(r.tenantId, "tenantId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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
			var v Status
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiDataRelationshipsReadRequest struct {
	ctx context.Context
	ApiService DataAPI
	tenantId string
	body *ReadRelationshipsBody
}

func (r ApiDataRelationshipsReadRequest) Body(body ReadRelationshipsBody) ApiDataRelationshipsReadRequest {
	r.body = &body
	return r
}

func (r ApiDataRelationshipsReadRequest) Execute() (*RelationshipReadResponse, *http.Response, error) {
	return r.ApiService.DataRelationshipsReadExecute(r)
}

/*
DataRelationshipsRead read relationships

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tenantId Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
 @return ApiDataRelationshipsReadRequest
*/
func (a *DataAPIService) DataRelationshipsRead(ctx context.Context, tenantId string) ApiDataRelationshipsReadRequest {
	return ApiDataRelationshipsReadRequest{
		ApiService: a,
		ctx: ctx,
		tenantId: tenantId,
	}
}

// Execute executes the request
//  @return RelationshipReadResponse
func (a *DataAPIService) DataRelationshipsReadExecute(r ApiDataRelationshipsReadRequest) (*RelationshipReadResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RelationshipReadResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataAPIService.DataRelationshipsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/tenants/{tenant_id}/data/relationships/read"
	localVarPath = strings.Replace(localVarPath, "{"+"tenant_id"+"}", url.PathEscape(parameterValueToString(r.tenantId, "tenantId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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
			var v Status
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiDataWriteRequest struct {
	ctx context.Context
	ApiService DataAPI
	tenantId string
	body *DataWriteBody
}

func (r ApiDataWriteRequest) Body(body DataWriteBody) ApiDataWriteRequest {
	r.body = &body
	return r
}

func (r ApiDataWriteRequest) Execute() (*DataWriteResponse, *http.Response, error) {
	return r.ApiService.DataWriteExecute(r)
}

/*
DataWrite write data

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tenantId Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
 @return ApiDataWriteRequest
*/
func (a *DataAPIService) DataWrite(ctx context.Context, tenantId string) ApiDataWriteRequest {
	return ApiDataWriteRequest{
		ApiService: a,
		ctx: ctx,
		tenantId: tenantId,
	}
}

// Execute executes the request
//  @return DataWriteResponse
func (a *DataAPIService) DataWriteExecute(r ApiDataWriteRequest) (*DataWriteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DataWriteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataAPIService.DataWrite")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/tenants/{tenant_id}/data/write"
	localVarPath = strings.Replace(localVarPath, "{"+"tenant_id"+"}", url.PathEscape(parameterValueToString(r.tenantId, "tenantId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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
			var v Status
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiRelationshipsDeleteRequest struct {
	ctx context.Context
	ApiService DataAPI
	tenantId string
	body *DeleteRelationshipsBody
}

func (r ApiRelationshipsDeleteRequest) Body(body DeleteRelationshipsBody) ApiRelationshipsDeleteRequest {
	r.body = &body
	return r
}

func (r ApiRelationshipsDeleteRequest) Execute() (*RelationshipDeleteResponse, *http.Response, error) {
	return r.ApiService.RelationshipsDeleteExecute(r)
}

/*
RelationshipsDelete delete relationships

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tenantId Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
 @return ApiRelationshipsDeleteRequest
*/
func (a *DataAPIService) RelationshipsDelete(ctx context.Context, tenantId string) ApiRelationshipsDeleteRequest {
	return ApiRelationshipsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		tenantId: tenantId,
	}
}

// Execute executes the request
//  @return RelationshipDeleteResponse
func (a *DataAPIService) RelationshipsDeleteExecute(r ApiRelationshipsDeleteRequest) (*RelationshipDeleteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RelationshipDeleteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataAPIService.RelationshipsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/tenants/{tenant_id}/relationships/delete"
	localVarPath = strings.Replace(localVarPath, "{"+"tenant_id"+"}", url.PathEscape(parameterValueToString(r.tenantId, "tenantId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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
			var v Status
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiRelationshipsWriteRequest struct {
	ctx context.Context
	ApiService DataAPI
	tenantId string
	body *WriteRelationshipsBody
}

func (r ApiRelationshipsWriteRequest) Body(body WriteRelationshipsBody) ApiRelationshipsWriteRequest {
	r.body = &body
	return r
}

func (r ApiRelationshipsWriteRequest) Execute() (*RelationshipWriteResponse, *http.Response, error) {
	return r.ApiService.RelationshipsWriteExecute(r)
}

/*
RelationshipsWrite write relationships

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tenantId Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
 @return ApiRelationshipsWriteRequest
*/
func (a *DataAPIService) RelationshipsWrite(ctx context.Context, tenantId string) ApiRelationshipsWriteRequest {
	return ApiRelationshipsWriteRequest{
		ApiService: a,
		ctx: ctx,
		tenantId: tenantId,
	}
}

// Execute executes the request
//  @return RelationshipWriteResponse
func (a *DataAPIService) RelationshipsWriteExecute(r ApiRelationshipsWriteRequest) (*RelationshipWriteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RelationshipWriteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataAPIService.RelationshipsWrite")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/tenants/{tenant_id}/relationships/write"
	localVarPath = strings.Replace(localVarPath, "{"+"tenant_id"+"}", url.PathEscape(parameterValueToString(r.tenantId, "tenantId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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
			var v Status
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
