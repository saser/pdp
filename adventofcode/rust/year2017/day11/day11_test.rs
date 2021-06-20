use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::String("ne,ne,ne"),
        aoc::Data::String("3"),
        crate::part1
    );
    aoc::testfn!(
        example2,
        aoc::Data::String("ne,ne,sw,sw"),
        aoc::Data::String("0"),
        crate::part1
    );
    aoc::testfn!(
        example3,
        aoc::Data::String("ne,ne,s,s"),
        aoc::Data::String("2"),
        crate::part1
    );
    aoc::testfn!(
        example4,
        aoc::Data::String("se,sw,se,sw,sw"),
        aoc::Data::String("3"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2017/11"),
        aoc::Data::String("761"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2017/11"),
        aoc::Data::String("1542"),
        crate::part2
    );
}
