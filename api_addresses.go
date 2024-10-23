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
	"reflect"
	"strings"
)

type AddressesAPI interface {

	/*
		DeleteAddress Delete address

		Deletes an address label.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param chain The blockchain chain label.
		@param addressOrLabel An address or the label of an address.
		@return ApiDeleteAddressRequest
	*/
	DeleteAddress(ctx context.Context, chain ChainName, addressOrLabel string) ApiDeleteAddressRequest

	// DeleteAddressExecute executes the request
	//  @return BaseResponse
	DeleteAddressExecute(r ApiDeleteAddressRequest) (*BaseResponse, *http.Response, error)

	/*
		GetAddress Get address

		Returns details about an address.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param chain The blockchain chain label.
		@param addressOrLabel An address or the label of an address.
		@return ApiGetAddressRequest
	*/
	GetAddress(ctx context.Context, chain ChainName, addressOrLabel string) ApiGetAddressRequest

	// GetAddressExecute executes the request
	//  @return SetAddress201Response
	GetAddressExecute(r ApiGetAddressRequest) (*SetAddress201Response, *http.Response, error)

	/*
		ListAddresses List addresses

		Returns all the labeled addresses.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param chain The blockchain chain label.
		@return ApiListAddressesRequest
	*/
	ListAddresses(ctx context.Context, chain ChainName) ApiListAddressesRequest

	// ListAddressesExecute executes the request
	//  @return ListAddresses200Response
	ListAddressesExecute(r ApiListAddressesRequest) (*ListAddresses200Response, *http.Response, error)

	/*
		SetAddress Create or update address

		Associates an address with a label.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param chain The blockchain chain label.
		@return ApiSetAddressRequest
	*/
	SetAddress(ctx context.Context, chain ChainName) ApiSetAddressRequest

	// SetAddressExecute executes the request
	//  @return SetAddress201Response
	SetAddressExecute(r ApiSetAddressRequest) (*SetAddress201Response, *http.Response, error)
}

// AddressesAPIService AddressesAPI service
type AddressesAPIService service

type ApiDeleteAddressRequest struct {
	ctx            context.Context
	ApiService     AddressesAPI
	chain          ChainName
	addressOrLabel string
}

func (r ApiDeleteAddressRequest) Execute() (*BaseResponse, *http.Response, error) {
	return r.ApiService.DeleteAddressExecute(r)
}

/*
DeleteAddress Delete address

Deletes an address label.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param chain The blockchain chain label.
	@param addressOrLabel An address or the label of an address.
	@return ApiDeleteAddressRequest
*/
func (a *AddressesAPIService) DeleteAddress(ctx context.Context, chain ChainName, addressOrLabel string) ApiDeleteAddressRequest {
	return ApiDeleteAddressRequest{
		ApiService:     a,
		ctx:            ctx,
		chain:          chain,
		addressOrLabel: addressOrLabel,
	}
}

