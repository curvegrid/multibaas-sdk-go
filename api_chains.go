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

type ChainsApi interface {

	/*
		GetBlock Get a block

		Returns a block.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param chain The blockchain chain label.
		@param block A block number, hash or 'latest' for the latest block.
		@return ApiGetBlockRequest
	*/
	GetBlock(ctx context.Context, chain ChainName, block string) ApiGetBlockRequest

	// GetBlockExecute executes the request
	//  @return GetBlock200Response
	GetBlockExecute(r ApiGetBlockRequest) (*GetBlock200Response, *http.Response, error)

	/*
		GetChainStatus Get chain status

		Returns the chain status.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param chain The blockchain chain label.
		@return ApiGetChainStatusRequest
	*/
	GetChainStatus(ctx context.Context, chain ChainName) ApiGetChainStatusRequest

	// GetChainStatusExecute executes the request
	//  @return GetChainStatus200Response
	GetChainStatusExecute(r ApiGetChainStatusRequest) (*GetChainStatus200Response, *http.Response, error)

	/*
		GetTransaction Get transaction

		Returns a transaction.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param chain The blockchain chain label.
		@param hash Transaction hash.
		@return ApiGetTransactionRequest
	*/
	GetTransaction(ctx context.Context, chain ChainName, hash string) ApiGetTransactionRequest

	// GetTransactionExecute executes the request
	//  @return GetTransaction200Response
	GetTransactionExecute(r ApiGetTransactionRequest) (*GetTransaction200Response, *http.Response, error)

	/*
		GetTransactionReceipt Get transaction receipt

		Returns the receipt of a transaction that's been successfully added to the blockchain.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param chain The blockchain chain label.
		@param hash Transaction hash.
		@return ApiGetTransactionReceiptRequest
	*/
	GetTransactionReceipt(ctx context.Context, chain ChainName, hash string) ApiGetTransactionReceiptRequest

	// GetTransactionReceiptExecute executes the request
	//  @return GetTransactionReceipt200Response
	GetTransactionReceiptExecute(r ApiGetTransactionReceiptRequest) (*GetTransactionReceipt200Response, *http.Response, error)

	/*
		SubmitSignedTransaction Submit signed transaction

		Receives a pre-signed raw transaction and submits it to the blockchain.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param chain The blockchain chain label.
		@return ApiSubmitSignedTransactionRequest
	*/
	SubmitSignedTransaction(ctx context.Context, chain ChainName) ApiSubmitSignedTransactionRequest

	// SubmitSignedTransactionExecute executes the request
	//  @return SubmitSignedTransaction200Response
	SubmitSignedTransactionExecute(r ApiSubmitSignedTransactionRequest) (*SubmitSignedTransaction200Response, *http.Response, error)

	/*
		TransferEth Transfer ETH

		Returns a transaction for sending the native token between addresses.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param chain The blockchain chain label.
		@return ApiTransferEthRequest
	*/
	TransferEth(ctx context.Context, chain ChainName) ApiTransferEthRequest

	// TransferEthExecute executes the request
	//  @return TransferEth200Response
	TransferEthExecute(r ApiTransferEthRequest) (*TransferEth200Response, *http.Response, error)
}

// ChainsApiService ChainsApi service
type ChainsApiService service

type ApiGetBlockRequest struct {
	ctx        context.Context
	ApiService ChainsApi
	chain      ChainName
	block      string
}

func (r ApiGetBlockRequest) Execute() (*GetBlock200Response, *http.Response, error) {
	return r.ApiService.GetBlockExecute(r)
}

/*
GetBlock Get a block

Returns a block.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param chain The blockchain chain label.
	@param block A block number, hash or 'latest' for the latest block.
	@return ApiGetBlockRequest
*/
func (a *ChainsApiService) GetBlock(ctx context.Context, chain ChainName, block string) ApiGetBlockRequest {
	return ApiGetBlockRequest{
		ApiService: a,
		ctx:        ctx,
		chain:      chain,
		block:      block,
	}
}

// Execute executes the request
//
//	@return GetBlock200Response
func (a *ChainsApiService) GetBlockExecute(r ApiGetBlockRequest) (*GetBlock200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetBlock200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChainsApiService.GetBlock")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chains/{chain}/blocks/{block}"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", url.PathEscape(parameterValueToString(r.chain, "chain")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"block"+"}", url.PathEscape(parameterValueToString(r.block, "block")), -1)

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

type ApiGetChainStatusRequest struct {
	ctx        context.Context
	ApiService ChainsApi
	chain      ChainName
}

func (r ApiGetChainStatusRequest) Execute() (*GetChainStatus200Response, *http.Response, error) {
	return r.ApiService.GetChainStatusExecute(r)
}

/*
GetChainStatus Get chain status

Returns the chain status.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param chain The blockchain chain label.
	@return ApiGetChainStatusRequest
*/
func (a *ChainsApiService) GetChainStatus(ctx context.Context, chain ChainName) ApiGetChainStatusRequest {
	return ApiGetChainStatusRequest{
		ApiService: a,
		ctx:        ctx,
		chain:      chain,
	}
}

