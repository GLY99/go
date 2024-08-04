package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Printf("sql open err, err=%v", err)
	}
	fmt.Printf("sql open succeed\n")
}
