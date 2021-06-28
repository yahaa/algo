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
