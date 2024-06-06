# \TxmAPI

All URIs are relative to *https://your_deployment.multibaas.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelTransaction**](TxmAPI.md#CancelTransaction) | **Post** /chains/{chain}/txm/{wallet_address}/nonce/{nonce}/cancel | Cancel transaction
[**CountWalletTransactions**](TxmAPI.md#CountWalletTransactions) | **Get** /chains/{chain}/txm/{wallet_address}/count | Count all transactions for a wallet
[**ListWalletTransactions**](TxmAPI.md#ListWalletTransactions) | **Get** /chains/{chain}/txm/{wallet_address} | List transactions for a wallet
[**SpeedUpTransaction**](TxmAPI.md#SpeedUpTransaction) | **Post** /chains/{chain}/txm/{wallet_address}/nonce/{nonce}/speed_up | Speed up transaction



## CancelTransaction

> TransferEth200Response CancelTransaction(ctx, chain, walletAddress, nonce).GasParams(gasParams).Execute()

Cancel transaction



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
	walletAddress := "walletAddress_example" // string | An HSM ethereum address.
	nonce := int64(789) // int64 | Transaction nonce.
	gasParams := *openapiclient.NewGasParams() // GasParams |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TxmAPI.CancelTransaction(context.Background(), chain, walletAddress, nonce).GasParams(gasParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TxmAPI.CancelTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelTransaction`: TransferEth200Response
	fmt.Fprintf(os.Stdout, "Response from `TxmAPI.CancelTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 
**walletAddress** | **string** | An HSM ethereum address. | 
**nonce** | **int64** | Transaction nonce. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **gasParams** | [**GasParams**](GasParams.md) |  | 

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


## CountWalletTransactions

> CountWalletTransactions200Response CountWalletTransactions(ctx, chain, walletAddress).Execute()

Count all transactions for a wallet



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
	walletAddress := "walletAddress_example" // string | An HSM ethereum address.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TxmAPI.CountWalletTransactions(context.Background(), chain, walletAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TxmAPI.CountWalletTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountWalletTransactions`: CountWalletTransactions200Response
	fmt.Fprintf(os.Stdout, "Response from `TxmAPI.CountWalletTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 
**walletAddress** | **string** | An HSM ethereum address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountWalletTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CountWalletTransactions200Response**](CountWalletTransactions200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWalletTransactions

> ListWalletTransactions200Response ListWalletTransactions(ctx, chain, walletAddress).Hash(hash).Nonce(nonce).Status(status).Limit(limit).Offset(offset).Execute()

List transactions for a wallet



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
	walletAddress := "walletAddress_example" // string | An HSM ethereum address.
	hash := "hash_example" // string | Filter transactions by transaction hash. To filter for multiple hashes, use ampersands: `?hash=HASH1&hash=HASH2&hash=HASH3` (optional)
	nonce := int64(789) // int64 | Filter transactions by nonce (optional)
	status := openapiclient.TransactionStatus("pending") // TransactionStatus | Filter transactions by status (optional)
	limit := int64(789) // int64 |  (optional)
	offset := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TxmAPI.ListWalletTransactions(context.Background(), chain, walletAddress).Hash(hash).Nonce(nonce).Status(status).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TxmAPI.ListWalletTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWalletTransactions`: ListWalletTransactions200Response
	fmt.Fprintf(os.Stdout, "Response from `TxmAPI.ListWalletTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 
**walletAddress** | **string** | An HSM ethereum address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWalletTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hash** | **string** | Filter transactions by transaction hash. To filter for multiple hashes, use ampersands: &#x60;?hash&#x3D;HASH1&amp;hash&#x3D;HASH2&amp;hash&#x3D;HASH3&#x60; | 
 **nonce** | **int64** | Filter transactions by nonce | 
 **status** | [**TransactionStatus**](TransactionStatus.md) | Filter transactions by status | 
 **limit** | **int64** |  | 
 **offset** | **int64** |  | 

### Return type

[**ListWalletTransactions200Response**](ListWalletTransactions200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpeedUpTransaction

> TransferEth200Response SpeedUpTransaction(ctx, chain, walletAddress, nonce).GasParams(gasParams).Execute()

Speed up transaction



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
	walletAddress := "walletAddress_example" // string | An HSM ethereum address.
	nonce := int64(789) // int64 | Transaction nonce.
	gasParams := *openapiclient.NewGasParams() // GasParams |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TxmAPI.SpeedUpTransaction(context.Background(), chain, walletAddress, nonce).GasParams(gasParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TxmAPI.SpeedUpTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpeedUpTransaction`: TransferEth200Response
	fmt.Fprintf(os.Stdout, "Response from `TxmAPI.SpeedUpTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 
**walletAddress** | **string** | An HSM ethereum address. | 
**nonce** | **int64** | Transaction nonce. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpeedUpTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **gasParams** | [**GasParams**](GasParams.md) |  | 

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

