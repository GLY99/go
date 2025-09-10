package main

// https://leetcode.cn/problems/minimum-number-of-people-to-teach/?envType=daily-question&envId=2025-09-10

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	// 记录不能正常交流的人
	mapping := make(map[int]interface{})
	for _, friendship := range friendships {
		// 判断每组好友能否正常交流
		flag := false
		tmpMapping := make(map[int]interface{})
		for _, lang := range languages[friendship[0]-1] {
			tmpMapping[lang] = struct{}{}
		}
		// 如果语言存在交集就能交流
		for _, lang := range languages[friendship[1]-1] {
			if _, ok := tmpMapping[lang]; ok {
				flag = true
				break
			}
		}
		if !flag {
			mapping[friendship[0]] = struct{}{}
			mapping[friendship[1]] = struct{}{}
		}
	}

	maxCount := 0
	// 记录不能正常交流的人中会不同语言的人的数量
	counter := make([]int, n+1)
	for person, _ := range mapping {
		for _, lang := range languages[person-1] {
			counter[lang]++
			maxCount = max(maxCount, counter[lang])
		}
	}
	// 不能正常的交流的人减去最多人会的语言人数就是需要教的人数
	return len(mapping) - maxCount
}
