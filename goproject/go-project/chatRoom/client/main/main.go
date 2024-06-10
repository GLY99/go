package main

import (
	"chatRoom/client/process"
	"fmt"
	"os"
)

func mainMenu() int {
	var num int
	var loop bool = true
	for {
		fmt.Printf("----------欢迎登陆多人聊天室----------\n")
		fmt.Printf("          1.登录聊天室\n")
		fmt.Printf("          2.注册用户\n")
		fmt.Printf("          3.退出系统\n")
		fmt.Printf("请选择1-3:")
		_, err := fmt.Scanf("%d\n", &num)
		if err != nil {
			fmt.Println("输入错误,退出系统!!!")
			loop = false
			continue
		}
		switch num {
		case 1:
			loop = false
		case 2:
			loop = false
		case 3:
			loop = false
		default:
			fmt.Println("输入有误,请重新输入")
		}
		if !loop {
			break
		}
	}
	return num
}

func getYOrN() bool {
	count := 3
	var choose string
	for i := 1; i <= count; i++ {
		fmt.Printf("请输入Y or N:")
		_, err := fmt.Scanf("%s\n", &choose)
		if err != nil {
			fmt.Printf("err=%v", err)
			continue
		}
		if choose != "Y" && choose != "N" {
			continue
		} else {
			if choose == "Y" {
				return true
			} else {
				return false
			}
		}
	}
	return true
}

func loginMenu() {
	var userId int
	var passWd string
	loop := true
	for {
		fmt.Printf("请输入用户账号:")
		_, err1 := fmt.Scanf("%d\n", &userId)
		fmt.Printf("请输入用户密码:")
		_, err2 := fmt.Scanf("%s\n", &passWd)
		if err1 != nil || err2 != nil {
			fmt.Println("账号或者密码错误请重新输入,是否重新登录?")
			if getYOrN() {
				continue
			} else {
				num := mainMenu()
				if num == 1 {
					continue
				} else if num == 2 {
					registerMenu()
				} else if num == 3 {
					exitSys()
				}
			}
		}

		userProcess := &process.UserProcess{}
		err := userProcess.Login(userId, passWd)
		if err != nil {
			fmt.Println("登录失败,是否重新登录?")
			if getYOrN() {
				continue
			} else {
				num := mainMenu()
				if num == 1 {
					continue
				} else if num == 2 {
					registerMenu()
				} else if num == 3 {
					exitSys()
				}
			}
		} else {
			loop = false
		}
		if !loop {
			break
		}
	}
}

func registerMenu() {
	fmt.Println("注册用户")
	var userId int
	var userName string
	var userPwd string
	loop := true
	for {
		fmt.Printf("请输入要注册的用户id:")
		_, err1 := fmt.Scanf("%d\n", &userId)
		fmt.Printf("请输入要注册的用户密码:")
		_, err2 := fmt.Scanf("%s\n", &userPwd)
		fmt.Printf("请输入要注册的用户昵称:")
		_, err3 := fmt.Scanf("%s\n", &userName)
		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("填写注册信息失败,是否重新填写？")
			if getYOrN() {
				continue
			} else {
				num := mainMenu()
				if num == 1 {
					loginMenu()
				} else if num == 2 {
					continue
				} else if num == 3 {
					exitSys()
				}
			}
		}
		userProcess := &process.UserProcess{}
		err := userProcess.Register(userId, userPwd, userName)
		if err != nil {
			fmt.Printf("register user info fail, err=%v\n", err)
			fmt.Printf("是否重新注册？")
			if getYOrN() {
				continue
			} else {
				num := mainMenu()
				if num == 1 {
					continue
				} else if num == 2 {
					registerMenu()
				} else if num == 3 {
					exitSys()
				}
			}
		} else {
			loop = false
		}
		if !loop {
			break
		}
	}
}

func exitSys() {
	os.Exit(1)
}

func main() {
	for {
		num := mainMenu()
		if num == 1 {
			loginMenu()
		} else if num == 2 {
			registerMenu()
		} else if num == 3 {
			exitSys()
		}
	}
}
