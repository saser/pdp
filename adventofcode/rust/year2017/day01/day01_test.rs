use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::String("1122"),
        aoc::Data::String("3"),
        crate::part1
    );
    aoc::testfn!(
        example2,
        aoc::Data::String("1111"),
        aoc::Data::String("4"),
        crate::part1
    );
    aoc::testfn!(
        example3,
        aoc::Data::String("1234"),
        aoc::Data::String("0"),
        crate::part1
    );
    aoc::testfn!(
        example4,
        aoc::Data::String("91212129"),
        aoc::Data::String("9"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2017/01"),
        aoc::Data::String("1044"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::String("1212"),
        aoc::Data::String("6"),
        crate::part2
    );
    aoc::testfn!(
        example2,
        aoc::Data::String("1221"),
        aoc::Data::String("0"),
        crate::part2
    );
    aoc::testfn!(
        example3,
        aoc::Data::String("123425"),
        aoc::Data::String("4"),
        crate::part2
    );
    aoc::testfn!(
        example4,
        aoc::Data::String("123123"),
        aoc::Data::String("12"),
        crate::part2
    );
    aoc::testfn!(
        example5,
        aoc::Data::String("12131415"),
        aoc::Data::String("4"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2017/01"),
        aoc::Data::String("1054"),
        crate::part2
    );
}
