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

pub enum Data<'a> {
    String(&'a str),
    File(&'a str),
}

impl Data<'_> {
    fn to_string(&self) -> Result<String, String> {
        match *self {
            Data::String(s) => Ok(s.to_string()),
            Data::File(path) => fs::read_to_string(path).map_err(|e| e.to_string()),
        }
    }
}

pub fn run_test(input: Data, want: Data, s: Solution) {
    let input_string = input.to_string().expect("Couldn't convert input to string");
    let want_string = want.to_string().expect("Couldn't convert want to string");
    let got = s(&input_string).expect("Error in running solution");
    if got != want_string {
        panic!("got = {:?}; want {:?}", got, want_string)
    }
}

#[macro_export]
macro_rules! testfn {
    ($name:ident, $input:expr, $want:expr, $solution:expr) => {
        #[test]
        fn $name() {
            use adventofcode_rust_aoc as aoc;
            let input: aoc::Data = $input;
            let want: aoc::Data = $want;
            let solution: aoc::Solution = $solution;
            aoc::run_test(input, want, solution);
        }
    };
}
