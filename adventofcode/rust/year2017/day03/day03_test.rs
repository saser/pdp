use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(example1, aoc::Input::String("1"), "0", crate::part1);
    aoc::testfn!(example2, aoc::Input::String("12"), "3", crate::part1);
    aoc::testfn!(example3, aoc::Input::String("23"), "2", crate::part1);
    aoc::testfn!(example4, aoc::Input::String("1024"), "31", crate::part1);
    aoc::testfn!(actual, aoc::Input::File("adventofcode/inputs/2017/03"), "371", crate::part1);
}

mod part2 {
    use super::*;
    aoc::testfn!(actual, aoc::Input::File("adventofcode/inputs/2017/03"), "369601", crate::part2);
}
