/*
 * @lc app=leetcode id=137 lang=rust
 *
 * [137] Single Number II
 *
 * https://leetcode.com/problems/single-number-ii/description/
 *
 * algorithms
 * Medium (53.14%)
 * Total Accepted:    265.4K
 * Total Submissions: 499.2K
 * Testcase Example:  '[2,2,3,2]'
 *
 * Given a non-empty array of integers, every element appears three times
 * except for one, which appears exactly once. Find that single one.
 *
 * Note:
 *
 * Your algorithm should have a linear runtime complexity. Could you implement
 * it without using extra memory?
 *
 * Example 1:
 *
 *
 * Input: [2,2,3,2]
 * Output: 3
 *
 *
 * Example 2:
 *
 *
 * Input: [0,1,0,1,0,1,99]
 * Output: 99
 *
 */
impl Solution {
    pub fn single_number(nums: Vec<i32>) -> i32 {
        let mut res = 0;
        for i in 0..32 {
            let mut sum = 0;
            for n in nums.iter() {
                sum += (n >> i) & 1;
            }
            res ^= (sum % 3) << i;
        }
        res
    }
}
