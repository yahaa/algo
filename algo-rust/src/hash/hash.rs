#![allow(dead_code)]
use std::collections::HashMap;

struct Solution {}

impl Solution {
    // leetcode 1282
    pub fn group_the_people(group_sizes: Vec<i32>) -> Vec<Vec<i32>> {
        let mut result = Vec::new();
        let mut group_map = HashMap::new();

        for (i, size) in group_sizes.iter().enumerate() {
            let group = group_map.entry(size).or_insert(vec![]);

            group.push(i as i32);

            if group.len() >= *size as usize {
                if let Some((_, group)) = group_map.remove_entry(size) {
                    result.push(group);
                }
            }
        }

        result
    }

    // Longest Consecutive Sequence
    pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
        let (mut map, mut ans) = (HashMap::new(), 0);

        for n in nums {
            if map.contains_key(&n) {
                continue;
            }

            map.insert(n, 1);

            let (mut left, mut right, mut length) = (n, n, 1);
            if let Some(v) = map.get(&(n - 1)) {
                left = n - v;
                length += map.get(&left).cloned().unwrap_or(0);
            };

            if let Some(v) = map.get(&(n + 1)) {
                right = n + v;
                length += map.get(&right).cloned().unwrap_or(0);
            }

            map.insert(left, length);
            map.insert(right, length);
            ans = std::cmp::max(ans, length);
        }
        ans
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn group_the_people() {
        let group_sizes = vec![3, 3, 3, 3, 3, 1, 3];
        let want = vec![vec![0, 1, 2], vec![5], vec![3, 4, 6]];

        assert_eq!(want, Solution::group_the_people(group_sizes));
    }
    #[test]
    fn longest_consecutive() {
        let nums = vec![100, 4, 200, 1, 3, 2];
        let want = 4;

        assert_eq!(want, Solution::longest_consecutive(nums));
    }
}
