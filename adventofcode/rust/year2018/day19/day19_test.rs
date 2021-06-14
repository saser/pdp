use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/19"),
        aoc::Data::String("1344"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/19"),
        aoc::Data::String("16078144"),
        crate::part2
    );
}
