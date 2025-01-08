package main

import "strconv"

func largestGoodInteger(num string) string {
	ans := -1
	for i := 0; i < len(num)-2; i++ {
		tmp_num, _ := strconv.Atoi(num[i : i+3])
		if num[i] == num[i+1] && num[i+1] == num[i+2] && ans < tmp_num {
			ans = tmp_num
		}
	}
	if ans == -1 {
		return ""
	} else {
		if ans == 0 {
			return "000"
		} else {
			return strconv.Itoa(ans)
		}
	}
}
