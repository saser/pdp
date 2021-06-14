use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(example1, aoc::Input::String("{}"), "1", crate::part1);
    aoc::testfn!(example2, aoc::Input::String("{{{}}}"), "6", crate::part1);
    aoc::testfn!(example3, aoc::Input::String("{{},{}}"), "5", crate::part1);
    aoc::testfn!(
        example4,
        aoc::Input::String("{{{},{},{{}}}}"),
        "16",
        crate::part1
    );
    aoc::testfn!(
        example5,
        aoc::Input::String("{<a>,<a>,<a>,<a>}"),
        "1",
        crate::part1
    );
    aoc::testfn!(
        example6,
        aoc::Input::String("{{<ab>},{<ab>},{<ab>},{<ab>}}"),
        "9",
        crate::part1
    );
    aoc::testfn!(
        example7,
        aoc::Input::String("{{<!!>},{<!!>},{<!!>},{<!!>}}"),
        "9",
        crate::part1
    );
    aoc::testfn!(
        example8,
        aoc::Input::String("{{<a!>},{<a!>},{<a!>},{<ab>}}"),
        "3",
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/09"),
        "21037",
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(example1, aoc::Input::String("<>"), "0", crate::part2);
    aoc::testfn!(
        example2,
        aoc::Input::String("<random characters>"),
        "17",
        crate::part2
    );
    aoc::testfn!(example3, aoc::Input::String("<<<<>"), "3", crate::part2);
    aoc::testfn!(example4, aoc::Input::String("<{!>}>"), "2", crate::part2);
    aoc::testfn!(example5, aoc::Input::String("<!!>"), "0", crate::part2);
    aoc::testfn!(example6, aoc::Input::String("<!!!>>"), "0", crate::part2);
    aoc::testfn!(
        example7,
        aoc::Input::String("<{o\"i!a,<{i<a>"),
        "10",
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/09"),
        "9495",
        crate::part2
    );
}
