use adventofcode_rust_aoc as aoc;

mod part1 {
    use super::*;
    aoc::testfn!(
        example1,
        aoc::Data::String("^WNE$"),
        aoc::Data::String("3"),
        crate::part1
    );
    aoc::testfn!(
        example2,
        aoc::Data::String("^ENWWW(NEEE|SSE(EE|N))$"),
        aoc::Data::String("10"),
        crate::part1
    );
    aoc::testfn!(
        example3,
        aoc::Data::String("^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$"),
        aoc::Data::String("18"),
        crate::part1
    );
    aoc::testfn!(
        example4,
        aoc::Data::String("^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$"),
        aoc::Data::String("23"),
        crate::part1
    );
    aoc::testfn!(
        example5,
        aoc::Data::String("^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$"),
        aoc::Data::String("31"),
        crate::part1
    );
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2018/day20/actual.in"),
        aoc::Data::String("4214"),
        crate::part1
    );
}

mod part2 {
    use super::*;
    aoc::testfn!(
        actual,
        aoc::Data::File("adventofcode/data/year2018/day20/actual.in"),
        aoc::Data::String("8492"),
        crate::part2
    );
}
