package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*type student struct {
	name string `db:name`
	age int `db:age`
}*/

func main(){
//connect
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/go_test?charset=utf8")
	if err!=nil {
		fmt.Println(err)

	}

	//insert
	res, err := db.Exec("insert into student (name,age) values ('qian',20) ")
	if err!=nil {
		fmt.Println(err)
		return
	}
	id, err := res.LastInsertId()
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println("insert sucess ", id)

//query
	rows, err := db.Query("select * from student")
	if err != nil {
		fmt.Println("query failed1",err)
		return
	}

	for rows.Next() {
		var n string
		var age int
		err = rows.Scan(&n,&age)
		if err != nil {
			fmt.Println("query failed2",err)
			return
		}
		fmt.Println(n,age)
	}
//delete
	_, err = db.Exec("delete from student where name='qian'")
	if err != nil {
		fmt.Println("delete failed",err)
		return
	}

//update
	_, err = db.Exec("update student set age=18 where name='li'")
	if err != nil {
		fmt.Println("update failed",err)
		return
	}

}
