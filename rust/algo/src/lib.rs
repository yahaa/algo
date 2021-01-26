mod bit;
mod design;
mod heap;
mod linkedlist;
mod recursive;
mod search;

use std::cmp;
use std::collections::HashMap;
use std::collections::LinkedList;

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
    // lc 26
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

    // lc 27
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

    // lc 28
    pub fn str_str1(haystack: String, needle: String) -> i32 {
        match haystack.find(&needle) {
            None => -1,
            Some(pos) => pos as i32,
        }
    }

    // lc 28
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

    // lc 125
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

    // lc 125
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

    // lc 88
    pub fn merge1(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
        let (mut k, mut j, mut index) = (m as usize, n as usize, (n + m) as usize);

        if index <= 0 {
            return;
        }

        while k > 0 && j > 0 {
            if nums1[k - 1] > nums2[j - 1] {
                nums1[index - 1] = nums1[k - 1];
                k -= 1;
            } else {
                nums1[index - 1] = nums2[j - 1];
                j -= 1;
            }
            index -= 1;
        }

        while j > 0 {
            nums1[index - 1] = nums2[j - 1];
            j -= 1;
            index -= 1;
        }
    }

    pub fn is_palindrome2(head: Option<Box<ListNode>>) -> bool {
        let mut arr: Vec<i32> = Vec::new();
        let mut cur = head;
        while let Some(n) = cur {
            arr.push(n.val);
            cur = n.next;
        }

        let (mut i, mut j) = (0 as usize, arr.len() as usize);
        while i < j {
            if arr[i] != arr[j - 1] {
                return false;
            }
            i += 1;
            j -= 1;
        }
        true
    }

    pub fn reverse_vowels(s: String) -> String {
        if s.len() == 0 {
            return s;
        }

        let mut chars: Vec<char> = s.chars().collect();

        let mut mp: HashMap<char, bool> = HashMap::new();
        mp.insert('a', true);
        mp.insert('e', true);
        mp.insert('i', true);
        mp.insert('o', true);
        mp.insert('u', true);
        mp.insert('A', true);
        mp.insert('E', true);
        mp.insert('I', true);
        mp.insert('O', true);
        mp.insert('U', true);

        let (mut i, mut j) = (0, s.len() - 1);

        while i < j {
            while i < j {
                match mp.get(&chars[i]) {
                    None => i += 1,
                    Some(_) => break,
                }
            }
            while i < j {
                match mp.get(&chars[j]) {
                    None => j -= 1,
                    Some(_) => break,
                }
            }

            if i != j {
                chars.swap(i, j);
                i += 1;
                j -= 1;
            }
        }

        return chars.iter().collect();
    }

    // lc 884
    pub fn backspace_compare(s: String, t: String) -> bool {
        let s = Solution::build(s);
        let t = Solution::build(t);
        return s == t;
    }

    pub fn build(s: String) -> String {
        let mut stack: LinkedList<char> = LinkedList::new();

        for c in s.chars() {
            if c != '#' {
                stack.push_back(c);
            } else {
                stack.pop_back();
            }
        }

        return stack.into_iter().collect();
    }

    pub fn cmp(a: i32, b: i32) -> bool {
        a < b
    }

    // lc 56
    pub fn merge(mut intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        if intervals.len() == 0 {
            return intervals;
        }
        intervals.sort_by(|a, b| a[0].cmp(&b[0]));
        let mut i: usize = 0;
        while i < intervals.len() - 1 {
            if intervals[i][1] >= intervals[i + 1][0] {
                if intervals[i][1] < intervals[i + 1][1] {
                    intervals[i][1] = intervals[i + 1][1];
                }
                intervals.remove(i + 1);
            } else {
                i += 1;
            }
        }
        intervals
    }

    // lc 442
    pub fn find_duplicates(mut nums: Vec<i32>) -> Vec<i32> {
        let mut res: Vec<i32> = Vec::new();
        for i in 0..nums.len() {
            let index = nums[i].abs() - 1;
            if nums[index as usize] < 0 {
                res.push(index + 1);
            }
            nums[index as usize] = -nums[index as usize]
        }
        res
    }

    // lc 1395
    pub fn num_teams(rating: Vec<i32>) -> i32 {
        let mut res = 0;
        for i in 0..rating.len() {
            for j in i + 1..rating.len() {
                for k in j + 1..rating.len() {
                    if rating[i] < rating[j] && rating[j] < rating[k]
                        || rating[i] > rating[j] && rating[j] > rating[k]
                    {
                        res += 1;
                    }
                }
            }
        }
        res
    }

    // lc 38
    pub fn count_and_say(n: i32) -> String {
        if n == 1 {
            return "1".to_string();
        }

        let sub = Self::count_and_say(n - 1).chars().collect::<Vec<char>>();
        let (mut cur, mut count, mut s) = (&sub[0], 1, Vec::new());

        for c in &sub[1..] {
            if cur != c {
                s.push(format!("{}{}", count, cur));
                count = 1;
            } else {
                count += 1;
            }
            cur = c;
        }
        s.push(format!("{}{}", count, cur));
        s.join("")
    }

    // leetcoee 383
    pub fn can_construct(ransom_note: String, magazine: String) -> bool {
        let mut mp: HashMap<char, i32> = HashMap::new();
        for c in magazine.chars() {
            *mp.entry(c).or_insert(0) += 1;
        }

        for c in ransom_note.chars() {
            match mp.get_mut(&c) {
                Some(count) => {
                    *count -= 1;
                    if *count < 0 {
                        return false;
                    }
                }
                None => return false,
            }
        }
        true
    }

    // lc 387
    pub fn first_uniq_char(s: String) -> i32 {
        let mut mp: HashMap<char, i32> = HashMap::new();

        for c in s.chars() {
            *mp.entry(c).or_insert(0) += 1;
        }

        let mut res = 0;
        for c in s.chars() {
            match mp.get(&c) {
                Some(count) => {
                    if *count == 1 {
                        return res;
                    }
                    res += 1;
                }
                None => return 0,
            }
        }
        return -1;
    }

    // lc 415
    pub fn add_strings(num1: String, num2: String) -> String {
        let min = cmp::min(num1.len(), num2.len());
        let chs1 = num1.chars().rev().collect::<Vec<char>>();
        let chs2 = num2.chars().rev().collect::<Vec<char>>();

        let mut res = Vec::new();
        let mut carry = 0;
        for i in 0..min {
            let a = chs1[i].to_digit(10).unwrap() as i32;
            let b = chs2[i].to_digit(10).unwrap() as i32;
            let v = (a + b + carry) % 10;
            carry = (a + b + carry) / 10;
            res.push(v.to_string());
        }

        for i in min..chs1.len() {
            let a = chs1[i].to_digit(10).unwrap() as i32;
            let v = (a + carry) % 10;
            carry = (a + carry) / 10;
            res.push(v.to_string());
        }

        for i in min..chs2.len() {
            let b = chs2[i].to_digit(10).unwrap() as i32;
            let v = (b + carry) % 10;
            carry = (b + carry) / 10;
            res.push(v.to_string());
        }

        if carry > 0 {
            res.push(carry.to_string());
        }

        res.join("").chars().rev().collect()
    }

    // lc 434
    pub fn count_segments(s: String) -> i32 {
        s.split_whitespace().count() as i32
    }

    // lc 520
    pub fn detect_capital_use(word: String) -> bool {
        word.chars().all(|a| a.is_uppercase()) || word.chars().skip(0).all(|a| a.is_lowercase())
    }

    pub fn min_window(s: String, t: String) -> String {
        let mut mps: HashMap<char, i32> = HashMap::new();
        let mut mpt: HashMap<char, i32> = HashMap::new();

        for c in s.chars() {
            *mps.entry(c).or_insert(0) += 1;
        }

        for c in t.chars() {
            *mpt.entry(c).or_insert(0) += 1;
        }

        let (mut i, mut j, mut found) = (0, s.len() - 1, false);

        while i < j && !found {
            let c = s.chars().nth(i).unwrap();
            match mpt.get(&c) {
                Some(count) => {
                    if let Some(t) = mps.get(&c) {
                        if t > count {
                            i += 1;
                            *mps.entry(c).or_insert(0) -= 1;
                        }
                    } else {
                        return String::from("");
                    }
                }
                None => i += 1,
            }
        }

        return s;
    }

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

    // lc 136
    pub fn single_number(nums: Vec<i32>) -> i32 {
        if nums.len() <= 0 {
            return 0;
        }

        let mut ans = nums[0];

        for i in 1..nums.len() {
            ans ^= nums[i];
        }
        ans
    }

    pub fn single_number2(nums: Vec<i32>) -> i32 {
        let mut res = 0;
        for i in 0..32 {
            let mut sum = 0;
            for n in nums.iter() {
                sum += (n >> i) & 1;
            }
            res ^= (sum % 3) << i;
        }
        res
    }

    pub fn find_repeated_dna_sequences(s: String) -> Vec<String> {
        if s.len() < 10 {
            return vec![];
        }
        let mut rep: HashMap<&str, i32> = HashMap::new();

        for i in 0..=s.len() - 10 {
            let sb = &s[i..i + 10];
            *rep.entry(&sb).or_insert(0) += 1;
        }

        let mut res: Vec<String> = Vec::new();

        for (k, v) in rep.iter() {
            if *v > 1 {
                res.push(k.to_string());
            }
        }
        res
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
    }

    #[test]
    fn reverse_vowels() {
        let input = String::from("a.");
        let want = String::from("a.");

        let res = Solution::reverse_vowels(input);
        assert_eq!(res, want);
    }

    #[test]
    fn merge() {
        let mut arrays: Vec<Vec<i32>> = vec![vec![2, 6], vec![1, 4], vec![3, 5], vec![-1, 9]];
        arrays = Solution::merge(arrays);
        assert_eq!(vec![vec![-1, 9]], arrays);
    }

    #[test]
    fn find_duplicates() {
        let nums = vec![4, 3, 2, 7, 8, 2, 3, 1];
        let res = Solution::find_duplicates(nums);
        assert_eq!(vec![2, 3], res);
    }

    #[test]
    fn num_teams() {
        let nums = vec![2, 5, 3, 4, 1];
        let res = Solution::num_teams(nums);
        assert_eq!(res, 3);
    }

    #[test]
    fn count_and_say() {
        let res = Solution::count_and_say(1);
        assert_eq!("1", res);
        let res = Solution::count_and_say(2);
        assert_eq!("11", res);
        let res = Solution::count_and_say(4);
        assert_eq!("1211", res);
    }

    #[test]
    fn can_construct() {
        assert_eq!(
            Solution::can_construct(String::from("a"), String::from("b")),
            false
        );

        assert_eq!(
            Solution::can_construct(String::from("aa"), String::from("ab")),
            false
        );

        assert_eq!(
            Solution::can_construct(String::from("aa"), String::from("aab")),
            true
        );
    }

    #[test]
    fn first_uniq_char() {
        // assert_eq!(0, Solution::first_uniq_char(String::from("lc")));
        // assert_eq!(2, Solution::first_uniq_char(String::from("lovelc")));
    }

    #[test]
    fn add_strings() {
        // assert_eq!(
        //     "11",
        //     Solution::add_strings(String::from("10"), String::from("1"))
        // );

        // assert_eq!(
        //     "2",
        //     Solution::add_strings(String::from("1"), String::from("1"))
        // );
        assert_eq!(
            "602",
            Solution::add_strings(String::from("584"), String::from("18"))
        );
    }

    #[test]
    fn dominant_index() {
        let nums = vec![3, 6, 1, 0];
        let want = 1;
    }
}
