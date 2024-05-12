package main

import (
	"dataStructure/expStack/exp"
	"fmt"
)

func main() {
	str := "10+20*30-1"
	res, err := exp.Exp(str)
	fmt.Printf("%v, %v", res, err)
}
