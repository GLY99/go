package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	// 连接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Printf("conn redis fail, err=%v\n", err)
		return
	}
	defer conn.Close()

	// 向redis写入数据string key-val
	_, err = conn.Do("Set", "name", "tom")
	if err != nil {
		fmt.Printf("Set fail, err=%v\n", err)
		return
	}
	// 从redis读取写入的string,读取回来会进行一次类型转换
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Printf("Get fail, err=%v\n", err)
		return
	}
	// 下面这个类型转换会报错
	// r = r.(string) panic: interface conversion: interface {} is []uint8, not string
	fmt.Printf("get name succ, val=%v\n", r) //get name succ, val=tom

	// 向redis写入hash
	_, err = conn.Do("HSet", "user1", "name", "tom")
	if err != nil {
		fmt.Printf("HSet fail, err=%v\n", err)
		return
	}
	_, err = conn.Do("HMset", "user1", "age", 18, "job", "yaofan")
	if err != nil {
		fmt.Printf("HMset fail, err=%v\n", err)
		return
	}
	// 从redis读hash
	r, err = redis.String(conn.Do("HGet", "user1", "name"))
	if err != nil {
		fmt.Printf("Hget fail, err=%v\n", err) // HGet succ, val=tom
		return
	}
	fmt.Printf("HGet succ, val=%v\n", r)
	r_slice, err := redis.Strings(conn.Do("HMget", "user1", "name", "age", "job"))
	if err != nil {
		fmt.Printf("HMget fail, err=%v\n", err)
		return
	}
	fmt.Printf("HMGet succ, val=%v\n", r_slice) // HMGet succ, val=[tom 18 yaofan]

	// 给数据设置有效时间
	conn.Do("Set", "test1", "test1Val")
	conn.Do("expire", "test1", 3)
	r, _ = redis.String(conn.Do("Get", "test1"))
	fmt.Printf("before 3s, val=%s\n", r) // before 10s, val=test1Val
	time.Sleep(5 * time.Second)
	r, _ = redis.String(conn.Do("Get", "test1"))
	fmt.Printf("after 3s, val=%s\n", r) // after 10s, val=

	// 向redis写入list
	_, err = conn.Do("Rpush", "names", "tom", "jack")
	if err != nil {
		fmt.Printf("Rpush fail, err=%v\n", err)
		return
	}
	// 读取list
	r_slice, err = redis.Strings(conn.Do("Lrange", "names", 0, -1))
	if err != nil {
		fmt.Printf("Lrange fail, err=%v\n", err)
		return
	}
	fmt.Printf("Lrange succ, val=%v\n", r_slice) // Lrange succ, val=[tom jack]
}
