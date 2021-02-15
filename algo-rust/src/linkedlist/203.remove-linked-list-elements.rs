/*
 * @lc app=leetcode id=203 lang=rust
 *
 * [203] Remove Linked List Elements
 *
 * https://leetcode.com/problems/remove-linked-list-elements/description/
 *
 * algorithms
 * Easy (38.91%)
 * Total Accepted:    405.6K
 * Total Submissions: 1M
 * Testcase Example:  '[1,2,6,3,4,5,6]\n6'
 *
 * Remove all elements from a linked list of integers that have value val.
 *
 * Example:
 *
 *
 * Input:  1->2->6->3->4->5->6, val = 6
 * Output: 1->2->3->4->5
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
impl Solution {
    pub fn remove_elements(head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
        let mut new_head = Box::new(ListNode::new(0));
        new_head.next = head;

        let mut pre = new_head.as_mut();

        while let Some(cur) = pre.next.as_mut() {
            if cur.val == val {
                pre.next = cur.next.take();
            } else {
                pre = pre.next.as_mut().unwrap();
            }
        }

        new_head.next
    }
}
