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

#[derive(Debug, Eq, PartialEq)]
struct Particle {
    x_pos: i64,
    x_vel: i64,
    x_acc: i64,

    y_pos: i64,
    y_vel: i64,
    y_acc: i64,

    z_pos: i64,
    z_vel: i64,
    z_acc: i64,
}

impl FromStr for Particle {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        lazy_static::lazy_static! {
            static ref PARTICLE_RE: regex::Regex =
                regex::Regex::new(r"p=<(?P<x_pos>-?\d+),(?P<y_pos>-?\d+),(?P<z_pos>-?\d+)>, v=<(?P<x_vel>-?\d+),(?P<y_vel>-?\d+),(?P<z_vel>-?\d+)>, a=<(?P<x_acc>-?\d+),(?P<y_acc>-?\d+),(?P<z_acc>-?\d+)>")
                    .unwrap();
        }
        let captures = PARTICLE_RE.captures(s).unwrap();
        Ok(Particle {
            x_pos: i64::from_str(&captures["x_pos"]).map_err(|e| e.to_string())?,
            y_pos: i64::from_str(&captures["y_pos"]).map_err(|e| e.to_string())?,
            z_pos: i64::from_str(&captures["z_pos"]).map_err(|e| e.to_string())?,

            x_vel: i64::from_str(&captures["x_vel"]).map_err(|e| e.to_string())?,
            y_vel: i64::from_str(&captures["y_vel"]).map_err(|e| e.to_string())?,
            z_vel: i64::from_str(&captures["z_vel"]).map_err(|e| e.to_string())?,

            x_acc: i64::from_str(&captures["x_acc"]).map_err(|e| e.to_string())?,
            y_acc: i64::from_str(&captures["y_acc"]).map_err(|e| e.to_string())?,
            z_acc: i64::from_str(&captures["z_acc"]).map_err(|e| e.to_string())?,
        })
    }
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let particles = input
        .lines()
        .map(Particle::from_str)
        .map(Result::unwrap);
    if part == aoc::Part::One {
        // My intuition: in the long term, the position of the particle will be
        // approximately along an axis determined by the direction of the
        // acceleration vector. The length of the acceleration vector determines
        // how "quickly" the particle will go along that axis. So, the particle
        // with the shortest acceleration vector will stay the closest in the
        // long term. The positions are updated in discrete ticks in a Manhattan
        // fashion, so the length of the vector is the Manhattan length (I think
        // it's called the L1-norm).
        return Ok(particles
            .enumerate()
            .min_by_key(|(_, p)| p.x_acc.abs() + p.y_acc.abs() + p.z_acc.abs())
            .unwrap()
            .0
            .to_string());
    }
    Err(format!("no implementation for part {}", part))
}

#[cfg(test)]
mod day20_test;
