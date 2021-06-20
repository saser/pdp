use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::String("3"),
        aoc::Data::String("638"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2017/day17/actual.in"),
        aoc::Data::String("1311"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2017/day17/actual.in"),
        aoc::Data::String("39170601"),
        crate::part2
    );
}
