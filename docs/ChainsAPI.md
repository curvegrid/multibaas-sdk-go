# \ChainsAPI

All URIs are relative to *https://your_deployment.multibaas.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlock**](ChainsAPI.md#GetBlock) | **Get** /chains/{chain}/blocks/{block} | Get a block
[**GetChainStatus**](ChainsAPI.md#GetChainStatus) | **Get** /chains/{chain}/status | Get chain status
[**GetTransaction**](ChainsAPI.md#GetTransaction) | **Get** /chains/{chain}/transactions/{hash} | Get transaction
[**GetTransactionReceipt**](ChainsAPI.md#GetTransactionReceipt) | **Get** /chains/{chain}/transactions/receipt/{hash} | Get transaction receipt
[**SubmitSignedTransaction**](ChainsAPI.md#SubmitSignedTransaction) | **Post** /chains/{chain}/transactions/submit | Submit signed transaction
[**TransferEth**](ChainsAPI.md#TransferEth) | **Post** /chains/{chain}/transfers | Transfer ETH



## GetBlock

> GetBlock200Response GetBlock(ctx, chain, block).Execute()

Get a block



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/curvegrid/multibaas-sdk-go"
)

func main() {
	chain := openapiclient.ChainName("ethereum") // ChainName | The blockchain chain label.
	block := "block_example" // string | A block number, hash or 'latest' for the latest block.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChainsAPI.GetBlock(context.Background(), chain, block).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChainsAPI.GetBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlock`: GetBlock200Response
	fmt.Fprintf(os.Stdout, "Response from `ChainsAPI.GetBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 
**block** | **string** | A block number, hash or &#39;latest&#39; for the latest block. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetBlock200Response**](GetBlock200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChainStatus

> GetChainStatus200Response GetChainStatus(ctx, chain).Execute()

Get chain status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/curvegrid/multibaas-sdk-go"
)

func main() {
	chain := openapiclient.ChainName("ethereum") // ChainName | The blockchain chain label.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChainsAPI.GetChainStatus(context.Background(), chain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChainsAPI.GetChainStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChainStatus`: GetChainStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `ChainsAPI.GetChainStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChainStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetChainStatus200Response**](GetChainStatus200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransaction

> GetTransaction200Response GetTransaction(ctx, chain, hash).Include(include).Execute()

Get transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/curvegrid/multibaas-sdk-go"
)

func main() {
	chain := openapiclient.ChainName("ethereum") // ChainName | The blockchain chain label.
	hash := "hash_example" // string | Transaction hash.
	include := "include_example" // string | Include contract and method call details, if available. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChainsAPI.GetTransaction(context.Background(), chain, hash).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChainsAPI.GetTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransaction`: GetTransaction200Response
	fmt.Fprintf(os.Stdout, "Response from `ChainsAPI.GetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 
**hash** | **string** | Transaction hash. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | **string** | Include contract and method call details, if available. | 

### Return type

[**GetTransaction200Response**](GetTransaction200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionReceipt

> GetTransactionReceipt200Response GetTransactionReceipt(ctx, chain, hash).Include(include).Execute()

Get transaction receipt



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/curvegrid/multibaas-sdk-go"
)

func main() {
	chain := openapiclient.ChainName("ethereum") // ChainName | The blockchain chain label.
	hash := "hash_example" // string | Transaction hash.
	include := "include_example" // string | Include contract and event details, if available. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChainsAPI.GetTransactionReceipt(context.Background(), chain, hash).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChainsAPI.GetTransactionReceipt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionReceipt`: GetTransactionReceipt200Response
	fmt.Fprintf(os.Stdout, "Response from `ChainsAPI.GetTransactionReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 
**hash** | **string** | Transaction hash. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | **string** | Include contract and event details, if available. | 

### Return type

[**GetTransactionReceipt200Response**](GetTransactionReceipt200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitSignedTransaction

> BaseResponse SubmitSignedTransaction(ctx, chain).SignedTransactionSubmission(signedTransactionSubmission).Execute()

Submit signed transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/curvegrid/multibaas-sdk-go"
)

func main() {
	chain := openapiclient.ChainName("ethereum") // ChainName | The blockchain chain label.
	signedTransactionSubmission := *openapiclient.NewSignedTransactionSubmission("SignedTx_example") // SignedTransactionSubmission | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChainsAPI.SubmitSignedTransaction(context.Background(), chain).SignedTransactionSubmission(signedTransactionSubmission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChainsAPI.SubmitSignedTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitSignedTransaction`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ChainsAPI.SubmitSignedTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitSignedTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signedTransactionSubmission** | [**SignedTransactionSubmission**](SignedTransactionSubmission.md) |  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferEth

> TransferEth200Response TransferEth(ctx, chain).PostMethodArgs(postMethodArgs).Execute()

Transfer ETH



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/curvegrid/multibaas-sdk-go"
)

func main() {
	chain := openapiclient.ChainName("ethereum") // ChainName | The blockchain chain label.
	postMethodArgs := *openapiclient.NewPostMethodArgs() // PostMethodArgs | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChainsAPI.TransferEth(context.Background(), chain).PostMethodArgs(postMethodArgs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChainsAPI.TransferEth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferEth`: TransferEth200Response
	fmt.Fprintf(os.Stdout, "Response from `ChainsAPI.TransferEth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferEthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postMethodArgs** | [**PostMethodArgs**](PostMethodArgs.md) |  | 

### Return type

[**TransferEth200Response**](TransferEth200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

