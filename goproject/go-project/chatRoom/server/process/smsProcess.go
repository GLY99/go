package process

import (
	"chatRoom/client/utils"
	"chatRoom/common/message"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct{}

func (smsProcess *SmsProcess) SendGroupMsg(msg *message.Message) (err error) {
	var smsMsg message.SmsMsg
	err = json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Printf("json unmarshal msg fail in SendGroupMsg, err=%v\n", err)
		return
	}
	user := smsMsg.User
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("json marshal msg fail in SendGroupMsg, err=%v\n", err)
		return
	}
	for userId, up := range globalUserMgr.onlieUsers {
		if userId == user.UserId {
			continue
		}
		smsProcess.SendMsgToEachOnlieUser(data, up.Conn)
	}
	return
}

func (smsProcess *SmsProcess) SendMsgToEachOnlieUser(data []byte, conn net.Conn) {
	tf := utils.Transfer{Conn: conn}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Printf("send msg to each onlie user fail, err=%v\n", err)
		return
	}
	return
}
