package main

func minTimeToType(word string) int {
	res := 0
	prev := 0
	length := len(word)
	for i := 0; i < length; i++ {
		cur := int(word[i] - 'a')
		res += 1 + min(abs(cur, prev), 26-abs(cur, prev))
		prev = cur
	}
	return res
}

// func abs(num1, num2 int) int {
// 	var res int
// 	if num1 > num2 {
// 		res = num1 - num2
// 	} else {
// 		res = num2 - num1
// 	}
// 	return res
// }
