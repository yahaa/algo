#![allow(dead_code)]

use std::collections::HashMap;
// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

struct Solution {}

impl Solution {
    pub fn get_decimal_value(head: Option<Box<ListNode>>) -> i32 {
        let (mut res, mut head) = (0, head);
        while let Some(cur) = head {
            match cur.val {
                0 => res <<= 1,
                1 => res = (res << 1) + 1,
                _ => {}
            }
            head = cur.next;
        }
        res
    }

    pub fn merge_two_lists(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        match (l1, l2) {
            (None, None) => None,
            (Some(n1), None) => Some(n1),
            (None, Some(n2)) => Some(n2),
            (Some(mut n1), Some(mut n2)) => {
                if n1.val < n2.val {
                    let n = n1.next;
                    n1.next = Solution::merge_two_lists(n, Some(n2));
                    Some(n1)
                } else {
                    let n = n2.next;
                    n2.next = Solution::merge_two_lists(n, Some(n1));
                    Some(n2)
                }
            }
        }
    }

    pub fn merge_two_lists2(
        mut l1: Option<Box<ListNode>>,
        mut l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut head = Box::new(ListNode::new(0));
        let mut tail = &mut head;

        loop {
            match (l1, l2) {
                (Some(mut n1), Some(mut n2)) => {
                    if n1.val < n2.val {
                        l1 = n1.next.take();
                        l2 = Some(n2);
                        tail.next = Some(n1);
                    } else {
                        l1 = Some(n1);
                        l2 = n2.next.take();
                        tail.next = Some(n2);
                    }
                }
                (Some(n1), None) => {
                    tail.next = Some(n1);
                    break;
                }
                (None, Some(n2)) => {
                    tail.next = Some(n2);
                    break;
                }
                (None, None) => break,
            }
            tail = tail.next.as_mut().unwrap();
        }

        head.next
    }

    pub fn merge_two_lists3(
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

    pub fn delete_duplicates(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return None;
        }

        let mut head = head;
        let mut pre = head.as_mut().unwrap();
        while let Some(cur) = pre.next.as_mut() {
            if pre.val == cur.val {
                pre.next = cur.next.take();
            } else {
                pre = pre.next.as_mut().unwrap();
            }
        }
        head
    }

    pub fn delete_duplicates2(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut map: HashMap<i32, i32> = HashMap::new();
        let mut p = head.as_mut();

        while let Some(cur) = p {
            *map.entry(cur.val).or_insert(0) += 1;
            p = cur.next.as_mut();
        }

        for (k, v) in map.iter() {
            println!("key: {},val: {}", k, v);
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

    pub fn remove_elements2(head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
        let mut new_head = Box::new(ListNode::new(0));
        new_head.next = head;

        let mut cur = new_head.as_mut();

        while let Some(next) = cur.next.as_mut() {
            if next.val == val {
                cur.next = next.next.take();
            } else {
                cur = cur.next.as_mut().unwrap();
            }
        }

        new_head.next
    }

    pub fn partition(mut head: Option<Box<ListNode>>, x: i32) -> Option<Box<ListNode>> {
        let mut fake_less_head = Box::new(ListNode::new(0));
        let mut less_tail = &mut fake_less_head;

        let mut fake_big_head = Box::new(ListNode::new(0));
        let mut big_tail = &mut fake_big_head;

        while let Some(mut cur) = head {
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

    pub fn mid_of_linkedlist(head: &Option<Box<ListNode>>) -> &Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }

        let mut slow = head;
        let mut fast = head;
        while fast.as_ref().unwrap().next.is_some()
            && fast.as_ref().unwrap().next.as_ref().unwrap().next.is_some()
        {
            slow = &slow.as_ref().unwrap().next;
            fast = &fast.as_ref().unwrap().next.as_ref().unwrap().next;
        }
        slow
    }

    pub fn has_cycle(head: &Option<Box<ListNode>>) -> bool {
        if head.is_none() {
            return false;
        }

        let (mut slow, mut fast) = (head, head);
        while fast.as_ref().unwrap().next.is_some()
            && fast.as_ref().unwrap().next.as_ref().unwrap().next.is_some()
        {
            slow = &slow.as_ref().unwrap().next;
            fast = &fast.as_ref().unwrap().next.as_ref().unwrap().next;

            if slow == fast {
                return true;
            }
        }

        false
    }

    pub fn sort_list(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() || head.as_ref()?.next.is_none() {
            return head;
        }

        let mut slow = head.as_mut()? as *mut Box<ListNode>;
        let mut fast = head.as_mut()? as *mut Box<ListNode>;

        unsafe {
            while (*fast).next.is_some() && (*fast).next.as_mut()?.next.is_some() {
                slow = (*slow).next.as_mut()?;
                fast = (*fast).next.as_mut()?.next.as_mut()?;
            }
        }

        let right = unsafe { Solution::sort_list((*slow).next.take()) };
        let left = Solution::sort_list(head);

        Solution::merge_two_lists(left, right)
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]

    fn delete_duplicates2() {
        let head = Some(Box::new(ListNode {
            val: 1,
            next: Some(Box::new(ListNode {
                val: 2,
                next: Some(Box::new(ListNode {
                    val: 3,
                    next: Some(Box::new(ListNode {
                        val: 3,
                        next: Some(Box::new(ListNode {
                            val: 4,
                            next: Some(Box::new(ListNode {
                                val: 4,
                                next: Some(Box::new(ListNode { val: 5, next: None })),
                            })),
                        })),
                    })),
                })),
            })),
        }));

        let head = Solution::delete_duplicates2(head);

        println!("{:?}", head);
    }
}
