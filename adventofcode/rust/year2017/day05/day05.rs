use adventofcode_rust_aoc as aoc;

pub fn part1(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::One)
}

pub fn part2(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::Two)
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let instructions = parse_input(&input);
    let next_value = match part {
        aoc::Part::One => increase_by_one,
        aoc::Part::Two => decrement_if_three_or_more,
    };
    Ok(steps_until_escape(&mut instructions.clone(), next_value).to_string())
}

fn parse_input(input: &str) -> Vec<i64> {
    input.lines().map(str::parse).map(Result::unwrap).collect()
}

fn increase_by_one(i: i64) -> i64 {
    i + 1
}

fn decrement_if_three_or_more(i: i64) -> i64 {
    if i >= 3 {
        i - 1
    } else {
        i + 1
    }
}

fn steps_until_escape(instructions: &mut [i64], next_value: fn(i64) -> i64) -> u64 {
    let mut idx = 0;
    let mut counter = 0;
    while idx < instructions.len() {
        let offset = instructions[idx];
        instructions[idx] = next_value(offset);
        if offset < 0 {
            idx -= offset.abs() as usize;
        } else {
            idx += offset as usize;
        }
        counter += 1;
    }
    counter
}

#[cfg(test)]
mod day05_test;
