package model

import (
	"fmt"
	"goWebStudy/connectDB/utils"
)

// User 结构体
type User struct {
	Id       int
	Username string
	Password string
	EMail    string
}

// AddUser 添加User方法，带预编译
func (user *User) AddUser() error {
	// 插入的sql
	sqlStr := "insert into users (username, password, e_mail) values (?, ?, ?)"
	// 预编译
	inStmt, err := utils.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("utils.DB.Prepare failed in AddUser, err=%v", err)
		return err
	}
	// 执行
	_, err = inStmt.Exec(user.Username, user.Password, user.EMail)
	if err != nil {
		fmt.Printf("inStmt.Exec failed in AddUser, err=%v", err)
		return err
	}
	return nil
}

// AddUser2 添加User方法，不带预编译
func (user *User) AddUser2() error {
	// 插入的sql
	sqlStr := "insert into users (username, password, e_mail) values (?, ?, ?)"

	// 执行
	_, err := utils.DB.Exec(sqlStr, user.Username, user.Password, user.EMail)
	if err != nil {
		fmt.Printf("utils.DB.Exec failed in AddUser2, err=%v", err)
		return err
	}
	return nil
}
