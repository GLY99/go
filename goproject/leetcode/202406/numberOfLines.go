package main

func numberOfLines(widths []int, s string) []int {
	maxSize := 100
	lines, count := 1, 0
	for _, c := range s {
		count += widths[c-'a']
		if count > maxSize {
			count = widths[c-'a']
			lines++
		}
	}
	return []int{lines, count}
}
