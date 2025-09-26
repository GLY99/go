package main

import "reflect"

// https://leetcode.cn/problems/odd-string-difference/?envType=problem-list-v2&envId=array

func getOdd(words string) []int {
	n := len(words)
	diff := make([]int, n-1)
	for i := 0; i+1 < n; i++ {
		diff[i] = int(words[i+1]) - int(words[i])
	}
	return diff
}

func oddString(words []string) string {
	n := len(words)
	diff0 := getOdd(words[0])
	diff1 := getOdd(words[1])
	if reflect.DeepEqual(diff0, diff1) {
		for i := 2; i < n; i++ {
			if !reflect.DeepEqual(diff0, getOdd(words[i])) {
				return words[i]
			}
		}
	}
	if reflect.DeepEqual(diff0, getOdd(words[2])) {
		return words[1]
	}
	return words[0]
}

func oddString(words []string) string {
    diff0 := get(words[0])
    diff1 := get(words[1])
    if reflect.DeepEqual(diff0, diff1) {
        for i:= 2; i < len(words); i++ {
            if !reflect.DeepEqual(diff0, get(words[i])) {
                return words[i]
            }
        }
    }
    if reflect.DeepEqual(diff0, get(words[2])) {
        return words[1]
    }
    return words[0]
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/odd-string-difference/solutions/2281743/chai-zhi-shu-zu-bu-tong-de-zi-fu-chuan-b-3rhg/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。