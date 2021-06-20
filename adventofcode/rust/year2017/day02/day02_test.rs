use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::String("5 1 9 5"),
        aoc::Data::String("8"),
        crate::part1
    );
    aoc::testfn!(
        example2,
        aoc::Data::String("7 5 3"),
        aoc::Data::String("4"),
        crate::part1
    );
    aoc::testfn!(
        example3,
        aoc::Data::String("2 4 6 8"),
        aoc::Data::String("6"),
        crate::part1
    );
    aoc::testfn!(
        example_all,
        aoc::Data::File("adventofcode/rust/year2017/testdata/day02/p1ex"),
        aoc::Data::String("18"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2017/02"),
        aoc::Data::String("36766"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::String("5 9 2 8"),
        aoc::Data::String("4"),
        crate::part2
    );
    aoc::testfn!(
        example2,
        aoc::Data::String("9 4 7 3"),
        aoc::Data::String("3"),
        crate::part2
    );
    aoc::testfn!(
        example3,
        aoc::Data::String("3 8 6 5"),
        aoc::Data::String("2"),
        crate::part2
    );
    aoc::testfn!(
        example_all,
        aoc::Data::File("adventofcode/rust/year2017/testdata/day02/p2ex"),
        aoc::Data::String("9"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2017/02"),
        aoc::Data::String("261"),
        crate::part2
    );
}
