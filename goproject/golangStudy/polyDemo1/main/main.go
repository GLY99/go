package main

import "fmt"

// 定义一个Usb的接口
type Usb interface {
	Start()
	Stop()
}

// 定义Phone结构体，实现Usb接口的所有方法
type Phone struct {
	Name string
}

func (phone *Phone) Start() {
	fmt.Println("phone start work")
}

func (phone *Phone) Stop() {
	fmt.Println("phone stop work")
}

func (phone *Phone) Call() {
	fmt.Println("phone call")
}

// 定义Camera结构体，实现Usb接口的所有方法
type Camera struct {
	Name string
}

func (camera *Camera) Start() {
	fmt.Println("camera start work")
}

func (camera *Camera) Stop() {
	fmt.Println("camera stop work")
}

// 定义Computer结构体，实现Working方法接收一个Usb接口作为参数，所有实现了接口方法的结构体对象都可以传入
type Computer struct {
}

func (computer Computer) Working(usb Usb) {
	usb.Start()
	// 使用接口类型断言，如果是*Phone类型，赋值给phone, 然后调用Call()
	if phone, ok := usb.(*Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	var usbArr [3]Usb
	usbArr[0] = &Phone{"xiaomi"}
	usbArr[1] = &Phone{"vivo"}
	usbArr[2] = &Camera{"niko"}
	computer := Computer{}
	for _, usb := range usbArr {
		computer.Working(usb)
	}
}
