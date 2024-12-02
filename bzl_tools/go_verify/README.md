# Go format test generator

`main.go` implements the [Language interface](https://pkg.go.dev/github.com/bazelbuild/bazel-gazelle/language#Language)
and tells gazelle how to generate test targets for running go fmt and goimports.

This basically allows us to validate the formatting of files though the `bazel test` command.

The generated target is a `sh_test` rule that runs the `gofmt.sh` script in this directory.

See [rules_shell](https://github.com/bazelbuild/rules_shell) for more info on `sh_test`.