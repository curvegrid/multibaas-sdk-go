/*
MultiBaas API

MultiBaas's REST APIv0.

API version: 0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package multibaas

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type WebhooksAPI interface {

	/*
		CountWebhookEvents Count webhook events

		Count the events for the given webhook endpoint.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param webhookID
		@return ApiCountWebhookEventsRequest
	*/
	CountWebhookEvents(ctx context.Context, webhookID int64) ApiCountWebhookEventsRequest

	// CountWebhookEventsExecute executes the request
	//  @return CountWebhookEvents200Response
	CountWebhookEventsExecute(r ApiCountWebhookEventsRequest) (*CountWebhookEvents200Response, *http.Response, error)

	/*
		CountWebhooks Count webhooks

		Count all webhook endpoints.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiCountWebhooksRequest
	*/
	CountWebhooks(ctx context.Context) ApiCountWebhooksRequest

	// CountWebhooksExecute executes the request
	//  @return CountWebhooks200Response
	CountWebhooksExecute(r ApiCountWebhooksRequest) (*CountWebhooks200Response, *http.Response, error)

	/*
		CreateWebhook Create webhook

		Create a webhook.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiCreateWebhookRequest
	*/
	CreateWebhook(ctx context.Context) ApiCreateWebhookRequest

	// CreateWebhookExecute executes the request
	//  @return CreateWebhook200Response
	CreateWebhookExecute(r ApiCreateWebhookRequest) (*CreateWebhook200Response, *http.Response, error)

	/*
		DeleteWebhook Delete webhook

		Delete a webhook endpoint.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param webhookID
		@return ApiDeleteWebhookRequest
	*/
	DeleteWebhook(ctx context.Context, webhookID int64) ApiDeleteWebhookRequest

	// DeleteWebhookExecute executes the request
	//  @return BaseResponse
	DeleteWebhookExecute(r ApiDeleteWebhookRequest) (*BaseResponse, *http.Response, error)

	/*
		GetWebhook Get webhook

		Get a webhook endpoint.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param webhookID
		@return ApiGetWebhookRequest
	*/
	GetWebhook(ctx context.Context, webhookID int64) ApiGetWebhookRequest

	// GetWebhookExecute executes the request
	//  @return CreateWebhook200Response
	GetWebhookExecute(r ApiGetWebhookRequest) (*CreateWebhook200Response, *http.Response, error)

	/*
		ListWebhookEvents List webhook events

		List events for the given webhook endpoint.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param webhookID
		@return ApiListWebhookEventsRequest
	*/
	ListWebhookEvents(ctx context.Context, webhookID int64) ApiListWebhookEventsRequest

	// ListWebhookEventsExecute executes the request
	//  @return ListWebhookEvents200Response
	ListWebhookEventsExecute(r ApiListWebhookEventsRequest) (*ListWebhookEvents200Response, *http.Response, error)

	/*
		ListWebhooks List webhooks

		List all webhook endpoints.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiListWebhooksRequest
	*/
	ListWebhooks(ctx context.Context) ApiListWebhooksRequest

	// ListWebhooksExecute executes the request
	//  @return ListWebhooks200Response
	ListWebhooksExecute(r ApiListWebhooksRequest) (*ListWebhooks200Response, *http.Response, error)

	/*
		UpdateWebhook Update webhook

		Update a webhook endpoint.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param webhookID
		@return ApiUpdateWebhookRequest
	*/
	UpdateWebhook(ctx context.Context, webhookID int64) ApiUpdateWebhookRequest

	// UpdateWebhookExecute executes the request
	//  @return CreateWebhook200Response
	UpdateWebhookExecute(r ApiUpdateWebhookRequest) (*CreateWebhook200Response, *http.Response, error)
}

// WebhooksAPIService WebhooksAPI service
type WebhooksAPIService service

type ApiCountWebhookEventsRequest struct {
	ctx        context.Context
	ApiService WebhooksAPI
	webhookID  int64
}

func (r ApiCountWebhookEventsRequest) Execute() (*CountWebhookEvents200Response, *http.Response, error) {
	return r.ApiService.CountWebhookEventsExecute(r)
}

/*
CountWebhookEvents Count webhook events

Count the events for the given webhook endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param webhookID
	@return ApiCountWebhookEventsRequest
*/
func (a *WebhooksAPIService) CountWebhookEvents(ctx context.Context, webhookID int64) ApiCountWebhookEventsRequest {
	return ApiCountWebhookEventsRequest{
		ApiService: a,
		ctx:        ctx,
		webhookID:  webhookID,
	}
}

