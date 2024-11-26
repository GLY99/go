package main

func numberOfAlternatingGroups(colors []int) int {
	ans := 0
	for i := 0; i < len(colors); i++ {
		if i == 0 {
			if colors[i] != colors[len(colors)-1] && colors[i] != colors[i+1] {
				ans++
			}
		} else {
			if colors[i] != colors[i-1] && colors[i] != colors[(i+1)%len(colors)] {
				ans++
			}
		}
	}
	return ans
}
