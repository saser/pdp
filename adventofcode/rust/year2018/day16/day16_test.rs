use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/16"),
        aoc::Data::String("596"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/16"),
        aoc::Data::String("554"),
        crate::part2
    );
}
