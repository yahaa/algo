#![allow(dead_code)]
use std::collections::HashMap;

struct FindSumPairs {
    nums1: Vec<i32>,
    nums2: Vec<i32>,
    map: HashMap<i32, i32>,
}

impl FindSumPairs {
    fn new(nums1: Vec<i32>, nums2: Vec<i32>) -> Self {
        let mut map = HashMap::new();

        for num in nums2.iter() {
            *map.entry(*num).or_insert(0) += 1;
        }

        FindSumPairs { nums1, nums2, map }
    }

    fn add(&mut self, index: i32, val: i32) {
        *self.map.entry(self.nums2[index as usize]).or_insert(1) -= 1;
        self.nums2[index as usize] += val;
        *self.map.entry(self.nums2[index as usize]).or_insert(0) += 1;
    }

    fn count(&self, tot: i32) -> i32 {
        let mut res = 0;

        for a in self.nums1.iter() {
            if let Some(v) = self.map.get(&(tot - a)) {
                res += v;
            }
        }
        res
    }
}

struct Solution {}

impl Solution {
    // leetcode 2325
    pub fn decode_message(key: String, message: String) -> String {
        let mut map = HashMap::new();
        let mut index: u8 = 97;
        for c in key.chars() {
            if !c.is_alphabetic() || map.get(&c).is_some() {
                continue;
            }

            map.insert(c, index as char);
            index += 1;
        }

        let mut result = String::new();
        for c in message.chars() {
            if c.is_alphabetic() && map.contains_key(&c) {
                result.push(map[&c]);
            } else {
                result.push(c);
            }
        }

        result
    }
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

    // Count Number of Pairs With Absolute Difference K
    pub fn count_k_difference(nums: Vec<i32>, k: i32) -> i32 {
        let (mut res, mut map) = (0, HashMap::new());
        for num in nums {
            if let Some(v) = map.get(&(num - k)) {
                res += v;
            }

            if let Some(v) = map.get(&(num + k)) {
                res += v;
            }

            *map.entry(num).or_insert(0) += 1;
        }

        res
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
