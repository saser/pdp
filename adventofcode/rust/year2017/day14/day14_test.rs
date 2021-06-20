use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::String("flqrgnkx"),
        aoc::Data::String("8108"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2017/14"),
        aoc::Data::String("8214"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::String("flqrgnkx"),
        aoc::Data::String("1242"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2017/14"),
        aoc::Data::String("1093"),
        crate::part2
    );
}
