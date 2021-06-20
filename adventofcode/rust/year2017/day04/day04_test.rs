use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::String("aa bb cc dd ee"),
        aoc::Data::String("1"),
        crate::part1
    );
    aoc::testfn!(
        example2,
        aoc::Data::String("aa bb cc dd aa"),
        aoc::Data::String("0"),
        crate::part1
    );
    aoc::testfn!(
        example3,
        aoc::Data::String("aa bb cc dd aaa"),
        aoc::Data::String("1"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2017/04"),
        aoc::Data::String("337"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::String("abcde fghij"),
        aoc::Data::String("1"),
        crate::part2
    );
    aoc::testfn!(
        example2,
        aoc::Data::String("abcde xyz ecdab"),
        aoc::Data::String("0"),
        crate::part2
    );
    aoc::testfn!(
        example3,
        aoc::Data::String("a ab abc abd abf abj"),
        aoc::Data::String("1"),
        crate::part2
    );
    aoc::testfn!(
        example4,
        aoc::Data::String("iiii oiii ooii oooi oooo"),
        aoc::Data::String("1"),
        crate::part2
    );
    aoc::testfn!(
        example5,
        aoc::Data::String("oiii ioii iioi iiio"),
        aoc::Data::String("0"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2017/04"),
        aoc::Data::String("231"),
        crate::part2
    );
}
