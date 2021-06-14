use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Input::File("adventofcode/rust/year2018/testdata/day06/ex"),
        "17",
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2018/06"),
        "3687",
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2018/06"),
        "40134",
        crate::part2
    );
}
