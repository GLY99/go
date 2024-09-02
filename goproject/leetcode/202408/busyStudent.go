package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var count int
	count = 0
	length := len(startTime)
	for i := 0; i < length; i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			count++
		}
	}
	return count
}
