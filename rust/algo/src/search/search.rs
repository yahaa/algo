#![allow(dead_code)]

use std::{
    cmp::{max, min},
    i32::MAX,
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
}

#[cfg(test)]
mod test {
    use super::*;

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
}
