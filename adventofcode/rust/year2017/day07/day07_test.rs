use adventofcode_rust_aoc as aoc;

mod parse_tests {
    use std::str::FromStr;

    #[test]
    fn program_not_holding_up() {
        let input = "pbga (66)";
        let program = crate::Program::from_str(input).unwrap();
        assert_eq!("pbga", &program.name);
        assert_eq!(66, program.weight);
    }

    #[test]
    fn program_holding_up() {
        let input = "fwft (72) -> ktlj, cntj, xhth";
        let program = crate::Program::from_str(input).unwrap();
        assert_eq!("fwft", &program.name);
        assert_eq!(72, program.weight);
        assert_eq!(&["ktlj", "cntj", "xhth"], &program.holding_up.unwrap()[..]);
    }
}

mod part1 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2017/testdata/day07/ex"),
        aoc::Data::String("tknk"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2017/07"),
        aoc::Data::String("bpvhwhh"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example,
        aoc::Data::File("adventofcode/rust/year2017/testdata/day07/ex"),
        aoc::Data::String("60"),
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/inputs/2017/07"),
        aoc::Data::String("256"),
        crate::part2
    );
}
