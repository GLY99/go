package main

import "fmt"

type structDemo struct {
	slice []int
	num   int
	s     string
}

func main() {
	// 1.struct类型的变量中存储的是值，而不是地址，在变量传递时会cpoy一份值传递，修改新旧变量可能互相影响（如果值内容是地址就会互相影响）
	// 2.struct类型的指针变量中存储的时地址，在变量传递是会cpoy一份地址传递，修改新旧变量都会互相影响
	// 3.struct的字段对应的值存储在连续的内存中
	// 4.struct变量的地址就是第一个字段的地址

	struct1 := structDemo{
		slice: make([]int, 1),
		num:   10,
		s:     "abce1234",
	}
	fmt.Printf("struct1 addr=%p, struct1.slice addr=%p, struct1.slice[0] addr=%p, struct1.num addr=%p, struct1.s addr=%p\n",
		&struct1, &struct1.slice, &struct1.slice[0], &struct1.num, &struct1.s,
	)
	// struct1 addr=0xc0000200f0, struct1.slice addr=0xc0000200f0, struct1.slice[0] addr=0xc00000a0a8,
	// struct1.num addr=0xc000020108, struct1.s addr=0xc000020110

	struct2 := struct1
	fmt.Printf("struct2 addr=%p, struct2.slice addr=%p, struct1.slice[0] addr=%p, struct2.num addr=%p, struct2.s addr=%p\n",
		&struct2, &struct2.slice, &struct2.slice[0], &struct2.num, &struct2.s,
	)
	// struct2 addr=0xc000020120, struct2.slice addr=0xc000020120, struct1.slice[0] addr=0xc00000a0a8,
	// struct2.num addr=0xc000020138, struct2.s addr=0xc000020140

	struct3 := &struct1
	fmt.Printf("struct3 addr=%p, struct3.slice addr=%p, struct1.slice[0] addr=%p, struct3.num addr=%p, struct3.s addr=%p\n",
		&struct3, &struct3.slice, &struct3.slice[0], &struct3.num, &struct3.s,
	)
	// struct3 addr=0xc00006c028, struct3.slice addr=0xc0000200f0, struct1.slice[0] addr=0xc00000a0a8,
	// struct3.num addr=0xc000020108, struct3.s addr=0xc000020110
}
