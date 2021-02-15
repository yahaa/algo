use std::collections::HashMap;
struct Solution {}

impl Solution {
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
