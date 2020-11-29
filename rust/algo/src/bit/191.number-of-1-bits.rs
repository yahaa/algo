struct Solution {}

impl Solution {
    pub fn hammingWeight(n: u32) -> i32 {
        let (mut n, mut count) = (n, 0);
        while n > 0 {
            count += 1;
            n = n & (n - 1);
        }
        count
    }
}
