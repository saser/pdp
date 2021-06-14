use adventofcode_rust_aoc as aoc;

pub fn part1(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::One)
}

pub fn part2(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::Two)
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let tokens = parse_tokens(input.trim())?;
    let (score, removed_garbage) = process_tokens(&tokens);
    match part {
        aoc::Part::One => Ok(score.to_string()),
        aoc::Part::Two => Ok(removed_garbage.to_string()),
    }
}

#[derive(Clone, Copy, Debug, Eq, Hash, PartialEq)]
enum Token {
    StartGroup,
    EndGroup,
    StartGarbage,
    EndGarbage,
    Garbage(char),
    Ignore(char),
    Separator,
}

fn parse_tokens(s: &str) -> Result<Vec<Token>, String> {
    let mut chars = s.chars();
    let mut tokens = Vec::with_capacity(s.len());
    let mut is_garbage = false;

    while let Some(c) = chars.next() {
        let token = if is_garbage {
            match c {
                '>' => Token::EndGarbage,
                '!' => Token::Ignore(chars.next().unwrap()),
                garbage => Token::Garbage(garbage),
            }
        } else {
            match c {
                '{' => Token::StartGroup,
                '}' => Token::EndGroup,
                '<' => Token::StartGarbage,
                ',' => Token::Separator,
                invalid => {
                    return Err(format!(
                        "could not parse to token: invalid character '{}'",
                        invalid
                    ));
                }
            }
        };

        if token == Token::StartGarbage {
            is_garbage = true;
        } else if token == Token::EndGarbage {
            is_garbage = false;
        }

        tokens.push(token);
    }

    Ok(tokens)
}

fn process_tokens(tokens: &[Token]) -> (u64, u64) {
    let mut score = 0;
    let mut current_group = 0;
    let mut removed_garbage = 0;
    for &token in tokens {
        match token {
            Token::StartGroup => {
                current_group += 1;
                score += current_group;
            }
            Token::EndGroup => current_group -= 1,
            Token::Garbage(_) => removed_garbage += 1,
            _ => {}
        };
    }
    (score, removed_garbage)
}

#[cfg(test)]
mod day09_test;
