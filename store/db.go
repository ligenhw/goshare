package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "gen:1234@tcp(192.168.199.108)/goshare?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
}
