package main

import (
	"fmt"
	"unsafe"
)

// bool使用
func main() {
	// bool 只能是true或者false
	var flag1 bool = true
	fmt.Println("flag1 is:", flag1)
	// bool 占用1个字节
	fmt.Println("bool size is:", unsafe.Sizeof(flag1))
}
