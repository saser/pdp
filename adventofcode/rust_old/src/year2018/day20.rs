use std::collections::{HashMap, HashSet, VecDeque};
use std::io;

use crate::base::Part;

pub fn part1(r: &mut dyn io::Read) -> Result<String, String> {
    solve(r, Part::One)
}

pub fn part2(r: &mut dyn io::Read) -> Result<String, String> {
    solve(r, Part::Two)
}

fn solve(r: &mut dyn io::Read, part: Part) -> Result<String, String> {
    let mut input = String::new();
    r.read_to_string(&mut input).map_err(|e| e.to_string())?;
    let regex = parse(input.trim());
    let graph = construct(&regex);
    match part {
        Part::One => Ok(furthest(&graph).to_string()),
        Part::Two => {
            let distances = distances(&graph);
            let count = distances
                .values()
                .filter(|&&distance| distance >= 1000)
                .count();
            Ok(count.to_string())
        }
    }
}

type Graph = HashMap<Position, HashSet<Position>>;

#[derive(Clone, Copy, Debug, Eq, Hash, Ord, PartialEq, PartialOrd)]
struct Position {
    x: i64,
    y: i64,
}

impl Position {
    fn origin() -> Self {
        Position { x: 0, y: 0 }
    }

    fn north(&self) -> Self {
        Position {
            x: self.x,
            y: self.y + 1,
        }
    }

    fn south(&self) -> Self {
        Position {
            x: self.x,
            y: self.y - 1,
        }
    }

    fn east(&self) -> Self {
        Position {
            x: self.x + 1,
            y: self.y,
        }
    }

    fn west(&self) -> Self {
        Position {
            x: self.x - 1,
            y: self.y,
        }
    }
}

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
struct Regex {
    tokens: Vec<Token>,
}

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
enum Token {
    Terminals(Vec<char>),
    Branch(Vec<Regex>),
}

fn parse(input: &str) -> Regex {
    let regex_chars = input.chars().collect::<Vec<char>>();
    let (regex, _offset) = parse_regex(&regex_chars);
    regex
}

fn parse_regex(regex_chars: &[char]) -> (Regex, usize) {
    let mut tokens = Vec::new();
    let mut i = 0;
    while i < regex_chars.len() {
        match regex_chars[i] {
            '^' | '$' => i += 1,
            '|' | ')' => break,
            '(' => {
                let (regexes, offset) = parse_branch(&regex_chars[i..]);
                tokens.push(Token::Branch(regexes));
                i += offset;
            }
            t if is_terminal(t) => {
                let (terminals, offset) = parse_terminals(&regex_chars[i..]);
                tokens.push(Token::Terminals(terminals));
                i += offset;
            }
            c => panic!("parse_regex: unexpected char: {}", c),
        };
    }
    (Regex { tokens }, i)
}

fn parse_terminals(regex_chars: &[char]) -> (Vec<char>, usize) {
    let terminals = regex_chars
        .iter()
        .cloned()
        .take_while(|&c| is_terminal(c))
        .collect::<Vec<char>>();
    let offset = terminals.len();
    (terminals, offset)
}

fn parse_branch(regex_chars: &[char]) -> (Vec<Regex>, usize) {
    let mut regexes = Vec::new();
    let mut i = 0;
    while i < regex_chars.len() {
        match regex_chars[i] {
            '(' | '|' => {
                i += 1;
                let (regex, offset) = parse_regex(&regex_chars[i..]);
                regexes.push(regex);
                i += offset;
            }
            ')' => {
                i += 1;
                break;
            }
            c => panic!("parse_branch: unexpected char: {}", c),
        };
    }
    (regexes, i)
}

fn is_terminal(c: char) -> bool {
    ['N', 'E', 'S', 'W'].contains(&c)
}

fn construct(regex: &Regex) -> Graph {
    let mut graph = Graph::new();
    let mut positions = HashSet::new();
    positions.insert(Position::origin());
    construct_regex(regex, &positions, &mut graph);
    graph
}

fn construct_regex(
    regex: &Regex,
    positions: &HashSet<Position>,
    graph: &mut Graph,
) -> HashSet<Position> {
    let mut new_positions = positions.clone();
    for token in &regex.tokens {
        match token {
            Token::Terminals(ref terminals) => {
                new_positions = construct_terminals(terminals, &new_positions, graph);
            }
            Token::Branch(ref regexes) => {
                new_positions = construct_branch(regexes, &new_positions, graph);
            }
        }
    }
    new_positions
}

fn construct_terminals(
    terminals: &[char],
    positions: &HashSet<Position>,
    graph: &mut Graph,
) -> HashSet<Position> {
    let mut new_positions = HashSet::new();
    for &position in positions {
        let mut current_position = position;
        for t in terminals {
            let next_position = match t {
                'N' => current_position.north(),
                'E' => current_position.east(),
                'S' => current_position.south(),
                'W' => current_position.west(),
                _ => unreachable!(),
            };
            graph
                .entry(current_position)
                .or_insert_with(HashSet::new)
                .insert(next_position);
            graph
                .entry(next_position)
                .or_insert_with(HashSet::new)
                .insert(current_position);
            current_position = next_position;
        }
        new_positions.insert(current_position);
    }
    new_positions
}

fn construct_branch(
    regexes: &[Regex],
    positions: &HashSet<Position>,
    graph: &mut Graph,
) -> HashSet<Position> {
    let mut new_positions = HashSet::new();
    for regex in regexes {
        new_positions.extend(construct_regex(regex, positions, graph));
    }
    new_positions
}

fn distances(graph: &Graph) -> HashMap<Position, u64> {
    let mut distances = HashMap::new();
    let mut furthest_distance = 0;
    let mut queue = VecDeque::new();
    queue.push_back((Position::origin(), 0));
    while let Some((position, distance)) = queue.pop_front() {
        if distances.contains_key(&position) {
            continue;
        }
        distances.insert(position, distance);
        furthest_distance = furthest_distance.max(distance);
        if let Some(neighbors) = graph.get(&position) {
            for &neighbor in neighbors {
                queue.push_back((neighbor, distance + 1));
            }
        }
    }
    distances
}

fn furthest(graph: &Graph) -> u64 {
    *distances(graph).values().max().unwrap()
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::test;

    mod part1 {
        use super::*;

        test!(example1, "^WNE$", "3", part1);
        test!(example2, "^ENWWW(NEEE|SSE(EE|N))$", "10", part1);
        test!(
            example3,
            "^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$",
            "18",
            part1
        );
        test!(
            example4,
            "^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$",
            "23",
            part1
        );
        test!(
            example5,
            "^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$",
            "31",
            part1
        );
        test!(actual, file "../../../inputs/2018/20", "4214", part1);
    }

    mod part2 {
        use super::*;

        test!(actual, file "../../../inputs/2018/20", "8492", part2);
    }
}
