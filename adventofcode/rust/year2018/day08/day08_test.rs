use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::String("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"),
        aoc::Data::String("138"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2018/day08/actual.in"),
        aoc::Data::String("40908"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::String("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"),
        aoc::Data::String("66"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2018/day08/actual.in"),
        aoc::Data::String("25910"),
        crate::part2
    );
}
