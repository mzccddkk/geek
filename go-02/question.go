// 1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做？
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type User struct {
	Id   int64
	Name string
}

var Db *sql.DB

// Custom Error
var CustomizeNotFound = errors.New("resource not found")

func init() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4"
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = Db.Ping()
	if err != nil {
		panic(err)
	}
}

// DAO
func QueryUserById(id int64) (User, error) {
	var user User
	s := fmt.Sprintf("select id, name from user where id = %d", id)
	row := Db.QueryRow(s)
	err := row.Scan(&user.Id, &user.Name)
	if err != nil {
		return user, errors.Wrapf(CustomizeNotFound, "dao#QueryUserById#err=%v#sql=%v", err, s)
	}
	return user, nil
}

func main() {
	user, err := QueryUserById(1111)
	if errors.Is(err, CustomizeNotFound) {
		fmt.Printf("err: %+v", err)
		return
	}
	fmt.Printf("user: %v", user)
}
