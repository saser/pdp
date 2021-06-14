use std::collections;

use adventofcode_rust_aoc as aoc;
use adventofcode_rust_grid as grid;

pub fn part1(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::One)
}

pub fn part2(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::Two)
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let number = parse_input(input.trim());
    match part {
        aoc::Part::One => Ok(distance_to_center(number).to_string()),
        aoc::Part::Two => Ok(first_after_number_by_summing(number).to_string()),
    }
}

fn parse_input(input: &str) -> u64 {
    str::parse(input).unwrap()
}

struct SpiralTraveler {
    traveler: grid::Traveler,
    current_layer: u64,
    steps_to_corners: Vec<u64>,
    to_corner: u64,
    to_next_layer: u64,
}

impl SpiralTraveler {
    fn new() -> SpiralTraveler {
        let mut traveler = grid::Traveler::default();
        traveler.direction = grid::Direction::East;
        SpiralTraveler {
            traveler,
            current_layer: 1,
            steps_to_corners: Vec::with_capacity(4),
            to_corner: 1,
            to_next_layer: 1,
        }
    }
}

impl Iterator for SpiralTraveler {
    type Item = grid::Point;

    fn next(&mut self) -> Option<grid::Point> {
        self.traveler.step();

        self.to_next_layer -= 1;
        if self.to_next_layer == 0 {
            self.current_layer += 1;
            let layer_side = 2 * self.current_layer - 1;
            let steps = layer_side - 1;
            self.steps_to_corners
                .append(&mut vec![steps + 1, steps, steps, steps - 1]);
            self.to_next_layer = 4 * steps;
        }

        self.to_corner -= 1;
        if self.to_corner == 0 {
            self.traveler.turn(grid::Turn::CounterClockwise);
            self.to_corner = self.steps_to_corners.pop().unwrap();
        }
        Some(self.traveler.pos)
    }
}

struct Spiral {
    spiral_traveler: SpiralTraveler,
    grid: collections::HashMap<grid::Point, u64>,
    pos: grid::Point,
    value: u64,
    next_value_fun: Box<dyn Fn(&Spiral) -> u64>,
}

impl Spiral {
    fn new(next_value_fun: Box<dyn Fn(&Spiral) -> u64>) -> Spiral {
        let initial_pos = grid::Point { x: 0, y: 0 };
        let mut initial_grid = collections::HashMap::new();
        initial_grid.insert(initial_pos, 1);
        Spiral {
            spiral_traveler: SpiralTraveler::new(),
            grid: initial_grid,
            pos: initial_pos,
            value: 1,
            next_value_fun,
        }
    }
}

impl Iterator for Spiral {
    type Item = (grid::Point, u64);

    fn next(&mut self) -> Option<(grid::Point, u64)> {
        let ret = (self.pos, self.value);
        self.pos = self.spiral_traveler.next().unwrap();
        self.value = (self.next_value_fun)(self);
        self.grid.insert(self.pos, self.value);
        Some(ret)
    }
}

fn first_after_number_by_summing(target_number: u64) -> u64 {
    let mut spiral = Spiral::new(Box::new(|spiral| {
        let current_pos = spiral.pos;
        let deltas = [
            grid::Point { x: 1, y: 0 },
            grid::Point { x: 1, y: 1 },
            grid::Point { x: 0, y: 1 },
            grid::Point { x: -1, y: 1 },
            grid::Point { x: -1, y: 0 },
            grid::Point { x: -1, y: -1 },
            grid::Point { x: 0, y: -1 },
            grid::Point { x: 1, y: -1 },
        ];
        deltas
            .iter()
            .map(|&delta| current_pos + delta)
            .map(|pos| spiral.grid.get(&pos))
            .map(|op_val| op_val.unwrap_or(&0))
            .sum()
    }));
    spiral.find(|&(_, value)| value > target_number).unwrap().1
}

fn distance_to_center(target_number: u64) -> u64 {
    let mut spiral = Spiral::new(Box::new(|spiral| spiral.value + 1));
    let (pos, _) = spiral.find(|&(_, value)| value == target_number).unwrap();
    pos.manhattan_distance()
}

#[cfg(test)]
mod day03_test;
