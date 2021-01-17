/*
 * @lc app=leetcode id=215 lang=rust
 *
 * [215] Kth Largest Element in an Array
 *
 * https://leetcode.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (57.03%)
 * Total Accepted:    756.8K
 * Total Submissions: 1.3M
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * Find the kth largest element in an unsorted array. Note that it is the kth
 * largest element in the sorted order, not the kth distinct element.
 *
 * Example 1:
 *
 *
 * Input: [3,2,1,5,6,4] and k = 2
 * Output: 5
 *
 *
 * Example 2:
 *
 *
 * Input: [3,2,3,1,2,4,5,5,6] and k = 4
 * Output: 4
 *
 * Note:
 * You may assume k is always valid, 1 ≤ k ≤ array's length.
 *
 */

use std::collections::BinaryHeap;
use std::cmp::Reverse;
impl Solution {
    pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
        let mut heap = BinaryHeap::new();

        for n in nums {
            heap.push(Reverse(n));

            if heap.len() > k as usize {
                heap.pop();
            }
        }

        heap.peek().unwrap().0
    }
}
