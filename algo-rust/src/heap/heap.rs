#![allow(dead_code)]

use std::collections::{BinaryHeap, HashMap};
use std::{
    cmp::{min, Ord, Ordering, Reverse},
    vec,
};

struct Solution {}

impl Solution {
    pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
        let mut heap = BinaryHeap::new();

        for n in nums {
            heap.push(Reverse(n));

            if heap.len() > k as usize {
                heap.pop();
            }
        }

        heap.peek().unwrap().0
    }

    pub fn top_k_frequent(words: Vec<String>, k: i32) -> Vec<String> {
        let mut map: HashMap<String, i32> = HashMap::new();

        for w in words {
            *map.entry(w).or_default() += 1;
        }

        let mut heap = BinaryHeap::new();

        for entry in map.into_iter() {
            heap.push((-entry.1, entry.0));

            if heap.len() as i32 > k {
                heap.pop();
            }
        }

        let mut res = vec![];
        while let Some(entry) = heap.pop() {
            res.push(entry.1);
        }

        res.reverse();
        res
    }

    pub fn is_possible(nums: Vec<i32>) -> bool {
        let mut map: HashMap<i32, BinaryHeap<Reverse<i32>>> = HashMap::new();
        for n in nums {
            if let None = map.get(&n) {
                map.insert(n, BinaryHeap::new());
            }

            let mut v = 0;
            if let Some(heap) = map.get_mut(&(n - 1)) {
                match heap.pop() {
                    Some(min) => v = min.0,
                    None => {
                        map.remove(&(n - 1));
                    }
                }
            }

            if let Some(h) = map.get_mut(&n) {
                h.push(Reverse(v + 1));
            }
        }

        for (_, v) in map {
            if let Some(min) = v.peek() {
                if min.0 < 3 {
                    return false;
                }
            }
        }

        true
    }

    pub fn furthest_building(heights: Vec<i32>, bricks: i32, ladders: i32) -> i32 {
        let (mut heap, mut t) = (BinaryHeap::new(), 0);

        for i in 1..heights.len() {
            let delta = heights[i] - heights[i - 1];
            if delta > 0 {
                heap.push(Reverse(delta));
                if heap.len() as i32 > ladders {
                    t += heap.pop().unwrap().0;
                }

                if t > bricks {
                    return (i - 1) as i32;
                }
            }
        }
        (heights.len() - 1) as i32
    }

    pub fn rearrange_barcodes(barcodes: Vec<i32>) -> Vec<i32> {
        let n = barcodes.len();
        if n <= 1 {
            return barcodes;
        }

        let mut map = HashMap::new();

        for key in barcodes {
            *map.entry(key).or_insert(0) += 1
        }

        let mut heap = BinaryHeap::new();

        for (key, count) in map {
            heap.push(KeyCount::new(key, count));
        }

        let mut res = Vec::with_capacity(n);

        while heap.len() > 0 {
            match (heap.pop(), heap.pop()) {
                (Some(mut a), Some(mut b)) => {
                    res.push(a.key);
                    res.push(b.key);

                    if a.count > 1 {
                        a.count -= 1;
                        heap.push(a)
                    }

                    if b.count > 1 {
                        b.count -= 1;
                        heap.push(b)
                    }
                }
                (Some(a), None) => {
                    if a.count > 1 {
                        return res;
                    }

                    res.push(a.key);
                }
                _ => {}
            }
        }
        res
    }

    pub fn reorganize_string(s: String) -> String {
        let mut map = HashMap::new();

        for c in s.chars() {
            *map.entry(c).or_insert(0) += 1;
        }

        let mut heap = BinaryHeap::new();

        for (key, count) in map {
            heap.push(KeyCount::new(key, count));
        }

        let mut res = String::new();

        while heap.len() > 0 {
            match (heap.pop(), heap.pop()) {
                (Some(mut max), Some(max_1)) => {
                    res += &(max.key.to_string() + &max_1.key.to_string())
                        .repeat(max_1.count as usize);

                    if max.count != max_1.count {
                        max.count -= max_1.count;
                        heap.push(max);
                    }
                }
                (Some(max), None) => {
                    if max.count > 1 {
                        return String::new();
                    }

                    res += &max.key.to_string();
                }
                _ => {}
            }
        }

        res
    }
    // https://leetcode-cn.com/problems/network-delay-time/solution/dan-yuan-zui-duan-lu-po-su-de-dijkstra-dui-you-hua/
    pub fn network_delay_time(times: Vec<Vec<i32>>, n: i32, k: i32) -> i32 {
        let m = (n + 1) as usize;
        let mut graph = vec![vec![std::i32::MAX; m]; m];

        for item in times {
            let (u, v) = (item[0] as usize, item[1] as usize);
            graph[u][v] = item[2];
        }

        let mut dist = vec![std::i32::MAX; m];
        let mut vist = vec![false; m];

        dist[k as usize] = 0;

        for _i in 0..n - 1 {
            let mut t: usize = 0;
            for j in 1..m {
                if !vist[j] && dist[j] < dist[t] {
                    t = j;
                }
            }
            vist[t] = true;

            for j in 1..m {
                if graph[t][j] != std::i32::MAX {
                    dist[j] = min(dist[j], dist[t] + graph[t][j]);
                }
            }
        }

        match dist.into_iter().skip(1).max() {
            Some(std::i32::MAX) => -1,
            Some(res) => res,
            None => -1,
        }
    }
    pub fn network_delay_time2(times: Vec<Vec<i32>>, n: i32, k: i32) -> i32 {
        let m = (n + 1) as usize;
        let mut graph: HashMap<i32, Vec<[i32; 2]>> = HashMap::new();

        for item in times.into_iter() {
            graph
                .entry(item[0])
                .or_insert(Vec::new())
                .push([item[1], item[2]]);
        }

        let mut dist = vec![std::i32::MAX; m];
        let mut vist = vec![false; m];

        let mut heap = BinaryHeap::new();
        heap.push(Pairs::new(0, k));

        dist[k as usize] = 0;

        while heap.len() > 0 {
            let t = heap.pop().unwrap();
            if vist[t.last as usize] {
                continue;
            }

            vist[t.last as usize] = true;

            if let Some(sides) = graph.get(&t.last) {
                for s in sides.into_iter() {
                    if dist[s[0] as usize] > t.first + s[1] {
                        dist[s[0] as usize] = t.first + s[1];
                        heap.push(Pairs::new(dist[s[0] as usize], s[0]));
                    }
                }
            }
        }

        match dist.into_iter().skip(1).max() {
            Some(std::i32::MAX) => -1,
            Some(res) => res,
            None => -1,
        }
    }
    // https://leetcode.com/problems/find-k-pairs-with-smallest-sums/
    pub fn k_smallest_pairs(nums1: Vec<i32>, nums2: Vec<i32>, k: i32) -> Vec<Vec<i32>> {
        let mut res: Vec<Vec<i32>> = Vec::new();
        let (n, m) = (nums1.len(), nums2.len());

        if n == 0 || m == 0 {
            return res;
        }

        let mut heap = BinaryHeap::new();

        for i in 0..n {
            heap.push(Pairs::new(nums1[i] + nums2[0], Pairs::new(i, 0 as usize)));
        }

        while res.len() < k as usize && heap.len() > 0 {
            let t = heap.pop().unwrap();
            let i = t.last.first;
            let j = t.last.last;

            res.push(vec![nums1[i], nums2[j]]);
            if j + 1 < m {
                heap.push(Pairs::new(nums1[i] + nums2[j + 1], Pairs::new(i, j + 1)));
            }
        }
        res
    }
    // leetcode 1705
    pub fn eaten_apples(apples: Vec<i32>, days: Vec<i32>) -> i32 {
        if apples.len() != days.len() {
            return 0;
        }

        let mut heap = BinaryHeap::new();

        let (mut res, n, mut i) = (0, apples.len(), 0 as usize);

        while i < n || heap.len() > 0 {
            if i < n && apples[i] > 0 {
                heap.push(Pairs::new(days[i] + i as i32, apples[i]));
            }

            while heap.len() > 0 && heap.peek().unwrap().first as usize <= i {
                heap.pop();
            }

            if let Some(mut p) = heap.pop() {
                res += 1;
                p.last -= 1;
                if p.last > 0 {
                    heap.push(p)
                }
            }
            i += 1;
        }
        res
    }

    pub fn most_competitive(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let mut stack = Vec::new();
        for i in 0..nums.len() {
            while stack.len() > 0
                && stack.last().unwrap() > &nums[i]
                && k as usize - stack.len() < nums.len() - i
            {
                stack.pop();
            }

            if stack.len() < k as usize {
                stack.push(nums[i]);
            }
        }

        stack
    }
}

