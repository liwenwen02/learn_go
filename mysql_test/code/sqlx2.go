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


	//db.Get返回一个结构体
		var sql_query_one string = "select * from student where name='li'"
		var s Stu
		err  = db.Get(&s,sql_query_one)
		if err != nil {
			fmt.Println("query_one failed", err)
			return
		}
		fmt.Println(s)

}
