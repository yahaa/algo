#![allow(dead_code)]

use std::cmp::{max, min};

struct Solution {}

impl Solution {
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
}

#[cfg(test)]
mod test {
    use super::*;

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
}
