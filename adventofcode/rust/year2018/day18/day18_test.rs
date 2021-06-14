use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day18/ex"),
        aoc::Data::String("1147"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/18"),
        aoc::Data::String("545600"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/18"),
        aoc::Data::String("202272"),
        crate::part2
    );
}
