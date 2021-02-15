/*
 * @lc app=leetcode id=86 lang=rust
 *
 * [86] Partition List
 *
 * https://leetcode.com/problems/partition-list/description/
 *
 * algorithms
 * Medium (42.51%)
 * Total Accepted:    235.4K
 * Total Submissions: 553.7K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * Given a linked list and a value x, partition it such that all nodes less
 * than x come before nodes greater than or equal to x.
 *
 * You should preserve the original relative order of the nodes in each of the
 * two partitions.
 *
 * Example:
 *
 *
 * Input: head = 1->4->3->2->5->2, x = 3
 * Output: 1->2->2->4->3->5
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
    pub fn partition(mut head: Option<Box<ListNode>>, x: i32) -> Option<Box<ListNode>> {
        let mut fake_less_head = Box::new(ListNode::new(0));
        let mut less_tail = &mut fake_less_head;

        let mut fake_big_head = Box::new(ListNode::new(0));
        let mut big_tail = &mut fake_big_head;

        while let Some(mut cur) = head{
            head = cur.next.take();
            if cur.val < x {
                less_tail.next = Some(cur);
                less_tail = less_tail.next.as_mut().unwrap();
            } else {
                big_tail.next = Some(cur);
                big_tail = big_tail.next.as_mut().unwrap();
            }
        }

        less_tail.next = fake_big_head.next;

        fake_less_head.next
    }
}
