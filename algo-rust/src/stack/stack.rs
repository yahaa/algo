#![allow(dead_code)]

use std::cmp::{max, min};

struct Solution {}

impl Solution {
    // 32. Longest Valid Parentheses
    pub fn longest_valid_parentheses(s: String) -> i32 {
        let (mut stack, mut ans) = (Vec::new(), 0);
        stack.push(-1);

        for (i, c) in s.chars().enumerate() {
            if c == '(' {
                stack.push(i as i32);
            } else {
                stack.pop();

                if stack.is_empty() {
                    stack.push(i as i32);
                } else {
                    ans = max(ans, i as i32 - stack.last().unwrap());
                }
            }
        }
        ans
    }
    // leetcode 921
    pub fn min_add_to_make_valid(s: String) -> i32 {
        let (mut stack, mut iter) = (Vec::new(), s.chars());

        while let Some(c) = iter.next() {
            if c == '(' {
                stack.push(c);
            } else {
                match stack.last() {
                    Some(l) => {
                        if l == &'(' {
                            stack.pop();
                        } else {
                            stack.push(c);
                        }
                    }
                    None => stack.push(c),
                }
            }
        }

        stack.len() as i32
    }

    // leetcode 1130
    pub fn mct_from_leaf_values(arr: Vec<i32>) -> i32 {
        let (mut stack, mut mct) = (Vec::new(), 0);
        stack.push(i32::MAX);

        for i in 0..arr.len() {
            while &arr[i] >= stack.last().unwrap() {
                mct += stack.pop().unwrap() * min(stack.last().unwrap(), &arr[i]);
            }

            stack.push(arr[i]);
        }

        while stack.len() > 2 {
            mct += stack.pop().unwrap() * stack.last().unwrap();
        }

        mct
    }

    //leetcode 856
    pub fn score_of_parentheses(s: String) -> i32 {
        let (mut res, mut stack) = (0, Vec::new());
        let mut iter = s.chars();

        while let Some(c) = iter.next() {
            if c == '(' {
                stack.push(res);
                res = 0;
            } else {
                res = stack.pop().unwrap() + max(res << 1, 1);
            }
        }

        res
    }

    // leetcode 739
    pub fn daily_temperatures(t: Vec<i32>) -> Vec<i32> {
        let (mut stack, mut res) = (Vec::new(), vec![0; t.len()]);

        stack.push((t.len(), i32::MAX));

        for i in (0..t.len()).rev() {
            let mut tmp = 0;

            while t[i] >= stack.last().unwrap().1 {
                let top = stack.pop().unwrap();

                tmp += res[top.0];
            }

            if stack.last().unwrap().1 < i32::MAX {
                tmp += 1;
                res[i] = tmp;
            }

            stack.push((i, t[i]));
        }

        res
    }

    // leetcode 946
    pub fn validate_stack_sequences(pushed: Vec<i32>, popped: Vec<i32>) -> bool {
        let mut stack = Vec::new();
        let mut index = 0;

        for i in 0..pushed.len() {
            stack.push(pushed[i]);
            while !stack.is_empty() && *stack.last().unwrap() == popped[index] {
                index += 1;
                stack.pop();
            }
        }

        index == popped.len()
    }

    // leetcode 1249
    pub fn min_remove_to_make_valid(s: String) -> String {
        let mut stack = Vec::new();
        let mut chars: Vec<char> = s.chars().collect();

        for i in 0..chars.len() {
            if chars[i] == '(' {
                stack.push(i);
            } else if chars[i] == ')' {
                if !stack.is_empty() && chars[*stack.last().unwrap()] == '(' {
                    stack.pop();
                } else {
                    stack.push(i);
                }
            }
        }

        while let Some(i) = stack.pop() {
            chars.remove(i);
        }

        chars.iter().collect()
    }

    // leetcode 1190
    pub fn reverse_parentheses(s: String) -> String {
        let mut chars: Vec<char> = s.chars().collect();
        let mut stack = Vec::new();

        for i in 0..chars.len() {
            if chars[i] == '(' {
                stack.push(i);
            } else if chars[i] == ')' {
                if let Some(s) = stack.pop() {
                    (&mut chars[s + 1..i]).reverse();
                }
            }
        }

        chars.iter().filter(|x| **x != '(' && **x != ')').collect()
    }

    // leetcode 503
    pub fn next_greater_elements(nums: Vec<i32>) -> Vec<i32> {
        let mut res = vec![-1; nums.len()];
        let mut stack: Vec<usize> = Vec::new();

        for i in 0..2 * nums.len() {
            while !stack.is_empty() && nums[*stack.last().unwrap()] < nums[i % nums.len()] {
                let index = stack.pop().unwrap();
                res[index] = nums[i % nums.len()];
            }

            stack.push(i % nums.len());
        }

        res
    }

    // leetcode 1209
    pub fn remove_duplicates(s: String, k: i32) -> String {
        let mut stack: Vec<(char, i32)> = Vec::new();
        let chars: Vec<char> = s.chars().collect();

        for i in 0..chars.len() {
            let mut count = 1;

            if !stack.is_empty() && stack.last().unwrap().0 == chars[i] {
                count += stack.last().unwrap().1;
            }

            stack.push((chars[i], count));

            while !stack.is_empty() && stack.last().unwrap().1 == k {
                let mut t = k;
                while t > 0 && !stack.is_empty() {
                    stack.pop();
                    t -= 1;
                }
            }
        }

        stack.iter().map(|item| item.0).collect()
    }

