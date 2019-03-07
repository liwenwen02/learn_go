package main
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
	"database/sql"
)
func clearTransaction(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil {
		fmt.Println(err)
		return
	}
}
//事务的连接生命周期从Begin函数调用起，直到Commit和Rollback函数的调用结束
func main () {
	db, err := sqlx.Open("mysql","root:@tcp(127.0.0.1:3306)/go_test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer clearTransaction(tx)
	_, err = tx.Exec("update student set age = 40 where name='li'")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}

}


