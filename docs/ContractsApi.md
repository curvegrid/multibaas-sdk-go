# \ContractsApi

All URIs are relative to *https://your_deployment.multibaas.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallContractFunction**](ContractsApi.md#CallContractFunction) | **Post** /chains/{chain}/addresses/{address-or-label}/contracts/{contract}/methods/{method} | Call a contract function
[**CreateContract**](ContractsApi.md#CreateContract) | **Post** /contracts/{contract} | Create a contract
[**CreateContracts**](ContractsApi.md#CreateContracts) | **Post** /contracts | Create multiple contracts
[**DeleteContract**](ContractsApi.md#DeleteContract) | **Delete** /contracts/{contract} | Delete a contract
[**DeleteContractVersion**](ContractsApi.md#DeleteContractVersion) | **Delete** /contracts/{contract}/{version} | Delete a contract version
[**DeployContract**](ContractsApi.md#DeployContract) | **Post** /contracts/{contract}/deploy | Deploy a contract
[**DeployContractVersion**](ContractsApi.md#DeployContractVersion) | **Post** /contracts/{contract}/{version}/deploy | Deploy a contract version
[**GetContract**](ContractsApi.md#GetContract) | **Get** /contracts/{contract} | Get a contract
[**GetContractVersion**](ContractsApi.md#GetContractVersion) | **Get** /contracts/{contract}/{version} | Get a contract version
[**GetContractVersions**](ContractsApi.md#GetContractVersions) | **Get** /contracts/{contract}/all | Get all contract versions
[**GetEventMonitorStatus**](ContractsApi.md#GetEventMonitorStatus) | **Get** /chains/{chain}/addresses/{address-or-label}/contracts/{contract}/status | Get event monitor status
[**GetEventTypeConversions**](ContractsApi.md#GetEventTypeConversions) | **Get** /contracts/{contract}/{version}/events/{event} | Get event type conversions
[**GetFunctionTypeConversions**](ContractsApi.md#GetFunctionTypeConversions) | **Get** /contracts/{contract}/{version}/methods/{method} | Get function type conversions
[**LinkAddressContract**](ContractsApi.md#LinkAddressContract) | **Post** /chains/{chain}/addresses/{address-or-label}/contracts | Link address and contract
[**ListContractVersions**](ContractsApi.md#ListContractVersions) | **Get** /contracts/{contract}/versions | List all contract versions
[**ListContracts**](ContractsApi.md#ListContracts) | **Get** /contracts | List contracts
[**SetEventTypeConversions**](ContractsApi.md#SetEventTypeConversions) | **Post** /contracts/{contract}/{version}/events/{event} | Set event type conversions
[**SetFunctionTypeConversions**](ContractsApi.md#SetFunctionTypeConversions) | **Post** /contracts/{contract}/{version}/methods/{method} | Set function type conversions
[**UnlinkAddressContract**](ContractsApi.md#UnlinkAddressContract) | **Delete** /chains/{chain}/addresses/{address-or-label}/contracts/{contract} | Unlink address and contract



## CallContractFunction

> CallContractFunction200Response CallContractFunction(ctx, chain, addressOrLabel, contract, method).PostMethodArgs(postMethodArgs).Execute()

Call a contract function



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
    addressOrLabel := "addressOrLabel_example" // string | An address or the label of an address.
    contract := "contract_example" // string | 
    method := "method_example" // string | Contract function.
    postMethodArgs := *openapiclient.NewPostMethodArgs() // PostMethodArgs |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.CallContractFunction(context.Background(), chain, addressOrLabel, contract, method).PostMethodArgs(postMethodArgs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.CallContractFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallContractFunction`: CallContractFunction200Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.CallContractFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 
**addressOrLabel** | **string** | An address or the label of an address. | 
**contract** | **string** |  | 
**method** | **string** | Contract function. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCallContractFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **postMethodArgs** | [**PostMethodArgs**](PostMethodArgs.md) |  | 

### Return type

[**CallContractFunction200Response**](CallContractFunction200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContract

> GetContract200Response CreateContract(ctx, contract).BaseContract(baseContract).Execute()

Create a contract



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
    contract := "contract_example" // string | 
    baseContract := *openapiclient.NewBaseContract("Label_example", "ContractName_example", "Version_example", "RawAbi_example", "UserDoc_example", "DeveloperDoc_example") // BaseContract |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.CreateContract(context.Background(), contract).BaseContract(baseContract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.CreateContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContract`: GetContract200Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.CreateContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseContract** | [**BaseContract**](BaseContract.md) |  | 

### Return type

[**GetContract200Response**](GetContract200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContracts

> SubmitSignedTransaction200Response CreateContracts(ctx).BaseContract(baseContract).Execute()

Create multiple contracts



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
    baseContract := []openapiclient.BaseContract{*openapiclient.NewBaseContract("Label_example", "ContractName_example", "Version_example", "RawAbi_example", "UserDoc_example", "DeveloperDoc_example")} // []BaseContract |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.CreateContracts(context.Background()).BaseContract(baseContract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.CreateContracts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContracts`: SubmitSignedTransaction200Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.CreateContracts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContractsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseContract** | [**[]BaseContract**](BaseContract.md) |  | 

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


## DeleteContract

> BaseResponse DeleteContract(ctx, contract).Execute()

Delete a contract



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
    contract := "contract_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.DeleteContract(context.Background(), contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.DeleteContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteContract`: BaseResponse
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.DeleteContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContractRequest struct via the builder pattern


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


## DeleteContractVersion

> BaseResponse DeleteContractVersion(ctx, contract, version).Execute()

Delete a contract version



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
    contract := "contract_example" // string | 
    version := "version_example" // string | Contract Version.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.DeleteContractVersion(context.Background(), contract, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.DeleteContractVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteContractVersion`: BaseResponse
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.DeleteContractVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | **string** |  | 
**version** | **string** | Contract Version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContractVersionRequest struct via the builder pattern


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


## DeployContract

> DeployContract200Response DeployContract(ctx, contract).PostMethodArgs(postMethodArgs).Execute()

Deploy a contract



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
    contract := "contract_example" // string | 
    postMethodArgs := *openapiclient.NewPostMethodArgs() // PostMethodArgs |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.DeployContract(context.Background(), contract).PostMethodArgs(postMethodArgs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.DeployContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployContract`: DeployContract200Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.DeployContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postMethodArgs** | [**PostMethodArgs**](PostMethodArgs.md) |  | 

### Return type

[**DeployContract200Response**](DeployContract200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeployContractVersion

> DeployContract200Response DeployContractVersion(ctx, contract, version).PostMethodArgs(postMethodArgs).Execute()

Deploy a contract version



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
    contract := "contract_example" // string | 
    version := "version_example" // string | Contract Version.
    postMethodArgs := *openapiclient.NewPostMethodArgs() // PostMethodArgs |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.DeployContractVersion(context.Background(), contract, version).PostMethodArgs(postMethodArgs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.DeployContractVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployContractVersion`: DeployContract200Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.DeployContractVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | **string** |  | 
**version** | **string** | Contract Version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployContractVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postMethodArgs** | [**PostMethodArgs**](PostMethodArgs.md) |  | 

### Return type

[**DeployContract200Response**](DeployContract200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContract

> GetContract200Response GetContract(ctx, contract).Execute()

Get a contract



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
    contract := "contract_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.GetContract(context.Background(), contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.GetContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContract`: GetContract200Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.GetContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContract200Response**](GetContract200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractVersion

> GetContract200Response GetContractVersion(ctx, contract, version).Execute()

Get a contract version



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
    contract := "contract_example" // string | 
    version := "version_example" // string | Contract Version.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.GetContractVersion(context.Background(), contract, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.GetContractVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractVersion`: GetContract200Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.GetContractVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | **string** |  | 
**version** | **string** | Contract Version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetContract200Response**](GetContract200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractVersions

> GetContractVersions200Response GetContractVersions(ctx, contract).Execute()

Get all contract versions



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
    contract := "contract_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.GetContractVersions(context.Background(), contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.GetContractVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractVersions`: GetContractVersions200Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.GetContractVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContractVersions200Response**](GetContractVersions200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventMonitorStatus

> GetEventMonitorStatus200Response GetEventMonitorStatus(ctx, chain, addressOrLabel, contract).Execute()

Get event monitor status



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
    addressOrLabel := "addressOrLabel_example" // string | An address or the label of an address.
    contract := "contract_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.GetEventMonitorStatus(context.Background(), chain, addressOrLabel, contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.GetEventMonitorStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventMonitorStatus`: GetEventMonitorStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.GetEventMonitorStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 
**addressOrLabel** | **string** | An address or the label of an address. | 
**contract** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventMonitorStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetEventMonitorStatus200Response**](GetEventMonitorStatus200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventTypeConversions

> GetEventTypeConversions200Response GetEventTypeConversions(ctx, contract, version, event).Execute()

Get event type conversions



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
    contract := "contract_example" // string | 
    version := "version_example" // string | Contract Version.
    event := "event_example" // string | Contract Event.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.GetEventTypeConversions(context.Background(), contract, version, event).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.GetEventTypeConversions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventTypeConversions`: GetEventTypeConversions200Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.GetEventTypeConversions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | **string** |  | 
**version** | **string** | Contract Version. | 
**event** | **string** | Contract Event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventTypeConversionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetEventTypeConversions200Response**](GetEventTypeConversions200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionTypeConversions

> GetFunctionTypeConversions200Response GetFunctionTypeConversions(ctx, contract, version, method).Execute()

Get function type conversions



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
    contract := "contract_example" // string | 
    version := "version_example" // string | Contract Version.
    method := "method_example" // string | Contract function.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.GetFunctionTypeConversions(context.Background(), contract, version, method).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.GetFunctionTypeConversions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFunctionTypeConversions`: GetFunctionTypeConversions200Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.GetFunctionTypeConversions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | **string** |  | 
**version** | **string** | Contract Version. | 
**method** | **string** | Contract function. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionTypeConversionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetFunctionTypeConversions200Response**](GetFunctionTypeConversions200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkAddressContract

> SetAddress201Response LinkAddressContract(ctx, chain, addressOrLabel).LinkAddressContractRequest(linkAddressContractRequest).Execute()

Link address and contract



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
    addressOrLabel := "addressOrLabel_example" // string | An address or the label of an address.
    linkAddressContractRequest := *openapiclient.NewLinkAddressContractRequest("Label_example") // LinkAddressContractRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.LinkAddressContract(context.Background(), chain, addressOrLabel).LinkAddressContractRequest(linkAddressContractRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.LinkAddressContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkAddressContract`: SetAddress201Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.LinkAddressContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 
**addressOrLabel** | **string** | An address or the label of an address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkAddressContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **linkAddressContractRequest** | [**LinkAddressContractRequest**](LinkAddressContractRequest.md) |  | 

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


## ListContractVersions

> ListContractVersions200Response ListContractVersions(ctx, contract).Execute()

List all contract versions



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
    contract := "contract_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.ListContractVersions(context.Background(), contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.ListContractVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContractVersions`: ListContractVersions200Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.ListContractVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContractVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListContractVersions200Response**](ListContractVersions200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContracts

> ListContracts200Response ListContracts(ctx).Execute()

List contracts



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
    resp, r, err := apiClient.ContractsApi.ListContracts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.ListContracts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContracts`: ListContracts200Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.ListContracts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListContractsRequest struct via the builder pattern


### Return type

[**ListContracts200Response**](ListContracts200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetEventTypeConversions

> BaseResponse SetEventTypeConversions(ctx, contract, version, event).ContractEventOptions(contractEventOptions).Execute()

Set event type conversions



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
    contract := "contract_example" // string | 
    version := "version_example" // string | Contract Version.
    event := "event_example" // string | Contract Event.
    contractEventOptions := *openapiclient.NewContractEventOptions([]openapiclient.ContractParameter{*openapiclient.NewContractParameter("TODO")}) // ContractEventOptions |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.SetEventTypeConversions(context.Background(), contract, version, event).ContractEventOptions(contractEventOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.SetEventTypeConversions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetEventTypeConversions`: BaseResponse
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.SetEventTypeConversions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | **string** |  | 
**version** | **string** | Contract Version. | 
**event** | **string** | Contract Event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetEventTypeConversionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contractEventOptions** | [**ContractEventOptions**](ContractEventOptions.md) |  | 

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


## SetFunctionTypeConversions

> BaseResponse SetFunctionTypeConversions(ctx, contract, version, method).ContractMethodOptions(contractMethodOptions).Execute()

Set function type conversions



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
    contract := "contract_example" // string | 
    version := "version_example" // string | Contract Version.
    method := "method_example" // string | Contract function.
    contractMethodOptions := *openapiclient.NewContractMethodOptions([]openapiclient.ContractParameter{*openapiclient.NewContractParameter("TODO")}) // ContractMethodOptions |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.SetFunctionTypeConversions(context.Background(), contract, version, method).ContractMethodOptions(contractMethodOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.SetFunctionTypeConversions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetFunctionTypeConversions`: BaseResponse
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.SetFunctionTypeConversions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | **string** |  | 
**version** | **string** | Contract Version. | 
**method** | **string** | Contract function. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFunctionTypeConversionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contractMethodOptions** | [**ContractMethodOptions**](ContractMethodOptions.md) |  | 

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


## UnlinkAddressContract

> SetAddress201Response UnlinkAddressContract(ctx, chain, addressOrLabel, contract).Execute()

Unlink address and contract



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
    addressOrLabel := "addressOrLabel_example" // string | An address or the label of an address.
    contract := "contract_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.UnlinkAddressContract(context.Background(), chain, addressOrLabel, contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.UnlinkAddressContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlinkAddressContract`: SetAddress201Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.UnlinkAddressContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | [**ChainName**](.md) | The blockchain chain label. | 
**addressOrLabel** | **string** | An address or the label of an address. | 
**contract** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkAddressContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

