package main

func finalString(s string) string {
	queue := []rune{}
	head := true
	for _, c := range s {
		if c != 'i' {
			if head {
				queue = append(queue, c)
			} else {
				queue = append([]rune{c}, queue...)
			}
		} else {
			head = !head
		}
	}
	if !head {
		reserve(queue)
	}
	return string(queue)
}

func reserve(arr []rune) {
	start := 0
	end := len(arr) - 1
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
