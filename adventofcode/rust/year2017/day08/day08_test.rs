use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Input::File("adventofcode/rust/year2017/testdata/day08/ex"),
        "1",
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/08"),
        "6012",
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Input::File("adventofcode/rust/year2017/testdata/day08/ex"),
        "10",
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/08"),
        "6369",
        crate::part2
    );
}
