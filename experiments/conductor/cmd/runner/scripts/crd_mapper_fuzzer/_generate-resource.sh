#!/bin/bash
# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


set -o errexit
set -o nounset
set -o pipefail
set -x

# TODO: Can we infer / default RESOURCE from PROTO_MESSAGE
export SERVICE=composer
export RESOURCE=environment
export CRD_KIND=ComposerEnvironment
export CRD_VERSION=v1alpha1
export CRD_GROUP=composer.cnrm.cloud.google.com
export PROTO_PACKAGE=./third_party/googleapis/google/cloud/orchestration/airflow/service/v1/environments.proto
export WORKDIR=/usr/local/google/home/maqiuyu/go/src/github.com/maqiuyujoyce/3-k8s-config-connector
# This time the PROTO_SERVICE is the prefix of the PROTO_MESSAGE.
export PROTO_SERVICE=google.cloud.orchestration.airflow.service.v1
export PROTO_RESOURCE=Environment
export PROTO_MESSAGE=google.cloud.orchestration.airflow.service.v1.Environment
export PATH=${PATH}:/usr/local/google/home/maqiuyu/go/src/github.com/maqiuyujoyce/2-k8s-config-connector/dev/tools/controllerbuilder/cmd/codebot:/usr/local/google/home/maqiuyu/go/src/github.com/maqiuyujoyce/2-k8s-config-connector/dev/tools/controllerbuilder

export BRANCH_NAME=types_${SERVICE}_${CRD_KIND}
export LOG_DIR=/tmp/conductor/${BRANCH_NAME}

#./01-write-generator-script.sh
#
#./02-generate-spec-and-status.sh

./03-generate-fuzzer.sh


cat <<EOF

Workflow is now to iterate:

apis/dataproc/v1beta1/generate.sh  && dev/tasks/generate-crds 
go run -mod=readonly golang.org/x/tools/cmd/goimports@latest -w  pkg/controller/direct/dataproc/
go test -v ./pkg/fuzztesting/fuzztests/ -fuzz=FuzzAllMappers -fuzztime 600s
<Make changes so fuzzer passes>

If the CRD already exists we need to make sure that there are only description changes in 

git diff origin/master -- config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_dataprocclusters.dataproc.cnrm.cloud.google.com.yaml

diff -u3 \
<(git show origin/master:config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_dataprocclusters.dataproc.cnrm.cloud.google.com.yaml | crd-tools remove-descriptions) \
<(cat config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_dataprocclusters.dataproc.cnrm.cloud.google.com.yaml | crd-tools remove-descriptions) | less
EOF

