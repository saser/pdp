use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;

    #[test]
    fn example() {
        let input = "3,4,1,5";

        let lengths = crate::parse_input_as_lengths(input);
        let mut vector = vec![0, 1, 2, 3, 4];
        crate::knot_hash(&mut vector, &lengths, 0, 0);
        let product = vector[0] * vector[1];

        assert_eq!(12, product);
    }

    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/10"),
        "1980",
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Input::String(""),
        "a2582a3a0e66e6e86e3812dcb672a272",
        crate::part2
    );
    aoc::testfn!(
        example2,
        aoc::Input::String("AoC 2017"),
        "33efeb34ea91902bb2f59c9920caa6cd",
        crate::part2
    );
    aoc::testfn!(
        example3,
        aoc::Input::String("1,2,3"),
        "3efbe78a8d82f29979031a4aa0b16a9d",
        crate::part2
    );
    aoc::testfn!(
        example4,
        aoc::Input::String("1,2,4"),
        "63960835bcdc130f0b66d7ff4f6a5a8e",
        crate::part2
    );
    aoc::testfn!(
        actual,
        aoc::Input::File("adventofcode/inputs/2017/10"),
        "899124dac21012ebc32e2f4d11eaec55",
        crate::part2
    );
}
