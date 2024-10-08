package main

func destCity(paths [][]string) string {
	var mapping map[string]*[2]int = make(map[string]*[2]int)
	for _, path := range paths {
		cA := path[0]
		cB := path[1]
		// 出
		if _, ok := mapping[cA]; !ok {
			mapping[cA] = &[2]int{0, 1}
		} else {
			mapping[cA][1] += 1
		}
		// 进
		if _, ok := mapping[cB]; !ok {
			mapping[cB] = &[2]int{1, 0}
		} else {
			mapping[cB][0] += 1
		}
	}
	ans := ""
	for k, v := range mapping {
		if (*v)[1] == 0 {
			ans = k
			break
		}
	}
	return ans
}
