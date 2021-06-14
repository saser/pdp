use adventofcode_rust_aoc as aoc;
use nalgebra;

type PowerGrid = nalgebra::DMatrix<i64>;
type RowSumGrid = nalgebra::DMatrix<i64>;
type ColSumGrid = nalgebra::DMatrix<i64>;
type StencilGrid = nalgebra::DMatrix<(usize, i64)>;

pub fn part1(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::One)
}

pub fn part2(input: &str) -> Result<String, String> {
    solve(input, aoc::Part::Two)
}

fn solve(input: &str, part: aoc::Part) -> Result<String, String> {
    let serial = input.trim().parse::<i64>().unwrap();
    let power_grid = PowerGrid::from_fn(300, 300, power_level(serial));
    match part {
        aoc::Part::One => {
            let stencil_grid = stencil_grid(&power_grid, 3);
            let (x, y, (_size, _value)) = max_stencil_xy(&stencil_grid);
            Ok(format!("{},{}", x, y))
        }
        aoc::Part::Two => {
            let max_stencil_grid = max_stencil_grid(&power_grid);
            let (x, y, (size, _value)) = max_stencil_xy(&max_stencil_grid);
            Ok(format!("{},{},{}", x, y, size))
        }
    }
}

fn ij_to_xy((i, j): (usize, usize)) -> (usize, usize) {
    // `x` denotes column, and thus depends on `j`. Likewise, `y` denotes row, and thus depends on `i`.
    let x = 1 + j;
    let y = 1 + i;
    (x, y)
}

fn power_level(serial: i64) -> impl Fn(usize, usize) -> i64 {
    move |i, j| {
        let (x, y) = ij_to_xy((i, j));
        let x = x as i64;
        let y = y as i64;
        let rack_id = x + 10;
        let mut power = y * rack_id;
        power += serial;
        power *= rack_id;
        power /= 100;
        power %= 10;
        power -= 5;
        power
    }
}

fn sum_grid(power_grid: &PowerGrid) -> (RowSumGrid, ColSumGrid) {
    let (nrows, ncols) = power_grid.shape();
    let mut row_sum_grid = RowSumGrid::zeros(nrows, ncols);
    let mut col_sum_grid = ColSumGrid::zeros(nrows, ncols);
    for i in (0..nrows).rev() {
        for j in (0..ncols).rev() {
            let value = power_grid[(i, j)];
            let below = (i + 1).min(nrows - 1);
            let right = (j + 1).min(ncols - 1);
            let below_sum = col_sum_grid[(below, j)];
            let right_sum = row_sum_grid[(i, right)];
            col_sum_grid[(i, j)] = value + below_sum;
            row_sum_grid[(i, j)] = value + right_sum;
        }
    }
    (row_sum_grid, col_sum_grid)
}

fn all_stencils_up_to(
    power_grid: &PowerGrid,
    row_sum_grid: &RowSumGrid,
    col_sum_grid: &ColSumGrid,
    (i, j): (usize, usize),
    max_size: usize,
) -> Vec<(usize, i64)> {
    let mut stencils = Vec::with_capacity(max_size);
    let mut previous_stencil = 0;
    for size in 1..=max_size {
        let row = i + size - 1;
        let col = j + size - 1;
        let corner_pos = (row, col);

        let row_sum_start = (row, j);
        let row_sum = row_sum_grid[row_sum_start] - row_sum_grid[corner_pos];

        let col_sum_start = (i, col);
        let col_sum = col_sum_grid[col_sum_start] - col_sum_grid[corner_pos];

        let corner = power_grid[corner_pos];
        let stencil = previous_stencil + row_sum + col_sum + corner;

        stencils.push((size, stencil));
        previous_stencil = stencil;
    }
    stencils
}

fn max_stencil_up_to(
    power_grid: &PowerGrid,
    row_sum_grid: &RowSumGrid,
    col_sum_grid: &ColSumGrid,
    (i, j): (usize, usize),
    max_size: usize,
) -> (usize, i64) {
    all_stencils_up_to(power_grid, &row_sum_grid, &col_sum_grid, (i, j), max_size)
        .into_iter()
        .max_by_key(|&(_size, value)| value)
        .unwrap()
}

fn max_stencil(
    power_grid: &PowerGrid,
    row_sum_grid: &RowSumGrid,
    col_sum_grid: &ColSumGrid,
    (i, j): (usize, usize),
) -> (usize, i64) {
    let size = power_grid.nrows();
    let max_size = size - i.max(j);
    max_stencil_up_to(power_grid, row_sum_grid, col_sum_grid, (i, j), max_size)
}

fn stencil_grid(power_grid: &PowerGrid, size: usize) -> StencilGrid {
    let matrix_size = 300 - size + 1;
    let stencil = |i, j| {
        let sum = power_grid.slice((i, j), (size, size)).iter().sum();
        (size, sum)
    };
    StencilGrid::from_fn(matrix_size, matrix_size, stencil)
}

fn max_stencil_grid(power_grid: &PowerGrid) -> StencilGrid {
    let (nrows, ncols) = power_grid.shape();
    let (row_sum_grid, col_sum_grid) = sum_grid(power_grid);
    StencilGrid::from_fn(nrows, ncols, |i, j| {
        max_stencil(power_grid, &row_sum_grid, &col_sum_grid, (i, j))
    })
}

fn max_stencil_xy(stencil_grid: &StencilGrid) -> (usize, usize, (usize, i64)) {
    let (nrows, ncols) = stencil_grid.shape();
    let (mut max_i, mut max_j) = (0, 0);
    let (mut max_size, mut max_value) = (0, 0);
    for i in 0..nrows {
        for j in 0..ncols {
            let (size, value) = stencil_grid[(i, j)];
            if value > max_value {
                max_size = size;
                max_value = value;
                max_i = i;
                max_j = j;
            }
        }
    }
    let (x, y) = ij_to_xy((max_i, max_j));
    (x, y, (max_size, max_value))
}

#[cfg(test)]
mod day11_test;
