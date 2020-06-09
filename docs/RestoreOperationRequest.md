# RestoreOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | type of operation, one has to set it to &#39;restore&#39; in case he wants this request to be considered as a backup one  | 
**StorageLocation** | **string** | similar to field in backup request but used for telling from where files should be downloaded, not uploaded, in case globalRequest field is set to true, it does not matter what dc and node id is used, these components in storageLocation path will be automatically changed.  | 
**ConcurrentConnections** | **int32** | similar to field in backup request but used for downloading files, not uploading them  | [optional] 
**LockFile** | **string** | similar to field in backup request  | [optional] 
**CassandraDirectory** | **string** | similar to field in backup request  | [optional] 
**CassandraConfigDirectory** | **string** | directory where one expects to find &#39;conf/cassandra.yaml&#39; file in case we need to update it with initial tokens in case restoration strategy is IN_PLACE.  | [optional] 
**RestoreSystemKeyspace** | **bool** | a flag saying if we should restore system keyspaces as well, relevant only for IN_PLACE restoration  | [optional] 
**SnapshotTag** | **string** | name of snapshot to restore  | 
**Entities** | **string** | similar to field in backup request, when empty, all entities in given snapshot will be restored  | [optional] 
**UpdateCassandraYaml** | **bool** | flag telling if cassandra.yaml should be updated with initial_tokens, relevant only in case of IN_PLACE strategy  | [optional] 
**RestorationStrategyType** | **string** | strategy telling how we should go about restoration, please refer to details in backup and sidecar documentation  | [optional] 
**RestorationPhase** | **string** | phase telling what should we do, this field has to be set just once as DOWNLOAD if globalRequest if true and coordinator of that request will take care of all other phases automatically on its own  | [optional] 
**NoDeleteTruncates** | **bool** | flag saying if we should not delete truncated SSTables after they are imported, as part of CLEANUP phase, defaults to false  | [optional] 
**NoDeleteDownloads** | **bool** | flag saying if we should not delete downloaded SSTables from remote location, as part of CLEANUP phase, defaults to false  | [optional] 
**NoDownloadData** | **bool** | flag saying if we should not download data from remote location as we expect them to be there already, defaults to false, setting this to true has sense only in case noDeleteDownloads was set to true in previous restoration requests  | [optional] 
**SchemaVersion** | **string** | version of schema we want to restore from, upon backup, a schema version is automatically appended to snapshot name and its manifest is uploaded under that name. In case we have two snapshots having same name, we might distinguish between them by this schema version. If schema version is not specified, we expect that there will be one and only one backup taken with respective snapshot name. This schema version has to match the version of a Cassandra node we are doing restore for (hence, by proxy, when global request mode is used, all nodes have to be on exact same schema version).  | [optional] 
**ExactSchemaVersion** | **bool** | flag saying if we indeed want a schema version of a running node match with schema version a snapshot is taken on. there might be cases when we want to restore a table for which its CQL schema has not changed but it has changed for other table / keyspace but a schema for that node has changed by doing that.  | [optional] 
**Import** | [**ImportOperationRequest**](ImportOperationRequest.md) |  | [optional] 
**GlobalRequest** | **bool** | flag saying that this request is a global one, meaning that a Sidecar this request is sent to will act as a restoration coordinator sending all other requests to each node in a cluster, for each phase.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


