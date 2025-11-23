package main

import (
	"math/rand"
)

func randomBrithday(count int) (sameCounter, diffCounter int) {
	// 模拟count次
	for i := 0; i < count; i++ {
		set := make(map[int]interface{})
		hasSame := false
		// 每次模拟50个人的生日
		for j := 0; j < 50; j++ {
			d := rand.Intn(365) + 1
			if _, ok := set[d]; ok {
				// 有相同的生日就退出，不用模拟了
				hasSame = true
				break
			} else {
				set[d] = struct{}{}
			}
		}
		// 根据是否存在相同生日的人给计数器+1
		if !hasSame {
			diffCounter++
		} else {
			sameCounter++
		}
	}
	return
}
