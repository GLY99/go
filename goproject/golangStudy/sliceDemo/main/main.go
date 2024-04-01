package main

import "fmt"

func add1(numSlice []int) []int {
	return append(numSlice, 0)
}

func add2(numSlice []int) {
	numSlice = append(numSlice, 0)
}

func add3(numSlice *[]int) {
	*numSlice = append(*numSlice, 0)
}

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	slice := arr[1:3]                         // slice指向arr数组的下标为[1, 3)的元素
	fmt.Printf("arr=%v\n", arr)               // arr=[1 2 3 4 5]
	fmt.Printf("slice的元素是%v\n", slice)        // slice的元素是[2 3]
	fmt.Printf("slice的元素个数是%d\n", len(slice)) // slice的元素个数是2
	fmt.Printf("slice的容量是%d\n", cap(slice))   // slice的容量是4, 容量可以动态变化

	// slice是引用类型，修改数组和slice都会互相影响
	// slice从底层说是一个结构体
	// type slice struct {
	// 	ptr *[2]int
	// 	len int
	// 	cap int
	// }
	slice[0] = -1
	arr[2] = -1
	fmt.Printf("arr=%v\n", arr)        // arr=[1 -1 -1 4 5]
	fmt.Printf("slice的元素是%v\n", slice) // slice的元素是[-1 -1]

	// 切片的使用
	// 方式1 定义一个切片，然后让他去引用一个数组
	// 方式2 使用make创建切片 var param = make([]int, size, [cap]) cap可选 cap >= size， cap默认等于size
	var slice1 = make([]int, 4) // 默认采用默认值
	slice1[0] = 1
	slice1[1] = 2
	fmt.Printf("slice的元素是%v\n", slice1)        // slice的元素是[1 2 0 0]
	fmt.Printf("slice的元素个数是%d\n", len(slice1)) // slice的元素个数是4
	fmt.Printf("slice的容量是%d\n", cap(slice1))   // slice的容量是4
	slice1 = make([]int, 4, 10)
	slice1[0] = 1
	slice1[1] = 2
	fmt.Printf("slice的元素是%v\n", slice1)        // slice的元素是[1 2 0 0]
	fmt.Printf("slice的元素个数是%d\n", len(slice1)) // slice的元素个数是4
	fmt.Printf("slice的容量是%d\n", cap(slice1))   // slice的容量是10
	// 方式3 定义切片时直接使用数组
	slice2 := []int{1, 2, 3}
	fmt.Printf("slice的元素是%v\n", slice2)        // slice的元素是[1 2 3]
	fmt.Printf("slice的元素个数是%d\n", len(slice2)) // slice的元素个数是3
	fmt.Printf("slice的容量是%d\n", cap(slice2))   // slice的容量是3

	// 切片的遍历
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d\n", slice[i])
	}

	for idx, val := range slice1 {
		fmt.Printf("slice1[%d]=%d\n", idx, val)
	}

	// 通过append对切片进行动态增加, 当append后的内容超过cap会重新创建一个新数组来引用
	slice3 := make([]int, 2, 10)
	fmt.Printf("slice3的地址是%v\n", &slice3[0]) // slice3的地址是0xc000012190
	slice3 = append(slice3, 1)
	fmt.Printf("slice3的地址是%v\n", &slice3[0]) // slice3的地址是0xc000012190
	slice3 = append(slice3, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("slice3的地址是%v\n", &slice3[0]) // slice3的地址是0xc0000240a0
	fmt.Printf("slice3的元素是%v\n", slice3)     // slice3的元素是[0 0 1 2 3 4 5 6 7 8 9]
	var slice4 []int
	slice4 = append(slice4, slice3...)
	fmt.Printf("slice4的元素是%v\n", slice4) // slice4的元素是[1 2 3]

	// copy切片
	slice5 := []int{1, 2, 3}
	var sliceCopy []int = []int{4}
	count := copy(sliceCopy, slice5)
	fmt.Printf("copy了%d个元素\n", count)              // copy了1个元素
	fmt.Printf("slice5的地址是%v\n", &slice5[0])       // slice5的地址是0xc00000e240
	fmt.Printf("sliceCopy的地址是%v\n", &sliceCopy[0]) // sliceCopy的地址是0xc00000e258
	fmt.Println(slice5, sliceCopy)                 // [1 2 3] [4]

	// 在其他函数里面对切片append虽然传递了地址过去，但是不会影响外面的值
	// 原因是size和cap是传递的值,只有切片里面指向数组首地址的指针传递的是地址。
	numSlice1 := make([]int, 0, 5)
	fmt.Println(numSlice1, len(numSlice1), cap(numSlice1)) // [] 0 5
	numSlice1 = add1(numSlice1)
	fmt.Println(numSlice1, len(numSlice1), cap(numSlice1)) // [0] 1 5
	add3(&numSlice1)
	fmt.Println(numSlice1, len(numSlice1), cap(numSlice1)) // [0， 0] 2 5

	numSlice2 := make([]int, 0, 5)
	fmt.Println(numSlice2, len(numSlice2), cap(numSlice2)) // [] 0 5
	add2(numSlice2)
	fmt.Println(numSlice2, len(numSlice2), cap(numSlice2)) // [] 0 5
}
