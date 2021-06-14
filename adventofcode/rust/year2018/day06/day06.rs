use std::collections;
use std::str::FromStr;

use adventofcode_rust_aoc as aoc;
use adventofcode_rust_grid as grid;
use rayon::iter::IntoParallelRefIterator;
use rayon::iter::ParallelIterator;

type Coordinates = collections::HashMap<char, grid::Point>;
type Distances = collections::HashMap<char, u64>;

pub fn part1(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::One)
}

pub fn part2(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::Two)
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let coordinates = parse_input(input.trim());
    let bb = BoundingBox::from_points(
        coordinates
            .values()
            .cloned()
            .collect::<Vec<grid::Point>>()
            .as_slice(),
    );
    let distances = bounding_box_distances(&bb, &coordinates);
    match part {
        aoc::Part::One => {
            let minimal_distances = all_minimal_distances(&distances);
            let edge_points = bb.edge_points();
            let mut infinite_coordinates = collections::HashSet::new();
            let mut closest_points = collections::HashMap::new();
            for (point, coordinates) in minimal_distances.iter() {
                if coordinates.len() == 1 {
                    let c = coordinates[0];
                    closest_points.entry(c).or_insert_with(Vec::new).push(point);
                    if edge_points.contains(point) {
                        infinite_coordinates.insert(c);
                    }
                }
            }
            let max_area = closest_points
                .par_iter()
                .filter(|(c, _points)| !infinite_coordinates.contains(c))
                .map(|(_c, points)| points.len())
                .max()
                .unwrap();
            Ok(max_area.to_string())
        }
        aoc::Part::Two => {
            let limit = 10_000;
            let count = distances
                .par_iter()
                .map(|(_point, distance_map)| distance_map.values().cloned().collect::<Vec<u64>>())
                .map(|ds| ds.iter().sum::<u64>())
                .filter(|&sum| sum < limit)
                .count();
            Ok(count.to_string())
        }
    }
}

fn parse_input(input: &str) -> Coordinates {
    let alphabet = (b'A'..).map(char::from);
    let points = input.lines().map(grid::Point::from_str).map(Result::unwrap);
    alphabet.zip(points).collect()
}

#[derive(Clone, Copy, Debug, Eq, Hash, PartialEq)]
struct BoundingBox {
    x_min: i64,
    y_min: i64,
    x_max: i64,
    y_max: i64,
}

impl BoundingBox {
    fn from_points(points: &[grid::Point]) -> Self {
        let first_point = points[0];
        let mut x_min = first_point.x;
        let mut x_max = first_point.x;
        let mut y_min = first_point.y;
        let mut y_max = first_point.y;
        for &point in &points[1..] {
            x_min = x_min.min(point.x);
            y_min = y_min.min(point.y);
            x_max = x_max.max(point.x);
            y_max = y_max.max(point.y);
        }
        BoundingBox {
            x_min,
            y_min,
            x_max,
            y_max,
        }
    }

    fn height(&self) -> u64 {
        1 + (self.y_min - self.y_max).abs() as u64
    }

    fn width(&self) -> u64 {
        1 + (self.x_min - self.x_max).abs() as u64
    }

    fn points(&self) -> collections::HashSet<grid::Point> {
        let mut points =
            collections::HashSet::with_capacity((self.width() * self.height()) as usize);
        for x in self.x_min..=self.x_max {
            for y in self.y_min..=self.y_max {
                points.insert(grid::Point { x, y });
            }
        }
        points
    }

    fn edge_points(&self) -> collections::HashSet<grid::Point> {
        let mut points = Vec::with_capacity(2 * (self.width() + self.height()) as usize - 4);
        for x in self.x_min..self.x_max {
            points.push(grid::Point { x, y: self.y_min });
            points.push(grid::Point { x, y: self.y_max });
        }
        for y in self.y_min..self.y_max {
            points.push(grid::Point { x: self.x_min, y });
            points.push(grid::Point { x: self.x_max, y });
        }
        let mut set = collections::HashSet::with_capacity(points.len());
        set.extend(points);
        set
    }
}

fn distances(point: &grid::Point, coordinates: &Coordinates) -> Distances {
    coordinates
        .par_iter()
        .map(|(&c, &coord_point)| (c, point.manhattan_distance_to(coord_point)))
        .collect()
}

fn bounding_box_distances(
    bb: &BoundingBox,
    coordinates: &Coordinates,
) -> collections::HashMap<grid::Point, Distances> {
    bb.points()
        .par_iter()
        .map(|&point| (point, distances(&point, coordinates)))
        .collect()
}

fn minimal_distances(distances: &Distances) -> Vec<char> {
    let minimal_distance = *distances.values().min().unwrap();
    distances
        .keys()
        .cloned()
        .filter(|k| distances[k] == minimal_distance)
        .collect()
}

fn all_minimal_distances(
    map: &collections::HashMap<grid::Point, Distances>,
) -> collections::HashMap<grid::Point, Vec<char>> {
    map.par_iter()
        .map(|(&point, distances)| (point, minimal_distances(distances)))
        .collect()
}

#[cfg(test)]
mod day06_test;
