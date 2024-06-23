package process

import (
	"chatRoom/common/message"
	"chatRoom/server/model"
	"chatRoom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

func (userProcess *UserProcess) ServerProcessLogin(msg *message.Message) (err error) {
	// 处理登录
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Printf("json.Unmarshal([]byte(msg.Data), &loginMsg) fail, err=%v\n", err)
		return err
	}
	var responseMsg message.Message
	responseMsg.Type = message.LoginRspMsgType
	var loginRspMsg message.LoginRspMsg
	user, err := model.GlobalUserDao.Login(loginMsg.UserId, loginMsg.UserPwd)
	if err != nil {
		fmt.Printf("user %v login fail, err is %v\n", loginMsg.UserId, err)
		if err == model.ErrorUserNotExists {
			loginRspMsg.Code = 404
			loginRspMsg.Msg = err.Error()
		} else if err == model.ErrorUserPwd {
			loginRspMsg.Code = 421
			loginRspMsg.Msg = err.Error()
		} else {
			loginRspMsg.Code = 500
			loginRspMsg.Msg = "unknown error!"
		}
	} else {
		loginRspMsg.Code = 200
		userProcess.UserId = loginMsg.UserId
		globalUserMgr.AddOnlieUser(userProcess)
		userProcess.NofityOthersOnlieUser(loginMsg.UserId, message.UserOnlie)
		for userId, _ := range globalUserMgr.onlieUsers {
			loginRspMsg.UserIds = append(loginRspMsg.UserIds, userId)
		}
		fmt.Printf("user %v 登录成功\n", user.UserId)
	}
	data, err := json.Marshal(loginRspMsg)
	if err != nil {
		fmt.Printf("json.Marshal(loginRspMsg) fail, err=%v\n", err)
		return err
	}
	responseMsg.Data = string(data)
	data, err = json.Marshal(responseMsg)
	if err != nil {
		fmt.Printf("json.Marshal(responseMsg) fail, err=%v\n", err)
		return err
	}
	transfer := &utils.Transfer{Conn: userProcess.Conn}
	err = transfer.WritePkg(data)
	return
}

func (userProcess *UserProcess) ServerProcessRegister(msg *message.Message) (err error) {
	// 处理注册
	var registerMsg message.RegisterMsg
	err = json.Unmarshal([]byte(msg.Data), &registerMsg)
	if err != nil {
		fmt.Printf("unmarshal register msg fail, err=%v", err)
		return
	}
	var responseMsg message.Message
	var registerRspMsg message.RegisterRspMsg
	responseMsg.Type = message.RegisterRspMsgType
	user, err := model.GlobalUserDao.Register(registerMsg.UserId, registerMsg.UserPwd, registerMsg.UserName)
	if err != nil {
		fmt.Printf("register user fail, err=%v\n", err)
		if err == model.ErrorUserExists {
			registerRspMsg.Code = 421
			registerRspMsg.Msg = err.Error()
		} else {
			registerRspMsg.Code = 500
			registerRspMsg.Msg = "unknown error!"
		}
	} else {
		registerRspMsg.Code = 200
		fmt.Printf("用户%d注册成功\n", user.UserId)
	}
	data, err := json.Marshal(registerRspMsg)
	if err != nil {
		fmt.Printf("marshal register rsp msg fail, err=%v\n", err)
		return err
	}
	responseMsg.Data = string(data)
	data, err = json.Marshal(responseMsg)
	if err != nil {
		fmt.Printf("marshal rsp msg fail, err=%v\n", err)
		return err
	}
	transfer := &utils.Transfer{Conn: userProcess.Conn}
	err = transfer.WritePkg(data)
	return
}

func (userProcess *UserProcess) NofityOthersOnlieUser(userId int, status int) {
	for id, up := range globalUserMgr.onlieUsers {
		if id == userId {
			continue
		}
		up.NotifyStatus(userId, status)
	}
}

func (userProcess *UserProcess) NotifyStatus(userId int, status int) {
	var msg message.Message
	msg.Type = message.UserStatusNofityMsgType
	var notidyStatusMsg message.UserStatusNotifyMsg
	notidyStatusMsg.UserId = userId
	notidyStatusMsg.Status = status
	data, err := json.Marshal(notidyStatusMsg)
	if err != nil {
		fmt.Printf("json marshal fail in NotifyStatus, err=%v\n", err)
		return
	}
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Printf("json marshal fail in NotifyStatus, err=%v\n", err)
		return
	}
	tf := &utils.Transfer{Conn: userProcess.Conn}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Printf("send user status msg fail, err=%v\n", err)
		return
	}
}

func (userProcess *UserProcess) ServerProcessLogout(msg *message.Message) (err error) {
	// 处理退出
	var logoutMes message.LogoutMsg
	err = json.Unmarshal([]byte(msg.Data), &logoutMes)
	if err != nil {
		fmt.Printf("json.Unmarshal([]byte(msg.Data), &logoutMes) fail, err=%v\n", err)
		return err
	}
	var responseMsg message.Message
	responseMsg.Type = message.LogoutRspMsgType
	var logoutRspMsg message.LogoutRspMsg
	userProcess.UserId = logoutMes.UserId
	globalUserMgr.DeleteOnlieUser(logoutMes.UserId)
	userProcess.NofityOthersOnlieUser(logoutMes.UserId, message.UserOffline)
	logoutRspMsg.Code = 200
	data, err := json.Marshal(logoutRspMsg)
	if err != nil {
		fmt.Printf("json.Marshal(loginRspMsg) fail, err=%v\n", err)
		return err
	}
	responseMsg.Data = string(data)
	data, err = json.Marshal(responseMsg)
	if err != nil {
		fmt.Printf("json.Marshal(responseMsg) fail, err=%v\n", err)
		return err
	}
	transfer := &utils.Transfer{Conn: userProcess.Conn}
	err = transfer.WritePkg(data)
	return
}
