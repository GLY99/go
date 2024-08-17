package dao

import (
	"goWebStudy/bookStore/model"
	"goWebStudy/bookStore/utils"
)

// 验证用户名和密码
func CheckUserNameAndPassWord(userName, passWord string) (*model.User, error) {
	sqlStr := "select id, user_name, pass_word, e_mail from users where user_name=? and pass_word=?"
	row := utils.DB.QueryRow(sqlStr, userName, passWord)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.UserName, &user.PassWord, &user.EMail)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 验证用户名
func CheckUserName(userName string) (*model.User, error) {
	sqlStr := "select id, user_name, pass_word, e_mail from users where user_name=?"
	row := utils.DB.QueryRow(sqlStr, userName)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.UserName, &user.PassWord, &user.EMail)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 保存用户
func SaveUser(user *model.User) error {
	sqlStr := "insert into users (user_name, pass_word, e_mail) values (?,?,?)"
	_, err := utils.DB.Exec(sqlStr, user.UserName, user.PassWord, user.EMail)
	return err
}
