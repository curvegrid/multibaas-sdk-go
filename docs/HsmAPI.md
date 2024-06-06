# \HsmAPI

All URIs are relative to *https://your_deployment.multibaas.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddHsmConfig**](HsmAPI.md#AddHsmConfig) | **Post** /hsm/config | Add HSM config
[**AddHsmKey**](HsmAPI.md#AddHsmKey) | **Post** /hsm/key | Add HSM key
[**CreateHsmKey**](HsmAPI.md#CreateHsmKey) | **Post** /hsm/key/new | Create HSM key
[**ListHsm**](HsmAPI.md#ListHsm) | **Get** /hsm | List HSM configs and wallets
[**ListHsmWallets**](HsmAPI.md#ListHsmWallets) | **Get** /hsm/wallets | List HSM wallets
[**RemoveHsmConfig**](HsmAPI.md#RemoveHsmConfig) | **Delete** /hsm/config/{client_id} | Remove HSM config
[**RemoveHsmKey**](HsmAPI.md#RemoveHsmKey) | **Delete** /hsm/key/{wallet_address} | Remove HSM key
[**SetLocalNonce**](HsmAPI.md#SetLocalNonce) | **Post** /chains/{chain}/hsm/nonce/{wallet_address} | Set local nonce
[**SignAndSubmitTransaction**](HsmAPI.md#SignAndSubmitTransaction) | **Post** /chains/{chain}/hsm/submit | Sign and submit transaction
[**SignData**](HsmAPI.md#SignData) | **Post** /chains/{chain}/hsm/sign | Sign data



## AddHsmConfig

> BaseResponse AddHsmConfig(ctx).BaseAzureAccount(baseAzureAccount).Execute()

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
	baseAzureAccount := *openapiclient.NewBaseAzureAccount("Label_example", "ClientID_example", "ClientSecret_example", "TenantID_example", "SubscriptionID_example", "BaseGroupName_example") // BaseAzureAccount |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HsmAPI.AddHsmConfig(context.Background()).BaseAzureAccount(baseAzureAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmAPI.AddHsmConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddHsmConfig`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmAPI.AddHsmConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddHsmConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseAzureAccount** | [**BaseAzureAccount**](BaseAzureAccount.md) |  | 

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


## AddHsmKey

> BaseResponse AddHsmKey(ctx).AddKey(addKey).Execute()

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
	resp, r, err := apiClient.HsmAPI.AddHsmKey(context.Background()).AddKey(addKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmAPI.AddHsmKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddHsmKey`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmAPI.AddHsmKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddHsmKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addKey** | [**AddKey**](AddKey.md) |  | 

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
	resp, r, err := apiClient.HsmAPI.CreateHsmKey(context.Background()).CreateKey(createKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmAPI.CreateHsmKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHsmKey`: CreateHsmKey200Response
	fmt.Fprintf(os.Stdout, "Response from `HsmAPI.CreateHsmKey`: %v\n", resp)
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
	resp, r, err := apiClient.HsmAPI.ListHsm(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmAPI.ListHsm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHsm`: ListHsm200Response
	fmt.Fprintf(os.Stdout, "Response from `HsmAPI.ListHsm`: %v\n", resp)
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
	resp, r, err := apiClient.HsmAPI.ListHsmWallets(context.Background()).KeyName(keyName).KeyVersion(keyVersion).VaultName(vaultName).BaseGroupName(baseGroupName).ClientId(clientId).PublicAddress(publicAddress).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmAPI.ListHsmWallets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHsmWallets`: ListHsmWallets200Response
	fmt.Fprintf(os.Stdout, "Response from `HsmAPI.ListHsmWallets`: %v\n", resp)
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


## RemoveHsmConfig

> BaseResponse RemoveHsmConfig(ctx, clientId).Execute()

Remove HSM config



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
	resp, r, err := apiClient.HsmAPI.RemoveHsmConfig(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmAPI.RemoveHsmConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveHsmConfig`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmAPI.RemoveHsmConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The HSM client ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveHsmConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveHsmKey

> BaseResponse RemoveHsmKey(ctx, walletAddress).Execute()

Remove HSM key



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
	resp, r, err := apiClient.HsmAPI.RemoveHsmKey(context.Background(), walletAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmAPI.RemoveHsmKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveHsmKey`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmAPI.RemoveHsmKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletAddress** | **string** | An HSM ethereum address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveHsmKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLocalNonce

> BaseResponse SetLocalNonce(ctx, chain, walletAddress).SetNonceRequest(setNonceRequest).Execute()

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
	resp, r, err := apiClient.HsmAPI.SetLocalNonce(context.Background(), chain, walletAddress).SetNonceRequest(setNonceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmAPI.SetLocalNonce``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetLocalNonce`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmAPI.SetLocalNonce`: %v\n", resp)
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

[**BaseResponse**](BaseResponse.md)

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
	resp, r, err := apiClient.HsmAPI.SignAndSubmitTransaction(context.Background(), chain).BaseTransactionToSign(baseTransactionToSign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmAPI.SignAndSubmitTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignAndSubmitTransaction`: TransferEth200Response
	fmt.Fprintf(os.Stdout, "Response from `HsmAPI.SignAndSubmitTransaction`: %v\n", resp)
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
	resp, r, err := apiClient.HsmAPI.SignData(context.Background(), chain).HSMSignRequest(hSMSignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmAPI.SignData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignData`: SignData200Response
	fmt.Fprintf(os.Stdout, "Response from `HsmAPI.SignData`: %v\n", resp)
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

