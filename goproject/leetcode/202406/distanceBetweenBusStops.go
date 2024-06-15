package main

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	length := len(distance)
	idx := start
	d1 := 0
	for idx != destination {
		d1 += distance[idx]
		idx = (idx + 1) % length
	}
	idx = destination
	d2 := 0
	for idx != start {
		d2 += distance[idx]
		idx = (idx + 1) % length
	}
	return min(d1, d2)
}
