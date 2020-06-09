#! /bin/bash

echo Generating Cassandra Sidecar Go client
export wv_go_path="module github.com\/erdrix\/cassandrasidecar-go-client"
export workspace_dir=$(pwd)
export wv_tmp_dir=${workspace_dir}/pkg
rm -fr ${workspace_dir}/pkg && mkdir -p ${workspace_dir}/pkg/cassandrasidecar
cd resources/client_gen
./generate_api_client.sh

cd ../../

echo Cleaning Cassandra Sidecar Go client
mv ${wv_tmp_dir}/cassandrasidecar/go-client/* ${wv_tmp_dir}/cassandrasidecar && rm -fr ${wv_tmp_dir}/cassandrasidecar/go-client
rm -f ${wv_tmp_dir}/cassandrasidecar/go.sum
rm -f ${wv_tmp_dir}/cassandrasidecar/git_push.sh

echo Moving documentations files to root dirs
rm -fr docs && mv ${wv_tmp_dir}/cassandrasidecar/docs .
rm -fr README.md && mv ${wv_tmp_dir}/cassandrasidecar/README.md .
rm -fr go.mod && mv ${wv_tmp_dir}/cassandrasidecar/go.mod .
sed -i.bak "1s/.*/$wv_go_path/" go.mod && echo "go 1.13" >> go.mod
