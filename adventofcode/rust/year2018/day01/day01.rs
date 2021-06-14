use std::collections;
use std::str::FromStr;

use adventofcode_rust_aoc as aoc;

pub fn part1(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::One)
}

pub fn part2(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::Two)
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let changes = parse_input(&input);
    match part {
        aoc::Part::One => Ok(final_frequency(&changes).to_string()),
        aoc::Part::Two => Ok(first_duplicate_frequency(&changes).to_string()),
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
    let mut seen = collections::HashSet::new();
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
mod day01_test;
