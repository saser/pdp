use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2018/day19/actual.in"),
        aoc::Data::String("1344"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2018/day19/actual.in"),
        aoc::Data::String("16078144"),
        crate::part2
    );
}
