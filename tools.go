// +build tools

package tools

import (
	_ "github.com/bazelbuild/buildtools/buildifier"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
