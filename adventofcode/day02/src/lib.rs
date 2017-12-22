extern crate base;

use base::{Part, Solver};

use std::str::FromStr;

struct Day02;

pub fn get_solver() -> Box<Solver> {
    Box::new(Day02)
}

impl Solver for Day02 {
    fn solve(&self, part: Part, input: &str) -> Result<String, String> {
        let fun = match part {
            Part::One => min_max,
            Part::Two => divisors,
        };
        Ok(input.lines()
            .map(parse_line)
            .map(|v| fun(&v))
            .sum::<u32>()
            .to_string())
    }
}

fn parse_line(line: &str) -> Vec<u32> {
    line.split_whitespace()
        .map(u32::from_str)
        .map(Result::unwrap)
        .collect()
}

fn min_max(nums: &[u32]) -> u32 {
    unimplemented!()
}

fn divisors(nums: &[u32]) -> u32 {
    unimplemented!()
}

#[cfg(test)]
mod tests {
    use super::*;

    mod part1 {
        use super::*;

        #[test]
        fn example_1() {
            let solver = get_solver();
            let input = "5 1 9 5";
            let expected = "8";
            assert_eq!(expected, solver.solve(Part::One, input).unwrap());
        }

        #[test]
        fn example_2() {
            let solver = get_solver();
            let input = "7 5 3";
            let expected = "4";
            assert_eq!(expected, solver.solve(Part::One, input).unwrap());
        }

        #[test]
        fn example_3() {
            let solver = get_solver();
            let input = "2 4 6 8";
            let expected = "6";
            assert_eq!(expected, solver.solve(Part::One, input).unwrap());
        }

        #[test]
        fn example_all() {
            let solver = get_solver();
            let input = "\
5 1 9 5
7 5 3
2 4 6 8\
            ";
            let expected = "18";
            assert_eq!(expected, solver.solve(Part::One, input).unwrap());
        }
    }

    mod part2 {
        use super::*;

        #[test]
        fn example_1() {
            let solver = get_solver();
            let input = "5 9 2 8";
            let expected = "4";
            assert_eq!(expected, solver.solve(Part::Two, input).unwrap());
        }

        #[test]
        fn example_2() {
            let solver = get_solver();
            let input = "9 4 7 3";
            let expected = "3";
            assert_eq!(expected, solver.solve(Part::Two, input).unwrap());
        }

        #[test]
        fn example_3() {
            let solver = get_solver();
            let input = "3 8 6 5";
            let expected = "2";
            assert_eq!(expected, solver.solve(Part::Two, input).unwrap());
        }

        #[test]
        fn example_all() {
            let solver = get_solver();
            let input = "\
5 9 2 8
9 4 7 3
3 8 6 5\
            ";
            let expected = "4";
            assert_eq!(expected, solver.solve(Part::Two, input).unwrap());
        }
    }

}
