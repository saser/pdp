use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Input::File("adventofcode/rust/year2017/testdata/day13/ex"),
        "24",
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/13"),
        "2508",
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Input::File("adventofcode/rust/year2017/testdata/day13/ex"),
        "10",
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/13"),
        "3913186",
        crate::part2
    );
}
