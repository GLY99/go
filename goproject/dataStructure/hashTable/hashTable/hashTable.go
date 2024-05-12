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
