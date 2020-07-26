package lc208

// TreeNode 树的节点
type TreeNode struct {
	value     rune
	isWord    bool
	wordCount int
	childs    map[rune]*TreeNode
}

// Trie 前缀树
type Trie struct {
	root *TreeNode
}

// Constructor leetcode 自动生成
func Constructor() Trie {
	root := &TreeNode{childs: make(map[rune]*TreeNode)}
	return Trie{root}
}

// Insert 插入一个单词
func (tr *Trie) Insert(word string) {
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

// Search 查找一个单词
func (tr *Trie) Search(word string) bool {
	root := tr.root

	for _, r := range word {
		ch, ok := root.childs[r]
		if !ok {
			return false
		}
		root = ch
	}

	if root != tr.root && root.isWord {
		return true
	}

	return false
}

// StartsWith 查找前缀
func (tr *Trie) StartsWith(prefix string) bool {
	root := tr.root

	for _, r := range prefix {
		ch, ok := root.childs[r]
		if !ok {
			return false
		}
		root = ch
	}

	return true
}
