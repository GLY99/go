package message

const (
	LoginMsgType    = "LoginMsgType"
	LoginRspMsgType = "LoginRspMsgType"
	RegisterMsgType = "RegisterMsgType"
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
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type RegisterMsg struct {
	// 注册
}
