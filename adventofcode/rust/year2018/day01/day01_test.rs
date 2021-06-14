use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;

    aoc::testfn!(
        example1,
        aoc::Input::File("adventofcode/rust/year2018/testdata/day01/p1ex1"),
        "3",
        crate::part1
    );
    aoc::testfn!(
        example2,
        aoc::Input::File("adventofcode/rust/year2018/testdata/day01/p1ex2"),
        "3",
        crate::part1
    );
    aoc::testfn!(
        example3,
        aoc::Input::File("adventofcode/rust/year2018/testdata/day01/p1ex3"),
        "0",
        crate::part1
    );
    aoc::testfn!(
        example4,
        aoc::Input::File("adventofcode/rust/year2018/testdata/day01/p1ex4"),
        "-6",
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2018/01"),
        "416",
        crate::part1
    );
}

mod part2 {
    use super::*;

    aoc::testfn!(
        example1,
        aoc::Input::File("adventofcode/rust/year2018/testdata/day01/p2ex1"),
        "2",
        crate::part2
    );
    aoc::testfn!(
        example2,
        aoc::Input::File("adventofcode/rust/year2018/testdata/day01/p2ex2"),
        "0",
        crate::part2
    );
    aoc::testfn!(
        example3,
        aoc::Input::File("adventofcode/rust/year2018/testdata/day01/p2ex3"),
        "10",
        crate::part2
    );
    aoc::testfn!(
        example4,
        aoc::Input::File("adventofcode/rust/year2018/testdata/day01/p2ex4"),
        "5",
        crate::part2
    );
    aoc::testfn!(
        example5,
        aoc::Input::File("adventofcode/rust/year2018/testdata/day01/p2ex5"),
        "14",
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2018/01"),
        "56752",
        crate::part2
    );
}
