use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2017/testdata/day05/ex"),
        aoc::Data::String("5"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2017/day05/actual.in"),
        aoc::Data::String("358131"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2017/testdata/day05/ex"),
        aoc::Data::String("10"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2017/day05/actual.in"),
        aoc::Data::String("25558839"),
        crate::part2
    );
}
