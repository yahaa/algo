#![allow(dead_code)]

struct Solution {}

impl Solution {
    pub fn gray_code(n: i32) -> Vec<i32> {
        let mut res = vec![0; 1 << n];

        for i in 1..(1 << n) {
            res[i as usize] = i ^ i >> 1;
        }

        res
    }
}
