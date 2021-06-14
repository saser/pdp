use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::String("18"),
        aoc::Data::String("33,45"),
        crate::part1
    );
    aoc::testfn!(
        example2,
        aoc::Data::String("42"),
        aoc::Data::String("21,61"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/11"),
        aoc::Data::String("233,36"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::String("18"),
        aoc::Data::String("90,269,16"),
        crate::part2
    );
    aoc::testfn!(
        example2,
        aoc::Data::String("42"),
        aoc::Data::String("232,251,12"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/11"),
        aoc::Data::String("231,107,14"),
        crate::part2
    );
}
