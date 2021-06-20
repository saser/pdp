use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2017/testdata/day15/ex"),
        aoc::Data::String("588"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2017/day15/actual.in"),
        aoc::Data::String("609"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2017/testdata/day15/ex"),
        aoc::Data::String("309"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2017/day15/actual.in"),
        aoc::Data::String("253"),
        crate::part2
    );
}
