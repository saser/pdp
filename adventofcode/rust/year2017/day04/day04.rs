use std::collections;

use adventofcode_rust_aoc as aoc;

pub fn part1(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::One)
}

pub fn part2(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::Two)
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let passphrases = parse_input(&input);
    let validator = match part {
        aoc::Part::One => contains_unique_passwords,
        aoc::Part::Two => contains_no_anagrams,
    };
    Ok(count_valid(validator, &passphrases).to_string())
}

fn parse_input(input: &str) -> Vec<Vec<String>> {
    input
        .lines()
        .map(|line| line.split_whitespace())
        .map(|iter| iter.map(String::from))
        .map(|iter| iter.collect())
        .collect()
}

fn count_valid(validator: fn(&[String]) -> bool, passphrases: &[Vec<String>]) -> usize {
    passphrases
        .iter()
        .filter(|&phrase| validator(&phrase))
        .count()
}

fn contains_unique_passwords(passphrase: &[String]) -> bool {
    let words = passphrase.len();
    let set: collections::HashSet<String> = passphrase.iter().cloned().collect();
    set.len() == words
}

fn contains_no_anagrams(passphrase: &[String]) -> bool {
    fn sort_string(s: &str) -> String {
        let mut chars: Vec<char> = s.chars().collect();
        chars.sort();
        chars.into_iter().collect()
    }
    let words = passphrase.len();
    let set: collections::HashSet<String> =
        passphrase.into_iter().map(|s| sort_string(&s)).collect();
    set.len() == words
}

#[cfg(test)]
mod day04_test;
