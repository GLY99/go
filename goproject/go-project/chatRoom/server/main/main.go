package main

import (
	"chatRoom/server/model"
	"chatRoom/server/myRedis"
	"chatRoom/server/process"
	"fmt"
	"net"
	"time"
)

func processMsg(conn net.Conn) {
	defer conn.Close()
	// 读客户端发送的信息
	process := &process.Processor{Conn: conn}
	err := process.Process()
	if err != nil {
		fmt.Printf("服务器和客户端协程错误, err=%v", err)
		return
	}
}

func initGlobalUserDao() {
	model.GlobalUserDao = model.NewUserDao(myRedis.GlobalRedisPool)
}

func main() {
	myRedis.InitPool("127.0.0.1:6379", 16, 0, 300*time.Second)
	initGlobalUserDao()
	process.InitGlobalUserMgr()
	fmt.Println("服务器再8848端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8848")
	if err != nil {
		fmt.Printf("创建监听失败,err=%v", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("listen.Accept() err=%v", err)
		}
		go processMsg(conn)
	}
}
