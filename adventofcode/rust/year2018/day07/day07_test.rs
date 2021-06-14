use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Input::File("adventofcode/rust/year2018/testdata/day07/ex"),
        "CABDFE",
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2018/07"),
        "MNQKRSFWGXPZJCOTVYEBLAHIUD",
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Input::File("adventofcode/rust/year2018/testdata/day07/ex"),
        "253",
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2018/07"),
        "948",
        crate::part2
    );
}
