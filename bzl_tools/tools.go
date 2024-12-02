//go:build tools
// +build tools

package tools

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/genproto/googleapis/api"

	_ "golang.org/x/tools/cmd/goimports"
)

// This file imports packages that are used when running go generate, or used
// during the development process but not otherwise depended on by built code.
