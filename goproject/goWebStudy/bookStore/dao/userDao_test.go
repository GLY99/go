package dao

import (
	"fmt"
	"goWebStudy/bookStore/model"
	"testing"
)

func TestMain(m *testing.M) {
	// 在TestMain中可以为测试做一些前置操作
	m.Run() // 如果不写这个下面的测试方法不会执行
}

func TestUserDao(t *testing.T) {
	t.Run("test SaveUser func", testSaveUser)
	t.Run("test CheckUserNameAndPassWord func", testCheckUserNameAndPassWord)
	t.Run("test CheckUserName func", testCheckUserName)
}

func testSaveUser(t *testing.T) {
	user := &model.User{
		UserName: "GGbond",
		PassWord: "123456",
		EMail:    "GGbond@163.com",
	}
	SaveUser(user)
	t.Log("testSaveUser succeed")
}

func testCheckUserNameAndPassWord(t *testing.T) {
	user, _ := CheckUserNameAndPassWord("GGbond", "123456")
	fmt.Println(user)
	user, _ = CheckUserNameAndPassWord("GGbond", "111111")
	fmt.Println(user)
	t.Log("testCheckUserNameAndPassWord succeed")
}

func testCheckUserName(t *testing.T) {
	user, _ := CheckUserName("GGbond")
	t.Log(user)
	user, _ = CheckUserName("GGbond1")
	t.Log(user)
	t.Logf("testCheckUserName succeed")
}
