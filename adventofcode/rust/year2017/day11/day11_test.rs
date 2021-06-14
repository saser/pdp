use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(example1, aoc::Input::String("ne,ne,ne"), "3", crate::part1);
    aoc::testfn!(
        example2,
        aoc::Input::String("ne,ne,sw,sw"),
        "0",
        crate::part1
    );
    aoc::testfn!(example3, aoc::Input::String("ne,ne,s,s"), "2", crate::part1);
    aoc::testfn!(
        example4,
        aoc::Input::String("se,sw,se,sw,sw"),
        "3",
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/11"),
        "761",
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/11"),
        "1542",
        crate::part2
    );
}
