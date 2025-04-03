# \EventsAPI

All URIs are relative to *https://your_deployment.multibaas.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventCount**](EventsAPI.md#GetEventCount) | **Get** /events/count | Get event count
[**ListEvents**](EventsAPI.md#ListEvents) | **Get** /events | List events



## GetEventCount

> GetEventCount200Response GetEventCount(ctx).BlockHash(blockHash).BlockNumber(blockNumber).TxIndexInBlock(txIndexInBlock).EventIndexInLog(eventIndexInLog).TxHash(txHash).FromConstructor(fromConstructor).Chain(chain).ContractAddress(contractAddress).ContractLabel(contractLabel).EventSignature(eventSignature).Limit(limit).Offset(offset).Execute()

Get event count



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
    blockHash := "blockHash_example" // string | Filter events by a block hash. (optional)
    blockNumber := int64(789) // int64 | Filter events by a block number. (optional)
    txIndexInBlock := int64(789) // int64 | Filter events by a transaction's index in the block. (optional)
    eventIndexInLog := int64(789) // int64 | Filter events by index in the log. (optional)
    txHash := "txHash_example" // string | Filter events by a transaction hash. (optional)
    fromConstructor := true // bool | Filter events by whether they were emitted from the constructor function. (optional)
    chain := openapiclient.ChainName("ethereum") // ChainName | Filter events by a chain name. (optional)
    contractAddress := "contractAddress_example" // string | Filter events by a contract address. (optional)
    contractLabel := "contractLabel_example" // string | Filter events by a contract label. (optional)
    eventSignature := "eventSignature_example" // string | Filter events by the signature. (optional)
    limit := int64(789) // int64 |  (optional) (default to 10)
    offset := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.GetEventCount(context.Background()).BlockHash(blockHash).BlockNumber(blockNumber).TxIndexInBlock(txIndexInBlock).EventIndexInLog(eventIndexInLog).TxHash(txHash).FromConstructor(fromConstructor).Chain(chain).ContractAddress(contractAddress).ContractLabel(contractLabel).EventSignature(eventSignature).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEventCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventCount`: GetEventCount200Response
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEventCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockHash** | **string** | Filter events by a block hash. | 
 **blockNumber** | **int64** | Filter events by a block number. | 
 **txIndexInBlock** | **int64** | Filter events by a transaction&#39;s index in the block. | 
 **eventIndexInLog** | **int64** | Filter events by index in the log. | 
 **txHash** | **string** | Filter events by a transaction hash. | 
 **fromConstructor** | **bool** | Filter events by whether they were emitted from the constructor function. | 
 **chain** | [**ChainName**](ChainName.md) | Filter events by a chain name. | 
 **contractAddress** | **string** | Filter events by a contract address. | 
 **contractLabel** | **string** | Filter events by a contract label. | 
 **eventSignature** | **string** | Filter events by the signature. | 
 **limit** | **int64** |  | [default to 10]
 **offset** | **int64** |  | 

### Return type

[**GetEventCount200Response**](GetEventCount200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvents

> ListEvents200Response ListEvents(ctx).BlockHash(blockHash).BlockNumber(blockNumber).TxIndexInBlock(txIndexInBlock).EventIndexInLog(eventIndexInLog).TxHash(txHash).FromConstructor(fromConstructor).Chain(chain).ContractAddress(contractAddress).ContractLabel(contractLabel).EventSignature(eventSignature).Limit(limit).Offset(offset).Execute()

List events



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
    blockHash := "blockHash_example" // string | Filter events by a block hash. (optional)
    blockNumber := int64(789) // int64 | Filter events by a block number. (optional)
    txIndexInBlock := int64(789) // int64 | Filter events by a transaction's index in the block. (optional)
    eventIndexInLog := int64(789) // int64 | Filter events by index in the log. (optional)
    txHash := "txHash_example" // string | Filter events by a transaction hash. (optional)
    fromConstructor := true // bool | Filter events by whether they were emitted from the constructor function. (optional)
    chain := openapiclient.ChainName("ethereum") // ChainName | Filter events by a chain name. (optional)
    contractAddress := "contractAddress_example" // string | Filter events by a contract address. (optional)
    contractLabel := "contractLabel_example" // string | Filter events by a contract label. (optional)
    eventSignature := "eventSignature_example" // string | Filter events by the signature. (optional)
    limit := int64(789) // int64 |  (optional) (default to 10)
    offset := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.ListEvents(context.Background()).BlockHash(blockHash).BlockNumber(blockNumber).TxIndexInBlock(txIndexInBlock).EventIndexInLog(eventIndexInLog).TxHash(txHash).FromConstructor(fromConstructor).Chain(chain).ContractAddress(contractAddress).ContractLabel(contractLabel).EventSignature(eventSignature).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ListEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEvents`: ListEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ListEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockHash** | **string** | Filter events by a block hash. | 
 **blockNumber** | **int64** | Filter events by a block number. | 
 **txIndexInBlock** | **int64** | Filter events by a transaction&#39;s index in the block. | 
 **eventIndexInLog** | **int64** | Filter events by index in the log. | 
 **txHash** | **string** | Filter events by a transaction hash. | 
 **fromConstructor** | **bool** | Filter events by whether they were emitted from the constructor function. | 
 **chain** | [**ChainName**](ChainName.md) | Filter events by a chain name. | 
 **contractAddress** | **string** | Filter events by a contract address. | 
 **contractLabel** | **string** | Filter events by a contract label. | 
 **eventSignature** | **string** | Filter events by the signature. | 
 **limit** | **int64** |  | [default to 10]
 **offset** | **int64** |  | 

### Return type

[**ListEvents200Response**](ListEvents200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

