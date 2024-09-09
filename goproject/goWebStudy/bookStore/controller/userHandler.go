package controller

import (
	"fmt"
	"goWebStudy/bookStore/dao"
	"goWebStudy/bookStore/model"
	"goWebStudy/bookStore/utils"
	"net/http"
)

// Login
func Login(w http.ResponseWriter, r *http.Request) {
	// 获取用户名和密码
	userName := r.PostFormValue("user_name")
	passWord := r.PostFormValue("pass_word")

	// 验证用户名和密码是否正确
	user, _ := dao.CheckUserNameAndPassWord(userName, passWord)
	if user != nil {
		// 用户名密码正确
		msg := model.Message{Code: 0, Msg: "Login succeed!"}
		data, _ := utils.JsonMarshal(msg)
		w.WriteHeader(200)
		w.Write(data)
	} else {
		// 用户名和密码错误
		msg := model.Message{Code: 1, Msg: "Login failed, check username and password failed!"}
		data, _ := utils.JsonMarshal(msg)
		w.WriteHeader(500)
		w.Write(data)
	}
}

// regsit
func Regsit(w http.ResponseWriter, r *http.Request) {
	// 获取用户名和密码邮箱
	userName := r.PostFormValue("user_name")
	passWord := r.PostFormValue("pass_word")
	EMail := r.PostFormValue("e_mail")

	// 验证用户名时候存在
	user, _ := dao.CheckUserName(userName)
	if user != nil {
		msg := model.Message{Code: 1, Msg: "Regist failed, username alreay exist!"}
		data, _ := utils.JsonMarshal(msg)
		w.WriteHeader(400)
		w.Write(data)
		return
	}

	// 用户名不存在写入db注册成功
	user = &model.User{UserName: userName, PassWord: passWord, EMail: EMail}
	err := dao.SaveUser(user)
	if err != nil {
		msg := model.Message{Code: 1, Msg: "Regist failed!"}
		data, _ := utils.JsonMarshal(msg)
		w.WriteHeader(500)
		w.Write(data)
	} else {
		msg := model.Message{Code: 0, Msg: "Regist succeed!"}
		data, _ := utils.JsonMarshal(msg)
		w.WriteHeader(200)
		w.Write(data)
	}
}

// check username
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	// 获取用户名
	userName := r.PostFormValue("user_name")

	// 验证用户名是否存在
	user, _ := dao.CheckUserName(userName)
	if user != nil {
		msg := model.Message{Code: 1, Msg: fmt.Sprintf("%s is exist", userName)}
		data, _ := utils.JsonMarshal(msg)
		w.WriteHeader(400)
		w.Write(data)
	}
}
