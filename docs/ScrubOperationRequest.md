# ScrubOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Keyspace** | **string** | keyspace to scrub  | 
**Tables** | **[]string** | tables to scrub, empty or not provided will scrub all tables in respective keyspace  | [optional] 
**Jobs** | **int32** | number of sstables to scrub simultanously, set to 0 to use all available compaction threads  | [optional] 
**DisableSnapshot** | **bool** | scrubbed CFs will be snapshotted first, defaults to false  | [optional] 
**SkipCorrupted** | **bool** | skip corrupted partitions even when scrubbing counter tables, defaults to false  | [optional] 
**NoValidate** | **bool** | do not validate columns using column validator, defaults to false  | [optional] 
**ReinsertOverflowedTTL** | **bool** | Rewrites rows with overflowed expiration date affected by CASSANDRA-14092 with the maximum supported expiration date of 2038-01-19T03:14:06+00:00. The rows are rewritten with the original timestamp incremented by one millisecond to override/supersede any potential tombstone that may have been generated during compaction of the affected rows.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


