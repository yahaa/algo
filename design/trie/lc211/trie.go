package lc211

// TreeNode 树的节点
type TreeNode struct {
	value     rune
	isWord    bool
	wordCount int
	childs    map[rune]*TreeNode
}

// WordDictionary 字典
type WordDictionary struct {
	root *TreeNode
}

// Constructor 构造器
func Constructor() WordDictionary {
	root := &TreeNode{childs: make(map[rune]*TreeNode)}
	return WordDictionary{root}

}

// AddWord 添加单词
func (tr *WordDictionary) AddWord(word string) {
	root := tr.root

	for _, r := range word {
		if ch, ok := root.childs[r]; ok {
			root = ch
		} else {
			n := &TreeNode{
				value:  r,
				childs: make(map[rune]*TreeNode),
			}

			root.childs[r] = n
			root = n
		}
	}

	if root != tr.root {
		root.isWord = true
		root.wordCount++
	}
}

// Search 索搜单词
func (tr *WordDictionary) Search(word string) bool {
	var dfs func(i int, root *TreeNode) bool

	dfs = func(i int, root *TreeNode) bool {
		if i == len(word) {
			return root.isWord
		}

		c := rune(word[i])

		if c == '.' {
			for _, n := range root.childs {
				if dfs(i+1, n) {
					return true
				}
			}
			return false
		}

		n, ok := root.childs[c]
		if ok {
			return dfs(i+1, n)
		}
		return false
	}

	return dfs(0, tr.root)
}
