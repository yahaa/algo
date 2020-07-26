package lfu

// node 链表节点
type node struct {
	value int
	freq  int
	key   int
	next  *node
	pre   *node
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

	if n == ln.head && n == ln.tail {
		ln.head = nil
		ln.tail = ln.head
		return
	}

	if n == ln.head {
		ln.head = n.next
		if ln.head != nil {
			ln.head.pre = nil
		}
		return
	}


	if n == ln.tail {
		ln.tail = n.pre
		if ln.tail != nil {
			ln.tail.next = nil
		}
		return
	}

	if n.pre != nil {
		n.pre.next = n.next
	}

	if n.next != nil {
		n.next.pre = n.pre
	}
}

// LFUCache  Least Frequently Used  最不常使用淘汰算法，当最不常使用频率有多个按照 LRU 淘汰
type LFUCache struct {
	cache    map[int]*node
	freqMap  map[int]*link
	minFreq  int
	capacity int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cache:    make(map[int]*node),
		freqMap:  make(map[int]*link),
		capacity: capacity,
		minFreq:  int(^uint(0) >> 1),
	}
}

// Get 查询 key
// 1. 当 key 不存在返回 -1
// 2. 当 key 存在, 通过 cache 获取 key 所在的 node
//    a. 通过 key 所在的 node 获取 key 的 freq，在 freqMap[freq] 的双向链表中删除改 node
//    b. 在 freqMap[freq+1] 的双向链表头结点中插入使用频率更新后的 node
//    c. 更新 minFreq 返回 node.value
func (fu *LFUCache) Get(key int) int {
	n, ok := fu.cache[key]
	if !ok {
		return -1
	}

	fu.freqMap[n.freq].Del(n)

	n.next = nil
	n.pre = nil
	n.freq++

	fu.setFreqLink(n)

	if fu.freqMap[fu.minFreq].head == nil {
		fu.minFreq = n.freq
	}

	return n.value
}

// Put 插入 key value
// 1. 当 key 不存在
//    a. 当 cache 已满
//       1. 删除 freqMap[minFreq] 尾结点
//       2. 在 freqMap[1] 头结点中插入 node,cache 记录下 node 的地址
//       3. 更新 minFrq 为 1
//    b. 当 cache 未满
//       2. 在 freqMap[1] 头结点中插入 node,cache 记录下 node 的地址
//       3. 更新 minFrq 为 1
// 2. key 已经存在，通过 cache 获取 key 所在的 node
//    a. 通过 key 所在的 node 获取 key 的 freq, 在 freqMap[freq] 的双向链表中删除该 node
//    b. 在 freqMap[freq+1] 的双向链表头结点中插入使用频率更新后的 node
func (fu *LFUCache) Put(key, value int) {
	if fu.capacity <= 0 {
		return
	}

	n, ok := fu.cache[key]
	if ok {
		fu.freqMap[n.freq].Del(n)

		n.freq++
		n.value = value

		fu.setFreqLink(n)

		if fu.freqMap[fu.minFreq].head == nil {
			fu.minFreq = n.freq
		}

		return
	}

	n = &node{value: value, key: key, freq: 1}
	fu.setFreqLink(n)
	fu.cache[key] = n

	if len(fu.cache) > fu.capacity {
		tail := fu.freqMap[fu.minFreq].tail
		fu.freqMap[fu.minFreq].Del(tail)

		delete(fu.cache, tail.key)
	}

	fu.minFreq = 1
}

func (fu *LFUCache) setFreqLink(n *node) {
	if fu.freqMap[n.freq] == nil {
		dlink := link{}
		dlink.Insert(n)
		fu.freqMap[n.freq] = &dlink
	} else {
		fu.freqMap[n.freq].Insert(n)
	}
}
