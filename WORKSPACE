load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

PROTOBUF_VERSION = "3.15.6"

http_archive(
    name = "com_google_protobuf",
    sha256 = "985bb1ca491f0815daad825ef1857b684e0844dc68123626a08351686e8d30c9",
    strip_prefix = "protobuf-%s" % PROTOBUF_VERSION,
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v%s.zip" % PROTOBUF_VERSION],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

RULES_PROTO_COMMIT = "97d8af4dc474595af3900dd85cb3a29ad28cc313"

http_archive(
    name = "rules_proto",
    sha256 = "602e7161d9195e50246177e7c55b2f39950a9cf7366f74ed5f22fd45750cd208",
    strip_prefix = "rules_proto-%s" % RULES_PROTO_COMMIT,
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/%s.tar.gz" % RULES_PROTO_COMMIT,
        "https://github.com/bazelbuild/rules_proto/archive/%s.tar.gz" % RULES_PROTO_COMMIT,
    ],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

RULES_GO_VERSION = "v0.27.0"

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "69de5c704a05ff37862f7e0f5534d4f479418afc21806c887db544a316f3cb6b",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/{v}/rules_go-{v}.tar.gz".format(v = RULES_GO_VERSION),
        "https://github.com/bazelbuild/rules_go/releases/download/{v}/rules_go-{v}.tar.gz".format(v = RULES_GO_VERSION),
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.16.2")

GAZELLE_VERSION = "v0.23.0"

http_archive(
    name = "bazel_gazelle",
    sha256 = "62ca106be173579c0a167deb23358fdfe71ffa1e4cfdddf5582af26520f1c66f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/{v}/bazel-gazelle-{v}.tar.gz".format(v = GAZELLE_VERSION),
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/{v}/bazel-gazelle-{v}.tar.gz".format(v = GAZELLE_VERSION),
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

load("//:go_repositories.bzl", "go_repositories")

# gazelle:repository_macro go_repositories.bzl%go_repositories
go_repositories()

gazelle_dependencies()
