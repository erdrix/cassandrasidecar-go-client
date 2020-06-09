# UpgradeSsTablesOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Keyspace** | **string** | keyspace to upgrade SSTables of  | 
**Tables** | **[]string** | an array of tables to upgrade SSTables of, empty or not provided array will default to upgrading of SSTables of all tables in respective keyspace  | [optional] 
**Jobs** | **int32** | the number of threads to use - 0 means use all available, it never uses more than concurrent_compactor threads  | [optional] 
**IncludeAllSStables** | **bool** | include all sstables, even those already on the current version, defaults to false | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


