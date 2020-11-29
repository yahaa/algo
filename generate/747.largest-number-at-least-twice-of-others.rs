/*
 * @lc app=leetcode id=747 lang=rust
 *
 * [747] Largest Number At Least Twice of Others
 *
 * https://leetcode.com/problems/largest-number-at-least-twice-of-others/description/
 *
 * algorithms
 * Easy (42.62%)
 * Total Accepted:    100.5K
 * Total Submissions: 235.9K
 * Testcase Example:  '[0,0,0,1]'
 *
 * In a given integer array nums, there is always exactly one largest element.
 *
 * Find whether the largest element in the array is at least twice as much as
 * every other number in the array.
 *
 * If it is, return the index of the largest element, otherwise return -1.
 *
 * Example 1:
 *
 *
 * Input: nums = [3, 6, 1, 0]
 * Output: 1
 * Explanation: 6 is the largest integer, and for every other number in the
 * array x,
 * 6 is more than twice as big as x.  The index of value 6 is 1, so we return
 * 1.
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1, 2, 3, 4]
 * Output: -1
 * Explanation: 4 isn't at least as big as twice the value of 3, so we return
 * -1.
 *
 *
 *
 *
 * Note:
 *
 *
 * nums will have a length in the range [1, 50].
 * Every nums[i] will be an integer in the range [0, 99].
 *
 *
 *
 *
 */
impl Solution {
    pub fn dominant_index(nums: Vec<i32>) -> i32 {
        if nums.len() < 2 {
            return (nums.len() - 1) as i32;
        }

        let (mut max_i, mut max_v, mut max_v_1) = (0 as usize, &std::i32::MIN, &std::i32::MIN);
        for (i, v) in nums.iter().enumerate() {
            if v > max_v {
                max_i = i;
                max_v_1 = max_v;
                max_v = v;
            } else if v > max_v_1 {
                max_v_1 = v;
            }
        }

        if *max_v_1 == 0 {
            return max_i as i32;
        }

        if max_v / max_v_1 >= 2 {
            return max_i as i32;
        }
        return -1;
    }
}
