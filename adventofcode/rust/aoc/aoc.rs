use std::fmt;
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

#[macro_export]
macro_rules! test {
    ($name:ident, $input:expr, $output:expr, $solution:expr) => {
        #[test]
        fn $name() {
            let input: &str = $input;
            let output: &str = $output;
            let solution: adventofcode_rust_aoc::Solution = $solution;
            assert_eq!(output, solution(&mut Box::new(input.as_bytes())).unwrap());
        }
    };
    ($name:ident, file $file:expr, $output:expr, $solution:expr) => {
        test!($name, include_str!($file), $output, $solution);
    };
    ($name:ident, $input:expr, file $file:expr, $solution:expr) => {
        test!($name, $input, include_str!($file), $solution);
    };
    ($name:ident, file $infile:expr, file $outfile:expr, $solution:expr) => {
        test!(
            $name,
            include_str!($infile),
            include_str!($outfile),
            $solution
        );
    };
}
