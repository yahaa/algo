#![allow(dead_code)]
struct Solution {}

impl Solution {
    // leetcode 1920
    pub fn build_array(nums: Vec<i32>) -> Vec<i32> {
        nums.iter().map(|n| nums[*n as usize]).collect()
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
}
