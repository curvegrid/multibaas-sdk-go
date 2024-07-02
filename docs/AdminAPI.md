# \AdminAPI

All URIs are relative to *https://your_deployment.multibaas.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptInvite**](AdminAPI.md#AcceptInvite) | **Post** /invites/{inviteID} | Accept invite
[**AddCorsOrigin**](AdminAPI.md#AddCorsOrigin) | **Post** /cors | Add CORS origin
[**AddGroupApiKey**](AdminAPI.md#AddGroupApiKey) | **Put** /groups/{groupID}/api_keys/{apiKeyID} | Add API key to group
[**AddGroupRole**](AdminAPI.md#AddGroupRole) | **Put** /groups/{groupID}/roles/{roleShortName} | Add role to group
[**AddGroupUser**](AdminAPI.md#AddGroupUser) | **Put** /groups/{groupID}/users/{userID} | Add user to group
[**CheckInvite**](AdminAPI.md#CheckInvite) | **Get** /invites/{inviteID} | Check invite
[**CreateApiKey**](AdminAPI.md#CreateApiKey) | **Post** /api_keys | Create API key
[**DeleteApiKey**](AdminAPI.md#DeleteApiKey) | **Delete** /api_keys/{apiKeyID} | Delete API key
[**DeleteUser**](AdminAPI.md#DeleteUser) | **Delete** /users/{userID} | Delete user
[**GetApiKey**](AdminAPI.md#GetApiKey) | **Get** /api_keys/{apiKeyID} | Get API Key
[**InviteUser**](AdminAPI.md#InviteUser) | **Post** /invites | Invite user
[**ListApiKeys**](AdminAPI.md#ListApiKeys) | **Get** /api_keys | List API keys
[**ListAuditLogs**](AdminAPI.md#ListAuditLogs) | **Get** /systemactivities | List audit logs
[**ListCorsOrigins**](AdminAPI.md#ListCorsOrigins) | **Get** /cors | List CORS origins
[**ListGroups**](AdminAPI.md#ListGroups) | **Get** /groups | List groups
[**ListUserSigners**](AdminAPI.md#ListUserSigners) | **Get** /users/{userID}/signers | List user signers
[**ListUsers**](AdminAPI.md#ListUsers) | **Get** /users | List users
[**RemoveCorsOrigin**](AdminAPI.md#RemoveCorsOrigin) | **Delete** /cors/{originID} | Remove CORS Origin
[**RemoveGroupApiKey**](AdminAPI.md#RemoveGroupApiKey) | **Delete** /groups/{groupID}/api_keys/{apiKeyID} | Remove API key from group
[**RemoveGroupRole**](AdminAPI.md#RemoveGroupRole) | **Delete** /groups/{groupID}/roles/{roleShortName} | Remove role from group
[**RemoveGroupUser**](AdminAPI.md#RemoveGroupUser) | **Delete** /groups/{groupID}/users/{userID} | Remove user from group
[**RemoveUserSignerCloudWallet**](AdminAPI.md#RemoveUserSignerCloudWallet) | **Delete** /users/{userID}/cloudwallets/{wallet_address} | Remove user cloud wallet signer
[**RemoveUserSignerSafeAccount**](AdminAPI.md#RemoveUserSignerSafeAccount) | **Delete** /users/{userID}/safeaccounts/{wallet_address} | Remove user safe account signer
[**RemoveUserSignerWeb3Wallet**](AdminAPI.md#RemoveUserSignerWeb3Wallet) | **Delete** /users/{userID}/web3wallets/{wallet_address} | Remove user web3 wallet signer
[**SetUserSignerCloudWallet**](AdminAPI.md#SetUserSignerCloudWallet) | **Put** /users/{userID}/cloudwallets/{wallet_address} | Add or update user cloud wallet signer
[**SetUserSignerSafeAccount**](AdminAPI.md#SetUserSignerSafeAccount) | **Put** /users/{userID}/safeaccounts/{wallet_address} | Add or update user safe account signer
[**SetUserSignerWeb3Wallet**](AdminAPI.md#SetUserSignerWeb3Wallet) | **Put** /users/{userID}/web3wallets/{wallet_address} | Add or update user web3 wallet signer
[**UpdateApiKey**](AdminAPI.md#UpdateApiKey) | **Put** /api_keys/{apiKeyID} | Update API key



## AcceptInvite

> AcceptInvite200Response AcceptInvite(ctx, inviteID).AcceptInviteRequest(acceptInviteRequest).Execute()

Accept invite



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
	inviteID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	acceptInviteRequest := *openapiclient.NewAcceptInviteRequest() // AcceptInviteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AcceptInvite(context.Background(), inviteID).AcceptInviteRequest(acceptInviteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AcceptInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcceptInvite`: AcceptInvite200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AcceptInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptInviteRequest** | [**AcceptInviteRequest**](AcceptInviteRequest.md) |  | 

### Return type

[**AcceptInvite200Response**](AcceptInvite200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddCorsOrigin

> BaseResponse AddCorsOrigin(ctx).CORSOrigin(cORSOrigin).Execute()

Add CORS origin



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
	cORSOrigin := *openapiclient.NewCORSOrigin() // CORSOrigin |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AddCorsOrigin(context.Background()).CORSOrigin(cORSOrigin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AddCorsOrigin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddCorsOrigin`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AddCorsOrigin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCorsOriginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cORSOrigin** | [**CORSOrigin**](CORSOrigin.md) |  | 

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


## AddGroupApiKey

> BaseResponse AddGroupApiKey(ctx, groupID, apiKeyID).Execute()

Add API key to group



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
	groupID := int64(789) // int64 | 
	apiKeyID := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AddGroupApiKey(context.Background(), groupID, apiKeyID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AddGroupApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGroupApiKey`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AddGroupApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupID** | **int64** |  | 
**apiKeyID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupApiKeyRequest struct via the builder pattern


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


## AddGroupRole

> BaseResponse AddGroupRole(ctx, groupID, roleShortName).Execute()

Add role to group



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
	groupID := int64(789) // int64 | 
	roleShortName := "roleShortName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AddGroupRole(context.Background(), groupID, roleShortName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AddGroupRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGroupRole`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AddGroupRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupID** | **int64** |  | 
**roleShortName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupRoleRequest struct via the builder pattern


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


## AddGroupUser

> BaseResponse AddGroupUser(ctx, groupID, userID).Execute()

Add user to group



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
	groupID := int64(789) // int64 | 
	userID := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AddGroupUser(context.Background(), groupID, userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AddGroupUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGroupUser`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AddGroupUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupID** | **int64** |  | 
**userID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupUserRequest struct via the builder pattern


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


## CheckInvite

> BaseResponse CheckInvite(ctx, inviteID).Execute()

Check invite



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
	inviteID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CheckInvite(context.Background(), inviteID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CheckInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckInvite`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CheckInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApiKey

> CreateApiKey200Response CreateApiKey(ctx).CreateApiKeyRequest(createApiKeyRequest).Execute()

Create API key



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
	createApiKeyRequest := *openapiclient.NewCreateApiKeyRequest("Label_example") // CreateApiKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CreateApiKey(context.Background()).CreateApiKeyRequest(createApiKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CreateApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiKey`: CreateApiKey200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CreateApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApiKeyRequest** | [**CreateApiKeyRequest**](CreateApiKeyRequest.md) |  | 

### Return type

[**CreateApiKey200Response**](CreateApiKey200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiKey

> BaseResponse DeleteApiKey(ctx, apiKeyID).Execute()

Delete API key



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
	apiKeyID := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.DeleteApiKey(context.Background(), apiKeyID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.DeleteApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApiKey`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.DeleteApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiKeyRequest struct via the builder pattern


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


## DeleteUser

> BaseResponse DeleteUser(ctx, userID).Execute()

Delete user



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
	userID := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.DeleteUser(context.Background(), userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUser`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## GetApiKey

> GetApiKey200Response GetApiKey(ctx, apiKeyID).Execute()

Get API Key



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
	apiKeyID := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetApiKey(context.Background(), apiKeyID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiKey`: GetApiKey200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetApiKey200Response**](GetApiKey200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteUser

> BaseResponse InviteUser(ctx).Invite(invite).Execute()

Invite user



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
	invite := *openapiclient.NewInvite("Email_example") // Invite |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.InviteUser(context.Background()).Invite(invite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.InviteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InviteUser`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.InviteUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInviteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invite** | [**Invite**](Invite.md) |  | 

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


## ListApiKeys

> ListApiKeys200Response ListApiKeys(ctx).All(all).Execute()

List API keys



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
	all := true // bool | If true, returns all API keys on the system, otherwise, returns only the API keys owned by the calling user. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.ListApiKeys(context.Background()).All(all).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ListApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApiKeys`: ListApiKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ListApiKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **bool** | If true, returns all API keys on the system, otherwise, returns only the API keys owned by the calling user. | 

### Return type

[**ListApiKeys200Response**](ListApiKeys200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuditLogs

> ListAuditLogs200Response ListAuditLogs(ctx).Execute()

List audit logs



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
	resp, r, err := apiClient.AdminAPI.ListAuditLogs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ListAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAuditLogs`: ListAuditLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ListAuditLogs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAuditLogsRequest struct via the builder pattern


### Return type

[**ListAuditLogs200Response**](ListAuditLogs200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCorsOrigins

> ListCorsOrigins200Response ListCorsOrigins(ctx).Execute()

List CORS origins



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
	resp, r, err := apiClient.AdminAPI.ListCorsOrigins(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ListCorsOrigins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCorsOrigins`: ListCorsOrigins200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ListCorsOrigins`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCorsOriginsRequest struct via the builder pattern


### Return type

[**ListCorsOrigins200Response**](ListCorsOrigins200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroups

> ListGroups200Response ListGroups(ctx).UserID(userID).ApiKeyID(apiKeyID).Assignable(assignable).Execute()

List groups



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
	userID := int64(789) // int64 |  (optional)
	apiKeyID := int64(789) // int64 |  (optional)
	assignable := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.ListGroups(context.Background()).UserID(userID).ApiKeyID(apiKeyID).Assignable(assignable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ListGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroups`: ListGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ListGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userID** | **int64** |  | 
 **apiKeyID** | **int64** |  | 
 **assignable** | **bool** |  | 

### Return type

[**ListGroups200Response**](ListGroups200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserSigners

> ListUserSigners200Response ListUserSigners(ctx, userID).Execute()

List user signers



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
	userID := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.ListUserSigners(context.Background(), userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ListUserSigners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserSigners`: ListUserSigners200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ListUserSigners`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserSignersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListUserSigners200Response**](ListUserSigners200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> ListUsers200Response ListUsers(ctx).GroupID(groupID).Execute()

List users



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
	groupID := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.ListUsers(context.Background()).GroupID(groupID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsers`: ListUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupID** | **int64** |  | 

### Return type

[**ListUsers200Response**](ListUsers200Response.md)

### Authorization

[cookie](../README.md#cookie), [bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCorsOrigin

> BaseResponse RemoveCorsOrigin(ctx, originID).Execute()

Remove CORS Origin



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
	originID := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.RemoveCorsOrigin(context.Background(), originID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.RemoveCorsOrigin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveCorsOrigin`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.RemoveCorsOrigin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**originID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCorsOriginRequest struct via the builder pattern


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


## RemoveGroupApiKey

> BaseResponse RemoveGroupApiKey(ctx, groupID, apiKeyID).Execute()

Remove API key from group



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
	groupID := int64(789) // int64 | 
	apiKeyID := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.RemoveGroupApiKey(context.Background(), groupID, apiKeyID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.RemoveGroupApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveGroupApiKey`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.RemoveGroupApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupID** | **int64** |  | 
**apiKeyID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGroupApiKeyRequest struct via the builder pattern


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


## RemoveGroupRole

> BaseResponse RemoveGroupRole(ctx, groupID, roleShortName).Execute()

Remove role from group



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
	groupID := int64(789) // int64 | 
	roleShortName := "roleShortName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.RemoveGroupRole(context.Background(), groupID, roleShortName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.RemoveGroupRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveGroupRole`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.RemoveGroupRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupID** | **int64** |  | 
**roleShortName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGroupRoleRequest struct via the builder pattern


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


## RemoveGroupUser

> BaseResponse RemoveGroupUser(ctx, groupID, userID).Execute()

Remove user from group



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
	groupID := int64(789) // int64 | 
	userID := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.RemoveGroupUser(context.Background(), groupID, userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.RemoveGroupUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveGroupUser`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.RemoveGroupUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupID** | **int64** |  | 
**userID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGroupUserRequest struct via the builder pattern


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


## RemoveUserSignerCloudWallet

> BaseResponse RemoveUserSignerCloudWallet(ctx, userID, walletAddress).Execute()

Remove user cloud wallet signer



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
	userID := int64(789) // int64 | 
	walletAddress := "walletAddress_example" // string | An HSM ethereum address.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.RemoveUserSignerCloudWallet(context.Background(), userID, walletAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.RemoveUserSignerCloudWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveUserSignerCloudWallet`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.RemoveUserSignerCloudWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **int64** |  | 
**walletAddress** | **string** | An HSM ethereum address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserSignerCloudWalletRequest struct via the builder pattern


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


## RemoveUserSignerSafeAccount

> BaseResponse RemoveUserSignerSafeAccount(ctx, userID, walletAddress).Execute()

Remove user safe account signer



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
	userID := int64(789) // int64 | 
	walletAddress := "walletAddress_example" // string | An HSM ethereum address.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.RemoveUserSignerSafeAccount(context.Background(), userID, walletAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.RemoveUserSignerSafeAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveUserSignerSafeAccount`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.RemoveUserSignerSafeAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **int64** |  | 
**walletAddress** | **string** | An HSM ethereum address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserSignerSafeAccountRequest struct via the builder pattern


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


## RemoveUserSignerWeb3Wallet

> BaseResponse RemoveUserSignerWeb3Wallet(ctx, userID, walletAddress).Execute()

Remove user web3 wallet signer



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
	userID := int64(789) // int64 | 
	walletAddress := "walletAddress_example" // string | An HSM ethereum address.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.RemoveUserSignerWeb3Wallet(context.Background(), userID, walletAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.RemoveUserSignerWeb3Wallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveUserSignerWeb3Wallet`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.RemoveUserSignerWeb3Wallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **int64** |  | 
**walletAddress** | **string** | An HSM ethereum address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserSignerWeb3WalletRequest struct via the builder pattern


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


## SetUserSignerCloudWallet

> BaseResponse SetUserSignerCloudWallet(ctx, userID, walletAddress).Execute()

Add or update user cloud wallet signer



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
	userID := int64(789) // int64 | 
	walletAddress := "walletAddress_example" // string | An HSM ethereum address.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.SetUserSignerCloudWallet(context.Background(), userID, walletAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.SetUserSignerCloudWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetUserSignerCloudWallet`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.SetUserSignerCloudWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **int64** |  | 
**walletAddress** | **string** | An HSM ethereum address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetUserSignerCloudWalletRequest struct via the builder pattern


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


## SetUserSignerSafeAccount

> BaseResponse SetUserSignerSafeAccount(ctx, userID, walletAddress).SignerLabel(signerLabel).Execute()

Add or update user safe account signer



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
	userID := int64(789) // int64 | 
	walletAddress := "walletAddress_example" // string | An HSM ethereum address.
	signerLabel := *openapiclient.NewSignerLabel("Label_example") // SignerLabel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.SetUserSignerSafeAccount(context.Background(), userID, walletAddress).SignerLabel(signerLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.SetUserSignerSafeAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetUserSignerSafeAccount`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.SetUserSignerSafeAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **int64** |  | 
**walletAddress** | **string** | An HSM ethereum address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetUserSignerSafeAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **signerLabel** | [**SignerLabel**](SignerLabel.md) |  | 

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


## SetUserSignerWeb3Wallet

> BaseResponse SetUserSignerWeb3Wallet(ctx, userID, walletAddress).SignerLabel(signerLabel).Execute()

Add or update user web3 wallet signer



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
	userID := int64(789) // int64 | 
	walletAddress := "walletAddress_example" // string | An HSM ethereum address.
	signerLabel := *openapiclient.NewSignerLabel("Label_example") // SignerLabel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.SetUserSignerWeb3Wallet(context.Background(), userID, walletAddress).SignerLabel(signerLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.SetUserSignerWeb3Wallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetUserSignerWeb3Wallet`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.SetUserSignerWeb3Wallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **int64** |  | 
**walletAddress** | **string** | An HSM ethereum address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetUserSignerWeb3WalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **signerLabel** | [**SignerLabel**](SignerLabel.md) |  | 

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


## UpdateApiKey

> BaseResponse UpdateApiKey(ctx, apiKeyID).BaseAPIKey(baseAPIKey).Execute()

Update API key



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
	apiKeyID := int64(789) // int64 | 
	baseAPIKey := *openapiclient.NewBaseAPIKey("Label_example") // BaseAPIKey |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.UpdateApiKey(context.Background(), apiKeyID).BaseAPIKey(baseAPIKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.UpdateApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiKey`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.UpdateApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyID** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseAPIKey** | [**BaseAPIKey**](BaseAPIKey.md) |  | 

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

