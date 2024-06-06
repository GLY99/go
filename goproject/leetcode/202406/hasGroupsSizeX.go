package main

func hasGroupsSizeX(deck []int) bool {
	myMap := make(map[int]int)
	for _, v := range deck {
		myMap[v]++
	}
	for x := 2; x <= len(deck); x++ {
		if len(deck)%x == 0 {
			flag := true
			for _, v := range myMap {
				if v%x != 0 {
					flag = false
				}
			}
			if flag {
				return true
			}
		}
	}
	return false
}
