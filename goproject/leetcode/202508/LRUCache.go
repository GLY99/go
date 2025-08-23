package main

// https://leetcode.cn/problems/lru-cache-lcci/
// https://leetcode.cn/problems/lru-cache/description/

// LRU 缓存机制可以通过哈希表辅以双向链表实现，我们用一个哈希表和一个双向链表维护所有在缓存中的键值对。
// 双向链表按照被使用的顺序存储了这些键值对，靠近头部的键值对是最近使用的，而靠近尾部的键值对是最久未使用的。
// 哈希表即为普通的哈希映射（HashMap），通过缓存数据的键映射到其在双向链表中的位置。

type LRUCache struct {
	size       int // 当前的数据数量
	capacity   int // 整体容量
	cache      map[int]*LinkNode
	head, tail *LinkNode
}

type LinkNode struct {
	key, value int
	prev, next *LinkNode
}

func initLinkNode(key, value int) *LinkNode {
	return &LinkNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    map[int]*LinkNode{},
		head:     &LinkNode{key: 0, value: 0},
		tail:     &LinkNode{key: 0, value: 0},
	}
	// 将头尾节点连接起来
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	// 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	// 写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。
	// 当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间
	if _, ok := this.cache[key]; !ok {
		// 如果key不在缓存中，那么初始化一个新的节点，将其设置为最近使用的元素，如果size大于容量就删除最后面的节点
		node := initLinkNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removedNode := this.removeToTail()
			delete(this.cache, removedNode.key)
			this.size--
		}
	} else {
		// 如果key存在缓存中，直接替换新的val并且将其设置为最近使用元素
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) removeNode(node *LinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
}

func (this *LRUCache) addToHead(node *LinkNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *LinkNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeToTail() *LinkNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
