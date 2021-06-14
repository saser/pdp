use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(example1, aoc::Input::String("5 1 9 5"), "8", crate::part1);
    aoc::testfn!(example2, aoc::Input::String("7 5 3"), "4", crate::part1);
    aoc::testfn!(example3, aoc::Input::String("2 4 6 8"), "6", crate::part1);
    aoc::testfn!(
        example_all,
        aoc::Input::File("adventofcode/rust/year2017/testdata/day02/p1ex"),
        "18",
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/02"),
        "36766",
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(example1, aoc::Input::String("5 9 2 8"), "4", crate::part2);
    aoc::testfn!(example2, aoc::Input::String("9 4 7 3"), "3", crate::part2);
    aoc::testfn!(example3, aoc::Input::String("3 8 6 5"), "2", crate::part2);
    aoc::testfn!(
        example_all,
        aoc::Input::File("adventofcode/rust/year2017/testdata/day02/p2ex"),
        "9",
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/02"),
        "261",
        crate::part2
    );
}
