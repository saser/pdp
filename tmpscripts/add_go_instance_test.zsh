#!/usr/bin/env zsh

# This script is for the //adventofcode/go packages. It adds an
# invocation of `go_instance_test` to each solution package (e.g.,
# //adventofcode/go/year2015/day01) called "instance_test", sets
# attributes based on some assumptions, and adds all relevant
# textprotos as instances.

set -euo pipefail

for day in {1..25}; do
    padded_day="$(printf %02d ${day})"

    # Create the new `instance_test` target.
    buildozer \
	'fix movePackageToTop' \
	'new_load //adventofcode/data:generate_testfile.bzl go_instance_test' \
	'new go_instance_test instance_test' \
	'fix unusedLoads' \
	"//adventofcode/go/year2015/day${padded_day}:__pkg__"

    # Set single-valued attributes on the `instance_test` target.
    buildozer \
	"set go_package \"day${padded_day}\"" \
	"set library \":day${padded_day}\"" \
	'set part1 "Part1"' \
	"//adventofcode/go/year2015/day${padded_day}:instance_test"
    if [[ "${day}" != "25" ]]; then
	buildozer \
	    'set part2 "Part2"' \
	    "//adventofcode/go/year2015/day${padded_day}:instance_test"
    fi

    # Find all textproto targets for the given day and add them to the
    # `instances` attribute of the `instance_test` target.
    for instance in $(bazel query "filter(textproto, //adventofcode/data/year2015/day${padded_day}:*)"); do
	buildozer \
	    "add instances \"${instance}\"" \
	    "//adventofcode/go/year2015/day${padded_day}:instance_test"
    done
done

make buildifier
