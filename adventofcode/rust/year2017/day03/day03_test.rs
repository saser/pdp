use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::String("1"),
        aoc::Data::String("0"),
        crate::part1
    );
    aoc::testfn!(
        example2,
        aoc::Data::String("12"),
        aoc::Data::String("3"),
        crate::part1
    );
    aoc::testfn!(
        example3,
        aoc::Data::String("23"),
        aoc::Data::String("2"),
        crate::part1
    );
    aoc::testfn!(
        example4,
        aoc::Data::String("1024"),
        aoc::Data::String("31"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2017/day03/actual.in"),
        aoc::Data::String("371"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2017/day03/actual.in"),
        aoc::Data::String("369601"),
        crate::part2
    );
}
