use std::str::FromStr;

use adventofcode_rust_aoc as aoc;

pub fn part1(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::One)
}

pub fn part2(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::Two)
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let fun = match part {
        aoc::Part::One => min_max,
        aoc::Part::Two => divisors,
    };
    Ok(input
        .lines()
        .map(parse_line)
        .map(|v| fun(&v))
        .sum::<u32>()
        .to_string())
}

fn parse_line(line: &str) -> Vec<u32> {
    line.split_whitespace()
        .map(u32::from_str)
        .map(Result::unwrap)
        .collect()
}

fn min_max(nums: &[u32]) -> u32 {
    // Can also be solved using the `Iterator::max` and `Iterator::min` methods, but that's no fun.

    let mut min = nums[0];
    let mut max = nums[0];
    for &n in nums.iter() {
        if n < min {
            min = n;
        } else if n > max {
            max = n;
        }
    }
    max - min
}

fn divisors(nums: &[u32]) -> u32 {
    // This is a very bad, brute-force approach to the problem. But I can't think of any other way
    // to do it that is more optimal.

    for (i, &x) in nums.iter().enumerate() {
        for (j, &y) in nums.iter().enumerate() {
            if i == j {
                continue;
            };

            if x % y == 0 {
                return x / y;
            } else if y % x == 0 {
                return y / x;
            }
        }
    }

    // This is unreachable since the input is guaranteed to have exactly one pair of numbers that
    // divide each other, so the function is guaranteed to return in the above loops.
    unreachable!()
}


#[cfg(test)]
mod day02_test;
