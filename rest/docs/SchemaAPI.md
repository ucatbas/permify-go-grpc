# client\SchemaAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SchemasList**](SchemaAPI.md#SchemasList) | **Post** /v1/tenants/{tenant_id}/schemas/list | list schema
[**SchemasPartialWrite**](SchemaAPI.md#SchemasPartialWrite) | **Patch** /v1/tenants/{tenant_id}/schemas/partial-write | partially update your authorization model
[**SchemasRead**](SchemaAPI.md#SchemasRead) | **Post** /v1/tenants/{tenant_id}/schemas/read | read schema
[**SchemasWrite**](SchemaAPI.md#SchemasWrite) | **Post** /v1/tenants/{tenant_id}/schemas/write | write schema



## SchemasList

> SchemaListResponse SchemasList(ctx, tenantId).Body(body).Execute()

list schema

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
	body := *openapiclient.NewSchemaListBody() // SchemaListBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchemaAPI.SchemasList(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchemaAPI.SchemasList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SchemasList`: SchemaListResponse
	fmt.Fprintf(os.Stdout, "Response from `SchemaAPI.SchemasList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant &lt;code&gt;t1&lt;/code&gt; for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemasListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SchemaListBody**](SchemaListBody.md) |  | 

### Return type

[**SchemaListResponse**](SchemaListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasPartialWrite

> SchemaPartialWriteResponse SchemasPartialWrite(ctx, tenantId).Body(body).Execute()

partially update your authorization model

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
	tenantId := "tenantId_example" // string | tenant_id is a string that identifies the tenant. It must match the pattern \"[a-zA-Z0-9-,]+\", be a maximum of 64 bytes, and must not be empty.
	body := *openapiclient.NewPartialWriteBody() // PartialWriteBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchemaAPI.SchemasPartialWrite(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchemaAPI.SchemasPartialWrite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SchemasPartialWrite`: SchemaPartialWriteResponse
	fmt.Fprintf(os.Stdout, "Response from `SchemaAPI.SchemasPartialWrite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | tenant_id is a string that identifies the tenant. It must match the pattern \&quot;[a-zA-Z0-9-,]+\&quot;, be a maximum of 64 bytes, and must not be empty. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemasPartialWriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PartialWriteBody**](PartialWriteBody.md) |  | 

### Return type

[**SchemaPartialWriteResponse**](SchemaPartialWriteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasRead

> SchemaReadResponse SchemasRead(ctx, tenantId).Body(body).Execute()

read schema

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
	body := *openapiclient.NewSchemaReadBody() // SchemaReadBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchemaAPI.SchemasRead(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchemaAPI.SchemasRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SchemasRead`: SchemaReadResponse
	fmt.Fprintf(os.Stdout, "Response from `SchemaAPI.SchemasRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant &lt;code&gt;t1&lt;/code&gt; for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemasReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SchemaReadBody**](SchemaReadBody.md) |  | 

### Return type

[**SchemaReadResponse**](SchemaReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemasWrite

> SchemaWriteResponse SchemasWrite(ctx, tenantId).Body(body).Execute()

write schema

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
	body := *openapiclient.NewSchemaWriteBody() // SchemaWriteBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchemaAPI.SchemasWrite(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchemaAPI.SchemasWrite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SchemasWrite`: SchemaWriteResponse
	fmt.Fprintf(os.Stdout, "Response from `SchemaAPI.SchemasWrite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant &lt;code&gt;t1&lt;/code&gt; for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemasWriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SchemaWriteBody**](SchemaWriteBody.md) |  | 

### Return type

[**SchemaWriteResponse**](SchemaWriteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

