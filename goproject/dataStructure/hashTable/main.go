package main

import (
	"dataStructure/hashTable/hashTable"
	"fmt"
)

func main() {
	myHashTable := hashTable.NewHashTable(1)
	fmt.Println(myHashTable.List())
	myHashTable.Insert(1)
	myHashTable.Insert(0)
	myHashTable.Insert(3)
	myHashTable.Insert(2)
	fmt.Println(myHashTable.List())
}
