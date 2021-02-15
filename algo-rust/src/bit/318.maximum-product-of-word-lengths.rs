/*
 * @lc app=leetcode id=318 lang=rust
 *
 * [318] Maximum Product of Word Lengths
 *
 * https://leetcode.com/problems/maximum-product-of-word-lengths/description/
 *
 * algorithms
 * Medium (51.73%)
 * Total Accepted:    103K
 * Total Submissions: 198.8K
 * Testcase Example:  '["abcw","baz","foo","bar","xtfn","abcdef"]'
 *
 * Given a string array words, find the maximum value of length(word[i]) *
 * length(word[j]) where the two words do not share common letters. You may
 * assume that each word will contain only lower case letters. If no such two
 * words exist, return 0.
 *
 * Example 1:
 *
 *
 * Input: ["abcw","baz","foo","bar","xtfn","abcdef"]
 * Output: 16
 * Explanation: The two words can be "abcw", "xtfn".
 *
 * Example 2:
 *
 *
 * Input: ["a","ab","abc","d","cd","bcd","abcd"]
 * Output: 4
 * Explanation: The two words can be "ab", "cd".
 *
 * Example 3:
 *
 *
 * Input: ["a","aa","aaa","aaaa"]
 * Output: 0
 * Explanation: No such pair of words.
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= words.length <= 10^3
 * 0 <= words[i].length <= 10^3
 * words[i] consists only of lowercase English letters.
 *
 *
 */
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
}
