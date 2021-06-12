load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

################################################################################
# Go
################################################################################

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "69de5c704a05ff37862f7e0f5534d4f479418afc21806c887db544a316f3cb6b",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.27.0/rules_go-v0.27.0.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.27.0/rules_go-v0.27.0.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "62ca106be173579c0a167deb23358fdfe71ffa1e4cfdddf5582af26520f1c66f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("//:go_repositories.bzl", "go_repositories")

# gazelle:repository_macro go_repositories.bzl%go_repositories
go_repositories()

go_rules_dependencies()

go_register_toolchains(version = "1.16.2")

gazelle_dependencies()

################################################################################
# C++
################################################################################

http_archive(
    name = "rules_cc",
    sha256 = "b295cad8c5899e371dde175079c0a2cdc0151f5127acc92366a8c986beb95c76",
    strip_prefix = "rules_cc-daf6ace7cfeacd6a83e9ff2ed659f416537b6c74",
    urls = ["https://github.com/bazelbuild/rules_cc/archive/daf6ace7cfeacd6a83e9ff2ed659f416537b6c74.zip"],
)

git_repository(
    name = "googletest",
    # `commit` and `shallow_since` was given by first specifying:
    #     tag = "release-1.10.0"
    # and then following the debug messages given by Bazel.
    commit = "703bd9caab50b139428cea1aaff9974ebee5742e",
    remote = "https://github.com/google/googletest",
    shallow_since = "1570114335 -0400",
)

git_repository(
    name = "googlebenchmark",
    # `commit` and `shallow_since` was given by first specifying:
    #     tag = "v1.5.0"
    # and then following the debug messages given by Bazel.
    commit = "090faecb454fbd6e6e17a75ef8146acb037118d4",
    remote = "https://github.com/google/benchmark",
    shallow_since = "1557776538 +0300",
)

git_repository(
    name = "abseil",
    # `commit` and `shallow_since` was given by first specifying:
    #     tag = "20190808"
    # and then following the debug messages given by Bazel.
    commit = "aa844899c937bde5d2b24f276b59997e5b668bde",
    remote = "https://github.com/abseil/abseil-cpp",
    shallow_since = "1565288385 -0400",
)

################################################################################
# Java
################################################################################

git_repository(
    name = "rules_jvm_external",
    # `commit` and `shallow_since` was given by first specifying:
    #     tag = "3.1"
    # and then following the debug messages given by Bazel.
    commit = "9aec21a7eff032dfbdcf728bb608fe1a02c54124",
    remote = "https://github.com/bazelbuild/rules_jvm_external",
    shallow_since = "1577467222 -0500",
)

load("@rules_jvm_external//:defs.bzl", "maven_install")

maven_install(
    artifacts = [
        "junit:junit:4.13",
        "org.openjdk.jmh:jmh-core:1.22",
        "org.openjdk.jmh:jmh-generator-annprocess:1.22",
    ],
    repositories = [
        "https://repo1.maven.org/maven2",
    ],
)
