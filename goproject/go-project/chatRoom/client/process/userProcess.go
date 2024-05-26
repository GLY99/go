package process

import (
	"chatRoom/client/utils"
	"chatRoom/common/message"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type UserProcess struct{}

func (userProcess *UserProcess) Login(userId int, passWd string) (err error) {
	// 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8848")
	if err != nil {
		fmt.Printf("net.Dial fail, err=%v", err)
		return
	}
	// 延时关闭
	defer conn.Close()
	var msg message.Message
	msg.Type = message.LoginMsgType
	var loginMsg message.LoginMsg
	loginMsg.UserId = userId
	loginMsg.Passwd = passWd
	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Printf("json.Marshal(loginMsg) err, err=%v", err)
		return
	}
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Printf(" json.Marshal(msg) err, err=%v", err)
		return
	}
	transfer := &utils.Transfer{Conn: conn}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Printf("transfer.WritePkg(data) err, err=%v", err)
		return
	}
	msg, err = transfer.ReadPkg()
	if err != nil {
		fmt.Printf("transfer.ReadPkg() err, err=%v", err)
		return
	}
	var loginRspMsg message.LoginRspMsg
	err = json.Unmarshal([]byte(msg.Data), &loginRspMsg)
	if err != nil {
		fmt.Printf("json.Unmarshal([]byte(msg.Data), &loginRspMsg) err, err=%v", err)
		return
	}
	if loginRspMsg.Code == 200 {
		// 启动协程监听服务器发送的数据
		go ProcessServerMsg(conn)
		ShowMenu()
	} else if loginRspMsg.Code == 500 {
		fmt.Printf("%v\n", loginRspMsg.Msg)
		err = errors.New(loginRspMsg.Msg)
		return
	}
	return

}
