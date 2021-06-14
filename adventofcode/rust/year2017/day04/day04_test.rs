use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Input::String("aa bb cc dd ee"),
        "1",
        crate::part1
    );
    aoc::testfn!(
        example2,
        aoc::Input::String("aa bb cc dd aa"),
        "0",
        crate::part1
    );
    aoc::testfn!(
        example3,
        aoc::Input::String("aa bb cc dd aaa"),
        "1",
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/04"),
        "337",
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Input::String("abcde fghij"),
        "1",
        crate::part2
    );
    aoc::testfn!(
        example2,
        aoc::Input::String("abcde xyz ecdab"),
        "0",
        crate::part2
    );
    aoc::testfn!(
        example3,
        aoc::Input::String("a ab abc abd abf abj"),
        "1",
        crate::part2
    );
    aoc::testfn!(
        example4,
        aoc::Input::String("iiii oiii ooii oooi oooo"),
        "1",
        crate::part2
    );
    aoc::testfn!(
        example5,
        aoc::Input::String("oiii ioii iioi iiio"),
        "0",
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/04"),
        "231",
        crate::part2
    );
}
