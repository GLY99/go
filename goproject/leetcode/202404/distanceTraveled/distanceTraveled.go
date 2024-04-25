package main

import "fmt"

func distanceTraveled(mainTank int, additionalTank int) int {
	res := 0
	for mainTank >= 5 {
		mainTank -= 5
		res += 50
		if additionalTank > 0 {
			additionalTank--
			mainTank++
		}
	}
	return res + mainTank*10
}

func main() {
	fmt.Println(distanceTraveled(5, 1))
}
