package lt760

import "container/list"

type MyHashMap struct {
	bucket []*list.List
	mod    int
}

type KV struct {
	key   int
	value int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		bucket: make([]*list.List, 2048),
		mod:    2048,
	}
}

/** value will always be non-negative. */
func (mp *MyHashMap) Put(key int, value int) {
	m := key % mp.mod

	if mp.bucket[m] == nil {
		l := list.New()
		l.PushBack(KV{key, value})
		mp.bucket[m] = l
		return
	}

	l := mp.bucket[m]
	head := l.Front()
	found := false
	for head != nil {
		v := head.Value.(KV)
		if v.key == key {
			v.value = value
			head.Value = v
			found = true
			break
		}

		head = head.Next()
	}

	if !found {
		l.PushBack(KV{key, value})
	}
}

/** Returns the value to which the specified key is mapped, or -1 if mp map contains no mapping for the key */
func (mp *MyHashMap) Get(key int) int {
	m := key % mp.mod

	if mp.bucket[m] == nil {
		return -1
	}

	l := mp.bucket[m]
	head := l.Front()
	for head != nil {
		v := head.Value.(KV)
		if v.key == key {
			return v.value
		}

		head = head.Next()
	}

	return -1
}

/** Removes the mapping of the specified value key if mp map contains a mapping for the key */
func (mp *MyHashMap) Remove(key int) {
	m := key % mp.mod

	if mp.bucket[m] == nil {
		return
	}
	l := mp.bucket[m]
	head := l.Front()
	for head != nil {
		v := head.Value.(KV)
		if v.key == key {
			l.Remove(head)
			return
		}

		head = head.Next()
	}
}
