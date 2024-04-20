package main

import (
	"fmt"

	"gopkg.in/redis.v4"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Printf("conn redis fail, err=%v\n", err)
	}
	fmt.Println(pong, err) // PONG <nil>

	result1 := client.Set("name", "tom", 0) // 0表示不过期
	if result1.Err() != nil {
		fmt.Printf("set name fail, err=%v\n", result1.Err())
	}
	result2 := client.Get("name")
	if result2.Err() != nil {
		fmt.Printf("get name fail, err=%v\n", result2.Err())
	}
	fmt.Println(result2.Val()) // tom
}
