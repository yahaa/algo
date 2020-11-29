/*
 * @lc app=leetcode id=201 lang=rust
 *
 * [201] Bitwise AND of Numbers Range
 *
 * https://leetcode.com/problems/bitwise-and-of-numbers-range/description/
 *
 * algorithms
 * Medium (39.54%)
 * Total Accepted:    163K
 * Total Submissions: 412.3K
 * Testcase Example:  '5\n7'
 *
 * Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND
 * of all numbers in this range, inclusive.
 *
 * Example 1:
 *
 *
 * Input: [5,7]
 * Output: 4
 *
 *
 * Example 2:
 *
 *
 * Input: [0,1]
 * Output: 0
 */
impl Solution {
    pub fn range_bitwise_and(m: i32, n: i32) -> i32 {
        let mut c = 0;
        let (mut n, mut m) = (n, m);
        while m < n {
            n = n >> 1;
            m = m >> 1;
            c += 1;
        }

        return m << c;
    }
}
