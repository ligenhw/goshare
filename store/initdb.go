package store

import (
	"io/ioutil"
	"log"
	"strings"
)

const dbName = "goshare"

func InitDb(file string, drop bool) (err error) {
	var bs []byte
	bs, err = ioutil.ReadFile(file)
	if err != nil {
		return
	}

	if drop {
		dropDb()
	}

	if !dbExits(dbName) {
		log.Println("db : goshare not exist, creating ...")
		createDb(string(bs))
	} else {
		log.Println("db : goshare exist, skip db use db :", dbName)
		useDb(dbName)
	}
	return
}

func dropDb() {
	sql := "DROP DATABASE goshare"
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
		if strings.TrimSpace(stat) != "" {
			log.Println("exec : ", stat)
			if _, err := Db.Exec(stat); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func useDb(dbName string) {
	sql := "USE " + dbName
	_, err := Db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}
