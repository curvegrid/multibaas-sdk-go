# \WebhooksAPI

All URIs are relative to *https://your_deployment.multibaas.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountWebhookEvents**](WebhooksAPI.md#CountWebhookEvents) | **Get** /webhooks/{webhookID}/events/count | Count webhook events
[**CountWebhooks**](WebhooksAPI.md#CountWebhooks) | **Get** /webhooks/count | Count webhooks
[**CreateWebhook**](WebhooksAPI.md#CreateWebhook) | **Post** /webhooks | Create webhook
[**DeleteWebhook**](WebhooksAPI.md#DeleteWebhook) | **Delete** /webhooks/{webhookID} | Delete webhook
[**GetWebhook**](WebhooksAPI.md#GetWebhook) | **Get** /webhooks/{webhookID} | Get webhook
[**ListWebhookEvents**](WebhooksAPI.md#ListWebhookEvents) | **Get** /webhooks/{webhookID}/events | List webhook events
[**ListWebhooks**](WebhooksAPI.md#ListWebhooks) | **Get** /webhooks | List webhooks
[**UpdateWebhook**](WebhooksAPI.md#UpdateWebhook) | **Put** /webhooks/{webhookID} | Update webhook



## CountWebhookEvents

> CountWebhookEvents200Response CountWebhookEvents(ctx, webhookID).Execute()

Count webhook events



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
	webhookID := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.CountWebhookEvents(context.Background(), webhookID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CountWebhookEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountWebhookEvents`: CountWebhookEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CountWebhookEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountWebhookEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CountWebhookEvents200Response**](CountWebhookEvents200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountWebhooks

> CountWebhooks200Response CountWebhooks(ctx).Execute()

Count webhooks



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
	resp, r, err := apiClient.WebhooksAPI.CountWebhooks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CountWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountWebhooks`: CountWebhooks200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CountWebhooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCountWebhooksRequest struct via the builder pattern


### Return type

[**CountWebhooks200Response**](CountWebhooks200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebhook

> CreateWebhook200Response CreateWebhook(ctx).BaseWebhookEndpoint(baseWebhookEndpoint).Execute()

Create webhook



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
	baseWebhookEndpoint := *openapiclient.NewBaseWebhookEndpoint("Url_example", "Label_example", []openapiclient.WebhookEventsType{openapiclient.WebhookEventsType("transaction.included")}) // BaseWebhookEndpoint |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.CreateWebhook(context.Background()).BaseWebhookEndpoint(baseWebhookEndpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CreateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhook`: CreateWebhook200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseWebhookEndpoint** | [**BaseWebhookEndpoint**](BaseWebhookEndpoint.md) |  | 

### Return type

[**CreateWebhook200Response**](CreateWebhook200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> BaseResponse DeleteWebhook(ctx, webhookID).Execute()

Delete webhook



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
	webhookID := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.DeleteWebhook(context.Background(), webhookID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.DeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebhook`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.DeleteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


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


## GetWebhook

> CreateWebhook200Response GetWebhook(ctx, webhookID).Execute()

Get webhook



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
	webhookID := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.GetWebhook(context.Background(), webhookID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhook`: CreateWebhook200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateWebhook200Response**](CreateWebhook200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhookEvents

> ListWebhookEvents200Response ListWebhookEvents(ctx, webhookID).Limit(limit).Offset(offset).Execute()

List webhook events



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
	webhookID := int64(789) // int64 | 
	limit := int64(789) // int64 |  (optional)
	offset := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.ListWebhookEvents(context.Background(), webhookID).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ListWebhookEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhookEvents`: ListWebhookEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ListWebhookEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhookEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** |  | 
 **offset** | **int64** |  | 

### Return type

[**ListWebhookEvents200Response**](ListWebhookEvents200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhooks

> ListWebhooks200Response ListWebhooks(ctx).Limit(limit).Offset(offset).Execute()

List webhooks



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
	limit := int64(789) // int64 |  (optional)
	offset := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.ListWebhooks(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ListWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhooks`: ListWebhooks200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ListWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** |  | 
 **offset** | **int64** |  | 

### Return type

[**ListWebhooks200Response**](ListWebhooks200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> CreateWebhook200Response UpdateWebhook(ctx, webhookID).BaseWebhookEndpoint(baseWebhookEndpoint).Execute()

Update webhook



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
	webhookID := int64(789) // int64 | 
	baseWebhookEndpoint := *openapiclient.NewBaseWebhookEndpoint("Url_example", "Label_example", []openapiclient.WebhookEventsType{openapiclient.WebhookEventsType("transaction.included")}) // BaseWebhookEndpoint |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.UpdateWebhook(context.Background(), webhookID).BaseWebhookEndpoint(baseWebhookEndpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.UpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhook`: CreateWebhook200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseWebhookEndpoint** | [**BaseWebhookEndpoint**](BaseWebhookEndpoint.md) |  | 

### Return type

[**CreateWebhook200Response**](CreateWebhook200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

