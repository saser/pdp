def _aoctool_generate_instance(
        name,
        year = None,
        day = None,
        part = None,
        input = None,
        answer = None,
        out_file = None):
    if year == None:
        fail("year is required")
    if day == None:
        fail("day is required")
    if part == None:
        fail("part is required")
    if input == None:
        fail("input is required")
    if answer == None:
        fail("answer is required")
    if out_file == None:
        fail("out_file is required")

    native.genrule(
        name = name,
        srcs = [],
        outs = [out_file],
        cmd = " ".join([
            "$(location //adventofcode/tools/aoctool)",
            "generate",
            "instance",
            "--year=%d" % year,
            "--day=%d" % day,
            "--part=%d" % part,
            "--input=\"%s\"" % input,
            "--answer=\"%s\"" % answer,
            "--format=prototext",
            "--out_file=$(location %s)" % out_file,
        ]),
        exec_tools = ["//adventofcode/tools/aoctool"],
    )

def generate_instances(
        year = None,
        day = None,
        part_one_instances = {},
        part_two_instances = {}):
    """
    generate_instances creates textproto files of adventofcode.Instance messages.

    Args:
        year [int]: which year the instance is for. Required.
        day [int]: which day the instance is for. Required.
        part_one_instances [dict[string, List[string]]]: a mapping from name (the key) to a 2-element list representing, in order, the input and the answer, for the part 1 version of the puzzle.
        part_two_instances [dict[string, List[string]]]: see part_one_instances.
    """
    for name, data in part_one_instances.items():
        _aoctool_generate_instance(
            name = name,
            year = year,
            day = day,
            part = 1,
            input = data[0],
            answer = data[1],
            out_file = "%s.textproto" % name,
        )
    for name, data in part_two_instances.items():
        _aoctool_generate_instance(
            name = name,
            year = year,
            day = day,
            part = 2,
            input = data[0],
            answer = data[1],
            out_file = "%s.textproto" % name,
        )
