use std::collections;
use std::io;
use std::io::BufRead;
use std::str::FromStr;

use adventofcode_rust_aoc as aoc;

pub fn part1(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::One)
}

pub fn part2(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::Two)
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let instructions = io::BufReader::new(input.as_bytes())
        .lines()
        .map(|line| line.expect("input line read failed"))
        .map(|ref line| Instruction::from_str(line))
        .map(|i| i.expect("instruction parsing failed"))
        .collect::<Vec<Instruction>>();
    let mut vm0 = VM::new(&instructions);
    let mut outs0 = vm0.run();
    if part == aoc::Part::One {
        return Ok(outs0.back().unwrap().to_string());
    }
    let mut vm0_sent = outs0.len();
    let mut vm1 = VM::new(&instructions);
    *vm1.register('p') = 1;
    vm1.inputs.append(&mut outs0);
    let mut vm1_sent = 0;
    let mut vm0_running = false;
    loop {
        let (current, sent, next) = match vm0_running {
            true => (&mut vm0, &mut vm0_sent, &mut vm1),
            false => (&mut vm1, &mut vm1_sent, &mut vm0),
        };
        let mut outs = current.run();
        if outs.len() == 0 {
            break;
        }
        *sent += outs.len();
        next.inputs.append(&mut outs);
        vm0_running = !vm0_running;
    }
    Ok(vm1_sent.to_string())
}

struct VM {
    instructions: Vec<Instruction>,
    registers: collections::HashMap<char, i64>,
    inputs: collections::VecDeque<i64>,
    pc: isize,
}

impl VM {
    fn new(instructions: &[Instruction]) -> Self {
        VM {
            instructions: instructions.to_vec(),
            registers: collections::HashMap::new(),
            inputs: collections::VecDeque::new(),
            pc: 0,
        }
    }

    fn run(&mut self) -> collections::VecDeque<i64> {
        let mut outs = collections::VecDeque::new();
        while let Some(opt_out) = self.step() {
            if let Some(i) = opt_out {
                outs.push_back(i);
            }
        }
        outs
    }

    fn step(&mut self) -> Option<Option<i64>> {
        if self.pc as usize >= self.instructions.len() {
            return None;
        }
        let mut new_pc = self.pc + 1;
        let mut out = None;
        match self.instructions[self.pc as usize] {
            Instruction::Snd(op) => out = Some(self.eval_op(op)),
            Instruction::Set(op1, op2) => *self.must_register(op1) = self.eval_op(op2),
            Instruction::Add(op1, op2) => *self.must_register(op1) += self.eval_op(op2),
            Instruction::Mul(op1, op2) => *self.must_register(op1) *= self.eval_op(op2),
            Instruction::Mod(op1, op2) => *self.must_register(op1) %= self.eval_op(op2),
            Instruction::Rcv(op) => {
                match self.inputs.pop_front() {
                    Some(i) => *self.must_register(op) = i,
                    None => return None,
                };
            }
            Instruction::Jgz(op1, op2) => {
                if self.eval_op(op1) > 0 {
                    new_pc = self.pc + self.eval_op(op2) as isize;
                }
            }
        };
        self.pc = new_pc;
        Some(out)
    }

    fn eval_op(&mut self, op: Operand) -> i64 {
        match op {
            Operand::Integer(i) => i,
            Operand::Register(c) => *self.register(c),
        }
    }

    fn register(&mut self, c: char) -> &mut i64 {
        self.registers.entry(c).or_insert(0)
    }

    fn must_register(&mut self, op: Operand) -> &mut i64 {
        if let Operand::Register(c) = op {
            self.register(c)
        } else {
            panic!("non-register operand {:?}", op)
        }
    }
}

#[derive(Clone, Debug)]
enum Instruction {
    Snd(Operand),
    Set(Operand, Operand),
    Add(Operand, Operand),
    Mul(Operand, Operand),
    Mod(Operand, Operand),
    Rcv(Operand),
    Jgz(Operand, Operand),
}

impl FromStr for Instruction {
    type Err = String;
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let parts = s.split(" ").collect::<Vec<&str>>();
        if parts.len() == 0 {
            return Err("empty instruction".to_string());
        }
        if parts.len() == 1 {
            return Err(format!("missing operands: {}", s));
        }
        let op = parts[0];
        let operand1 = Operand::from_str(parts[1])?;
        match op {
            "snd" => Ok(Instruction::Snd(operand1)),
            "rcv" => Ok(Instruction::Rcv(operand1)),
            "set" | "add" | "mul" | "mod" | "jgz" if parts.len() == 3 => {
                let operand2 = Operand::from_str(parts[2])?;
                Ok(match op {
                    "set" => Instruction::Set(operand1, operand2),
                    "add" => Instruction::Add(operand1, operand2),
                    "mul" => Instruction::Mul(operand1, operand2),
                    "mod" => Instruction::Mod(operand1, operand2),
                    "jgz" => Instruction::Jgz(operand1, operand2),
                    _ => unreachable!(),
                })
            }
            _ => Err(format!("invalid op: {}", op)),
        }
    }
}

#[derive(Clone, Copy, Debug)]
enum Operand {
    Integer(i64),
    Register(char),
}

impl FromStr for Operand {
    type Err = String;
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        if let Ok(i) = i64::from_str(s) {
            return Ok(Operand::Integer(i));
        }
        if let Some(c) = s.chars().next() {
            return Ok(Operand::Register(c));
        }
        Err(format!("invalid operand: {}", s))
    }
}

#[cfg(test)]
mod day18_test;
