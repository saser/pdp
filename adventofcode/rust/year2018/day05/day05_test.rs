use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::String("aA"),
        aoc::Data::String("0"),
        crate::part1
    );
    aoc::testfn!(
        example2,
        aoc::Data::String("abBA"),
        aoc::Data::String("0"),
        crate::part1
    );
    aoc::testfn!(
        example3,
        aoc::Data::String("abAB"),
        aoc::Data::String("4"),
        crate::part1
    );
    aoc::testfn!(
        example4,
        aoc::Data::String("aabAAB"),
        aoc::Data::String("6"),
        crate::part1
    );
    aoc::testfn!(
        example5,
        aoc::Data::String("dabAcCaCBAcCcaDA"),
        aoc::Data::String("10"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/05"),
        aoc::Data::String("9686"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::String("dabAcCaCBAcCcaDA"),
        aoc::Data::String("4"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/05"),
        aoc::Data::String("5524"),
        crate::part2
    );
}
