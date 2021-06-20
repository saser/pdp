use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2017/testdata/day18/p1ex"),
        aoc::Data::String("4"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2017/day18/actual.in"),
        aoc::Data::String("3188"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2017/testdata/day18/p2ex"),
        aoc::Data::String("3"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2017/day18/actual.in"),
        aoc::Data::String("7112"),
        crate::part2
    );
}
