# CleanupOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Keyspace** | **string** | keyspace to cleanup  | 
**Tables** | **[]string** | tables to cleanup, when not specified, all tables in a keyspace will be cleaned up  | [optional] 
**Jobs** | **int32** | number of jobs to use, never uses more that concurrent_compactor threads  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


