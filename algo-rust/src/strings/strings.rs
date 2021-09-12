#![allow(dead_code)]
#![allow(unused_macros)]

macro_rules! map(
    { $($key:expr => $value:expr),+ } => {
        {
            let mut m = ::std::collections::HashMap::new();
            $(m.insert($key, $value);)+
            m
        }
     };
);
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

    // leetcode 43
    pub fn multiply(num1: &String, num2: &String) -> String {
        let n1: Vec<u32> = num1
            .chars()
            .rev()
            .map(|x| x.to_digit(10).unwrap())
            .collect();
        let n2: Vec<u32> = num2
            .chars()
            .rev()
            .map(|x| x.to_digit(10).unwrap())
            .collect();

        let n = n1.len();
        let m = n2.len();

        let mut result = vec![0; n + m];
        let mut c = 0;

        for i in 0..n {
            for j in 0..m {
                let sum = result[i + j] + c + n1[i] * n2[j];
                result[i + j] = sum % 10;
                c = sum / 10;
            }
            if c > 0 {
                result[i + m] = c;
                c = 0;
            }
        }

        while result.len() > 1 && result.last() == Some(&0) {
            result.pop();
        }

        result.into_iter().rev().map(|x| x.to_string()).collect()
    }

    // leetcode 168
    pub fn convert_to_title(mut column_number: i32) -> String {
        let base = 'A' as u8;
        let mut result = String::new();

        while column_number > 0 {
            column_number -= 1;

            let c = (column_number % 26) as u8 + base;
            result.push(c as char);

            column_number /= 26;
        }
        result.chars().rev().collect()
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
    #[test]
    fn convert_to_title() {
        assert_eq!("A", Solution::convert_to_title(1));
        assert_eq!("AB", Solution::convert_to_title(28));
        assert_eq!("ZY", Solution::convert_to_title(701));
        assert_eq!("FXSHRXW", Solution::convert_to_title(2147483647));
        assert_eq!("AZ", Solution::convert_to_title(52));
    }

    #[test]
    fn multiply() {
        assert_eq!(
            "10",
            Solution::multiply(&"1".to_string(), &"10".to_string())
        );
        assert_eq!(
            "9801",
            Solution::multiply(&"99".to_string(), &"99".to_string())
        );
        assert_eq!("6", Solution::multiply(&"2".to_string(), &"3".to_string()));
        assert_eq!(
            "56088",
            Solution::multiply(&"123".to_string(), &"456".to_string())
        );
        assert_eq!("0", Solution::multiply(&"1".to_string(), &"0".to_string()));
    }
}
