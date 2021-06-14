use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;

    #[test]
    fn example() {
        let input = "s1,x3/4,pe/b";
        let mut programs = crate::generate_programs(5);
        let moves = crate::parse_input(input);
        crate::perform_moves_n(&mut programs, &moves, 1);
        let expected = "baedc";
        assert_eq!(expected, crate::programs_to_string(&programs));
    }

    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/16"),
        "kgdchlfniambejop",
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/16"),
        "fjpmholcibdgeakn",
        crate::part2
    );
}
