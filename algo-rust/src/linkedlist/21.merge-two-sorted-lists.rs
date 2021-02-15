/*
 * @lc app=leetcode id=21 lang=rust
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (54.74%)
 * Total Accepted:    1.2M
 * Total Submissions: 2.2M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new sorted list. The new
 * list should be made by splicing together the nodes of the first two
 * lists.
 *
 *
 * Example 1:
 *
 *
 * Input: l1 = [1,2,4], l2 = [1,3,4]
 * Output: [1,1,2,3,4,4]
 *
 *
 * Example 2:
 *
 *
 * Input: l1 = [], l2 = []
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: l1 = [], l2 = [0]
 * Output: [0]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in both lists is in the range [0, 50].
 * -100 <= Node.val <= 100
 * Both l1 and l2 are sorted in non-decreasing order.
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
    pub fn merge_two_lists(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let (mut l1, mut l2) = (l1, l2);
        let mut head = Box::new(ListNode::new(0));
        let mut tail = Some(&mut head);

        while let Some(mut cur) = tail {
            match (l1, l2) {
                (Some(mut n1), Some(mut n2)) => {
                    if n1.val < n2.val {
                        l1 = n1.next.take();
                        l2 = Some(n2);
                        cur.next = Some(n1);
                    } else {
                        l1 = Some(n1);
                        l2 = n2.next.take();
                        cur.next = Some(n2);
                    }
                }
                (Some(n1), None) => {
                    cur.next = Some(n1);
                    break;
                }
                (None, Some(n2)) => {
                    cur.next = Some(n2);
                    break;
                }
                (None, None) => {
                    cur.next = None;
                    break;
                }
            }
            tail = cur.next.as_mut();
        }

        head.next
    }
}
