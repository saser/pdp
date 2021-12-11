"""
@generated
cargo-raze generated Bazel file.

DO NOT EDIT! Replaced on runs of cargo-raze
"""

load("@bazel_tools//tools/build_defs/repo:git.bzl", "new_git_repository")  # buildifier: disable=load
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")  # buildifier: disable=load
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")  # buildifier: disable=load

def raze_fetch_remote_crates():
    """This function defines a collection of repos and should be called in a WORKSPACE file"""
    maybe(
        http_archive,
        name = "raze__aho_corasick__0_6_9",
        url = "https://crates.io/api/v1/crates/aho-corasick/0.6.9/download",
        type = "tar.gz",
        sha256 = "1e9a933f4e58658d7b12defcf96dc5c720f20832deebe3e0a19efd3b6aaeeb9e",
        strip_prefix = "aho-corasick-0.6.9",
        build_file = Label("//cargo/remote:BUILD.aho-corasick-0.6.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__alga__0_7_2",
        url = "https://crates.io/api/v1/crates/alga/0.7.2/download",
        type = "tar.gz",
        sha256 = "24bb00eeca59f2986c747b8c2f271d52310ce446be27428fc34705138b155778",
        strip_prefix = "alga-0.7.2",
        build_file = Label("//cargo/remote:BUILD.alga-0.7.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__approx__0_3_2",
        url = "https://crates.io/api/v1/crates/approx/0.3.2/download",
        type = "tar.gz",
        sha256 = "f0e60b75072ecd4168020818c0107f2857bb6c4e64252d8d3983f6263b40a5c3",
        strip_prefix = "approx-0.3.2",
        build_file = Label("//cargo/remote:BUILD.approx-0.3.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__autocfg__0_1_7",
        url = "https://crates.io/api/v1/crates/autocfg/0.1.7/download",
        type = "tar.gz",
        sha256 = "1d49d90015b3c36167a20fe2810c5cd875ad504b39cff3d4eae7977e6b7c1cb2",
        strip_prefix = "autocfg-0.1.7",
        build_file = Label("//cargo/remote:BUILD.autocfg-0.1.7.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__bitflags__1_0_1",
        url = "https://crates.io/api/v1/crates/bitflags/1.0.1/download",
        type = "tar.gz",
        sha256 = "b3c30d3802dfb7281680d6285f2ccdaa8c2d8fee41f93805dba5c4cf50dc23cf",
        strip_prefix = "bitflags-1.0.1",
        build_file = Label("//cargo/remote:BUILD.bitflags-1.0.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__cfg_if__0_1_6",
        url = "https://crates.io/api/v1/crates/cfg-if/0.1.6/download",
        type = "tar.gz",
        sha256 = "082bb9b28e00d3c9d39cc03e64ce4cea0f1bb9b3fde493f0cbc008472d22bdf4",
        strip_prefix = "cfg-if-0.1.6",
        build_file = Label("//cargo/remote:BUILD.cfg-if-0.1.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__chrono__0_4_9",
        url = "https://crates.io/api/v1/crates/chrono/0.4.9/download",
        type = "tar.gz",
        sha256 = "e8493056968583b0193c1bb04d6f7684586f3726992d6c573261941a895dbd68",
        strip_prefix = "chrono-0.4.9",
        build_file = Label("//cargo/remote:BUILD.chrono-0.4.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__cloudabi__0_0_3",
        url = "https://crates.io/api/v1/crates/cloudabi/0.0.3/download",
        type = "tar.gz",
        sha256 = "ddfc5b9aa5d4507acaf872de71051dfd0e309860e88966e1051e462a077aac4f",
        strip_prefix = "cloudabi-0.0.3",
        build_file = Label("//cargo/remote:BUILD.cloudabi-0.0.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__crossbeam_deque__0_7_2",
        url = "https://crates.io/api/v1/crates/crossbeam-deque/0.7.2/download",
        type = "tar.gz",
        sha256 = "c3aa945d63861bfe624b55d153a39684da1e8c0bc8fba932f7ee3a3c16cea3ca",
        strip_prefix = "crossbeam-deque-0.7.2",
        build_file = Label("//cargo/remote:BUILD.crossbeam-deque-0.7.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__crossbeam_epoch__0_8_0",
        url = "https://crates.io/api/v1/crates/crossbeam-epoch/0.8.0/download",
        type = "tar.gz",
        sha256 = "5064ebdbf05ce3cb95e45c8b086f72263f4166b29b97f6baff7ef7fe047b55ac",
        strip_prefix = "crossbeam-epoch-0.8.0",
        build_file = Label("//cargo/remote:BUILD.crossbeam-epoch-0.8.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__crossbeam_queue__0_1_2",
        url = "https://crates.io/api/v1/crates/crossbeam-queue/0.1.2/download",
        type = "tar.gz",
        sha256 = "7c979cd6cfe72335896575c6b5688da489e420d36a27a0b9eb0c73db574b4a4b",
        strip_prefix = "crossbeam-queue-0.1.2",
        build_file = Label("//cargo/remote:BUILD.crossbeam-queue-0.1.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__crossbeam_utils__0_6_6",
        url = "https://crates.io/api/v1/crates/crossbeam-utils/0.6.6/download",
        type = "tar.gz",
        sha256 = "04973fa96e96579258a5091af6003abde64af786b860f18622b82e026cca60e6",
        strip_prefix = "crossbeam-utils-0.6.6",
        build_file = Label("//cargo/remote:BUILD.crossbeam-utils-0.6.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__crossbeam_utils__0_7_0",
        url = "https://crates.io/api/v1/crates/crossbeam-utils/0.7.0/download",
        type = "tar.gz",
        sha256 = "ce446db02cdc3165b94ae73111e570793400d0794e46125cc4056c81cbb039f4",
        strip_prefix = "crossbeam-utils-0.7.0",
        build_file = Label("//cargo/remote:BUILD.crossbeam-utils-0.7.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__either__1_5_3",
        url = "https://crates.io/api/v1/crates/either/1.5.3/download",
        type = "tar.gz",
        sha256 = "bb1f6b1ce1c140482ea30ddd3335fc0024ac7ee112895426e0a629a6c20adfe3",
        strip_prefix = "either-1.5.3",
        build_file = Label("//cargo/remote:BUILD.either-1.5.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__fuchsia_cprng__0_1_1",
        url = "https://crates.io/api/v1/crates/fuchsia-cprng/0.1.1/download",
        type = "tar.gz",
        sha256 = "a06f77d526c1a601b7c4cdd98f54b5eaabffc14d5f2f0296febdc7f357c6d3ba",
        strip_prefix = "fuchsia-cprng-0.1.1",
        build_file = Label("//cargo/remote:BUILD.fuchsia-cprng-0.1.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__generic_array__0_11_1",
        url = "https://crates.io/api/v1/crates/generic-array/0.11.1/download",
        type = "tar.gz",
        sha256 = "8107dafa78c80c848b71b60133954b4a58609a3a1a5f9af037ecc7f67280f369",
        strip_prefix = "generic-array-0.11.1",
        build_file = Label("//cargo/remote:BUILD.generic-array-0.11.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__hermit_abi__0_1_3",
        url = "https://crates.io/api/v1/crates/hermit-abi/0.1.3/download",
        type = "tar.gz",
        sha256 = "307c3c9f937f38e3534b1d6447ecf090cafcc9744e4a6360e8b037b2cf5af120",
        strip_prefix = "hermit-abi-0.1.3",
        build_file = Label("//cargo/remote:BUILD.hermit-abi-0.1.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__lazy_static__1_2_0",
        url = "https://crates.io/api/v1/crates/lazy_static/1.2.0/download",
        type = "tar.gz",
        sha256 = "a374c89b9db55895453a74c1e38861d9deec0b01b405a82516e9d5de4820dea1",
        strip_prefix = "lazy_static-1.2.0",
        build_file = Label("//cargo/remote:BUILD.lazy_static-1.2.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__libc__0_2_65",
        url = "https://crates.io/api/v1/crates/libc/0.2.65/download",
        type = "tar.gz",
        sha256 = "1a31a0627fdf1f6a39ec0dd577e101440b7db22672c0901fe00a9a6fbb5c24e8",
        strip_prefix = "libc-0.2.65",
        build_file = Label("//cargo/remote:BUILD.libc-0.2.65.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__libm__0_1_4",
        url = "https://crates.io/api/v1/crates/libm/0.1.4/download",
        type = "tar.gz",
        sha256 = "7fc7aa29613bd6a620df431842069224d8bc9011086b1db4c0e0cd47fa03ec9a",
        strip_prefix = "libm-0.1.4",
        build_file = Label("//cargo/remote:BUILD.libm-0.1.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__matrixmultiply__0_1_15",
        url = "https://crates.io/api/v1/crates/matrixmultiply/0.1.15/download",
        type = "tar.gz",
        sha256 = "dcad67dcec2d58ff56f6292582377e6921afdf3bfbd533e26fb8900ae575e002",
        strip_prefix = "matrixmultiply-0.1.15",
        build_file = Label("//cargo/remote:BUILD.matrixmultiply-0.1.15.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__memchr__2_1_1",
        url = "https://crates.io/api/v1/crates/memchr/2.1.1/download",
        type = "tar.gz",
        sha256 = "0a3eb002f0535929f1199681417029ebea04aadc0c7a4224b46be99c7f5d6a16",
        strip_prefix = "memchr-2.1.1",
        build_file = Label("//cargo/remote:BUILD.memchr-2.1.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__memoffset__0_5_3",
        url = "https://crates.io/api/v1/crates/memoffset/0.5.3/download",
        type = "tar.gz",
        sha256 = "75189eb85871ea5c2e2c15abbdd541185f63b408415e5051f5cac122d8c774b9",
        strip_prefix = "memoffset-0.5.3",
        build_file = Label("//cargo/remote:BUILD.memoffset-0.5.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__nalgebra__0_16_14",
        url = "https://crates.io/api/v1/crates/nalgebra/0.16.14/download",
        type = "tar.gz",
        sha256 = "ccb86df349ecf5f549f6e12f6de4972cdd912d0bc290c1ca4d34d4b4b21a6f98",
        strip_prefix = "nalgebra-0.16.14",
        build_file = Label("//cargo/remote:BUILD.nalgebra-0.16.14.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__num_complex__0_2_3",
        url = "https://crates.io/api/v1/crates/num-complex/0.2.3/download",
        type = "tar.gz",
        sha256 = "fcb0cf31fb3ff77e6d2a6ebd6800df7fdcd106f2ad89113c9130bcd07f93dffc",
        strip_prefix = "num-complex-0.2.3",
        build_file = Label("//cargo/remote:BUILD.num-complex-0.2.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__num_integer__0_1_41",
        url = "https://crates.io/api/v1/crates/num-integer/0.1.41/download",
        type = "tar.gz",
        sha256 = "b85e541ef8255f6cf42bbfe4ef361305c6c135d10919ecc26126c4e5ae94bc09",
        strip_prefix = "num-integer-0.1.41",
        build_file = Label("//cargo/remote:BUILD.num-integer-0.1.41.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__num_traits__0_2_9",
        url = "https://crates.io/api/v1/crates/num-traits/0.2.9/download",
        type = "tar.gz",
        sha256 = "443c53b3c3531dfcbfa499d8893944db78474ad7a1d87fa2d94d1a2231693ac6",
        strip_prefix = "num-traits-0.2.9",
        build_file = Label("//cargo/remote:BUILD.num-traits-0.2.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__num_cpus__1_11_1",
        url = "https://crates.io/api/v1/crates/num_cpus/1.11.1/download",
        type = "tar.gz",
        sha256 = "76dac5ed2a876980778b8b85f75a71b6cbf0db0b1232ee12f826bccb00d09d72",
        strip_prefix = "num_cpus-1.11.1",
        build_file = Label("//cargo/remote:BUILD.num_cpus-1.11.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__rand__0_5_6",
        url = "https://crates.io/api/v1/crates/rand/0.5.6/download",
        type = "tar.gz",
        sha256 = "c618c47cd3ebd209790115ab837de41425723956ad3ce2e6a7f09890947cacb9",
        strip_prefix = "rand-0.5.6",
        build_file = Label("//cargo/remote:BUILD.rand-0.5.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__rand_core__0_3_1",
        url = "https://crates.io/api/v1/crates/rand_core/0.3.1/download",
        type = "tar.gz",
        sha256 = "7a6fdeb83b075e8266dcc8762c22776f6877a63111121f5f8c7411e5be7eed4b",
        strip_prefix = "rand_core-0.3.1",
        build_file = Label("//cargo/remote:BUILD.rand_core-0.3.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__rand_core__0_4_2",
        url = "https://crates.io/api/v1/crates/rand_core/0.4.2/download",
        type = "tar.gz",
        sha256 = "9c33a3c44ca05fa6f1807d8e6743f3824e8509beca625669633be0acbdf509dc",
        strip_prefix = "rand_core-0.4.2",
        build_file = Label("//cargo/remote:BUILD.rand_core-0.4.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__rawpointer__0_1_0",
        url = "https://crates.io/api/v1/crates/rawpointer/0.1.0/download",
        type = "tar.gz",
        sha256 = "ebac11a9d2e11f2af219b8b8d833b76b1ea0e054aa0e8d8e9e4cbde353bdf019",
        strip_prefix = "rawpointer-0.1.0",
        build_file = Label("//cargo/remote:BUILD.rawpointer-0.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__rayon__1_2_0",
        url = "https://crates.io/api/v1/crates/rayon/1.2.0/download",
        type = "tar.gz",
        sha256 = "83a27732a533a1be0a0035a111fe76db89ad312f6f0347004c220c57f209a123",
        strip_prefix = "rayon-1.2.0",
        build_file = Label("//cargo/remote:BUILD.rayon-1.2.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__rayon_core__1_6_0",
        url = "https://crates.io/api/v1/crates/rayon-core/1.6.0/download",
        type = "tar.gz",
        sha256 = "98dcf634205083b17d0861252431eb2acbfb698ab7478a2d20de07954f47ec7b",
        strip_prefix = "rayon-core-1.6.0",
        build_file = Label("//cargo/remote:BUILD.rayon-core-1.6.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__redox_syscall__0_1_32",
        url = "https://crates.io/api/v1/crates/redox_syscall/0.1.32/download",
        type = "tar.gz",
        sha256 = "ab105df655884ede59d45b7070c8a65002d921461ee813a024558ca16030eea0",
        strip_prefix = "redox_syscall-0.1.32",
        build_file = Label("//cargo/remote:BUILD.redox_syscall-0.1.32.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__regex__1_0_5",
        url = "https://crates.io/api/v1/crates/regex/1.0.5/download",
        type = "tar.gz",
        sha256 = "2069749032ea3ec200ca51e4a31df41759190a88edca0d2d86ee8bedf7073341",
        strip_prefix = "regex-1.0.5",
        build_file = Label("//cargo/remote:BUILD.regex-1.0.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__regex_syntax__0_6_2",
        url = "https://crates.io/api/v1/crates/regex-syntax/0.6.2/download",
        type = "tar.gz",
        sha256 = "747ba3b235651f6e2f67dfa8bcdcd073ddb7c243cb21c442fc12395dfcac212d",
        strip_prefix = "regex-syntax-0.6.2",
        build_file = Label("//cargo/remote:BUILD.regex-syntax-0.6.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__rustc_version__0_2_3",
        url = "https://crates.io/api/v1/crates/rustc_version/0.2.3/download",
        type = "tar.gz",
        sha256 = "138e3e0acb6c9fb258b19b67cb8abd63c00679d2851805ea151465464fe9030a",
        strip_prefix = "rustc_version-0.2.3",
        build_file = Label("//cargo/remote:BUILD.rustc_version-0.2.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__scopeguard__1_0_0",
        url = "https://crates.io/api/v1/crates/scopeguard/1.0.0/download",
        type = "tar.gz",
        sha256 = "b42e15e59b18a828bbf5c58ea01debb36b9b096346de35d941dcb89009f24a0d",
        strip_prefix = "scopeguard-1.0.0",
        build_file = Label("//cargo/remote:BUILD.scopeguard-1.0.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__semver__0_9_0",
        url = "https://crates.io/api/v1/crates/semver/0.9.0/download",
        type = "tar.gz",
        sha256 = "1d7eb9ef2c18661902cc47e535f9bc51b78acd254da71d375c2f6720d9a40403",
        strip_prefix = "semver-0.9.0",
        build_file = Label("//cargo/remote:BUILD.semver-0.9.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__semver_parser__0_7_0",
        url = "https://crates.io/api/v1/crates/semver-parser/0.7.0/download",
        type = "tar.gz",
        sha256 = "388a1df253eca08550bef6c72392cfe7c30914bf41df5269b68cbd6ff8f570a3",
        strip_prefix = "semver-parser-0.7.0",
        build_file = Label("//cargo/remote:BUILD.semver-parser-0.7.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__thread_local__0_3_6",
        url = "https://crates.io/api/v1/crates/thread_local/0.3.6/download",
        type = "tar.gz",
        sha256 = "c6b53e329000edc2b34dbe8545fd20e55a333362d0a321909685a19bd28c3f1b",
        strip_prefix = "thread_local-0.3.6",
        build_file = Label("//cargo/remote:BUILD.thread_local-0.3.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__time__0_1_42",
        url = "https://crates.io/api/v1/crates/time/0.1.42/download",
        type = "tar.gz",
        sha256 = "db8dcfca086c1143c9270ac42a2bbd8a7ee477b78ac8e45b19abfb0cbede4b6f",
        strip_prefix = "time-0.1.42",
        build_file = Label("//cargo/remote:BUILD.time-0.1.42.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__typenum__1_11_2",
        url = "https://crates.io/api/v1/crates/typenum/1.11.2/download",
        type = "tar.gz",
        sha256 = "6d2783fe2d6b8c1101136184eb41be8b1ad379e4657050b8aaff0c79ee7575f9",
        strip_prefix = "typenum-1.11.2",
        build_file = Label("//cargo/remote:BUILD.typenum-1.11.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__ucd_util__0_1_2",
        url = "https://crates.io/api/v1/crates/ucd-util/0.1.2/download",
        type = "tar.gz",
        sha256 = "d0f8bfa9ff0cadcd210129ad9d2c5f145c13e9ced3d3e5d948a6213487d52444",
        strip_prefix = "ucd-util-0.1.2",
        build_file = Label("//cargo/remote:BUILD.ucd-util-0.1.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__utf8_ranges__1_0_2",
        url = "https://crates.io/api/v1/crates/utf8-ranges/1.0.2/download",
        type = "tar.gz",
        sha256 = "796f7e48bef87609f7ade7e06495a87d5cd06c7866e6a5cbfceffc558a243737",
        strip_prefix = "utf8-ranges-1.0.2",
        build_file = Label("//cargo/remote:BUILD.utf8-ranges-1.0.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__version_check__0_1_5",
        url = "https://crates.io/api/v1/crates/version_check/0.1.5/download",
        type = "tar.gz",
        sha256 = "914b1a6776c4c929a602fafd8bc742e06365d4bcbe48c30f9cca5824f70dc9dd",
        strip_prefix = "version_check-0.1.5",
        build_file = Label("//cargo/remote:BUILD.version_check-0.1.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__winapi__0_3_8",
        url = "https://crates.io/api/v1/crates/winapi/0.3.8/download",
        type = "tar.gz",
        sha256 = "8093091eeb260906a183e6ae1abdba2ef5ef2257a21801128899c3fc699229c6",
        strip_prefix = "winapi-0.3.8",
        build_file = Label("//cargo/remote:BUILD.winapi-0.3.8.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__winapi_i686_pc_windows_gnu__0_4_0",
        url = "https://crates.io/api/v1/crates/winapi-i686-pc-windows-gnu/0.4.0/download",
        type = "tar.gz",
        sha256 = "ac3b87c63620426dd9b991e5ce0329eff545bccbbb34f3be09ff6fb6ab51b7b6",
        strip_prefix = "winapi-i686-pc-windows-gnu-0.4.0",
        build_file = Label("//cargo/remote:BUILD.winapi-i686-pc-windows-gnu-0.4.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__winapi_x86_64_pc_windows_gnu__0_4_0",
        url = "https://crates.io/api/v1/crates/winapi-x86_64-pc-windows-gnu/0.4.0/download",
        type = "tar.gz",
        sha256 = "712e227841d057c1ee1cd2fb22fa7e5a5461ae8e48fa2ca79ec42cfc1931183f",
        strip_prefix = "winapi-x86_64-pc-windows-gnu-0.4.0",
        build_file = Label("//cargo/remote:BUILD.winapi-x86_64-pc-windows-gnu-0.4.0.bazel"),
    )
