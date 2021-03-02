#![allow(dead_code)]

struct Solution {}

impl Solution {
    // leetcode 1347
    pub fn min_steps(s: String, t: String) -> i32 {
        let mut count = vec![0; 26];

        let ss: Vec<char> = s.chars().collect();
        let tt: Vec<char> = t.chars().collect();
        let base = 'a' as usize;

        for i in 0..ss.len() {
            count[ss[i] as usize - base] += 1;
            count[tt[i] as usize - base] -= 1;
        }

        count.into_iter().filter(|x| *x > 0).sum()
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn min_steps() {
        assert_eq!(0, Solution::min_steps("abc".to_string(), "abc".to_string()));
        assert_eq!(1, Solution::min_steps("bab".to_string(), "aba".to_string()));
        assert_eq!(
            5,
            Solution::min_steps("leetcode".to_string(), "practice".to_string())
        );
    }
}
