# RestoreOperationResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | type of operation, one has to set it to &#x27;restore&#x27; in case he wants this request to be considered as a backup one  | [default to null]
**StorageLocation** | **string** | similar to field in backup request but used for telling from where files should be downloaded, not uploaded,  in case globalRequest field is set to true, it does not matter what dc and node id is used, these components  in storageLocation path will be automatically changed.  | [default to null]
**ConcurrentConnections** | **int32** | similar to field in backup request but used for downloading files, not uploading them  | [optional] [default to null]
**LockFile** | **string** | similar to field in backup request  | [optional] [default to null]
**CassandraDirectory** | **string** | similar to field in backup request  | [optional] [default to null]
**CassandraConfigDirectory** | **string** | directory where one expects to find &#x27;conf/cassandra.yaml&#x27; file in case we need to update it with initial tokens  in case restoration strategy is IN_PLACE.  | [optional] [default to null]
**RestoreSystemKeyspace** | **bool** | a flag saying if we should restore system keyspaces as well, relevant only for IN_PLACE restoration  | [optional] [default to null]
**SnapshotTag** | **string** | name of snapshot to restore  | [default to null]
**Entities** | **string** | similar to field in backup request, when empty, all entities in given snapshot will be restored  | [optional] [default to null]
**UpdateCassandraYaml** | **bool** | flag telling if cassandra.yaml should be updated with initial_tokens, relevant only in case of IN_PLACE strategy  | [optional] [default to null]
**RestorationStrategyType** | **string** | strategy telling how we should go about restoration, please refer to details in backup and sidecar documentation  | [optional] [default to null]
**RestorationPhase** | **string** | phase telling what should we do, this field has to be set just once as DOWNLOAD if globalRequest if true and  coordinator of that request will take care of all other phases automatically on its own  | [optional] [default to null]
**NoDeleteTruncates** | **bool** | flag saying if we should not delete truncated SSTables after they are imported, as part of CLEANUP phase, defaults to false  | [optional] [default to null]
**NoDeleteDownloads** | **bool** | flag saying if we should not delete downloaded SSTables from remote location, as part of CLEANUP phase, defaults to false  | [optional] [default to null]
**NoDownloadData** | **bool** | flag saying if we should not download data from remote location as we expect them to be there already, defaults to false,  setting this to true has sense only in case noDeleteDownloads was set to true in previous restoration requests  | [optional] [default to null]
**SchemaVersion** | **string** | version of schema we want to restore from, upon backup, a schema version is automatically appended to snapshot name and  its manifest is uploaded under that name. In case we have two snapshots having same name, we might distinguish between them by this schema version. If schema version is not specified, we expect that there will be one and only one backup taken with  respective snapshot name. This schema version has to match the version of a Cassandra node we are doing restore for (hence,  by proxy, when global request mode is used, all nodes have to be on exact same schema version).  | [optional] [default to null]
**ExactSchemaVersion** | **bool** | flag saying if we indeed want a schema version of a running node match with schema version a snapshot is taken on. there might be cases when we want to restore a table for which its CQL schema has not changed but it has changed for  other table / keyspace but a schema for that node has changed by doing that.  | [optional] [default to null]
**Import_** | [***Object**](.md) |  | [optional] [default to null]
**GlobalRequest** | **bool** | flag saying that this request is a global one, meaning that a Sidecar this request is sent to will act as a restoration  coordinator sending all other requests to each node in a cluster, for each phase.  | [optional] [default to null]
**Id** | **string** | unique identifier of an operation, a random id is assigned to each operation after  a request is submitted, from caller&#x27;s perspective, an id is sent back as a response  to his request so he can further query state of that operation, referencing id, by operations/{id} endpoint  | [default to null]
**State** | **string** | state of an operation, operation might be in various states, PENDING - this operation is pending for being submitted. RUNNING - this operation is actively  doing its job, COMPLETED - this operation has finished successfully, CANCELLED -  this operation was interrupted while being run, FAILED - this operation has  finished errorneously  | [default to null]
**Progress** | **string** | float from 0.0 to 1.0, 1.0 telling that operation is completed,  either successfully or with errors.  | [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | timestamp telling when this operation was created on Sidecar&#x27;s side  | [default to null]
**StartTime** | [**time.Time**](time.Time.md) | timestamp telling when this operation was started by Sidecar, if an operation  is created, it does not necessarily mean that it will be started right away,  in most cases it is the case but if e.g. ExecutorService is full on its working thread, an execution of an operation is postponed and start time is updated only after that  | [optional] [default to null]
**CompletionTime** | [**time.Time**](time.Time.md) | timestamp telling when an operation has finished, irrelevant of its result, an  operation can be failed and it would still have this field populated.  | [optional] [default to null]
**FailureCause** | [***interface{}**](interface{}.md) | This field contains serialized java.lang.Throwable in case this operation has failed  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

