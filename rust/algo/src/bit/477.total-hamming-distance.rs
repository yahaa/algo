/*
 * @lc app=leetcode id=477 lang=rust
 *
 * [477] Total Hamming Distance
 *
 * https://leetcode.com/problems/total-hamming-distance/description/
 *
 * algorithms
 * Medium (50.50%)
 * Total Accepted:    67.8K
 * Total Submissions: 134.2K
 * Testcase Example:  '[4,14,2]'
 *
 * The Hamming distance between two integers is the number of positions at
 * which the corresponding bits are different.
 *
 * Now your job is to find the total Hamming distance between all pairs of the
 * given numbers.
 *
 *
 * Example:
 *
 * Input: 4, 14, 2
 *
 * Output: 6
 *
 * Explanation: In binary representation, the 4 is 0100, 14 is 1110, and 2 is
 * 0010 (just
 * showing the four bits relevant in this case). So the answer will be:
 * HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2
 * + 2 + 2 = 6.
 *
 *
 *
 * Note:
 *
 * Elements of the given array are in the range of 0  to 10^9
 * Length of the array will not exceed 10^4.
 *
 *
 */
impl Solution {
    pub fn total_hamming_distance(nums: Vec<i32>) -> i32 {
        let mut sum = 0;
        let n = nums.len();
        for i in 0..32 {
            let mut count = 0;
            for m in nums.iter() {
                if m & (1 << i) > 0 {
                    count += 1;
                }
            }
            sum += (n - count) * count;
        }
        sum as i32
    }
}
