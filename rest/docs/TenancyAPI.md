# client\TenancyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenantsCreate**](TenancyAPI.md#TenantsCreate) | **Post** /v1/tenants/create | create tenant
[**TenantsDelete**](TenancyAPI.md#TenantsDelete) | **Delete** /v1/tenants/{id} | delete tenant
[**TenantsList**](TenancyAPI.md#TenantsList) | **Post** /v1/tenants/list | list tenants



## TenantsCreate

> TenantCreateResponse TenantsCreate(ctx).Body(body).Execute()

create tenant

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ucatbas/permify-go-grpc/rest"
)

func main() {
	body := *openapiclient.NewTenantCreateRequest() // TenantCreateRequest | TenantCreateRequest is the message used for the request to create a tenant.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenantsCreate(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenantsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantsCreate`: TenantCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenantsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantCreateRequest**](TenantCreateRequest.md) | TenantCreateRequest is the message used for the request to create a tenant. | 

### Return type

[**TenantCreateResponse**](TenantCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsDelete

> TenantDeleteResponse TenantsDelete(ctx, id).Execute()

delete tenant

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ucatbas/permify-go-grpc/rest"
)

func main() {
	id := "id_example" // string | id is the unique identifier of the tenant to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenantsDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenantsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantsDelete`: TenantDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenantsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id is the unique identifier of the tenant to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TenantDeleteResponse**](TenantDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsList

> TenantListResponse TenantsList(ctx).Body(body).Execute()

list tenants

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ucatbas/permify-go-grpc/rest"
)

func main() {
	body := *openapiclient.NewTenantListRequest() // TenantListRequest | TenantListRequest is the message used for the request to list all tenants.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenancyAPI.TenantsList(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenancyAPI.TenantsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantsList`: TenantListResponse
	fmt.Fprintf(os.Stdout, "Response from `TenancyAPI.TenantsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantListRequest**](TenantListRequest.md) | TenantListRequest is the message used for the request to list all tenants. | 

### Return type

[**TenantListResponse**](TenantListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

