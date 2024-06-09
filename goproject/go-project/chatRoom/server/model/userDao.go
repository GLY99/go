package model

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// 在服务器启动后就初始化一个userDao
var (
	GlobalUserDao *UserDao
)

// 完成user这个结构体和redis交互
type UserDao struct {
	redisPool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{redisPool: pool}
	return
}

func (userDao *UserDao) getUserById(conn redis.Conn, userId int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", userId))
	if err == redis.ErrNil {
		err = ErrorUserNotExists
		return
	}
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Printf("json.Unmarshal([]byte(res), user) fail, err=%v", err)
		return
	}
	return
}

func (userDao *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := userDao.redisPool.Get()
	defer conn.Close()
	user, err = userDao.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ErrorUserPwd
		return
	}
	return
}
