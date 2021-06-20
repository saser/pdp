use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;

    aoc::testfn!(
        example1,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day01/p1ex1"),
        aoc::Data::String("3"),
        crate::part1
    );
    aoc::testfn!(
        example2,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day01/p1ex2"),
        aoc::Data::String("3"),
        crate::part1
    );
    aoc::testfn!(
        example3,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day01/p1ex3"),
        aoc::Data::String("0"),
        crate::part1
    );
    aoc::testfn!(
        example4,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day01/p1ex4"),
        aoc::Data::String("-6"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/01"),
        aoc::Data::String("416"),
        crate::part1
    );
}

mod part2 {
    use super::*;

    aoc::testfn!(
        example1,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day01/p2ex1"),
        aoc::Data::String("2"),
        crate::part2
    );
    aoc::testfn!(
        example2,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day01/p2ex2"),
        aoc::Data::String("0"),
        crate::part2
    );
    aoc::testfn!(
        example3,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day01/p2ex3"),
        aoc::Data::String("10"),
        crate::part2
    );
    aoc::testfn!(
        example4,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day01/p2ex4"),
        aoc::Data::String("5"),
        crate::part2
    );
    aoc::testfn!(
        example5,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day01/p2ex5"),
        aoc::Data::String("14"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/01"),
        aoc::Data::String("56752"),
        crate::part2
    );
}
