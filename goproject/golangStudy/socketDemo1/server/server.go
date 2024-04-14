package main

import (
	"fmt"
	"io"
	"net"
)

func ProcMessage(conn net.Conn) {
	defer func() {
		conn.Close()
		fmt.Printf("finash process client %v message\n", conn.RemoteAddr())
	}()
	for {
		buf := make([]byte, 1024)
		// 等待client发送数据，如果client没有发送东西，这里会一直阻塞
		fmt.Printf("server waitting client %v send data\n", conn.RemoteAddr())
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				// 客户端主动关闭连接
				fmt.Printf("client %v close conn\n", conn.RemoteAddr())
			} else {
				fmt.Printf("server read data fail, err=%v", err)
			}
			return
		}
		fmt.Printf("%s\n", string(buf[:n]))
	}
}

func main() {
	fmt.Printf("server start listen...\n")
	// 使用tcp协议监听本地8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Printf("server start listen fail, err=%v\n", err)
		return
	}
	// 延时关闭
	defer listen.Close()

	// 等待客户端连接
	for {
		fmt.Printf("wait client connect...\n")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept err=%v\n", err)
		} else {
			fmt.Printf("accept conn %v, client addr is %s\n", conn, conn.RemoteAddr())
		}
		// 这里起协程处理客户端请求
		go ProcMessage(conn)
	}
}
