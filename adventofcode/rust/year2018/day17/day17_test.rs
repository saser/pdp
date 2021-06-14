use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day17/ex"),
        aoc::Data::String("57"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/17"),
        aoc::Data::String("31471"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day17/ex"),
        aoc::Data::String("29"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/17"),
        aoc::Data::String("24169"),
        crate::part2
    );
}
