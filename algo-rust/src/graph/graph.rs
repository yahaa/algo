#![allow(dead_code)]

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
}
#[cfg(test)]
mod test {
    use super::*;

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
