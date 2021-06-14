use std::cmp;
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
    let instructions = parse_input(&input);
    let registers = initialize_registers(&instructions);
    let (final_registers, highest_value) = perform_instructions(&registers, &instructions);
    let max_register = final_registers.values().max().unwrap();
    match part {
        aoc::Part::One => Ok(max_register.to_string()),
        aoc::Part::Two => Ok(highest_value.to_string()),
    }
}

fn parse_input(input: &str) -> Vec<Instruction> {
    input
        .lines()
        .map(Instruction::from_str)
        .map(Result::unwrap)
        .collect()
}

fn initialize_registers(instructions: &[Instruction]) -> collections::HashMap<String, i64> {
    let mut map = collections::HashMap::new();
    for instruction in instructions {
        map.entry(instruction.register.clone()).or_insert(0);
    }
    map
}

fn do_comparison(comparison: Comparison, a: i64, b: i64) -> bool {
    match comparison {
        Comparison::Lt => a < b,
        Comparison::Gt => a > b,
        Comparison::EqLt => a <= b,
        Comparison::EqGt => a >= b,
        Comparison::Eq => a == b,
        Comparison::Neq => a != b,
    }
}

fn check_condition(registers: &collections::HashMap<String, i64>, condition: &Condition) -> bool {
    let register_value = *registers.get(&condition.register).unwrap();
    do_comparison(condition.cmp, register_value, condition.value)
}

fn perform_operation(operation: Operation, previous_value: i64) -> i64 {
    match operation {
        Operation::Inc(delta) => previous_value + delta,
        Operation::Dec(delta) => previous_value - delta,
    }
}

fn perform_instruction(
    (registers, highest_value): (&collections::HashMap<String, i64>, i64),
    instruction: &Instruction,
) -> (collections::HashMap<String, i64>, i64) {
    let mut map = registers.clone();
    let mut new_highest_value = highest_value;
    if check_condition(registers, &instruction.cond) {
        let register_value = map[&instruction.register];
        let new_value = perform_operation(instruction.op, register_value);
        new_highest_value = cmp::max(highest_value, new_value);
        map.insert(
            instruction.register.clone(),
            perform_operation(instruction.op, register_value),
        );
    }
    (map, new_highest_value)
}

fn perform_instructions(
    registers: &collections::HashMap<String, i64>,
    instructions: &[Instruction],
) -> (collections::HashMap<String, i64>, i64) {
    instructions
        .iter()
        .fold((registers.clone(), 0), |(regs, highest), instruction| {
            perform_instruction((&regs, highest), instruction)
        })
}

#[derive(Clone, Copy, Debug, Eq, Hash, PartialEq)]
enum Operation {
    Inc(i64),
    Dec(i64),
}

impl FromStr for Operation {
    type Err = String;

    fn from_str(s: &str) -> Result<Operation, String> {
        lazy_static::lazy_static! {
            static ref RE: regex::Regex = regex::Regex::new(r"(?P<op>inc|dec) (?P<value>-?\d+)").unwrap();
        }
        let captures = RE.captures(s).unwrap();
        let value = i64::from_str(&captures["value"]).unwrap();
        match &captures["op"] {
            "inc" => Ok(Operation::Inc(value)),
            "dec" => Ok(Operation::Dec(value)),
            _ => Err(format!("invalid operation: {}", s)),
        }
    }
}

#[derive(Clone, Copy, Debug, Eq, Hash, PartialEq)]
enum Comparison {
    Lt,
    Gt,
    EqLt,
    EqGt,
    Eq,
    Neq,
}

impl FromStr for Comparison {
    type Err = String;

    fn from_str(s: &str) -> Result<Comparison, String> {
        match s {
            "<" => Ok(Comparison::Lt),
            ">" => Ok(Comparison::Gt),
            "<=" => Ok(Comparison::EqLt),
            ">=" => Ok(Comparison::EqGt),
            "==" => Ok(Comparison::Eq),
            "!=" => Ok(Comparison::Neq),
            _ => Err(format!("invalid comparison: {}", s)),
        }
    }
}

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
struct Condition {
    register: String,
    cmp: Comparison,
    value: i64,
}

impl FromStr for Condition {
    type Err = String;

    fn from_str(s: &str) -> Result<Condition, String> {
        lazy_static::lazy_static! {
            static ref RE: regex::Regex =
                regex::Regex::new(r"(?P<register>\w+) (?P<cmp><|>|<=|>=|==|!=) (?P<value>-?\d+)").unwrap();
        }
        let captures = RE.captures(s).unwrap();
        let register = captures["register"].to_string();
        let cmp = Comparison::from_str(&captures["cmp"]).unwrap();
        let value = i64::from_str(&captures["value"]).unwrap();
        Ok(Condition {
            register,
            cmp,
            value,
        })
    }
}

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
struct Instruction {
    register: String,
    op: Operation,
    cond: Condition,
}

impl FromStr for Instruction {
    type Err = String;

    fn from_str(s: &str) -> Result<Instruction, String> {
        let parts: Vec<&str> = s.split(" if ").collect();
        let before_condition: Vec<&str> = parts[0].splitn(2, ' ').collect();
        let register = before_condition[0].to_string();
        let op = Operation::from_str(before_condition[1])?;
        let cond = Condition::from_str(parts[1])?;
        Ok(Instruction { register, op, cond })
    }
}

#[cfg(test)]
mod day08_test;
