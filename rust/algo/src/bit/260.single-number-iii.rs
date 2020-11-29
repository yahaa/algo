/*
 * @lc app=leetcode id=260 lang=rust
 *
 * [260] Single Number III
 *
 * https://leetcode.com/problems/single-number-iii/description/
 *
 * algorithms
 * Medium (65.05%)
 * Total Accepted:    177.2K
 * Total Submissions: 272.3K
 * Testcase Example:  '[1,2,1,3,2,5]'
 *
 * Given an integer array nums, in which exactly two elements appear only once
 * and all the other elements appear exactly twice. Find the two elements that
 * appear only once. You can return the answer in any order.
 *
 * Follow up: Your algorithm should run in linear runtime complexity. Could you
 * implement it using only constant space complexity?
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,1,3,2,5]
 * Output: [3,5]
 * Explanation:  [5, 3] is also a valid answer.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [-1,0]
 * Output: [-1,0]
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [0,1]
 * Output: [1,0]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 30000
 * Each integer in nums will appear twice, only two integers will appear once.
 *
 *
 */

struct Solution {}

impl Solution {
    pub fn single_number(nums: Vec<i32>) -> Vec<i32> {
        let mut t = 0;
        for n in nums.iter() {
            t ^= n;
        }

        let f = t & -t;

        let (mut a, mut b) = (0, 0);
        for n in nums.iter() {
            if (n & f) == 0 {
                a ^= n;
            } else {
                b ^= n;
            }
        }

        vec![a, b]
    }
}
/*
9 = 1001
原码：1001
反码：0110
补码：0111

-9 = 0111

9 &-9= 0001


10= 1010
1010
0101
0110
*/
