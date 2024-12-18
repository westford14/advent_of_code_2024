# These should mostly be convenience targets for bazel commands
.PHONY: fix
fix:
	bazel run //bzl_tools/go_fix:go_fix
	bazel run //:buildifier
	bazel run //bzl_tools/protofmt:fmt
	bazel run //:gazelle

.PHONY: gen
gen:
	bazel run @rules_go//go mod tidy
	bazel mod tidy
	bazel run //:gazelle
	bazel build //proto/...
	bazel query "attr(name, '.*proto_link$$', //...)" | xargs -I {} bazel run {}

.PHONY: proto
proto:
	bazel build //proto/...
	bazel query "attr(name, '.*proto_link$$', //...)" | xargs -I {} bazel run {}

.PHONY: test
test:
	bazel test //... --test_output=all