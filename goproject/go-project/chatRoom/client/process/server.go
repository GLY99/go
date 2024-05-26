package process

import (
	"chatRoom/client/utils"
	"fmt"
	"net"
	"os"
)

// 显示登录成功后的界面
func ShowMenu() {
	for {
		fmt.Println("----------恭喜xxx登录成功----------")
		fmt.Println("----------1.显示在线用户列表----------")
		fmt.Println("----------2.发送消息----------")
		fmt.Println("----------3.消息列表----------")
		fmt.Println("----------4.退出系统----------")
		fmt.Println("请选择1-4:")
		var key int
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("显示在线用户列表")
		case 2:
			fmt.Println("发送消息")
		case 3:
			fmt.Println("消息列表")
		case 4:
			fmt.Println("退出系统")
			os.Exit(1)
		default:
			fmt.Println("你输入的选项不对")
		}
	}
}

func ProcessServerMsg(conn net.Conn) {
	transfer := utils.Transfer{Conn: conn}
	for {
		fmt.Printf("客户端正在等待服务器消息\n")
		msg, err := transfer.ReadPkg()
		if err != nil {
			fmt.Printf("client transfer.ReadPkg() fail, err=%v", err)
			return
		}
		fmt.Printf("receive server msg: %v\n", msg)
	}
}
