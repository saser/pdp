use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day13/ex1"),
        aoc::Data::String("7,3"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/13"),
        aoc::Data::String("16,45"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day13/ex2"),
        aoc::Data::String("6,4"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/13"),
        aoc::Data::String("21,91"),
        crate::part2
    );
}
