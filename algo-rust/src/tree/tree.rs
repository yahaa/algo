#![allow(dead_code)]

// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

use std::cell::RefCell;
use std::cmp::max;
use std::rc::Rc;

struct Solution {}

impl Solution {
    // leetcode 222
    pub fn count_nodes_1(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        return match root {
            None => 0,
            Some(r) => {
                1 + Solution::count_nodes_1(r.borrow().left.clone())
                    + Solution::count_nodes_1(r.borrow().right.clone())
            }
        };
    }

    // leetcode 222
    pub fn count_nodes(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let (mut r, mut h) = (root.clone(), 0);

        while let Some(left) = r {
            r = left.borrow().left.clone();
            h += 1;
        }

        if h <= 1 {
            return h;
        }

        let (mut low, mut hight) = (1 << (h - 1), (1 << h) - 1);

        while low < hight {
            let mid = (hight - low + 1) / 2 + low;

            match Solution::find_node(root.clone(), mid, h) {
                true => low = mid,
                false => hight = mid - 1,
            }
        }

        low
    }

    fn find_node(root: Option<Rc<RefCell<TreeNode>>>, val: i32, hight: i32) -> bool {
        let (mut r, mut bit) = (root.clone(), 1 << (hight - 2));

        while bit > 0 && r.is_some() {
            if val & bit == 0 {
                r = r.unwrap().borrow().left.clone();
            } else {
                r = r.unwrap().borrow().right.clone();
            }

            bit >>= 1;
        }

        r.is_some()
    }

    // leetcode 230
    pub fn kth_smallest(root: Option<Rc<RefCell<TreeNode>>>, k: i32) -> i32 {
        let mut res = 0;
        let mut k = k;

        fn inorder(root: &Option<Rc<RefCell<TreeNode>>>, k: &mut i32, res: &mut i32) {
            if root.is_none() {
                return;
            }

            inorder(&root.as_ref().unwrap().borrow().left.clone(), k, res);

            *k -= 1;
            if *k == 0 {
                *res = root.as_ref().unwrap().borrow().val;
                return;
            }

            return inorder(&root.as_ref().unwrap().borrow().right.clone(), k, res);
        }

        inorder(&root, &mut k, &mut res);
        res
    }

    // leetcode 114
    pub fn flatten(root: &mut Option<Rc<RefCell<TreeNode>>>) {
        let mut list = Vec::new();

        fn post_order(root: Option<Rc<RefCell<TreeNode>>>, list: &mut Vec<Rc<RefCell<TreeNode>>>) {
            if let Some(r) = root {
                list.push(r.clone());

                post_order(r.borrow_mut().left.take(), list);
                post_order(r.borrow_mut().right.take(), list);
            }
        }

        post_order(root.clone(), &mut list);

        for i in 1..list.len() {
            let pre = list[i - 1].clone();
            let cur = list[i].clone();

            pre.borrow_mut().left.take();

            pre.borrow_mut().right = Some(cur);
        }
    }

    // leetcode 144
    pub fn preorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut list = Vec::new();

        fn pre_order(root: Option<Rc<RefCell<TreeNode>>>, list: &mut Vec<i32>) {
            if let Some(r) = root {
                list.push(r.borrow().val);

                pre_order(r.borrow().left.clone(), list);
                pre_order(r.borrow().right.clone(), list);
            }
        }

        pre_order(root, &mut list);

