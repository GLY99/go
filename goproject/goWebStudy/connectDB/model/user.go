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

// GetUserById 通过id获取User方法
func (user *User) GetUserById() (*User, error) {
	sqlStr := "select id, username, password, e_mail from users where id=?"
	row := utils.DB.QueryRow(sqlStr, user.Id)
	var userName string
	var passWord string
	var eMail string
	var id int
	// 这里变量的字段顺序要和sql结果里面字段顺序一致
	err := row.Scan(&id, &userName, &passWord, &eMail)
	if err != nil {
		return nil, err
	}
	u := &User{
		Id:       id,
		Username: userName,
		Password: passWord,
		EMail:    eMail,
	}
	return u, nil
}

// GetUsers 获取所有User记录
func (user *User) GetUsers() ([]*User, error) {
	sqlStr := "select id, username, password, e_mail from users"
	rows, err := utils.DB.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	users := make([]*User, 0)
	var userName string
	var passWord string
	var eMail string
	var id int
	// 判断有没有下一行，如果有移动游标到下一行
	for rows.Next() {
		err := rows.Scan(&id, &userName, &passWord, &eMail)
		if err != nil {
			return users, err
		}
		u := &User{
			Id:       id,
			Username: userName,
			Password: passWord,
			EMail:    eMail,
		}
		users = append(users, u)
	}
	return users, nil
}
