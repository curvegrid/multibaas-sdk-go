# \AddressesAPI

All URIs are relative to *https://your_deployment.multibaas.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAddress**](AddressesAPI.md#DeleteAddress) | **Delete** /chains/{chain}/addresses/{address-or-alias} | Delete address
[**GetAddress**](AddressesAPI.md#GetAddress) | **Get** /chains/{chain}/addresses/{address-or-alias} | Get address
[**ListAddresses**](AddressesAPI.md#ListAddresses) | **Get** /chains/{chain}/addresses | List addresses
[**SetAddress**](AddressesAPI.md#SetAddress) | **Post** /chains/{chain}/addresses | Create or update address



## DeleteAddress

> BaseResponse DeleteAddress(ctx, chain, addressOrAlias).Execute()

Delete address



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
	addressOrAlias := "addressOrAlias_example" // string | An address or the alias of an address.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.DeleteAddress(context.Background(), chain, addressOrAlias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.DeleteAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAddress`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.DeleteAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 
**addressOrAlias** | **string** | An address or the alias of an address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAddressRequest struct via the builder pattern


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


## GetAddress

> SetAddress201Response GetAddress(ctx, chain, addressOrAlias).Include(include).Execute()

Get address



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
	addressOrAlias := "addressOrAlias_example" // string | An address or the alias of an address.
	include := []string{"Include_example"} // []string | Optional data to fetch from the blockchain: - `balance` to get the balance of this address. - `code` to get the code at this address. - `nonce` to get the next available transaction nonce for this address. - `contractLookup` to get the contract(s) details for this address.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.GetAddress(context.Background(), chain, addressOrAlias).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.GetAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddress`: SetAddress201Response
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.GetAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 
**addressOrAlias** | **string** | An address or the alias of an address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | **[]string** | Optional data to fetch from the blockchain: - &#x60;balance&#x60; to get the balance of this address. - &#x60;code&#x60; to get the code at this address. - &#x60;nonce&#x60; to get the next available transaction nonce for this address. - &#x60;contractLookup&#x60; to get the contract(s) details for this address.  | 

### Return type

[**SetAddress201Response**](SetAddress201Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAddresses

> ListAddresses200Response ListAddresses(ctx, chain).Execute()

List addresses



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
	resp, r, err := apiClient.AddressesAPI.ListAddresses(context.Background(), chain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.ListAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAddresses`: ListAddresses200Response
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.ListAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListAddresses200Response**](ListAddresses200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAddress

> SetAddress201Response SetAddress(ctx, chain).AddressAlias(addressAlias).Execute()

Create or update address



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
	addressAlias := *openapiclient.NewAddressAlias("Alias_example", "Address_example") // AddressAlias | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.SetAddress(context.Background(), chain).AddressAlias(addressAlias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.SetAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAddress`: SetAddress201Response
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.SetAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addressAlias** | [**AddressAlias**](AddressAlias.md) |  | 

### Return type

[**SetAddress201Response**](SetAddress201Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

