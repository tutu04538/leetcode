/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */
package main

// @lc code=start
type DoubleLinkedList struct {
	next       *DoubleLinkedList
	prev       *DoubleLinkedList
	key, value int
}

type LRUCache struct {
	head, tail   *DoubleLinkedList
	capacity     int
	currentItems int
	keyItemMap   map[int]*DoubleLinkedList
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		capacity: capacity,
		head:     &DoubleLinkedList{},
		tail:     &DoubleLinkedList{},
	}
	c.head.next = c.tail
	c.head.prev = c.tail
	c.tail.next = c.head
	c.tail.prev = c.head
	c.keyItemMap = make(map[int]*DoubleLinkedList, capacity)
	return c
}

func (this *LRUCache) Get(key int) int {
	if item, ok := this.keyItemMap[key]; ok {
		// unlink
		item.prev.next = item.next
		item.next.prev = item.prev
		// link item after head
		item.next = this.head.next
		item.prev = this.head
		this.head.next = item
		item.next.prev = item
		return item.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if item, ok := this.keyItemMap[key]; ok {
		// update value
		item.value = value
		// unlink
		item.prev.next = item.next
		item.next.prev = item.prev
		// link item after head
		item.next = this.head.next
		item.prev = this.head
		this.head.next = item
		item.next.prev = item
		return
	}
	if this.currentItems < this.capacity {
		this.currentItems++
	} else {
		LeastUsedItem := this.tail.prev
		delete(this.keyItemMap, LeastUsedItem.key)
		this.tail.prev = LeastUsedItem.prev
		this.tail.prev.next = this.tail
	}
	// link item after head
	newItem := DoubleLinkedList{
		key:   key,
		value: value,
	}
	newItem.next = this.head.next
	newItem.prev = this.head
	this.head.next = &newItem
	newItem.next.prev = &newItem
	// update map
	this.keyItemMap[key] = &newItem
	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
