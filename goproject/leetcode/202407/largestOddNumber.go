package main

func largestOddNumber(num string) string {
	length := len(num)
	for i := length - 1; i >= 0; i-- {
		if (num[i]-'0')%2 != 0 {
			return string(num[:i+1])
		}
	}
	return ""
}
