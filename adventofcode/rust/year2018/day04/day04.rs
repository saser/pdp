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
    let mut sorted_events = parse_input(&input);
    sorted_events.sort();
    match part {
        aoc::Part::One => Ok(strategy_1(&sorted_events).to_string()),
        aoc::Part::Two => Ok(strategy_2(&sorted_events).to_string()),
    }
}

fn parse_input(input: &str) -> Vec<Event> {
    input
        .lines()
        .map(Event::from_str)
        .map(Result::unwrap)
        .collect()
}

fn strategy_1(sorted_events: &[Event]) -> u64 {
    let guard_events = gather_guard_events(sorted_events);
    let (id, (_total_sleep, most_sleeping_minute, _most_times_asleep)) = guard_events
        .iter()
        .map(|(id, events)| (id, calculate_sleeping(events)))
        .max_by_key(|&(_id, (total_sleep, _most_sleeping_minute, _most_times_asleep))| total_sleep)
        .unwrap();
    id * u64::from(most_sleeping_minute)
}

fn strategy_2(sorted_events: &[Event]) -> u64 {
    let guard_events = gather_guard_events(sorted_events);
    let (id, (_total_sleep, most_sleeping_minute, _most_times_asleep)) = guard_events
        .iter()
        .map(|(id, events)| (id, calculate_sleeping(events)))
        .max_by_key(
            |&(_id, (_total_sleep, _most_sleeping_minute, most_times_asleep))| most_times_asleep,
        )
        .unwrap();
    id * u64::from(most_sleeping_minute)
}

fn gather_guard_events(events: &[Event]) -> collections::HashMap<u64, Vec<Vec<(u32, EventType)>>> {
    let first_event = events[0];
    let first_event_minute = first_event.time.minute;
    let first_event_type = first_event.event_type;
    let mut current_guard = if let EventType::BeginsShift(id) = first_event_type {
        id
    } else {
        panic!("First event is not a begins shift event");
    };
    let mut current_events = vec![(first_event_minute, first_event_type)];
    let mut map = collections::HashMap::new();
    for &event in &events[1..] {
        let event_minute = event.time.minute;
        let event_type = event.event_type;
        if let EventType::BeginsShift(id) = event_type {
            map.entry(current_guard)
                .or_insert_with(Vec::new)
                .push(current_events);
            current_guard = id;
            current_events = Vec::new();
        } else {
            current_events.push((event_minute, event_type));
        }
    }
    map.entry(current_guard)
        .or_insert_with(Vec::new)
        .push(current_events);
    map
}

fn calculate_sleeping(events: &[Vec<(u32, EventType)>]) -> (u32, u32, u32) {
    let mut combined = events
        .into_iter()
        .cloned()
        .flatten()
        .filter(|(_event_minute, event_type)| {
            if let EventType::BeginsShift(_) = event_type {
                false
            } else {
                true
            }
        })
        .collect::<Vec<(u32, EventType)>>();
    combined.sort();
    let mut last_event_minute = 0;
    let mut total_sleep = 0;
    let mut times_asleep = 0;
    let mut most_sleeping_minute = 0;
    let mut most_times_asleep = 0;
    for &(event_minute, event_type) in &combined {
        total_sleep += (event_minute - last_event_minute) * times_asleep;
        match event_type {
            EventType::BeginsShift(_) => unreachable!(),
            EventType::FallsAsleep => times_asleep += 1,
            EventType::WakesUp => times_asleep -= 1,
        };
        if times_asleep > most_times_asleep {
            most_sleeping_minute = event_minute;
            most_times_asleep = times_asleep;
        }
        last_event_minute = event_minute;
    }
    (total_sleep, most_sleeping_minute, most_times_asleep)
}

#[derive(Clone, Copy, Debug, Eq, Hash, Ord, PartialEq, PartialOrd)]
struct Time {
    year: u32,
    month: u32,
    day: u32,
    hour: u32,
    minute: u32,
}

#[derive(Clone, Copy, Debug, Eq, Hash, Ord, PartialEq, PartialOrd)]
struct Event {
    time: Time,
    event_type: EventType,
}

impl FromStr for Event {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        lazy_static::lazy_static! {
            static ref EVENT_RE: regex::Regex =
                regex::Regex::new(r"\[(?P<Y>\d{4})\-(?P<m>\d{2})\-(?P<d>\d{2}) (?P<H>\d{2}):(?P<M>\d{2})\] (?P<event_type>.+)")
                    .unwrap();
        }
        let captures = EVENT_RE.captures(s).unwrap();
        let time = Time {
            year: u32::from_str(&captures["Y"]).unwrap(),
            month: u32::from_str(&captures["m"]).unwrap(),
            day: u32::from_str(&captures["d"]).unwrap(),
            hour: u32::from_str(&captures["H"]).unwrap(),
            minute: u32::from_str(&captures["M"]).unwrap(),
        };
        let event_type = EventType::from_str(&captures["event_type"]).unwrap();
        Ok(Event {
            time,
            event_type,
        })
    }
}

#[derive(Clone, Copy, Debug, Eq, Hash, Ord, PartialEq, PartialOrd)]
enum EventType {
    BeginsShift(u64),
    FallsAsleep,
    WakesUp,
}

impl FromStr for EventType {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        lazy_static::lazy_static! {
            static ref BEGIN_RE: regex::Regex = regex::Regex::new(r"Guard #(?P<id>\d+) begins shift").unwrap();
        }
        if let Some(caps) = BEGIN_RE.captures(s) {
            let id = u64::from_str(&caps["id"]).unwrap();
            Ok(EventType::BeginsShift(id))
        } else {
            match s {
                "falls asleep" => Ok(EventType::FallsAsleep),
                "wakes up" => Ok(EventType::WakesUp),
                _ => Err(format!("invalid event: \"{}\"", s)),
            }
        }
    }
}

#[cfg(test)]
mod day04_test;
