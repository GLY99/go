package process

import (
	"chatRoom/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(msg *message.Message) {
	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Printf("json unmarshal msg fail in SendGroupMsg, err=%v\n", err)
		return
	}
	content := smsMsg.Content
	userId := smsMsg.UserId
	info := fmt.Sprintf("用户\t%d 对大家说\t%s", userId, content)
	fmt.Println(info)
	fmt.Println()
}
