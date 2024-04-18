package main

import (
	"fmt"
	"math"
)

func getArrAvgAndMinAndMax(arr *[8]float64) (avg float64, min float64, max float64) {
	length := len(*arr)
	avg = (*arr)[0] / float64(length)
	min = (*arr)[0]
	max = (*arr)[0]
	for i := 1; i < length; i++ {
		avg += (*arr)[i] / float64(length)
		if (*arr)[i] > max {
			max = (*arr)[i]
		}
		if (*arr)[i] < min {
			min = (*arr)[i]
		}
	}
	return
}

func getGreatestAndLowest(arr *[8]float64, avg float64) ([]int, []int) {
	length := len(*arr)
	slice1 := make([]int, 0)
	slice2 := make([]int, 0)

	for i := 0; i < length; i++ {
		tmp := math.Abs((*arr)[i] - avg)
		if len(slice1) == 0 || math.Abs((*arr)[slice1[0]]-avg) == tmp {
			slice1 = append(slice1, i)
		} else {
			if math.Abs((*arr)[slice1[0]]-avg) > tmp {
				slice1 = make([]int, 0)
				slice1 = append(slice1, i)
			}
		}
		if len(slice2) == 0 || math.Abs((*arr)[slice2[0]]-avg) == tmp {
			slice2 = append(slice2, i)
		} else {
			if math.Abs((*arr)[slice2[0]]-avg) < tmp {
				slice2 = make([]int, 0)
				slice2 = append(slice2, i)
			}
		}
	}
	return slice1, slice2
}

func main() {
	var arr = [8]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	avg, min, max := getArrAvgAndMinAndMax(&arr)
	fmt.Println(avg, min, max)
	greatSlice, lowSlice := getGreatestAndLowest(&arr, avg)
	fmt.Println(greatSlice, lowSlice)
}
