/*
 * @lc app=leetcode id=342 lang=rust
 *
 * [342] Power of Four
 *
 * https://leetcode.com/problems/power-of-four/description/
 *
 * algorithms
 * Easy (41.44%)
 * Total Accepted:    212.9K
 * Total Submissions: 513.5K
 * Testcase Example:  '16'
 *
 * Given an integer n, return true if it is a power of four. Otherwise, return
 * false.
 *
 * An integer n is a power of four, if there exists an integer x such that n ==
 * 4^x.
 *
 *
 * Example 1:
 * Input: n = 16
 * Output: true
 * Example 2:
 * Input: n = 5
 * Output: false
 * Example 3:
 * Input: n = 1
 * Output: true
 *
 *
 * Constraints:
 *
 *
 * -2^31 <= n <= 2^31 - 1
 *
 *
 *
 * Follow up: Could you solve it without loops/recursion?
 */
impl Solution {
    pub fn is_power_of_two(n: i32) -> bool {
        if n == 0 {
            return false;
        }

        if n & -n == n {
            return true;
        }
        return false;
    }

    pub fn is_power_of_four(n: i32) -> bool {
        if n < 0 {
            return false;
        }

        if Solution::is_power_of_two(n) {
            let (mut n, mut count) = (n, 0);
            while n > 1 {
                n >>= 1;
                count += 1;
            }

            if count % 2 == 0 {
                return true;
            }
            return false;
        }
        return false;
    }
}
