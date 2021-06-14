use std::cmp;
use std::collections;
use std::fmt;
use std::hash;

use adventofcode_rust_aoc as aoc;
use adventofcode_rust_grid as grid;

type Tiles = grid::Grid<Tile>;

pub fn part1(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::One)
}

pub fn part2(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::Two)
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let tiles = parse_input(&input);
    let iterations = match part {
        aoc::Part::One => 10,
        aoc::Part::Two => 1_000_000_000,
    };
    let final_tiles = run_iterations(iterations, &tiles);
    let counts = count(final_tiles.iter());
    let resource_value = counts[&Tile::Tree] * counts[&Tile::Lumberyard];
    Ok(resource_value.to_string())
}

#[derive(Clone, Copy, Debug, Eq, hash::Hash, Ord, PartialEq, PartialOrd)]
enum Tile {
    Open,
    Tree,
    Lumberyard,
}

impl Tile {
    fn next(&self, surrounding: &[Tile]) -> Self {
        let mut counts = count(surrounding.iter());
        let trees = *counts.entry(&Tile::Tree).or_insert(0);
        let lumberyards = *counts.entry(&Tile::Lumberyard).or_insert(0);
        match *self {
            Tile::Open => {
                if trees >= 3 {
                    Tile::Tree
                } else {
                    *self
                }
            }
            Tile::Tree => {
                if lumberyards >= 3 {
                    Tile::Lumberyard
                } else {
                    *self
                }
            }
            Tile::Lumberyard => {
                if lumberyards >= 1 && trees >= 1 {
                    *self
                } else {
                    Tile::Open
                }
            }
        }
    }
}

impl fmt::Display for Tile {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        let c = match *self {
            Tile::Open => '.',
            Tile::Tree => '|',
            Tile::Lumberyard => '#',
        };
        write!(f, "{}", c)
    }
}

impl From<char> for Tile {
    fn from(c: char) -> Self {
        match c {
            '|' => Tile::Tree,
            '#' => Tile::Lumberyard,
            _ => Tile::Open,
        }
    }
}

fn parse_input(input: &str) -> Tiles {
    let chars = input
        .lines()
        .map(|line| line.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();
    let nrows = chars.len();
    let ncols = chars[0].len();
    let mut grid = Tiles::with(nrows, ncols, &Tile::Open);
    for row in 0..nrows {
        for col in 0..ncols {
            grid[(row, col)] = Tile::from(chars[row][col]);
        }
    }
    grid
}

fn surrounding(row: usize, col: usize, tiles: &Tiles) -> Vec<Tile> {
    let start_row = cmp::max(0, row as isize - 1) as usize;
    let start_col = cmp::max(0, col as isize - 1) as usize;
    let end_row = cmp::min((tiles.nrows() - 1) as isize, (row + 1) as isize) as usize;
    let end_col = cmp::min((tiles.ncols() - 1) as isize, (col + 1) as isize) as usize;

    let mut surrounding = Vec::new();
    for s_row in start_row..=end_row {
        for s_col in start_col..=end_col {
            if s_row == row && s_col == col {
                continue;
            }
            surrounding.push(tiles[(s_row, s_col)]);
        }
    }
    surrounding
}

fn count<T, I>(iter: I) -> collections::HashMap<T, usize>
where
    T: Eq + hash::Hash + Copy,
    I: Iterator<Item = T>,
{
    let mut map = collections::HashMap::new();
    for item in iter {
        *map.entry(item).or_insert(0) += 1;
    }
    map
}

fn iteration(tiles: &Tiles) -> Tiles {
    let mut new_tiles = tiles.clone();
    for row in 0..tiles.nrows() {
        for col in 0..tiles.ncols() {
            let tile = tiles[(row, col)];
            let surrounding = surrounding(row, col, tiles);
            new_tiles[(row, col)] = tile.next(&surrounding);
        }
    }
    new_tiles
}

fn run_iterations(iterations: usize, tiles: &Tiles) -> Tiles {
    let mut current_tiles = tiles.clone();
    let mut seen = collections::HashMap::new();
    seen.insert(current_tiles.clone(), 0);
    for i in 1..=iterations {
        let new_tiles = iteration(&current_tiles);
        if let Some(seen_i) = seen.get(&new_tiles) {
            let loop_length = i - seen_i;
            let iterations_left = (iterations - seen_i) % loop_length;
            return run_iterations(iterations_left, &new_tiles);
        }
        seen.insert(new_tiles.clone(), i);
        current_tiles = new_tiles;
    }
    current_tiles
}

#[cfg(test)]
mod day18_test;
