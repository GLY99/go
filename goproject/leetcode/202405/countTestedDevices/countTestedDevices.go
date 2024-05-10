package main

func countTestedDevices(batteryPercentages []int) int {
	count := 0
	legnth := len(batteryPercentages)
	for i, batteryPercentage := range batteryPercentages {
		if batteryPercentage > 0 {
			count++
			for j := i + 1; j < legnth; j++ {
				batteryPercentages[j] = max(0, batteryPercentages[j]-1)
			}
		}
	}
	return count
}

func main() {
	countTestedDevices([]int{1, 2})
}
