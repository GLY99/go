package main

import "fmt"

func SetWay(myMap *[8][7]int, i int, j int) bool {
	if (*myMap)[6][5] == 2 {
		return true
	} else {
		if (*myMap)[i][j] == 0 {
			(*myMap)[i][j] = 2
			if SetWay(myMap, i+1, j) {
				return true
			} else if SetWay(myMap, i, j+1) {
				return true
			} else if SetWay(myMap, i-1, j) {
				return true
			} else if SetWay(myMap, i, j-1) {
				return true
			} else {
				(*myMap)[i][j] = 3
				return false
			}
		} else {
			return false
		}
	}
}

func main() {
	// map[i][j] == 1表示是一个墙
	// map[i][j] == 0表示没有走过的路
	// map[i][j] == 2表示是一个通路
	// map[i][j] == 3表示是一个曾经走过但是不通的路
	var myMap [8][7]int
	for i := 0; i < len(myMap); i++ {
		for j := 0; j < len(myMap[0]); j++ {
			if i == 0 || i == len(myMap)-1 {
				myMap[i][j] = 1
			} else {
				if j == 0 || j == len(myMap[0])-1 {
					myMap[i][j] = 1
				}
			}
		}
	}
	for _, arr := range myMap {
		fmt.Println(arr)
	}
	fmt.Println(SetWay(&myMap, 1, 1))
	for _, arr := range myMap {
		fmt.Println(arr)
	}
}
