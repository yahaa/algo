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
use std::collections::HashMap;
use std::collections::HashSet;
use std::collections::VecDeque;

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

    // 863. All Nodes Distance K in Binary Tree
    pub fn distance_k(
        root: Option<Rc<RefCell<TreeNode>>>,
        p: Option<Rc<RefCell<TreeNode>>>,
        k: i32,
    ) -> Vec<i32> {
        fn find_parent(
            root: Option<Rc<RefCell<TreeNode>>>,
            parent: &mut HashMap<i32, Rc<RefCell<TreeNode>>>,
        ) {
            if let Some(r) = root {
                if r.borrow().left.is_some() {
                    parent.insert(r.borrow().left.clone().unwrap().borrow().val, r.clone());
                }

                if r.borrow().right.is_some() {
                    parent.insert(r.borrow().right.clone().unwrap().borrow().val, r.clone());
                }

                find_parent(r.borrow().left.clone(), parent);
                find_parent(r.borrow().right.clone(), parent);
            }
        }

        fn dfs(
            root: Option<Rc<RefCell<TreeNode>>>,
            k: i32,
            parent: &mut HashMap<i32, Rc<RefCell<TreeNode>>>,
            visited: &mut HashSet<i32>,
            res: &mut Vec<i32>,
        ) {
            if let Some(r) = root {
                if visited.contains(&r.borrow().val) {
                    return;
                }

                visited.insert(r.borrow().val);

                if k == 0 {
                    res.push(r.borrow().val);
                }

                dfs(r.borrow().left.clone(), k - 1, parent, visited, res);
                dfs(r.borrow().right.clone(), k - 1, parent, visited, res);

                if let Some(p) = parent.get(&r.borrow().val) {
                    dfs(Some(p.clone()), k - 1, parent, visited, res);
                }
            }
        }

        let mut res = Vec::new();
        let mut parent = HashMap::new();
        let mut visited = HashSet::new();

        find_parent(root, &mut parent);
        dfs(p, k, &mut parent, &mut visited, &mut res);

        res
    }

    // 1339. Maximum Product of Splitted Binary Tree todo
    pub fn max_product(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        fn build_sum(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
            match root {
                Some(r) => {
                    build_sum(r.borrow().left.clone())
                        + build_sum(r.borrow().right.clone())
                        + r.borrow().val
                }
                None => 0,
            }
        }

        fn find_max(root: Option<Rc<RefCell<TreeNode>>>, sum: &i64, max: &mut i64) -> i64 {
            match root {
                Some(r) => {
                    let left = find_max(r.borrow().left.clone(), sum, max);
                    let right = find_max(r.borrow().right.clone(), sum, max);

                    *max = std::cmp::max(*max, (sum - left) * left);
                    *max = std::cmp::max(*max, (sum - right) * right);

                    left + right + r.borrow().val as i64
                }
                None => 0 as i64,
            }
        }

        let sum = build_sum(root.clone()) as i64;
        let mut max = 0 as i64;

        find_max(root, &sum, &mut max);

        (max % 1000000007) as i32
    }

    // 1161. Maximum Level Sum of a Binary Tree
    pub fn max_level_sum(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        if root.is_none() {
            return 0;
        }

        let mut queue = VecDeque::new();
        let (mut max_sum, mut max_level, mut level) = (root.clone().unwrap().borrow().val, 1, 1);

        queue.push_back(root);

        while !queue.is_empty() {
            let mut level_size = queue.len();
            let mut level_sum = 0;

            while level_size > 0 {
                let r = queue.pop_front().unwrap().unwrap();

                level_sum += r.borrow().val;

                if r.borrow().left.is_some() {
                    queue.push_back(r.borrow().left.clone());
                }

                if r.borrow().right.is_some() {
                    queue.push_back(r.borrow().right.clone());
                }

                level_size -= 1;
            }

            if level_sum > max_sum {
                max_sum = level_sum;
                max_level = level;
            }

            level += 1;
        }

        max_level
    }

    // 988. Smallest String Starting From Leaf
    pub fn smallest_from_leaf(root: Option<Rc<RefCell<TreeNode>>>) -> String {
        fn dfs(root: Option<Rc<RefCell<TreeNode>>>, cur: String, ans: &mut String) {
            if let Some(r) = root {
                let c = std::char::from_u32(r.borrow().val as u32 + 97)
                    .unwrap()
                    .to_string();

                let cur = c + &cur;

                if r.borrow().left.is_none() && r.borrow().right.is_none() {
                    *ans = std::cmp::min(ans.clone(), cur.clone());
                }

                dfs(r.borrow().left.clone(), cur.clone(), ans);
                dfs(r.borrow().right.clone(), cur.clone(), ans);
            }
        }

        let mut ans = String::from("~");
        dfs(root, "".to_string(), &mut ans);

        ans
    }

    // 701. Insert into a Binary Search Tree todo
    pub fn insert_into_bst(
        root: Option<Rc<RefCell<TreeNode>>>,
        val: i32,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        unimplemented!()
    }

    // 617. Merge Two Binary Trees todo
    pub fn merge_trees(
        root1: Option<Rc<RefCell<TreeNode>>>,
        root2: Option<Rc<RefCell<TreeNode>>>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        fn merge(
            root1: Option<Rc<RefCell<TreeNode>>>,
            root2: Option<Rc<RefCell<TreeNode>>>,
        ) -> Option<Rc<RefCell<TreeNode>>> {
            let r = Rc::new(RefCell::new(TreeNode::new(0)));

            match (root1, root2) {
                (Some(r1), Some(r2)) => {
                    r.borrow_mut().val = r1.borrow().val + r2.borrow().val;
                    r.borrow_mut().left = merge(r1.borrow().left.clone(), r2.borrow().left.clone());
                    r.borrow_mut().right =
                        merge(r1.borrow().right.clone(), r2.borrow().right.clone());

                    Some(r)
                }
                (Some(r1), None) => {
                    r.borrow_mut().val = r1.borrow().val;
                    r.borrow_mut().left = merge(r1.borrow().left.clone(), None);
                    r.borrow_mut().right = merge(None, r1.borrow().right.clone());

                    Some(r)
                }
                (None, Some(r2)) => {
                    r.borrow_mut().val = r2.borrow().val;
                    r.borrow_mut().left = merge(r2.borrow().left.clone(), None);
                    r.borrow_mut().right = merge(None, r2.borrow().right.clone());

                    Some(r)
                }
                _ => None,
            }
        }

        merge(root1, root2)
    }

    // 1325. Delete Leaves With a Given Value
    pub fn remove_leaf_nodes(
        root: Option<Rc<RefCell<TreeNode>>>,
        target: i32,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        match root {
            None => None,
            Some(r) => {
                let left = Solution::remove_leaf_nodes(r.borrow().left.clone(), target);
                let right = Solution::remove_leaf_nodes(r.borrow().right.clone(), target);

                r.borrow_mut().left = left;
                r.borrow_mut().right = right;

                if r.borrow().left.is_none()
                    && r.borrow().right.is_none()
                    && r.borrow().val == target
                {
                    return None;
                }

                Some(r)
            }
        }
    }

    // 814. Binary Tree Pruning
    pub fn prune_tree(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        match root {
            None => None,
            Some(r) => {
                let left = Solution::prune_tree(r.borrow().left.clone());
                let right = Solution::prune_tree(r.borrow().right.clone());

                r.borrow_mut().left = left;
                r.borrow_mut().right = right;

                if r.borrow().left.is_none() && r.borrow().right.is_none() && r.borrow().val == 0 {
                    return None;
                }

                Some(r)
            }
        }
    }

    // 513. Find Bottom Left Tree Value
    pub fn find_bottom_left_value(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        if root.is_none() {
            return 0;
        }

        let (mut res, mut queue) = (0, VecDeque::new());

        queue.push_back(root.unwrap().clone());

        while queue.len() > 0 {
            let size = queue.len();
            for i in 0..size {
                if let Some(top) = queue.pop_front() {
                    if top.borrow().left.is_some() {
                        queue.push_back(top.borrow().left.clone().unwrap());
                    }

                    if top.borrow().right.is_some() {
                        queue.push_back(top.borrow().right.clone().unwrap());
                    }

                    if i == 0 {
                        res = top.borrow().val;
                    }
                }
            }
        }

        res
    }
}
