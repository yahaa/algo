package lc1472

type node struct {
	value string
	next  *node
	pre   *node
}

//BrowserHistory 浏览器历史
type BrowserHistory struct {
	root *node
	cur  *node
}

//Constructor 构造函数
func Constructor(homepage string) BrowserHistory {
	root := &node{value: homepage}
	cur := root
	return BrowserHistory{root: root, cur: cur}
}

// Visit 浏览
func (b *BrowserHistory) Visit(url string) {
	n := &node{value: url, pre: b.cur}
	b.cur.next = n

	b.cur = n
}

// Back 回退
func (b *BrowserHistory) Back(steps int) string {
	for steps > 0 && b.cur.pre != nil {
		b.cur = b.cur.pre
		steps--
	}

	return b.cur.value
}

// Forward 前进
func (b *BrowserHistory) Forward(steps int) string {
	for steps > 0 && b.cur.next != nil {
		b.cur = b.cur.next
		steps--
	}

	return b.cur.value
}
