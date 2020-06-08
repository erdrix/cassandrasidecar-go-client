#! /bin/bash

echo Generating Cassandra Sidecar Go client
export workspace_dir=$(pwd)
export wv_tmp_dir=${workspace_dir}/pkg
cd resources/client_gen
./generate_api_client.sh

cd ../../

echo Cleaning Cassandra Sidecar Go client
rm -fr ${wv_tmp_dir}/nifi/.swagger-codegen \
  && rm -fr ${wv_tmp_dir}/cassandrasidecar/.gitignore \
  && rm -fr ${wv_tmp_dir}/cassandrasidecar/.swagger-codegen-ignore \
  && rm -fr ${wv_tmp_dir}/cassandrasidecar/.travis.yml \
  && rm -fr ${wv_tmp_dir}/swagger-codegen-cli-*.jar \
  && rm -fr ${wv_tmp_dir}/cassandrasidecar.conf.json

echo Moving documentations files to root dirs
rm -fr ${wv_tmp_dir}/docs && mv ${wv_tmp_dir}/cassandrasidecar/docs .
rm -fr ${wv_tmp_dir}/README.md && mv ${wv_tmp_dir}/cassandrasidecar/README.md .