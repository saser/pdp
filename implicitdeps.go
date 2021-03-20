// +build implicitdeps

// This file copies the imports from implicit dependencies required by, for example, generated
// protobuf code, or external Bazel repositories. This is so that the Go module will define the
// appropriate dependencies in `go.mod`/`go.sum`, which in turn can be picked up by Gazelles
// `update-repos` command.

package implicitdeps

import (
	// Imports required by generated protobuf code.
	_ "context"
	_ "reflect"
	_ "sync"

	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/reflect/protoreflect"
	_ "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"

	// Packages required by `rules_go`.
	_ "golang.org/x/tools/go/analysis"
)
