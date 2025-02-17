use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day15/ex1"),
        aoc::Data::String("27730"),
        crate::part1
    );
    aoc::testfn!(
        example2,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day15/ex2"),
        aoc::Data::String("36334"),
        crate::part1
    );
    aoc::testfn!(
        example3,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day15/ex3"),
        aoc::Data::String("39514"),
        crate::part1
    );
    aoc::testfn!(
        example4,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day15/ex4"),
        aoc::Data::String("27755"),
        crate::part1
    );
    aoc::testfn!(
        example5,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day15/ex5"),
        aoc::Data::String("28944"),
        crate::part1
    );
    aoc::testfn!(
        example6,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day15/ex6"),
        aoc::Data::String("18740"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2018/day15/actual.in"),
        aoc::Data::String("201638"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day15/ex1"),
        aoc::Data::String("4988"),
        crate::part2
    );
    aoc::testfn!(
        example3,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day15/ex3"),
        aoc::Data::String("31284"),
        crate::part2
    );
    aoc::testfn!(
        example4,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day15/ex4"),
        aoc::Data::String("3478"),
        crate::part2
    );
    aoc::testfn!(
        example5,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day15/ex5"),
        aoc::Data::String("6474"),
        crate::part2
    );
    aoc::testfn!(
        example6,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day15/ex6"),
        aoc::Data::String("1140"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2018/day15/actual.in"),
        aoc::Data::String("95764"),
        crate::part2
    );
}
