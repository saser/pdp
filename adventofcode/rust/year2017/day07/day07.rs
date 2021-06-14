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
    let programs = parse_input(&input);
    let tower = construct_tower(&programs);
    let bottom_program = find_bottom_program(&tower);
    if part == aoc::Part::One {
        Ok(bottom_program.name.clone())
    } else {
        let tower_weights = calculate_tower_weights(&tower, &bottom_program);
        Ok(find_correct_weight(&tower, &tower_weights, &bottom_program)
            .unwrap()
            .to_string())
    }
}

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
struct Program {
    name: String,
    weight: u64,
    holding_up: Option<Vec<String>>,
    held_up_by: Option<String>,
}

impl FromStr for Program {
    type Err = String;
    fn from_str(s: &str) -> Result<Program, String> {
        lazy_static::lazy_static! {
            static ref NAME_AND_WEIGHT: regex::Regex =
                regex::Regex::new(r"(?P<name>\w+) \((?P<weight>\d+)\)").unwrap();
        }
        let parts: Vec<&str> = s.split(" -> ").collect();
        let (name_and_weight, programs) = (
            parts[0],
            if parts.len() == 2 {
                Some(parts[1])
            } else {
                None
            },
        );
        let name_and_weight_caps = NAME_AND_WEIGHT.captures(name_and_weight).unwrap();
        let name = name_and_weight_caps["name"].to_string();
        let weight: u64 = name_and_weight_caps["weight"].parse().unwrap();
        let holding_up =
            programs.map(|program_str| program_str.split(", ").map(String::from).collect());

        Ok(Program {
            name,
            weight,
            holding_up,
            held_up_by: None,
        })
    }
}

fn parse_input(input: &str) -> collections::HashMap<String, Program> {
    input
        .lines()
        .map(Program::from_str)
        .map(Result::unwrap)
        .map(|prog| (prog.name.clone(), prog))
        .collect()
}

fn construct_tower(
    programs: &collections::HashMap<String, Program>,
) -> collections::HashMap<String, Program> {
    let mut tower = programs.clone();
    let progs_holding_up: Vec<Program> = tower
        .values()
        .cloned()
        .filter(|prog| prog.holding_up.is_some())
        .collect();
    for holding_prog in progs_holding_up {
        for prog in holding_prog.holding_up.unwrap() {
            tower.get_mut(&prog).unwrap().held_up_by = Some(holding_prog.name.clone());
        }
    }
    tower
}

fn find_bottom_program(tower: &collections::HashMap<String, Program>) -> Program {
    tower
        .values()
        .find(|prog| prog.held_up_by.is_none())
        .unwrap()
        .clone()
}

fn calculate_tower_weights(
    tower: &collections::HashMap<String, Program>,
    root: &Program,
) -> collections::HashMap<String, u64> {
    if root.holding_up.is_none() {
        let mut map = collections::HashMap::new();
        map.insert(root.name.clone(), root.weight);
        return map;
    }

    let held_up_progs = root.holding_up.clone().unwrap();
    let mut map = collections::HashMap::new();
    let mut weight = 0;
    for prog in &held_up_progs {
        map.extend(calculate_tower_weights(tower, tower.get(prog).unwrap()));
        weight += map[prog];
    }
    weight += root.weight;
    map.insert(root.name.clone(), weight);
    map
}

#[allow(clippy::question_mark)]
fn find_correct_weight(
    tower: &collections::HashMap<String, Program>,
    tower_weights: &collections::HashMap<String, u64>,
    root: &Program,
) -> Option<u64> {
    if root.holding_up.is_none() {
        return None;
    }

    let mut map: collections::HashMap<u64, Vec<String>> = collections::HashMap::new();
    for held_up_prog in &root.holding_up.clone().unwrap() {
        let correct_weight =
            find_correct_weight(tower, tower_weights, tower.get(held_up_prog).unwrap());
        if correct_weight.is_some() {
            return correct_weight;
        }
        map.entry(*tower_weights.get(held_up_prog).unwrap())
            .or_insert_with(Vec::new)
            .push(held_up_prog.clone());
    }

    if map.len() == 1 {
        None
    } else {
        let offending_prog_name = &map.iter().find(|&(_, progs)| progs.len() == 1).unwrap().1[0];
        let offending_prog_weight = tower.get(offending_prog_name).unwrap().weight;
        let offending_subtower_weight = tower_weights.get(offending_prog_name).unwrap();
        let desired_subtower_weight = map.iter().find(|&(_, progs)| progs.len() > 1).unwrap().0;
        if offending_subtower_weight > desired_subtower_weight {
            Some(offending_prog_weight - (offending_subtower_weight - desired_subtower_weight))
        } else {
            Some(offending_prog_weight + (desired_subtower_weight - offending_subtower_weight))
        }
    }
}

#[cfg(test)]
mod day07_test;
