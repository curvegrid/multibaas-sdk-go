# \EventQueriesAPI

All URIs are relative to *https://your_deployment.multibaas.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountEventQueryRecords**](EventQueriesAPI.md#CountEventQueryRecords) | **Get** /queries/{event_query}/count | Count event query records
[**DeleteEventQuery**](EventQueriesAPI.md#DeleteEventQuery) | **Delete** /queries/{event_query} | Delete event query
[**ExecuteArbitraryEventQuery**](EventQueriesAPI.md#ExecuteArbitraryEventQuery) | **Post** /queries | Execute arbitrary event query
[**ExecuteEventQuery**](EventQueriesAPI.md#ExecuteEventQuery) | **Get** /queries/{event_query}/results | Execute event query
[**GetEventQuery**](EventQueriesAPI.md#GetEventQuery) | **Get** /queries/{event_query} | Get event query
[**ListEventQueries**](EventQueriesAPI.md#ListEventQueries) | **Get** /queries | List event queries
[**SetEventQuery**](EventQueriesAPI.md#SetEventQuery) | **Put** /queries/{event_query} | Create or update event query



## CountEventQueryRecords

> CountEventQueryRecords200Response CountEventQueryRecords(ctx, eventQuery).Execute()

Count event query records



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
	eventQuery := "eventQuery_example" // string | An event query label.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventQueriesAPI.CountEventQueryRecords(context.Background(), eventQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventQueriesAPI.CountEventQueryRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountEventQueryRecords`: CountEventQueryRecords200Response
	fmt.Fprintf(os.Stdout, "Response from `EventQueriesAPI.CountEventQueryRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventQuery** | **string** | An event query label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountEventQueryRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CountEventQueryRecords200Response**](CountEventQueryRecords200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEventQuery

> BaseResponse DeleteEventQuery(ctx, eventQuery).Execute()

Delete event query



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
	eventQuery := "eventQuery_example" // string | An event query label.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventQueriesAPI.DeleteEventQuery(context.Background(), eventQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventQueriesAPI.DeleteEventQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEventQuery`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `EventQueriesAPI.DeleteEventQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventQuery** | **string** | An event query label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventQueryRequest struct via the builder pattern


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


## ExecuteArbitraryEventQuery

> ExecuteArbitraryEventQuery200Response ExecuteArbitraryEventQuery(ctx).Offset(offset).Limit(limit).EventQuery(eventQuery).Execute()

Execute arbitrary event query



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
	offset := int64(789) // int64 |  (optional)
	limit := int64(789) // int64 |  (optional)
	eventQuery := *openapiclient.NewEventQuery([]openapiclient.EventQueryEvent{*openapiclient.NewEventQueryEvent("EventName_example", []openapiclient.EventQueryField{*openapiclient.NewEventQueryField(openapiclient.FieldType("input"))})}) // EventQuery |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventQueriesAPI.ExecuteArbitraryEventQuery(context.Background()).Offset(offset).Limit(limit).EventQuery(eventQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventQueriesAPI.ExecuteArbitraryEventQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteArbitraryEventQuery`: ExecuteArbitraryEventQuery200Response
	fmt.Fprintf(os.Stdout, "Response from `EventQueriesAPI.ExecuteArbitraryEventQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteArbitraryEventQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** |  | 
 **limit** | **int64** |  | 
 **eventQuery** | [**EventQuery**](EventQuery.md) |  | 

### Return type

[**ExecuteArbitraryEventQuery200Response**](ExecuteArbitraryEventQuery200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteEventQuery

> ExecuteArbitraryEventQuery200Response ExecuteEventQuery(ctx, eventQuery).Offset(offset).Limit(limit).Execute()

Execute event query



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
	eventQuery := "eventQuery_example" // string | An event query label.
	offset := int64(789) // int64 |  (optional)
	limit := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventQueriesAPI.ExecuteEventQuery(context.Background(), eventQuery).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventQueriesAPI.ExecuteEventQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteEventQuery`: ExecuteArbitraryEventQuery200Response
	fmt.Fprintf(os.Stdout, "Response from `EventQueriesAPI.ExecuteEventQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventQuery** | **string** | An event query label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteEventQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int64** |  | 
 **limit** | **int64** |  | 

### Return type

[**ExecuteArbitraryEventQuery200Response**](ExecuteArbitraryEventQuery200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventQuery

> GetEventQuery200Response GetEventQuery(ctx, eventQuery).Execute()

Get event query



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
	eventQuery := "eventQuery_example" // string | An event query label.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventQueriesAPI.GetEventQuery(context.Background(), eventQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventQueriesAPI.GetEventQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventQuery`: GetEventQuery200Response
	fmt.Fprintf(os.Stdout, "Response from `EventQueriesAPI.GetEventQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventQuery** | **string** | An event query label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEventQuery200Response**](GetEventQuery200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEventQueries

> ListEventQueries200Response ListEventQueries(ctx).Execute()

List event queries



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
	resp, r, err := apiClient.EventQueriesAPI.ListEventQueries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventQueriesAPI.ListEventQueries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEventQueries`: ListEventQueries200Response
	fmt.Fprintf(os.Stdout, "Response from `EventQueriesAPI.ListEventQueries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListEventQueriesRequest struct via the builder pattern


### Return type

[**ListEventQueries200Response**](ListEventQueries200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetEventQuery

> BaseResponse SetEventQuery(ctx, eventQuery).EventQuery2(eventQuery2).Execute()

Create or update event query



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
	eventQuery := "eventQuery_example" // string | An event query label.
	eventQuery2 := *openapiclient.NewEventQuery([]openapiclient.EventQueryEvent{*openapiclient.NewEventQueryEvent("EventName_example", []openapiclient.EventQueryField{*openapiclient.NewEventQueryField(openapiclient.FieldType("input"))})}) // EventQuery |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventQueriesAPI.SetEventQuery(context.Background(), eventQuery).EventQuery2(eventQuery2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventQueriesAPI.SetEventQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetEventQuery`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `EventQueriesAPI.SetEventQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventQuery** | **string** | An event query label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetEventQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventQuery2** | [**EventQuery**](EventQuery.md) |  | 

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

