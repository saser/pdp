use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Input::File("adventofcode/rust/year2017/testdata/day12/ex"),
        "6",
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/12"),
        "141",
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Input::File("adventofcode/rust/year2017/testdata/day12/ex"),
        "2",
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/12"),
        "171",
        crate::part2
    );
}
