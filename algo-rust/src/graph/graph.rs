#![allow(dead_code)]
use std::cmp::min;
use std::collections::HashMap;
use std::collections::VecDeque;

struct Solution {}

impl Solution {
    // leetcode 207
    pub fn can_finish(num: i32, prerequisites: Vec<Vec<i32>>) -> bool {
        let mut graph = vec![vec![]; num as usize];
        let mut indegree = vec![0; num as usize];

        for edge in prerequisites {
            graph[edge[1] as usize].push(edge[0]);
            indegree[edge[0] as usize] += 1;
        }

        let (mut queue, mut count) = (Vec::new(), 0);

        for i in 0..num {
            if indegree[i as usize] == 0 {
                queue.push(i);
            }
        }

        while let Some(s) = queue.pop() {
            count += 1;

            for e in graph[s as usize].iter() {
                indegree[*e as usize] -= 1;

                if indegree[*e as usize] <= 0 {
                    queue.push(*e);
                }
            }
        }

        count == num
    }

    // leetcode 210
    pub fn find_order(num: i32, prerequisites: Vec<Vec<i32>>) -> Vec<i32> {
        let mut graph = vec![vec![]; num as usize];
        let mut indegree = vec![0; num as usize];

        for edge in prerequisites {
            graph[edge[1] as usize].push(edge[0]);
            indegree[edge[0] as usize] += 1;
        }

        let (mut queue, mut count, mut result) = (Vec::new(), 0, Vec::new());

        for i in 0..num {
            if indegree[i as usize] == 0 {
                queue.push(i);
            }
        }

        while let Some(s) = queue.pop() {
            count += 1;
            result.push(s);

            for e in graph[s as usize].iter() {
                indegree[*e as usize] -= 1;

                if indegree[*e as usize] <= 0 {
                    queue.push(*e);
                }
            }
        }

        if count == num {
            return result;
        }

        return vec![];
    }

    // leetcode 310
    pub fn find_min_height_trees(n: i32, edges: Vec<Vec<i32>>) -> Vec<i32> {
        let mut graph = vec![vec![]; n as usize];

        for e in edges {
            graph[e[0] as usize].push(e[1] as usize);
            graph[e[1] as usize].push(e[0] as usize);
        }

        let mut hmap: HashMap<i32, Vec<i32>> = HashMap::new();
        let mut min_high = n;

        for root in 0..n as usize {
            let high = Solution::tree_hight(root, &graph, n);

            let list = hmap.entry(high).or_insert(vec![]);
            list.push(root as i32);

            min_high = min(min_high, high);
        }

        match hmap.remove(&min_high) {
            Some(res) => res,
            None => vec![],
        }
    }

    // leetcode 310 help, find the high of the tree
    fn tree_hight(root: usize, graph: &Vec<Vec<usize>>, n: i32) -> i32 {
        let mut visit = vec![false; n as usize];
        let (mut queue, mut high) = (VecDeque::new(), 0);

        queue.push_back(root);

        while !queue.is_empty() {
            let mut size = queue.len();

            while size > 0 {
                let s = queue.pop_front().unwrap();
                visit[s] = true;

                for e in graph[s].iter() {
                    if !visit[*e] {
                        queue.push_back(*e);
                    }
                }

                size -= 1;
            }

            high += 1;
        }

        high
    }

    // leetcode 310
    pub fn find_min_height_trees2(n: i32, edges: Vec<Vec<i32>>) -> Vec<i32> {
        if n == 1 {
            return vec![0];
        }

        let mut graph = vec![vec![]; n as usize];
        let mut degree = vec![0; n as usize];
        let mut res = Vec::new();

        for e in edges {
            degree[e[0] as usize] += 1;
            degree[e[1] as usize] += 1;

            graph[e[0] as usize].push(e[1] as usize);
            graph[e[1] as usize].push(e[0] as usize);
        }
        let mut queue = VecDeque::new();

        for i in 0..n as usize {
            if degree[i] == 1 {
                queue.push_back(i);
            }
        }

        while !queue.is_empty() {
            res = vec![];
            let mut size = queue.len();

            while size > 0 {
                let s = queue.pop_front().unwrap();
                res.push(s as i32);

                for e in graph[s].iter() {
                    degree[*e] -= 1;
                    if degree[*e] == 1 {
                        queue.push_back(*e);
                    }
                }
                size -= 1;
            }
        }

        res
    }
}
#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn find_min_height_trees() {
        let edges = vec![vec![1, 0], vec![1, 2], vec![1, 3]];

        assert_eq!(vec![1], Solution::find_min_height_trees(4, edges));

        let edges = vec![vec![3, 0], vec![3, 1], vec![3, 2], vec![3, 4], vec![5, 4]];

        assert_eq!(vec![3, 4], Solution::find_min_height_trees(6, edges));

        let edges = vec![vec![0, 1], vec![0, 2], vec![0, 3], vec![3, 4], vec![4, 5]];
        assert_eq!(vec![3], Solution::find_min_height_trees(6, edges));
    }
    #[test]
    fn find_min_height_trees2() {
        let edges = vec![vec![1, 0], vec![1, 2], vec![1, 3]];

        assert_eq!(vec![1], Solution::find_min_height_trees2(4, edges));

        let edges = vec![vec![3, 0], vec![3, 1], vec![3, 2], vec![3, 4], vec![5, 4]];

        assert_eq!(vec![3, 4], Solution::find_min_height_trees2(6, edges));

        let edges = vec![vec![0, 1], vec![0, 2], vec![0, 3], vec![3, 4], vec![4, 5]];
        assert_eq!(vec![3], Solution::find_min_height_trees2(6, edges));
    }

    #[test]
    fn can_finish() {
        let graph = vec![vec![1, 0]];
        assert_eq!(true, Solution::can_finish(2, graph));

        let graph = vec![vec![1, 0], vec![0, 1]];
        assert_eq!(false, Solution::can_finish(2, graph));

        let graph = vec![vec![0, 1]];
        assert_eq!(true, Solution::can_finish(2, graph));
    }
}
