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

type TxmAPI interface {

	/*
		CancelTransaction Cancel transaction

		Cancels a transaction by resubmitting it as no-op transaction and with a higher gas price.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param chain The blockchain chain label.
		@param walletAddress An Ethereum address.
		@param nonce Transaction nonce.
		@return ApiCancelTransactionRequest
	*/
	CancelTransaction(ctx context.Context, chain ChainName, walletAddress string, nonce int64) ApiCancelTransactionRequest

	// CancelTransactionExecute executes the request
	//  @return TransferEth200Response
	CancelTransactionExecute(r ApiCancelTransactionRequest) (*TransferEth200Response, *http.Response, error)

	/*
		CountWalletTransactions Count all transactions for a wallet

		Count all transactions for the given wallet address.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param chain The blockchain chain label.
		@param walletAddress An Ethereum address.
		@return ApiCountWalletTransactionsRequest
	*/
	CountWalletTransactions(ctx context.Context, chain ChainName, walletAddress string) ApiCountWalletTransactionsRequest

	// CountWalletTransactionsExecute executes the request
	//  @return CountWalletTransactions200Response
	CountWalletTransactionsExecute(r ApiCountWalletTransactionsRequest) (*CountWalletTransactions200Response, *http.Response, error)

	/*
		ListWalletTransactions List transactions for a wallet

		List the transactions submitted by the given wallet address.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param chain The blockchain chain label.
		@param walletAddress An Ethereum address.
		@return ApiListWalletTransactionsRequest
	*/
	ListWalletTransactions(ctx context.Context, chain ChainName, walletAddress string) ApiListWalletTransactionsRequest

	// ListWalletTransactionsExecute executes the request
	//  @return ListWalletTransactions200Response
	ListWalletTransactionsExecute(r ApiListWalletTransactionsRequest) (*ListWalletTransactions200Response, *http.Response, error)

	/*
		SpeedUpTransaction Speed up transaction

		Speeds up a transaction by resubmitting it with a higher gas price.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param chain The blockchain chain label.
		@param walletAddress An Ethereum address.
		@param nonce Transaction nonce.
		@return ApiSpeedUpTransactionRequest
	*/
	SpeedUpTransaction(ctx context.Context, chain ChainName, walletAddress string, nonce int64) ApiSpeedUpTransactionRequest

	// SpeedUpTransactionExecute executes the request
	//  @return TransferEth200Response
	SpeedUpTransactionExecute(r ApiSpeedUpTransactionRequest) (*TransferEth200Response, *http.Response, error)
}

// TxmAPIService TxmAPI service
type TxmAPIService service

type ApiCancelTransactionRequest struct {
	ctx           context.Context
	ApiService    TxmAPI
	chain         ChainName
	walletAddress string
	nonce         int64
	gasParams     *GasParams
}

func (r ApiCancelTransactionRequest) GasParams(gasParams GasParams) ApiCancelTransactionRequest {
	r.gasParams = &gasParams
	return r
}

func (r ApiCancelTransactionRequest) Execute() (*TransferEth200Response, *http.Response, error) {
	return r.ApiService.CancelTransactionExecute(r)
}

/*
CancelTransaction Cancel transaction

Cancels a transaction by resubmitting it as no-op transaction and with a higher gas price.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param chain The blockchain chain label.
	@param walletAddress An Ethereum address.
	@param nonce Transaction nonce.
	@return ApiCancelTransactionRequest
*/
func (a *TxmAPIService) CancelTransaction(ctx context.Context, chain ChainName, walletAddress string, nonce int64) ApiCancelTransactionRequest {
	return ApiCancelTransactionRequest{
		ApiService:    a,
		ctx:           ctx,
		chain:         chain,
		walletAddress: walletAddress,
		nonce:         nonce,
	}
}

