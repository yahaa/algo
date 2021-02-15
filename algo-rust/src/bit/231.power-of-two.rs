/*
 * @lc app=leetcode id=231 lang=rust
 *
 * [231] Power of Two
 *
 * https://leetcode.com/problems/power-of-two/description/
 *
 * algorithms
 * Easy (43.76%)
 * Total Accepted:    382.9K
 * Total Submissions: 874.8K
 * Testcase Example:  '1'
 *
 * Given an integer n, return true if it is a power of two. Otherwise, return
 * false.
 *
 * An integer n is a power of two, if there exists an integer x such that n ==
 * 2^x.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 1
 * Output: true
 * Explanation: 2^0 = 1
 *
 *
 * Example 2:
 *
 *
 * Input: n = 16
 * Output: true
 * Explanation: 2^4 = 16
 *
 *
 * Example 3:
 *
 *
 * Input: n = 3
 * Output: false
 *
 *
 * Example 4:
 *
 *
 * Input: n = 4
 * Output: true
 *
 *
 * Example 5:
 *
 *
 * Input: n = 5
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * -2^31 <= n <= 2^31 - 1
 *
 *
 */
impl Solution {
    pub fn is_power_of_two(n: i32) -> bool {
        let n = n as i64;

        if n == 0 {
            return false;
        }

        if n & -n == n {
            return true;
        }
        return false;
    }
}
