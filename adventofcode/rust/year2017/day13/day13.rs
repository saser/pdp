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
    let layers = parse_input(&input);
    match part {
        aoc::Part::One => {
            let total_severity: u64 = layers
                .iter()
                .map(|(&layer, &depth)| severity(layer, depth, 0))
                .sum();
            Ok(total_severity.to_string())
        }
        aoc::Part::Two => Ok(find_min_delay(&layers).to_string()),
    }
}

fn parse_input(input: &str) -> collections::HashMap<u64, u64> {
    input.lines().map(parse_line).collect()
}

fn parse_line(line: &str) -> (u64, u64) {
    let parts: Vec<&str> = line.split(": ").collect();
    let layer = u64::from_str(parts[0]).unwrap();
    let depth = u64::from_str(parts[1]).unwrap();
    (layer, depth)
}

fn detected_when_entering(picosecond: u64, depth: u64, delay: u64) -> bool {
    (picosecond + delay) % ((depth - 1) * 2) == 0
}

fn severity(layer: u64, depth: u64, delay: u64) -> u64 {
    if detected_when_entering(layer, depth, delay) {
        layer * depth
    } else {
        0
    }
}

fn any_detection_with_delay(layers: &collections::HashMap<u64, u64>, delay: u64) -> bool {
    layers
        .iter()
        .any(|(&layer, &depth)| detected_when_entering(layer, depth, delay))
}

fn find_min_delay(layers: &collections::HashMap<u64, u64>) -> u64 {
    (0..)
        .find(|&delay| !any_detection_with_delay(layers, delay))
        .unwrap()
}

#[cfg(test)]
mod day13_test;
