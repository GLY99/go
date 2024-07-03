package main

func removeTrailingZeros(num string) string {
	length := len(num)
	i := length - 1
	for ; i >= 0; i-- {
		if num[i] != '0' {
			break
		}
	}
	return num[0 : i+1]
}
