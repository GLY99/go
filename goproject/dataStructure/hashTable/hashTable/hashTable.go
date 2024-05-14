package hashTable

type Node struct {
	Val  int
	Next *Node
}

type HashTable struct {
	Size int
	arr  []*Node
}

func NewHashTable(size int) HashTable {
	return HashTable{Size: size, arr: make([]*Node, size)}
}

func (hashTable *HashTable) hash(num int) int {
	return num % hashTable.Size
}

func (hashTable *HashTable) Insert(num int) {
	node := Node{Val: num}
	hash_code := hashTable.hash(num)
	if hashTable.arr[hash_code] == nil {
		hashTable.arr[hash_code] = &node
	} else {
		head := hashTable.arr[hash_code]
		if head.Val >= node.Val {
			node.Next = head
			hashTable.arr[hash_code] = &node
		} else {
			curNode := head
			for curNode.Next != nil {
				if curNode.Next.Val >= node.Val {
					node.Next = curNode.Next
					curNode.Next = &node
					return
				}
				curNode = curNode.Next
			}
			curNode.Next = &node
		}
	}
}

func (hashTable *HashTable) Delete(num int) {
	hash_code := hashTable.hash(num)
	if hashTable.arr[hash_code] == nil {
		return
	} else {
		head := hashTable.arr[hash_code]
		if head.Val == num {
			hashTable.arr[hash_code] = head.Next
			return
		} else {
			perNode := head
			curNode := head.Next
			for curNode != nil {
				if curNode.Val == num {
					perNode.Next = curNode.Next
				}
				curNode = curNode.Next
			}
		}
	}
}

func (hashTable *HashTable) Get(num int) *Node {
	hash_code := hashTable.hash(num)
	if hashTable.arr[hash_code] == nil {
		return nil
	} else {
		node := hashTable.arr[hash_code]
		for node != nil {
			if node.Val == num {
				return node
			}
			node = node.Next
		}
		return nil
	}
}

func (hashTable *HashTable) Update(num int, newNum int) {
	node := hashTable.Get(num)
	if node == nil {
		return
	}
	hashTable.Delete(num)
	hashTable.Insert(newNum)
}

func (hashTable *HashTable) List() map[int][]int {
	myMap := make(map[int][]int, hashTable.Size)
	for idx, node := range hashTable.arr {
		for node != nil {
			myMap[idx] = append(myMap[idx], node.Val)
			node = node.Next
		}
	}
	return myMap
}
