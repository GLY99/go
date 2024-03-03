package main

import "fmt"

func getPrimeNumber(num int) {
	count := 0
	for i := 1; i <= num; i++ {
		f := true
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				f = false
			}
		}
		if f {
			if count < 4 {
				fmt.Printf("%d ", i)
				count++
			} else {
				fmt.Printf("%d\n", i)
				count = 0
			}
		}
	}
}

func main() {
	getPrimeNumber(100)
}
