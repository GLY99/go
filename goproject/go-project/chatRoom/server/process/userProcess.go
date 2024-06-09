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
	Conn net.Conn
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
