package main

import "fmt"

type InterfaceA interface {
	testa()
}

type InterfaceB interface {
	testb()
}

type InterfaceC interface {
	InterfaceA
	InterfaceB
	testc()
}

type StructA struct {
}

func (a *StructA) testa() {
	fmt.Println("testa")
}

func (a *StructA) testb() {
	fmt.Println("testb")
}

type StructB struct {
}

func (b *StructB) testa() {
	fmt.Println("testa")
}

func (b *StructB) testb() {
	fmt.Println("testb")
}

func (b *StructB) testc() {
	fmt.Println("testc")
}

type T interface{}

func main() {
	var interfaceA InterfaceA
	var interfaceB InterfaceB
	var interfaceC InterfaceC
	var structA StructA
	// 结构体可以实现多个接口
	// 结构体A实现了接口A B的所有方法因此可以赋值给接口A B的变量, 但是没有实现接口C的方法因此不能赋值给接口C
	interfaceA = &structA
	interfaceB = &structA
	interfaceA.testa() // testa
	interfaceB.testb() // testb

	// 当一个接口继承了别的接口，结构体想要实现这个接口就要实现本接口和继承的接口所有方法
	var structB StructB
	interfaceA = &structB
	interfaceB = &structB
	interfaceC = &structB
	interfaceA.testa() // testa
	interfaceB.testb() // testb
	interfaceC.testc() // testc

	// 可以将任何值赋值给空接口
	var nullInterface1 T = structB
	var nullInterface2 interface{} = structB
	fmt.Println(nullInterface1, nullInterface2) // {} {}
}
