# tools.mk: rules for installing tools used by this project.

# The underscore is to prevent the `go` command from considering any Go files that may exist in
# downloaded third-party dependencies.
tools := _tools
$(tools):
	mkdir -p '$@'

# protoc: the protobuf compiler
protoc_version := 3.19.1
protoc_dir := $(tools)/protoc-$(protoc_version)
protoc_archive := $(protoc_dir).zip
protoc := $(protoc_dir)/bin/protoc
# TODO: make this more platform-independent (`linux` is specified in the archive URL.)
$(protoc_archive): | $(tools)
	curl \
		--fail \
		--location \
		--show-error \
		--silent \
		--output '$@' \
		'https://github.com/protocolbuffers/protobuf/releases/download/v$(protoc_version)/protoc-$(protoc_version)-linux-x86_64.zip'

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

# googleapis: a repository of protobuf definitions, for example the useful `google.api` and
# `google.type` packages. It's not a tool per se, but I'm hacking it in this way for now.
#
# branch "master" as of 2021-03-28
googleapis_version := c3918feb2415878ac428d728fb473ed4187b7819
googleapis_dir := $(tools)/googleapis-$(googleapis_version)
googleapis_archive := $(googleapis_dir).zip
$(googleapis_archive): | $(tools)
	curl \
		--fail \
		--location \
		--show-error \
		--silent \
		--output '$@' \
		'https://github.com/googleapis/googleapis/archive/$(googleapis_version).zip'

$(googleapis_dir): | $(googleapis_archive)
	unzip \
		'$(googleapis_archive)' \
		-d '$(@D)'

# buildifier: a formatter and linter for BUILD.bazel files.
buildifier := $(tools)/buildifier
$(buildifier): go.mod go.sum
	go \
		build \
		-o='$@' \
		github.com/bazelbuild/buildtools/buildifier
