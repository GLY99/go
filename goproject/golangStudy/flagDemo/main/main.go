package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var password string
	var host string
	var port int
	flag.StringVar(&user, "u", "", "user name default is \"\"")
	flag.StringVar(&password, "p", "", "password default is \"\"")
	flag.StringVar(&host, "h", "localhost", "host default is localhost")
	flag.IntVar(&port, "P", 3306, "port default is 3306")
	// 从os.Args[1:]中解析参数,必须调用这个方法
	flag.Parse()

	fmt.Printf("user=%s,password=%s,host=%s,port=%d\n", user, password, host, port)
}
