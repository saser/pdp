use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day10/ex.in"),
        aoc::Data::File("adventofcode/rust/year2018/testdata/day10/ex.out"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/10"),
        aoc::Data::File("adventofcode/rust/year2018/testdata/day10/actual.out"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day10/ex.in"),
        aoc::Data::String("3"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/10"),
        aoc::Data::String("10355"),
        crate::part2
    );
}
