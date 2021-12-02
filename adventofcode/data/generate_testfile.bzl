"""Build macros to generate test files based on adventofcode.Instance proto
messages."""

load("@io_bazel_rules_go//go:def.bzl", "go_test")

def go_instance_test(
        name,
        library = None,
        go_package = None,
        part1 = None,
        part2 = None,
        instances = []):
    """
    go_instance_test generates a go_test target for the given instances.

    The go_test target will be generated as a regular unit test
    target, by embedding the library.

    Args:
        name: String. Name of the generated go_test target.
        library: Label. Must be a go_library target. It will embedded
            in the go_test target using the `embed` attribute.
        go_package: String. The name of the embedded package.
        part1: String. The name of the solver function for part 1.
        part2: String. The name of the solver function for part 2.
        instances: List of labels. The instance files from which test
            cases will be created. At least one file must be given.
    """

    if library == None:
        fail("library is required")
    if go_package == None:
        fail("go_package is required")
    if len(instances) == 0:
        fail("instances is required to be non-empty")
    if part1 == None and part2 == None:
        fail("at least one of part1 and part2 is required")

    srcs = instances
    out_file = "%s.go" % name
    outs = [out_file]
    cmd = [
        "$(location //adventofcode/tools/aoctool)",
        "generate",
        "testfile",
        "--go_out='$(location %s)'" % out_file,
        "--go_package='%s'" % go_package,
    ]
    if part1 != None:
        cmd.append("--go_part1='%s'" % part1)
    if part2 != None:
        cmd.append("--go_part2='%s'" % part2)
    for instance in instances:
        cmd.append("'$(location %s)'" % instance)

    native.genrule(
        name = "%s_go" % name,
        srcs = srcs,
        outs = outs,
        cmd = " ".join(cmd),
        exec_tools = ["//adventofcode/tools/aoctool"],
    )

    go_test(
        name = name,
        srcs = [out_file],
        embed = [library],
        deps = [
            "//adventofcode/adventofcode_go_proto",
            "@org_golang_google_protobuf//proto",
        ],
    )