    // leetcode 1047
    pub fn remove_duplicates_1(s: String) -> String {
        let mut stack: Vec<char> = Vec::new();

        for c in s.chars() {
            if stack.is_empty() {
                stack.push(c);
                continue;
            }

            if stack.last().unwrap() == &c {
                stack.pop();
                continue;
            }

            stack.push(c);
        }

        stack.iter().map(|item| item).collect()
    }

    // leetcode 1003
    pub fn is_valid(s: String) -> bool {
        let mut s = s;
        while s.contains("abc") {
            s = s.replace("abc", "");
        }

        s.len() == 0
    }

    // leetcode 1614
    pub fn max_depth(s: String) -> i32 {
        let (mut stack, mut ans) = (Vec::new(), 0);

        for c in s.chars() {
            if c == '(' {
                stack.push(c);
            } else if c == ')' {
                ans = std::cmp::max(ans, stack.len());
                stack.pop();
            }
        }

        ans as i32
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn is_valid() {
        assert_eq!(true, Solution::is_valid("aabcbc".to_string()));
    }

    #[test]
    fn max_depth() {
        assert_eq!(3, Solution::max_depth("(1+(2*3)+((8)/4))+1".to_string()));
        assert_eq!(3, Solution::max_depth("(1)+((2))+(((3)))".to_string()));
        assert_eq!(1, Solution::max_depth("1+(2*3)/(2-1)".to_string()));
    }

    #[test]
    fn remove_duplicates_1() {
        assert_eq!("ca", Solution::remove_duplicates_1("abbaca".to_string()));
        assert_eq!("a", Solution::remove_duplicates_1("aaa".to_string()));
        assert_eq!("", Solution::remove_duplicates_1("aaaa".to_string()));
    }

    #[test]
    fn remove_duplicates() {
        assert_eq!("abcd", Solution::remove_duplicates("abcd".to_string(), 2));
        assert_eq!(
            "aa",
            Solution::remove_duplicates("deeedbbcccbdaa".to_string(), 3)
        );

        assert_eq!(
            "ps",
            Solution::remove_duplicates("pbbcggttciiippooaais".to_string(), 2)
        );
    }

    #[test]
    fn next_greater_elements() {
        let nums = vec![1, 2, 1];
        assert_eq!(vec![2, -1, 2], Solution::next_greater_elements(nums));
    }
    #[test]
    fn reverse_parentheses() {
        assert_eq!("dcba", Solution::reverse_parentheses("(abcd)".to_string()));
        assert_eq!(
            "iloveu",
            Solution::reverse_parentheses("(u(love)i)".to_string())
        );
        assert_eq!(
            "leetcode",
            Solution::reverse_parentheses("(ed(et(oc))el)".to_string())
        );
        assert_eq!(
            "apmnolkjihgfedcbq",
            Solution::reverse_parentheses("a(bcdefghijkl(mno)p)q".to_string())
        );
    }

    #[test]
    fn min_remove_to_make_valid() {
        assert_eq!(
            "lee(t(c)o)de",
            Solution::min_remove_to_make_valid("lee(t(c)o)de)".to_string())
        )
    }
    #[test]
    fn validate_stack_sequences() {
        let pushed = vec![1, 2, 3, 4, 5];
        let poped = vec![4, 5, 3, 2, 1];
        assert_eq!(true, Solution::validate_stack_sequences(pushed, poped));

        let pushed = vec![1, 2, 3, 4, 5];
        let poped = vec![4, 3, 5, 1, 2];
        assert_eq!(false, Solution::validate_stack_sequences(pushed, poped));
    }

    #[test]
    fn mmin_add_to_make_valid() {
        assert_eq!(1, Solution::min_add_to_make_valid("())".to_string()));
        assert_eq!(3, Solution::min_add_to_make_valid("(((".to_string()));
    }

    #[test]
    fn mct_from_leaf_values() {
        let array = vec![6, 2, 4];
        assert_eq!(32, Solution::mct_from_leaf_values(array));
    }

    #[test]
    fn score_of_parentheses() {
        assert_eq!(1, Solution::score_of_parentheses("()".to_string()));
        assert_eq!(2, Solution::score_of_parentheses("(())".to_string()));
        assert_eq!(4, Solution::score_of_parentheses("((()))".to_string()));
        assert_eq!(5, Solution::score_of_parentheses("((()))()".to_string()));
    }

    #[test]
    fn daily_temperatures() {
        let t = vec![73, 74, 75, 71, 69, 72, 76, 73];
        let res = vec![1, 1, 4, 2, 1, 1, 0, 0];

        assert_eq!(res, Solution::daily_temperatures(t));

        let t = vec![55, 38, 53, 81, 61, 93, 97, 32, 43, 78];
        let res = vec![3, 1, 1, 2, 1, 1, 0, 1, 1, 0];
        assert_eq!(res, Solution::daily_temperatures(t));
    }

    #[test]
    fn longest_valid_parentheses() {
        assert_eq!(2, Solution::longest_valid_parentheses("(()".to_string()));
        assert_eq!(4, Solution::longest_valid_parentheses("()()".to_string()));
        assert_eq!(6, Solution::longest_valid_parentheses("()(())".to_string()));
        assert_eq!(8, Solution::longest_valid_parentheses("()(())()".to_string()));
    }
}
