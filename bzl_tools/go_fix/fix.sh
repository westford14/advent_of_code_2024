#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

gofmt=$(realpath "external/rules_go~~go_sdk~go_sdk/bin/gofmt")
if [[ ! -x "${gofmt}" ]]; then
  gofmt=$(which gofmt)
fi
goimports=$(realpath "external/gazelle~~go_deps~org_golang_x_tools/cmd/goimports/goimports_/goimports")
if [[ ! -x "${goimports}" ]]; then
  goimports=$(which goimports)
fi

cd $BUILD_WORKSPACE_DIRECTORY

for f in $(git ls-files '*.go' | grep -v 'bzl_tools/tools.go') ; do
    "$gofmt" -l -w $f
    "$goimports" -local github.com/westford14/advent_of_code_2024 -w "$f"
done