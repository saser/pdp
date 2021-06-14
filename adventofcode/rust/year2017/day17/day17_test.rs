use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(example, aoc::Input::String("3"), "638", crate::part1);
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/17"),
        "1311",
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/17"),
        "39170601",
        crate::part2
    );
}
