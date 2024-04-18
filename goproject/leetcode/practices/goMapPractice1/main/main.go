package main

import "fmt"

func modifyUser(users map[string]map[string]string, user string) {
	if users[user] != nil {
		users[user]["password"] = "888888"
	} else {
		users[user] = make(map[string]string, 2)
		users[user]["name"] = user
		users[user]["password"] = "000000"
	}
}

func main() {
	users := make(map[string]map[string]string, 1)
	modifyUser(users, "Tom")
	modifyUser(users, "jerry")
	fmt.Println(users) // map[Tom:map[nickname:Tom password:000000] jerry:map[nickname:jerry password:000000]]
	modifyUser(users, "Tom")
	fmt.Println(users) // map[Tom:map[nickname:Tom password:888888] jerry:map[nickname:jerry password:000000]]
}
