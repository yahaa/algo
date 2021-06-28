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
}