// Execute executes the request
//
//	@return TransferEth200Response
func (a *TxmAPIService) CancelTransactionExecute(r ApiCancelTransactionRequest) (*TransferEth200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TransferEth200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TxmAPIService.CancelTransaction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chains/{chain}/txm/{wallet_address}/nonce/{nonce}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", url.PathEscape(parameterValueToString(r.chain, "chain")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_address"+"}", url.PathEscape(parameterValueToString(r.walletAddress, "walletAddress")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nonce"+"}", url.PathEscape(parameterValueToString(r.nonce, "nonce")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.gasParams == nil {
		return localVarReturnValue, nil, reportError("gasParams is required and must be specified")
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
	localVarPostBody = r.gasParams
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

type ApiCountWalletTransactionsRequest struct {
	ctx           context.Context
	ApiService    TxmAPI
	chain         ChainName
	walletAddress string
}

func (r ApiCountWalletTransactionsRequest) Execute() (*CountWalletTransactions200Response, *http.Response, error) {
	return r.ApiService.CountWalletTransactionsExecute(r)
}

/*
CountWalletTransactions Count all transactions for a wallet

Count all transactions for the given wallet address.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param chain The blockchain chain label.
	@param walletAddress An Ethereum address.
	@return ApiCountWalletTransactionsRequest
*/
func (a *TxmAPIService) CountWalletTransactions(ctx context.Context, chain ChainName, walletAddress string) ApiCountWalletTransactionsRequest {
	return ApiCountWalletTransactionsRequest{
		ApiService:    a,
		ctx:           ctx,
		chain:         chain,
		walletAddress: walletAddress,
	}
}

// Execute executes the request
//
//	@return CountWalletTransactions200Response
func (a *TxmAPIService) CountWalletTransactionsExecute(r ApiCountWalletTransactionsRequest) (*CountWalletTransactions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CountWalletTransactions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TxmAPIService.CountWalletTransactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chains/{chain}/txm/{wallet_address}/count"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", url.PathEscape(parameterValueToString(r.chain, "chain")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_address"+"}", url.PathEscape(parameterValueToString(r.walletAddress, "walletAddress")), -1)

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

type ApiListWalletTransactionsRequest struct {
	ctx           context.Context
	ApiService    TxmAPI
	chain         ChainName
	walletAddress string
	hash          *string
	nonce         *int64
	status        *TransactionStatus
	limit         *int64
	offset        *int64
}

// Filter transactions by transaction hash. To filter for multiple hashes, use ampersands: &#x60;?hash&#x3D;HASH1&amp;hash&#x3D;HASH2&amp;hash&#x3D;HASH3&#x60;
func (r ApiListWalletTransactionsRequest) Hash(hash string) ApiListWalletTransactionsRequest {
	r.hash = &hash
	return r
}

// Filter transactions by nonce
func (r ApiListWalletTransactionsRequest) Nonce(nonce int64) ApiListWalletTransactionsRequest {
	r.nonce = &nonce
	return r
}

// Filter transactions by status
func (r ApiListWalletTransactionsRequest) Status(status TransactionStatus) ApiListWalletTransactionsRequest {
	r.status = &status
	return r
}

func (r ApiListWalletTransactionsRequest) Limit(limit int64) ApiListWalletTransactionsRequest {
	r.limit = &limit
	return r
}

func (r ApiListWalletTransactionsRequest) Offset(offset int64) ApiListWalletTransactionsRequest {
	r.offset = &offset
	return r
}

func (r ApiListWalletTransactionsRequest) Execute() (*ListWalletTransactions200Response, *http.Response, error) {
	return r.ApiService.ListWalletTransactionsExecute(r)
}

/*
ListWalletTransactions List transactions for a wallet

List the transactions submitted by the given wallet address.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param chain The blockchain chain label.
	@param walletAddress An Ethereum address.
	@return ApiListWalletTransactionsRequest
*/
func (a *TxmAPIService) ListWalletTransactions(ctx context.Context, chain ChainName, walletAddress string) ApiListWalletTransactionsRequest {
	return ApiListWalletTransactionsRequest{
		ApiService:    a,
		ctx:           ctx,
		chain:         chain,
		walletAddress: walletAddress,
	}
}

// Execute executes the request
//
//	@return ListWalletTransactions200Response
func (a *TxmAPIService) ListWalletTransactionsExecute(r ApiListWalletTransactionsRequest) (*ListWalletTransactions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListWalletTransactions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TxmAPIService.ListWalletTransactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chains/{chain}/txm/{wallet_address}"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", url.PathEscape(parameterValueToString(r.chain, "chain")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_address"+"}", url.PathEscape(parameterValueToString(r.walletAddress, "walletAddress")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.hash != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hash", r.hash, "")
	}
	if r.nonce != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nonce", r.nonce, "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int64 = 10
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
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

type ApiSpeedUpTransactionRequest struct {
	ctx           context.Context
	ApiService    TxmAPI
	chain         ChainName
	walletAddress string
	nonce         int64
	gasParams     *GasParams
}

func (r ApiSpeedUpTransactionRequest) GasParams(gasParams GasParams) ApiSpeedUpTransactionRequest {
	r.gasParams = &gasParams
	return r
}

func (r ApiSpeedUpTransactionRequest) Execute() (*TransferEth200Response, *http.Response, error) {
	return r.ApiService.SpeedUpTransactionExecute(r)
}

/*
SpeedUpTransaction Speed up transaction

Speeds up a transaction by resubmitting it with a higher gas price.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param chain The blockchain chain label.
	@param walletAddress An Ethereum address.
	@param nonce Transaction nonce.
	@return ApiSpeedUpTransactionRequest
*/
func (a *TxmAPIService) SpeedUpTransaction(ctx context.Context, chain ChainName, walletAddress string, nonce int64) ApiSpeedUpTransactionRequest {
	return ApiSpeedUpTransactionRequest{
		ApiService:    a,
		ctx:           ctx,
		chain:         chain,
		walletAddress: walletAddress,
		nonce:         nonce,
	}
}

// Execute executes the request
//
//	@return TransferEth200Response
func (a *TxmAPIService) SpeedUpTransactionExecute(r ApiSpeedUpTransactionRequest) (*TransferEth200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TransferEth200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TxmAPIService.SpeedUpTransaction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chains/{chain}/txm/{wallet_address}/nonce/{nonce}/speed_up"
	localVarPath = strings.Replace(localVarPath, "{"+"chain"+"}", url.PathEscape(parameterValueToString(r.chain, "chain")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_address"+"}", url.PathEscape(parameterValueToString(r.walletAddress, "walletAddress")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nonce"+"}", url.PathEscape(parameterValueToString(r.nonce, "nonce")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.gasParams == nil {
		return localVarReturnValue, nil, reportError("gasParams is required and must be specified")
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
	localVarPostBody = r.gasParams
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
