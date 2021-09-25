#![allow(dead_code)]
use std::cmp::{max, min};
struct Solution {}

impl Solution {
    // leetcode 1920
    pub fn build_array(nums: Vec<i32>) -> Vec<i32> {
        nums.iter().map(|n| nums[*n as usize]).collect()
    }

    // leetcode 36
    pub fn is_valid_sudoku(board: Vec<Vec<char>>) -> bool {
        let mut row = vec![0 as u16; 9];
        let mut col = vec![0 as u16; 9];
        let mut cen = vec![0 as u16; 9];

        for i in 0..9 {
            for j in 0..9 {
                if board[i][j] == '.' {
                    continue;
                }

                let v = 1 << (board[i][j] as u8 - '0' as u8);
                let k = i / 3 * 3 + j / 3;

                if row[i] != 0 && row[i] & v != 0 {
                    return false;
                }

                if col[j] != 0 && col[j] & v != 0 {
                    return false;
                }

                if cen[k] != 0 && cen[k] & v != 0 {
                    return false;
                }

                row[i] |= v;
                col[j] |= v;
                cen[k] |= v;
            }
        }
        true
    }

    // leetcode 42 dp
    pub fn trap(height: Vec<i32>) -> i32 {
        let n = height.len();
        let mut a1 = vec![0; n];
        let mut a2 = vec![0; n];

        a1[0] = height[0];
        a2[n - 1] = height[n - 1];

        for i in 1..n {
            a1[i] = max(a1[i], max(a1[i - 1], height[i]));
        }

        for i in (0..n - 1).rev() {
            a2[i] = max(a2[i], max(a2[i + 1], height[i]));
        }

        let mut result = 0;

        for i in 0..n {
            result += min(a1[i], a2[i]) - height[i];
        }

        result
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn build_array() {
        assert_eq!(
            vec![4, 5, 0, 1, 2, 3],
            Solution::build_array(vec![5, 0, 1, 2, 3, 4])
        );
    }

    #[test]
    fn is_valid_sudoku() {
        let input = vec![
            vec!['5', '3', '.', '.', '7', '.', '.', '.', '.'],
            vec!['6', '.', '.', '1', '9', '5', '.', '.', '.'],
            vec!['.', '9', '8', '.', '.', '.', '.', '6', '.'],
            vec!['8', '.', '.', '.', '6', '.', '.', '.', '3'],
            vec!['4', '.', '.', '8', '.', '3', '.', '.', '1'],
            vec!['7', '.', '.', '.', '2', '.', '.', '.', '6'],
            vec!['.', '6', '.', '.', '.', '.', '2', '8', '.'],
            vec!['.', '.', '.', '4', '1', '9', '.', '.', '5'],
            vec!['.', '.', '.', '.', '8', '.', '.', '7', '9'],
        ];

        assert_eq!(true, Solution::is_valid_sudoku(input));
    }
}
