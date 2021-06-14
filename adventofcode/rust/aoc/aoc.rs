use std::fmt;
use std::fs;
use std::str::FromStr;

pub type Solution = fn(input: &str) -> Result<String, String>;

#[derive(Copy, Clone, Debug, Eq, PartialEq)]
pub enum Part {
    One,
    Two,
}

impl FromStr for Part {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "1" => Ok(Part::One),
            "2" => Ok(Part::Two),
            _ => Err(format!("part must be 1 or 2; was: {}", s)),
        }
    }
}

impl fmt::Display for Part {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let s = match *self {
            Part::One => "1",
            Part::Two => "2",
        };
        write!(f, "{}", s)
    }
}

pub enum Input<'a> {
    String(&'a str),
    File(&'a str),
}

impl Input<'_> {
    fn to_string(&self) -> Result<String, String> {
        match *self {
            Input::String(s) => Ok(s.to_string()),
            Input::File(path) => fs::read_to_string(path).map_err(|e| e.to_string()),
        }
    }
}

pub fn run_test(input: Input, want: &str, s: Solution) {
    let input_string = input.to_string().expect("Couldn't convert input to string");
    let got = s(&input_string).expect("Error in running solution");
    if got != want {
        panic!("got = {:?}; want {:?}", got, want)
    }
}

#[macro_export]
macro_rules! testfn {
    ($name:ident, $input:expr, $want:expr, $solution:expr) => {
        #[test]
        fn $name() {
            use adventofcode_rust_aoc as aoc;
            let input: aoc::Input = $input;
            let want: &str = $want;
            let solution: aoc::Solution = $solution;
            aoc::run_test(input, want, solution);
        }
    };
}
