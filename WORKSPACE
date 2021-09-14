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

COM_GOOGLE_GOOGLETEST_COMMIT = "e2239ee6043f73722e7aa812a459f54a28552929"  # master as of 2021-09-14

http_archive(
    name = "com_google_googletest",
    sha256 = "8daa1a71395892f7c1ec5f7cb5b099a02e606be720d62f1a6a98f8f8898ec826",
    strip_prefix = "googletest-%s" % COM_GOOGLE_GOOGLETEST_COMMIT,
    urls = ["https://github.com/google/googletest/archive/%s.zip" % COM_GOOGLE_GOOGLETEST_COMMIT],
)

COM_GITHUB_GOOGLE_BENCHMARK_COMMIT = "713b9177183375c8b1b25595e33daf2a1625df5b"  # master as of 2021-09-14

http_archive(
    name = "com_github_google_benchmark",
    sha256 = "5cffc1043e6ba948e7fd16b0f97472ab17af07615726aa0cbf5361e4dca27597",
    strip_prefix = "benchmark-%s" % COM_GITHUB_GOOGLE_BENCHMARK_COMMIT,
    urls = ["https://github.com/google/benchmark/archive/%s.zip" % COM_GITHUB_GOOGLE_BENCHMARK_COMMIT],
)

COM_GOOGLE_ABSL_COMMIT = "b2dc72c17ac663885b62334d334da9f8970543b5"  # master as of 2021-09-14

http_archive(
    name = "com_google_absl",
    sha256 = "5247e92a222cf39ce1dfa12bf5aad452d27bab3f051f4e81f5d78d1d7ede9306",
    strip_prefix = "abseil-cpp-%s" % COM_GOOGLE_ABSL_COMMIT,
    urls = ["https://github.com/abseil/abseil-cpp/archive/%s.zip" % COM_GOOGLE_ABSL_COMMIT],
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

RULES_JVM_EXTERNAL_COMMIT = "786947c47918c44b5d740df500bb3090607df04d"  # master as of 2021-09-14

http_archive(
    name = "rules_jvm_external",
    sha256 = "3d00a53394e0e856f6a97fab75855a3fd6552190ca004f79bfac0cbfd3c1e5d1",
    strip_prefix = "rules_jvm_external-%s" % RULES_JVM_EXTERNAL_COMMIT,
    urls = ["https://github.com/bazelbuild/rules_jvm_external/archive/%s.zip" % RULES_JVM_EXTERNAL_COMMIT],
)

load("@rules_jvm_external//:defs.bzl", "maven_install")

maven_install(
    artifacts = [
        "junit:junit:4.13",
        "org.openjdk.jmh:jmh-core:1.22",
        "org.openjdk.jmh:jmh-generator-annprocess:1.22",
    ],
    maven_install_json = "//:maven_install.json",
    repositories = [
        "https://repo1.maven.org/maven2",
    ],
)

load("@maven//:defs.bzl", "pinned_maven_install")

pinned_maven_install()

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
