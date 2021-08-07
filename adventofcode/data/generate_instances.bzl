load("@bazel_skylib//lib:shell.bzl", "shell")

def instance_data(
        input = None,
        input_file = None,
        answer = None,
        answer_file = None):
    """
    instance_data generates a struct suitable for generate_instances.

    The arguments correspond to those in generate_instance.
    """
    if (input == None and input_file == None) or (input != None and input_file != None):
        fail("exactly one of input and input_file is required; got input = \"%s\", input_file = \"%s\"" % (input, input_file))
    if (answer == None and answer_file == None) or (answer != None and answer_file != None):
        fail("exactly one of answer and answer_file is required; got answer = %s, answer_file = %s" % (answer, answer_file))

    return struct(
        input = input,
        input_file = input_file,
        answer = answer,
        answer_file = answer_file,
    )

def generate_instance(
        name,
        year = None,
        day = None,
        part = None,
        input = None,
        input_file = None,
        answer = None,
        answer_file = None):
    if year == None:
        fail("year is required")
    if day == None:
        fail("day is required")
    if part == None:
        fail("part is required")
    """
    generate_instance creates a textproto file containing an adventofcode.Instance message.

    Args:
        year: Int. Required.
        day: Int. Required.
        part: Int. Required.
        input: String. The input to the problem. Exactly one of input and input_file is required.
        input_file: Label. A file containing the input to the problem. Exactly one of input and input_file is required.
        answer: String. The answer to the problem. Exactly one of answer and answer_file is required.
        answer_file: Label. A file containing the answer to the problem. Exactly one of answer and answer_file is required.
    """

    if (input == None and input_file == None) or (input != None and input_file != None):
        fail("exactly one of input and input_file must be specified; got input = \"%s\" and input_file = \"%s\"" % (input, input_file))
    if (answer == None and answer_file == None) or (answer != None and answer_file != None):
        fail("exactly one of answer and answer_file must be specified; got answer = \"%s\" and answer_file = \"%s\"" % (answer, answer_file))

    srcs = []
    out_file = "%s.textproto" % name
    outs = [out_file]
    cmd = [
        "$(location //adventofcode/tools/aoctool)",
        "generate",
        "instance",
        "--year=%d" % year,
        "--day=%d" % day,
        "--part=%d" % part,
        "--format=prototext",
        "--out_file=\"$(location %s)\"" % out_file,
    ]
    if input != None:
        cmd.append("--input=%s" % shell.quote(input))
    if input_file != None:
        cmd.append("--input_file=\"$(location %s)\"" % input_file)
        srcs.append(input_file)
    if answer != None:
        cmd.append("--answer=%s" % shell.quote(answer))
    if answer_file != None:
        cmd.append("--answer_file=\"$(location %s)\"" % answer_file)
        srcs.append(answer_file)

    native.genrule(
        name = name,
        srcs = srcs,
        outs = outs,
        cmd = " ".join(cmd),
        exec_tools = ["//adventofcode/tools/aoctool"],
    )

def generate_instances(
        year = None,
        day = None,
        part1_instances = {},
        part2_instances = {}):
    """generate_instances is a convenience wrapper around generate_instance.

    Args:
        year: Int. Required.
        day: Int. Required.
        part1_instances: dict[string, struct]. A dict mapping from name to a struct. The struct should be created using instance_data. For each item in the dict, a target called part1_{name} will be generated.
        part2_instances: dict[string, struct]. See part1_instances.
    """
    for part, instances in {
        1: part1_instances,
        2: part2_instances,
    }.items():
        for name, data in instances.items():
            generate_instance(
                name = "part%d_%s" % (part, name),
                year = year,
                day = day,
                part = part,
                input = data.input,
                input_file = data.input_file,
                answer = data.answer,
                answer_file = data.answer_file,
            )
