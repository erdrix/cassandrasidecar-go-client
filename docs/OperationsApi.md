# \OperationsApi

All URIs are relative to *http://localhost:4567*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OperationsGet**](OperationsApi.md#OperationsGet) | **Get** /operations | All operations of Sidecar
[**OperationsOperationIdGet**](OperationsApi.md#OperationsOperationIdGet) | **Get** /operations/{operationId} | abc
[**OperationsPost**](OperationsApi.md#OperationsPost) | **Post** /operations | Submits an operation to this Sidecar



## OperationsGet

> OperationsGet(ctx, optional)

All operations of Sidecar

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OperationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OperationsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**optional.Interface of []string**](string.md)| type of operations to filter on | 
 **status** | [**optional.Interface of []string**](string.md)| status of operations to filter on | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsOperationIdGet

> OperationsOperationIdGet(ctx, operationId)

abc

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | [**string**](.md)| ID of operation to return | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsPost

> OneOfCleanupOperationResponseDecommissionOperationResponseDrainOperationResponseFlushOperationResponseRebuildOperationResponseRefreshOperationResponseRestartOperationResponseScrubOperationResponseUpgradeSsTablesOperationResponseImportOperationResponseTruncateOperationResponseBackupOperationResponseRestoreOperationResponse OperationsPost(ctx, optional)

Submits an operation to this Sidecar

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OperationsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OperationsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**optional.Interface of UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**OneOfCleanupOperationResponseDecommissionOperationResponseDrainOperationResponseFlushOperationResponseRebuildOperationResponseRefreshOperationResponseRestartOperationResponseScrubOperationResponseUpgradeSsTablesOperationResponseImportOperationResponseTruncateOperationResponseBackupOperationResponseRestoreOperationResponse**](oneOf&lt;CleanupOperationResponse,DecommissionOperationResponse,DrainOperationResponse,FlushOperationResponse,RebuildOperationResponse,RefreshOperationResponse,RestartOperationResponse,ScrubOperationResponse,UpgradeSSTablesOperationResponse,ImportOperationResponse,TruncateOperationResponse,BackupOperationResponse,RestoreOperationResponse&gt;.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