// Execute executes the request
//
//	@return GetChainStatus200Response
func (a *ChainsApiService) GetChainStatusExecute(r ApiGetChainStatusRequest) (*GetChainStatus200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetChainStatus200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChainsApiService.GetChainStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chains/{chain}/status"
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

type ApiGetTransactionRequest struct {
	ctx        context.Context
	ApiService ChainsApi
	chain      ChainName
	hash       string
	include    *string
}

// Include contract and method call details, if available.
func (r ApiGetTransactionRequest) Include(include string) ApiGetTransactionRequest {
	r.include = &include
	return r
}

func (r ApiGetTransactionRequest) Execute() (*GetTransaction200Response, *http.Response, error) {
	return r.ApiService.GetTransactionExecute(r)
}

/*
GetTransaction Get transaction

Returns a transaction.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param chain The blockchain chain label.
	@param hash Transaction hash.
	@return ApiGetTransactionRequest
*/
func (a *ChainsApiService) GetTransaction(ctx context.Context, chain ChainName, hash string) ApiGetTransactionRequest {
	return ApiGetTransactionRequest{
		ApiService: a,
		ctx:        ctx,
		chain:      chain,
		hash:       hash,
	}
}

// Execute executes the request
//
//	@return GetTransaction200Response
func (a *ChainsApiService) GetTransactionExecute(r ApiGetTransactionRequest) (*GetTransaction200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetTransaction200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChainsApiService.GetTransaction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chains/{chain}/transactions/{hash}"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", url.PathEscape(parameterValueToString(r.chain, "chain")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"hash"+"}", url.PathEscape(parameterValueToString(r.hash, "hash")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "")
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

type ApiGetTransactionReceiptRequest struct {
	ctx        context.Context
	ApiService ChainsApi
	chain      ChainName
	hash       string
	include    *string
}

// Include contract and event details, if available.
func (r ApiGetTransactionReceiptRequest) Include(include string) ApiGetTransactionReceiptRequest {
	r.include = &include
	return r
}

func (r ApiGetTransactionReceiptRequest) Execute() (*GetTransactionReceipt200Response, *http.Response, error) {
	return r.ApiService.GetTransactionReceiptExecute(r)
}

/*
GetTransactionReceipt Get transaction receipt

Returns the receipt of a transaction that's been successfully added to the blockchain.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param chain The blockchain chain label.
	@param hash Transaction hash.
	@return ApiGetTransactionReceiptRequest
*/
func (a *ChainsApiService) GetTransactionReceipt(ctx context.Context, chain ChainName, hash string) ApiGetTransactionReceiptRequest {
	return ApiGetTransactionReceiptRequest{
		ApiService: a,
		ctx:        ctx,
		chain:      chain,
		hash:       hash,
	}
}

// Execute executes the request
//
//	@return GetTransactionReceipt200Response
func (a *ChainsApiService) GetTransactionReceiptExecute(r ApiGetTransactionReceiptRequest) (*GetTransactionReceipt200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetTransactionReceipt200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChainsApiService.GetTransactionReceipt")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chains/{chain}/transactions/receipt/{hash}"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", url.PathEscape(parameterValueToString(r.chain, "chain")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"hash"+"}", url.PathEscape(parameterValueToString(r.hash, "hash")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "")
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

type ApiSubmitSignedTransactionRequest struct {
	ctx                         context.Context
	ApiService                  ChainsApi
	chain                       ChainName
	signedTransactionSubmission *SignedTransactionSubmission
}

func (r ApiSubmitSignedTransactionRequest) SignedTransactionSubmission(signedTransactionSubmission SignedTransactionSubmission) ApiSubmitSignedTransactionRequest {
	r.signedTransactionSubmission = &signedTransactionSubmission
	return r
}

func (r ApiSubmitSignedTransactionRequest) Execute() (*SubmitSignedTransaction200Response, *http.Response, error) {
	return r.ApiService.SubmitSignedTransactionExecute(r)
}

/*
SubmitSignedTransaction Submit signed transaction

Receives a pre-signed raw transaction and submits it to the blockchain.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param chain The blockchain chain label.
	@return ApiSubmitSignedTransactionRequest
*/
func (a *ChainsApiService) SubmitSignedTransaction(ctx context.Context, chain ChainName) ApiSubmitSignedTransactionRequest {
	return ApiSubmitSignedTransactionRequest{
		ApiService: a,
		ctx:        ctx,
		chain:      chain,
	}
}

// Execute executes the request
//
//	@return SubmitSignedTransaction200Response
func (a *ChainsApiService) SubmitSignedTransactionExecute(r ApiSubmitSignedTransactionRequest) (*SubmitSignedTransaction200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SubmitSignedTransaction200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChainsApiService.SubmitSignedTransaction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chains/{chain}/transactions/submit"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", url.PathEscape(parameterValueToString(r.chain, "chain")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.signedTransactionSubmission
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

type ApiTransferEthRequest struct {
	ctx            context.Context
	ApiService     ChainsApi
	chain          ChainName
	postMethodArgs *PostMethodArgs
}

func (r ApiTransferEthRequest) PostMethodArgs(postMethodArgs PostMethodArgs) ApiTransferEthRequest {
	r.postMethodArgs = &postMethodArgs
	return r
}

func (r ApiTransferEthRequest) Execute() (*TransferEth200Response, *http.Response, error) {
	return r.ApiService.TransferEthExecute(r)
}

/*
TransferEth Transfer ETH

Returns a transaction for sending the native token between addresses.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param chain The blockchain chain label.
	@return ApiTransferEthRequest
*/
func (a *ChainsApiService) TransferEth(ctx context.Context, chain ChainName) ApiTransferEthRequest {
	return ApiTransferEthRequest{
		ApiService: a,
		ctx:        ctx,
		chain:      chain,
	}
}

// Execute executes the request
//
//	@return TransferEth200Response
func (a *ChainsApiService) TransferEthExecute(r ApiTransferEthRequest) (*TransferEth200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TransferEth200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChainsApiService.TransferEth")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chains/{chain}/transfers"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", url.PathEscape(parameterValueToString(r.chain, "chain")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.postMethodArgs
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
