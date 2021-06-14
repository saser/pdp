use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;

    aoc::testfn!(
        example1,
        aoc::Input::String("9 players; last marble is worth 25 points"),
        "32",
        crate::part1
    );
    aoc::testfn!(
        example2,
        aoc::Input::String("10 players; last marble is worth 1618 points"),
        "8317",
        crate::part1
    );
    aoc::testfn!(
        example3,
        aoc::Input::String("13 players; last marble is worth 7999 points"),
        "146373",
        crate::part1
    );
    aoc::testfn!(
        example4,
        aoc::Input::String("17 players; last marble is worth 1104 points"),
        "2764",
        crate::part1
    );
    aoc::testfn!(
        example5,
        aoc::Input::String("21 players; last marble is worth 6111 points"),
        "54718",
        crate::part1
    );
    aoc::testfn!(
        example6,
        aoc::Input::String("30 players; last marble is worth 5807 points"),
        "37305",
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2018/09"),
        "436720",
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2018/09"),
        "3527845091",
        crate::part2
    );
}
