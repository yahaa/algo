package lru

import (
	"container/list"
)

type node struct {
	v    int
	key  int
	next *node
	pre  *node
}

// link 双向链表
type link struct {
	head *node
	tail *node
}

// Insert 在链表头插入新节点
func (ln *link) Insert(n *node) {
	if n == nil {
		return
	}

	if ln.head == nil {
		ln.head = n
		ln.tail = n
		return
	}

	n.next = ln.head
	ln.head.pre = n
	ln.head = n
}

// Del 删除双向链表中的指定节点
func (ln *link) Del(n *node) {
	if n == nil {
		return
	}

	if n == ln.head {
		ln.head = n.next
		return
	}

	if n == ln.tail {
		ln.tail = n.pre
		return
	}

	if n.pre != nil {
		n.pre.next = n.next
	}

	if n.next != nil {
		n.next.pre = n.pre
	}
}

// LRUCache  Least Recently Used 最近最少使用算法
type LRUCache struct {
	cache    map[int]*node
	capacity int
	dlink    link
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]*node),
		capacity: capacity,
	}
}

// Get 向 LRUCache 中获取元素
// 1. key 存在，把该 key 所在的节点删除，重新插入表头
// 2. key 不存在，返回 -1
func (ru *LRUCache) Get(key int) int {
	n, ok := ru.cache[key]
	if !ok {
		return -1
	}

	ru.dlink.Del(n)
	ru.dlink.Insert(n)

	return n.v
}

// Put 向 LRUCache 中添加元素
// 1. key 已经存在，直接删除 key 所在节点,新节点插入链表头
// 2. key 不存在分为两种情况
//    a. Cache 未满，直接在表头插入
//    b. Cache 已满，删除链表尾节点，再表头插入新节点
func (ru *LRUCache) Put(key int, value int) {
	if ru.capacity <= 0 {
		return
	}

	old, ok := ru.cache[key]
	if ok {
		ru.dlink.Del(old)
		delete(ru.cache, key)
	}

	n := &node{v: value, key: key}
	ru.dlink.Insert(n)
	ru.cache[key] = n

	if len(ru.cache) > ru.capacity {
		delete(ru.cache, ru.dlink.tail.key)
		ru.dlink.Del(ru.dlink.tail)
	}
}

type pair struct {
	key   int
	value int
}

type LRUCache2 struct {
	capacity int
	keys     map[int]*list.Element
	link     *list.List
}

func Constructor2(capacity int) LRUCache2 {
	return LRUCache2{
		capacity: capacity,
		keys:     make(map[int]*list.Element),
		link:     list.New(),
	}
}

func (this *LRUCache2) Get(key int) int {
	e, ok := this.keys[key]
	if !ok {
		return -1
	}

	p := e.Value.(pair)

	this.link.Remove(e)
	delete(this.keys, p.key)

	e = this.link.PushFront(p)
	this.keys[key] = e

	return p.value
}

func (this *LRUCache2) Put(key int, value int) {
	e, ok := this.keys[key]
	if ok {
		this.link.Remove(e)

		e = this.link.PushFront(pair{key, value})
		this.keys[key] = e
		return
	}

	e = this.link.PushFront(pair{key, value})
	this.keys[key] = e

	if this.link.Len() > this.capacity {
		b := this.link.Back()
		this.link.Remove(b)
		delete(this.keys, b.Value.(pair).key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
