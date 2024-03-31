package main

import "fmt"

type Point struct {
	a int
	b int
}

type NilInterface interface {
}

func testAssert(nilInterface NilInterface) {
	fmt.Println(nilInterface)     // {1 2}
	b, ok := nilInterface.(Point) // 类型断言
	if ok {
		fmt.Println(b) // {1 2}
	} else {
		fmt.Println("类型断言失败")
	}
}

// switch type
func TypeJudge(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("idx %d val is %v, type is bool\n", i, x)
		case int:
			fmt.Printf("idx %d val is %v, type is int\n", i, x)
		case float64:
			fmt.Printf("idx %d val is %v, type is float32\n", i, x)
		case byte:
			fmt.Printf("idx %d val is %v, type is byte\n", i, x)
		case []int:
			fmt.Printf("idx %d val is %v, type is []int\n", i, x)
		default:
			fmt.Printf("idx %d val is %v, unknown type\n", i, x)
		}
	}
}

func main() {
	a := Point{1, 2}
	testAssert(a)
	var b byte = 'a'
	TypeJudge(1, 1.0, true, "A", []int{1, 2, 3}, b)
}
