use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2017/testdata/day08/ex"),
        aoc::Data::String("1"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2017/day08/actual.in"),
        aoc::Data::String("6012"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2017/testdata/day08/ex"),
        aoc::Data::String("10"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2017/day08/actual.in"),
        aoc::Data::String("6369"),
        crate::part2
    );
}
