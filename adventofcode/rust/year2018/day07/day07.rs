use std::cmp;
use std::collections;

use adventofcode_rust_aoc as aoc;
use lazy_static;
use regex;

type Dependencies = collections::HashMap<char, collections::HashSet<char>>;
type Dependants = collections::HashMap<char, collections::HashSet<char>>;

#[derive(Debug, Clone, Copy, Hash, Eq, PartialEq)]
struct RevChar(char);

impl PartialOrd for RevChar {
    fn partial_cmp(&self, other: &RevChar) -> Option<cmp::Ordering> {
        Some(self.cmp(other))
    }
}

impl Ord for RevChar {
    fn cmp(&self, other: &RevChar) -> cmp::Ordering {
        match self.0.cmp(&other.0) {
            cmp::Ordering::Less => cmp::Ordering::Greater,
            cmp::Ordering::Equal => cmp::Ordering::Equal,
            cmp::Ordering::Greater => cmp::Ordering::Less,
        }
    }
}

pub fn part1(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::One)
}

pub fn part2(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::Two)
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let (dependencies, dependants) = parse_input(&input);
    let workers = 5;
    match part {
        aoc::Part::One => Ok(determine_order(&dependencies, &dependants)),
        aoc::Part::Two => Ok(seconds_with_workers(workers, &dependencies, &dependants).to_string()),
    }
}

fn parse_input(input: &str) -> (Dependencies, Dependants) {
    let mut dependencies = collections::HashMap::new();
    let mut dependants = collections::HashMap::new();
    input
        .lines()
        .map(parse_instruction)
        .for_each(|(dependency, dependant)| {
            dependencies
                .entry(dependant)
                .or_insert_with(collections::HashSet::new)
                .insert(dependency);
            dependants
                .entry(dependency)
                .or_insert_with(collections::HashSet::new)
                .insert(dependant);
        });
    (dependencies, dependants)
}

fn parse_instruction(instruction: &str) -> (char, char) {
    lazy_static::lazy_static! {
        static ref INSTR_RE: regex::Regex = regex::Regex::new(
            r"Step (?P<dependency>\w) must be finished before step (?P<dependant>\w) can begin."
        )
        .unwrap();
    }
    let captures = INSTR_RE.captures(instruction).unwrap();
    let dependency = captures["dependency"].chars().collect::<Vec<char>>()[0];
    let dependant = captures["dependant"].chars().collect::<Vec<char>>()[0];
    (dependency, dependant)
}

fn determine_order(dependencies: &Dependencies, dependants: &Dependants) -> String {
    let mut order = String::new();
    let mut done = collections::HashSet::new();
    let mut available = (&dependants
        .keys()
        .cloned()
        .collect::<collections::HashSet<char>>()
        - &dependencies
            .keys()
            .cloned()
            .collect::<collections::HashSet<char>>())
        .iter()
        .cloned()
        .map(RevChar)
        .collect::<collections::BinaryHeap<RevChar>>();
    while !available.is_empty() {
        println!("available: {:?}", available);
        println!("done: {:?}", done);
        let next_step = available.pop().unwrap();
        if done.contains(&next_step.0) {
            continue;
        }
        order.push(next_step.0);
        done.insert(next_step.0);
        for (&c, local_dependencies) in dependencies.iter() {
            if local_dependencies.is_subset(&done) && !done.contains(&c) {
                available.push(RevChar(c));
            }
        }
    }
    order
}

fn seconds_with_workers(workers: u64, dependencies: &Dependencies, dependants: &Dependants) -> u64 {
    let dependencies_keys = dependencies
        .keys()
        .cloned()
        .collect::<collections::HashSet<char>>();
    let dependants_keys = dependants
        .keys()
        .cloned()
        .collect::<collections::HashSet<char>>();
    let all_steps = &dependencies_keys | &dependants_keys;

    let mut done = collections::HashSet::new();
    let mut current_time = 0;
    let initially_available = (&dependants_keys - &dependencies_keys)
        .iter()
        .cloned()
        .collect::<Vec<char>>();
    let mut events = initially_available
        .iter()
        .map(|&c| (duration(c), c))
        .collect::<Vec<(u64, char)>>();
    events.sort();
    let mut available = Vec::new();
    let mut available_workers = workers - events.len() as u64;

    while done.len() < all_steps.len() {
        let (new_time, completed_step) = events.remove(0);
        current_time = new_time;
        done.insert(completed_step);
        available_workers += 1;

        if dependants.contains_key(&completed_step) {
            let newly_available = dependants[&completed_step]
                .iter()
                .cloned()
                .filter(|&dependant| {
                    is_available(dependant, dependencies, &done) && !available.contains(&dependant)
                })
                .collect::<Vec<char>>();
            available.extend(newly_available);

            available.sort();
        }

        while available.len() > 0 && available_workers > 0 {
            let next_step = available.remove(0);
            available_workers -= 1;
            events.push((current_time + duration(next_step), next_step));
        }
        events.sort();
    }
    current_time
}

fn duration(c: char) -> u64 {
    (c as u64) - 4
}

fn is_available(c: char, dependencies: &Dependencies, done: &collections::HashSet<char>) -> bool {
    !dependencies.contains_key(&c) || !done.contains(&c) && dependencies[&c].is_subset(done)
}

#[cfg(test)]
mod day07_test;
