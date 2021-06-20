use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;

    aoc::testfn!(
        example1,
        aoc::Data::String("9 players; last marble is worth 25 points"),
        aoc::Data::String("32"),
        crate::part1
    );
    aoc::testfn!(
        example2,
        aoc::Data::String("10 players; last marble is worth 1618 points"),
        aoc::Data::String("8317"),
        crate::part1
    );
    aoc::testfn!(
        example3,
        aoc::Data::String("13 players; last marble is worth 7999 points"),
        aoc::Data::String("146373"),
        crate::part1
    );
    aoc::testfn!(
        example4,
        aoc::Data::String("17 players; last marble is worth 1104 points"),
        aoc::Data::String("2764"),
        crate::part1
    );
    aoc::testfn!(
        example5,
        aoc::Data::String("21 players; last marble is worth 6111 points"),
        aoc::Data::String("54718"),
        crate::part1
    );
    aoc::testfn!(
        example6,
        aoc::Data::String("30 players; last marble is worth 5807 points"),
        aoc::Data::String("37305"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2018/day09/actual.in"),
        aoc::Data::String("436720"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2018/day09/actual.in"),
        aoc::Data::String("3527845091"),
        crate::part2
    );
}
