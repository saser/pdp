use std::ops;
use std::str::FromStr;

use adventofcode_rust_aoc as aoc;

pub fn part1(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::One)
}

pub fn part2(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::Two)
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let directions = parse_input(input.trim());
    let (final_position, furthest) = directions
        .as_slice()
        .iter()
        .map(|hex_dir| hex_dir.as_point())
        .fold((Point3D::origin(), 0), |(point, furthest), dir| {
            let new_point = point + dir;
            let new_furthest = std::cmp::max(furthest, new_point.manhattan_distance() / 2);
            (new_point, new_furthest)
        });
    match part {
        aoc::Part::One => {
            let distance = final_position.manhattan_distance() / 2;
            Ok(distance.to_string())
        }
        aoc::Part::Two => Ok(furthest.to_string()),
    }
}

fn parse_input(input: &str) -> Vec<HexDirection> {
    input
        .split(',')
        .map(HexDirection::from_str)
        .map(Result::unwrap)
        .collect()
}

#[derive(Clone, Copy, Debug, Eq, Hash, PartialEq)]
struct Point3D {
    x: i64,
    y: i64,
    z: i64,
}

impl Point3D {
    fn from(x: i64, y: i64, z: i64) -> Point3D {
        Point3D { x, y, z }
    }

    fn origin() -> Point3D {
        Point3D::from(0, 0, 0)
    }

    fn manhattan_distance(&self) -> u64 {
        (self.x.abs() + self.y.abs() + self.z.abs()) as u64
    }
}

impl ops::Add for Point3D {
    type Output = Point3D;

    fn add(self, other: Point3D) -> Point3D {
        Point3D {
            x: self.x + other.x,
            y: self.y + other.y,
            z: self.z + other.z,
        }
    }
}

#[derive(Clone, Copy, Debug, Eq, Hash, PartialEq)]
enum HexDirection {
    North,
    NorthEast,
    SouthEast,
    South,
    SouthWest,
    NorthWest,
}

impl HexDirection {
    fn as_point(self) -> Point3D {
        // A hexgrid can be represented as a "stack of boxes" in a kind of staircase pattern.
        match self {
            HexDirection::NorthEast => Point3D::from(1, 0, 1),
            HexDirection::SouthWest => Point3D::from(-1, 0, -1),
            HexDirection::North => Point3D::from(0, 1, 1),
            HexDirection::South => Point3D::from(0, -1, -1),
            HexDirection::NorthWest => Point3D::from(-1, 1, 0),
            HexDirection::SouthEast => Point3D::from(1, -1, 0),
        }
    }
}

impl FromStr for HexDirection {
    type Err = String;

    fn from_str(s: &str) -> Result<HexDirection, String> {
        match s {
            "n" => Ok(HexDirection::North),
            "ne" => Ok(HexDirection::NorthEast),
            "se" => Ok(HexDirection::SouthEast),
            "s" => Ok(HexDirection::South),
            "sw" => Ok(HexDirection::SouthWest),
            "nw" => Ok(HexDirection::NorthWest),
            _ => Err(format!("invalid hex-direction: {}", s)),
        }
    }
}

#[cfg(test)]
mod day11_test;
