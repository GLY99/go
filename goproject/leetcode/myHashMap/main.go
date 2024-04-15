package main

import "container/list"

const base = 789

type entry struct {
	key int
	val int
}

type MyHashMap struct {
	data []list.List
}

func Constructor() MyHashMap {
	return MyHashMap{data: make([]list.List, base)}
}

func (myHashMap *MyHashMap) hash(k int) int {
	return k % base
}

func (myHashMap *MyHashMap) Put(key int, value int) {
	hash_key := myHashMap.hash(key)
	for e := myHashMap.data[hash_key].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			e.Value = entry{key, value}
			return
		}
	}
	myHashMap.data[hash_key].PushBack(entry{key, value})
}

func (myHashMap *MyHashMap) Get(key int) int {
	hash_key := myHashMap.hash(key)
	for e := myHashMap.data[hash_key].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			return et.val
		}
	}
	return -1
}

func (myHashMap *MyHashMap) Remove(key int) {
	hash_key := myHashMap.hash(key)
	for e := myHashMap.data[hash_key].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			myHashMap.data[hash_key].Remove(e)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
