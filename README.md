# Advent of Code 2024

## Bazel

* bazel run //bzl_tools/go_fix:go_fix
* bazel run @rules_go//go mod tidy
* bazel mod tidy
* bazel run //:gazelle
* bazel run //:buildifier.check
* bazel build //...
* bazel test //... --test_output=all