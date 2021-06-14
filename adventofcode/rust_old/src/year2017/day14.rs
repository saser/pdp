use std::collections::{HashSet, VecDeque};
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
    let strings = strings_to_hash(input.trim());
    match part {
        Part::One => Ok(total_bits(&hash_all(&strings)).to_string()),
        Part::Two => {
            //let hashes = hash_all(&strings);
            Ok(count_groups(&hashes_to_binary(&hash_all(&strings))).to_string())
        }
    }
}

fn strings_to_hash(input: &str) -> Vec<String> {
    (0..128).map(|i| format!("{}-{}", input, i)).collect()
}

fn hex_digit_to_binary(digit: char) -> String {
    format!("{:04b}", digit.to_digit(16).unwrap())
}

fn hash_to_binary(hash: &str) -> String {
    hash.chars()
        .map(hex_digit_to_binary)
        .collect::<Vec<String>>()
        .join("")
}

fn binary_to_vec(binary: &str) -> Vec<bool> {
    binary
        .chars()
        .map(|c| match c {
            '1' => true,
            '0' => false,
            _ => panic!("invalid digit in binary string: {}", c),
        })
        .collect()
}

fn bits_in_hash(hash: &str) -> usize {
    hash_to_binary(hash).chars().filter(|&c| c == '1').count()
}

fn hash_all(strings: &[String]) -> Vec<String> {
    strings
        .iter()
        .map(|string| day10::full_hash_str(&string))
        .collect()
}

fn total_bits(hashes: &[String]) -> usize {
    hashes.iter().map(|hash| bits_in_hash(&hash)).sum()
}

fn hashes_to_binary(hashes: &[String]) -> HashSet<(usize, usize)> {
    let mut set = HashSet::new();
    let binary_vectors: Vec<Vec<bool>> = hashes
        .iter()
        .map(|hash| hash_to_binary(&hash))
        .map(|binstring| binary_to_vec(&binstring))
        .collect();
    for (i, binvec) in binary_vectors.as_slice().iter().enumerate() {
        for (j, &binval) in binvec.as_slice().iter().enumerate() {
            if binval {
                set.insert((i, j));
            }
        }
    }
    set
}

fn mark_group(
    visited: &mut HashSet<(usize, usize)>,
    set: &HashSet<(usize, usize)>,
    (start_x, start_y): (usize, usize),
) {
    let mut queue = VecDeque::new();
    queue.push_back((start_x, start_y));

    while let Some((x, y)) = queue.pop_front() {
        if visited.contains(&(x, y)) {
            continue;
        }
        visited.insert((x, y));

        let mut adjacent = vec![(x + 1, y), (x, y + 1)];
        if x > 0 {
            adjacent.push((x - 1, y));
        }

        if y > 0 {
            adjacent.push((x, y - 1));
        }

        for &pos in &adjacent {
            if set.contains(&pos) {
                queue.push_back(pos);
            }
        }
    }
}

fn count_groups(set: &HashSet<(usize, usize)>) -> u64 {
    let mut visited = HashSet::with_capacity(set.len());
    let mut counter = 0;
    for &pos in set.iter() {
        if visited.contains(&pos) {
            continue;
        }
        mark_group(&mut visited, set, pos);
        counter += 1;
    }
    counter
}

mod day10 {
    use std::str::FromStr;

    fn parse_input_as_lengths(input: &str) -> Vec<u8> {
        input
            .split(',')
            .map(u8::from_str)
            .map(Result::unwrap)
            .collect()
    }

    fn parse_input_as_bytes(input: &str) -> Vec<u8> {
        Vec::from(input.as_bytes())
    }

    fn initialize_vector() -> Vec<u8> {
        let mut i = 0;
        let highest_value = std::u8::MAX;
        let mut v = Vec::with_capacity(highest_value as usize + 1);
        while i < highest_value {
            v.push(i);
            i += 1;
        }
        v.push(highest_value);
        v
    }

    fn indices_wrapping(slice_length: usize, start: usize, length: usize) -> Vec<usize> {
        (start..start + length).map(|i| i % slice_length).collect()
    }

    fn reverse_by_indices<T: Copy>(slice: &mut [T], indices: &[usize]) {
        if indices.is_empty() {
            return;
        }

        let mut i = 0;
        let mut j = indices.len() - 1;
        while i < j {
            let early = slice[indices[i]];
            let late = slice[indices[j]];
            slice[indices[i]] = late;
            slice[indices[j]] = early;

            i += 1;
            j -= 1;
        }
    }

    fn perform_knot<T: Copy>(slice: &mut [T], start: usize, length: usize) {
        let indices = indices_wrapping(slice.len(), start, length);
        reverse_by_indices(slice, &indices);
    }

    fn knot_hash<T: Copy>(
        slice: &mut [T],
        lengths: &[u8],
        mut current: usize,
        mut skip_size: usize,
    ) -> (usize, usize) {
        let len = slice.len();
        for &length in lengths {
            perform_knot(slice, current, length as usize);
            current = (current + length as usize + skip_size) % len;
            skip_size += 1;
        }
        (current, skip_size)
    }

    fn knot_hash_n<T: Copy>(slice: &mut [T], lengths: &[u8], n: u64) {
        let mut current = 0;
        let mut skip_size = 0;
        for _ in 0..n {
            let (new_current, new_skip_size) = knot_hash(slice, lengths, current, skip_size);
            current = new_current;
            skip_size = new_skip_size;
        }
    }

    fn hash_and_multiply(slice: &mut [u8], lengths: &[u8]) -> u64 {
        knot_hash(slice, lengths, 0, 0);
        u64::from(slice[0]) * u64::from(slice[1])
    }

    fn add_suffix(lengths: &[u8]) -> Vec<u8> {
        let mut vec = Vec::from(lengths);
        vec.append(&mut vec![17, 31, 73, 47, 23]);
        vec
    }

    fn byte_as_hexadecimal(byte: u8) -> String {
        format!("{:02x}", byte)
    }

    fn full_hash(slice: &mut [u8], lengths: &[u8]) -> String {
        let lengths_suffixed = add_suffix(lengths);
        knot_hash_n(slice, &lengths_suffixed, 64);
        slice
            .chunks(16)
            .map(|chunk| chunk.iter())
            .map(|iter| iter.fold(0, |acc, x| acc ^ x))
            .map(byte_as_hexadecimal)
            .collect::<Vec<String>>()
            .join("")
    }

    pub fn full_hash_str(input: &str) -> String {
        let mut vector = initialize_vector();
        let lengths = parse_input_as_bytes(input);
        full_hash(&mut vector, &lengths)
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::test;

    mod part1 {
        use super::*;

        test!(example, "flqrgnkx", "8108", part1);
        test!(actual, file "../../../inputs/2017/14", "8214", part1);
    }

    mod part2 {
        use super::*;

        test!(example, "flqrgnkx", "1242", part2);
        test!(actual, file "../../../inputs/2017/14", "1093", part2);
    }
}
