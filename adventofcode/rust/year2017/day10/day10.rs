use std::str::FromStr;

use adventofcode_rust_aoc as aoc;

pub fn part1(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::One)
}

pub fn part2(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::Two)
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let mut vector = initialize_vector();
    match part {
        aoc::Part::One => {
            let lengths = parse_input_as_lengths(input.trim());
            Ok(hash_and_multiply(&mut vector, &lengths).to_string())
        }
        aoc::Part::Two => {
            let lengths = parse_input_as_bytes(input.trim());
            Ok(full_hash(&mut vector, &lengths))
        }
    }
}

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

#[cfg(test)]
mod day10_test;
