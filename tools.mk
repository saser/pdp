# tools.mk: rules for installing tools used by this project.

# The underscore is to prevent the `go` command from considering any Go files that may exist in
# downloaded third-party dependencies.
tools := _tools
$(tools):
	mkdir -p '$@'

# protoc: the protobuf compiler
protoc_version := 3.15.6
protoc_dir := $(tools)/protoc-$(protoc_version)
protoc_archive := $(protoc_dir).zip
protoc := $(protoc_dir)/bin/protoc
# TODO: make this more platform-independent (`osx` is specified in the archive URL.)
$(protoc_archive): | $(tools)
	curl \
		--fail \
		--location \
		--show-error \
		--silent \
		--output '$@' \
		'https://github.com/protocolbuffers/protobuf/releases/download/v$(protoc_version)/protoc-$(protoc_version)-osx-x86_64.zip'

$(protoc_dir): $(protoc_archive)
	unzip \
		'$<' \
		-d '$@'

$(protoc): | $(protoc_dir)

# protoc-gen-go: protoc plugin to generate Go code for protobufs.
protoc-gen-go := $(tools)/protoc-gen-go
$(protoc-gen-go): go.mod go.sum
	go \
		build \
		-o='$@' \
		google.golang.org/protobuf/cmd/protoc-gen-go

# protoc-gen-go-grpc: protoc plugin to generate Go code for protobuf services.
protoc-gen-go-grpc := $(tools)/protoc-gen-go-grpc
$(protoc-gen-go-grpc): go.mod go.sum
	go \
		build \
		-o='$@' \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc
