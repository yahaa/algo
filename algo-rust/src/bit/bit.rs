#![allow(dead_code)]
pub struct Solution {}

impl Solution {
    pub fn max_product(words: Vec<String>) -> i32 {
        if words.len() == 0 {
            return 0;
        }
        let n = words.len();
        let mut bit_map: Vec<u32> = vec![0; n];

        for i in 0..n {
            for c in words[i].chars() {
                bit_map[i] |= 1 << c as u32 - 'a' as u32;
            }
        }

        let mut ans: i32 = 0;

        for i in 0..n {
            for j in i + 1..n {
                if bit_map[i] & bit_map[j] == 0 {
                    ans = std::cmp::max(ans, (words[i].len() * words[j].len()) as i32);
                }
            }
        }
        ans
    }

    pub fn is_power_of_two(n: i32) -> bool {
        if n == 0 {
            return false;
        }

        if n & -n == n {
            return true;
        }
        return false;
    }

    pub fn is_power_of_four(n: i32) -> bool {
        if n < 0 {
            return false;
        }
        if Solution::is_power_of_two(n) {
            let (mut n, mut count) = (n, 0);
            while n > 0 {
                n >>= 1;
                count += 1;
            }

            if count % 2 == 0 {
                return true;
            }
            return false;
        }
        return false;
    }

    pub fn hamming_weight(n: u32) -> i32 {
        let (mut n, mut count) = (n, 0);
        while n > 0 {
            count += 1;
            n = n & (n - 1);
        }
        count
    }

    pub fn total_hamming_distance(nums: Vec<i32>) -> i32 {
        let mut sum = 0;
        let n = nums.len();
        for i in 0..32 {
            let mut count = 0;
            for n in nums.iter() {
                if n & (1 << i) > 0 {
                    count += 1;
                }
            }
            sum += (n - count) * count;
        }
        sum as i32
    }
}
