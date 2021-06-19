use std::collections::HashSet;
use std::io;
use std::str::FromStr;

use crate::base::Part;

pub fn part1(r: &mut dyn io::Read) -> Result<String, String> {
    solve(r, Part::One)
}

pub fn part2(r: &mut dyn io::Read) -> Result<String, String> {
    solve(r, Part::Two)
}

fn solve(r: &mut dyn io::Read, part: Part) -> Result<String, String> {
    let mut input = String::new();
    r.read_to_string(&mut input).map_err(|e| e.to_string())?;
    let changes = parse_input(&input);
    match part {
        Part::One => Ok(final_frequency(&changes).to_string()),
        Part::Two => Ok(first_duplicate_frequency(&changes).to_string()),
    }
}

fn parse_input(input: &str) -> Vec<i64> {
    input
        .lines()
        .map(|line| i64::from_str(line).unwrap())
        .collect()
}

fn final_frequency(changes: &[i64]) -> i64 {
    changes.iter().sum()
}

fn first_duplicate_frequency(changes: &[i64]) -> i64 {
    let looped_frequencies = changes.iter().cycle().scan(0, |acc, &x| {
        *acc += x;
        Some(*acc)
    });
    let mut seen = HashSet::new();
    seen.insert(0);
    for freq in looped_frequencies {
        if seen.contains(&freq) {
            return freq;
        }
        seen.insert(freq);
    }
    unreachable!()
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::test;

    mod part1 {
        use super::*;

        test!(example1, file env!("YEAR2018_DAY01_P1EX1"), "3", part1);
        test!(example2, file env!("YEAR2018_DAY01_P1EX2"), "3", part1);
        test!(example3, file env!("YEAR2018_DAY01_P1EX3"), "0", part1);
        test!(example4, file env!("YEAR2018_DAY01_P1EX4"), "-6", part1);
        test!(actual, file env!("YEAR2018_DAY01"), "416", part1);
    }

    mod part2 {
        use super::*;

        test!(example1, file env!("YEAR2018_DAY01_P2EX1"), "2", part2);
        test!(example2, file env!("YEAR2018_DAY01_P2EX2"), "0", part2);
        test!(example3, file env!("YEAR2018_DAY01_P2EX3"), "10", part2);
        test!(example4, file env!("YEAR2018_DAY01_P2EX4"), "5", part2);
        test!(example5, file env!("YEAR2018_DAY01_P2EX5"), "14", part2);
        test!(actual, file env!("YEAR2018_DAY01"), "56752", part2);
    }
}
