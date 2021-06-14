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
    let claims = input
        .lines()
        .map(Claim::from_str)
        .map(Result::unwrap)
        .collect::<Vec<Claim>>();
    let map = build_map(&claims);
    match part {
        aoc::Part::One => {
            let count = map
                .values()
                .filter(|point_claims| point_claims.len() > 1)
                .count();
            Ok(count.to_string())
        }
        aoc::Part::Two => {
            let mut candidate_claims = map
                .values()
                .filter(|point_claims| point_claims.len() == 1)
                .map(|point_claims| point_claims[0]);
            let lonely_claim = candidate_claims
                .find(|&claim| {
                    claim.covered_points().iter().all(|point| {
                        let point_claims = &map[point];
                        point_claims.len() == 1 && point_claims[0] == claim
                    })
                })
                .unwrap();
            Ok(lonely_claim.id.to_string())
        }
    }
}

#[derive(Clone, Copy, Debug, Eq, PartialEq)]
struct Claim {
    id: usize,
    x: usize,
    y: usize,
    dx: usize,
    dy: usize,
}

impl Claim {
    fn covered_points(&self) -> Vec<(usize, usize)> {
        let mut points = Vec::with_capacity(self.dx * self.dy);
        for i in self.x..self.x + self.dx {
            for j in self.y..self.y + self.dy {
                points.push((i, j));
            }
        }
        points
    }
}

impl FromStr for Claim {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        lazy_static::lazy_static! {
            static ref CLAIM_RE: regex::Regex =
                regex::Regex::new(r"#(?P<id>\d+) @ (?P<x>\d+),(?P<y>\d+): (?P<dx>\d+)x(?P<dy>\d+)")
                    .unwrap();
        }
        let captures = CLAIM_RE.captures(s).unwrap();
        let id = usize::from_str(&captures["id"]).unwrap();
        let x = usize::from_str(&captures["x"]).unwrap();
        let y = usize::from_str(&captures["y"]).unwrap();
        let dx = usize::from_str(&captures["dx"]).unwrap();
        let dy = usize::from_str(&captures["dy"]).unwrap();
        Ok(Self { id, x, y, dx, dy })
    }
}

fn build_map(claims: &[Claim]) -> collections::HashMap<(usize, usize), Vec<&Claim>> {
    let mut map = collections::HashMap::new();
    for claim in claims {
        for i in claim.x..claim.x + claim.dx {
            for j in claim.y..claim.y + claim.dy {
                let point_claims = map.entry((i, j)).or_insert_with(Vec::new);
                point_claims.push(claim);
            }
        }
    }
    map
}

#[cfg(test)]
mod day03_test;
