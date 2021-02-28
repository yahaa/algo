#![allow(dead_code)]

use std::collections::{BTreeMap, HashMap, HashSet, LinkedList};

struct SubrectangleQueries {
    rectangle: Vec<Vec<i32>>,
}

impl SubrectangleQueries {
    fn new(rectangle: Vec<Vec<i32>>) -> Self {
        return SubrectangleQueries { rectangle };
    }
    fn update_subrectangle(&mut self, row1: i32, col1: i32, row2: i32, col2: i32, new_value: i32) {
        for i in row1..=row2 {
            for j in col1..=col2 {
                self.rectangle[i as usize][j as usize] = new_value;
            }
        }
    }
    fn get_value(&self, row: i32, col: i32) -> i32 {
        self.rectangle[row as usize][col as usize]
    }
}

struct Twitter {
    tweets: LinkedList<(i32, i32)>, // user_id,tweet_id
    relations: HashMap<i32, HashSet<i32>>,
}

impl Twitter {
    fn new() -> Self {
        Twitter {
            tweets: LinkedList::new(),
            relations: HashMap::new(),
        }
    }

    fn post_tweet(&mut self, user_id: i32, tweet_id: i32) {
        self.tweets.push_front((user_id, tweet_id));
    }

    fn get_news_feed(&self, user_id: i32) -> Vec<i32> {
        self.tweets
            .iter()
            .filter(|item| {
                if item.0 == user_id {
                    return true;
                }
                match self.relations.get(&user_id) {
                    Some(set) => return set.contains(&item.0),
                    None => return false,
                }
            })
            .take(10)
            .map(|item| item.1)
            .collect()
    }

    fn follow(&mut self, follower_id: i32, followee_id: i32) {
        self.relations
            .entry(follower_id)
            .or_insert(HashSet::new())
            .insert(followee_id);
    }

    fn unfollow(&mut self, follower_id: i32, followee_id: i32) {
        self.relations
            .entry(follower_id)
            .or_insert(HashSet::new())
            .remove(&followee_id);
    }
}

struct MyCalendarTwo {
    delta: BTreeMap<i32, i32>,
}

impl MyCalendarTwo {
    fn new() -> Self {
        MyCalendarTwo {
            delta: BTreeMap::new(),
        }
    }

    fn book(&mut self, start: i32, end: i32) -> bool {
        *self.delta.entry(start).or_insert(0) += 1;
        *self.delta.entry(end).or_insert(0) += -1;

        let mut sum = 0;

        for v in self.delta.values() {
            sum += *v;
            if sum >= 3 {
                *self.delta.entry(start).or_insert(0) -= 1;
                *self.delta.entry(end).or_insert(0) += 1;

                return false;
            }
        }

        true
    }
}

struct CustomStack {
    max_size: i32,
    stack: Vec<i32>,
}

impl CustomStack {
    fn new(max_size: i32) -> Self {
        return CustomStack {
            max_size,
            stack: Vec::new(),
        };
    }

    fn push(&mut self, x: i32) {
        if self.stack.len() >= self.max_size as usize {
            return;
        }

        self.stack.push(x);
    }

    fn pop(&mut self) -> i32 {
        match self.stack.pop() {
            Some(v) => v,
            None => -1,
        }
    }

    fn increment(&mut self, k: i32, val: i32) {
        let mut n = k as usize;

        if self.stack.len() <= n {
            n = self.stack.len();
        }

        for i in (0..n).rev() {
            self.stack[i] += val;
        }
    }
}

// leetcode 901
struct StockSpanner {
    history: Vec<(i32, i32)>,
}

impl StockSpanner {
    fn new() -> Self {
        let mut history = Vec::new();
        history.push((i32::MAX, 0));

        StockSpanner { history }
    }

    fn next(&mut self, price: i32) -> i32 {
        let mut tmp = 1;

        while !self.history.is_empty() && self.history.last().unwrap().0 <= price {
            tmp += self.history.pop().unwrap().1;
        }

        self.history.push((price, tmp));

        tmp
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_subrectangle_queries() {
        let nums: Vec<Vec<i32>> = vec![vec![1, 2, 2, 3], vec![1, 2, 2, 3]];
        let mut rec = SubrectangleQueries::new(nums);
        assert_eq!(1, rec.get_value(0, 0));
        rec.update_subrectangle(0, 0, 0, 0, 100);
        assert_eq!(100, rec.get_value(0, 0));
    }
}
