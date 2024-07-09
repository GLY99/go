package main

func titleToNumber(columnTitle string) int {
	length := len(columnTitle)
	ans := 0
	for i, mult := length-1, 1; i >= 0; i-- {
		k := columnTitle[i] - 'A' + 1
		ans += int(k) * mult
		mult *= 26
	}
	return ans
}
