package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ligenhw/goshare/configration"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", configration.Conf.Dsn)
	if err != nil {
		panic(err)
	}
}
