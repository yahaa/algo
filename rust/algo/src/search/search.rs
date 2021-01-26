#![allow(dead_code)]

struct Solution {}

impl Solution {
    pub fn count_negatives(grid: Vec<Vec<i32>>) -> i32 {
        let mut res = 0;
        let (mut i, mut j) = (0 as i32, (grid[0].len() - 1) as i32);
        while i < grid.len() as i32 && j >= 0 {
            if grid[i as usize][j as usize] >= 0 {
                i += 1;
            } else {
                res += grid.len() as i32 - i;
                j -= 1;
            }
        }

        res
    }
}
