use std::str::FromStr;

use adventofcode_rust_aoc as aoc;

const MUL_A: u64 = 16807;
const MUL_B: u64 = 48271;
const MOD: u64 = 2_147_483_647;

pub fn part1(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::One)
}

pub fn part2(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::Two)
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let (start_a, start_b) = parse_input(&input);
    let (repetitions, constraints) = match part {
        aoc::Part::One => (40_000_000, (1, 1)),
        aoc::Part::Two => (5_000_000, (4, 8)),
    };
    let matching_pairs = (0..repetitions)
        .scan((start_a, start_b), |pair, _| {
            *pair = next_values(*pair, constraints);
            Some(*pair)
        })
        //.inspect(|pair| println!("pair: {:?}", pair))
        .filter(|&pair| lowest_16_bits_matching(pair))
        .count();
    Ok(matching_pairs.to_string())
}

fn parse_input(input: &str) -> (u64, u64) {
    let lines: Vec<&str> = input.lines().collect();
    (parse_line(&lines[0]), parse_line(&lines[1]))
}

fn parse_line(line: &str) -> u64 {
    let parts: Vec<&str> = line.split(' ').collect();
    u64::from_str(parts[parts.len() - 1]).unwrap()
}

fn next_values(
    (value_a, value_b): (u64, u64),
    (constraint_a, constraint_b): (u64, u64),
) -> (u64, u64) {
    let mut next_value_a = (value_a * MUL_A) % MOD;
    while next_value_a % constraint_a != 0 {
        next_value_a = (next_value_a * MUL_A) % MOD;
    }
    let mut next_value_b = (value_b * MUL_B) % MOD;
    while next_value_b % constraint_b != 0 {
        next_value_b = (next_value_b * MUL_B) % MOD;
    }
    (next_value_a, next_value_b)
}

fn lowest_16_bits_matching((value_a, value_b): (u64, u64)) -> bool {
    value_a & 0xFFFF == value_b & 0xFFFF
}

#[cfg(test)]
mod day15_test;
