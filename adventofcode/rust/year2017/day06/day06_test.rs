use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::String("0 2 7 0"),
        aoc::Data::String("5"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2017/06"),
        aoc::Data::String("5042"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::String("0 2 7 0"),
        aoc::Data::String("4"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2017/06"),
        aoc::Data::String("1086"),
        crate::part2
    );
}
