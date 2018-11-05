extern crate regex;
#[macro_use]
extern crate lazy_static;

extern crate base;

use base::{Solver, YearDispatcher};

pub mod day01;

pub struct Year2016;

pub fn get_dispatcher() -> Box<dyn YearDispatcher> {
    Box::new(Year2016)
}

impl YearDispatcher for Year2016 {
    fn get_solver(&self, day: u8) -> Result<Box<dyn Solver>, String> {
        match day {
            01 => Ok(day01::get_solver()),
            _ => Err(format!("no solver for day {}", day)),
        }
    }
}
