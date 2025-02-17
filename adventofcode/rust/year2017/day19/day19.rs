use adventofcode_rust_aoc as aoc;
use adventofcode_rust_grid as grid;

pub fn part1(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::One)
}

pub fn part2(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::Two)
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let grid = parse_input(&input);
    let traveler = TileTraveler::from(&grid);
    let (count, letters) = travel(&traveler);
    match part {
        aoc::Part::One => Ok(letters),
        aoc::Part::Two => Ok(count.to_string()),
    }
}

fn parse_input(input: &str) -> Vec<Vec<Tile>> {
    input.lines().map(parse_line).collect()
}

fn parse_line(line: &str) -> Vec<Tile> {
    line.chars().map(Tile::from).collect()
}

fn find_starting_point(grid: &[Vec<Tile>]) -> grid::Point {
    let (column, _) = grid[0]
        .iter()
        .enumerate()
        .find(|&(_, &pipe)| pipe == Tile::Vertical)
        .unwrap();
    grid::Point {
        x: column as i64,
        y: 0,
    }
}

fn travel(traveler: &TileTraveler) -> (u64, String) {
    let mut count = 1;
    let mut letters = String::new();
    for tile in traveler.tiles() {
        count += 1;
        if let Tile::Letter(c) = tile {
            letters.push(c);
        }
    }
    (count, letters)
}

#[derive(Clone, Eq, PartialEq)]
struct TileTraveler {
    grid: Vec<Vec<Tile>>,
    traveler: grid::Traveler,
}

impl TileTraveler {
    fn from(grid: &[Vec<Tile>]) -> TileTraveler {
        let grid = grid.to_owned();
        let traveler = grid::Traveler {
            pos: find_starting_point(&grid),
            direction: grid::Direction::North,
        };
        TileTraveler { grid, traveler }
    }

    fn tiles(&self) -> Tiles {
        Tiles {
            tile_traveler: self.clone(),
        }
    }
}

struct Tiles {
    tile_traveler: TileTraveler,
}

impl Iterator for Tiles {
    type Item = Tile;

    fn next(&mut self) -> Option<Tile> {
        let grid::Point { x, y } = self.tile_traveler.traveler.pos;
        let current_tile = self.tile_traveler.grid[y as usize][x as usize];

        let next_dir = if let Tile::Corner = current_tile {
            [grid::Turn::Clockwise, grid::Turn::CounterClockwise]
                .iter()
                .map(|&turn| self.tile_traveler.traveler.direction.turn(turn))
                .find(|dir| {
                    let grid::Point { x, y } = self.tile_traveler.traveler.pos + dir.as_point();
                    self.tile_traveler.grid[y as usize][x as usize] != Tile::Empty
                })
        } else {
            None
        }
        .unwrap_or(self.tile_traveler.traveler.direction);
        self.tile_traveler.traveler.direction = next_dir;

        let grid::Point {
            x: next_x,
            y: next_y,
        } = self.tile_traveler.traveler.peek_step();

        let next_tile = self.tile_traveler.grid[next_y as usize][next_x as usize];
        match next_tile {
            Tile::Empty => None,
            _ => {
                self.tile_traveler.traveler.step();
                Some(next_tile)
            }
        }
    }
}

#[derive(Clone, Copy, Debug, Eq, Hash, PartialEq)]
enum Tile {
    Empty,
    Horizontal,
    Vertical,
    Corner,
    Letter(char),
}

impl From<char> for Tile {
    fn from(c: char) -> Tile {
        match c {
            '-' => Tile::Horizontal,
            '|' => Tile::Vertical,
            '+' => Tile::Corner,
            'A'..='Z' => Tile::Letter(c),
            _ => Tile::Empty,
        }
    }
}

#[cfg(test)]
mod day19_test;
