# client\PermissionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermissionsCheck**](PermissionAPI.md#PermissionsCheck) | **Post** /v1/tenants/{tenant_id}/permissions/check | check api
[**PermissionsExpand**](PermissionAPI.md#PermissionsExpand) | **Post** /v1/tenants/{tenant_id}/permissions/expand | expand api
[**PermissionsLookupEntity**](PermissionAPI.md#PermissionsLookupEntity) | **Post** /v1/tenants/{tenant_id}/permissions/lookup-entity | lookup entity
[**PermissionsLookupEntityStream**](PermissionAPI.md#PermissionsLookupEntityStream) | **Post** /v1/tenants/{tenant_id}/permissions/lookup-entity-stream | lookup entity stream
[**PermissionsLookupSubject**](PermissionAPI.md#PermissionsLookupSubject) | **Post** /v1/tenants/{tenant_id}/permissions/lookup-subject | lookup-subject
[**PermissionsSubjectPermission**](PermissionAPI.md#PermissionsSubjectPermission) | **Post** /v1/tenants/{tenant_id}/permissions/subject-permission | subject permission



## PermissionsCheck

> PermissionCheckResponse PermissionsCheck(ctx, tenantId).Body(body).Execute()

check api

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
	body := *openapiclient.NewCheckBody() // CheckBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionAPI.PermissionsCheck(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionAPI.PermissionsCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionsCheck`: PermissionCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `PermissionAPI.PermissionsCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant &lt;code&gt;t1&lt;/code&gt; for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CheckBody**](CheckBody.md) |  | 

### Return type

[**PermissionCheckResponse**](PermissionCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsExpand

> PermissionExpandResponse PermissionsExpand(ctx, tenantId).Body(body).Execute()

expand api

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
	body := *openapiclient.NewPermissionExpandBody() // PermissionExpandBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionAPI.PermissionsExpand(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionAPI.PermissionsExpand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionsExpand`: PermissionExpandResponse
	fmt.Fprintf(os.Stdout, "Response from `PermissionAPI.PermissionsExpand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant &lt;code&gt;t1&lt;/code&gt; for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsExpandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PermissionExpandBody**](PermissionExpandBody.md) |  | 

### Return type

[**PermissionExpandResponse**](PermissionExpandResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsLookupEntity

> PermissionLookupEntityResponse PermissionsLookupEntity(ctx, tenantId).Body(body).Execute()

lookup entity

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
	body := *openapiclient.NewLookupEntityBody() // LookupEntityBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionAPI.PermissionsLookupEntity(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionAPI.PermissionsLookupEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionsLookupEntity`: PermissionLookupEntityResponse
	fmt.Fprintf(os.Stdout, "Response from `PermissionAPI.PermissionsLookupEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant &lt;code&gt;t1&lt;/code&gt; for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsLookupEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LookupEntityBody**](LookupEntityBody.md) |  | 

### Return type

[**PermissionLookupEntityResponse**](PermissionLookupEntityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsLookupEntityStream

> StreamResultOfPermissionLookupEntityStreamResponse PermissionsLookupEntityStream(ctx, tenantId).Body(body).Execute()

lookup entity stream

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
	body := *openapiclient.NewLookupEntityStreamBody() // LookupEntityStreamBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionAPI.PermissionsLookupEntityStream(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionAPI.PermissionsLookupEntityStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionsLookupEntityStream`: StreamResultOfPermissionLookupEntityStreamResponse
	fmt.Fprintf(os.Stdout, "Response from `PermissionAPI.PermissionsLookupEntityStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant &lt;code&gt;t1&lt;/code&gt; for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsLookupEntityStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LookupEntityStreamBody**](LookupEntityStreamBody.md) |  | 

### Return type

[**StreamResultOfPermissionLookupEntityStreamResponse**](StreamResultOfPermissionLookupEntityStreamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsLookupSubject

> PermissionLookupSubjectResponse PermissionsLookupSubject(ctx, tenantId).Body(body).Execute()

lookup-subject

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
	body := *openapiclient.NewLookupSubjectBody() // LookupSubjectBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionAPI.PermissionsLookupSubject(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionAPI.PermissionsLookupSubject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionsLookupSubject`: PermissionLookupSubjectResponse
	fmt.Fprintf(os.Stdout, "Response from `PermissionAPI.PermissionsLookupSubject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant &lt;code&gt;t1&lt;/code&gt; for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsLookupSubjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LookupSubjectBody**](LookupSubjectBody.md) |  | 

### Return type

[**PermissionLookupSubjectResponse**](PermissionLookupSubjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsSubjectPermission

> PermissionSubjectPermissionResponse PermissionsSubjectPermission(ctx, tenantId).Body(body).Execute()

subject permission

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
	body := *openapiclient.NewSubjectPermissionBody() // SubjectPermissionBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionAPI.PermissionsSubjectPermission(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionAPI.PermissionsSubjectPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionsSubjectPermission`: PermissionSubjectPermissionResponse
	fmt.Fprintf(os.Stdout, "Response from `PermissionAPI.PermissionsSubjectPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Identifier of the tenant, if you are not using multi-tenancy (have only one tenant) use pre-inserted tenant &lt;code&gt;t1&lt;/code&gt; for this field. Required, and must match the pattern \\“[a-zA-Z0-9-,]+\\“, max 64 bytes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsSubjectPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SubjectPermissionBody**](SubjectPermissionBody.md) |  | 

### Return type

[**PermissionSubjectPermissionResponse**](PermissionSubjectPermissionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

