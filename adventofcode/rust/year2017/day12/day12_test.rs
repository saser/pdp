use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2017/testdata/day12/ex"),
        aoc::Data::String("6"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2017/day12/actual.in"),
        aoc::Data::String("141"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2017/testdata/day12/ex"),
        aoc::Data::String("2"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2017/day12/actual.in"),
        aoc::Data::String("171"),
        crate::part2
    );
}
