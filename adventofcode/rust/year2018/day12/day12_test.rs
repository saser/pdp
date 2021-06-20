use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day12/ex"),
        aoc::Data::String("325"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2018/day12/actual.in"),
        aoc::Data::String("3221"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2018/day12/actual.in"),
        aoc::Data::String("2600000001872"),
        crate::part2
    );
}
