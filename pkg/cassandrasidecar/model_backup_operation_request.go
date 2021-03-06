/*
 * Instaclustr Cassandra Sidecar
 *
 * REST API for Cassandra Sidecar from Instaclustr
 *
 * API version: 1.1.0
 * Contact: support@instaclustr.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cassandrasidecar

type BackupOperationRequest struct {
	// type of operation, one has to set it to 'backup' in case he wants this  request to be considered as a backup one 
	Type_ string `json:"type"`
	// location where SSTables will be uploaded. A value of the storageLocation property  has to have exact format which is 'protocol://bucket/clusterName/dcName/nodeName'. protocol is either 'gcp', 's3', 'azure' or 'file:/'. For global requests, dcName and  nodeName are changed automatically as these values are read from Cassandra and  storageLocation is updated automatically for every node a specific backup request  will be submitted to so the value of dcName and nodeName is irrelevant for global requests  as they will be modified every time, a bucket does not need to exist, it will be created automatically if it does not, clusterName has to be specified. There might be automatic  resolution of clusterName in the future however for now, one has to supply this  property on his own. 
	StorageLocation string `json:"storageLocation"`
	// directory of Cassandra, by default it is /var/lib/cassandra, in this path, one expects there is 'data' directory 
	CassandraDirectory string `json:"cassandraDirectory,omitempty"`
	// Based on this field, there will be throughtput per second computed based on what size  data we want to upload we have. The formula is \"size / duration\". The lower the duration is,  the higher throughput per second we will need and vice versa. This will influence e.g. responsiveness  of a node to its business requests so one can control how much bandwidth is used for backup  purposes in case a cluster is fully operational. The format of this field is \"amount unit\". 'unit' is just a (case-insensitive) java.util.concurrent.TimeUnit enum value. If not used, there will not be any restrictions as how fast an upload can be. 
	Duration string `json:"duration,omitempty"`
	Bandwidth *DataRate `json:"bandwidth,omitempty"`
	// number of threads used for upload, there might be at most so many uploading threads at any given time,  when not set, it defaults to 10 
	ConcurrentConnections int32 `json:"concurrentConnections,omitempty"`
	// path to file which will be used for locking the critical logic dealing with backups, this locking is done  by locking a file on a file system so other execution will not proceed until the former one has finished. By  default, this path is System.getProperty(\"java.io.tmpdir\") + \"/global-transfer-lock\". 
	LockFile string `json:"lockFile,omitempty"`
	// name of snapshot to make so this snapshot will be uploaded to storage location. If not specified, the  name of snapshot will be automatically generated and it will have name 'autosnap-milliseconds-since-epoch' 
	SnapshotTag string `json:"snapshotTag,omitempty"`
	// name of datacenter to backup, nodes in the other datacenter(s) will not be involved 
	Dc string `json:"dc,omitempty"`
	// database entities to backup, it might be either only keyspaces or only tables (from different keyspaces if needed), e.g. 'k1,k2' if one wants to backup whole keyspaces and 'ks1.t1,ks2,t2' if one wants to backup tables. These formats can not be used together so 'k1,k2.t2' is invalid. If this field is empty, all keyspaces are backed up.
	Entities string `json:"entities,omitempty"`
	// name of Kubernetes namespace to fetch Kubernetes secret for backups from, when not specified, it defaults to 'default' 
	K8sNamespace string `json:"k8sNamespace,omitempty"`
	// name of Kubernetes secret from which credentials used for the communication to cloud  storage providers are read, if not specified, secret name to be read will be automatically  derived in form 'cassandra-backup-restore-secret-cluster-{name-of-cluster}'. These secrets are used only in case  protocol in storageLocation is gcp, azure or s3. 
	K8sBackupSecretName string `json:"k8sBackupSecretName,omitempty"`
	// flag saying if this request is meant to be global or not, once a global backup request is submitted to Sidecar, it will coordinate backup for all other nodes in a cluster (including itself) so from a point of view of a caller,  one might just backup whole cluster by one request and repeatedly query its status based on returned operation id. 
	GlobalRequest bool `json:"globalRequest,omitempty"`
}
