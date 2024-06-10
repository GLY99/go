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
		fmt.Printf("json.Unmarshal([]byte(res), user) fail, err=%v\n", err)
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

func (UserDao *UserDao) Register(userId int, userPwd, userName string) (user *User, err error) {
	conn := UserDao.redisPool.Get()
	defer conn.Close()
	user, err = UserDao.getUserById(conn, userId)
	if err == nil {
		err = ErrorUserExists
		return
	}
	user = &User{UserId: userId, UserPwd: userPwd, UserName: userName}
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("marshal user info fail when register user info to redis, err=%v\n", err)
		return
	}
	_, err = conn.Do("HSet", "users", fmt.Sprintf("%d", userId), string(data))
	if err != nil {
		fmt.Printf("register user info to redis fail, err=%v\n", err)
		return
	}
	return
}
