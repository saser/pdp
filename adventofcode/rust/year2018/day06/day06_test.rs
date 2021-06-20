use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day06/ex"),
        aoc::Data::String("17"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2018/day06/actual.in"),
        aoc::Data::String("3687"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2018/day06/actual.in"),
        aoc::Data::String("40134"),
        crate::part2
    );
}
