package main

import "fmt"

type A struct {
	Num int
}

type inter int

func (a A) test1() { // a的值是从a.传进去的
	a.Num = 100
	fmt.Println(a.Num) // 100
}

func (a *A) test2() {
	(*a).Num = 100
	fmt.Println(a.Num) // 100
}

func (a A) test3(num int) int {
	return a.Num + num
}

func (a *A) String() string {
	str := fmt.Sprintf("A.Num=%v", (*a).Num)
	return str
}

func (num *inter) test1() {
	*num++
}

func main() {
	var a A
	a.Num = 99
	a.test1()          // 传递值
	fmt.Println(a.Num) // 99
	(&a).test2()       // 传递地址
	fmt.Println(a.Num) // 100

	// 方法参数传递
	res := a.test3(100)
	fmt.Println(res) // 200

	// 除了struct 任何自定义类型都可以有方法
	var num inter = 1
	(&num).test1()
	fmt.Println(num) // 2

	// 如果一个结构体实现了String()，在println时会调用这个方法输出
	// 因为String绑定的时a的地址，所以在println时要输出地址
	fmt.Println(&a) // A.Num=100

	// 对于方法接受者是值类型可以用值类型传递也可以用指针类型传递，反之亦然
	// 具体是传递的指针还是值，看方法定义的类型，如果方法定义的是值类型，即使用指针调用依然是值拷贝，反之亦然
	var b A = A{1}
	b.test1()          // 100
	fmt.Println(b.Num) // 1
	(&b).test1()       // 100 这里用指针调用，但是因为方法定义的是值类型，所以实际传递的还是值
	fmt.Println(b.Num) // 1

	b.test2()          // 100 这里用值调用，但是因为方法定义的是指针类型，所以实际传递的是指针
	fmt.Println(b.Num) // 100
	b.Num = 1
	(&b).test2()       // 100
	fmt.Println(b.Num) // 100
}
