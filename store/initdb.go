package store

import (
	"io/ioutil"
	"log"
	"strings"
)

func InitDb(file string, drop bool) (err error) {
	var bs []byte
	bs, err = ioutil.ReadFile(file)
	if err != nil {
		return
	}

	if drop {
		dropDb()
	}

	if !dbExits("goshare") {
		log.Println("db : goshare not exist, creating ...")
		createDb(string(bs))
	} else {
		log.Println("db : goshare exist, skip db init ...")
	}
	return
}

func dropDb() {
	sql := "DROP DATABASE test1"
	_, err := Db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

func dbExits(db string) bool {
	sql := "SHOW DATABASES"
	r, err := Db.Query(sql)
	if err != nil {
		return false
	}

	for r.Next() {
		var dbName string
		if err = r.Scan(&dbName); err != nil {
			return false
		}
		if dbName == db {
			return true
		}
	}

	return false
}

func createDb(sql string) {
	stats := strings.Split(sql, ";")
	for _, stat := range stats {
		if stat != "" {
			if _, err := Db.Exec(stat); err != nil {
				log.Fatal(err)
			}
		}
	}
}
