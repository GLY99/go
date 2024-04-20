package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func redisDial() (redis.Conn, error) {
	return redis.Dial("tcp", "127.0.0.1:6379")
}

func main() {
	pool := &redis.Pool{
		Dial:        redisDial,
		MaxIdle:     10, // 最大空闲连接数
		MaxActive:   0,  // 和数据库的最大连接数，0是无限制
		IdleTimeout: 10, // 最大空闲时间
	}
	defer pool.Close() // 关闭连接池
	c := pool.Get()    // 从连接池中取出连接
	_, err := c.Do("set", "k1", "v1")
	if err != nil {
		fmt.Printf("set k1 v1 fail, err=%v\n", err)
	}
	v, err := redis.String(c.Do("get", "k1"))
	if err != nil {
		fmt.Printf("get k1 fail, err=%v\n", err)
	}
	fmt.Printf("get k1 succ, val=%s\n", v) // get k1 succ, val=v1
}
