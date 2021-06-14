use std::collections;
use std::str::FromStr;

use adventofcode_rust_aoc as aoc;
use lazy_static;
use regex;

pub fn part1(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::One)
}

pub fn part2(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::Two)
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let (players, last_marble) = parse_input(input.trim());
    match part {
        aoc::Part::One => {
            let scores = play_game(players, last_marble);
            let winner = scores.iter().max().unwrap();
            Ok(winner.to_string())
        }
        aoc::Part::Two => {
            let scores = play_game(players, last_marble * 100);
            let winner = scores.iter().max().unwrap();
            Ok(winner.to_string())
        }
    }
}

fn parse_input(input: &str) -> (usize, usize) {
    lazy_static::lazy_static! {
        static ref INPUT_RE: regex::Regex = regex::Regex::new(
            r"(?P<players>\d+) players; last marble is worth (?P<last_marble>\d+) points"
        )
        .unwrap();
    }
    let captures = INPUT_RE.captures(input).unwrap();
    let players = usize::from_str(&captures["players"]).unwrap();
    let last_marble = usize::from_str(&captures["last_marble"]).unwrap();
    (players, last_marble)
}

fn play_game(players: usize, last_marble: usize) -> Vec<usize> {
    let mut scores = vec![0; players];
    let mut ring = collections::VecDeque::new();
    ring.push_front(0);
    for marble in 1..=last_marble {
        if marble % 23 == 0 {
            for _ in 0..7 {
                let popped = ring.pop_back().unwrap();
                ring.push_front(popped);
            }
            scores[marble % players] += marble + ring.pop_front().unwrap();
        } else {
            for _ in 0..2 {
                let popped = ring.pop_front().unwrap();
                ring.push_back(popped);
            }
            ring.push_front(marble);
        }
    }
    scores
}

#[cfg(test)]
mod day09_test;
