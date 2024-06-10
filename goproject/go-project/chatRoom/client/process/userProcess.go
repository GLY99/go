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
	loginMsg.UserPwd = passWd
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
	fmt.Printf("send login user data %s\n", data)
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Printf("send login user data fail, err=%v", err)
		return
	}
	msg, err = transfer.ReadPkg()
	if err != nil {
		fmt.Printf("read login rsp msg fail, err=%v", err)
		return
	}
	var loginRspMsg message.LoginRspMsg
	err = json.Unmarshal([]byte(msg.Data), &loginRspMsg)
	if err != nil {
		fmt.Printf("unmarshal login rsp msg fail, err=%v", err)
		return
	}
	if loginRspMsg.Code == 200 {
		// 启动协程监听服务器发送的数据
		fmt.Println("登录成功！")
		go ProcessServerMsg(conn)
		ShowMenu()
	} else {
		fmt.Printf("user login fail, err=%v\n", loginRspMsg.Msg)
		err = errors.New(loginRspMsg.Msg)
		return
	}
	return
}

func (userProcess *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	// 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8848")
	if err != nil {
		fmt.Printf("net.Dial fail, err=%v", err)
		return
	}
	// 延时关闭
	defer conn.Close()
	// 构造需要发送给server的注册信息
	var msg message.Message
	var registerMsg message.RegisterMsg
	msg.Type = message.RegisterMsgType
	registerMsg.UserId = userId
	registerMsg.UserPwd = userPwd
	registerMsg.UserName = userName
	data, err := json.Marshal(registerMsg)
	if err != nil {
		fmt.Printf("marshal registerMsg fail, err=%v\n", err)
		return
	}
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Printf("marshal msg fail when register, err=%v\n", err)
		return
	}
	// 发送数据
	transfer := &utils.Transfer{Conn: conn}
	fmt.Printf("send register user data %s\n", data)
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Printf("send register user data fail, err=%v", err)
		return
	}
	// 接受服务器响应
	msg, err = transfer.ReadPkg()
	if err != nil {
		fmt.Printf("read register rsp msg fail, err=%v", err)
		return
	}
	var registerRspMsg message.RegisterRspMsg
	err = json.Unmarshal([]byte(msg.Data), &registerRspMsg)
	if err != nil {
		fmt.Printf("unmarshal register rsp msg fail, err=%v", err)
		return
	}
	if registerRspMsg.Code == 200 {
		fmt.Println("注册成功！")
	} else {
		fmt.Printf("register user fail, err=%v\n", registerRspMsg.Msg)
		err = errors.New(registerRspMsg.Msg)
		return
	}
	return
}
