package main

import "fmt"

func main() {
	var arr1 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr1)
	// 第二维不可以省略
	var arr2 = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2)
	// 二位数组遍历
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			fmt.Printf("%d", arr1[i][j])
		}
		fmt.Println()
	}

	for _, v1 := range arr2 {
		for _, v2 := range v1 {
			fmt.Printf("%d", v2)
		}
		fmt.Println()
	}
}
