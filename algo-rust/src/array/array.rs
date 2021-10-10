#![allow(dead_code)]
use std::cmp::{max, min};
use std::collections::BinaryHeap;
use std::{
    cmp::{Ord, Ordering},
    vec,
};
struct Solution {}

#[derive(Copy, Clone, Eq, PartialEq, Ord, Debug)]
struct Val {
    pub v: i32,
    pub index: (i32, i32),
}

impl Val {
    pub fn new(v: i32, index: (i32, i32)) -> Self {
        Val { v, index }
    }
}

impl PartialOrd for Val {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        other.v.partial_cmp(&self.v)
    }
}

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

    // leetcode 407
    pub fn trap_rain_water(mut height: Vec<Vec<i32>>) -> i32 {
        let (n, m) = (height.len(), height[0].len());
        let mut heap = BinaryHeap::new();

        for i in 0..n {
            heap.push(Val::new(height[i][0], (i as i32, 0)));
            heap.push(Val::new(height[i][m - 1], (i as i32, (m - 1) as i32)));
            height[i][0] = -1;
            height[i][m - 1] = -1;
        }

        for j in 1..m - 1 {
            heap.push(Val::new(height[0][j], (0, j as i32)));
            heap.push(Val::new(height[n - 1][j], ((n - 1) as i32, j as i32)));

            height[0][j] = -1;
            height[n - 1][j] = -1;
        }

        let dir = vec![[-1, 0], [1, 0], [0, -1], [0, 1]];
        let mut result = 0;

        while let Some(top) = heap.pop() {
            dir.iter().for_each(|d| {
                let x = d[0] + top.index.0;
                let y = d[1] + top.index.1;

                if x >= n as i32 || x < 0 || y >= m as i32 || y < 0 {
                    return;
                }

                let x = x as usize;
                let y = y as usize;

                if height[x][y] == -1 {
                    return;
                }

                if height[x][y] < top.v {
                    result += top.v - height[x][y];
                    heap.push(Val::new(top.v, (x as i32, y as i32)));
                } else {
                    heap.push(Val::new(height[x][y], (x as i32, y as i32)));
                }
                height[x][y] = -1;
            });
        }
        result
    }

    // leetcode 51
    pub fn solve_n_queens(n: i32) -> Vec<Vec<String>> {
        fn validate(board: &Vec<Vec<char>>, i: i32, j: i32, n: i32) -> bool {
            let (ri, rj) = (i - min(i, j), j - min(i, j));
            let (li, lj) = (i - min(i, n - 1 - j), j + min(i, n - 1 - j));

            for k in 0..n {
                if board[i as usize][k as usize] == 'Q'
                    || board[k as usize][j as usize] == 'Q'
                    || (ri + k < n
                        && rj + k < n
                        && board[(ri + k) as usize][(rj + k) as usize] == 'Q')
                    || (li + k < n
                        && lj - k >= 0
                        && board[(li + k) as usize][(lj - k) as usize] == 'Q')
                {
                    return false;
                }
            }

            true
        }

        fn dfs(i: i32, n: i32, result: &mut Vec<Vec<String>>, board: &mut Vec<Vec<char>>) {
            if i == n {
                let mut res = vec![];
                board.iter().for_each(|row| {
                    res.push(row.into_iter().collect());
                });

                result.push(res);
                return;
            }

            for j in 0..n {
                if validate(board, i, j, n) {
                    board[i as usize][j as usize] = 'Q';
                    dfs(i + 1, n, result, board);
                    board[i as usize][j as usize] = '.';
                }
            }
        }

        let mut result = Vec::new();
        let mut board = vec![vec!['.'; n as usize]; n as usize];

        dfs(0, n, &mut result, &mut board);

        result
    }

    // leetcode 52
    pub fn total_n_queens(n: i32) -> i32 {
        fn validate(board: &Vec<Vec<char>>, i: i32, j: i32, n: i32) -> bool {
            let (ri, rj) = (i - min(i, j), j - min(i, j));
            let (li, lj) = (i - min(i, n - 1 - j), j + min(i, n - 1 - j));

            for k in 0..n {
                if board[i as usize][k as usize] == 'Q'
                    || board[k as usize][j as usize] == 'Q'
                    || (ri + k < n
                        && rj + k < n
                        && board[(ri + k) as usize][(rj + k) as usize] == 'Q')
                    || (li + k < n
                        && lj - k >= 0
                        && board[(li + k) as usize][(lj - k) as usize] == 'Q')
                {
                    return false;
                }
            }

            true
        }

        fn dfs(i: i32, n: i32, result: &mut i32, board: &mut Vec<Vec<char>>) {
            if i == n {
                *result += 1;
                return;
            }

            for j in 0..n {
                if validate(board, i, j, n) {
                    board[i as usize][j as usize] = 'Q';
                    dfs(i + 1, n, result, board);
                    board[i as usize][j as usize] = '.';
                }
            }
        }

        let (mut board, mut result) = (vec![vec!['.'; n as usize]; n as usize], 0);

        dfs(0, n, &mut result, &mut board);

        result
    }
}

#[cfg(test)]
mod test {
    use std::collections::BinaryHeap;

    use super::*;

    #[test]
    fn solve_n_queens() {
        let res = vec![vec!["Q"]];
        assert_eq!(res, Solution::solve_n_queens(1));

        let res = vec![
            vec![".Q..", "...Q", "Q...", "..Q."],
            vec!["..Q.", "Q...", "...Q", ".Q.."],
        ];
        assert_eq!(res, Solution::solve_n_queens(4));

        let res = vec![
            vec!["Q....", "..Q..", "....Q", ".Q...", "...Q."],
            vec!["Q....", "...Q.", ".Q...", "....Q", "..Q.."],
            vec![".Q...", "...Q.", "Q....", "..Q..", "....Q"],
            vec![".Q...", "....Q", "..Q..", "Q....", "...Q."],
            vec!["..Q..", "Q....", "...Q.", ".Q...", "....Q"],
            vec!["..Q..", "....Q", ".Q...", "...Q.", "Q...."],
            vec!["...Q.", "Q....", "..Q..", "....Q", ".Q..."],
            vec!["...Q.", ".Q...", "....Q", "..Q..", "Q...."],
            vec!["....Q", ".Q...", "...Q.", "Q....", "..Q.."],
            vec!["....Q", "..Q..", "Q....", "...Q.", ".Q..."],
        ];

        assert_eq!(res, Solution::solve_n_queens(5));
    }

    #[test]
    fn build_array() {
        assert_eq!(
            vec![4, 5, 0, 1, 2, 3],
            Solution::build_array(vec![5, 0, 1, 2, 3, 4])
        );
    }

    #[test]
    fn trap_rain_water() {
        let input = vec![
            vec![9, 9, 9, 9, 9, 9, 8, 9, 9, 9, 9],
            vec![9, 0, 0, 0, 0, 0, 1, 0, 0, 0, 9],
            vec![9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9],
            vec![9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9],
            vec![9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9],
        ];
        assert_eq!(215, Solution::trap_rain_water(input));
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

    #[test]
    fn val_struct() {
        let mut heap = BinaryHeap::new();
        heap.push(Val::new(1, (0, 0)));
        heap.push(Val::new(3, (0, 0)));
        heap.push(Val::new(2, (0, 0)));

        assert_eq!(Val::new(1, (0, 0)), heap.pop().unwrap());
        assert_eq!(Val::new(2, (0, 0)), heap.pop().unwrap());
        assert_eq!(Val::new(3, (0, 0)), heap.pop().unwrap());
    }
}
