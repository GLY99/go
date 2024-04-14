package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("client conn 127.0.0.1:8888 fail, err=%v\n", err)
		return
	} else {
		fmt.Printf("conn server %v, server addr is %v\n", conn, conn.RemoteAddr())
		defer conn.Close()
	}

	reader := bufio.NewReader(os.Stdin)
	// 从终端读取用户输入并且发送
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read string fail, err=%v", err)
			return
		}
		str = strings.TrimRight(str, " \r\n")
		if str == "exit" {
			fmt.Printf("client exit!\n")
			break
		}
		n, err := conn.Write([]byte(str))
		if err != nil {
			fmt.Printf("send data fail, err=%v\n", err)
			return
		} else {
			fmt.Printf("succ send %d byte data\n", n)
		}
	}
}
