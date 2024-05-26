package process

import (
	"chatRoom/common/message"
	"chatRoom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (processor *Processor) ServerProcessMessage(msg *message.Message) (err error) {
	switch msg.Type {
	case message.LoginMsgType:
		// 处理登录
		userProcess := &UserProcess{Conn: processor.Conn}
		err = userProcess.ServerProcessLogin(msg)
	case message.RegisterMsgType:
		// 处理注册
	default:
		fmt.Println("消息类型不存在,无法处理...")
	}
	return
}

func (process *Processor) Process() (err error) {
	transfer := utils.Transfer{Conn: process.Conn}
	for {
		var msg message.Message
		msg, err = transfer.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("客户端关闭了连接,服务端也退出\n")
				return
			} else {
				fmt.Printf("transfer.ReadPkg() fail, err=%v", err)
				return
			}
		}
		fmt.Printf("server receive msg %v\n", msg)
		err = process.ServerProcessMessage(&msg)
		if err != nil {
			return
		}
	}
}
