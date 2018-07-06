#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

vendor/k8s.io/code-generator/generate-groups.sh \
deepcopy \
github.com/sense12/sense12-operator/pkg/generated \
github.com/sense12/sense12-operator/pkg/apis \
sense12:v1 \
--go-header-file "./tmp/codegen/boilerplate.go.txt"
