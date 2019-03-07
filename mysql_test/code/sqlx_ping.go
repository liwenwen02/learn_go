package main

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sqlx.Open("mysql","root:@tcp(127.0.0.1:3306)/go_test?charset=utf8")
	//占位符？
	if err != nil {
		fmt.Println("connect failed", err)
		return
	}
	defer db.Close()
	fmt.Println("connect sucess")
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ping sucess")
//调用ping,连接池会初始化一个连接





}
