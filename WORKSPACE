load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

################################################################################
# bazel-skylib
################################################################################

http_archive(
    name = "bazel_skylib",
    sha256 = "1c531376ac7e5a180e0237938a2536de0c54d93f5c278634818e0efc952dd56c",
    urls = [
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.0.3/bazel-skylib-1.0.3.tar.gz",
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.0.3/bazel-skylib-1.0.3.tar.gz",
    ],
)

load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

bazel_skylib_workspace()

################################################################################
# Go
################################################################################

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8e968b5fcea1d2d64071872b12737bbb5514524ee5f0a4f54f5920266c261acb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
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

go_register_toolchains(version = "1.16.5")

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

http_archive(
    name = "rules_java",
    sha256 = "34b41ec683e67253043ab1a3d1e8b7c61e4e8edefbcad485381328c934d072fe",
    url = "https://github.com/bazelbuild/rules_java/releases/download/4.0.0/rules_java-4.0.0.tar.gz",
)

load("@rules_java//java:repositories.bzl", "rules_java_dependencies", "rules_java_toolchains")

rules_java_dependencies()

rules_java_toolchains()

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

################################################################################
# Rust
################################################################################

http_archive(
    name = "rules_rust",
    sha256 = "b47bb71d60ed92ea8c07b9c841291af38e0f265b7f1b37912c90cce0428c2ce7",
    strip_prefix = "rules_rust-087bcab8154f5c0d79980ad32cb6ffb8158de649",
    urls = [
        # main branch as of 2021-06-12
        "https://github.com/bazelbuild/rules_rust/archive/087bcab8154f5c0d79980ad32cb6ffb8158de649.tar.gz",
    ],
)

load("@rules_rust//rust:repositories.bzl", "rust_repositories")

rust_repositories(
    edition = "2018",
    version = "1.52.1",
)

http_archive(
    name = "cargo_raze",
    sha256 = "0a7986b1a8ec965ee7aa317ac61e82ea08568cfdf36b7ccc4dd3d1aff3b36e0b",
    strip_prefix = "cargo-raze-0.12.0",
    url = "https://github.com/google/cargo-raze/archive/v0.12.0.tar.gz",
)

load("@cargo_raze//:repositories.bzl", "cargo_raze_repositories")

cargo_raze_repositories()

load("@cargo_raze//:transitive_deps.bzl", "cargo_raze_transitive_deps")

cargo_raze_transitive_deps()

load("//cargo:crates.bzl", "raze_fetch_remote_crates")

raze_fetch_remote_crates()