// Execute executes the request
//
//	@return CountWebhookEvents200Response
func (a *WebhooksAPIService) CountWebhookEventsExecute(r ApiCountWebhookEventsRequest) (*CountWebhookEvents200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CountWebhookEvents200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.CountWebhookEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/{webhookID}/events/count"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookID"+"}", url.PathEscape(parameterValueToString(r.webhookID, "webhookID")), -1)

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

type ApiCountWebhooksRequest struct {
	ctx        context.Context
	ApiService WebhooksAPI
}

func (r ApiCountWebhooksRequest) Execute() (*CountWebhooks200Response, *http.Response, error) {
	return r.ApiService.CountWebhooksExecute(r)
}

/*
CountWebhooks Count webhooks

Count all webhook endpoints.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCountWebhooksRequest
*/
func (a *WebhooksAPIService) CountWebhooks(ctx context.Context) ApiCountWebhooksRequest {
	return ApiCountWebhooksRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CountWebhooks200Response
func (a *WebhooksAPIService) CountWebhooksExecute(r ApiCountWebhooksRequest) (*CountWebhooks200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CountWebhooks200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.CountWebhooks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/count"

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
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiCreateWebhookRequest struct {
	ctx                 context.Context
	ApiService          WebhooksAPI
	baseWebhookEndpoint *BaseWebhookEndpoint
}

func (r ApiCreateWebhookRequest) BaseWebhookEndpoint(baseWebhookEndpoint BaseWebhookEndpoint) ApiCreateWebhookRequest {
	r.baseWebhookEndpoint = &baseWebhookEndpoint
	return r
}

func (r ApiCreateWebhookRequest) Execute() (*CreateWebhook200Response, *http.Response, error) {
	return r.ApiService.CreateWebhookExecute(r)
}

/*
CreateWebhook Create webhook

Create a webhook.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateWebhookRequest
*/
func (a *WebhooksAPIService) CreateWebhook(ctx context.Context) ApiCreateWebhookRequest {
	return ApiCreateWebhookRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateWebhook200Response
func (a *WebhooksAPIService) CreateWebhookExecute(r ApiCreateWebhookRequest) (*CreateWebhook200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateWebhook200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.CreateWebhook")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.baseWebhookEndpoint == nil {
		return localVarReturnValue, nil, reportError("baseWebhookEndpoint is required and must be specified")
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
	localVarPostBody = r.baseWebhookEndpoint
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
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiDeleteWebhookRequest struct {
	ctx        context.Context
	ApiService WebhooksAPI
	webhookID  int64
}

func (r ApiDeleteWebhookRequest) Execute() (*BaseResponse, *http.Response, error) {
	return r.ApiService.DeleteWebhookExecute(r)
}

/*
DeleteWebhook Delete webhook

Delete a webhook endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param webhookID
	@return ApiDeleteWebhookRequest
*/
func (a *WebhooksAPIService) DeleteWebhook(ctx context.Context, webhookID int64) ApiDeleteWebhookRequest {
	return ApiDeleteWebhookRequest{
		ApiService: a,
		ctx:        ctx,
		webhookID:  webhookID,
	}
}

// Execute executes the request
//
//	@return BaseResponse
func (a *WebhooksAPIService) DeleteWebhookExecute(r ApiDeleteWebhookRequest) (*BaseResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BaseResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.DeleteWebhook")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/{webhookID}"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookID"+"}", url.PathEscape(parameterValueToString(r.webhookID, "webhookID")), -1)

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
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiGetWebhookRequest struct {
	ctx        context.Context
	ApiService WebhooksAPI
	webhookID  int64
}

func (r ApiGetWebhookRequest) Execute() (*CreateWebhook200Response, *http.Response, error) {
	return r.ApiService.GetWebhookExecute(r)
}

/*
GetWebhook Get webhook

Get a webhook endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param webhookID
	@return ApiGetWebhookRequest
*/
func (a *WebhooksAPIService) GetWebhook(ctx context.Context, webhookID int64) ApiGetWebhookRequest {
	return ApiGetWebhookRequest{
		ApiService: a,
		ctx:        ctx,
		webhookID:  webhookID,
	}
}

// Execute executes the request
//
//	@return CreateWebhook200Response
func (a *WebhooksAPIService) GetWebhookExecute(r ApiGetWebhookRequest) (*CreateWebhook200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateWebhook200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.GetWebhook")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/{webhookID}"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookID"+"}", url.PathEscape(parameterValueToString(r.webhookID, "webhookID")), -1)

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
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiListWebhookEventsRequest struct {
	ctx        context.Context
	ApiService WebhooksAPI
	webhookID  int64
	limit      *int64
	offset     *int64
}

func (r ApiListWebhookEventsRequest) Limit(limit int64) ApiListWebhookEventsRequest {
	r.limit = &limit
	return r
}

func (r ApiListWebhookEventsRequest) Offset(offset int64) ApiListWebhookEventsRequest {
	r.offset = &offset
	return r
}

func (r ApiListWebhookEventsRequest) Execute() (*ListWebhookEvents200Response, *http.Response, error) {
	return r.ApiService.ListWebhookEventsExecute(r)
}

/*
ListWebhookEvents List webhook events

List events for the given webhook endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param webhookID
	@return ApiListWebhookEventsRequest
*/
func (a *WebhooksAPIService) ListWebhookEvents(ctx context.Context, webhookID int64) ApiListWebhookEventsRequest {
	return ApiListWebhookEventsRequest{
		ApiService: a,
		ctx:        ctx,
		webhookID:  webhookID,
	}
}

// Execute executes the request
//
//	@return ListWebhookEvents200Response
func (a *WebhooksAPIService) ListWebhookEventsExecute(r ApiListWebhookEventsRequest) (*ListWebhookEvents200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListWebhookEvents200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.ListWebhookEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/{webhookID}/events"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookID"+"}", url.PathEscape(parameterValueToString(r.webhookID, "webhookID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
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

type ApiListWebhooksRequest struct {
	ctx        context.Context
	ApiService WebhooksAPI
	limit      *int64
	offset     *int64
}

func (r ApiListWebhooksRequest) Limit(limit int64) ApiListWebhooksRequest {
	r.limit = &limit
	return r
}

func (r ApiListWebhooksRequest) Offset(offset int64) ApiListWebhooksRequest {
	r.offset = &offset
	return r
}

func (r ApiListWebhooksRequest) Execute() (*ListWebhooks200Response, *http.Response, error) {
	return r.ApiService.ListWebhooksExecute(r)
}

/*
ListWebhooks List webhooks

List all webhook endpoints.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListWebhooksRequest
*/
func (a *WebhooksAPIService) ListWebhooks(ctx context.Context) ApiListWebhooksRequest {
	return ApiListWebhooksRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListWebhooks200Response
func (a *WebhooksAPIService) ListWebhooksExecute(r ApiListWebhooksRequest) (*ListWebhooks200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListWebhooks200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.ListWebhooks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
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
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiUpdateWebhookRequest struct {
	ctx                 context.Context
	ApiService          WebhooksAPI
	webhookID           int64
	baseWebhookEndpoint *BaseWebhookEndpoint
}

func (r ApiUpdateWebhookRequest) BaseWebhookEndpoint(baseWebhookEndpoint BaseWebhookEndpoint) ApiUpdateWebhookRequest {
	r.baseWebhookEndpoint = &baseWebhookEndpoint
	return r
}

func (r ApiUpdateWebhookRequest) Execute() (*CreateWebhook200Response, *http.Response, error) {
	return r.ApiService.UpdateWebhookExecute(r)
}

/*
UpdateWebhook Update webhook

Update a webhook endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param webhookID
	@return ApiUpdateWebhookRequest
*/
func (a *WebhooksAPIService) UpdateWebhook(ctx context.Context, webhookID int64) ApiUpdateWebhookRequest {
	return ApiUpdateWebhookRequest{
		ApiService: a,
		ctx:        ctx,
		webhookID:  webhookID,
	}
}

// Execute executes the request
//
//	@return CreateWebhook200Response
func (a *WebhooksAPIService) UpdateWebhookExecute(r ApiUpdateWebhookRequest) (*CreateWebhook200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateWebhook200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.UpdateWebhook")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/{webhookID}"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookID"+"}", url.PathEscape(parameterValueToString(r.webhookID, "webhookID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.baseWebhookEndpoint == nil {
		return localVarReturnValue, nil, reportError("baseWebhookEndpoint is required and must be specified")
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
	localVarPostBody = r.baseWebhookEndpoint
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
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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
