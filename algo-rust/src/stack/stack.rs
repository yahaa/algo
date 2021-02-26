#![allow(dead_code)]

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
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn mmin_add_to_make_valid() {
        assert_eq!(1, Solution::min_add_to_make_valid("())".to_string()));
        assert_eq!(3, Solution::min_add_to_make_valid("(((".to_string()));
    }
}
