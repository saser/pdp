use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(example1, aoc::Input::String("aA"), "0", crate::part1);
    aoc::testfn!(example2, aoc::Input::String("abBA"), "0", crate::part1);
    aoc::testfn!(example3, aoc::Input::String("abAB"), "4", crate::part1);
    aoc::testfn!(example4, aoc::Input::String("aabAAB"), "6", crate::part1);
    aoc::testfn!(
        example5,
        aoc::Input::String("dabAcCaCBAcCcaDA"),
        "10",
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2018/05"),
        "9686",
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Input::String("dabAcCaCBAcCcaDA"),
        "4",
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2018/05"),
        "5524",
        crate::part2
    );
}
