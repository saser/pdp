use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day04/ex"),
        aoc::Data::String("240"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2018/day04/actual.in"),
        aoc::Data::String("125444"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day04/ex"),
        aoc::Data::String("4455"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2018/day04/actual.in"),
        aoc::Data::String("18325"),
        crate::part2
    );
}
