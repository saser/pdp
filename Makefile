include tools.mk

proto_files := $(shell git ls-files -- '*.proto')
go_module := $(shell go list -m)

.PHONY: generate
generate: \
	$(googleapis_dir) \
	$(proto_files) \
	$(protoc) \
	$(protoc-gen-go) \
	$(protoc-gen-go-grpc) \
	$(protoc_dir)
generate:
	$(protoc) \
		--proto_path='.' \
		--proto_path='$(googleapis_dir)' \
		--proto_path='$(protoc_dir)/include' \
		--plugin='$(protoc-gen-go)' \
		--go_out=. \
		--go_opt=module='$(go_module)' \
		--plugin='$(protoc-gen-go-grpc)' \
		--go-grpc_out=. \
		--go-grpc_opt=module='$(go_module)' \
		$(proto_files)

.PHONY: lint
buildifier: \
	$(buildifier)
buildifier:
	$(buildifier) \
		-lint=fix \
		-warnings=all \
		-r \
		.
