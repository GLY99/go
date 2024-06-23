package process

import (
	"chatRoom/client/utils"
	"chatRoom/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct{}

func (smsProcess *SmsProcess) SendGroupMsg(content string) (err error) {
	msg := message.Message{Type: message.SmsMsgType}
	smsMsg := message.SmsMsg{Content: content, User: CurUser.User}

	data, err := json.Marshal(smsMsg)
	if err != nil {
		fmt.Printf("json Marshal smsMsg fail in SendGroupMsg, err=%v\n", err)
		return
	}

	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Printf("json Marshal msg fail in SendGroupMsg. err=%v\n", err)
		return
	}

	tf := utils.Transfer{Conn: CurUser.Conn}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Printf("tf WritePkg fail in SendGroupMsg, err=%v\n", err)
		return
	}
	fmt.Printf("user %d send msg %s successfully\n", CurUser.UserId, content)
	return
}
