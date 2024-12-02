#!/bin/bash
set -e

export PATH="$PATH:`realpath external/rules_go~~go_sdk~go_sdk/bin/`"
gofmt=$(realpath "external/rules_go~~go_sdk~go_sdk/bin/gofmt")
if [[ ! -x "${gofmt}" ]]; then
  gofmt=$(which gofmt)
fi
golang=$(realpath "external/rules_go~~go_sdk~go_sdk/bin/go")
if [[ ! -x "${golang}" ]]; then
  golang=$(which go)
fi
export GOPATH=$(pwd -P)
mkdir -p $GOPATH
export GOCACHE="$(mktemp -d)"
trap "rm -rf '${GOCACHE}'" EXIT

cd $BUILD_WORKSPACE_DIRECTORY

for f in $(git ls-files '*.go' | grep -v 'bzl_tools/tools.go') ; do
    "$gofmt" -l -w $f
    "$golang" run golang.org/x/tools/cmd/goimports -local github.com/Metronome-Industries/usage-service -w "$f"
done