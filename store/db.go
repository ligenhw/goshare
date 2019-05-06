package store

import (
	"database/sql"
	"path"
	"runtime"

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

	InitDb(getSqlFile(), false)
}

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)

	return path.Dir(filename)
}

func getSqlFile() string {
	currentPath := getCurrentPath()

	return currentPath + "/schema.sql"
}
