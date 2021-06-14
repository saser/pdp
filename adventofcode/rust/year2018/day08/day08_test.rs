use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Input::String("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"),
        "138",
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2018/08"),
        "40908",
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Input::String("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"),
        "66",
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2018/08"),
        "25910",
        crate::part2
    );
}