#[derive(Copy, Clone, Eq, PartialEq, Ord, Debug)]
struct Pairs<T: Ord, U: Ord> {
    first: T,
    last: U,
}

impl<T: Ord, U: Ord> Pairs<T, U> {
    pub fn new(first: T, last: U) -> Self {
        Pairs { first, last }
    }
}

impl<T: Ord, U: Ord> PartialOrd for Pairs<T, U> {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        if self.first == other.first {
            return other.last.partial_cmp(&self.last);
        }

        other.first.partial_cmp(&self.first)
    }
}

#[derive(Eq)]
struct KeyCount<T: Ord> {
    key: T,
    count: i32,
}

impl<T: Ord> KeyCount<T> {
    pub fn new(key: T, count: i32) -> Self {
        KeyCount { key, count }
    }
}

impl<T: Ord> Ord for KeyCount<T> {
    fn cmp(&self, other: &Self) -> Ordering {
        if self.count == other.count {
            return self.key.cmp(&other.key);
        }
        self.count.cmp(&other.count)
    }
}

impl<T: Ord> PartialOrd for KeyCount<T> {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        if self.count == other.count {
            return self.key.partial_cmp(&other.key);
        }

        self.count.partial_cmp(&other.count)
    }
}

impl<T: Ord> PartialEq for KeyCount<T> {
    fn eq(&self, other: &Self) -> bool {
        self.count.eq(&other.count) && self.key.eq(&other.key)
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn top_k_frequent() {
        let mut words: Vec<String> = Vec::new();
        words.push("aaa".to_string());
        Solution::top_k_frequent(words, 2);
    }

    #[test]
    fn is_possible() {
        let nums = vec![1, 2, 3, 3, 4, 5];
        assert_eq!(true, Solution::is_possible(nums));

        let nums = vec![1, 2, 3, 3, 4, 4, 5, 5];

        assert_eq!(true, Solution::is_possible(nums));

        let nums = vec![1, 2, 3, 4, 4, 5];

        assert_eq!(false, Solution::is_possible(nums));
    }

    #[test]
    fn furthest_building() {
        let heights = vec![4, 2, 7, 6, 9, 14, 12];

        assert_eq!(4, Solution::furthest_building(heights, 5, 1))
    }

    #[test]
    fn rearrange_barcodes() {
        // let input = vec![1, 1, 1, 2, 2, 2];
        //
        // assert_eq!(vec![1, 2, 1, 2, 1, 2], Solution::rearrange_barcodes(input));
        //
        // let input = vec![1, 1, 1, 1, 2, 2, 3, 3];
        //
        // assert_eq!(
        //     vec![1, 2, 1, 2, 1, 3, 1, 3],
        //     Solution::rearrange_barcodes(input)
        // );

        // let input = vec![1, 1, 2];
        //
        // assert_eq!(vec![1, 2, 1], Solution::rearrange_barcodes(input));
    }

    #[test]
    fn pairs() {
        let mut heap = BinaryHeap::new();
        heap.push(Pairs::new(1, 2));
        heap.push(Pairs::new(1, 3));
        heap.push(Pairs::new(2, 2));
        heap.push(Pairs::new(2, 3));

        assert_eq!(Some(Pairs::new(1, 2)), heap.pop());
        assert_eq!(Some(Pairs::new(1, 3)), heap.pop());
        assert_eq!(Some(Pairs::new(2, 2)), heap.pop());
        assert_eq!(Some(Pairs::new(2, 3)), heap.pop());
    }

    #[test]
    fn most_competitive() {
        let nums = vec![3, 5, 2, 6];
        let k = 2;
        assert_eq!(vec![2, 6], Solution::most_competitive(nums, k));

        let nums = vec![2, 4, 3, 3, 5, 4, 9, 6];
        let k = 4;
        assert_eq!(vec![2, 3, 3, 4], Solution::most_competitive(nums, k));
    }
}