        list
    }

    // leetcode 124
    pub fn max_path_sum(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut sum = i32::MIN;

        fn dfs(root: Option<Rc<RefCell<TreeNode>>>, sum: &mut i32) -> i32 {
            if let Some(r) = root {
                let left = max(dfs(r.borrow().left.clone(), sum), 0);
                let right = max(dfs(r.borrow().right.clone(), sum), 0);

                *sum = max(*sum, r.borrow().val + left + right);

                return r.borrow().val + max(left, right);
            }

            return 0;
        }

        dfs(root, &mut sum);
        sum
    }

    // leetcode 129
    pub fn sum_numbers(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut sum = 0;
        fn dfs(root: Option<Rc<RefCell<TreeNode>>>, mut pre: i32, sum: &mut i32) {
            if let Some(r) = root {
                pre = pre * 10 + r.borrow().val;
                if r.borrow().left.is_none() && r.borrow().right.is_none() {
                    *sum += pre;
                }

                dfs(r.borrow().left.clone(), pre, sum);
                dfs(r.borrow().right.clone(), pre, sum);
            }
        }

        dfs(root, 0, &mut sum);
        sum
    }

    // leetcode 145
    pub fn postorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut list = Vec::new();

        fn post_order(root: Option<Rc<RefCell<TreeNode>>>, list: &mut Vec<i32>) {
            if let Some(r) = root {
                post_order(r.borrow().left.clone(), list);
                post_order(r.borrow().right.clone(), list);

                list.push(r.borrow().val);
            }
        }

        post_order(root, &mut list);

        list
    }
    // leetcode 199
    pub fn right_side_view(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut list = Vec::new();

        fn dfs(root: Option<Rc<RefCell<TreeNode>>>, level: i32, list: &mut Vec<i32>) {
            if let Some(r) = root {
                if list.len() <= level as usize {
                    list.push(r.borrow().val);
                } else {
                    list[level as usize] = r.borrow().val;
                }

                dfs(r.borrow().left.clone(), level + 1, list);
                dfs(r.borrow().right.clone(), level + 1, list);
            }
        }

        dfs(root, 0, &mut list);
        list
    }

    // leetcode 226
    pub fn invert_tree(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        fn dfs(root: Option<Rc<RefCell<TreeNode>>>) {
            if let Some(r) = root {
                let left = r.borrow_mut().left.take();
                let right = r.borrow_mut().right.take();

                dfs(left.clone());
                dfs(right.clone());

                r.borrow_mut().left = right;
                r.borrow_mut().right = left;
            }
        }

        dfs(root.clone());
        root
    }

    // leetcode 257
    pub fn binary_tree_paths(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<String> {
        let mut list = Vec::new();
        fn dfs(root: Option<Rc<RefCell<TreeNode>>>, mut pre: String, list: &mut Vec<String>) {
            if let Some(r) = root {
                if pre.len() == 0 {
                    pre = format!("{}", r.borrow().val);
                } else {
                    pre = format!("{}->{}", pre, r.borrow().val);
                }

                if r.borrow().left.is_none() && r.borrow().right.is_none() {
                    list.push(pre.clone());
                }

                dfs(r.borrow().left.clone(), pre.clone(), list);
                dfs(r.borrow().right.clone(), pre.clone(), list);
            }
        }

        dfs(root, String::new(), &mut list);
        list
    }

    // leetcode 1302
    pub fn deepest_leaves_sum(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut list = Vec::new();

        fn dfs(root: Option<Rc<RefCell<TreeNode>>>, level: i32, list: &mut Vec<i32>) {
            if let Some(r) = root {
                if list.len() <= level as usize {
                    list.push(r.borrow().val);
                } else {
                    list[level as usize] += r.borrow().val;
                }

                dfs(r.borrow().left.clone(), level + 1, list);
                dfs(r.borrow().right.clone(), level + 1, list);
            }
        }

        dfs(root, 0, &mut list);
        match list.last() {
            None => 0,
            Some(v) => *v,
        }
    }

    // leetcode 404
    pub fn sum_of_left_leaves(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut sum = 0;
        fn dfs(root: Option<Rc<RefCell<TreeNode>>>, sum: &mut i32, is_left: bool) {
            if let Some(r) = root {
                if is_left && r.borrow().left.is_none() && r.borrow().right.is_none() {
                    *sum += r.borrow().val;
                }

                dfs(r.borrow().left.clone(), sum, true);
                dfs(r.borrow().right.clone(), sum, false);
            }
        }

        dfs(root, &mut sum, false);

        sum
    }

    // leetcode 501
    pub fn find_mode(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        //        let mut map: HashMap<i32, Vec<i32>> = HashMap::new();
        //        let mut max_count = 0;
        //
        //        fn dfs(
        //            root: Option<Rc<RefCell<TreeNode>>>,
        //            pre: Option<&mut i32>,
        //            max_count: &mut i32,
        //            map: &mut HashMap<i32, Vec<i32>>,
        //        ) {
        //            match root {
        //                Some(r) => {
        //                    dfs(r.borrow().left.clone(), pre, max_count, map);
        //
        //                    if let Some(prev) = pre {
        //                    } else {
        //                        let mut val = r.borrow().val;
        //
        //                        pre = Some(prev.borrow().val);
        //                    }
        //                    dfs(r.borrow().right.clone(), pre, max_count, map);
        //                }
        //                None => {
        //                    return;
        //                }
        //            }
        //        }
        //
        unimplemented!()
    }

    // 863. All Nodes Distance K in Binary Tree todo
    pub fn distance_k(
        root: Option<Rc<RefCell<TreeNode>>>,
        p: Option<Rc<RefCell<TreeNode>>>,
        K: i32,
    ) -> Vec<i32> {
        unimplemented!()
    }
}
