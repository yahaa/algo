pub struct Solution {}

// #[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }

    pub fn last(&mut self) -> &mut Self {
        if let Some(ref mut box_node) = self.next {
            return box_node.last();
        }

        return self;
    }

    pub fn append(&mut self, val: i32) {
        let _node = ListNode::new(val);
        self.last().next = Some(Box::new(_node));
    }
}

impl Solution {
    // leetcode 26
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        if nums.len() == 0 {
            return 0;
        }
        let mut i = 0;

        for j in 1..nums.len() {
            if nums[i] != nums[j] {
                i += 1;
                nums[i] = nums[j];
            }
        }

        i as i32 + 1
    }

    // leetcode 27
    pub fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
        let mut i = 0;
        for j in 0..nums.len() {
            if nums[j] != val {
                nums[i] = nums[j];
                i += 1;
            }
        }
        i as i32
    }

    // leetcode 28
    pub fn str_str1(haystack: String, needle: String) -> i32 {
        match haystack.find(&needle) {
            None => -1,
            Some(pos) => pos as i32,
        }
    }

    // leetcode 28
    pub fn str_str(haystack: String, needle: String) -> i32 {
        let (hlen, nlen) = (haystack.len(), needle.len());

        if nlen <= 0 {
            return 0;
        } else if hlen < nlen {
            return -1;
        }

        for i in 0..=(hlen - nlen) {
            if haystack[i..i + nlen] == needle {
                return i as i32;
            }
        }
        -1
    }

    // leetcode 125
    pub fn is_palindrome1(s: String) -> bool {
        let chars: Vec<char> = s
            .to_lowercase()
            .chars()
            .filter(|c| c.is_alphanumeric())
            .collect();

        if chars.len() == 0 {
            return true;
        }

        let (mut i, mut j) = (0, chars.len() - 1);
        while i < j {
            let (left, right) = (chars.get(i), chars.get(j));
            if left == right {
                i += 1;
                j -= 1;
            } else {
                return false;
            }
        }
        return true;
    }

    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut pre = None;
        let mut cur = head;

        while let Some(mut _node) = cur {
            cur = _node.next.take();
            _node.next = pre;
            pre = Some(_node);
        }

        return pre;
    }

    pub fn remove_nth_from_end1(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        let mut head = head;
        let (mut cur, mut count) = (head.as_mut(), 0);
        while let Some(mut _node) = cur {
            count += 1;
            cur = _node.next.as_mut();
        }

        let mut head = Some(Box::new(ListNode { val: 0, next: head }));
        let (mut cur, mut count) = (head.as_mut(), count - n);
        while let Some(mut _node) = cur {
            if count == 0 {
                if let Some(mut next) = _node.next.take() {
                    _node.next = next.next.take();
                    break;
                }
            }
            count -= 1;
            cur = _node.next.as_mut();
        }

        return head.unwrap().next;
    }

    pub fn remove_nth_from_end(mut head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        let mut l: *const Option<Box<ListNode>> = &head;
        let (mut d, mut count) = (head.as_mut().unwrap(), 0);
        while let Some(sl) = unsafe { &*l } {
            l = &sl.next;
            if count > n {
                d = d.next.as_mut().unwrap()
            }
            count += 1;
        }
        if count == n {
            return head.unwrap().next;
        }
        d.next = d.next.take().unwrap().next;
        head
    }

    // leetcode 125
    pub fn is_palindrome(s: String) -> bool {
        let s: String = s
            .to_lowercase()
            .chars()
            .filter(|c| c.is_alphanumeric())
            .collect();
        s.chars()
            .zip(s.chars().rev())
            .fold(true, |acc, (a, b)| acc & (a == b))
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn rm() {
        let mut nums: Vec<i32> = vec![1, 2, 2, 3];
        let res = Solution::remove_duplicates(&mut nums);
        assert_eq!(res, 3);

        nums = vec![1, 2, 3, 3];

        let res = Solution::remove_duplicates(&mut nums);
        assert_eq!(res, 3);
    }

    #[test]
    fn rm_el() {
        let mut nums = vec![2, 3, 3, 4];
        let (target, want) = (3, 2);
        assert_eq!(Solution::remove_element(&mut nums, target), want)
    }

    #[test]
    fn str() {
        let h = String::from("abcabcabcd");
        let n = String::from("abcd");

        assert_eq!(Solution::str_str(h, n), 6)
    }
    #[test]
    fn is_palindrome() {
        let s = String::from("A man, a plan, a canal: Panama");
        assert_eq!(Solution::is_palindrome1(s), true);

        let s = String::from("A man, a plan, a canal: Panamaggggg");
        assert_eq!(Solution::is_palindrome1(s), false);

        let s = String::from("");
        assert_eq!(Solution::is_palindrome1(s), true);

        let a = String::from("ab");
        let b = String::from("bbb");

        let mut c = a.chars().zip(b.chars()).fold(init: B, mut f: F)

        while let Some(t) = c.next() {
            print!("{:?}", t);
        }
    }
}
