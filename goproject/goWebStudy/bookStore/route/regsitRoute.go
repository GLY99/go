package route

import (
	"goWebStudy/bookStore/controller"
	"net/http"
)

func RegistRoute(serverMux *http.ServeMux) {
	// 注册路由
	// 登录
	serverMux.HandleFunc("/login", controller.Login)
	// 注册
	serverMux.HandleFunc("/regist", controller.Regsit)
	// 检查用户名
	serverMux.HandleFunc("/check_user_name", controller.CheckUserName)
}
