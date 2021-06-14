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
    let connections = parse_input(&input);
    match part {
        aoc::Part::One => Ok(find_group(&connections, 0).len().to_string()),
        aoc::Part::Two => Ok(find_all_groups(&connections).len().to_string()),
    }
}

fn parse_input(input: &str) -> collections::HashMap<u64, Vec<u64>> {
    input.lines().map(parse_line).collect()
}

fn parse_line(line: &str) -> (u64, Vec<u64>) {
    let parts: Vec<&str> = line.split(" <-> ").collect();
    let program = u64::from_str(parts[0]).unwrap();
    let connected: Vec<u64> = parts[1]
        .split(", ")
        .map(u64::from_str)
        .map(Result::unwrap)
        .collect();
    (program, connected)
}

fn find_all_groups(connections: &collections::HashMap<u64, Vec<u64>>) -> Vec<Vec<u64>> {
    let number_of_programs = connections.len();
    let mut groups: Vec<Vec<u64>> = Vec::new();
    let mut visited = vec![false; number_of_programs];
    let mut queue: collections::VecDeque<u64> = (0..number_of_programs).map(|n| n as u64).collect();
    while let Some(unvisited_program) = queue.pop_front() {
        if visited[unvisited_program as usize] {
            continue;
        }

        let group = find_group(connections, unvisited_program as u64);
        for &program in &group {
            visited[program as usize] = true;
        }
        groups.push(group);
    }
    groups
}

fn find_group(
    connections: &collections::HashMap<u64, Vec<u64>>,
    included_program: u64,
) -> Vec<u64> {
    let mut queue: collections::VecDeque<u64> = collections::VecDeque::new();
    let mut group: Vec<u64> = Vec::new();
    let mut visited = vec![false; connections.len()];
    queue.push_back(included_program);

    while let Some(program) = queue.pop_front() {
        if visited[program as usize] {
            continue;
        }

        group.push(program);
        visited[program as usize] = true;
        let connected_programs = connections.get(&program).unwrap();
        for &connected_program in connected_programs {
            queue.push_back(connected_program);
        }
    }
    group
}

#[cfg(test)]
mod day12_test;
