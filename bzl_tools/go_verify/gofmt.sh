#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

export PATH="$PATH:`realpath external/rules_go~~go_sdk~go_sdk/bin/`"
gofmt="external/rules_go~~go_sdk~go_sdk/bin/gofmt"
if [[ ! -x "${gofmt}" ]]; then
  gofmt=$(which gofmt)
fi
golang="external/rules_go~~go_sdk~go_sdk/bin/go"
if [[ ! -x "${golang}" ]]; then
  golang=$(which go)
fi
export GOPATH=$(pwd -P)
mkdir -p $GOPATH
export GOCACHE="$(mktemp -d)"
trap "rm -rf '${GOCACHE}'" EXIT

diff=$(find . -name "*.go" | grep -v "\/external\/" | grep -v 'pkg/mod' | xargs "${gofmt}" -d 2>&1)
if [[ -n "${diff}" ]]; then
    echo "GO FMT DIFF"
    echo "${diff}"
    exit 1
fi

diff=$(find . -name "*.go" | grep -v "\/external\/" | grep -v 'pkg/mod' | xargs "${golang}" run golang.org/x/tools/cmd/goimports -local github.com/Metronome-Industries/usage-service -d)
if [[ -n "${diff}" ]]; then
    echo "GO IMPORTS DIFF"
    echo "${diff}"
    exit 1
fi