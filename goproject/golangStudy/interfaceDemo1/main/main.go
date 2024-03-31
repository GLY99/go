package main

import "fmt"

// 定义一个Usb的接口
type Usb interface {
	Start()
	Stop()
}

// 定义Phone结构体，实现Usb接口的所有方法
type Phone struct {
}

func (phone *Phone) Start() {
	fmt.Println("phone start work")
}

func (phone *Phone) Stop() {
	fmt.Println("phone stop work")
}

// 定义Camera结构体，实现Usb接口的所有方法
type Camera struct {
}

func (camera *Camera) Start() {
	fmt.Println("camera start work")
}

func (camera *Camera) Stop() {
	fmt.Println("camera stop work")
}

// 任何自定义类型都可以实现接口的方法
type integer int

func (i *integer) Start() {
	fmt.Println("i start")
}

func (i *integer) Stop() {
	fmt.Println("i stop")
}

// 定义Computer结构体，实现Working方法接收一个Usb接口作为参数，所有实现了接口方法的结构体对象都可以传入
type Computer struct {
}

func (computer Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	var i integer = 1
	computer.Working(&phone)  // phone start work phone stop work
	computer.Working(&camera) // camera start work camera stop work
	var usbInterface1 Usb = &i
	usbInterface1.Start() // i start
	usbInterface1.Stop()  // i stop
}
