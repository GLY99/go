package main

import (
	"dataStructure/hashTable/hashTable"
	"fmt"
)

func main() {
	myHashTable := hashTable.NewHashTable(5)
	fmt.Println(myHashTable.List())
	for i := 1; i < 20; i++ {
		myHashTable.Insert(i)
	}
	fmt.Println(myHashTable.List())
	for i := 5; i <= 15; i += 5 {
		myHashTable.Delete(i)
	}
	fmt.Println(myHashTable.List())
	for i := 5; i <= 15; i += 5 {
		myHashTable.Insert(i)
	}
	fmt.Println(myHashTable.List())
	node := myHashTable.Get(15)
	fmt.Printf("%p\n", node)
	myHashTable.Update(15, 100)
	fmt.Println(myHashTable.List())
}
