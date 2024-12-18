#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

gofmt="external/rules_go~~go_sdk~go_sdk/bin/gofmt"
if [[ ! -x "${gofmt}" ]]; then
  gofmt=$(which gofmt)
fi
goimports="external/gazelle~~go_deps~org_golang_x_tools/cmd/goimports/goimports_/goimports"
if [[ ! -x "${goimports}" ]]; then
  goimports=$(which goimports)
fi
gopath=$(realpath external/rules_go~~go_sdk~go_sdk/bin)
export PATH="${gopath}:$PATH"

diff=$(find . -name "*.go" | grep -v "\/external\/" | grep -v 'pkg/mod' | xargs "${gofmt}" -d 2>&1)
if [[ -n "${diff}" ]]; then
    echo "GO FMT DIFF"
    echo "${diff}"
    exit 1
fi

diff=$(find . -name "*.go" | grep -v "\/external\/" | grep -v 'pkg/mod' | xargs "${goimports}" -local github.com/westford14/advent_of_code_2024 -d)
if [[ -n "${diff}" ]]; then
    echo "GO IMPORTS DIFF"
    echo "${diff}"
    exit 1
fi