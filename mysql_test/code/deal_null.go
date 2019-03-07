package main

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)
//有一行记录的值为空值
func main() {
	db, err := sqlx.Open("mysql","root:@tcp(127.0.0.1:3306)/go_test?charset=utf8")
	//占位符？
	if err != nil {
		fmt.Println("connect failed", err)
		return
	}
	defer db.Close()
	fmt.Println("connect sucess")

	var name string
	var age sql.NullInt64 //没有sql.NullInt64,会报错
	//报错converting driver.Value type <nil> ("<nil>") to a int
	//var age []byte  //可以用byte切片接收，空值则为空
	rows := db.QueryRow("select * from student where name='yang'")
	err = rows.Scan(&name, &age)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name, age)

	if age.Valid {
		fmt.Println("res:", age.Int64)
	} else {
		fmt.Println("err", age.Int64)
	}







}
