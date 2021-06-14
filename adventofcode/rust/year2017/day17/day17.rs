use std::str::FromStr;

use adventofcode_rust_aoc as aoc;

pub fn part1(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::One)
}

pub fn part2(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::Two)
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let length = parse_input(input.trim());
    match part {
        aoc::Part::One => {
            let (vec, final_position) = build_ring_buffer(2017, length);
            Ok(vec[final_position + 1].to_string())
        }
        aoc::Part::Two => Ok(value_after_zero(50_000_000, length).to_string()),
    }
}

fn parse_input(input: &str) -> usize {
    usize::from_str(input).unwrap()
}

fn build_ring_buffer(final_value: usize, length: usize) -> (Vec<usize>, usize) {
    let mut vec = Vec::with_capacity(final_value + 1);
    vec.push(0);
    let mut current_position = 0;
    for i in 1..=final_value {
        let index_to_insert = ((current_position + length) % i) + 1;
        vec.insert(index_to_insert, i);
        current_position = index_to_insert;
    }
    (vec, current_position)
}

fn value_after_zero(final_value: usize, length: usize) -> usize {
    let mut index_for_zero = 0;
    let mut value_after_zero = 0;
    let mut current_position = index_for_zero;
    for i in 1..final_value {
        let index_to_insert = ((current_position + length) % i) + 1;
        if index_to_insert <= index_for_zero {
            index_for_zero += 1;
        } else if index_to_insert == index_for_zero + 1 {
            value_after_zero = i;
        }
        current_position = index_to_insert;
    }
    value_after_zero
}

#[cfg(test)]
mod day17_test;
