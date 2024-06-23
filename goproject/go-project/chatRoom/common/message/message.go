package message

const (
	LoginMsgType            = "LoginMsgType"
	LoginRspMsgType         = "LoginRspMsgType"
	RegisterMsgType         = "RegisterMsgType"
	RegisterRspMsgType      = "RegisterRspMsgType"
	UserStatusNofityMsgType = "UserStatusNofityMsgType"
	SmsMsgType              = "SmsMsgType"
	LogoutMsgType           = "LogoutMsgType"
	LogoutRspMsgType        = "LogoutRspMsgType"
)

const (
	UserOnlie = iota
	UserOffline
	UserBusy
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMsg struct {
	UserId   int    `json:"user_id"`
	UserPwd  string `json:"user_pwd"`
	UserName string `json:"user_name"`
}

type LoginRspMsg struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	UserIds []int  `json:"user_ids"`
}

type RegisterMsg struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	UserPwd  string `json:"user_pwd"`
}

type RegisterRspMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type UserStatusNotifyMsg struct {
	UserId int `json:"user_id"`
	Status int `json:"status"`
}

type User struct {
	UserId     int    `json:"user_id"`
	UserPwd    string `json:"user_pwd"`
	UserName   string `json:"user_name"`
	UserStatus int    `json:"user_status"`
}

type SmsMsg struct {
	Content string `json:"content"`
	User
}

type LogoutMsg struct {
	User
}

type LogoutRspMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
