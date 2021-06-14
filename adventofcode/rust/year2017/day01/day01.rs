use adventofcode_rust_aoc as aoc;

pub fn part1(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::One)
}

pub fn part2(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::Two)
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let digits = parse_input(input.trim());
    let offset = match part {
        aoc::Part::One => 1,
        aoc::Part::Two => digits.len() / 2,
    };
    Ok(sum_matching(&digits, offset).to_string())
}

fn parse_input(input: &str) -> Vec<u32> {
    input.chars().map(|c| c.to_digit(10).unwrap()).collect()
}

fn sum_matching(digits: &[u32], offset: usize) -> u32 {
    let n = digits.len();
    let mut sum = 0;
    for (idx, &d) in digits.iter().enumerate() {
        let u = d;
        let v_idx = (idx + offset) % n;
        let v = digits[v_idx];
        if u == v {
            sum += u;
        }
    }
    sum
}


#[cfg(test)]
mod day01_test;
