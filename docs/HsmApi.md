# \HsmApi

All URIs are relative to *https://your_deployment.multibaas.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddHsmConfig**](HsmApi.md#AddHsmConfig) | **Post** /hsm/config | Add HSM config
[**AddHsmKey**](HsmApi.md#AddHsmKey) | **Post** /hsm/key | Add HSM key
[**CreateHsmKey**](HsmApi.md#CreateHsmKey) | **Post** /hsm/key/new | Create HSM key
[**DeleteHsmConfig**](HsmApi.md#DeleteHsmConfig) | **Delete** /hsm/config/{client_id} | Delete HSM config
[**DeleteHsmKey**](HsmApi.md#DeleteHsmKey) | **Delete** /hsm/key/{wallet_address} | Delete HSM key
[**ListHsm**](HsmApi.md#ListHsm) | **Get** /hsm | List HSM configs and wallets
[**ListHsmWallets**](HsmApi.md#ListHsmWallets) | **Get** /hsm/wallets | List HSM wallets
[**SetLocalNonce**](HsmApi.md#SetLocalNonce) | **Post** /chains/{chain}/hsm/nonce/{wallet_address} | Set local nonce
[**SignAndSubmitTransaction**](HsmApi.md#SignAndSubmitTransaction) | **Post** /chains/{chain}/hsm/submit | Sign and submit transaction
[**SignData**](HsmApi.md#SignData) | **Post** /chains/{chain}/hsm/sign | Sign data



## AddHsmConfig

> SubmitSignedTransaction200Response AddHsmConfig(ctx).AzureAccountConfig(azureAccountConfig).Execute()

Add HSM config



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
    azureAccountConfig := *openapiclient.NewAzureAccountConfig("ClientID_example", "ClientSecret_example", "TenantID_example", "SubscriptionID_example", "BaseGroupName_example") // AzureAccountConfig |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HsmApi.AddHsmConfig(context.Background()).AzureAccountConfig(azureAccountConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HsmApi.AddHsmConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddHsmConfig`: SubmitSignedTransaction200Response
    fmt.Fprintf(os.Stdout, "Response from `HsmApi.AddHsmConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddHsmConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureAccountConfig** | [**AzureAccountConfig**](AzureAccountConfig.md) |  | 

### Return type

[**SubmitSignedTransaction200Response**](SubmitSignedTransaction200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddHsmKey

> SubmitSignedTransaction200Response AddHsmKey(ctx).AddKey(addKey).Execute()

Add HSM key



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
    addKey := *openapiclient.NewAddKey("ClientID_example", "KeyName_example", "KeyVersion_example", "VaultName_example") // AddKey |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HsmApi.AddHsmKey(context.Background()).AddKey(addKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HsmApi.AddHsmKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddHsmKey`: SubmitSignedTransaction200Response
    fmt.Fprintf(os.Stdout, "Response from `HsmApi.AddHsmKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddHsmKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addKey** | [**AddKey**](AddKey.md) |  | 

### Return type

[**SubmitSignedTransaction200Response**](SubmitSignedTransaction200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHsmKey

> CreateHsmKey200Response CreateHsmKey(ctx).CreateKey(createKey).Execute()

Create HSM key



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
    createKey := *openapiclient.NewCreateKey("ClientID_example", "KeyName_example", "VaultName_example", false) // CreateKey |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HsmApi.CreateHsmKey(context.Background()).CreateKey(createKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HsmApi.CreateHsmKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHsmKey`: CreateHsmKey200Response
    fmt.Fprintf(os.Stdout, "Response from `HsmApi.CreateHsmKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHsmKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createKey** | [**CreateKey**](CreateKey.md) |  | 

### Return type

[**CreateHsmKey200Response**](CreateHsmKey200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHsmConfig

> SubmitSignedTransaction200Response DeleteHsmConfig(ctx, clientId).Execute()

Delete HSM config



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
    clientId := "clientId_example" // string | The HSM client ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HsmApi.DeleteHsmConfig(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HsmApi.DeleteHsmConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteHsmConfig`: SubmitSignedTransaction200Response
    fmt.Fprintf(os.Stdout, "Response from `HsmApi.DeleteHsmConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The HSM client ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHsmConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubmitSignedTransaction200Response**](SubmitSignedTransaction200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHsmKey

> SubmitSignedTransaction200Response DeleteHsmKey(ctx, walletAddress).Execute()

Delete HSM key



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
    walletAddress := "walletAddress_example" // string | An HSM ethereum address.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HsmApi.DeleteHsmKey(context.Background(), walletAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HsmApi.DeleteHsmKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteHsmKey`: SubmitSignedTransaction200Response
    fmt.Fprintf(os.Stdout, "Response from `HsmApi.DeleteHsmKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletAddress** | **string** | An HSM ethereum address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHsmKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubmitSignedTransaction200Response**](SubmitSignedTransaction200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHsm

> ListHsm200Response ListHsm(ctx).Execute()

List HSM configs and wallets



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HsmApi.ListHsm(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HsmApi.ListHsm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHsm`: ListHsm200Response
    fmt.Fprintf(os.Stdout, "Response from `HsmApi.ListHsm`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListHsmRequest struct via the builder pattern


### Return type

[**ListHsm200Response**](ListHsm200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHsmWallets

> ListHsmWallets200Response ListHsmWallets(ctx).KeyName(keyName).KeyVersion(keyVersion).VaultName(vaultName).BaseGroupName(baseGroupName).ClientId(clientId).PublicAddress(publicAddress).Limit(limit).Offset(offset).Execute()

List HSM wallets



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
    keyName := "keyName_example" // string | Filter wallets by a key name. (optional)
    keyVersion := "keyVersion_example" // string | Filter wallets by a key version. (optional)
    vaultName := "vaultName_example" // string | Filter wallets by a vault name. (optional)
    baseGroupName := "baseGroupName_example" // string | Filter wallets by a base group name. (optional)
    clientId := "clientId_example" // string | Filter wallets by a client ID. (optional)
    publicAddress := "publicAddress_example" // string | Filter wallets by a public address. (optional)
    limit := int64(789) // int64 |  (optional)
    offset := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HsmApi.ListHsmWallets(context.Background()).KeyName(keyName).KeyVersion(keyVersion).VaultName(vaultName).BaseGroupName(baseGroupName).ClientId(clientId).PublicAddress(publicAddress).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HsmApi.ListHsmWallets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHsmWallets`: ListHsmWallets200Response
    fmt.Fprintf(os.Stdout, "Response from `HsmApi.ListHsmWallets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHsmWalletsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyName** | **string** | Filter wallets by a key name. | 
 **keyVersion** | **string** | Filter wallets by a key version. | 
 **vaultName** | **string** | Filter wallets by a vault name. | 
 **baseGroupName** | **string** | Filter wallets by a base group name. | 
 **clientId** | **string** | Filter wallets by a client ID. | 
 **publicAddress** | **string** | Filter wallets by a public address. | 
 **limit** | **int64** |  | 
 **offset** | **int64** |  | 

### Return type

[**ListHsmWallets200Response**](ListHsmWallets200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLocalNonce

> SubmitSignedTransaction200Response SetLocalNonce(ctx, chain, walletAddress).SetNonceRequest(setNonceRequest).Execute()

Set local nonce



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
    setNonceRequest := *openapiclient.NewSetNonceRequest() // SetNonceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HsmApi.SetLocalNonce(context.Background(), chain, walletAddress).SetNonceRequest(setNonceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HsmApi.SetLocalNonce``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetLocalNonce`: SubmitSignedTransaction200Response
    fmt.Fprintf(os.Stdout, "Response from `HsmApi.SetLocalNonce`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 
**walletAddress** | **string** | An HSM ethereum address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLocalNonceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setNonceRequest** | [**SetNonceRequest**](SetNonceRequest.md) |  | 

### Return type

[**SubmitSignedTransaction200Response**](SubmitSignedTransaction200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignAndSubmitTransaction

> TransferEth200Response SignAndSubmitTransaction(ctx, chain).BaseTransactionToSign(baseTransactionToSign).Execute()

Sign and submit transaction



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
    baseTransactionToSign := *openapiclient.NewBaseTransactionToSign(*openapiclient.NewBaseTransactionToSignTx(int64(123), "From_example", "Value_example", "Data_example", int64(123))) // BaseTransactionToSign |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HsmApi.SignAndSubmitTransaction(context.Background(), chain).BaseTransactionToSign(baseTransactionToSign).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HsmApi.SignAndSubmitTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignAndSubmitTransaction`: TransferEth200Response
    fmt.Fprintf(os.Stdout, "Response from `HsmApi.SignAndSubmitTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignAndSubmitTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseTransactionToSign** | [**BaseTransactionToSign**](BaseTransactionToSign.md) |  | 

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


## SignData

> SignData200Response SignData(ctx, chain).HSMSignRequest(hSMSignRequest).Execute()

Sign data



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
    hSMSignRequest := *openapiclient.NewHSMSignRequest("Address_example", "Data_example") // HSMSignRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HsmApi.SignData(context.Background(), chain).HSMSignRequest(hSMSignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HsmApi.SignData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignData`: SignData200Response
    fmt.Fprintf(os.Stdout, "Response from `HsmApi.SignData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hSMSignRequest** | [**HSMSignRequest**](HSMSignRequest.md) |  | 

### Return type

[**SignData200Response**](SignData200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

