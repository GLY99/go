package main

func decrypt(code []int, k int) []int {
	length := len(code)
	res := make([]int, length)
	if k == 0 {
		return res
	}
	for idx, _ := range code {
		new_code := 0
		count := 0
		if k > 0 {
			i := (idx + 1) % length
			for count < k {
				new_code += code[i]
				i = (i + 1) % length
				count++
			}
		} else {
			i := idx - 1
			if i < 0 {
				i = length - 1
			}
			for count < -k {
				new_code += code[i]
				i--
				if i < 0 {
					i = length - 1
				}
				count++
			}
		}
		res[idx] = new_code
	}
	return res
}
