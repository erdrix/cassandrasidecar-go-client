# RebuildOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Keyspace** | **string** | specific keyspace to rebuild, if not specified, all keyspaces are rebuilt  | [optional] 
**SourceDC** | **string** | name of DC from which to select sources for streaming, by default, pick any DC  | [optional] 
**SpecificTokens** | [**[]TokenRange**](TokenRange.md) | rebuild specific token ranges  | [optional] 
**SpecificSources** | **[]string** | specify hosts that this node should stream from when specificTokens are used  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


