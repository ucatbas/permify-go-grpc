# client\DataAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BundleRun**](DataAPI.md#BundleRun) | **Post** /v1/tenants/{tenant_id}/data/run-bundle | run bundle
[**DataAttributesRead**](DataAPI.md#DataAttributesRead) | **Post** /v1/tenants/{tenant_id}/data/attributes/read | read attributes
[**DataDelete**](DataAPI.md#DataDelete) | **Post** /v1/tenants/{tenant_id}/data/delete | delete data
[**DataRelationshipsRead**](DataAPI.md#DataRelationshipsRead) | **Post** /v1/tenants/{tenant_id}/data/relationships/read | read relationships
[**DataWrite**](DataAPI.md#DataWrite) | **Post** /v1/tenants/{tenant_id}/data/write | write data
[**RelationshipsDelete**](DataAPI.md#RelationshipsDelete) | **Post** /v1/tenants/{tenant_id}/relationships/delete | delete relationships
[**RelationshipsWrite**](DataAPI.md#RelationshipsWrite) | **Post** /v1/tenants/{tenant_id}/relationships/write | write relationships



## BundleRun

> BundleRunResponse BundleRun(ctx, tenantId).Body(body).Execute()

run bundle

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
	tenantId := "tenantId_example" // string | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
	body := *openapiclient.NewRunBundleBody() // RunBundleBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.BundleRun(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.BundleRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BundleRun`: BundleRunResponse
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.BundleRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant &lt;code&gt;t1&lt;/code&gt; for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBundleRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RunBundleBody**](RunBundleBody.md) |  | 

### Return type

[**BundleRunResponse**](BundleRunResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataAttributesRead

> AttributeReadResponse DataAttributesRead(ctx, tenantId).Body(body).Execute()

read attributes

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
	tenantId := "tenantId_example" // string | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
	body := *openapiclient.NewReadAttributesBody() // ReadAttributesBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.DataAttributesRead(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.DataAttributesRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataAttributesRead`: AttributeReadResponse
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.DataAttributesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant &lt;code&gt;t1&lt;/code&gt; for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataAttributesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ReadAttributesBody**](ReadAttributesBody.md) |  | 

### Return type

[**AttributeReadResponse**](AttributeReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataDelete

> DataDeleteResponse DataDelete(ctx, tenantId).Body(body).Execute()

delete data

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
	tenantId := "tenantId_example" // string | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
	body := *openapiclient.NewDataDeleteBody() // DataDeleteBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.DataDelete(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.DataDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataDelete`: DataDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.DataDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant &lt;code&gt;t1&lt;/code&gt; for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DataDeleteBody**](DataDeleteBody.md) |  | 

### Return type

[**DataDeleteResponse**](DataDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataRelationshipsRead

> RelationshipReadResponse DataRelationshipsRead(ctx, tenantId).Body(body).Execute()

read relationships

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
	tenantId := "tenantId_example" // string | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
	body := *openapiclient.NewReadRelationshipsBody() // ReadRelationshipsBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.DataRelationshipsRead(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.DataRelationshipsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataRelationshipsRead`: RelationshipReadResponse
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.DataRelationshipsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant &lt;code&gt;t1&lt;/code&gt; for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataRelationshipsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ReadRelationshipsBody**](ReadRelationshipsBody.md) |  | 

### Return type

[**RelationshipReadResponse**](RelationshipReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataWrite

> DataWriteResponse DataWrite(ctx, tenantId).Body(body).Execute()

write data

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
	tenantId := "tenantId_example" // string | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
	body := *openapiclient.NewDataWriteBody() // DataWriteBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.DataWrite(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.DataWrite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataWrite`: DataWriteResponse
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.DataWrite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant &lt;code&gt;t1&lt;/code&gt; for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataWriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DataWriteBody**](DataWriteBody.md) |  | 

### Return type

[**DataWriteResponse**](DataWriteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RelationshipsDelete

> RelationshipDeleteResponse RelationshipsDelete(ctx, tenantId).Body(body).Execute()

delete relationships

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
	tenantId := "tenantId_example" // string | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
	body := *openapiclient.NewDeleteRelationshipsBody() // DeleteRelationshipsBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.RelationshipsDelete(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.RelationshipsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelationshipsDelete`: RelationshipDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.RelationshipsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant &lt;code&gt;t1&lt;/code&gt; for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRelationshipsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DeleteRelationshipsBody**](DeleteRelationshipsBody.md) |  | 

### Return type

[**RelationshipDeleteResponse**](RelationshipDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RelationshipsWrite

> RelationshipWriteResponse RelationshipsWrite(ctx, tenantId).Body(body).Execute()

write relationships

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
	tenantId := "tenantId_example" // string | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant <code>t1</code> for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes.
	body := *openapiclient.NewWriteRelationshipsBody() // WriteRelationshipsBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAPI.RelationshipsWrite(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.RelationshipsWrite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelationshipsWrite`: RelationshipWriteResponse
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.RelationshipsWrite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant &lt;code&gt;t1&lt;/code&gt; for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRelationshipsWriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**WriteRelationshipsBody**](WriteRelationshipsBody.md) |  | 

### Return type

[**RelationshipWriteResponse**](RelationshipWriteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

