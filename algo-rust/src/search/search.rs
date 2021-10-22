#![allow(dead_code)]

use std::{
    cmp::{max, min},
    collections::HashMap,
};

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

    pub fn peak_index_in_mountain_array(arr: Vec<i32>) -> i32 {
        if arr.len() < 3 {
            return 0;
        }

        let mut ans = 0;
        for i in 1..arr.len() - 1 {
            if arr[i - 1] < arr[i] && arr[i] > arr[i + 1] {
                ans = i as i32;
                break;
            }
        }

        ans
    }

    pub fn my_sqrt(x: i32) -> i32 {
        if x == 0 {
            return x;
        }

        let (mut left, mut right) = (1, x);
        loop {
            let mid = (right - left) / 2 + left;

            if mid > x / mid {
                right = mid - 1;
            } else {
                if (mid + 1) > x / (mid + 1) {
                    return mid;
                }
                left = mid + 1;
            }
        }
    }

    // leetcode 35
    pub fn search_insert(nums: Vec<i32>, target: i32) -> i32 {
        let (mut left, mut right) = (0 as i32, nums.len() as i32 - 1);
        while left <= right {
            let mid = (right - left) / 2 + left;
            match nums[mid as usize].cmp(&target) {
                std::cmp::Ordering::Equal => return mid,
                std::cmp::Ordering::Less => left = mid + 1,
                std::cmp::Ordering::Greater => right = mid - 1,
            }
        }

        left as i32
    }

    // leetcode 33
    pub fn search_1(nums: Vec<i32>, target: i32) -> i32 {
        if nums.len() == 0 {
            return -1;
        }

        let (mut left, mut right) = (0, nums.len() as i32 - 1);

        while left <= right {
            let mid = (right - left) / 2 + left;

            if nums[mid as usize] == target {
                return mid;
            }

            if nums[left as usize] <= nums[mid as usize] {
                if nums[left as usize] <= target && nums[mid as usize] >= target {
                    right = mid - 1;
                } else {
                    left = mid + 1;
                }
            } else {
                if nums[mid as usize] <= target && nums[right as usize] >= target {
                    left = mid + 1;
                } else {
                    right = mid - 1;
                }
            }
        }

        -1
    }

    // leetcode 81
    pub fn search(nums: Vec<i32>, target: i32) -> bool {
        if nums.len() == 0 {
            return false;
        }

        let (mut left, mut right) = (0, nums.len() as i32 - 1);

        while left <= right {
            let mid = (right - left) / 2 + left;

            if nums[mid as usize] == target {
                return true;
            }

            match nums[left as usize].cmp(&nums[mid as usize]) {
                std::cmp::Ordering::Equal => {
                    left += 1;
                    continue;
                }
                std::cmp::Ordering::Less => {
                    if nums[left as usize] <= target && nums[mid as usize] > target {
                        right = mid - 1;
                    } else {
                        left = mid + 1;
                    }
                }
                std::cmp::Ordering::Greater => {
                    if nums[mid as usize] < target && nums[right as usize] >= target {
                        left = mid + 1;
                    } else {
                        right = mid - 1;
                    }
                }
            }
        }

        false
    }

    // leetcode 153
    pub fn find_min_1(nums: Vec<i32>) -> i32 {
        let (mut left, mut right) = (0, nums.len() as i32 - 1);

        while left < right {
            let mid = (right - left) / 2 + left;
            if nums[mid as usize] > nums[right as usize] {
                left = mid + 1;
            } else {
                right = mid;
            }
        }
        nums[left as usize]
    }

    // leetcode 154
    pub fn find_min(nums: Vec<i32>) -> i32 {
        let (mut left, mut right) = (0, nums.len() as i32 - 1);

        while left < right {
            let mid = (right - left) / 2 + left;
            match nums[mid as usize].cmp(&nums[right as usize]) {
                std::cmp::Ordering::Equal => right -= 1,
                std::cmp::Ordering::Less => right = mid,
                std::cmp::Ordering::Greater => left = mid + 1,
            }
        }
        nums[left as usize]
    }

    // leetcode 162
    pub fn find_peak_element(nums: Vec<i32>) -> i32 {
        if nums.len() <= 1 {
            return 0;
        }

        let mut index = 0;
        for i in 1..nums.len() {
            if i + 1 < nums.len() && nums[i] > nums[i - 1] && nums[i] > nums[i + 1] {
                return i as i32;
            }

            if nums[i] > nums[index] {
                index = i;
            }
        }

        return index as i32;
    }

    // leetcode 167
    pub fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
        let mut res = vec![0, 0];

        for i in 0..numbers.len() {
            let r = target - numbers[i];

            if let Ok(index) = numbers.binary_search(&r) {
                res[0] = i as i32 + 1;
                res[1] = index as i32 + 1;
                break;
            }
        }

        res
    }

    // leetcode 174
    pub fn calculate_minimum_hp(dungeon: Vec<Vec<i32>>) -> i32 {
        let (n, m) = (dungeon.len(), dungeon[0].len());
        let mut dp = vec![vec![i32::MAX; m + 1]; n + 1];
        dp[n][m - 1] = 1;
        dp[n - 1][m] = 1;

        for i in (0..n).rev() {
            for j in (0..m).rev() {
                let m = min(dp[i + 1][j], dp[i][j + 1]);
                dp[i][j] = max(m - dungeon[i][j], 1);
            }
        }

        dp[0][0]
    }

    // leetcode 209
    pub fn min_sub_array_len(target: i32, nums: Vec<i32>) -> i32 {
        let (mut sum, mut res) = (0, i32::MAX);
        let (mut i, mut j) = (0 as usize, 0 as usize);

        while i <= j && j < nums.len() {
            while sum < target && j < nums.len() {
                sum += nums[j];
                j += 1;
            }

            while sum >= target && i < nums.len() {
                if sum >= target {
                    res = min(res, (j - i) as i32);
                }
                sum -= nums[i];
                i += 1;
            }
        }

        if res == i32::MAX {
            res = 0;
        }

        res
    }

    // leetcode 275
    pub fn h_index(citations: Vec<i32>) -> i32 {
        let n = citations.len() as i32;
        let (mut left, mut right) = (0 as i32, n as i32 - 1);

        while left <= right {
            let mid = (right - left) / 2 + left;

            match citations[mid as usize].cmp(&(n - mid)) {
                std::cmp::Ordering::Equal => {
                    return n - mid;
                }
                std::cmp::Ordering::Less => {
                    left = mid + 1;
                }
                std::cmp::Ordering::Greater => {
                    right = mid - 1;
                }
            }
        }

        n - left
    }

    // leetcode 278
    pub fn first_bad_version(&self, n: i32) -> i32 {
        let (mut first, mut last) = (1, n);

        while first <= last {
            let mid = (last - first) / 2 + first;

            if self.is_bad_version(mid) {
                last = mid - 1;
            } else {
                first = mid + 1;
            }
        }

        first
    }

    // isBadVersion for leetcode 278
    fn is_bad_version(&self, _num: i32) -> bool {
        return false;
    }

    // leetcode 220
    pub fn contains_nearby_almost_duplicate(nums: Vec<i32>, k: i32, t: i32) -> bool {
        fn abs(a: i64) -> i64 {
            if a > 0 {
                return a;
            }

            -a
        }
        let t = t as i64;

        for i in 0..nums.len() {
            for j in 1..=k as usize {
                if i + j < nums.len() && abs(nums[i] as i64 - nums[i + j] as i64) <= t {
                    return true;
                }
            }
        }

        return false;
    }

    // leetcode 130
    pub fn solve(board: &mut Vec<Vec<char>>) {
        fn dfs(board: &mut Vec<Vec<char>>, x: usize, y: usize, n: usize, m: usize) {
            if x >= n || x < 0 as usize || y >= m || y < 0 as usize || board[x][y] != 'O' {
                return;
            }

            board[x][y] = 'A';
            dfs(board, x, y + 1, n, m);
            dfs(board, x, y - 1, n, m);
            dfs(board, x + 1, y, n, m);
            dfs(board, x - 1, y, n, m);
        }

        let (n, m) = (board.len(), board[0].len());

        for i in 0..n {
            dfs(board, i, 0, n, m);
            dfs(board, i, m - 1, n, m);
        }

        for j in 0..m {
            dfs(board, 0, j, n, m);
            dfs(board, n - 1, j, n, m);
        }

        for i in 0..n {
            for j in 0..m {
                if board[i][j] == 'A' {
                    board[i][j] = 'O';
                } else if board[i][j] == 'O' {
                    board[i][j] = 'X';
                }
            }
        }
    }
    // leetcode 200
    pub fn num_islands(grid: Vec<Vec<char>>) -> i32 {
        fn dfs(
            grid: &Vec<Vec<char>>,
            visited: &mut Vec<Vec<bool>>,
            x: i32,
            y: i32,
            n: i32,
            m: i32,
        ) {
            if x < 0
                || x >= n
                || y < 0
                || y >= m
                || visited[x as usize][y as usize]
                || grid[x as usize][y as usize] == '0'
            {
                return;
            }

            visited[x as usize][y as usize] = true;

            dfs(grid, visited, x, y + 1, n, m);
            dfs(grid, visited, x, y - 1, n, m);
            dfs(grid, visited, x + 1, y, n, m);
            dfs(grid, visited, x - 1, y, n, m);
        }

        let (n, m) = (grid.len(), grid[0].len());

        let (mut visited, mut count) = (vec![vec![false; m]; n], 0);

        for i in 0..n {
            for j in 0..m {
                if grid[i][j] == '1' && !visited[i][j] {
                    count += 1;
                    dfs(&grid, &mut visited, i as i32, j as i32, n as i32, m as i32);
                }
            }
        }

        count
    }
    // leetcode 131
    pub fn partition(s: String) -> Vec<Vec<String>> {
        fn dfs(
            chars: &Vec<char>,
            index: usize,
            len: usize,
            path: &mut Vec<String>,
            result: &mut Vec<Vec<String>>,
        ) {
            if index == len {
                result.push(path.clone());
                return;
            }

            for i in index..len {
                if !Solution::is_palindrome(&chars, index, i) {
                    continue;
                }
                path.push(chars[index..(i + 1)].iter().collect());
                dfs(chars, i + 1, len, path, result);
                path.pop();
            }
        }

        let mut result: Vec<Vec<String>> = Vec::new();

        if s.len() == 0 {
            return result;
        }

        let mut path: Vec<String> = Vec::new();

        dfs(&s.chars().collect(), 0, s.len(), &mut path, &mut result);

        result
    }

    // leetcode 131 help
    fn is_palindrome(chars: &Vec<char>, left: usize, right: usize) -> bool {
        let (mut left, mut right) = (left, right);

        while left < right {
            if chars[left] != chars[right] {
                return false;
            }
            left += 1;
            right -= 1;
        }

        true
    }

    // leetcode 74
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        let (n, m) = (matrix.len() as i32, matrix[0].len() as i32);
        let (mut start, mut end) = (0, n * m - 1);

        while start <= end {
            let mid = (end - start) / 2 + start;
            let (i, j) = (mid / m, mid % m);

            match target.cmp(&matrix[i as usize][j as usize]) {
                std::cmp::Ordering::Equal => return true,
                std::cmp::Ordering::Less => end = mid - 1,
                std::cmp::Ordering::Greater => start = mid + 1,
            }
        }
        return false;
    }

    // leetcode 771
    pub fn num_jewels_in_stones(jewels: String, stones: String) -> i32 {
        let (mut jes, mut res) = (HashMap::new(), 0);
        jewels.chars().for_each(|x| *jes.entry(x).or_insert(1) += 1);

        stones.chars().for_each(|x| {
            if let Some(_a) = jes.get(&x) {
                res += 1
            }
        });

        res
    }
    // leetcode 1528
    pub fn restore_string(s: String, indices: Vec<i32>) -> String {
        let mut result = vec!['0'; s.len()];
        let ss: Vec<char> = s.chars().collect();
        indices
            .into_iter()
            .enumerate()
            .for_each(|(i, x)| result[x as usize] = ss[i]);

        result.into_iter().collect()
    }

    // leetcode 1678
    pub fn interpret(command: String) -> String {
        command.replace("(al)", "al").replace("()", "o")
    }

    // leetcode 1173
    pub fn count_matches(items: Vec<Vec<String>>, rule_key: String, rule_value: String) -> i32 {
        let key = match rule_key.as_str() {
            "type" => 0,
            "color" => 1,
            _ => 2,
        };

        items.iter().filter(|x| x[key] == rule_value).count() as i32
    }

    // leetcode 1221
    pub fn balanced_string_split(s: String) -> i32 {
        let (mut l, mut r, mut ans) = (0, 0, 0);

        s.chars().for_each(|x| {
            match x {
                'L' => l += 1,
                'R' => r += 1,
                _ => {}
            }

            if l == r {
                l = 0;
                r = 0;
                ans += 1;
            }
        });

        ans
    }

    // leetcode 1859
    pub fn sort_sentence(s: String) -> String {
        let ss: Vec<&str> = s.split(" ").collect();
        let mut result = vec![""; ss.len()];

        ss.into_iter().filter(|x| return x.len() > 0).for_each(|x| {
            let index: usize = x[x.len() - 1..].parse().unwrap();
            result[index - 1] = &x[..x.len() - 1];
        });

        result.join(" ")
    }

    // leetcode 1920
    pub fn build_array(nums: Vec<i32>) -> Vec<i32> {
        nums.iter().map(|n| nums[*n as usize]).collect()
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn partition() {
        let result = vec![vec!["a", "a", "b"], vec!["aa", "b"]];
        assert_eq!(result, Solution::partition("aab".to_string()));

        let result = vec![vec!["a"]];
        assert_eq!(result, Solution::partition("a".to_string()));
    }
    #[test]
    fn num_islands() {
        let grid = vec![
            vec!['1', '1', '1', '1', '0'],
            vec!['1', '1', '0', '1', '0'],
            vec!['1', '1', '0', '0', '0'],
            vec!['0', '0', '0', '0', '0'],
        ];

        assert_eq!(1, Solution::num_islands(grid));
    }

    #[test]
    fn solve() {
        let mut board = vec![
            vec!['X', 'X', 'X', 'X'],
            vec!['X', 'O', 'O', 'X'],
            vec!['X', 'X', 'O', 'X'],
            vec!['X', 'O', 'X', 'X'],
        ];

        let want = vec![
            vec!['X', 'X', 'X', 'X'],
            vec!['X', 'X', 'X', 'X'],
            vec!['X', 'X', 'X', 'X'],
            vec!['X', 'O', 'X', 'X'],
        ];

        Solution::solve(&mut board);

        assert_eq!(want, board);
    }

    #[test]
    fn my_sqrt() {
        assert_eq!(2, Solution::my_sqrt(4));
        assert_eq!(2, Solution::my_sqrt(8));
    }

    #[test]
    fn search_insert() {
        let nums = vec![1, 3, 5, 6];
        assert_eq!(2, Solution::search_insert(nums, 5));

        let nums = vec![1, 3, 5, 6];
        assert_eq!(1, Solution::search_insert(nums, 2));
        let nums = vec![1, 3, 5, 6];
        assert_eq!(4, Solution::search_insert(nums, 7));
        let nums = vec![1, 3, 5, 6];
        assert_eq!(0, Solution::search_insert(nums, 0));
    }

    #[test]
    fn search_1() {
        let nums = vec![4, 5, 6, 7, 0, 1, 2];
        assert_eq!(4, Solution::search_1(nums, 0));
        let nums = vec![4, 5, 6, 7, 0, 1, 2];
        assert_eq!(-1, Solution::search_1(nums, 3));
        let nums = vec![1];
        assert_eq!(-1, Solution::search_1(nums, 0));
    }

    #[test]
    fn search() {
        let nums = vec![4, 5, 6, 7, 0, 1, 2];
        assert_eq!(true, Solution::search(nums, 0));
        let nums = vec![4, 5, 6, 7, 0, 1, 2];
        assert_eq!(false, Solution::search(nums, 3));
        let nums = vec![1];
        assert_eq!(false, Solution::search(nums, 0));

        let nums = vec![2, 5, 6, 0, 0, 1, 2];
        assert_eq!(true, Solution::search(nums, 0));

        let nums = vec![2, 5, 6, 0, 0, 1, 2];
        assert_eq!(false, Solution::search(nums, 3));
    }

    #[test]
    fn find_min_1() {
        let nums = vec![3, 4, 5, 1, 2];
        assert_eq!(1, Solution::find_min_1(nums));
        let nums = vec![4, 5, 6, 7, 0, 1, 2];
        assert_eq!(0, Solution::find_min_1(nums));
    }

    #[test]
    fn min_sub_array_len() {
        let nums = vec![2, 3, 1, 2, 4, 3];
        assert_eq!(2, Solution::min_sub_array_len(7, nums));
        let nums = vec![1, 4, 4];
        assert_eq!(1, Solution::min_sub_array_len(4, nums));

        let nums = vec![1, 1, 1, 1, 1, 1, 1, 1];
        assert_eq!(0, Solution::min_sub_array_len(11, nums));

        let nums = vec![1, 2, 3, 4, 5];
        assert_eq!(3, Solution::min_sub_array_len(11, nums));
    }
}
