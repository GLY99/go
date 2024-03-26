package main

import (
	"fmt"
	"golangStudy/encapsulateDemo/model"
)

func main() {
	p1 := model.NewPerson("xiaoming")
	p1.SetAge(18)
	p1.SetSal(30000)
	fmt.Println(p1.GetAge()) // 18
	fmt.Println(p1.GetSal()) // 30000
}
