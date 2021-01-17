struct Solution {}

impl Solution {
    pub fn tribonacci(n: i32) -> i32 {
        let (mut f0, mut f1, mut f2, mut f3) = (0, 1, 1, 2);
        let mut n = n;

        match n {
            0 => return 0,
            1 => return 1,
            2 => return 1,
            _ => {}
        }

        while n > 2 {
            f3 = f2 + f1 + f0;
            f0 = f1;
            f1 = f2;
            f2 = f3;
            n -= 1;
        }

        f3
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn tribonacci() {
        assert_eq!(0, Solution::tribonacci(0));
        assert_eq!(1, Solution::tribonacci(1));
        assert_eq!(1, Solution::tribonacci(2));
        assert_eq!(1, Solution::tribonacci(2));
    }
}
