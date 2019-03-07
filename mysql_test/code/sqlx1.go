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

	//insert返回的是sql.Result，相关函数有LastInserId、RowsAffected
	sql_insert := "insert into student (name,age) values ('qian',20)"
	res, err := db.Exec(sql_insert)
	id, _ := res.LastInsertId()
	ra, _ := res.RowsAffected()
	fmt.Println("id is ", id, ", affected rows ", ra)
//select返回的是sql.Rows
	sql_select := "select * from student"
	rows, err := db.Query(sql_select)
	if err != nil {
		fmt.Println("query failed 01", err)
		return
	}
	for rows.Next(){
		var n string
		var a int
		err = rows.Scan(&n,&a)//scan会自动推断部分目标变量
		if err != nil {
			fmt.Println("query rows failed",err)
			return
		}
		fmt.Println(n, a)
	}
	rows.Close()
	err = rows.Err()
	if err != nil {
		fmt.Println("rows error ", err)
		return
	}//防止遍历中出现错误，rows的连接未释放


//db.Get返回一个结构体
/*	var sql_query_one string = "select name,age from student where name='li'"
	var s stu
	err  = db.Get(&s, sql_query_one)
	if err != nil {
		fmt.Println("query_one failed", err)
		return
	}
	fmt.Println(s)
	*/


	//db.select返回多个结构体
/*	var s_slice []stu
	err = db.Select(&s_slice, sql_select)
	if err != nil {
		fmt.Println("query falied 02", err)
		return
	}
	for _, v := range s_slice {
		fmt.Println(v)
	}*/

	//




}
