/*
 * @lc app=leetcode id=82 lang=rust
 *
 * [82] Remove Duplicates from Sorted List II
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (37.59%)
 * Total Accepted:    277.2K
 * Total Submissions: 736K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * Given a sorted linked list, delete all nodes that have duplicate numbers,
 * leaving only distinct numbers from the original list.
 *
 * Return the linked list sorted as well.
 *
 * Example 1:
 *
 *
 * Input: 1->2->3->3->4->4->5
 * Output: 1->2->5
 *
 *
 * Example 2:
 *
 *
 * Input: 1->1->1->2->3
 * Output: 2->3
 *
 *
 */
// Definition for singly-linked list.
// #[derive(PartialEq, Eq, Clone, Debug)]
// pub struct ListNode {
//   pub val: i32,
//   pub next: Option<Box<ListNode>>
// }
//
// impl ListNode {
//   #[inline]
//   fn new(val: i32) -> Self {
//     ListNode {
//       next: None,
//       val
//     }
//   }
// }
use std::collections::HashMap;
impl Solution {
    pub fn delete_duplicates(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut map: HashMap<i32, i32> = HashMap::new();
        let mut p = head.as_mut();

        while let Some(cur) = p {
            *map.entry(cur.val).or_insert(0) += 1;
            p = cur.next.as_mut();
        }

        let mut new_head = Box::new(ListNode::new(0));
        new_head.next = head;

        let mut pre = new_head.as_mut();

        while let Some(cur) = pre.next.as_mut() {
            if let Some(v) = map.get(&cur.val) {
                if *v > 1 {
                    pre.next = cur.next.take();
                } else {
                    pre = pre.next.as_mut().unwrap();
                }
            } else {
                pre = pre.next.as_mut().unwrap();
            }
        }

        new_head.next
    }
}
