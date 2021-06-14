use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::String("9"),
        aoc::Data::String("5158916779"),
        crate::part1
    );
    aoc::testfn!(
        example2,
        aoc::Data::String("5"),
        aoc::Data::String("0124515891"),
        crate::part1
    );
    aoc::testfn!(
        example3,
        aoc::Data::String("18"),
        aoc::Data::String("9251071085"),
        crate::part1
    );
    aoc::testfn!(
        example4,
        aoc::Data::String("2018"),
        aoc::Data::String("5941429882"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/14"),
        aoc::Data::String("5371393113"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::String("51589"),
        aoc::Data::String("9"),
        crate::part2
    );
    aoc::testfn!(
        example2,
        aoc::Data::String("01245"),
        aoc::Data::String("5"),
        crate::part2
    );
    aoc::testfn!(
        example3,
        aoc::Data::String("92510"),
        aoc::Data::String("18"),
        crate::part2
    );
    aoc::testfn!(
        example4,
        aoc::Data::String("59414"),
        aoc::Data::String("2018"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/14"),
        aoc::Data::String("20286858"),
        crate::part2
    );
}
