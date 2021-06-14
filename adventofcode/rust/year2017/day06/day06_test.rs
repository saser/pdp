use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(example, aoc::Input::String("0 2 7 0"), "5", crate::part1);
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/06"),
        "5042",
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(example, aoc::Input::String("0 2 7 0"), "4", crate::part2);
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/06"),
        "1086",
        crate::part2
    );
}
