package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

// 包被导入时，执行完成全局变量的一些代码块后执行
func init() {
	var err error
	DB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/bookstore")
	if err != nil {
		fmt.Printf("sql.Open failed, err=%v\n", err)
		return
	}
	fmt.Printf("sql.Open succeed\n")
}
