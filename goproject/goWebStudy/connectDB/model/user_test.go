package model

import (
	"testing"
)

func TestMain(m *testing.M) {
	// 在TestMain中可以为测试做一些前置操作
	m.Run() // 如果不写这个下面的测试方法不会执行
}

func TestUser(t *testing.T) {
	t.Run("test AddUser func", testAddUser)
	t.Run("test Adduser1 func", testAddUser1)
	t.Run("test GetUserById func", testGetUserById)
	t.Run("test GetUsers func", testGetUsers)
}

func testAddUser(t *testing.T) {
	user1 := User{
		Username: "jack",
		Password: "123456",
	}
	err := user1.AddUser()
	if err != nil {
		t.Log("add user1 failed\n")
	}
}

func testAddUser1(t *testing.T) {
	user2 := User{
		Username: "tom",
		Password: "123456",
		EMail:    "111@163.com",
	}
	err := user2.AddUser2()
	if err != nil {
		t.Log("add user2 failed\n")
	}
}

func testGetUserById(t *testing.T) {
	user := &User{Id: 1}
	u, err := user.GetUserById()
	if err != nil {
		t.Logf("get user by id failed, err=%v", err)
	} else {
		t.Log(*u)
	}
}

func testGetUsers(t *testing.T) {
	user := &User{}
	users, err := user.GetUsers()
	if err != nil {
		t.Logf("get users failed, err=%v", err)
	} else {
		for _, u := range users {
			t.Log(*u)
		}
	}
}
