package main

import (
	"container/list"
	"fmt"
)

const base = 769

type MyHashSet struct {
	data []list.List
}

func Constructor() MyHashSet {
	return MyHashSet{data: make([]list.List, base)}
}

func (myHashSet *MyHashSet) hash(key int) int {
	return key % base
}

func (myHashSet *MyHashSet) Add(key int) {
	if !myHashSet.Contains(key) {
		hashKey := myHashSet.hash(key)
		myHashSet.data[hashKey].PushBack(key)
	}
}

func (myHashSet *MyHashSet) Remove(key int) {
	if myHashSet.Contains(key) {
		hashKey := myHashSet.hash(key)
		for node := myHashSet.data[hashKey].Front(); node != nil; node = node.Next() {
			if node.Value.(int) == key {
				myHashSet.data[hashKey].Remove(node)
			}
		}
	}
}

func (myHashSet *MyHashSet) Contains(key int) bool {
	hashKey := myHashSet.hash(key)
	for node := myHashSet.data[hashKey].Front(); node != nil; node = node.Next() {
		if node.Value.(int) == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
func main() {
	myHashSet := Constructor()
	myHashSet.Add(1)
	myHashSet.Add(2)
	myHashSet.Remove(2)
	fmt.Println(myHashSet.Contains(1))
	fmt.Println(myHashSet.Contains(2))
}
