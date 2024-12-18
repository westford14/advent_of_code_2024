# Proto -> go source

This gazelle plugin and bazel rule are responsible for copying the *.go files generated
from protoc from the bazel sandbox back into the source tree.

This is mostly for local development convenience. This allows editors to pick up the go
sources, and allows users to inspect the files.

This also allows for using the go toolchain for running and debugging things directly.

`main.go` implements the [Language interface](https://pkg.go.dev/github.com/bazelbuild/bazel-gazelle/language#Language)
and allows gazelle to generate targets with the `go_proto_link` rule defined in [golink.bzl](./golink.bzl)