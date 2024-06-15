package process

import (
	"chatRoom/common/message"
	"fmt"
)

type User struct {
	UserId     int    `json:"user_id"`
	UserPwd    string `json:"user_pwd"`
	UserName   string `json:"user_name"`
	UserStatus int    `json:"user_status"`
}

var onlineUserMap map[int]*User = make(map[int]*User, 1024)

func updateUserStatus(notifyUserStatusMsg *message.UserStatusNotifyMsg) {
	user, ok := onlineUserMap[notifyUserStatusMsg.UserId]
	if !ok {
		onlineUserMap[notifyUserStatusMsg.UserId] = &User{
			UserId:     notifyUserStatusMsg.UserId,
			UserStatus: notifyUserStatusMsg.Status,
		}
	} else {
		user.UserStatus = message.UserOnlie
		onlineUserMap[notifyUserStatusMsg.UserId] = user
	}
	outputOnlieUser()
}

func outputOnlieUser() {
	fmt.Println("当前在线用户列表:")
	for id, _ := range onlineUserMap {
		fmt.Printf("id: %d\n", id)
	}
}
