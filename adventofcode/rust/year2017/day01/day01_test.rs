use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(example1, aoc::Input::String("1122"), "3", crate::part1);
    aoc::testfn!(example2, aoc::Input::String("1111"), "4", crate::part1);
    aoc::testfn!(example3, aoc::Input::String("1234"), "0", crate::part1);
    aoc::testfn!(example4, aoc::Input::String("91212129"), "9", crate::part1);
    aoc::testfn!(actual, aoc::Input::File("adventofcode/inputs/2017/01"), "1044", crate::part1);
}

mod part2 {
    use super::*;
    aoc::testfn!(example1, aoc::Input::String("1212"), "6", crate::part2);
    aoc::testfn!(example2, aoc::Input::String("1221"), "0", crate::part2);
    aoc::testfn!(example3, aoc::Input::String("123425"), "4", crate::part2);
    aoc::testfn!(example4, aoc::Input::String("123123"), "12", crate::part2);
    aoc::testfn!(example5, aoc::Input::String("12131415"), "4", crate::part2);
    aoc::testfn!(actual, aoc::Input::File("adventofcode/inputs/2017/01"), "1054", crate::part2);
}