// Execute executes the request
//
//	@return BaseResponse
func (a *AddressesAPIService) DeleteAddressExecute(r ApiDeleteAddressRequest) (*BaseResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BaseResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressesAPIService.DeleteAddress")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chains/{chain}/addresses/{address-or-label}"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", url.PathEscape(parameterValueToString(r.chain, "chain")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address-or-label"+"}", url.PathEscape(parameterValueToString(r.addressOrLabel, "addressOrLabel")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.addressOrLabel) < 1 {
		return localVarReturnValue, nil, reportError("addressOrLabel must have at least 1 elements")
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

type ApiGetAddressRequest struct {
	ctx            context.Context
	ApiService     AddressesAPI
	chain          ChainName
	addressOrLabel string
	include        *[]string
}

// Optional data to fetch from the blockchain: - &#x60;balance&#x60; to get the balance of this address. - &#x60;code&#x60; to get the code at this address. - &#x60;nonce&#x60; to get the next available transaction nonce for this address. - &#x60;contractLookup&#x60; to get the contract(s) details for this address.
func (r ApiGetAddressRequest) Include(include []string) ApiGetAddressRequest {
	r.include = &include
	return r
}

func (r ApiGetAddressRequest) Execute() (*SetAddress201Response, *http.Response, error) {
	return r.ApiService.GetAddressExecute(r)
}

/*
GetAddress Get address

Returns details about an address.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param chain The blockchain chain label.
	@param addressOrLabel An address or the label of an address.
	@return ApiGetAddressRequest
*/
func (a *AddressesAPIService) GetAddress(ctx context.Context, chain ChainName, addressOrLabel string) ApiGetAddressRequest {
	return ApiGetAddressRequest{
		ApiService:     a,
		ctx:            ctx,
		chain:          chain,
		addressOrLabel: addressOrLabel,
	}
}

// Execute executes the request
//
//	@return SetAddress201Response
func (a *AddressesAPIService) GetAddressExecute(r ApiGetAddressRequest) (*SetAddress201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SetAddress201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressesAPIService.GetAddress")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chains/{chain}/addresses/{address-or-label}"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", url.PathEscape(parameterValueToString(r.chain, "chain")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address-or-label"+"}", url.PathEscape(parameterValueToString(r.addressOrLabel, "addressOrLabel")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.addressOrLabel) < 1 {
		return localVarReturnValue, nil, reportError("addressOrLabel must have at least 1 elements")
	}

	if r.include != nil {
		t := *r.include
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "include", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "include", t, "form", "multi")
		}
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

type ApiListAddressesRequest struct {
	ctx        context.Context
	ApiService AddressesAPI
	chain      ChainName
}

func (r ApiListAddressesRequest) Execute() (*ListAddresses200Response, *http.Response, error) {
	return r.ApiService.ListAddressesExecute(r)
}

/*
ListAddresses List addresses

Returns all the labeled addresses.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param chain The blockchain chain label.
	@return ApiListAddressesRequest
*/
func (a *AddressesAPIService) ListAddresses(ctx context.Context, chain ChainName) ApiListAddressesRequest {
	return ApiListAddressesRequest{
		ApiService: a,
		ctx:        ctx,
		chain:      chain,
	}
}

// Execute executes the request
//
//	@return ListAddresses200Response
func (a *AddressesAPIService) ListAddressesExecute(r ApiListAddressesRequest) (*ListAddresses200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListAddresses200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressesAPIService.ListAddresses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chains/{chain}/addresses"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", url.PathEscape(parameterValueToString(r.chain, "chain")), -1)

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

type ApiSetAddressRequest struct {
	ctx          context.Context
	ApiService   AddressesAPI
	chain        ChainName
	addressLabel *AddressLabel
}

func (r ApiSetAddressRequest) AddressLabel(addressLabel AddressLabel) ApiSetAddressRequest {
	r.addressLabel = &addressLabel
	return r
}

func (r ApiSetAddressRequest) Execute() (*SetAddress201Response, *http.Response, error) {
	return r.ApiService.SetAddressExecute(r)
}

/*
SetAddress Create or update address

Associates an address with a label.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param chain The blockchain chain label.
	@return ApiSetAddressRequest
*/
func (a *AddressesAPIService) SetAddress(ctx context.Context, chain ChainName) ApiSetAddressRequest {
	return ApiSetAddressRequest{
		ApiService: a,
		ctx:        ctx,
		chain:      chain,
	}
}

// Execute executes the request
//
//	@return SetAddress201Response
func (a *AddressesAPIService) SetAddressExecute(r ApiSetAddressRequest) (*SetAddress201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SetAddress201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressesAPIService.SetAddress")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chains/{chain}/addresses"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", url.PathEscape(parameterValueToString(r.chain, "chain")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.addressLabel == nil {
		return localVarReturnValue, nil, reportError("addressLabel is required and must be specified")
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
	localVarPostBody = r.addressLabel
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
