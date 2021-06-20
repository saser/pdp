use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day07/ex"),
        aoc::Data::String("CABDFE"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/07"),
        aoc::Data::String("MNQKRSFWGXPZJCOTVYEBLAHIUD"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2018/testdata/day07/ex"),
        aoc::Data::String("253"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2018/07"),
        aoc::Data::String("948"),
        crate::part2
    );
}
